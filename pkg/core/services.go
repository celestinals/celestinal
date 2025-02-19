/*
 * Copyright 2025 The Tickex Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// GRPCService is an interface for registering a gRPC service.
type GRPCService interface {
	Register(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
}

// IServiceServer is a service registrar.
type IServiceServer interface {
	AsServer() *grpc.Server
	Run() error
}

var _ IServiceServer = (*ServiceServer)(nil)

// ServiceServer is a gRPC server that registers services.
type ServiceServer struct {
	server *grpc.Server
}

// AsServer returns the underlying gRPC server.
func (s *ServiceServer) AsServer() *grpc.Server {
	return s.server
}

// Run starts the service registrar.
func (s *ServiceServer) Run() error {
	panic("unimplemented")
}

// New returns a new service registrar.
func New(opts ...grpc.ServerOption) *ServiceServer {
	return &ServiceServer{
		server: grpc.NewServer(opts...),
	}
}

// NewDefault returns a new service registrar with default options.
func NewDefault() *ServiceServer {
	return New()
}
