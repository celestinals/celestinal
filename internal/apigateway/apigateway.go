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

// Package apigateway provides the apigateway
package apigateway

import (
	"context"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	"github.com/celestinals/celestinal/pkg/striker"
	"github.com/celestinals/celestinal/pkg/striker/skhttp"
	"github.com/celestinals/celestinal/pkg/striker/skutils"

	"github.com/celestinals/celestinal/internal/apigateway/middleware"
	"github.com/celestinals/celestinal/internal/apigateway/services/v1"
	"github.com/celestinals/celestinal/internal/apigateway/utils/openapi"
	"github.com/celestinals/celestinal/internal/apigateway/utils/registry"
	"github.com/celestinals/celestinal/internal/apigateway/utils/watcher"
	"github.com/celestinals/celestinal/internal/pkg/version"

	"github.com/celestinals/celestinal/pkg/logger"
	"github.com/celestinals/celestinal/pkg/protobuf"
)

// make sure apigateway implement striker.Server
// striker.runner will be start application through striker.Server interface
var _ striker.Server = (*Edge)(nil)

// New creates a new gateway app, return striker.Server interface
func New(conf *celestinal.Config) striker.Server {
	return &Edge{
		server: skhttp.New(),
		config: conf,
	}
}

// Edge represents the celestinal app The apigateway application is the main
// entry point for the Celestinal. It will automatically connect to other
// services via gRPC. Run the application along with other services
// in the x/ directory. The application provides APIs for users through
// a single HTTP gateway following the REST API standard. The application
// uses gRPC to connect to other services. Additionally, the system provides
// a Swagger UI interface for users to easily interact with the system
// through a web interface.
type Edge struct {
	// config is the configuration of the apigateway app, load environment
	// variables from .env file
	config *celestinal.Config

	// server is the core server, manage http.ServeMux,
	// runtime.ServeMux and HTTP server
	server skhttp.Server
}

// registerServiceServer gRPC server endpoint.
func (edge *Edge) registerServiceServer(ctx context.Context) error {
	// NOTE: Make sure the gRPC server is running properly and accessible
	// Create folder at services, inherit base package, override function,
	// implement business logic
	// See: services/v1/greeter
	serviceList := []skutils.ServiceRegistrar{
		// Example: register the greeter service to the gateway
		services.NewGreeter(),
		// add more service here ...
	}

	return edge.visit(ctx, serviceList...)
}

// visit all service by Accept function
func (edge *Edge) visit(ctx context.Context, services ...skutils.ServiceRegistrar) error {
	for _, service := range services {
		if err := service.Accept(ctx, edge.server); err != nil {
			return err
		}
	}

	return nil
}

// use is a chain of functions to use when accepting the request
// serve is a function to use when accepting the request
func (edge *Edge) use(serve func(skhttp.Server, *celestinal.Config)) {
	serve(edge.server, edge.config)
}

// functions is chain of functions to use before starting the apigateway app
func (edge *Edge) functions(ctx context.Context) error {
	// new middleware handler
	edge.use(middleware.Serve)

	// serve swagger ui
	edge.use(openapi.Serve)

	// watch service change on service registry
	edge.use(watcher.Serve)

	edge.use(registry.Serve)

	return edge.registerServiceServer(ctx)
}

// Start the apigateway/gateway app
func (edge *Edge) Start(ctx context.Context) error {
	// service ascii art banner
	version.ASCII()
	if err := protobuf.Validate(edge.config); err != nil {
		return err
	}

	// add chain functions handler request
	if err := edge.functions(ctx); err != nil {
		return err
	}

	// Listen HTTP server (and apigateway calls to gRPC server endpoint)
	// log info in console and return register error if they exist
	logger.Infof("[http] starting server %s", edge.config.GetApiAddr())
	return edge.server.Listen(edge.config.GetApiAddr())
	// return errors.F("apigateway: failed to listen and serve")
}

// Shutdown implements striker.Server.
func (edge *Edge) Shutdown(ctx context.Context) error {
	return edge.server.Shutdown(ctx)
}
