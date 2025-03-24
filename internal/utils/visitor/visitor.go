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

	"github.com/tickexvn/tickex/internal/edge/services/base"
	"github.com/tickexvn/tickex/internal/utils/eventq"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/txlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// VisitService visits the greeter service.
func VisitService(ctx context.Context, namespace string, edge core.Edge,
	service base.IService) error {

	eventq.Subscribe(ctx, namespace, func(endpoint string) error {
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		txlog.Infof("[visitor.VisitService] %s %s", namespace, "******")
		return core.RegisterService(ctx, edge, service, endpoint, opts)
	})

	return nil
}
