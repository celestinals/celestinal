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

// Package apiserver provides the apiserver
package apiserver

import (
	"context"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"

	"github.com/celestinals/celestinal/internal/apiserver/apps/openapi"
	"github.com/celestinals/celestinal/internal/apiserver/middleware"
	"github.com/celestinals/celestinal/internal/apiserver/registrar/v1"
	dcvrcontrollers "github.com/celestinals/celestinal/internal/discovery/v1/controllers"
	"github.com/celestinals/celestinal/internal/pkg/version"

	"github.com/celestinals/celestinal/pkg/flag"
	"github.com/celestinals/celestinal/pkg/logger"
	"github.com/celestinals/celestinal/pkg/protobuf"
	"github.com/celestinals/celestinal/pkg/striker"
	"github.com/celestinals/celestinal/pkg/striker/skhttp"
	"github.com/celestinals/celestinal/pkg/striker/skutils"
)

// make sure apiserver implement striker.Server
// striker.runner will be start application through striker.Server interface
var _ striker.Server = (*Server)(nil)

// New creates a new gateway app and returns a striker.Server interface.
// This constructor is based on dependency injection. When you add parameters
// (e.g., svc discoverysvc.Discovery), you must use the striker.Inject
// to inject the constructor of the object into the striker framework.
// See service/discoverysvc/discoverysvc.go for reference.
//
// Example:
//
//	var _ = striker.Inject(discoverysvc.NewDiscoveryService)
func New(conf *celestinal.Config, dcv *dcvrcontrollers.Discovery) (striker.Server, error) {
	srv := &Server{
		server: skhttp.New(),
		config: conf,
	}

	// handler custom
	srv.use(openapi.New())
	srv.use(middleware.New(conf))

	// NOTE: Make sure the gRPC server is running properly and accessible
	// Create file at registrar, inherit base package, override function,
	// implement business logic
	// See: registrar/v1/greeter
	err := srv.visit(context.Background(),
		registrar.NewGreeter(),
		registrar.NewDiscovery(dcv),
	)

	return srv, err
}

// Handler interface must be implemented by object handler
//
// References:
//   - openapi
//   - middleware
type Handler interface {
	RegisterServer(server skhttp.Server, conf *celestinal.Config)
}

// Server represents the celestinal app The apiserver application is the main
// entry point for the Celestinal. It will automatically connect to other
// services via gRPC. Run the application along with other services
// in the cmd/ directory. The application provides APIs for users through
// a single HTTP gateway following the REST API standard. The application
// uses gRPC to connect to other services. Additionally, the system provides
// a Swagger UI interface for users to easily interact with the system
// through a web interface.
type Server struct {
	// config is the configuration of the apiserver app, load environment
	// variables from .env file
	config *celestinal.Config

	// server is the core server, manage http.ServeMux,
	// runtime.ServeMux and HTTP server
	server skhttp.Server
}

// visit all service by Accept function
func (srv *Server) visit(ctx context.Context, services ...skutils.ServiceRegistrar) error {
	for _, service := range services {
		if err := service.Accept(ctx, srv.server); err != nil {
			return err
		}
	}
	return nil
	// return errors.F("apiserver: failed to visit service")
}

// use is a chain of functions to use when accepting the request
// serve is a function to use when accepting the request
func (srv *Server) use(handler Handler) {
	handler.RegisterServer(srv.server, srv.config)
}

// Start the apiserver/gateway app
func (srv *Server) Start(_ context.Context) error {
	// service ascii art banner
	version.ASCII()
	if err := protobuf.Validate(srv.config); err != nil {
		return err
	}

	// Listen HTTP server (and apiserver calls to gRPC server endpoint)
	// log info in console and return register error if they exist
	logger.Infof("[http] starting server %s", flag.ParseAPIServer().GetAddress())
	return srv.server.Listen(flag.ParseAPIServer().GetAddress())
	// for DEBUG:
	// return errors.F("apiserver: failed to listen and serve")
}

// Shutdown implements striker.Server.
func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.server.Shutdown(ctx)
}
