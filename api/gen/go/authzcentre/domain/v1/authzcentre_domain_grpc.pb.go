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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: authzcentre_domain.proto

package domain

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthzCentreDomainService_SayHello_FullMethodName = "/tickex.authzcentre.domain.v1.AuthzCentreDomainService/SayHello"
)

// AuthzCentreDomainServiceClient is the client API for AuthzCentreDomainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthzCentreDomainServiceClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type authzCentreDomainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthzCentreDomainServiceClient(cc grpc.ClientConnInterface) AuthzCentreDomainServiceClient {
	return &authzCentreDomainServiceClient{cc}
}

func (c *authzCentreDomainServiceClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, AuthzCentreDomainService_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthzCentreDomainServiceServer is the server API for AuthzCentreDomainService service.
// All implementations must embed UnimplementedAuthzCentreDomainServiceServer
// for forward compatibility.
type AuthzCentreDomainServiceServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	mustEmbedUnimplementedAuthzCentreDomainServiceServer()
}

// UnimplementedAuthzCentreDomainServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthzCentreDomainServiceServer struct{}

func (UnimplementedAuthzCentreDomainServiceServer) SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedAuthzCentreDomainServiceServer) mustEmbedUnimplementedAuthzCentreDomainServiceServer() {
}
func (UnimplementedAuthzCentreDomainServiceServer) testEmbeddedByValue() {}

// UnsafeAuthzCentreDomainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthzCentreDomainServiceServer will
// result in compilation errors.
type UnsafeAuthzCentreDomainServiceServer interface {
	mustEmbedUnimplementedAuthzCentreDomainServiceServer()
}

func RegisterAuthzCentreDomainServiceServer(s grpc.ServiceRegistrar, srv AuthzCentreDomainServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthzCentreDomainServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthzCentreDomainService_ServiceDesc, srv)
}

func _AuthzCentreDomainService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzCentreDomainServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthzCentreDomainService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzCentreDomainServiceServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthzCentreDomainService_ServiceDesc is the grpc.ServiceDesc for AuthzCentreDomainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthzCentreDomainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tickex.authzcentre.domain.v1.AuthzCentreDomainService",
	HandlerType: (*AuthzCentreDomainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _AuthzCentreDomainService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authzcentre_domain.proto",
}
