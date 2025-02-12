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

// Package tkxruntime provides the runtime layer for the service.
package tkxruntime

import (
	"context"

	"github.com/tickexvn/tickex/pkg/core/tkxservice"

	"google.golang.org/grpc"
)

// RegisterService registers a service with the runtime.
func RegisterService(ctx context.Context, mux IServeMux, service tkxservice.GRPCServicer, endpoint string, opts []grpc.DialOption) error {
	if err := service.Register(ctx, mux.AsRuntimeMux(), endpoint, opts); err != nil {
		return err
	}

	return nil
}
