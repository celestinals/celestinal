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
	"context"

	"github.com/tickexvn/tickex/api/gen/go/tickex/greeter/v1"
	"github.com/tickexvn/tickex/api/gen/go/tickex/v1"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/flag"
	"github.com/tickexvn/tickex/pkg/namespace"
	"github.com/tickexvn/tickex/pkg/protobuf"
	"github.com/tickexvn/tickex/x/greeter/internal/controllers"
)

// make sure Greeter implement core.Server
// it will start by core.runner through core.Server
var _ core.Server = (*Greeter)(nil)

// New creates a new Greeter module.
func New(srv controllers.IGreeter, conf *tickex.Config) core.Server {
	return &Greeter{
		ServiceServer: core.NewDefault(),
		srv:           srv,
		config:        conf,
	}
}

// Greeter implements GreeterServiceServer.
type Greeter struct {
	*core.ServiceServer // inherit core.ServiceServer
	config              *tickex.Config
	srv                 greeter.GreeterServiceServer
}

// ListenAndServe implements IGreeter, override core.ServiceServer.ListenAndServe
func (g *Greeter) ListenAndServe(_ context.Context) error {
	greeter.PrintASCII()
	if err := protobuf.Validate(g.config); err != nil {
		return err
	}

	greeter.RegisterGreeterServiceServer(g.AsServer(), g.srv)

	return g.Serve(&core.ServiceInfo{
		Config: g.config,
		Addr:   flag.Parse().GetAddress(),
		Tags:   []string{"greeter", namespace.GreeterV1},
		Name:   namespace.GreeterV1,
	})
}
