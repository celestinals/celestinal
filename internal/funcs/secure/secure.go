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

// Package secure provide secure middleware layer
package secure

import (
	"net/http"

	"github.com/corazawaf/coraza/v3"
	txhttp "github.com/corazawaf/coraza/v3/http"
	"github.com/corazawaf/coraza/v3/types"
	"google.golang.org/grpc/grpclog"

	"github.com/tickexvn/tickex/api/gen/go/common/env/config/v1"
	"github.com/tickexvn/tickex/pkg/cli"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/logger"
)

// Serve the edge with WAF secure middleware layer
func Serve(edge core.Edge, _ *config.Config) {
	flags := cli.ParseEdge()
	if !flags.GetSecure() {
		if len(flags.GetRules()) > 0 {
			logger.Warn("[HTTP] OWASP rules was't applied (--secure=false)")
		}
		return
	}

	if len(flags.GetRules()) == 0 {
		logger.Warn("[HTTP] OWASP CRS .conf rules was't provided")
		return
	}

	waf, err := NewWAF(flags.GetRules()...)
	if err != nil {
		logger.Errorf("init secure layer err: %v", err)
		return
	}

	edge.Use(waf.Secure)
}

// NewWAF create a new WAF middleware layer
func NewWAF(filepaths ...string) (*WAF, error) {
	var wafconf = coraza.NewWAFConfig().
		WithErrorCallback(logError)

	for _, filepath := range filepaths {
		wafconf = wafconf.WithDirectivesFromFile(filepath)
	}

	cozarawaf, err := coraza.NewWAF(wafconf)
	if err != nil {
		return nil, err
	}

	return &WAF{
		cozarawaf: cozarawaf,
	}, nil
}

// WAF middleware layer
type WAF struct {
	cozarawaf coraza.WAF
}

// Secure HTTP middleware with Coraza WAF
func (waf *WAF) Secure(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txhttp.WrapHandler(waf.cozarawaf, next).ServeHTTP(w, r)
	})
}

func logError(err types.MatchedRule) {
	msg := err.ErrorLog()
	grpclog.Errorf("[%s] %s\n", err.Rule().Severity(), msg)
}
