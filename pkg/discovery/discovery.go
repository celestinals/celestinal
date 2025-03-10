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
	"github.com/tickexvn/tickex/api/gen/go/common/discovery/v1"
	configpb "github.com/tickexvn/tickex/api/gen/go/common/env/config/v1"
	servicepb "github.com/tickexvn/tickex/api/gen/go/common/service/v1"
	"github.com/tickexvn/tickex/pkg/pbtools"
)

var _ discovery.DiscoveryServiceServer = (*Discovery)(nil)

// New provide service registry of Tickex microservice network
func New(conf *configpb.Config) (*Discovery, error) {
	if err := pbtools.Validate(conf); err != nil {
		return nil, err
	}

	config := api.DefaultConfig()
	config.Address = conf.GetServiceRegistryAddr()

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
func (d *Discovery) Register(
	ctx context.Context, req *discovery.RegisterRequest) (*discovery.RegisterResponse, error) {

	_ = ctx
	if err := pbtools.Validate(req); err != nil {
		return nil, err
	}

	serviceInfo := req.GetService()
	registration := &api.AgentServiceRegistration{
		ID:      serviceInfo.GetId(),
		Name:    serviceInfo.GetName(),
		Address: serviceInfo.GetHost(),
		Port:    int(serviceInfo.GetPort()),
		Tags:    serviceInfo.GetTags(),
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: req.GetServiceCheck().GetDeregisterCriticalServiceAfter(),
			TTL:                            req.GetServiceCheck().GetTtl(),
			TLSSkipVerify:                  req.GetServiceCheck().GetTlsSkipVerify(),
		},
	}

	return nil, d.client.Agent().ServiceRegister(registration)
}

// Discover implements the Discover method of the ServiceRegistryService.
func (d *Discovery) Discover(
	ctx context.Context, req *discovery.DiscoverRequest) (*discovery.DiscoverResponse, error) {
	_ = ctx
	if err := pbtools.Validate(req); err != nil {
		return nil, err
	}

	services, _, err := d.client.Health().Service(req.GetName(), "", true, nil)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, fmt.Errorf("service %s not found", req.GetName())
	}

	var resp discovery.DiscoverResponse
	for _, service := range services {
		resp.Services = append(resp.Services, &servicepb.Service{
			Id:   service.Service.ID,
			Name: service.Service.Service,
			Host: service.Service.Address,
			Port: uint32(service.Service.Port),
			Tags: service.Service.Tags,
		})
	}

	return &resp, nil
}

// Heartbeat implements the Heartbeat method of the ServiceRegistryService.
func (d *Discovery) Heartbeat(
	ctx context.Context, req *discovery.HeartbeatRequest) (*discovery.HeartbeatResponse, error) {

	_ = ctx
	if err := pbtools.Validate(req); err != nil {
		return nil, err
	}

	err := d.client.Agent().
		UpdateTTL("service:"+req.GetId(), "Service is healthy", api.HealthPassing)

	return &discovery.HeartbeatResponse{
		Success: err == nil,
	}, err
}
