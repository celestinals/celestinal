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

// Package openapi serve cestcore.Edge to host swagger ui
package openapi

import (
	"net/http"
	"strings"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	cestcore "github.com/celestinals/celestinal/pkg/core"
	cestflag "github.com/celestinals/celestinal/pkg/flag"
)

// Serve return api json and swagger ui
func Serve(server cestcore.HTTPServer, _ *celestinal.Config) {
	flags := cestflag.ParseEdge()

	apifs := http.FileServer(http.Dir(flags.GetApiSpecsPath()))
	server.HTTPMux().Handle("/api/", http.StripPrefix("/api/", apifs))

	swaggerfs := http.FileServer(http.Dir(flags.GetSwaggerPath()))
	server.HTTPMux().Handle("/swagger/", apiSpecSwaggerHandler(swaggerfs))
	server.HTTPMux().HandleFunc("/swagger", swaggerHandler())
}

func swaggerHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/swagger/", http.StatusMovedPermanently)
	}
}

func apiSpecSwaggerHandler(swaggerfs http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".json") {
			target := "/api" + strings.TrimPrefix(r.URL.Path, "/swagger")
			http.Redirect(w, r, target, http.StatusTemporaryRedirect)
			return
		}

		http.StripPrefix("/swagger/", swaggerfs).ServeHTTP(w, r)
	})
}
