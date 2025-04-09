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

	"github.com/celestinals/celestinal/pkg/capsule/capsulegrpc"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/greeter/v1"
	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	"github.com/celestinals/celestinal/internal/greeter/v1/controllers"

	"github.com/celestinals/celestinal/pkg/capsule"
	"github.com/celestinals/celestinal/pkg/flag"
	"github.com/celestinals/celestinal/pkg/names"
	"github.com/celestinals/celestinal/pkg/protobuf"
)

// make sure Greeter implement capsule.Server
// it will start by capsule.runner through capsule.Server
var _ capsule.Server = (*Greeter)(nil)

// New creates a new Greeter module.
func New(srv controllers.IGreeter, conf *celestinal.Config) capsule.Server {
	return &Greeter{
		Server: capsulegrpc.NewDefault(),
		srv:    srv,
		config: conf,
	}
}

// Greeter implements GreeterServiceServer.
type Greeter struct {
	*capsulegrpc.Server // inherit capsulegrpc.Server
	config              *celestinal.Config
	srv                 greeter.GreeterServiceServer
}

// Start implements IGreeter, override capsule.Server.Start
func (g *Greeter) Start(_ context.Context) error {
	greeter.PrintASCII()
	if err := protobuf.Validate(g.config); err != nil {
		return err
	}

	greeter.RegisterGreeterServiceServer(g.AsServer(), g.srv)

	return g.Serve(&capsule.ServiceInfo{
		Config: g.config,
		Addr:   flag.Parse().GetAddress(),
		Tags:   []string{"greeter", names.GreeterV1.String()},
		Name:   names.GreeterV1.String(),
	})
}
