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
	cestcore "github.com/celestinals/celestinal/pkg/core"
)

const timeout = time.Second * 2

// Serve is watching function consul when service info was changed
func Serve(_ cestcore.HTTPServer, config *celestinal.Config) {

	go service()
}

func service() {
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()

	for {

		<-ticker.C
	}
}
