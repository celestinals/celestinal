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

// Package visitor provides an implementation of the visitor pattern.
package visitor

import (
	"context"

	"github.com/celestinals/celestinal/internal/utils/eventq"

	cestcore "github.com/celestinals/celestinal/pkg/core"
	cestlog "github.com/celestinals/celestinal/pkg/logger"
	cestns "github.com/celestinals/celestinal/pkg/names"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// VisitService visits the greeter service.
func VisitService(ctx context.Context, ns cestns.Namespace, server cestcore.HTTPServer,
	service cestcore.ServiceRegistrar) error {

	eventq.Subscribe(ctx, ns.String(), func(endpoint string) error {
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		cestlog.Infof("[visitor.VisitService] %s %s", ns.String(), "******")
		return cestcore.RegisterService(ctx, server, service, endpoint, opts)
	})

	return nil
}
