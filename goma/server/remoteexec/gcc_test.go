// Copyright 2018 The Goma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package remoteexec

import (
	"reflect"
	"strings"
	"testing"

	"go.chromium.org/goma/server/command/descriptor/posixpath"
)

func TestGccPathFlags(t *testing.T) {
	lastP := pathFlags[0]
	for _, p := range pathFlags[1:] {
		if len(lastP) < len(p) {
			t.Errorf("%q is longer than %q", p, lastP)
		}
		lastP = p
	}
}

func TestGccRelocatableReq(t *testing.T) {

	baseReleaseArgs := []string{
		"../../third_party/llvm-build/Release+Asserts/bin/clang++",
		"-MMD",
		"-MF",
		"obj/base/base/time.o.d",
		"-DUSE_SYMBOLIZE",
		"-I../..",
		"-fno-strict-aliasing",
		"--param=ssp-buffer-size=4",
		"-fPIC",
		"-pipe",
		"-B../../third_party/binutils/Linux_x64/Release/bin",
		"-pthread",
		"-Xclang",
		"-mllvm",
		"-Xclang",
		"-instcombine-lower-dbg-declare=0",
		"-no-canonical-prefixes",
		"-m64",
		"-march=x86-64",
		"-Wall",
		"-g0",
		"-Xclang",
		"-load",
		"-Xclang",
		"../../third_party/llvm-build/Release+Asserts/lib/libFindBadCustructs.so",
		"-Xclang",
		"-add-plugin",
		"-Xclang",
		"find-bad-constructs",
		"-Xclang",
		"-plugin-arg-find-bad-constructs",
		"-Xclang",
		"check-enum-max-value",
		"-isystem../../build/linux/debian_sid_amd64-sysroot/usr/include/glib-2.0",
		"-O2",
		"--sysroot=../../build/linux/debian_sid_amd64-sysroot",
		"-c",
		"../../base/time/time.cc",
		"-o",
		"obj/base/base/time.o",
	}

	baseDebugArgs := []string{
		"../../third_party/llvm-build/Release+Asserts/bin/clang++",
		"-MMD",
		"-MF",
		"obj/base/base/time.o.d",
		"-DUSE_SYMBOLIZE",
		"-I../..",
		"-fno-strict-aliasing",
		"--param=ssp-buffer-size=4",
		"-fPIC",
		"-pipe",
		"-B../../third_party/binutils/Linux_x64/Release/bin",
		"-pthread",
		"-Xclang",
		"-mllvm",
		"-Xclang",
		"-instcombine-lower-dbg-declare=0",
		"-no-canonical-prefixes",
		"-m64",
		"-march=x86-64",
		"-Wall",
		"-g2",
		"-gsplit-dwarf",
		"-Xclang",
		"-load",
		"-Xclang",
		"../../third_party/llvm-build/Release+Asserts/lib/libFindBadCustructs.so",
		"-Xclang",
		"-add-plugin",
		"-Xclang",
		"find-bad-constructs",
		"-Xclang",
		"-plugin-arg-find-bad-constructs",
		"-Xclang",
		"check-enum-max-value",
		"-isystem../../build/linux/debian_sid_amd64-sysroot/usr/include/glib-2.0",
		"-O2",
		"--sysroot=../../build/linux/debian_sid_amd64-sysroot",
		"-c",
		"../../base/time/time.cc",
		"-o",
		"obj/base/base/time.o",
	}

	modifyArgs := func(args []string, prefix, replace string) []string {
		var ret []string
		found := false
		for _, arg := range args {
			if strings.HasPrefix(arg, prefix) {
				ret = append(ret, replace)
				found = true
				continue
			}
			ret = append(ret, arg)
		}
		if !found {
			ret = append(ret, replace)
		}
		return ret
	}

	for _, tc := range []struct {
		desc        string
		args        []string
		envs        []string
		relocatable bool
	}{
		{
			desc: "chromium base release",
			args: baseReleaseArgs,
			envs: []string{
				"PWD=/b/c/b/linux/src/out/Release",
			},
			relocatable: true,
		},
		{
			desc: "chromium base debug",
			args: baseDebugArgs,
			envs: []string{
				"PWD=/b/c/b/linux/src/out/Debug",
			},
			relocatable: false,
		},
		{
			desc: "resource-dir relative",
			args: append(append([]string{}, baseReleaseArgs...),
				"-resource-dir=../../third_party/llvm-build-release+Asserts/bin/clang/10.0.0"),
			relocatable: true,
		},
		{
			desc: "resource-dir absolute",
			args: append(append([]string{}, baseReleaseArgs...),
				"-resource-dir=/b/c/b/linux/src/third_party/llvm-build-release+Asserts/bin/clang/10.0.0"),
			relocatable: false,
		},
		{
			desc: "isystem absolute",
			args: append(append([]string{}, baseReleaseArgs...),
				"-isystem/b/c/b/linux/src/build/linux/debian_sid_amd64-sysroot/usr/include/glib-2.0"),
			relocatable: false,
		},
		{
			desc: "include absolute",
			args: append(append([]string{}, baseReleaseArgs...),
				"-include",
				"/b/c/b/linux/header.h"),
			relocatable: false,
		},
		{
			desc: "include relative",
			args: append(append([]string{}, baseReleaseArgs...),
				"-include",
				"../b/c/b/linux/header.h"),
			relocatable: true,
		},
		{
			desc: "include= absolute",
			args: append(append([]string{}, baseReleaseArgs...),
				"-include=/b/c/b/linux/header.h"),
			relocatable: false,
		},
		{
			desc: "include= relative",
			args: append(append([]string{}, baseReleaseArgs...),
				"-include=../b/c/b/linux/header.h"),
			relocatable: true,
		},
		{
			desc: "sysroot absolute",
			args: modifyArgs(baseReleaseArgs,
				"--sysroot",
				"--sysroot=/b/c/b/linux/build/linux/debian_sid_amd64-sysroot"),
			relocatable: false,
		},
		{
			desc: "sysroot separate relative",
			args: append(modifyArgs(baseReleaseArgs,
				"--sysroot",
				"--sysroot"),
				"../../build/linux/debian_sid_amd64-sysroot"),
			relocatable: true,
		},
		{
			desc: "sysroot separate absolute",
			args: append(modifyArgs(baseReleaseArgs,
				"--sysroot",
				"--sysroot"),
				"/b/c/b/linux/build/linux/debian_sid_amd64-sysroot"),
			relocatable: false,
		},
		{
			desc: "llvm -asan option",
			// https://b/issues/141210713#comment3
			args: append(append([]string{}, baseReleaseArgs...),
				"-mllvm",
				"-asan-globals=0"),
			relocatable: true,
		},
		{
			desc: "llvm -regalloc option",
			// https://b/issues/141210713#comment4
			args: append(append([]string{}, baseReleaseArgs...),
				"-mllvm",
				"-regalloc=pbqp",
				"-mllvm",
				"-pbqp-coalescing"),
			relocatable: true,
		},
		{
			desc: "headermap file",
			// http://b/149448356#comment17
			args: append(append([]string{}, baseReleaseArgs...),
				"-Ifoo.hmap"),
			relocatable: false,
		},
		{
			desc: "headermap file 2",
			// http://b/149448356#comment17
			args: append(append([]string{}, baseReleaseArgs...),
				[]string{"-I", "foo.hmap"}...),
			relocatable: false,
		},
		{
			desc: "Qunused-arguments",
			args: append(append([]string{}, baseReleaseArgs...),
				"-Qunused-arguments"),
			relocatable: true,
		},
		{
			desc: "fprofile-instr-use relative",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fprofile-instr-use=../../out/data/default.profdata"),
			relocatable: true,
		},
		{
			desc: "fprofile-instr-use absolute",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fprofile-instr-use=/b/c/b/linux/src/out/data/default.profdataa"),
			relocatable: false,
		},
		{
			desc: "fprofile-sample-use relative",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fprofile-sample-use=../../chrome/android/profiles/afdo.prof"),
			relocatable: true,
		},
		{
			desc: "fprofile-sample-use absolute",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fprofile-sample-use=/b/c/b/linux/src/android/profiles/afdo.prof"),
			relocatable: false,
		},
		{
			desc: "fsatinitize-blacklist relative",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fsanitize-blacklist=../../tools/cfi/blacklist.txt"),
			relocatable: true,
		},
		{
			desc: "fsanitize-blacklist absolute",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fsanitize-blacklist=/b/c/b/linux/src/tools/cfi/blacklist.txt"),
			relocatable: false,
		},
		{
			desc: "fcrash-diagnostics-dir absolute",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fcrash-diagnostics-dir=../../tools/clang/crashreports"),
			relocatable: true,
		},
		{
			desc: "fcrash-diagnostics-dir absolute",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fcrash-diagnostics-dir=/b/c/b/linux/src/tools/clang/crashreports"),
			relocatable: false,
		},
		{
			desc: "static-libgcc",
			args: append(append([]string{}, baseReleaseArgs...),
				"-static-libgcc"),
			relocatable: true,
		},
		{
			desc: "idirafter relative",
			args: append(append([]string{}, baseReleaseArgs...),
				"-idirafter", "../../system/ulib/include"),
			relocatable: true,
		},
		{
			desc: "idirafter abs",
			args: append(append([]string{}, baseReleaseArgs...),
				"-idirafter", "/b/c/system/ulib/include"),
			relocatable: false,
		},
		{
			desc: "-mllvm -instcombine-lower-dbg-declare=0",
			args: append(append([]string{}, baseReleaseArgs...),
				"-mllvm", "-instcombine-lower-dbg-declare=0"),
			relocatable: true,
		},
		{
			desc: "-mllvm ml inliner relocatable",
			args: append(append([]string{}, baseReleaseArgs...),
				"-mllvm", "-enable-ml-inliner=development", "-mllvm", "-training-log=./train.log", "-mllvm", "-ml-inliner-model-under-training=./tf.policy"),
			relocatable: true,
		},
		{
			desc: "-mllvm ml inliner non-relocatable train log",
			args: append(append([]string{}, baseReleaseArgs...),
				"-mllvm", "-enable-ml-inliner=development", "-mllvm", "-training-log=/abs/train.log", "-mllvm", "-ml-inliner-model-under-training=./tf.policy"),
			relocatable: false,
		},
		{
			desc: "-mllvm ml inliner non-relocatable policy file",
			args: append(append([]string{}, baseReleaseArgs...),
				"-mllvm", "-enable-ml-inliner=development", "-mllvm", "-training-log=./train.log", "-mllvm", "-ml-inliner-model-under-training=/abs/tf.policy"),
			relocatable: false,
		},
		{
			desc: "-includeBuildConfig.h",
			args: append(append([]string{}, baseReleaseArgs...),
				"-includeBuildConfig.h"),
			relocatable: true,
		},
		{
			desc: "-include/usr/include/BuildConfig.h",
			args: append(append([]string{}, baseReleaseArgs...),
				"-include/usr/include/BuildConfig.h"),
			relocatable: false,
		},
		{
			desc: "--includeBuildConfig.h",
			args: append(append([]string{}, baseReleaseArgs...),
				"--includeBuildConfig.h"),
			relocatable: true,
		},
		{
			desc: "--include/usr/include/BuildConfig.h",
			args: append(append([]string{}, baseReleaseArgs...),
				"--include/usr/include/BuildConfig.h"),
			relocatable: false,
		},
		{
			desc: "--include=BuildConfig.h",
			args: append(append([]string{}, baseReleaseArgs...),
				"--include=BuildConfig.h"),
			relocatable: true,
		},
		{
			desc: "--include=/usr/include/BuildConfig.h",
			args: append(append([]string{}, baseReleaseArgs...),
				"--include=/usr/include/BuildConfig.h"),
			relocatable: false,
		},
		{
			desc: "-mllvm -basic-aa-recphi=0",
			args: append(append([]string{}, baseReleaseArgs...),
				"-mllvm", "-basic-aa-recphi=0"),
			relocatable: true,
		},
		{
			desc: "-pie",
			args: append(append([]string{}, baseReleaseArgs...),
				"-pie"),
			relocatable: true,
		},
		{
			desc: "-Xclang -debug-info-kind=constructor",
			args: append(append([]string{}, baseReleaseArgs...),
				"-Xclang", "-debug-info-kind=constructor"),
			relocatable: true,
		},
		{
			desc: "-Xclang -fno-experimental-new-pass-manager",
			args: append(append([]string{}, baseReleaseArgs...),
				"-Xclang", "-fno-experimental-new-pass-manager"),
			relocatable: true,
		},
		{
			desc: "-Xclang -no-opaque-pointers",
			args: append(append([]string{}, baseReleaseArgs...),
				"-Xclang", "-no-opaque-pointers"),
			relocatable: true,
		},
		{
			desc: "-mllvm -sanitizer-coverage-prune-blocks=1",
			args: append(append([]string{}, baseReleaseArgs...),
				"-mllvm", "-sanitizer-coverage-prune-blocks=1"),
			relocatable: true,
		},
		{
			desc: "framework include search path -F.",
			args: append(append([]string{}, baseReleaseArgs...),
				"-F."),
			relocatable: true,
		},
		{
			desc: "framework include search path -F/App/Framework",
			args: append(append([]string{}, baseReleaseArgs...),
				"-F/App/Framework"),
			relocatable: false,
		},
		{
			desc: "-target",
			args: append(append([]string{}, baseReleaseArgs...),
				"-target", "x86_64-linux-gnu"),
			relocatable: true,
		},
		{
			desc: "-mllvm -enable-dse-memoryssa=true",
			args: append(append([]string{}, baseReleaseArgs...),
				"-mllvm", "-enable-dse-memoryssa=true"),
			relocatable: true,
		},
		{
			desc: "-mllvm -enable-dse-memoryssa=false",
			args: append(append([]string{}, baseReleaseArgs...),
				"-mllvm", "-enable-dse-memoryssa=false"),
			relocatable: true,
		},
		{
			desc: "-mllvm -limited-coverage-experimental=true",
			args: append(append([]string{}, baseReleaseArgs...),
				"-mllvm", "-limited-coverage-experimental=true"),
			relocatable: true,
		},
		{
			desc: "-Wa,--fatal-warnings",
			args: append(append([]string{}, baseReleaseArgs...),
				"-Wa,--fatal-warnings"),
			relocatable: true,
		},
		{
			desc: "-fcoverage-compilation-dir= unrelocatable",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fcoverage-compilation-dir=/b/c/b/linux/src/out/data/coverage-compilation-dir"),
			relocatable: false,
		},
		{
			desc: "-fcoverage-compilation-dir= relocatable",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fcoverage-compilation-dir=./out/data/coverage-compilation-dir"),
			relocatable: true,
		},
		{
			desc: "-ffile-compilation-dir= unrelocatable with debug build",
			args: append(append([]string{}, baseDebugArgs...),
				"-ffile-compilation-dir=/b/c/b/linux/src/out/data/file-compilation-dir"),
			relocatable: false,
		},
		{
			desc: "-ffile-compilation-dir= relocatable with debug build",
			args: append(append([]string{}, baseDebugArgs...),
				"-ffile-compilation-dir=./out/data/file-compilation-dir"),
			relocatable: true,
		},
		{
			desc: "-ffile-compilation-dir= unrelocatable with release build",
			args: append(append([]string{}, baseReleaseArgs...),
				"-ffile-compilation-dir=/b/c/b/linux/src/out/data/file-compilation-dir"),
			relocatable: false,
		},
		{
			desc: "-ffile-compilation-dir= relocatable with release build",
			args: append(append([]string{}, baseReleaseArgs...),
				"-ffile-compilation-dir=./out/data/file-compilation-dir"),
			relocatable: true,
		},
		{
			desc: "-fprofile-list= unrelocatable",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fprofile-list=/b/c/b/linux/src/out/data/profile-list"),
			relocatable: false,
		},
		{
			desc: "-fprofile-list= relocatable",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fprofile-list=./out/data/profile-list"),
			relocatable: true,
		},
		{
			desc: "clang -target-feature",
			args: append(append([]string{}, baseReleaseArgs...),
				"-Xclang", "-target-feature",
				"-Xclang", "+crc",
				"-Xclang", "-target-feature",
				"-Xclang", "+crypto"),
			relocatable: true,
		},
		{
			desc: "clang --rtlib=",
			args: append(append([]string{}, baseReleaseArgs...),
				"--rtlib=libgcc"),
			relocatable: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := gccRelocatableReq(posixpath.FilePath{}, tc.args, tc.envs)
			if (err == nil) != tc.relocatable {
				t.Errorf("gccRelocatableReq(posixpath.FilePath, args, envs)=%v; relocatable=%t", err, tc.relocatable)
			}
		})
	}
}

