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

// Package server implements the Greeter service server.
package server

import (
	"github.com/tickexvn/tickex/api/gen/go/common/env/config/v1"
	"github.com/tickexvn/tickex/api/gen/go/greeter/v1"
	"github.com/tickexvn/tickex/pkg/cli"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/namespace"
	"github.com/tickexvn/tickex/pkg/pbtools"
	"github.com/tickexvn/tickex/x/greeter/internal/controllers"
)

var _ core.Server = (*Greeter)(nil)

// New creates a new Greeter module.
func New(srv controllers.IGreeter, conf *config.Config) core.Server {
	return &Greeter{
		ServiceServer: core.NewDefault(),
		srv:           srv,
		config:        conf,
	}
}

// Greeter implements GreeterServiceServer.
type Greeter struct {
	*core.ServiceServer
	config *config.Config
	srv    greeter.GreeterServiceServer
}

// ListenAndServe implements IGreeter.
func (g *Greeter) ListenAndServe() error {
	if err := pbtools.Validate(g.config); err != nil {
		return err
	}

	greeter.RegisterGreeterServiceServer(g.AsServer(), g.srv)

	return g.Serve(&core.ServiceInfo{
		Config: g.config,
		Addr:   cli.Parse().GetAddress(),
		Tags:   []string{"greeter", namespace.GreeterV1},
		Name:   namespace.GreeterV1,
	})
}
