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

// Package edge provides the tickex-edge
package edge

import (
	"context"

	"github.com/tickexvn/tickex/api/gen/go/stdx/v1"
	"github.com/tickexvn/tickex/internal/edge/services/v1"
	"github.com/tickexvn/tickex/internal/funcs/middleware"
	"github.com/tickexvn/tickex/internal/funcs/openapi"
	"github.com/tickexvn/tickex/internal/funcs/secure"
	"github.com/tickexvn/tickex/internal/funcs/watch"
	"github.com/tickexvn/tickex/internal/utils/version"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/protobuf"
	"github.com/tickexvn/tickex/pkg/txlog"
)

var _ core.Server = (*Edge)(nil)

// New creates a new gateway app
func New(conf *stdx.Config) core.Server {
	return &Edge{
		server: core.NewHTTPServer(),
		config: conf,
	}
}

// Edge represents the tickex app The edge application is the main
// entry point for the Tickex. It will automatically connect to other
// services via gRPC. Start the application along with other services
// in the x/ directory. The application provides APIs for users through
// a single HTTP gateway following the REST API standard. The application
// uses gRPC to connect to other services. Additionally, the system provides
// a Swagger UI interface for users to easily interact with the system
// through a web interface.
type Edge struct {
	// config is the configuration of the edge app, load environment
	// variables from .env file
	config *stdx.Config

	// server is the core server server, manage http.ServeMux,
	// runtime.ServeMux and HTTP server
	server core.HTTPServer
}

// registerServer gRPC server endpoint.
func (e *Edge) registerServer(ctx context.Context) error {
	// Note: Make sure the gRPC server is running properly and accessible
	// Create folder at services, inherit base package, override function,
	// implement business logic
	// See: services/v1/greeter
	serviceList := []core.GRPCServer{
		// Example: register the greeter service to the gateway
		services.NewGreeter(),
		// add more service here ...
		services.NewTicket(),
	}

	return e.visit(ctx, serviceList...)
}

// visit all service by Accept function
func (e *Edge) visit(ctx context.Context, services ...core.GRPCServer) error {
	for _, service := range services {
		if err := service.Accept(ctx, e.server); err != nil {
			return err
		}
	}

	return nil
}

// functions is chain of functions to use before starting the edge app
func (e *Edge) functions(ctx context.Context) error {
	// serve swagger ui
	openapi.Serve(e.server, e.config)

	// watch service change on service registry
	watch.Serve(e.server, e.config)

	// waf secure middleware layer
	secure.Serve(e.server, e.config)

	// new middleware handler
	middleware.Serve(e.server, e.config)

	return e.registerServer(ctx)
}

// ListenAndServe the edge/gateway app
func (e *Edge) ListenAndServe(ctx context.Context) error {
	// service ascii art banner
	version.ASCII()
	if err := protobuf.Validate(e.config); err != nil {
		return err
	}

	// add chain functions handler request
	if err := e.functions(ctx); err != nil {
		return err
	}

	// Listen HTTP server (and edge calls to gRPC server endpoint)
	// log info in console and return register error if they exist
	txlog.Infof("[http] starting server %s", e.config.GetApiAddr())
	return e.server.Listen(e.config.GetApiAddr())
	// return errors.F("edge: failed to listen and serve")
}

// Shutdown implements core.Server.
func (e *Edge) Shutdown(ctx context.Context) error {
	return e.server.Shutdown(ctx)
}
