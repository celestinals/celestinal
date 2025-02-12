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
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// IServeMux is an interface for a runtime mux.
type IServeMux interface {
	Listen(addr string) error
	AsRuntimeMux() *runtime.ServeMux
}

// NewServeMux creates a new runtime mux.
func NewServeMux(opts ...runtime.ServeMuxOption) IServeMux {
	return &ServeMux{
		mux: runtime.NewServeMux(opts...),
	}
}

// ServeMux is a runtime mux.
type ServeMux struct {
	mux *runtime.ServeMux
}

// Listen starts the runtime mux.
func (mux *ServeMux) Listen(addr string) error {
	return http.ListenAndServe(addr, mux.AsRuntimeMux())
}

// AsRuntimeMux returns the underlying runtime mux.
func (mux *ServeMux) AsRuntimeMux() *runtime.ServeMux {
	return mux.mux
}
