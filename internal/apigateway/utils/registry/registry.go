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

package registry

import (
	"encoding/json"
	"net/http"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	"github.com/celestinals/celestinal/pkg/capsule/capsulehttp"
	"github.com/celestinals/celestinal/pkg/errors"
	"github.com/celestinals/celestinal/pkg/logger"
	"github.com/celestinals/celestinal/pkg/protobuf"
	"github.com/gorilla/schema"
)

func Serve(server capsulehttp.Server, _ *celestinal.Config) {
	sr := &ServiceRegistry{}

	server.HTTPMux().HandleFunc("/discovery/register", sr.Register)
	server.HTTPMux().HandleFunc("/discovery/heartbeat", sr.Heartbeat)
	server.HTTPMux().HandleFunc("/discovery/discover", sr.Discover)
}

type ServiceRegistry struct{}

func (sr *ServiceRegistry) Register(w http.ResponseWriter, r *http.Request) {
	var req celestinal.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Errorf("ServiceRegistry.Register: error when decode %v", err)

		http.Error(w, errors.ErrInvalidData.Error(), http.StatusBadRequest)
		return
	}

	if err := protobuf.Validate(&req); err != nil {
		logger.Errorf("ServiceRegistry.Register: error when validate %v", err)

		http.Error(w, errors.ErrInvalidData.Error(), http.StatusBadRequest)
		return
	}

	//TODO:

	w.WriteHeader(http.StatusOK)
}

func (sr *ServiceRegistry) Heartbeat(w http.ResponseWriter, r *http.Request) {
	var req celestinal.HeartbeatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Errorf("ServiceRegistry.Register: error when decode %v", err)

		http.Error(w, errors.ErrInvalidData.Error(), http.StatusBadRequest)
		return
	}

	if err := protobuf.Validate(&req); err != nil {
		logger.Errorf("ServiceRegistry.Register: error when validate %v", err)

		http.Error(w, errors.ErrInvalidData.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (sr *ServiceRegistry) Discover(w http.ResponseWriter, r *http.Request) {
	var decoder = schema.NewDecoder()
	var req celestinal.DiscoverRequest

	if err := r.ParseForm(); err != nil {
		logger.Errorf("ServiceRegistry.Register: error when parse form %v", err)
		http.Error(w, errors.ErrInvalidData.Error(), http.StatusBadRequest)
		return
	}

	if err := decoder.Decode(&req, r.Form); err != nil {
		logger.Errorf("ServiceRegistry.Register: error when decode %v", err)
		http.Error(w, errors.ErrInvalidData.Error(), http.StatusBadRequest)
		return
	}

	logger.Debugf("ServiceRegistry.Register: decode %s", req.String())

	if err := protobuf.Validate(&req); err != nil {
		logger.Errorf("ServiceRegistry.Register: error when validate %v", err)
		http.Error(w, errors.ErrInvalidData.Error(), http.StatusBadRequest)
		return
	}

	var resp celestinal.DiscoverResponse
	resp.Name = req.Name

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&resp)
}
