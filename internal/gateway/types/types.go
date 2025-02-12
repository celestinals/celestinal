/*
 * Copyright 2024 The Tickex Authors.
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

// Package types provides the types for the gateway.
package types

import (
	"context"

	"github.com/tickexvn/tickex/pkg/core/tkxruntime"
	"github.com/tickexvn/tickex/pkg/core/tkxservice"
)

// IService represents the service interface.
type IService interface {
	tkxservice.GRPCServicer
	Accept(context.Context, tkxruntime.IServeMux, IVisitor) error
}

// IVisitor represents the visitor interface.
type IVisitor interface {
	VisitGreeterService(ctx context.Context, mux tkxruntime.IServeMux, service IService) error
}
