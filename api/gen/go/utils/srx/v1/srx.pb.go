// Copyright 2024 The Tickex Authors.

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
// source: srx.proto

package srx

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

// Detailed service information
type RegisterServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`       // Service name
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"` // IP address or hostname
	Port    int32  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`      // Service port
}

func (x *RegisterServiceRequest) Reset() {
	*x = RegisterServiceRequest{}
	mi := &file_srx_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterServiceRequest) ProtoMessage() {}

func (x *RegisterServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_srx_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterServiceRequest.ProtoReflect.Descriptor instead.
func (*RegisterServiceRequest) Descriptor() ([]byte, []int) {
	return file_srx_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterServiceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegisterServiceRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`       // Service name
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"` // IP address or hostname
	Port    int32  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`      // Service port
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	mi := &file_srx_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_srx_proto_msgTypes[1]
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
	return file_srx_proto_rawDescGZIP(), []int{1}
}

func (x *HeartbeatRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HeartbeatRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *HeartbeatRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// Request to find a service
type DiscoverServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Name of the service to discover
}

func (x *DiscoverServiceRequest) Reset() {
	*x = DiscoverServiceRequest{}
	mi := &file_srx_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoverServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverServiceRequest) ProtoMessage() {}

func (x *DiscoverServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_srx_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverServiceRequest.ProtoReflect.Descriptor instead.
func (*DiscoverServiceRequest) Descriptor() ([]byte, []int) {
	return file_srx_proto_rawDescGZIP(), []int{2}
}

func (x *DiscoverServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Response from the discovery process
type DiscoverServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*RegisterServiceRequest `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"` // List of service instances
}

func (x *DiscoverServiceResponse) Reset() {
	*x = DiscoverServiceResponse{}
	mi := &file_srx_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoverServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverServiceResponse) ProtoMessage() {}

func (x *DiscoverServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_srx_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverServiceResponse.ProtoReflect.Descriptor instead.
func (*DiscoverServiceResponse) Descriptor() ([]byte, []int) {
	return file_srx_proto_rawDescGZIP(), []int{3}
}

func (x *DiscoverServiceResponse) GetServices() []*RegisterServiceRequest {
	if x != nil {
		return x.Services
	}
	return nil
}

// Service registration response
type RegisterServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // Result message
}

func (x *RegisterServiceResponse) Reset() {
	*x = RegisterServiceResponse{}
	mi := &file_srx_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterServiceResponse) ProtoMessage() {}

func (x *RegisterServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_srx_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterServiceResponse.ProtoReflect.Descriptor instead.
func (*RegisterServiceResponse) Descriptor() ([]byte, []int) {
	return file_srx_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterServiceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Heartbeat response
type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Heartbeat status
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	mi := &file_srx_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_srx_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_srx_proto_rawDescGZIP(), []int{5}
}

func (x *HeartbeatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_srx_proto protoreflect.FileDescriptor

var file_srx_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x72, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x78, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x73, 0x72, 0x78, 0x2e, 0x76, 0x31,
	0x22, 0x5a, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x54, 0x0a, 0x10,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0x2c, 0x0a, 0x16, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x62, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x73, 0x72, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x33, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xd0, 0x02, 0x0a, 0x16, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e,
	0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x73, 0x72, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x75, 0x74, 0x69,
	0x6c, 0x73, 0x2e, 0x73, 0x72, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6c, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x75, 0x74,
	0x69, 0x6c, 0x73, 0x2e, 0x73, 0x72, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73,
	0x2e, 0x73, 0x72, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5a, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x25, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x73, 0x72, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x75, 0x74, 0x69,
	0x6c, 0x73, 0x2e, 0x73, 0x72, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78,
	0x76, 0x6e, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x73, 0x72, 0x78, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x72, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_srx_proto_rawDescOnce sync.Once
	file_srx_proto_rawDescData = file_srx_proto_rawDesc
)

func file_srx_proto_rawDescGZIP() []byte {
	file_srx_proto_rawDescOnce.Do(func() {
		file_srx_proto_rawDescData = protoimpl.X.CompressGZIP(file_srx_proto_rawDescData)
	})
	return file_srx_proto_rawDescData
}

var file_srx_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_srx_proto_goTypes = []any{
	(*RegisterServiceRequest)(nil),  // 0: tickex.utils.srx.v1.RegisterServiceRequest
	(*HeartbeatRequest)(nil),        // 1: tickex.utils.srx.v1.HeartbeatRequest
	(*DiscoverServiceRequest)(nil),  // 2: tickex.utils.srx.v1.DiscoverServiceRequest
	(*DiscoverServiceResponse)(nil), // 3: tickex.utils.srx.v1.DiscoverServiceResponse
	(*RegisterServiceResponse)(nil), // 4: tickex.utils.srx.v1.RegisterServiceResponse
	(*HeartbeatResponse)(nil),       // 5: tickex.utils.srx.v1.HeartbeatResponse
}
var file_srx_proto_depIdxs = []int32{
	0, // 0: tickex.utils.srx.v1.DiscoverServiceResponse.services:type_name -> tickex.utils.srx.v1.RegisterServiceRequest
	0, // 1: tickex.utils.srx.v1.ServiceRegistryService.RegisterService:input_type -> tickex.utils.srx.v1.RegisterServiceRequest
	2, // 2: tickex.utils.srx.v1.ServiceRegistryService.DiscoverService:input_type -> tickex.utils.srx.v1.DiscoverServiceRequest
	1, // 3: tickex.utils.srx.v1.ServiceRegistryService.Heartbeat:input_type -> tickex.utils.srx.v1.HeartbeatRequest
	4, // 4: tickex.utils.srx.v1.ServiceRegistryService.RegisterService:output_type -> tickex.utils.srx.v1.RegisterServiceResponse
	3, // 5: tickex.utils.srx.v1.ServiceRegistryService.DiscoverService:output_type -> tickex.utils.srx.v1.DiscoverServiceResponse
	5, // 6: tickex.utils.srx.v1.ServiceRegistryService.Heartbeat:output_type -> tickex.utils.srx.v1.HeartbeatResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_srx_proto_init() }
func file_srx_proto_init() {
	if File_srx_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_srx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_srx_proto_goTypes,
		DependencyIndexes: file_srx_proto_depIdxs,
		MessageInfos:      file_srx_proto_msgTypes,
	}.Build()
	File_srx_proto = out.File
	file_srx_proto_rawDesc = nil
	file_srx_proto_goTypes = nil
	file_srx_proto_depIdxs = nil
}
