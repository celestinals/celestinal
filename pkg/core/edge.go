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
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tickexvn/tickex/pkg/cli"
)

// Edge is an interface for a runtime mux with http server.
// default port is 9000
type Edge interface {
	Listen(address string) error
	AsRuntimeMux() *runtime.ServeMux
	AsMux() *http.ServeMux
	Use(handler func(http.Handler) http.Handler)
}

// NewEdge creates a new http server.
func NewEdge(opts ...runtime.ServeMuxOption) Edge {
	return &edge{
		mux:     runtime.NewServeMux(opts...),
		httpMux: http.NewServeMux(),
	}
}

// edge is a http server with http serve mux and grpc-gateway serve mux.
type edge struct {
	// grpc-gateway runtime mux
	mux *runtime.ServeMux

	// http server mux
	httpMux *http.ServeMux

	// middlewares for the http server
	middlewares []func(http.Handler) http.Handler
}

// handler wraps the http handler with the middlewares. middlewares
// will be executed in the order they are added, top to bottom.
func (e *edge) handler(httpHandler http.Handler) http.Handler {
	for _, middleware := range e.middlewares {
		httpHandler = middleware(httpHandler)
	}
	return httpHandler
}

// Use middleware for the http server. Middleware will be called
// in the order they are added, top to bottom. the middleware will
// be executed before the http handler.
func (e *edge) Use(handler func(http.Handler) http.Handler) {
	e.middlewares = append(e.middlewares, handler)
}

// Listen starts the runtime mux.
func (e *edge) Listen(address string) error {
	if address == "" {
		address = cli.Parse().GetAddress()
	}

	// handler runtime.Mux with http.ServeMux
	// serve grpc-gateway mux on the root path
	e.httpMux.Handle("/", e.mux)

	// create http server with address and http.Handler
	// httpMux was wrapped with the middlewares
	server := &http.Server{
		Addr:    address,
		Handler: e.handler(e.httpMux),
	}

	return server.ListenAndServe()
}

// AsRuntimeMux returns the underlying runtime mux.
func (e *edge) AsRuntimeMux() *runtime.ServeMux {
	return e.mux
}

// AsMux returns the underlying http mux
func (e *edge) AsMux() *http.ServeMux {
	return e.httpMux
}
