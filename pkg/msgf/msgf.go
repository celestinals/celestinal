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

// Package msgf contain log message title with format
package msgf

import (
	"github.com/tickexvn/tickex/api/gen/go/types/v1"
	"github.com/tickexvn/tickex/internal/version"
)

var (
	// InfoGrpcServer gRPC server listening on [PORT]
	InfoGrpcServer = version.Header(types.Status_I) + " [gRPC] Listening on %s"

	// InfoHTTPServer HTTP server listening on [PORT]
	InfoHTTPServer = version.Header(types.Status_I) + " [HTTP] Listening on %s"
)
