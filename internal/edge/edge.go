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

	configpb "github.com/tickexvn/tickex/api/gen/go/common/env/config/v1"
	"github.com/tickexvn/tickex/internal/edge/services/v1"
	"github.com/tickexvn/tickex/internal/edge/types"
	"github.com/tickexvn/tickex/internal/edge/visitor"
	"github.com/tickexvn/tickex/internal/funcs/middleware"
	"github.com/tickexvn/tickex/internal/funcs/openapi"
	"github.com/tickexvn/tickex/internal/funcs/watch"
	"github.com/tickexvn/tickex/pkg/constant"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/logger"
	"github.com/tickexvn/tickex/pkg/pbtools"
)

var _ core.Server = (*Edge)(nil)

// New creates a new gateway app
func New(conf *configpb.Config) core.Server {
	return &Edge{
		edge:    core.NewEdge(),
		visitor: visitor.New(),
		config:  conf,
	}
}

// Edge represents the tickex app The edge application is the main
// entry point for the Tickex. It will automatically connect to other
// services via gRPC. Start the application along with other services
// in the x/ directory. The application provides APIs for users through
// a single HTTP gateway following the RESTful API standard. The application
// uses gRPC to connect to other services. Additionally, the system provides
// a Swagger UI interface for users to easily interact with the system
// through a web interface.
type Edge struct {
	// config is the configuration of the edge app, load environment
	// variables from .env file
	config *configpb.Config

	// edge is the core edge server, manage http.ServeMux,
	// runtime.ServeMux and HTTP server
	edge core.Edge

	// visitor is the visitor to visit all services by visitor pattern
	// and register them to the edge server
	visitor types.IVisitor
}

// Register gRPC server endpoint.
// Declare function in edge/types at types.IVisitor interface
//
// Ex:
//
//	type IVisitor interface {
//		VisitGreeterService(
//		ctx context.Context, edge core.Edge, service IService) error
//	}
//
// Implement function at visitor.Visitor:
//
// Ex:
//
//	func (v *Visitor) VisitGreeterService(
//		ctx context.Context, edge core.Edge, service types.IService) error {
//
//		opts := []grpc.DialOption{
//			grpc.WithTransportCredentials(insecure.NewCredentials())}
//
//		greeterAddr := ":8000"
//		if err := core.RegisterService(
//			ctx, edge, service, greeterAddr, opts); err != nil {
//			return err
//		}
//
//		return nil
//	}
func (e *Edge) register(ctx context.Context) error {
	// Note: Make sure the gRPC server is running properly and accessible
	// Create folder at services, inherit base package, override function,
	// implement business logic
	// See: services/v1/greeter
	serviceList := []types.IService{
		// Example: register the greeter service to the gateway
		services.NewGreeter(),
		// TODO: add more service here ...
	}

	return e.visit(ctx, serviceList...)
}

// visit all service by Accept function
func (e *Edge) visit(ctx context.Context, services ...types.IService) error {
	for _, service := range services {
		if err := service.Accept(ctx, e.edge, e.visitor); err != nil {
			return err
		}
	}

	return nil
}

// functions is chain of functions to use before starting the edge app
func (e *Edge) functions(ctx context.Context) error {
	// serve swagger ui
	openapi.Serve(e.edge)

	// watch service change on service registry
	watch.Service(e.config)

	// new middleware handler
	// mdw.LogRequestBody(mdw.AllowCORS(e.edge.AsMux()))
	mdw := middleware.New(e.config)
	e.edge.Use(mdw.AllowCORS)
	e.edge.Use(mdw.LogRequestBody)

	// log info in console and return register error if they exist
	logger.Infof(constant.InfoHTTPServer, e.config.GetApiAddr())
	return e.register(ctx)
}

// ListenAndServe the edge/gateway app
func (e *Edge) ListenAndServe() error {
	if err := pbtools.Validate(e.config); err != nil {
		return err
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := e.functions(ctx); err != nil {
		return err
	}

	// Listen HTTP server (and edge calls to gRPC server endpoint)
	return e.edge.Listen(e.config.GetApiAddr())
}
