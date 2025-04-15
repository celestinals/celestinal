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

// Package watcher is watching service registry when service info was changed
package watcher

import (
	"time"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	"github.com/celestinals/celestinal/pkg/striker/skhttp"
)

const timeout = time.Second * 2

// Watcher is a watcher for service registry when service info was
// changed it will notify the service registry and update the service
// info in the service registry
type Watcher struct{}

// RegisterServer registers the watcher to the striker server
func (w *Watcher) RegisterServer(_ skhttp.Server, _ *celestinal.Config) {
	go service()
}

// service is a service watcher
// it will watch the service registry and update the service info
func service() {
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()

	for {

		<-ticker.C
	}
}
