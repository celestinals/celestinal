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
	"fmt"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	discoverypb "github.com/tickexvn/tickex/api/gen/go/common/discovery/v1"
	"github.com/tickexvn/tickex/api/gen/go/common/env/config/v1"
	servicepb "github.com/tickexvn/tickex/api/gen/go/common/service/v1"
	"github.com/tickexvn/tickex/pkg/core/net"
	"github.com/tickexvn/tickex/pkg/discovery"
	"github.com/tickexvn/tickex/pkg/logger"
	"google.golang.org/grpc"
)

var _ IServiceServer = (*ServiceServer)(nil)

// GRPCService is an interface for registering a gRPC service.
//
// usage:
//
//	type IService interface {
//		core.GRPCService
//	}
type GRPCService interface {
	Register(ctx context.Context,
		mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
}

// IServiceServer is a gRPC service server.
type IServiceServer interface {
	AsServer() *grpc.Server
	Serve(info *ServiceInfo) error
}

// service is register Properties
type service struct {
	Host string
	Port uint32
	Name string
	Tags []string
}

// ServiceInfo is Serve method properties
type ServiceInfo struct {
	Config *config.Config
	Addr   string
	Tags   []string
	Name   string
}

// ServiceServer is a gRPC server that registers services.
//
// usage:
//
//	type Greeter struct {
//		*core.ServiceServer
//		config *types.Config
//		srv    greeter.GreeterServiceServer
//	}
type ServiceServer struct {
	server    *grpc.Server
	discovery discoverypb.DiscoveryServiceServer
}

// AsServer returns the underlying gRPC server.
// return the underlying gRPC server.
func (s *ServiceServer) AsServer() *grpc.Server {
	return s.server
}

// Serve starts the http server.
// return error if the http server fails to start.
func (s *ServiceServer) Serve(info *ServiceInfo) error {
	if info == nil {
		return fmt.Errorf("info is nil")
	}

	listener, err := net.ListenNetworkTCP(info.Addr)
	if err != nil {
		return err
	}

	host, port, err := net.SplitHostPortListener(listener)
	if err != nil {
		return err
	}

	if err := s.register(info.Config, service{
		Host: host,
		Port: port,
		Name: info.Name,
		Tags: info.Tags,
	}); err != nil {
		return err
	}

	return s.AsServer().Serve(listener)
}

// register registers the service with the service discovery.
func (s *ServiceServer) register(conf *config.Config, service service) error {
	discover, err := discovery.New(conf)
	if err != nil {
		return err
	}

	s.discovery = discover
	return s.registerConsul(service)
}

func (s *ServiceServer) registerConsul(service service) error {
	if s.discovery == nil {
		return nil
	}

	serviceID := service.Name
	ttl := time.Second * 5

	if _, err := s.discovery.Register(
		context.Background(), &discoverypb.RegisterRequest{
			ServiceCheck: &discoverypb.ServiceCheck{
				Ttl:                            ttl.String(),
				TlsSkipVerify:                  true,
				DeregisterCriticalServiceAfter: ttl.String(),
			},
			Service: &servicepb.Service{
				Id:   serviceID,
				Name: service.Name,
				Host: service.Host,
				Port: service.Port,
				Tags: service.Tags,
			},
		}); err != nil {
		return err
	}

	go s.heartbeat(serviceID, ttl)

	return nil
}

func (s *ServiceServer) heartbeat(id string, ttl time.Duration) {
	if s.discovery == nil {
		return
	}

	ticker := time.NewTicker(ttl)
	for {
		_, err := s.discovery.Heartbeat(
			context.Background(), &discoverypb.HeartbeatRequest{Id: id})
		if err != nil {
			logger.Errorf("consul heartbeat error: %v", err)
		}

		<-ticker.C
	}

}

// New returns a new service registrar.
// opts are the gRPC server options.
func New(opts ...grpc.ServerOption) *ServiceServer {
	return &ServiceServer{
		server: grpc.NewServer(opts...),
	}
}

// NewDefault returns a new service registrar with default options.
func NewDefault() *ServiceServer {
	return New()
}
