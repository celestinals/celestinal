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

package discovery

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
)

type Discovery interface {
	Heartbeat(ctx context.Context, req *celestinal.HeartbeatRequest) error
	Register(ctx context.Context, req *celestinal.RegisterRequest) error
	Discover(ctx context.Context, req *celestinal.DiscoverRequest) (*celestinal.DiscoverResponse, error)
}

func New(baseUrl string) Discovery {
	return &discovery{
		baseUrl: baseUrl,
		client:  &http.Client{},
	}
}

type discovery struct {
	baseUrl string
	client  *http.Client
}

func (d *discovery) Heartbeat(ctx context.Context, req *celestinal.HeartbeatRequest) error {
	return d.post("/discovery/heartbeat", req)
}

func (d *discovery) Register(ctx context.Context, req *celestinal.RegisterRequest) error {
	return d.post("/discovery/register", RegisterRequest{
		Name:    req.Name,
		Address: req.Address,
		TTL:     fmt.Sprintf("%ds", req.GetTtl().Seconds),
	})

}

func (d *discovery) Discover(ctx context.Context, req *celestinal.DiscoverRequest) (*celestinal.DiscoverResponse, error) {
	resp, err := d.client.Get(fmt.Sprintf("%s/discovery/discover?name=%s", d.baseUrl, req.GetName()))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var result celestinal.DiscoverResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}

func (d *discovery) post(path string, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := d.client.Post(d.baseUrl+path, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %s", resp.Status)
	}
	return nil
}
