// Copyright 2017 The Goma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

/*
Binary exec_server provides goma exec service via gRPC.

*/
package main

import (
	"context"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"path"
	"strings"
	"time"

	rpb "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	gax "github.com/googleapis/gax-go/v2"
	"github.com/googleapis/google-cloud-go-testing/storage/stiface"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
	"go.opencensus.io/zpages"
	"google.golang.org/api/option"
	bspb "google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/encoding/prototext"

	"go.chromium.org/goma/server/cache/redis"
	"go.chromium.org/goma/server/command"
	"go.chromium.org/goma/server/exec"
	"go.chromium.org/goma/server/file"
	"go.chromium.org/goma/server/log"
	"go.chromium.org/goma/server/log/errorreporter"
	"go.chromium.org/goma/server/profiler"
	cmdpb "go.chromium.org/goma/server/proto/command"
	pb "go.chromium.org/goma/server/proto/exec"
	filepb "go.chromium.org/goma/server/proto/file"
	"go.chromium.org/goma/server/remoteexec"
	"go.chromium.org/goma/server/remoteexec/digest"
	"go.chromium.org/goma/server/rpc"
	"go.chromium.org/goma/server/server"
)

var (
	port                  = flag.Int("port", 5050, "rpc port")
	mport                 = flag.Int("mport", 8081, "monitor port")
	fileAddr              = flag.String("file-addr", "passthrough:///file-server:5050", "file server address")
	configMapURI          = flag.String("configmap_uri", "", "deprecated: configmap uri. e.g. gs://$project-toolchain-config/$name.config, text proto of command.ConfigMap.")
	configMap             = flag.String("configmap", "", "configmap text proto")
	toolchainConfigBucket = flag.String("toolchain-config-bucket", "", "cloud storage bucket for toolchain config")
	configMapFile         = flag.String("configmap_file", "", "filename for configmap text proto")

	traceProjectID     = flag.String("trace-project-id", "", "project id for cloud tracing")
	pubsubProjectID    = flag.String("pubsub-project-id", "", "project id for pubsub")
	serviceAccountFile = flag.String("service-account-file", "", "service account json file")

	remoteexecAddr         = flag.String("remoteexec-addr", "", "use remoteexec API endpoint")
	remoteInstancePrefix   = flag.String("remote-instance-prefix", "", "remote instance name path prefix.")
	remoteInstanceBaseName = flag.String("remote-instance-basename", "default_instance", "remote instance basename under remote-instance-prefix")

	// http://b/141901653
	execMaxRetryCount = flag.Int("exec-max-retry-count", 5, "max retry count for exec call. 0 is unlimited count, but bound to ctx timtout. Use small number for powerful clients to run local fallback quickly. Use large number for powerless clients to use remote more than local.")
	execActionTimeout = flag.Duration("exec-action-timeout", 15*time.Minute, "action timeout after which the execution should be killed.")

	cmdFilesBucket      = flag.String("cmd-files-bucket", "", "cloud storage bucket for command binary files")
	fetchConfigParallel = flag.Bool("fetch-config-parallel", true, "fetch toolchain configs in parallel")

	// Needed for b/120582303, but will be deprecated by b/80508682.
	fileLookupConcurrency = flag.Int("file-lookup-concurrency", 20, "concurrency to look up files from file-server")

	// chromium code as of July 2020 (*.c*, *.h) = 230k
	// also chromium clobber bulids has ~60k gomacc invocation.
	// thinlto would upload *.o and *.thinlto.
	// rbe-staging1 uses 2.2M keys (< 512MB memory usage in redis).
	maxDigestCacheEntries = flag.Int("max-digest-cache-entries", 2e6, "maximum entries in in-memory digest cache. 0 means unimited")

	// nsjail is applied in hardened request.
	// note windows and chroot reqs are out of scope for the ratio.
	// e.g.
	//   hadening-ratio=0
	//   nsjail-rario=<any>
	//   => no hardening (no runsc nor nsjail) at all
	//
	//   hardening-ratio=1
	//   nsjail-ratio=0
	//   => hardening by runsc only
	//
	//   hardening-ratio=1
	//   nsjail-ratio=1
	//   => hardening by nsjail only
	//
	//   hardening-ratio=0.5
	//   nsjail-ratio=0.5
	//   => no hardeing 50%
	//      hardening 50%
	//        nsjail 25%  (50% in hardening)
	//        runsc 25%   (50% in hardening)
	experimentHardeningRatio = flag.Float64("experiment-hardening-ratio", 0, "Ratio [0,1] to enable hardening. 0=no hardening. 1=all hardening.")
	experimentNsjailRatio    = flag.Float64("experiment-nsjail-ratio", 0, "Ratio [0,1] to use nsjail for hardening. 0=no nsjial (ie. runsc), 1=all nsjail.")
	disableHardenings        = flag.String("disable-hardenings", "", "comma separated sha256 file hashes of command to disable hardening (i.e. for ELF-32)")

	redisMaxIdleConns   = flag.Int("redis-max-idle-conns", redis.DefaultMaxIdleConns, "maximum number of idle connections to redis.")
	redisMaxActiveConns = flag.Int("redis-max-active-conns", redis.DefaultMaxActiveConns, "maximum number of active connections to redis.")
)

