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
	"fmt"
	"github.com/tickexvn/tickex/pkg/core/syslog"
	"github.com/tickexvn/tickex/pkg/errors"

	"github.com/tickexvn/tickex/pkg/pbtools"
	"github.com/tickexvn/tickex/pkg/utils"

	"github.com/tickexvn/tickex/api/gen/go/controllers/greeter/v1"
	"github.com/tickexvn/tickex/api/gen/go/types/v1"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/core/net"
	"github.com/tickexvn/tickex/pkg/logger"
	"github.com/tickexvn/tickex/pkg/msgf"
	"github.com/tickexvn/tickex/x/greeter/v1/internal/controllers"
)

var _ core.Server = (*Greeter)(nil)

// Greeter implements GreeterServiceServer.
type Greeter struct {
	*core.ServiceServer
	config *types.Config
	srv    greeter.GreeterServiceServer
}

// ListenAndServe implements IGreeter.
func (g *Greeter) ListenAndServe() error {
	if err := pbtools.Validate(g.config); err != nil {
		errs := errors.New(types.Errors_ERRORS_INVALID_DATA, "validation failed", err)
		syslog.Error(errs.Error())

		return errs.Unwrap()
	}

	listener, err := net.ListenTCP(fmt.Sprintf(":%d", 8000))
	if err != nil {
		return err
	}

	log := logger.New()
	log.Sugar().Infof(msgf.InfoGrpcServer, listener.Addr().String())
	defer utils.CallBack(log.Sync)

	// Listen gRPC srv here
	greeter.RegisterGreeterServiceServer(g.AsServer(), g.srv)

	return g.AsServer().Serve(listener)
}

// New creates a new Greeter module.
func New(srv controllers.IGreeter, conf *types.Config) core.Server {
	return &Greeter{
		ServiceServer: core.NewDefault(),
		srv:           srv,
		config:        conf,
	}
}
