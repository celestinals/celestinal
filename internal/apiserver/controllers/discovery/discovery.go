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

// Package discovery provides a service discovery for the celestinal.
package discovery

import (
	"context"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	discoverysvc "github.com/celestinals/celestinal/internal/apiserver/services/discovery"

	"github.com/celestinals/celestinal/pkg/errors"
	"github.com/celestinals/celestinal/pkg/logger"
	"github.com/celestinals/celestinal/pkg/striker/skhttp"
	"github.com/celestinals/celestinal/pkg/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ celestinal.DiscoveryServiceServer = (*Discovery)(nil)

// New create new instance
func New(svc discoverysvc.Discovery) *Discovery {
	return &Discovery{
		service: svc,
	}
}

// Discovery is a service registry for the celestinal.
type Discovery struct {
	celestinal.UnimplementedDiscoveryServiceServer
	service discoverysvc.Discovery
}

// RegisterServer registers the service registry to the server.
func (dcv *Discovery) RegisterServer(server skhttp.Server, _ *celestinal.Config) {
	if err := celestinal.RegisterDiscoveryServiceHandlerServer(
		context.Background(), server.RuntimeMux(), dcv); err != nil {
		logger.Errorf("discovery.Serve service handler server error: %s", err)
	}
}

// Register registers the service to the service registry.
func (dcv *Discovery) Register(ctx context.Context, request *celestinal.RegisterRequest) (*celestinal.RegisterResponse, error) {
	id := uuid.Generate()

	if err := dcv.service.RegisterService(ctx, id, request); err != nil {
		logger.Errorf("discovery.RegisterFromEndpoint service error: %s", err)
		return nil, errors.StatusInternalError
	}

	return &celestinal.RegisterResponse{
		Id:      id,
		Name:    request.GetName(),
		Address: request.GetAddress(),
	}, nil
}

// Heartbeat is used to send heartbeat to the service registry.
func (dcv *Discovery) Heartbeat(ctx context.Context, request *celestinal.HeartbeatRequest) (*emptypb.Empty, error) {
	_ = ctx
	_ = request
	//TODO implement me
	panic("implement me")
}

// Discover used to discover the service registry.
func (dcv *Discovery) Discover(ctx context.Context, request *celestinal.DiscoverRequest) (*celestinal.DiscoverResponse, error) {
	_ = ctx
	_ = request
	return &celestinal.DiscoverResponse{
		Name: request.GetName(),
	}, nil
}