var (
	configUpdate = stats.Int64("go.chromium.org/goma/server/cmd/exec_server.toolchain-config-updates", "toolchain-config updates", stats.UnitDimensionless)

	configStatusKey = tag.MustNewKey("status")

	configViews = []*view.View{
		{
			Description: "counts toolchain-config updates",
			TagKeys: []tag.Key{
				configStatusKey,
			},
			Measure:     configUpdate,
			Aggregation: view.Count(),
		},
	}
)

func recordConfigUpdate(ctx context.Context, err error) {
	logger := log.FromContext(ctx)
	status := "success"
	if err != nil {
		status = "failure"
	}
	ctx, cerr := tag.New(ctx, tag.Upsert(configStatusKey, status))
	if cerr != nil {
		logger.Fatal(cerr)
	}
	stats.Record(ctx, configUpdate.M(1))
	if err != nil {
		server.Flush()
	}
}

func configMapToConfigResp(ctx context.Context, cm *cmdpb.ConfigMap) *cmdpb.ConfigResp {
	resp := &cmdpb.ConfigResp{
		VersionId: time.Now().UTC().Format(time.RFC3339),
	}
	for _, rt := range cm.Runtimes {
		c := &cmdpb.Config{
			Target: &cmdpb.Target{
				Addr: *remoteexecAddr,
			},
			BuildInfo:          &cmdpb.BuildInfo{},
			RemoteexecPlatform: &cmdpb.RemoteexecPlatform{},
			Dimensions:         rt.PlatformRuntimeConfig.Dimensions,
		}
		for _, p := range rt.Platform.Properties {
			c.RemoteexecPlatform.Properties = append(c.RemoteexecPlatform.Properties, &cmdpb.RemoteexecPlatform_Property{
				Name:  p.Name,
				Value: p.Value,
			})
		}
		resp.Configs = append(resp.Configs, c)
	}
	return resp
}

func configureByLoader(ctx context.Context, loader *command.ConfigMapLoader, inventory *exec.Inventory, force bool) (string, error) {
	logger := log.FromContext(ctx)
	start := time.Now()
	resp, err := loader.Load(ctx, force)
	logger.Infof("loader.Load finished in %s: %v", time.Since(start), err)
	if err != nil {
		return "", err
	}
	start = time.Now()
	err = inventory.Configure(ctx, resp)
	logger.Infof("inventory.Configure finished in %s: %v", time.Since(start), err)
	if err != nil {
		return "", err
	}
	return resp.VersionId, nil
}

type cmdStorageBucket struct {
	Bucket *storage.BucketHandle
}

func (b cmdStorageBucket) Open(ctx context.Context, hash string) (io.ReadCloser, error) {
	return b.Bucket.Object(path.Join("sha256", hash)).NewReader(ctx)
}

type nullServer struct {
	ch chan error
}

func (s nullServer) ListenAndServe() error { return <-s.ch }
func (s nullServer) Shutdown(ctx context.Context) error {
	close(s.ch)
	return nil
}

type configServer struct {
	inventory *exec.Inventory
	configmap command.ConfigMap
	psclient  *pubsub.Client
	w         command.ConfigMapWatcher
	loader    *command.ConfigMapLoader
	cancel    func()
}

