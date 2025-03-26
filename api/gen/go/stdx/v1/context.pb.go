// Copyright 2025 The Tickex Authors.
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
// source: tickex/stdx/v1/context.proto

package stdx

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

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies the user making the request.
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Passes tokens or authentication information through the context.
	Authorization string `protobuf:"bytes,4,opt,name=authorization,proto3" json:"authorization,omitempty"`
	// Passes language information for multilingual systems.
	Locale string `protobuf:"bytes,5,opt,name=locale,proto3" json:"locale,omitempty"`
	// Identifies the current service processing the request.
	ServiceName string `protobuf:"bytes,7,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Identifies the IP address of the client sending the request.
	Ip string `protobuf:"bytes,10,opt,name=ip,proto3" json:"ip,omitempty"`
	// Identifies the environment in which the service is running.
	Environment string `protobuf:"bytes,11,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	mi := &file_tickex_stdx_v1_context_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_tickex_stdx_v1_context_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_tickex_stdx_v1_context_proto_rawDescGZIP(), []int{0}
}

func (x *Context) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Context) GetAuthorization() string {
	if x != nil {
		return x.Authorization
	}
	return ""
}

func (x *Context) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Context) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Context) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Context) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

var File_tickex_stdx_v1_context_proto protoreflect.FileDescriptor

var file_tickex_stdx_v1_context_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x73, 0x74, 0x64, 0x78, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x73, 0x74, 0x64, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x51, 0x0a, 0x0d, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x2b, 0xba, 0x48, 0x28, 0x72, 0x26, 0x32, 0x24, 0x5e, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x5b, 0x20, 0x09, 0x5d, 0x2b, 0x28, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5c, 0x2d, 0x2e, 0x5f, 0x7e, 0x2b, 0x2f, 0x5d, 0x2b, 0x3d, 0x2a, 0x29, 0x24, 0x52, 0x0d, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x06,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x36, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x13, 0xba, 0x48, 0x10, 0x72, 0x0e, 0x32, 0x0c, 0x5e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x78, 0x5c, 0x2e, 0x2e, 0x2a, 0x24, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x19, 0xba, 0x48, 0x16, 0x72, 0x14, 0x32, 0x12, 0x5e, 0x28, 0x5b, 0x61, 0x2d, 0x7a,
	0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2e, 0x2d, 0x5d, 0x2b, 0x29, 0x24, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x3d, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xba, 0x48, 0x18, 0x72, 0x16, 0x32, 0x14, 0x5e, 0x28,
	0x64, 0x65, 0x76, 0x7c, 0x70, 0x72, 0x6f, 0x64, 0x7c, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78,
	0x29, 0x24, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x42,
	0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x78, 0x76, 0x6e, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x64, 0x78, 0x2f, 0x76, 0x31,
	0x3b, 0x73, 0x74, 0x64, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tickex_stdx_v1_context_proto_rawDescOnce sync.Once
	file_tickex_stdx_v1_context_proto_rawDescData = file_tickex_stdx_v1_context_proto_rawDesc
)

func file_tickex_stdx_v1_context_proto_rawDescGZIP() []byte {
	file_tickex_stdx_v1_context_proto_rawDescOnce.Do(func() {
		file_tickex_stdx_v1_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_tickex_stdx_v1_context_proto_rawDescData)
	})
	return file_tickex_stdx_v1_context_proto_rawDescData
}

var file_tickex_stdx_v1_context_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tickex_stdx_v1_context_proto_goTypes = []any{
	(*Context)(nil), // 0: tickex.stdx.v1.Context
}
var file_tickex_stdx_v1_context_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tickex_stdx_v1_context_proto_init() }
func file_tickex_stdx_v1_context_proto_init() {
	if File_tickex_stdx_v1_context_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tickex_stdx_v1_context_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tickex_stdx_v1_context_proto_goTypes,
		DependencyIndexes: file_tickex_stdx_v1_context_proto_depIdxs,
		MessageInfos:      file_tickex_stdx_v1_context_proto_msgTypes,
	}.Build()
	File_tickex_stdx_v1_context_proto = out.File
	file_tickex_stdx_v1_context_proto_rawDesc = nil
	file_tickex_stdx_v1_context_proto_goTypes = nil
	file_tickex_stdx_v1_context_proto_depIdxs = nil
}
