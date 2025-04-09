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

// Package services provides all service declare.
package services

import (
	"context"

	"github.com/celestinals/celestinal/pkg/capsule/capsulehttp"
	"github.com/celestinals/celestinal/pkg/capsule/capsuleutils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	greeterpb "github.com/celestinals/celestinal/api/gen/go/celestinal/greeter/v1"
	"github.com/celestinals/celestinal/internal/apigateway/services/base"
	"github.com/celestinals/celestinal/internal/pkg/visitor"

	"github.com/celestinals/celestinal/pkg/names"
)

var _ capsuleutils.ServiceRegistrar = (*greeter)(nil)

// NewGreeter creates a new greeter service to register handler to gateway
func NewGreeter() capsuleutils.ServiceRegistrar {
	return greeter{}
}

// greeter represents the greeter service
type greeter struct {
	base.Base
}

// Accept accepts the greeter service
func (g greeter) Accept(ctx context.Context, server capsulehttp.Server) error {
	return visitor.VisitService(ctx, names.GreeterV1, server, g)
}

// Register registers the greeter service
func (g greeter) Register(ctx context.Context, mux *runtime.ServeMux,
	endpoint string, opts []grpc.DialOption) error {

	return greeterpb.
		RegisterGreeterServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
