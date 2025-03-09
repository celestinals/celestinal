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

	"github.com/tickexvn/tickex/internal/edge/types"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/eventq"
	"github.com/tickexvn/tickex/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// New returns a new visitor.
func New() types.IVisitor {
	return &Visitor{}
}

// implement types.IVisitor interfaces bellows
var _ types.IVisitor = (*Visitor)(nil)

// Visitor represents the visitor interface.
type Visitor struct{}

// VisitService visits the greeter service.
func (v Visitor) VisitService(ctx context.Context,
	namespace string, edge core.Edge, service types.IService) error {

	eventq.Subscribe(ctx, namespace, func(endpoint string) error {
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		logger.Infof("visiting %s %s", namespace, "******")
		return core.RegisterService(ctx, edge, service, endpoint, opts)
	})

	return nil
}
