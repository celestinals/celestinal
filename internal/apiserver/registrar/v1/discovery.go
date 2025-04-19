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

	discoverypb "github.com/celestinals/celestinal/api/gen/go/celestinal/discovery/v1"
	registrarbase "github.com/celestinals/celestinal/internal/apiserver/registrar/base"

	"github.com/celestinals/celestinal/internal/pkg/visitor"
	"github.com/celestinals/celestinal/pkg/frw/striker/skhttp"
	"github.com/celestinals/celestinal/pkg/frw/striker/skutils"
)

var _ skutils.ServiceRegistrar = (*discovery)(nil)

// NewDiscovery creates a new Discovery service to register handler to gateway
func NewDiscovery(srv discoverypb.DiscoveryServiceServer) skutils.ServiceRegistrar {
	return discovery{server: srv}
}

// discovery represents the discovery service
type discovery struct {
	registrarbase.Base
	server discoverypb.DiscoveryServiceServer
}

// Accept accepts the Discovery service
func (d discovery) Accept(ctx context.Context, server skhttp.Server) error {
	return visitor.VisitService(ctx, server, d)
}

// Register registers the Discovery service
func (d discovery) Register(ctx context.Context, mux *runtime.ServeMux) error {
	return discoverypb.RegisterDiscoveryServiceHandlerServer(ctx, mux, d.server)
}
