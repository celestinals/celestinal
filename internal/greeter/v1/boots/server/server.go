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

// Package server implements the Greeter service server.
package server

import (
	"context"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/greeter/v1"
	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	"github.com/celestinals/celestinal/internal/greeter/v1/controllers"

	cestcore "github.com/celestinals/celestinal/pkg/core"
	cestflag "github.com/celestinals/celestinal/pkg/flag"
	cestns "github.com/celestinals/celestinal/pkg/names"
	cestpb "github.com/celestinals/celestinal/pkg/protobuf"
)

// make sure Greeter implement cestcore.Server
// it will start by cestcore.runner through cestcore.Server
var _ cestcore.Server = (*Greeter)(nil)

// New creates a new Greeter module.
func New(srv controllers.IGreeter, conf *celestinal.Config) cestcore.Server {
	return &Greeter{
		GRPCServer: cestcore.NewGRPCServerDefault(),
		srv:        srv,
		config:     conf,
	}
}

// Greeter implements GreeterServiceServer.
type Greeter struct {
	*cestcore.GRPCServer // inherit cestcore.GRPCServer
	config               *celestinal.Config
	srv                  greeter.GreeterServiceServer
}

// Start implements IGreeter, override cestcore.GRPCServer.Start
func (g *Greeter) Start(_ context.Context) error {
	greeter.PrintASCII()
	if err := cestpb.Validate(g.config); err != nil {
		return err
	}

	greeter.RegisterGreeterServiceServer(g.AsServer(), g.srv)

	return g.Serve(&cestcore.ServiceInfo{
		Config: g.config,
		Addr:   cestflag.Parse().GetAddress(),
		Tags:   []string{"greeter", cestns.GreeterV1.String()},
		Name:   cestns.GreeterV1.String(),
	})
}
