// Copyright 2025 The Celestinal Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: celestinal/v1/flags.proto

package celestinal

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type Flag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Mode    string `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *Flag) Reset() {
	*x = Flag{}
	mi := &file_celestinal_v1_flags_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Flag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flag) ProtoMessage() {}

func (x *Flag) ProtoReflect() protoreflect.Message {
	mi := &file_celestinal_v1_flags_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flag.ProtoReflect.Descriptor instead.
func (*Flag) Descriptor() ([]byte, []int) {
	return file_celestinal_v1_flags_proto_rawDescGZIP(), []int{0}
}

func (x *Flag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Flag) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Flag) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type FlagEdge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Telegram bool `protobuf:"varint,1,opt,name=telegram,proto3" json:"telegram,omitempty"`
}

func (x *FlagEdge) Reset() {
	*x = FlagEdge{}
	mi := &file_celestinal_v1_flags_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FlagEdge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagEdge) ProtoMessage() {}

func (x *FlagEdge) ProtoReflect() protoreflect.Message {
	mi := &file_celestinal_v1_flags_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagEdge.ProtoReflect.Descriptor instead.
func (*FlagEdge) Descriptor() ([]byte, []int) {
	return file_celestinal_v1_flags_proto_rawDescGZIP(), []int{1}
}

func (x *FlagEdge) GetTelegram() bool {
	if x != nil {
		return x.Telegram
	}
	return false
}

var File_celestinal_v1_flags_proto protoreflect.FileDescriptor

var file_celestinal_v1_flags_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x65, 0x6c,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xba, 0x48, 0x1f, 0x72, 0x1d, 0x32, 0x1b, 0x5e, 0x5b,
	0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2e, 0x2d, 0x5d, 0x2b, 0x3a, 0x5b, 0x30,
	0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x35, 0x7d, 0x24, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1b, 0xba, 0x48, 0x18, 0x72, 0x16, 0x32, 0x14, 0x5e, 0x28, 0x64, 0x65, 0x76, 0x7c, 0x70,
	0x72, 0x6f, 0x64, 0x7c, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x29, 0x24, 0x52, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x22, 0x26, 0x0a, 0x08, 0x46, 0x6c, 0x61, 0x67, 0x45, 0x64, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x47, 0x5a, 0x45, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x65, 0x6c, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_celestinal_v1_flags_proto_rawDescOnce sync.Once
	file_celestinal_v1_flags_proto_rawDescData = file_celestinal_v1_flags_proto_rawDesc
)

func file_celestinal_v1_flags_proto_rawDescGZIP() []byte {
	file_celestinal_v1_flags_proto_rawDescOnce.Do(func() {
		file_celestinal_v1_flags_proto_rawDescData = protoimpl.X.CompressGZIP(file_celestinal_v1_flags_proto_rawDescData)
	})
	return file_celestinal_v1_flags_proto_rawDescData
}

var file_celestinal_v1_flags_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_celestinal_v1_flags_proto_goTypes = []any{
	(*Flag)(nil),     // 0: celestinal.v1.Flag
	(*FlagEdge)(nil), // 1: celestinal.v1.FlagEdge
}
var file_celestinal_v1_flags_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_celestinal_v1_flags_proto_init() }
func file_celestinal_v1_flags_proto_init() {
	if File_celestinal_v1_flags_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_celestinal_v1_flags_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_celestinal_v1_flags_proto_goTypes,
		DependencyIndexes: file_celestinal_v1_flags_proto_depIdxs,
		MessageInfos:      file_celestinal_v1_flags_proto_msgTypes,
	}.Build()
	File_celestinal_v1_flags_proto = out.File
	file_celestinal_v1_flags_proto_rawDesc = nil
	file_celestinal_v1_flags_proto_goTypes = nil
	file_celestinal_v1_flags_proto_depIdxs = nil
}
