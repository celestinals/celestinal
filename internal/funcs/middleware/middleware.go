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

// Package middleware provide http handler - net/http
package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc/grpclog"

	"github.com/tickexvn/tickex/api/gen/go/common/env/config/v1"
	robotpb "github.com/tickexvn/tickex/api/gen/go/common/robot/v1"
	"github.com/tickexvn/tickex/internal/utils/version"
	"github.com/tickexvn/tickex/pkg/cli"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/pbtools"
	"github.com/tickexvn/tickex/pkg/robot"
)

// New middleware handler
func New(conf *config.Config) *Middleware {
	return &Middleware{
		conf: conf,
	}
}

// Serve middleware handler for http handler in edge server
func Serve(edge core.Edge, conf *config.Config) {
	// new middleware handler
	mdw := New(conf)

	// mdw.LogRequestBody(mdw.AllowCORS(e.edge.AsMux()))
	edge.Use(mdw.AllowCORS)
	edge.Use(mdw.LogRequestBody)
}

// Middleware for http handler in grpc gateway
type Middleware struct {
	conf *config.Config
}

// LogRequestBody logs the request body when the response status code is not 200.
// This addresses the issue of being unable to retrieve the request body in the
// customErrorHandler middleware.
func (mdw *Middleware) LogRequestBody(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := newLogResponseWriter(w)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w,
				fmt.Sprintf("grpc server read request body err %+v", err),
				http.StatusBadRequest)

			return
		}

		clonedR := r.Clone(r.Context())
		clonedR.Body = io.NopCloser(bytes.NewReader(body))

		h.ServeHTTP(lw, clonedR)

		if lw.statusCode >= 400 {
			grpclog.Errorf("http error=%+v request body=%+v",
				lw.statusCode, string(body))

			// send log to telegram
			mdw.notify(lw.statusCode, string(body))
		}
	})
}

// AllowCORS allows Cross Origin Resource Sharing from any origin.
// Don't do this without consideration in production systems.
func (mdw *Middleware) AllowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")                           // Fix cache issue
			w.Header().Set("Access-Control-Allow-Credentials", "true") // If cookies need to be sent

			if r.Method == "OPTIONS" &&
				r.Header.Get("Access-Control-Request-Method") != "" {

				mdw.preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func (mdw *Middleware) preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))

	// Fix: Add status code to avoid error
	w.WriteHeader(http.StatusNoContent)

	// Fix: Add Max-Age to optimize
	w.Header().Set("Access-Control-Max-Age", "86400")

	grpclog.Infof("Preflight request for %s", r.URL.Path)
}

func (mdw *Middleware) notify(statusCode int, body string) {
	monitor, _ := robot.New(mdw.conf)

	footer := fmt.Sprintf("%s<%s<<%s<<<%s<<<<<<<<<<<<<<<<<<\nGITHUB<COM<<TICKEXVN<<<TICKEX<<<<<<<<<<<<<<<\n",
		"I", version.FullName, version.Version, version.GoVersion)
	_ = monitor.Send(&robotpb.Message{
		Metadata: &robotpb.MessageMetadata{
			CreatedAt: pbtools.ToTime(time.Now().Local()),
			Author:    cli.Parse().GetName(),
		},
		Header: fmt.Sprintf("http error %+v ", statusCode),
		Body:   fmt.Sprintf("http error %+v request body %+v", statusCode, string(body)),
		Footer: footer,
	})
}
