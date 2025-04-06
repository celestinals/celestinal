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

// Package controllers provides the controller for the greeter service.
package controllers

import (
	"context"

	"github.com/celestinals/celestinal/staging/mod/greeter/domain"

	domainpb "github.com/celestinals/celestinal/api/gen/go/celestinal/greeter/domain/v1"
	"github.com/celestinals/celestinal/api/gen/go/celestinal/greeter/v1"

	cestcopier "github.com/celestinals/celestinal/pkg/copier"
	cesterr "github.com/celestinals/celestinal/pkg/errors"
	cestlog "github.com/celestinals/celestinal/pkg/logger"
)

// IGreeter defines the interface for the Greeter controller.
type IGreeter interface {
	greeter.GreeterServiceServer
}

// Greeter is the module for Greeter.
type Greeter struct {
	greeter.UnimplementedGreeterServiceServer
	domain domainpb.GreeterDomainServiceServer
}

// SayHello implements GreeterServer.
func (g *Greeter) SayHello(
	ctx context.Context, msg *greeter.SayHelloRequest) (*greeter.SayHelloResponse, error) {

	var SayHelloReq domainpb.SayHelloRequest
	if err := cestcopier.CopyProtoMessage(msg, &SayHelloReq); err != nil {
		cestlog.Error(err)

		return nil, cesterr.StatusInvalidData
	}

	sayHelloResp, err := g.domain.SayHello(ctx, &SayHelloReq)
	if err != nil {
		return nil, err
	}

	return &greeter.SayHelloResponse{
		Response: sayHelloResp.Response,
	}, nil
}

// Status healthcheck for consul
func (g *Greeter) Status(
	ctx context.Context, msg *greeter.StatusRequest) (*greeter.StatusResponse, error) {

	_ = ctx
	_ = msg

	return &greeter.StatusResponse{
		Message: "ok",
	}, nil
}

// New creates a new Greeter module.
func New(biz domain.IGreeter) IGreeter {
	return &Greeter{
		domain: biz,
	}
}
