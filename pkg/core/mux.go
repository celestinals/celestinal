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
)

// Edge is an interface for a runtime mux.
type Edge interface {
	Listen(conf *EdgeConfig) error
	AsRuntimeMux() *runtime.ServeMux
	AsServeMux() *http.ServeMux
	HandleFunc(pattern string, handler http.HandlerFunc)
}

type EdgeConfig struct {
	Addr    string
	Handler http.Handler
}

// NewEdge creates a new runtime mux.
func NewEdge(opts ...runtime.ServeMuxOption) Edge {
	return &edge{
		mux:     runtime.NewServeMux(opts...),
		httpMux: http.NewServeMux(),
	}
}

// edge is a runtime mux.
type edge struct {
	mux     *runtime.ServeMux
	httpMux *http.ServeMux
	server  *http.Server
}

func (e *edge) HandleFunc(pattern string, handler http.HandlerFunc) {
	e.httpMux.HandleFunc(pattern, handler)
}

// Listen starts the runtime mux.
func (e *edge) Listen(conf *EdgeConfig) error {
	if conf == nil {
		conf = &EdgeConfig{
			Addr:    ":9000",
			Handler: e.mux,
		}
	}

	e.httpMux.Handle("/", e.mux)
	e.server = &http.Server{
		Addr:    conf.Addr,
		Handler: conf.Handler,
	}

	return e.server.ListenAndServe()
}

// AsRuntimeMux returns the underlying runtime mux.
func (e *edge) AsRuntimeMux() *runtime.ServeMux {
	return e.mux
}

func (e *edge) AsServeMux() *http.ServeMux {
	return e.httpMux
}
