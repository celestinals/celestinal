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
// source: celestinal/v1/discovery.proto

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

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	mi := &file_celestinal_v1_discovery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_celestinal_v1_discovery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_celestinal_v1_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *HeartbeatRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Ttl     int64  `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_celestinal_v1_discovery_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_celestinal_v1_discovery_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_celestinal_v1_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegisterRequest) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type DiscoverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DiscoverRequest) Reset() {
	*x = DiscoverRequest{}
	mi := &file_celestinal_v1_discovery_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverRequest) ProtoMessage() {}

func (x *DiscoverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_celestinal_v1_discovery_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverRequest.ProtoReflect.Descriptor instead.
func (*DiscoverRequest) Descriptor() ([]byte, []int) {
	return file_celestinal_v1_discovery_proto_rawDescGZIP(), []int{2}
}

func (x *DiscoverRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DiscoverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *DiscoverResponse) Reset() {
	*x = DiscoverResponse{}
	mi := &file_celestinal_v1_discovery_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverResponse) ProtoMessage() {}

func (x *DiscoverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_celestinal_v1_discovery_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverResponse.ProtoReflect.Descriptor instead.
func (*DiscoverResponse) Descriptor() ([]byte, []int) {
	return file_celestinal_v1_discovery_proto_rawDescGZIP(), []int{3}
}

func (x *DiscoverResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DiscoverResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_celestinal_v1_discovery_proto protoreflect.FileDescriptor

var file_celestinal_v1_discovery_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x10, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xa8, 0x01, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x7a, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xba, 0x48, 0x11, 0x72, 0x0f, 0x32, 0x0d, 0x5e, 0x63, 0x65,
	0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x5c, 0x2e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xa8, 0x01, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x07, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22,
	0x3b, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x14, 0xba, 0x48, 0x11, 0x72, 0x0f, 0x32, 0x0d, 0x5e, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x6c, 0x5c, 0x2e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x60, 0x0a, 0x10,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14,
	0xba, 0x48, 0x11, 0x72, 0x0f, 0x32, 0x0d, 0x5e, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x6c, 0x5c, 0x2e, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05,
	0x72, 0x03, 0xa8, 0x01, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x47,
	0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x6c,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x65, 0x6c,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_celestinal_v1_discovery_proto_rawDescOnce sync.Once
	file_celestinal_v1_discovery_proto_rawDescData = file_celestinal_v1_discovery_proto_rawDesc
)

func file_celestinal_v1_discovery_proto_rawDescGZIP() []byte {
	file_celestinal_v1_discovery_proto_rawDescOnce.Do(func() {
		file_celestinal_v1_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_celestinal_v1_discovery_proto_rawDescData)
	})
	return file_celestinal_v1_discovery_proto_rawDescData
}

var file_celestinal_v1_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_celestinal_v1_discovery_proto_goTypes = []any{
	(*HeartbeatRequest)(nil), // 0: celestinal.v1.HeartbeatRequest
	(*RegisterRequest)(nil),  // 1: celestinal.v1.RegisterRequest
	(*DiscoverRequest)(nil),  // 2: celestinal.v1.DiscoverRequest
	(*DiscoverResponse)(nil), // 3: celestinal.v1.DiscoverResponse
}
var file_celestinal_v1_discovery_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_celestinal_v1_discovery_proto_init() }
func file_celestinal_v1_discovery_proto_init() {
	if File_celestinal_v1_discovery_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_celestinal_v1_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_celestinal_v1_discovery_proto_goTypes,
		DependencyIndexes: file_celestinal_v1_discovery_proto_depIdxs,
		MessageInfos:      file_celestinal_v1_discovery_proto_msgTypes,
	}.Build()
	File_celestinal_v1_discovery_proto = out.File
	file_celestinal_v1_discovery_proto_rawDesc = nil
	file_celestinal_v1_discovery_proto_goTypes = nil
	file_celestinal_v1_discovery_proto_depIdxs = nil
}