func TestGccRelocatableReqForDebugCompilationDir(t *testing.T) {
	// Tests for supporting "-fdebug-compilation-dir", see b/135719929.
	// We could have merged the cases here into TestGccRelocatableReq, but decided
	// to separate them for clarity.

	// Do not set "-g*" options in baseReleaseArgs!
	baseReleaseArgs := []string{
		"../../third_party/llvm-build/Release+Asserts/bin/clang++",
		"../../base/time/time.cc",
	}

	// Since "-fdebug-compilation-dir" has been moved to clang driver flags in
	// https://reviews.llvm.org/D63387, we set cases both with and w/o "-Xclang"
	for _, tc := range []struct {
		desc        string
		args        []string
		envs        []string
		relocatable bool
	}{
		{
			desc: "basic",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fdebug-compilation-dir",
				"."),
			relocatable: true,
		},
		{
			desc: "-Xclang",
			args: append(append([]string{}, baseReleaseArgs...),
				"-Xclang",
				"-fdebug-compilation-dir",
				"-Xclang",
				"."),
			relocatable: true,
		},
		{
			desc: "With -g* DBG options",
			args: append(append([]string{}, baseReleaseArgs...),
				"-g2",
				"-gsplit-dwarf",
				"-fdebug-compilation-dir",
				"."),
			relocatable: true,
		},
		{
			desc: "-Xclang with -g* DBG option",
			args: append(append([]string{}, baseReleaseArgs...),
				"-g2",
				"-gsplit-dwarf",
				"-Xclang",
				"-fdebug-compilation-dir",
				"-Xclang",
				"."),
			relocatable: true,
		},
		{
			// Make sure the CWD agnostic still returns false if
			// "-fdebug-compilation-dir" is not specified.
			desc: "Only -g* DBG options",
			args: append(append([]string{}, baseReleaseArgs...),
				"-g2",
				"-gsplit-dwarf"),
			relocatable: false,
		},
		{
			// "-fdebug-compilation-dir" is not supported as LLVM flags.
			desc: "No LLVM",
			args: append(append([]string{}, baseReleaseArgs...),
				"-mllvm",
				"-fdebug-compilation-dir",
				"-mllvm",
				"."),
			relocatable: false,
		},
		{
			desc: "input abs path",
			args: append(append([]string{}, baseReleaseArgs...),
				"-fdebug-compilation-dir",
				".",
				"-isystem/b/c/b/linux/src/build/linux/debian_sid_amd64-sysroot/usr/include/glib-2.0"),
			relocatable: false,
		},
		{
			desc: "-fdebug-compilation-dir= unrelocatable",
			args: append(append([]string{},
				baseReleaseArgs...),
				"-fdebug-compilation-dir=/b/c/b/linux/src/out/data/debug_compilation_dir"),
			relocatable: false,
		},
		{
			desc: "-fdebug-compilation-dir= relocatable",
			args: append(append([]string{},
				baseReleaseArgs...),
				"-fdebug-compilation-dir=./debug_compilation_dir"),
			relocatable: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := gccRelocatableReq(posixpath.FilePath{}, tc.args, tc.envs)
			if (err == nil) != tc.relocatable {
				t.Errorf("gccRelocatableReq(posixpath.FilePath, args, envs)=%v; relocatable=%t", err, tc.relocatable)
			}
		})
	}
}

