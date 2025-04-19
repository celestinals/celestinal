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

// Package openapi serve apiserver to host swagger ui
package openapi

import (
	"net/http"
	"strings"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	"github.com/celestinals/celestinal/pkg/flags"
	"github.com/celestinals/celestinal/pkg/frw/striker/skhttp"
)

// New creates a new OpenAPI handler.
func New() *OpenAPI {
	return &OpenAPI{}
}

// OpenAPI serves the OpenAPI specification and Swagger UI.
type OpenAPI struct{}

// RegisterServer return api json and swagger ui
func (oapi *OpenAPI) RegisterServer(server skhttp.Server, _ *celestinal.Config) {
	flags := flags.ParseAPIServer()

	apifs := http.FileServer(http.Dir(flags.GetApiSpecsPath()))
	server.HTTPMux().Handle("/api/", http.StripPrefix("/api/", apifs))

	swaggerfs := http.FileServer(http.Dir(flags.GetSwaggerPath()))
	server.HTTPMux().Handle("/swagger/", oapi.apiSpecSwaggerHandler(swaggerfs))
	server.HTTPMux().HandleFunc("/swagger", oapi.swaggerHandler())
}

func (oapi *OpenAPI) swaggerHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/swagger/", http.StatusMovedPermanently)
	}
}

func (oapi *OpenAPI) apiSpecSwaggerHandler(swaggerfs http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".json") {
			target := "/api" + strings.TrimPrefix(r.URL.Path, "/swagger")
			http.Redirect(w, r, target, http.StatusTemporaryRedirect)
			return
		}

		http.StripPrefix("/swagger/", swaggerfs).ServeHTTP(w, r)
	})
}
