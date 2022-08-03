// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pkg/proto/iscc/iscc.proto

package iscc

import (
	context "context"
	v2 "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PreviousExecution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Outcome:
	//
	//	*PreviousExecution_Failed
	//	*PreviousExecution_TimedOut
	//	*PreviousExecution_Succeeded
	Outcome isPreviousExecution_Outcome `protobuf_oneof:"outcome"`
}

func (x *PreviousExecution) Reset() {
	*x = PreviousExecution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_iscc_iscc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviousExecution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviousExecution) ProtoMessage() {}

func (x *PreviousExecution) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_iscc_iscc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviousExecution.ProtoReflect.Descriptor instead.
func (*PreviousExecution) Descriptor() ([]byte, []int) {
	return file_pkg_proto_iscc_iscc_proto_rawDescGZIP(), []int{0}
}

func (m *PreviousExecution) GetOutcome() isPreviousExecution_Outcome {
	if m != nil {
		return m.Outcome
	}
	return nil
}

func (x *PreviousExecution) GetFailed() *emptypb.Empty {
	if x, ok := x.GetOutcome().(*PreviousExecution_Failed); ok {
		return x.Failed
	}
	return nil
}

func (x *PreviousExecution) GetTimedOut() *durationpb.Duration {
	if x, ok := x.GetOutcome().(*PreviousExecution_TimedOut); ok {
		return x.TimedOut
	}
	return nil
}

func (x *PreviousExecution) GetSucceeded() *durationpb.Duration {
	if x, ok := x.GetOutcome().(*PreviousExecution_Succeeded); ok {
		return x.Succeeded
	}
	return nil
}

type isPreviousExecution_Outcome interface {
	isPreviousExecution_Outcome()
}

type PreviousExecution_Failed struct {
	Failed *emptypb.Empty `protobuf:"bytes,1,opt,name=failed,proto3,oneof"`
}

type PreviousExecution_TimedOut struct {
	TimedOut *durationpb.Duration `protobuf:"bytes,2,opt,name=timed_out,json=timedOut,proto3,oneof"`
}

type PreviousExecution_Succeeded struct {
	Succeeded *durationpb.Duration `protobuf:"bytes,3,opt,name=succeeded,proto3,oneof"`
}

func (*PreviousExecution_Failed) isPreviousExecution_Outcome() {}

func (*PreviousExecution_TimedOut) isPreviousExecution_Outcome() {}

func (*PreviousExecution_Succeeded) isPreviousExecution_Outcome() {}

type PerSizeClassStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousExecutions         []*PreviousExecution `protobuf:"bytes,1,rep,name=previous_executions,json=previousExecutions,proto3" json:"previous_executions,omitempty"`
	InitialPageRankProbability float64              `protobuf:"fixed64,3,opt,name=initial_page_rank_probability,json=initialPageRankProbability,proto3" json:"initial_page_rank_probability,omitempty"`
}

func (x *PerSizeClassStats) Reset() {
	*x = PerSizeClassStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_iscc_iscc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerSizeClassStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerSizeClassStats) ProtoMessage() {}

func (x *PerSizeClassStats) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_iscc_iscc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerSizeClassStats.ProtoReflect.Descriptor instead.
func (*PerSizeClassStats) Descriptor() ([]byte, []int) {
	return file_pkg_proto_iscc_iscc_proto_rawDescGZIP(), []int{1}
}

func (x *PerSizeClassStats) GetPreviousExecutions() []*PreviousExecution {
	if x != nil {
		return x.PreviousExecutions
	}
	return nil
}

func (x *PerSizeClassStats) GetInitialPageRankProbability() float64 {
	if x != nil {
		return x.InitialPageRankProbability
	}
	return 0
}

type PreviousExecutionStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SizeClasses     map[uint32]*PerSizeClassStats `protobuf:"bytes,1,rep,name=size_classes,json=sizeClasses,proto3" json:"size_classes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LastSeenFailure *timestamppb.Timestamp        `protobuf:"bytes,2,opt,name=last_seen_failure,json=lastSeenFailure,proto3" json:"last_seen_failure,omitempty"`
}

func (x *PreviousExecutionStats) Reset() {
	*x = PreviousExecutionStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_iscc_iscc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviousExecutionStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviousExecutionStats) ProtoMessage() {}

func (x *PreviousExecutionStats) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_iscc_iscc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviousExecutionStats.ProtoReflect.Descriptor instead.
func (*PreviousExecutionStats) Descriptor() ([]byte, []int) {
	return file_pkg_proto_iscc_iscc_proto_rawDescGZIP(), []int{2}
}

func (x *PreviousExecutionStats) GetSizeClasses() map[uint32]*PerSizeClassStats {
	if x != nil {
		return x.SizeClasses
	}
	return nil
}

func (x *PreviousExecutionStats) GetLastSeenFailure() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSeenFailure
	}
	return nil
}

type GetPreviousExecutionStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceName        string     `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	ReducedActionDigest *v2.Digest `protobuf:"bytes,2,opt,name=reduced_action_digest,json=reducedActionDigest,proto3" json:"reduced_action_digest,omitempty"`
}

