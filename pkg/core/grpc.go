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

package cestcore

import (
	"context"
	"fmt"

	"github.com/celestinals/celestinal/pkg/core/net"
	cesterr "github.com/celestinals/celestinal/pkg/errors"
	"google.golang.org/grpc"
)

var (
	// Ensure GRPCServer implements ServiceServer.
	_ ServiceServer = (*GRPCServer)(nil)

	// Ensure GRPCServer implements Server.
	_ Server = (*GRPCServer)(nil)
)

// ServiceServer is a gRPC service server.
type ServiceServer interface {
	AsServer() *grpc.Server
	Serve(info *ServiceInfo) error
	Shutdown(ctx context.Context) error
}

// GRPCServer is a gRPC server that registers services.
// inherit in <Service>GRPCServer:
//
//	type Greeter struct {
//		*cestcore.GRPCServer
//		cestconf *types.Config
//		srv    greeter.GreeterServiceServer
//	}
type GRPCServer struct {
	server *grpc.Server
}

// Start implements Server.
func (s *GRPCServer) Start(ctx context.Context) error {
	_ = ctx
	return cesterr.ErrUnimplemented
}

// Shutdown implements ServiceServer.
func (s *GRPCServer) Shutdown(_ context.Context) error {
	s.server.GracefulStop()
	return nil
}

// AsServer returns the underlying gRPC server.
// return the underlying gRPC server.
func (s *GRPCServer) AsServer() *grpc.Server {
	return s.server
}

// Serve starts the http server.
// return error if the http server fails to start.
func (s *GRPCServer) Serve(info *ServiceInfo) error {
	if info == nil {
		return fmt.Errorf("info is nil")
	}

	listener, err := net.ListenNetworkTCP(info.Addr)
	if err != nil {
		return err
	}

	return s.AsServer().Serve(listener)
}

// NewGRPCServer returns a new service registrar.
// opts are the gRPC server options.
func NewGRPCServer(opts ...grpc.ServerOption) *GRPCServer {
	return &GRPCServer{
		server: grpc.NewServer(opts...),
	}
}

// NewGRPCServerDefault returns a new service registrar with default options.
func NewGRPCServerDefault() *GRPCServer {
	return NewGRPCServer()
}
