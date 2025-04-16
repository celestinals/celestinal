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

// Package dcvrctrls provides a service discovery for the celestinal.
package dcvrctrls

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	discoverypb "github.com/celestinals/celestinal/api/gen/go/celestinal/discovery/v1"
	dcvrdomain "github.com/celestinals/celestinal/internal/discovery/v1/domain"
	dcvrrepo "github.com/celestinals/celestinal/internal/discovery/v1/repos"

	"github.com/celestinals/celestinal/pkg/errors"
	"github.com/celestinals/celestinal/pkg/logger"
	"github.com/celestinals/celestinal/pkg/uuid"
)

var _ discoverypb.DiscoveryServiceServer = (*Discovery)(nil)

// New create new instance
func New(repo dcvrrepo.Discovery) *Discovery {
	return &Discovery{
		service: dcvrdomain.NewDiscovery(repo),
	}
}

// Discovery is a service registry for the celestinal.
type Discovery struct {
	discoverypb.UnimplementedDiscoveryServiceServer
	service dcvrdomain.Discovery
}

// Register registers the service to the service registry.
func (dcv *Discovery) Register(ctx context.Context, request *discoverypb.RegisterRequest) (*discoverypb.RegisterResponse, error) {
	id := uuid.Generate()

	if err := dcv.service.RegisterService(ctx, id, request); err != nil {
		logger.Errorf("discovery.Register service error: %s", err)
		return nil, errors.StatusInternalError
	}

	return &discoverypb.RegisterResponse{
		Id:      id,
		Name:    request.GetName(),
		Address: request.GetAddress(),
	}, nil
}

// Heartbeat is used to send heartbeat to the service registry.
func (dcv *Discovery) Heartbeat(ctx context.Context, request *discoverypb.HeartbeatRequest) (*emptypb.Empty, error) {
	_ = ctx
	_ = request
	//TODO implement me
	panic("implement me")
}

// Discover used to discover the service registry.
func (dcv *Discovery) Discover(ctx context.Context, request *discoverypb.DiscoverRequest) (*discoverypb.DiscoverResponse, error) {
	_ = ctx
	_ = request
	return &discoverypb.DiscoverResponse{
		Name: request.GetName(),
	}, nil
}