func (x *GetPreviousExecutionStatsRequest) Reset() {
	*x = GetPreviousExecutionStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_iscc_iscc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPreviousExecutionStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPreviousExecutionStatsRequest) ProtoMessage() {}

func (x *GetPreviousExecutionStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_iscc_iscc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPreviousExecutionStatsRequest.ProtoReflect.Descriptor instead.
func (*GetPreviousExecutionStatsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_iscc_iscc_proto_rawDescGZIP(), []int{3}
}

func (x *GetPreviousExecutionStatsRequest) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *GetPreviousExecutionStatsRequest) GetReducedActionDigest() *v2.Digest {
	if x != nil {
		return x.ReducedActionDigest
	}
	return nil
}

type UpdatePreviousExecutionStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceName           string                  `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	ReducedActionDigest    *v2.Digest              `protobuf:"bytes,2,opt,name=reduced_action_digest,json=reducedActionDigest,proto3" json:"reduced_action_digest,omitempty"`
	PreviousExecutionStats *PreviousExecutionStats `protobuf:"bytes,3,opt,name=previous_execution_stats,json=previousExecutionStats,proto3" json:"previous_execution_stats,omitempty"`
}

func (x *UpdatePreviousExecutionStatsRequest) Reset() {
	*x = UpdatePreviousExecutionStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_iscc_iscc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePreviousExecutionStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePreviousExecutionStatsRequest) ProtoMessage() {}

func (x *UpdatePreviousExecutionStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_iscc_iscc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePreviousExecutionStatsRequest.ProtoReflect.Descriptor instead.
func (*UpdatePreviousExecutionStatsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_iscc_iscc_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePreviousExecutionStatsRequest) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *UpdatePreviousExecutionStatsRequest) GetReducedActionDigest() *v2.Digest {
	if x != nil {
		return x.ReducedActionDigest
	}
	return nil
}

func (x *UpdatePreviousExecutionStatsRequest) GetPreviousExecutionStats() *PreviousExecutionStats {
	if x != nil {
		return x.PreviousExecutionStats
	}
	return nil
}

var File_pkg_proto_iscc_iscc_proto protoreflect.FileDescriptor

var file_pkg_proto_iscc_iscc_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x73, 0x63, 0x63,
	0x2f, 0x69, 0x73, 0x63, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x69, 0x73, 0x63, 0x63, 0x1a, 0x36, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48,
	0x00, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x64,
	0x4f, 0x75, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x42, 0x09,
	0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x11, 0x50, 0x65,
	0x72, 0x53, 0x69, 0x7a, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x52, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x69, 0x73, 0x63, 0x63, 0x2e, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x12, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x41, 0x0a, 0x1d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x1a, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x62, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x9f, 0x02, 0x0a,
	0x16, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x5a, 0x0a, 0x0c, 0x73, 0x69, 0x7a, 0x65, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x69, 0x73, 0x63, 0x63, 0x2e, 0x50,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x69, 0x7a, 0x65, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65, 0x6e,
	0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74,
	0x53, 0x65, 0x65, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x1a, 0x61, 0x0a, 0x10, 0x53,
	0x69, 0x7a, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x69, 0x73, 0x63,
	0x63, 0x2e, 0x50, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa4,
	0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x15, 0x72, 0x65, 0x64, 0x75,
	0x63, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x52, 0x13, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0x89, 0x02, 0x0a, 0x23, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x15, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x13, 0x72, 0x65, 0x64, 0x75,
	0x63, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12,
	0x60, 0x0a, 0x18, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x69, 0x73,
	0x63, 0x63, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x16, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x32, 0xfb, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x69, 0x7a,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x75, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x30, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x61, 0x72, 0x6e, 0x2e, 0x69, 0x73, 0x63, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x69, 0x73, 0x63, 0x63, 0x2e, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x6b, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x33, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x69,
	0x73, 0x63, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2f, 0x62, 0x62, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x73, 0x63,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_iscc_iscc_proto_rawDescOnce sync.Once
	file_pkg_proto_iscc_iscc_proto_rawDescData = file_pkg_proto_iscc_iscc_proto_rawDesc
)

func file_pkg_proto_iscc_iscc_proto_rawDescGZIP() []byte {
	file_pkg_proto_iscc_iscc_proto_rawDescOnce.Do(func() {
		file_pkg_proto_iscc_iscc_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_iscc_iscc_proto_rawDescData)
	})
	return file_pkg_proto_iscc_iscc_proto_rawDescData
}

var file_pkg_proto_iscc_iscc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_proto_iscc_iscc_proto_goTypes = []interface{}{
	(*PreviousExecution)(nil),                   // 0: buildbarn.iscc.PreviousExecution
	(*PerSizeClassStats)(nil),                   // 1: buildbarn.iscc.PerSizeClassStats
	(*PreviousExecutionStats)(nil),              // 2: buildbarn.iscc.PreviousExecutionStats
	(*GetPreviousExecutionStatsRequest)(nil),    // 3: buildbarn.iscc.GetPreviousExecutionStatsRequest
	(*UpdatePreviousExecutionStatsRequest)(nil), // 4: buildbarn.iscc.UpdatePreviousExecutionStatsRequest
	nil,                           // 5: buildbarn.iscc.PreviousExecutionStats.SizeClassesEntry
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
	(*durationpb.Duration)(nil),   // 7: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*v2.Digest)(nil),             // 9: build.bazel.remote.execution.v2.Digest
}
var file_pkg_proto_iscc_iscc_proto_depIdxs = []int32{
	6,  // 0: buildbarn.iscc.PreviousExecution.failed:type_name -> google.protobuf.Empty
	7,  // 1: buildbarn.iscc.PreviousExecution.timed_out:type_name -> google.protobuf.Duration
	7,  // 2: buildbarn.iscc.PreviousExecution.succeeded:type_name -> google.protobuf.Duration
	0,  // 3: buildbarn.iscc.PerSizeClassStats.previous_executions:type_name -> buildbarn.iscc.PreviousExecution
	5,  // 4: buildbarn.iscc.PreviousExecutionStats.size_classes:type_name -> buildbarn.iscc.PreviousExecutionStats.SizeClassesEntry
	8,  // 5: buildbarn.iscc.PreviousExecutionStats.last_seen_failure:type_name -> google.protobuf.Timestamp
	9,  // 6: buildbarn.iscc.GetPreviousExecutionStatsRequest.reduced_action_digest:type_name -> build.bazel.remote.execution.v2.Digest
	9,  // 7: buildbarn.iscc.UpdatePreviousExecutionStatsRequest.reduced_action_digest:type_name -> build.bazel.remote.execution.v2.Digest
	2,  // 8: buildbarn.iscc.UpdatePreviousExecutionStatsRequest.previous_execution_stats:type_name -> buildbarn.iscc.PreviousExecutionStats
	1,  // 9: buildbarn.iscc.PreviousExecutionStats.SizeClassesEntry.value:type_name -> buildbarn.iscc.PerSizeClassStats
	3,  // 10: buildbarn.iscc.InitialSizeClassCache.GetPreviousExecutionStats:input_type -> buildbarn.iscc.GetPreviousExecutionStatsRequest
	4,  // 11: buildbarn.iscc.InitialSizeClassCache.UpdatePreviousExecutionStats:input_type -> buildbarn.iscc.UpdatePreviousExecutionStatsRequest
	2,  // 12: buildbarn.iscc.InitialSizeClassCache.GetPreviousExecutionStats:output_type -> buildbarn.iscc.PreviousExecutionStats
	6,  // 13: buildbarn.iscc.InitialSizeClassCache.UpdatePreviousExecutionStats:output_type -> google.protobuf.Empty
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_proto_iscc_iscc_proto_init() }
func file_pkg_proto_iscc_iscc_proto_init() {
	if File_pkg_proto_iscc_iscc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_iscc_iscc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreviousExecution); i {
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
		file_pkg_proto_iscc_iscc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerSizeClassStats); i {
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
		file_pkg_proto_iscc_iscc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreviousExecutionStats); i {
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
		file_pkg_proto_iscc_iscc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPreviousExecutionStatsRequest); i {
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
		file_pkg_proto_iscc_iscc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePreviousExecutionStatsRequest); i {
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
	file_pkg_proto_iscc_iscc_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PreviousExecution_Failed)(nil),
		(*PreviousExecution_TimedOut)(nil),
		(*PreviousExecution_Succeeded)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_iscc_iscc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_iscc_iscc_proto_goTypes,
		DependencyIndexes: file_pkg_proto_iscc_iscc_proto_depIdxs,
		MessageInfos:      file_pkg_proto_iscc_iscc_proto_msgTypes,
	}.Build()
	File_pkg_proto_iscc_iscc_proto = out.File
	file_pkg_proto_iscc_iscc_proto_rawDesc = nil
	file_pkg_proto_iscc_iscc_proto_goTypes = nil
	file_pkg_proto_iscc_iscc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InitialSizeClassCacheClient is the client API for InitialSizeClassCache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InitialSizeClassCacheClient interface {
	GetPreviousExecutionStats(ctx context.Context, in *GetPreviousExecutionStatsRequest, opts ...grpc.CallOption) (*PreviousExecutionStats, error)
	UpdatePreviousExecutionStats(ctx context.Context, in *UpdatePreviousExecutionStatsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type initialSizeClassCacheClient struct {
	cc grpc.ClientConnInterface
}

func NewInitialSizeClassCacheClient(cc grpc.ClientConnInterface) InitialSizeClassCacheClient {
	return &initialSizeClassCacheClient{cc}
}

func (c *initialSizeClassCacheClient) GetPreviousExecutionStats(ctx context.Context, in *GetPreviousExecutionStatsRequest, opts ...grpc.CallOption) (*PreviousExecutionStats, error) {
	out := new(PreviousExecutionStats)
	err := c.cc.Invoke(ctx, "/buildbarn.iscc.InitialSizeClassCache/GetPreviousExecutionStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *initialSizeClassCacheClient) UpdatePreviousExecutionStats(ctx context.Context, in *UpdatePreviousExecutionStatsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/buildbarn.iscc.InitialSizeClassCache/UpdatePreviousExecutionStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InitialSizeClassCacheServer is the server API for InitialSizeClassCache service.
type InitialSizeClassCacheServer interface {
	GetPreviousExecutionStats(context.Context, *GetPreviousExecutionStatsRequest) (*PreviousExecutionStats, error)
	UpdatePreviousExecutionStats(context.Context, *UpdatePreviousExecutionStatsRequest) (*emptypb.Empty, error)
}

// UnimplementedInitialSizeClassCacheServer can be embedded to have forward compatible implementations.
type UnimplementedInitialSizeClassCacheServer struct {
}

func (*UnimplementedInitialSizeClassCacheServer) GetPreviousExecutionStats(context.Context, *GetPreviousExecutionStatsRequest) (*PreviousExecutionStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreviousExecutionStats not implemented")
}
func (*UnimplementedInitialSizeClassCacheServer) UpdatePreviousExecutionStats(context.Context, *UpdatePreviousExecutionStatsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePreviousExecutionStats not implemented")
}

func RegisterInitialSizeClassCacheServer(s grpc.ServiceRegistrar, srv InitialSizeClassCacheServer) {
	s.RegisterService(&_InitialSizeClassCache_serviceDesc, srv)
}

func _InitialSizeClassCache_GetPreviousExecutionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreviousExecutionStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InitialSizeClassCacheServer).GetPreviousExecutionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbarn.iscc.InitialSizeClassCache/GetPreviousExecutionStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InitialSizeClassCacheServer).GetPreviousExecutionStats(ctx, req.(*GetPreviousExecutionStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InitialSizeClassCache_UpdatePreviousExecutionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePreviousExecutionStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InitialSizeClassCacheServer).UpdatePreviousExecutionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbarn.iscc.InitialSizeClassCache/UpdatePreviousExecutionStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InitialSizeClassCacheServer).UpdatePreviousExecutionStats(ctx, req.(*UpdatePreviousExecutionStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InitialSizeClassCache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buildbarn.iscc.InitialSizeClassCache",
	HandlerType: (*InitialSizeClassCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPreviousExecutionStats",
			Handler:    _InitialSizeClassCache_GetPreviousExecutionStats_Handler,
		},
		{
			MethodName: "UpdatePreviousExecutionStats",
			Handler:    _InitialSizeClassCache_UpdatePreviousExecutionStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/iscc/iscc.proto",
}
