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
// source: celestinal/v1/errors.proto

package celestinal

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

type Errors int32

const (
	Errors_ERRORS_UNSPECIFIED    Errors = 0
	Errors_ERRORS_INTERNAL_ERROR Errors = 1
	Errors_ERRORS_NOT_FOUND      Errors = 2
	Errors_ERRORS_UNAUTHORIZED   Errors = 3
	Errors_ERRORS_FORBIDDEN      Errors = 4
	Errors_ERRORS_INVALID_DATA   Errors = 5
	Errors_ERRORS_UNIMPLEMENTED  Errors = 6
)

// Enum value maps for Errors.
var (
	Errors_name = map[int32]string{
		0: "ERRORS_UNSPECIFIED",
		1: "ERRORS_INTERNAL_ERROR",
		2: "ERRORS_NOT_FOUND",
		3: "ERRORS_UNAUTHORIZED",
		4: "ERRORS_FORBIDDEN",
		5: "ERRORS_INVALID_DATA",
		6: "ERRORS_UNIMPLEMENTED",
	}
	Errors_value = map[string]int32{
		"ERRORS_UNSPECIFIED":    0,
		"ERRORS_INTERNAL_ERROR": 1,
		"ERRORS_NOT_FOUND":      2,
		"ERRORS_UNAUTHORIZED":   3,
		"ERRORS_FORBIDDEN":      4,
		"ERRORS_INVALID_DATA":   5,
		"ERRORS_UNIMPLEMENTED":  6,
	}
)

func (x Errors) Enum() *Errors {
	p := new(Errors)
	*p = x
	return p
}

func (x Errors) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Errors) Descriptor() protoreflect.EnumDescriptor {
	return file_celestinal_v1_errors_proto_enumTypes[0].Descriptor()
}

func (Errors) Type() protoreflect.EnumType {
	return &file_celestinal_v1_errors_proto_enumTypes[0]
}

func (x Errors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Errors.Descriptor instead.
func (Errors) EnumDescriptor() ([]byte, []int) {
	return file_celestinal_v1_errors_proto_rawDescGZIP(), []int{0}
}

var File_celestinal_v1_errors_proto protoreflect.FileDescriptor

var file_celestinal_v1_errors_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x65,
	0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2a, 0xb3, 0x01, 0x0a, 0x06,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19,
	0x0a, 0x15, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12,
	0x17, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53, 0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48,
	0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x04, 0x12, 0x17,
	0x0a, 0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x53, 0x5f, 0x55, 0x4e, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10,
	0x06, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x63, 0x65, 0x6c, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_celestinal_v1_errors_proto_rawDescOnce sync.Once
	file_celestinal_v1_errors_proto_rawDescData = file_celestinal_v1_errors_proto_rawDesc
)

func file_celestinal_v1_errors_proto_rawDescGZIP() []byte {
	file_celestinal_v1_errors_proto_rawDescOnce.Do(func() {
		file_celestinal_v1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_celestinal_v1_errors_proto_rawDescData)
	})
	return file_celestinal_v1_errors_proto_rawDescData
}

var file_celestinal_v1_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_celestinal_v1_errors_proto_goTypes = []any{
	(Errors)(0), // 0: celestinal.v1.Errors
}
var file_celestinal_v1_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_celestinal_v1_errors_proto_init() }
func file_celestinal_v1_errors_proto_init() {
	if File_celestinal_v1_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_celestinal_v1_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_celestinal_v1_errors_proto_goTypes,
		DependencyIndexes: file_celestinal_v1_errors_proto_depIdxs,
		EnumInfos:         file_celestinal_v1_errors_proto_enumTypes,
	}.Build()
	File_celestinal_v1_errors_proto = out.File
	file_celestinal_v1_errors_proto_rawDesc = nil
	file_celestinal_v1_errors_proto_goTypes = nil
	file_celestinal_v1_errors_proto_depIdxs = nil
}
