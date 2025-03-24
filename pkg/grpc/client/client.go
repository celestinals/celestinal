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

// Package grpcclient provides the grpc client for the greeter service.
package grpcclient

import (
	"context"
	"fmt"

	"github.com/tickexvn/tickex/api/gen/go/common/env/config/v1"
	"github.com/tickexvn/tickex/api/gen/go/greeter/v1"
	"github.com/tickexvn/tickex/pkg/discovery"
	"github.com/tickexvn/tickex/pkg/namespace"
	"github.com/tickexvn/tickex/pkg/txlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewGreeterClient creates a new greeter grpc client.
func NewGreeterClient(
	ctx context.Context, conf *config.Config) (greeter.GreeterServiceClient, error) {

	endpoint, err := discovery.Visit(ctx, conf, namespace.GreeterV1)
	if err != nil {
		txlog.Errorf("failed to connect greeter service: %v", err)
		return nil, fmt.Errorf("discovery.Visit : %v", err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		txlog.Errorf("failed to create greeter client connection: %v", err)
		return nil, fmt.Errorf("grpc.NewClient : %v", err)
	}

	return greeter.NewGreeterServiceClient(conn), nil
}
