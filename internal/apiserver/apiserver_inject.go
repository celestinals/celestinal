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

// Package apiserver provides the apiserver
package apiserver

import (
	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	discoveryrepo "github.com/celestinals/celestinal/internal/apiserver/apps/discovery/v1/repos"
	discoverysvc "github.com/celestinals/celestinal/internal/apiserver/apps/discovery/v1/services"
	"github.com/celestinals/celestinal/pkg/cache/mem"
	"github.com/celestinals/celestinal/pkg/striker"
)

// inject all dependencies to the apiserver
// This is a dependency injection pattern.
var (
	// discovery dependency data - repo - service
	_ = striker.Inject(mem.NewDefault[*celestinal.Registrar])
	_ = striker.Inject(discoveryrepo.New)
	_ = striker.Inject(discoverysvc.NewDiscoveryService)
)
