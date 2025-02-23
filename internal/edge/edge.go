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

// Package edge provides the edge/gateway of tickex.
package edge

import (
	"context"

	typepb "github.com/tickexvn/tickex/api/gen/go/types/v1"
	"github.com/tickexvn/tickex/internal/edge/openapi"
	"github.com/tickexvn/tickex/internal/edge/services/greeter/v1"
	"github.com/tickexvn/tickex/internal/edge/types"
	"github.com/tickexvn/tickex/internal/edge/visitor"
	"github.com/tickexvn/tickex/internal/middleware"
	"github.com/tickexvn/tickex/pkg/constant"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/logger"
	"github.com/tickexvn/tickex/pkg/pbtools"
)

var _ core.Server = (*Edge)(nil)

// New creates a new gateway app
func New(conf *typepb.Config) core.Server {
	return &Edge{
		edge:    core.NewEdge(),
		visitor: visitor.New(conf),
		config:  conf,
	}
}

// Edge represents the gateway app
type Edge struct {
	config  *typepb.Config
	edge    core.Edge
	visitor types.IVisitor
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

// Declare function in edge/types at types.IVisitor interface
//
// Ex:
//
//	type IVisitor interface {
//		VisitGreeterService(ctx context.Context, edge core.Edge, service IService) error
//	}
//
// Implement function at visitor.Visitor:
//
// Ex:
//
//	func (v *Visitor) VisitGreeterService(ctx context.Context, edge core.Edge, service types.IService) error {
//		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
//
//		greeterAddr := ":8000"
//		if err := core.RegisterService(ctx, edge, service, greeterAddr, opts); err != nil {
//			return err
//		}
//
//		return nil
//	}
func (e *Edge) register(ctx context.Context) error {
	// TODO: Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	// Create folder at services, inherit base package, override function, implement business logic
	// See: edge/services/greeter
	services := []types.IService{
		// Example: register the greeter service to the edge/gateway
		&greeter.Greeter{},
		// Add more services here ...
	}

	return e.visit(ctx, services...)
}

// ListenAndServe the edge/gateway app
func (e *Edge) ListenAndServe() error {
	if err := pbtools.Validate(e.config); err != nil {
		return err
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// register service
	if err := e.register(ctx); err != nil {
		return err
	}

	// serve swagger ui
	openapi.Serve(e.edge)

	// log info in console
	logger.Infof(constant.InfoHTTPServer, e.config.GetGatewayAddress())

	// new middleware handler
	mdw := middleware.New(e.config)

	// Listen HTTP server (and edge calls to gRPC server endpoint)
	return e.edge.Listen(&core.EdgeConfig{
		Addr:    e.config.GetGatewayAddress(),
		Handler: mdw.LogRequestBody(mdw.AllowCORS(e.edge.AsMux())),
	})
}
