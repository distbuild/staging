// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pkg/proto/completedactionlogger/completed_action_logger.proto

package completedactionlogger

import (
	context "context"
	cas "github.com/buildbarn/bb-storage/pkg/proto/cas"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CompletedAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HistoricalExecuteResponse *cas.HistoricalExecuteResponse `protobuf:"bytes,1,opt,name=historical_execute_response,json=historicalExecuteResponse,proto3" json:"historical_execute_response,omitempty"`
	Uuid                      string                         `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	InstanceName              string                         `protobuf:"bytes,3,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
}

func (x *CompletedAction) Reset() {
	*x = CompletedAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_completedactionlogger_completed_action_logger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletedAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletedAction) ProtoMessage() {}

func (x *CompletedAction) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_completedactionlogger_completed_action_logger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletedAction.ProtoReflect.Descriptor instead.
func (*CompletedAction) Descriptor() ([]byte, []int) {
	return file_pkg_proto_completedactionlogger_completed_action_logger_proto_rawDescGZIP(), []int{0}
}

func (x *CompletedAction) GetHistoricalExecuteResponse() *cas.HistoricalExecuteResponse {
	if x != nil {
		return x.HistoricalExecuteResponse
	}
	return nil
}

func (x *CompletedAction) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CompletedAction) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

var File_pkg_proto_completedactionlogger_completed_action_logger_proto protoreflect.FileDescriptor

var file_pkg_proto_completedactionlogger_completed_action_logger_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x1a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x73, 0x2f,
	0x63, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x68, 0x0a, 0x1b, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x61, 0x73, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x19, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x7c, 0x0a,
	0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x63, 0x0a, 0x13, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x30, 0x01, 0x42, 0x4a, 0x5a, 0x48, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x61, 0x72, 0x6e, 0x2f, 0x62, 0x62, 0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_completedactionlogger_completed_action_logger_proto_rawDescOnce sync.Once
	file_pkg_proto_completedactionlogger_completed_action_logger_proto_rawDescData = file_pkg_proto_completedactionlogger_completed_action_logger_proto_rawDesc
)

func file_pkg_proto_completedactionlogger_completed_action_logger_proto_rawDescGZIP() []byte {
	file_pkg_proto_completedactionlogger_completed_action_logger_proto_rawDescOnce.Do(func() {
		file_pkg_proto_completedactionlogger_completed_action_logger_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_completedactionlogger_completed_action_logger_proto_rawDescData)
	})
	return file_pkg_proto_completedactionlogger_completed_action_logger_proto_rawDescData
}

var file_pkg_proto_completedactionlogger_completed_action_logger_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_completedactionlogger_completed_action_logger_proto_goTypes = []interface{}{
	(*CompletedAction)(nil),               // 0: buildbarn.completedactionlogger.CompletedAction
	(*cas.HistoricalExecuteResponse)(nil), // 1: buildbarn.cas.HistoricalExecuteResponse
	(*emptypb.Empty)(nil),                 // 2: google.protobuf.Empty
}
var file_pkg_proto_completedactionlogger_completed_action_logger_proto_depIdxs = []int32{
	1, // 0: buildbarn.completedactionlogger.CompletedAction.historical_execute_response:type_name -> buildbarn.cas.HistoricalExecuteResponse
	0, // 1: buildbarn.completedactionlogger.CompletedActionLogger.LogCompletedActions:input_type -> buildbarn.completedactionlogger.CompletedAction
	2, // 2: buildbarn.completedactionlogger.CompletedActionLogger.LogCompletedActions:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_completedactionlogger_completed_action_logger_proto_init() }
func file_pkg_proto_completedactionlogger_completed_action_logger_proto_init() {
	if File_pkg_proto_completedactionlogger_completed_action_logger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_completedactionlogger_completed_action_logger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletedAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_completedactionlogger_completed_action_logger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_completedactionlogger_completed_action_logger_proto_goTypes,
		DependencyIndexes: file_pkg_proto_completedactionlogger_completed_action_logger_proto_depIdxs,
		MessageInfos:      file_pkg_proto_completedactionlogger_completed_action_logger_proto_msgTypes,
	}.Build()
	File_pkg_proto_completedactionlogger_completed_action_logger_proto = out.File
	file_pkg_proto_completedactionlogger_completed_action_logger_proto_rawDesc = nil
	file_pkg_proto_completedactionlogger_completed_action_logger_proto_goTypes = nil
	file_pkg_proto_completedactionlogger_completed_action_logger_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CompletedActionLoggerClient is the client API for CompletedActionLogger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompletedActionLoggerClient interface {
	LogCompletedActions(ctx context.Context, opts ...grpc.CallOption) (CompletedActionLogger_LogCompletedActionsClient, error)
}

type completedActionLoggerClient struct {
	cc grpc.ClientConnInterface
}

func NewCompletedActionLoggerClient(cc grpc.ClientConnInterface) CompletedActionLoggerClient {
	return &completedActionLoggerClient{cc}
}

func (c *completedActionLoggerClient) LogCompletedActions(ctx context.Context, opts ...grpc.CallOption) (CompletedActionLogger_LogCompletedActionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CompletedActionLogger_serviceDesc.Streams[0], "/buildbarn.completedactionlogger.CompletedActionLogger/LogCompletedActions", opts...)
	if err != nil {
		return nil, err
	}
	x := &completedActionLoggerLogCompletedActionsClient{stream}
	return x, nil
}

type CompletedActionLogger_LogCompletedActionsClient interface {
	Send(*CompletedAction) error
	Recv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type completedActionLoggerLogCompletedActionsClient struct {
	grpc.ClientStream
}

func (x *completedActionLoggerLogCompletedActionsClient) Send(m *CompletedAction) error {
	return x.ClientStream.SendMsg(m)
}

func (x *completedActionLoggerLogCompletedActionsClient) Recv() (*emptypb.Empty, error) {
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CompletedActionLoggerServer is the server API for CompletedActionLogger service.
type CompletedActionLoggerServer interface {
	LogCompletedActions(CompletedActionLogger_LogCompletedActionsServer) error
}

// UnimplementedCompletedActionLoggerServer can be embedded to have forward compatible implementations.
type UnimplementedCompletedActionLoggerServer struct {
}

func (*UnimplementedCompletedActionLoggerServer) LogCompletedActions(CompletedActionLogger_LogCompletedActionsServer) error {
	return status.Errorf(codes.Unimplemented, "method LogCompletedActions not implemented")
}

func RegisterCompletedActionLoggerServer(s grpc.ServiceRegistrar, srv CompletedActionLoggerServer) {
	s.RegisterService(&_CompletedActionLogger_serviceDesc, srv)
}

func _CompletedActionLogger_LogCompletedActions_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CompletedActionLoggerServer).LogCompletedActions(&completedActionLoggerLogCompletedActionsServer{stream})
}

type CompletedActionLogger_LogCompletedActionsServer interface {
	Send(*emptypb.Empty) error
	Recv() (*CompletedAction, error)
	grpc.ServerStream
}

type completedActionLoggerLogCompletedActionsServer struct {
	grpc.ServerStream
}

func (x *completedActionLoggerLogCompletedActionsServer) Send(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *completedActionLoggerLogCompletedActionsServer) Recv() (*CompletedAction, error) {
	m := new(CompletedAction)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CompletedActionLogger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buildbarn.completedactionlogger.CompletedActionLogger",
	HandlerType: (*CompletedActionLoggerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LogCompletedActions",
			Handler:       _CompletedActionLogger_LogCompletedActions_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/proto/completedactionlogger/completed_action_logger.proto",
}