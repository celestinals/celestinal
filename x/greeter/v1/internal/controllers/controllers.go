/*
 * Copyright 2024 The Tickex Authors.
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

// Package controllers provides the controller for the greeter service.
package controllers

import (
	"context"

	"github.com/tickexvn/tickex/api/gen/go/controllers/greeter/v1"
	greeterdomain "github.com/tickexvn/tickex/api/gen/go/domain/greeter/v1"
	"github.com/tickexvn/tickex/pkg/copier"
	"github.com/tickexvn/tickex/x/greeter/v1/internal/domain"
)

// IGreeter defines the interface for the Greeter controller.
type IGreeter interface {
	greeter.GreeterServiceServer
}

// Greeter is the module for Greeter.
type Greeter struct {
	greeter.UnimplementedGreeterServiceServer
	domain greeterdomain.GreeterDomainServiceServer
}

// SayHello implements GreeterServer.
func (g *Greeter) SayHello(ctx context.Context, msg *greeter.SayHelloRequest) (*greeter.SayHelloResponse, error) {
	var SayHelloReq greeterdomain.SayHelloRequest
	if err := copier.CopyMsg(msg, &SayHelloReq); err != nil {
		return nil, err
	}

	sayHelloResp, err := g.domain.SayHello(ctx, &SayHelloReq)
	if err != nil {
		return nil, err
	}

	var resp greeter.SayHelloResponse
	if err := copier.CopyMsg(sayHelloResp, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// New creates a new Greeter module.
func New(biz domain.IGreeter) IGreeter {
	return &Greeter{
		domain: biz,
	}
}
