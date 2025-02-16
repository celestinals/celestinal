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

// Package discovery provides the function for the service registry.
package discovery

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/tickexvn/tickex/api/gen/go/types/v1"
	"github.com/tickexvn/tickex/api/gen/go/utils/discovery/v1"
	"github.com/tickexvn/tickex/pkg/pbtools"
)

var _ discovery.DiscoveryServiceServer = (*Discovery)(nil)

// New provide service registry of Tick microservice network
func New(conf *types.Config) (*Discovery, error) {
	_ = conf
	config := api.DefaultConfig()

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Discovery{
		client: client,
	}, nil
}

// Discovery implements the ServiceRegistryService.
type Discovery struct {
	discovery.UnimplementedDiscoveryServiceServer
	client *api.Client
}

// Register implements the Register method of the ServiceRegistryService.
func (d *Discovery) Register(_ context.Context, req *discovery.RegisterRequest) (*discovery.RegisterResponse, error) {
	if err := pbtools.Validate(req); err != nil {
		return nil, err
	}

	healthcheck := fmt.Sprintf("%s:%d%s", req.GetAddress(), req.GetPort(), req.GetStatusPath())

	registration := &api.AgentServiceRegistration{
		ID:      req.GetId(),
		Name:    req.GetName(),
		Address: req.GetAddress(),
		Port:    int(req.GetPort()),
		Tags:    req.GetTags(),
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + healthcheck,
			Interval: "10s",
			Timeout:  "5s",
		},
	}

	return nil, d.client.Agent().ServiceRegister(registration)
}

// Discover implements the Discover method of the ServiceRegistryService.
func (d *Discovery) Discover(_ context.Context, _ *discovery.DiscoverRequest) (*discovery.DiscoverResponse, error) {
	//TODO implement me
	panic("implement me")
}

// Heartbeat implements the Heartbeat method of the ServiceRegistryService.
func (d *Discovery) Heartbeat(_ context.Context, _ *discovery.HeartbeatRequest) (*discovery.HeartbeatResponse, error) {
	//TODO implement me
	panic("implement me")
}