func TestGccOutputs(t *testing.T) {
	for _, tc := range []struct {
		desc string
		args []string
		want []string
	}{
		{
			desc: "basic",
			args: []string{
				"gcc", "-c", "A/test.c",
				"-I", "A/B/C",
				"-ID/E/F",
				"-o", "A/test.o",
				"-MF", "test.d",
			},
			want: []string{"test.d", "A/test.o"},
		},
		{
			desc: "prefix",
			args: []string{
				"gcc", "-c", "A/test.c",
				"-I", "A/B/C",
				"-ID/E/F",
				"-oA/test.o",
				"-MFtest.d",
			},
			want: []string{"test.d", "A/test.o"},
		},
		{
			desc: "with dwo",
			args: []string{
				"gcc", "-c", "A/test.c",
				"-I", "A/B/C",
				"-ID/E/F",
				"-gsplit-dwarf",
				"-o", "A/test.o",
				"-MF", "test.d",
			},
			want: []string{"test.d", "A/test.o", "A/test.dwo"},
		},
		{
			desc: "prefix with dwo",
			args: []string{
				"gcc", "-c", "A/test.c",
				"-I", "A/B/C",
				"-ID/E/F",
				"-gsplit-dwarf",
				"-oA/test.o",
				"-MFtest.d",
			},
			want: []string{"test.d", "A/test.o", "A/test.dwo"},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			if got := gccOutputs(tc.args); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("gccOutputs(%q)=%q; want %q", tc.args, got, tc.want)
			}
		})
	}
}
