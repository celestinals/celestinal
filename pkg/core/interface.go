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

// Package core provides the core setting for the application.
package core

import (
	"context"
)

// Application represents the application when all constructor was build
// by core.Build() start the app, it will start the server and provide all
// constructor needed
type Application interface {
	Start(ctx context.Context) error
}

// Server represents the HTTP/gRPC server interface.
type Server interface {
	ListenAndServe() error
}
