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

// Package greeter implements the Greeter service server.
package greeter

import (
	"context"
	"time"

	"github.com/celestinals/celestinal/pkg/frw/striker/skgrpc"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/greeter/v1"
	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"

	greeterctrls "github.com/celestinals/celestinal/internal/greeter/v1/controllers"
	greeterdomain "github.com/celestinals/celestinal/internal/greeter/v1/domain"

	"github.com/celestinals/celestinal/pkg/flags"
	"github.com/celestinals/celestinal/pkg/frw/striker"
	"github.com/celestinals/celestinal/pkg/names"
	"github.com/celestinals/celestinal/pkg/protobuf"
)

// make sure Greeter implement striker.Server
// it will start by striker.runner through striker.Server
var _ striker.Server = (*Greeter)(nil)

// inject all dependencies to the greeter
// This is a dependency injection pattern.
var (
	_ = striker.Inject(greeterctrls.New)
	_ = striker.Inject(greeterdomain.New)
)

// New creates a new Greeter module.
func New(srv greeterctrls.IGreeter, conf *celestinal.Config) striker.Server {
	return &Greeter{
		Server: skgrpc.NewDefault(),
		srv:    srv,
		config: conf,
	}
}

// Greeter implements GreeterServiceServer.
type Greeter struct {
	*skgrpc.Server // inherit skgrpc.Server
	config         *celestinal.Config
	srv            greeter.GreeterServiceServer
}

// Start implements IGreeter, override striker.Server.Start
func (g *Greeter) Start(_ context.Context) error {
	greeter.PrintASCII()
	if err := protobuf.Validate(g.config); err != nil {
		return err
	}

	greeter.RegisterGreeterServiceServer(g.AsServer(), g.srv)

	return g.Serve(&striker.ServiceInfo{
		Config:      g.config,
		Addr:        flags.ParseGRPCService().GetAddress(),
		GatewayAddr: flags.ParseGRPCService().GetGatewayAddress(),
		Name:        names.GreeterV1.String(),
		TTL:         time.Minute * 5,
	})
}
