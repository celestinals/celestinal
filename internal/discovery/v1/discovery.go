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

// Package discovery implements the discovery service
package discovery

import (
	"github.com/celestinals/celestinal/api/gen/go/celestinal/discovery/v1"
	dcvrcontrollers "github.com/celestinals/celestinal/internal/discovery/v1/controllers"
	dcvrdomain "github.com/celestinals/celestinal/internal/discovery/v1/domain"
	dcvrrepo "github.com/celestinals/celestinal/internal/discovery/v1/repos"
	"github.com/celestinals/celestinal/pkg/cache/mem"
	"github.com/celestinals/celestinal/pkg/striker"
)

var (
	// discovery dependency data - repo - service
	_ = striker.Inject(mem.NewDefault[*discovery.Registrar])
	_ = striker.Inject(dcvrrepo.New)
	_ = striker.Inject(dcvrdomain.NewDiscovery)
	_ = striker.Inject(dcvrcontrollers.New)
)
