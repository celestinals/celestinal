/*
 * Copyright 2024 The Tickex Authors.
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

// Package gateway provides the gateway of tickex.
package gateway

import (
	"context"

	"github.com/tickexvn/tickex/pkg/core/tkxapp"
	"github.com/tickexvn/tickex/pkg/core/tkxruntime"

	"github.com/tickexvn/tickex/internal/gateway/services/greeter"
	"github.com/tickexvn/tickex/internal/gateway/types"
	"github.com/tickexvn/tickex/internal/gateway/visitor"

	"github.com/tickexvn/tickex/pkg/logger"
	"github.com/tickexvn/tickex/pkg/msgf"
)

var _ tkxapp.Server = (*Engine)(nil)

// Engine represents the gateway app
type Engine struct {
	mux     tkxruntime.IServeMux
	visitor types.IVisitor
}

func (e *Engine) visit(ctx context.Context, services ...types.IService) error {
	for _, service := range services {
		if err := service.Accept(ctx, e.mux, e.visitor); err != nil {
			return err
		}
	}

	return nil
}

func (e *Engine) register(ctx context.Context) error {
	// TODO: Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	services := []types.IService{
		// Example: register the greeter service to the gateway
		&greeter.Greeter{},
		// Add more services here ...
	}

	return e.visit(ctx, services...)
}

// ListenAndServe the gateway app
func (e *Engine) ListenAndServe() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := e.register(ctx); err != nil {
		return err
	}

	// Listen HTTP server (and mux calls to gRPC server endpoint)
	logger.Infof(msgf.InfoHTTPServer, ":9000")
	return e.mux.Listen(":9000")
}

// New creates a new gateway app
func New() tkxapp.Server {
	return &Engine{
		mux:     tkxruntime.NewServeMux(),
		visitor: visitor.New(),
	}
}
