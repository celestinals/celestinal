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

// Package registrar provides all service declare.
package registrar

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	greeterpb "github.com/celestinals/celestinal/api/gen/go/celestinal/greeter/v1"
	registrarbase "github.com/celestinals/celestinal/internal/apiserver/registrar/base"

	"github.com/celestinals/celestinal/internal/pkg/visitor"
	"github.com/celestinals/celestinal/pkg/names"
	"github.com/celestinals/celestinal/pkg/striker/skhttp"
	"github.com/celestinals/celestinal/pkg/striker/skutils"
)

var _ skutils.ServiceRegistrar = (*greeter)(nil)

// NewGreeter creates a new greeter service to register handler to gateway
func NewGreeter() skutils.ServiceRegistrar {
	return greeter{}
}

// greeter represents the greeter service
type greeter struct {
	registrarbase.Base
}

// Accept accepts the greeter service
func (g greeter) Accept(ctx context.Context, server skhttp.Server) error {
	return visitor.VisitServiceFromEndpoint(ctx, names.GreeterV1, server, g)
}

// RegisterFromEndpoint registers the greeter service
func (g greeter) RegisterFromEndpoint(ctx context.Context, mux *runtime.ServeMux,
	endpoint string, opts []grpc.DialOption) error {

	return greeterpb.
		RegisterGreeterServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