func newConfigServer(ctx context.Context, inventory *exec.Inventory, bucket, configMapFile string, cm *cmdpb.ConfigMap, gsclient *storage.Client, opts ...option.ClientOption) (*configServer, error) {
	cs := &configServer{
		inventory: inventory,
	}
	if *pubsubProjectID == "" {
		*pubsubProjectID = server.ProjectID(ctx)
		if *pubsubProjectID == "" {
			return nil, errors.New("--pubsub-project-id must be set")
		}
	}
	var err error
	cs.psclient, err = pubsub.NewClient(ctx, *pubsubProjectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("pubsub client failed: %v", err)
	}
	cs.configmap = command.ConfigMapBucket{
		URI:            fmt.Sprintf("gs://%s/", bucket),
		ConfigMap:      cm,
		ConfigMapFile:  configMapFile,
		StorageClient:  stiface.AdaptClient(gsclient),
		PubsubClient:   cs.psclient,
		SubscriberID:   fmt.Sprintf("toolchain-config-%s-%s", server.ClusterName(ctx), server.HostName(ctx)),
		RemoteexecAddr: *remoteexecAddr,
	}
	cs.w = cs.configmap.Watcher(ctx)
	cs.loader = &command.ConfigMapLoader{
		ConfigMap: cs.configmap,
		ConfigLoader: command.ConfigLoader{
			StorageClient:  stiface.AdaptClient(gsclient),
			EnableParallel: *fetchConfigParallel,
		},
	}
	return cs, nil
}

func (cs *configServer) configure(ctx context.Context, force bool) error {
	logger := log.FromContext(ctx)
	id, err := configureByLoader(ctx, cs.loader, cs.inventory, force)
	if errors.Is(err, context.Canceled) {
		logger.Errorf("canceled to configure: %v", err)
		return err
	}
	if err != nil {
		if err != command.ErrNoUpdate {
			recordConfigUpdate(ctx, err)
		}
		logger.Errorf("failed to configure: %v", err)
		return err
	}
	logger.Infof("configure %s", id)
	recordConfigUpdate(ctx, nil)
	return nil
}

func (cs *configServer) ListenAndServe() error {
	ctx, cancel := context.WithCancel(context.Background())
	cs.cancel = cancel
	logger := log.FromContext(ctx)
	var backoff *gax.Backoff
	for {
		wctx := ctx
		cancel := func() {}
		var timeout time.Duration
		if backoff != nil {
			timeout = backoff.Pause()
			wctx, cancel = context.WithTimeout(wctx, timeout)
		}
		logger.Infof("waiting for config update... timeout:%s", timeout)
		err := cs.w.Next(wctx)
		cancel()
		if err != nil {
			logger.Errorf("watch failed %v", err)
			if backoff == nil {
				return err
			}
			if errors.Is(err, command.ErrWatcherClosed) {
				return err
			}
			// if backoff != nil, Next may return context.Canceled
			// or so due to timeout in wctx.  Try loading anyway.
		}
		force := backoff != nil
		err = cs.configure(ctx, force)
		if errors.Is(err, context.Canceled) {
			// configServer was shutted down.
			return err
		}
		if err == command.ErrNoUpdate {
			backoff = nil
			continue
		}
		if err != nil {
			// loader  may not get all objects matched around storage@v1.15.0
			// https://github.com/googleapis/google-cloud-go/issues/4676
			logger.Errorf("config failed: %v", err)
			if backoff == nil {
				backoff = &gax.Backoff{
					Initial: time.Minute,
					Max:     time.Hour,
				}
			}
			continue
		}
		backoff = nil
	}
}

func (cs *configServer) Shutdown(ctx context.Context) error {
	defer errorreporter.Do(nil, nil)
	defer func() {
		if cs.psclient != nil {
			cs.psclient.Close()
		}
	}()
	if cs.cancel != nil {
		cs.cancel()
	}
	return cs.w.Close()
}

func newDigestCache(ctx context.Context) remoteexec.DigestCache {
	logger := log.FromContext(ctx)
	addr, err := redis.AddrFromEnv()
	if err != nil {
		logger.Warnf("redis disabled for gomafile-digest: %v", err)
		return digest.NewCache(nil, *maxDigestCacheEntries)
	}
	logger.Infof("redis enabled for gomafile-digest: %v idle=%d active=%d", addr, *redisMaxIdleConns, *redisMaxActiveConns)
	return digest.NewCache(redis.NewClient(ctx, addr, redis.Opts{
		Prefix:         "gomafile-digest:",
		MaxIdleConns:   *redisMaxIdleConns,
		MaxActiveConns: *redisMaxActiveConns,
	}), *maxDigestCacheEntries)
}

