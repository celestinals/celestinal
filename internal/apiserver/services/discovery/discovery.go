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

// Package discverysvc implements the discovery business logic
package discverysvc

import (
	"context"

	discoveryrepo "github.com/celestinals/celestinal/internal/apiserver/repos/discovery"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	"github.com/celestinals/celestinal/internal/pkg/eventq"
)

var _ Discovery = (*discovery)(nil)

// Discovery implement discovery registry business logic
type Discovery interface {
	RegisterService(ctx context.Context, id string, req *celestinal.RegisterRequest) error
}

// NewDiscoveryService creates a new discovery service
// and returns a Discovery interface.
func NewDiscoveryService(repo discoveryrepo.Discovery) Discovery {
	return &discovery{
		repo: repo,
	}
}

type discovery struct {
	repo discoveryrepo.Discovery
}

// RegisterService implements Discovery
func (d *discovery) RegisterService(ctx context.Context, id string, req *celestinal.RegisterRequest) error {

	_ = ctx
	_ = req
	_ = id

	eventq.Publish(req.GetName(), req.GetAddress())

	return nil
}
