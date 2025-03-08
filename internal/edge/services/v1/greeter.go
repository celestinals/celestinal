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

// Package services provides all service declare.
package services

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	greetergw "github.com/tickexvn/tickex/api/gen/go/greeter/v1"
	"github.com/tickexvn/tickex/internal/edge/services"
	"github.com/tickexvn/tickex/internal/edge/types"
	"github.com/tickexvn/tickex/pkg/core"
	"google.golang.org/grpc"
)

var _ types.IService = (*Greeter)(nil)

// NewGreeter creates a new Greeter service to register handler to gateway
func NewGreeter() Greeter {
	return Greeter{}
}

// Greeter represents the Greeter service
type Greeter struct {
	services.Service
}

// Accept accepts the Greeter service
func (g Greeter) Accept(
	ctx context.Context, edge core.Edge, v types.IVisitor) error {

	return v.VisitGreeterService(ctx, edge, g)
}

// Register registers the Greeter service
func (g Greeter) Register(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {

	return greetergw.RegisterGreeterServiceHandlerFromEndpoint(
		ctx, mux, endpoint, opts)
}