func main() {
	spanTimeout := remoteexec.DefaultSpanTimeout
	flag.DurationVar(&spanTimeout.Inventory, "exec-inventory-timeout", spanTimeout.Inventory, "timeout of exec-inventory")
	flag.DurationVar(&spanTimeout.InputTree, "exec-input-tree-timeout", spanTimeout.InputTree, "timeout of exec-iput-tree")
	flag.DurationVar(&spanTimeout.Setup, "exec-setup-timeout", spanTimeout.Setup, "timeout of exec-setup")
	flag.DurationVar(&spanTimeout.CheckCache, "exec-check-cache-timeout", spanTimeout.CheckCache, "timeout of exec-check-cache")
	flag.DurationVar(&spanTimeout.CheckMissing, "exec-check-missing-timeout", spanTimeout.CheckMissing, "timeout of exec-check-missing")
	flag.DurationVar(&spanTimeout.UploadBlobs, "exec-upload-blobs-timeout", spanTimeout.UploadBlobs, "timeout of exec-upload-blobs")
	flag.DurationVar(&spanTimeout.Execute, "exec-execute-timeout", spanTimeout.Execute, "timeout of exec-execute")
	flag.DurationVar(&spanTimeout.Response, "exec-response-timeout", spanTimeout.Response, "timeout of exec-response")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	ctx := context.Background()

	profiler.Setup(ctx)

	logger := log.FromContext(ctx)
	defer logger.Sync()

	if (*toolchainConfigBucket == "" || *configMapFile == "") && *configMap == "" {
		logger.Fatalf("--toolchain-config-bucket,--configmap_file or --configmap must be given")
	}
	if *remoteexecAddr == "" {
		logger.Fatalf("--remoteexec-addr must be given")
	}

	err := server.Init(ctx, *traceProjectID, "exec_server")
	if err != nil {
		logger.Fatal(err)
	}

	err = view.Register(configViews...)
	if err != nil {
		logger.Fatal(err)
	}
	err = view.Register(command.DefaultViews...)
	if err != nil {
		logger.Fatal(err)
	}
	err = view.Register(exec.DefaultViews...)
	if err != nil {
		logger.Fatal(err)
	}
	err = view.Register(remoteexec.DefaultViews...)
	if err != nil {
		logger.Fatal(err)
	}
	err = view.Register(digest.DefaultViews...)
	if err != nil {
		logger.Fatal(err)
	}
	trace.ApplyConfig(trace.Config{
		DefaultSampler: server.NewLimitedSampler(server.DefaultTraceFraction, server.DefaultTraceQPS),
	})

	s, err := server.NewGRPC(*port,
		grpc.MaxSendMsgSize(exec.DefaultMaxRespMsgSize),
		grpc.MaxRecvMsgSize(exec.DefaultMaxReqMsgSize))
	if err != nil {
		logger.Fatal(err)
	}

	fileConn, err := server.DialContext(ctx, *fileAddr,
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(file.DefaultMaxMsgSize), grpc.MaxCallSendMsgSize(file.DefaultMaxMsgSize)))
	if err != nil {
		logger.Fatalf("dial %s: %v", *fileAddr, err)
	}
	defer fileConn.Close()

	var gsclient *storage.Client
	var opts []option.ClientOption
	if *toolchainConfigBucket != "" || *cmdFilesBucket != "" {
		logger.Infof("toolchain-config-bucket or cmd-files-bucket is specified. use cloud storage")
		if *serviceAccountFile != "" {
			opts = append(opts, option.WithServiceAccountFile(*serviceAccountFile))
		}
		gsclient, err = storage.NewClient(ctx, opts...)
		if err != nil {
			logger.Fatalf("storage client failed: %v", err)
		}
		defer gsclient.Close()
	} else {
		logger.Infof("configmap_uri nor cmd-files-bucket is not specified. don't use cloud storage")
	}

	logger.Infof("use remoteexec API: %s", *remoteexecAddr)
	reConn, err := grpc.DialContext(ctx, *remoteexecAddr,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		logger.Fatalf("dial %s: %v", *remoteexecAddr, err)
	}
	defer reConn.Close()

	if *remoteInstancePrefix == "" {
		logger.Fatalf("--remote-instance-prefix must be given for remoteexec API")
	}

	if *fileLookupConcurrency == 0 {
		*fileLookupConcurrency = 1
	}
	casBlobLookupConcurrency := 20
	outputFileConcurrency := 20
	logger.Infof("span timeout = %#v", spanTimeout)
	re := &remoteexec.Adapter{
		InstancePrefix:   *remoteInstancePrefix,
		InstanceBaseName: *remoteInstanceBaseName,
		ExecTimeout:      *execActionTimeout,
		SpanTimeout:      spanTimeout,
		Client: remoteexec.Client{
			ClientConn: reConn,
			Retry: rpc.Retry{
				MaxRetry: *execMaxRetryCount,
			},
		},
		GomaFile:    filepb.NewFileServiceClient(fileConn),
		DigestCache: newDigestCache(ctx),
		ToolDetails: &rpb.ToolDetails{
			ToolName:    "goma/exec-server",
			ToolVersion: "0.0.0-experimental",
		},
		FileLookupSema:    make(chan struct{}, *fileLookupConcurrency),
		CASBlobLookupSema: make(chan struct{}, casBlobLookupConcurrency),
		OutputFileSema:    make(chan struct{}, outputFileConcurrency),
		HardeningRatio:    *experimentHardeningRatio,
		NsjailRatio:       *experimentNsjailRatio,
		DisableHardenings: strings.Split(*disableHardenings, ","),
	}
	logger.Infof("hardeniong=%f nsjail=%f", re.HardeningRatio, re.NsjailRatio)

	if *cmdFilesBucket == "" {
		logger.Warnf("--cmd-files-bucket is not given. support only ARBITRARY_TOOLCHAIN_SUPPORT enabled client")
	} else {
		logger.Infof("use gs://%s for cmd files", *cmdFilesBucket)
		re.CmdStorage = cmdStorageBucket{
			Bucket: gsclient.Bucket(*cmdFilesBucket),
		}
	}

	inventory := &re.Inventory

	// expose bytestream proxy.
	bs := &remoteexec.ByteStream{
		Adapter:      re,
		InstanceName: re.Instance(),
		// TODO: Create bytestreams for multiple instances.
	}
	bspb.RegisterByteStreamServer(s.Server, bs)

	var confServer server.Server
	ready := make(chan error)
	switch {
	case *configMap != "":
		go func() {
			cm := &cmdpb.ConfigMap{}
			err := prototext.Unmarshal([]byte(*configMap), cm)
			if err != nil {
				ready <- fmt.Errorf("parse configmap %q: %v", *configMap, err)
				return
			}
			resp := configMapToConfigResp(ctx, cm)
			err = inventory.Configure(ctx, resp)
			if err != nil {
				ready <- err
				return
			}
			logger.Infof("configure %s", resp.VersionId)
			ready <- nil
		}()
		confServer = nullServer{ch: make(chan error)}

	case *toolchainConfigBucket != "":
		cm := &cmdpb.ConfigMap{}
		if *configMap != "" {
			err := prototext.Unmarshal([]byte(*configMap), cm)
			if err != nil {
				ready <- fmt.Errorf("parse configmap %q: %v", *configMap, err)
				return
			}
		}
		cs, err := newConfigServer(ctx, inventory, *toolchainConfigBucket, *configMapFile, cm, gsclient, opts...)
		if err != nil {
			logger.Fatalf("configServer: %v", err)
		}
		go func() {
			ready <- cs.configure(ctx, true)
		}()
		confServer = cs
	}
	http.Handle("/configz", inventory)
	pb.RegisterExecServiceServer(s.Server, re)

	// as of Dec 14 2018, it takes about 45 seconds to be ready.
	// so wait 90-110 seconds with buffer.  b/120394151
	// assume initialDelaySeconds: 120.
	// TODO: split toolchain config server and exec server? b/120115232
	timeout := 90*time.Second + time.Duration(float64(20*time.Second)*rand.Float64())
	logger.Infof("wait %s to be ready", timeout)
	start := time.Now()
	select {
	case err = <-ready:
		if err != nil {
			logger.Errorf("configure: %v", err)
			confServer.Shutdown(ctx)
			server.Flush()
			logger.Fatalf("initial config failed: %v", err)
		}
		logger.Infof("exec-server ready in %s", time.Since(start))
	case <-time.After(timeout):
		logger.Errorf("initial loading timed out")
		confServer.Shutdown(ctx)
		server.Flush()
		logger.Fatalf("no configs available in %s", timeout)
	}
	hs := server.NewHTTP(*mport, nil)
	zpages.Handle(http.DefaultServeMux, "/debug")
	server.Run(ctx, s, hs, confServer)
}
