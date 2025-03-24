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

package eventq

import (
	"context"

	"github.com/tickexvn/tickex/pkg/txlog"
)

// QueueSpace is the size of the queue
const QueueSpace int = 2

var queue = New[string](QueueSpace)

// Subscribe to the event queue with a namespace and a handler function
func Subscribe(ctx context.Context, namespace string, handler func(value string) error) {
	go func() {
		ch := queue.Get(namespace)

		for {
			select {
			case <-ctx.Done():
				txlog.Infof("subscription to namespace %s stopped", namespace)
				return
			case event, ok := <-ch:
				txlog.Debugf("received event of namespace: %s", namespace)
				if !ok {
					txlog.Warnf("channel for namespace %s closed", namespace)
					return
				}

				func() {
					defer func() {
						if r := recover(); r != nil {
							txlog.Errorf("panic in handler: %v", r)
						}
					}()

					if err := handler(event); err != nil {
						txlog.Errorf("failed to handle event: %v", err)
					}
				}()
			}
		}
	}()
}

// Publish to the event queue with a namespace and a value
func Publish(namespace string, value string) {
	ch := queue.Get(namespace)

	select {
	case ch <- value:
		txlog.Debugf("published event to namespace: %s", namespace)
	default:
		txlog.Warnf("channel for namespace %s is full, dropping event", namespace)
	}
}
