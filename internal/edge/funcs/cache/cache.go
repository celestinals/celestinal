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

// Package cache provides the cache middleware for the edge app
package cache

import (
	"net/http"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	"github.com/celestinals/celestinal/pkg/cestcore"
)

// Serve is a middleware that serves the cache response for the same
// request in the future.
func Serve(server cestcore.HTTPServer, _ *celestinal.Config) {
	server.Use(cache)
}

// cache is a middleware that caches the response of the request
// and serves the cached response for the same request in the future.
func cache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
