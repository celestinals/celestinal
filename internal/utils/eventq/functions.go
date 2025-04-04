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

package eventq

import (
	"context"

	cestlog "github.com/celestinals/celestinal/pkg/logger"
)

// QueueSpace is the size of the queue
const QueueSpace int = 2

var queue = New[string](QueueSpace)

// Subscribe to the event queue with a cestns and a handler function
func Subscribe(ctx context.Context, ns string, handler func(value string) error) {
	go func() {
		ch := queue.Get(ns)

		for {
			select {
			case <-ctx.Done():
				cestlog.Infof("subscription to cestns %s stopped", ns)
				return
			case event, ok := <-ch:
				cestlog.Debugf("received event of cestns: %s", ns)
				if !ok {
					cestlog.Warnf("channel for cestns %s closed", ns)
					return
				}

				func() {
					defer func() {
						if r := recover(); r != nil {
							cestlog.Errorf("panic in handler: %v", r)
						}
					}()

					if err := handler(event); err != nil {
						cestlog.Errorf("failed to handle event: %v", err)
					}
				}()
			}
		}
	}()
}

// Publish to the event queue with a cestns and a value
func Publish(ns string, value string) {
	ch := queue.Get(ns)

	select {
	case ch <- value:
		cestlog.Debugf("published event to cestns: %s", ns)
	default:
		cestlog.Warnf("channel for cestns %s is full, dropping event", ns)
	}
}
