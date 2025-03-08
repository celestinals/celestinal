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

// Package visitor provides an implementation of the visitor pattern.
package visitor

import (
	"context"

	"github.com/tickexvn/tickex/api/gen/go/universal/env/config/v1"
	"github.com/tickexvn/tickex/internal/edge/types"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/discovery"
	"github.com/tickexvn/tickex/pkg/eventq"
	"github.com/tickexvn/tickex/pkg/logger"
	"github.com/tickexvn/tickex/pkg/namespace"
	"github.com/tickexvn/tickex/pkg/pbtools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// New returns a new visitor.
func New(conf *config.Config) types.IVisitor {
	return &Visitor{
		conf:    conf,
		visitor: discovery.NewVisitor(conf),
	}
}

// implement types.IVisitor interfaces bellows
var _ types.IVisitor = (*Visitor)(nil)

// Visitor represents the visitor interface.
type Visitor struct {
	conf    *config.Config
	visitor discovery.Visitor
}

// VisitGreeterService visits the greeter service.
func (v *Visitor) VisitGreeterService(
	ctx context.Context, edge core.Edge, service types.IService) error {
	if err := pbtools.Validate(v.conf); err != nil {
		return err
	}

	eventq.Subscribe(ctx, namespace.GreeterV1, func(endpoint string) error {
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		logger.Infof("visiting %s %s", namespace.GreeterV1, "******")
		return core.RegisterService(ctx, edge, service, endpoint, opts)
	})

	return nil
}
