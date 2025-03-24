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

package core

import (
	"context"

	coreinternal "github.com/tickexvn/tickex/pkg/core/internal"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// Build builds the application.
// The application is built by providing the constructors.
func Build(constructors ...any) Application {
	for _, constructor := range constructors {
		coreinternal.Provide(constructor)
	}

	// disable log: use fx.NopLogger
	return &container{
		engine: fx.New(coreinternal.Option(), fx.Invoke(runner)),
	}
}

// RegisterService registers a service with the runtime.
// The service is registered with the runtime mux.
//
// dependency:
//
// - github.com/grpc-ecosystem/grpc-gateway/v2/runtime
func RegisterService(
	ctx context.Context, edge Edge, service GRPCService, endpoint string,
	opts []grpc.DialOption) error {

	if err := service.Register(
		ctx, edge.AsRuntimeMux(), endpoint, opts); err != nil {
		return err
	}

	return nil
}
