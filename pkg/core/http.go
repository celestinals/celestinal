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

package cestcore

import (
	"context"
	"net/http"

	cesterr "github.com/celestinals/celestinal/pkg/errors"
	cestflag "github.com/celestinals/celestinal/pkg/flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var (
	// Ensure httpServer implements Server.
	_ Server = (*httpServer)(nil)

	// Ensure httpServer implements HttpServer.
	_ HTTPServer = (*httpServer)(nil)
)

// HTTPServer is an interface for a http server.
// default port is 9000
type HTTPServer interface {
	Listen(address string) error
	Shutdown(ctx context.Context) error
	RuntimeMux() *runtime.ServeMux
	HTTPMux() *http.ServeMux
	Use(handler func(http.Handler) http.Handler)
}

// NewHTTPServer creates a new http server.
func NewHTTPServer(opts ...runtime.ServeMuxOption) HTTPServer {
	return &httpServer{
		runtimeMux: runtime.NewServeMux(opts...),
		httpMux:    http.NewServeMux(),
	}
}

// httpServer is a http server with http serve mux and grpc-gateway serve mux.
type httpServer struct {
	// grpc-gateway runtime mux
	runtimeMux *runtime.ServeMux

	// http mux
	httpMux *http.ServeMux

	// middlewares for the http server
	middlewares []func(http.Handler) http.Handler

	// http server
	server *http.Server
}

// ListenAndServe implements Server.
func (h *httpServer) Start(ctx context.Context) error {
	_ = ctx
	return cesterr.ErrUnimplemented
}

// handler wraps the http handler with the middlewares. middlewares
// will be executed in the order they are added, top to bottom.
func (h *httpServer) handler(httpHandler http.Handler) http.Handler {
	for _, middleware := range h.middlewares {
		httpHandler = middleware(httpHandler)
	}
	return httpHandler
}

// Use middleware for the http server. Middleware will be called
// in the order they are added, top to bottom. the middleware will
// be executed before the http handler.
func (h *httpServer) Use(handler func(http.Handler) http.Handler) {
	h.middlewares = append(h.middlewares, handler)
}

// Listen starts the runtime mux.
func (h *httpServer) Listen(address string) error {
	if address == "" {
		address = cestflag.Parse().GetAddress()
	}

	// handler runtime.Mux with http.ServeMux
	// serve grpc-gateway mux on the root path
	h.httpMux.Handle("/", h.runtimeMux)

	// create http server with address and http.Handler
	// httpMux was wrapped with the middlewares
	h.server = &http.Server{
		Addr:    address,
		Handler: h.handler(h.httpMux),
	}

	return h.server.ListenAndServe()
}

// RuntimeMux returns the underlying runtime mux.
func (h *httpServer) RuntimeMux() *runtime.ServeMux {
	return h.runtimeMux
}

// HTTPMux returns the underlying http mux
func (h *httpServer) HTTPMux() *http.ServeMux {
	return h.httpMux
}

// Shutdown implements HttpServer.
func (h *httpServer) Shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
