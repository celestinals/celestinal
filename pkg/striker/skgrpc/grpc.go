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

// Package skgrpc provides a gRPC server for the celestinal.
package skgrpc

import (
	"context"
	"fmt"

	"github.com/celestinals/celestinal/api/discovery/v1"
	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	"github.com/celestinals/celestinal/pkg/errors"
	"github.com/celestinals/celestinal/pkg/logger"
	"github.com/celestinals/celestinal/pkg/protobuf"
	"github.com/celestinals/celestinal/pkg/striker"
	"github.com/celestinals/celestinal/pkg/striker/sknet"
	"google.golang.org/grpc"
)

var (
	// Ensure Server implements ServiceServer.
	_ ServiceServer = (*Server)(nil)

	// Ensure Server implements Server.
	_ striker.Server = (*Server)(nil)
)

// ServiceServer is a gRPC service server.
type ServiceServer interface {
	AsServer() *grpc.Server
	Serve(info *striker.ServiceInfo) error
	Shutdown(ctx context.Context) error
}

// Server is a gRPC server that registers services.
// inherit in <Service>Server:
//
//	type Greeter struct {
//		*core.Server
//		config *types.Config
//		srv    greeter.GreeterServiceServer
//	}
type Server struct {
	server *grpc.Server
}

// Start implements Server.
func (s *Server) Start(ctx context.Context) error {
	_ = ctx
	return errors.ErrUnimplemented
}

// Shutdown implements ServiceServer.
func (s *Server) Shutdown(_ context.Context) error {
	s.server.GracefulStop()
	return nil
}

// AsServer returns the underlying gRPC server.
// return the underlying gRPC server.
func (s *Server) AsServer() *grpc.Server {
	return s.server
}

// Serve starts the http server.
// return error if the http server fails to start.
func (s *Server) Serve(info *striker.ServiceInfo) error {
	if info == nil {
		return fmt.Errorf("info is nil")
	}

	listener, err := sknet.ListenNetworkTCP(info.Addr)
	if err != nil {
		return err
	}

	host, port, err := sknet.SplitHostPortListener(listener)
	if err != nil {
		return err
	}

	if err := discovery.New(info.GatewayAddr).Register(context.Background(),
		&celestinal.RegisterRequest{
			Name:    info.Name,
			Address: fmt.Sprintf("%s:%d", host, port),
			Ttl:     protobuf.ToDuration(info.TTL),
		}); err != nil {

		logger.Errorf("GRPC.Serve: error when register %v", err)
	}

	return s.AsServer().Serve(listener)
}

// New returns a new service registrar.
// opts are the gRPC server options.
func New(opts ...grpc.ServerOption) *Server {
	return &Server{
		server: grpc.NewServer(opts...),
	}
}

// NewDefault returns a new service registrar with default options.
func NewDefault() *Server {
	return New()
}
