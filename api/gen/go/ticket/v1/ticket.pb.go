// Copyright 2025 The Tickex Authors.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: ticket.proto

package ticket

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	v11 "github.com/tickexvn/tickex/api/gen/go/ticket/shared/v1"
	v1 "github.com/tickexvn/tickex/api/gen/go/types/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *v1.Pages `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Name string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	mi := &file_ticket_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloRequest) GetPage() *v1.Pages {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *SayHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *v11.SayHelloResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	mi := &file_ticket_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetResponse() *v11.SayHelloResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	mi := &file_ticket_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *StatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	mi := &file_ticket_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{3}
}

var File_ticket_proto protoreflect.FileDescriptor

var file_ticket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x0f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x59, 0x0a, 0x10, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x0a, 0x0e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xce, 0x01, 0x0a, 0x0d, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x08, 0x53,
	0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x21, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x78, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61,
	0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x06, 0x12, 0x04, 0x2f, 0x73, 0x61, 0x79, 0x12, 0x5c, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x09, 0x12, 0x07, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x9f, 0x01, 0x92, 0x41, 0x64,
	0x12, 0x15, 0x0a, 0x0e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x02, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x23, 0x0a,
	0x21, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x13, 0x08,
	0x02, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x02, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x00, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x76, 0x6e, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticket_proto_rawDescOnce sync.Once
	file_ticket_proto_rawDescData = file_ticket_proto_rawDesc
)

func file_ticket_proto_rawDescGZIP() []byte {
	file_ticket_proto_rawDescOnce.Do(func() {
		file_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticket_proto_rawDescData)
	})
	return file_ticket_proto_rawDescData
}

var file_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ticket_proto_goTypes = []any{
	(*SayHelloRequest)(nil),      // 0: tickex.ticket.v1.SayHelloRequest
	(*SayHelloResponse)(nil),     // 1: tickex.ticket.v1.SayHelloResponse
	(*StatusResponse)(nil),       // 2: tickex.ticket.v1.StatusResponse
	(*StatusRequest)(nil),        // 3: tickex.ticket.v1.StatusRequest
	(*v1.Pages)(nil),             // 4: tickex.types.v1.Pages
	(*v11.SayHelloResponse)(nil), // 5: tickex.ticket.shared.v1.SayHelloResponse
}
var file_ticket_proto_depIdxs = []int32{
	4, // 0: tickex.ticket.v1.SayHelloRequest.page:type_name -> tickex.types.v1.Pages
	5, // 1: tickex.ticket.v1.SayHelloResponse.response:type_name -> tickex.ticket.shared.v1.SayHelloResponse
	0, // 2: tickex.ticket.v1.TicketService.SayHello:input_type -> tickex.ticket.v1.SayHelloRequest
	3, // 3: tickex.ticket.v1.TicketService.Status:input_type -> tickex.ticket.v1.StatusRequest
	1, // 4: tickex.ticket.v1.TicketService.SayHello:output_type -> tickex.ticket.v1.SayHelloResponse
	2, // 5: tickex.ticket.v1.TicketService.Status:output_type -> tickex.ticket.v1.StatusResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ticket_proto_init() }
func file_ticket_proto_init() {
	if File_ticket_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticket_proto_goTypes,
		DependencyIndexes: file_ticket_proto_depIdxs,
		MessageInfos:      file_ticket_proto_msgTypes,
	}.Build()
	File_ticket_proto = out.File
	file_ticket_proto_rawDesc = nil
	file_ticket_proto_goTypes = nil
	file_ticket_proto_depIdxs = nil
}
