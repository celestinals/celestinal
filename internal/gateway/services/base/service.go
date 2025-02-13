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

	"github.com/tickexvn/tickex/pkg/core"

	"github.com/tickexvn/tickex/internal/gateway/types"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var _ types.IService = (*Service)(nil)

// Service represents the base service
type Service struct{}

// Register registers the base service
func (s *Service) Register(_ context.Context, _ *runtime.ServeMux, _ string, _ []grpc.DialOption) error {
	panic("unimplemented")
}

// Accept accepts the base service
func (s *Service) Accept(_ context.Context, _ core.IServeMux, _ types.IVisitor) error {
	panic("unimplemented")
}
