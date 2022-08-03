// Copyright 2018 The Goma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: command/package_opts.proto

package command

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PackageOpts is a package option.
// NEXT_ID_TO_USE: 8
type PackageOpts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// emulation_command specifies which emulation layer is necessary for this package.
	// If empty, it means no emulation layer is necessary.
	EmulationCommand string `protobuf:"bytes,7,opt,name=emulation_command,json=emulationCommand,proto3" json:"emulation_command,omitempty"`
	// output_file_filters is a set of regexp to filter output files.
	OutputFileFilters []string `protobuf:"bytes,5,rep,name=output_file_filters,json=outputFileFilters,proto3" json:"output_file_filters,omitempty"`
}

func (x *PackageOpts) Reset() {
	*x = PackageOpts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_package_opts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageOpts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageOpts) ProtoMessage() {}

func (x *PackageOpts) ProtoReflect() protoreflect.Message {
	mi := &file_command_package_opts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageOpts.ProtoReflect.Descriptor instead.
func (*PackageOpts) Descriptor() ([]byte, []int) {
	return file_command_package_opts_proto_rawDescGZIP(), []int{0}
}

func (x *PackageOpts) GetEmulationCommand() string {
	if x != nil {
		return x.EmulationCommand
	}
	return ""
}

func (x *PackageOpts) GetOutputFileFilters() []string {
	if x != nil {
		return x.OutputFileFilters
	}
	return nil
}

var File_command_package_opts_proto protoreflect.FileDescriptor

var file_command_package_opts_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x65, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x11, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04,
	0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x6f, 0x6d, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_package_opts_proto_rawDescOnce sync.Once
	file_command_package_opts_proto_rawDescData = file_command_package_opts_proto_rawDesc
)

func file_command_package_opts_proto_rawDescGZIP() []byte {
	file_command_package_opts_proto_rawDescOnce.Do(func() {
		file_command_package_opts_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_package_opts_proto_rawDescData)
	})
	return file_command_package_opts_proto_rawDescData
}

var file_command_package_opts_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_command_package_opts_proto_goTypes = []interface{}{
	(*PackageOpts)(nil), // 0: command.PackageOpts
}
var file_command_package_opts_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_command_package_opts_proto_init() }
func file_command_package_opts_proto_init() {
	if File_command_package_opts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_command_package_opts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageOpts); i {
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
			RawDescriptor: file_command_package_opts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_command_package_opts_proto_goTypes,
		DependencyIndexes: file_command_package_opts_proto_depIdxs,
		MessageInfos:      file_command_package_opts_proto_msgTypes,
	}.Build()
	File_command_package_opts_proto = out.File
	file_command_package_opts_proto_rawDesc = nil
	file_command_package_opts_proto_goTypes = nil
	file_command_package_opts_proto_depIdxs = nil
}