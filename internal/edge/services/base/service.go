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

// Package base provides the base service.
package base

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/tickexvn/tickex/pkg/core"
)

var _ IService = (*Base)(nil)

// IService represents the service gRPC interface.
type IService interface {
	core.GRPCService
	Accept(context.Context, core.Edge) error
}

// Base represents the base service
type Base struct{}

// Register registers the base service
func (b Base) Register(_ context.Context, _ *runtime.ServeMux, _ string,
	_ []grpc.DialOption) error {
	panic("unimplemented")
}

// Accept accepts the base service
func (b Base) Accept(_ context.Context, _ core.Edge) error {
	panic("unimplemented")
}
