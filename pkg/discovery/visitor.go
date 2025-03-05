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

package discovery

import (
	"context"
	"fmt"

	discoverypb "github.com/tickexvn/tickex/api/gen/go/discovery/v1"
	"github.com/tickexvn/tickex/api/gen/go/types/v1"
	"github.com/tickexvn/tickex/pkg/errors"
	"github.com/tickexvn/tickex/pkg/logger"
	"google.golang.org/grpc"
)

var discover *Discovery

// Visitor is the interface that wraps the Visit method.
type Visitor interface {
	Visit(ctx context.Context, service grpc.ServiceDesc) (string, error)
}

// NewVisitor returns a new Visitor. It uses the discovery service to
// discover the service.
func NewVisitor(conf *types.Config) Visitor {
	if discover == nil {
		discovery, err := New(conf)
		if err != nil {
			logger.Fatalf("discovery connect failed: %v", err)
		}

		discover = discovery
	}

	return visitor{}
}

type visitor struct{}

func (v visitor) Visit(ctx context.Context, service grpc.ServiceDesc) (string, error) {
	if discover == nil {
		return "", errors.ErrNotFound
	}

	services, err := discover.Discover(ctx, &discoverypb.DiscoverRequest{
		Name: service.ServiceName,
	})
	if err != nil {
		return "", err
	}

	host := services.GetServices()[0].GetHost()
	port := services.GetServices()[0].GetPort()

	return fmt.Sprintf("%s:%d", host, port), nil
}
