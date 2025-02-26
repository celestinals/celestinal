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

package core

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
)

// container represents the container with uber/fx frameworks.
// manage the lifecycle of the application.
type container struct {
	engine *fx.App
}

// Start the app with the given context.
func (c *container) Start(ctx context.Context) error {
	err := make(chan error)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	defer close(err)
	defer close(sig)

	// fork the goroutine 1 for start the app
	go c.start(ctx, err)

	// fork the goroutine 2 for stop the app
	go c.stop(ctx, sig, err)

	// wait for the error from the goroutine 1 or 2, end the app
	return <-err
}

// start the app with the given context.
func (c *container) start(ctx context.Context, err chan<- error) {
	// if the error is not nil, return the error to err channel end goroutine 1
	err <- c.engine.Start(ctx)
}

// stop the app with the given context.
func (c *container) stop(ctx context.Context, sig <-chan os.Signal, err chan<- error) {
	// wait for the signal interrupt from the OS
	<-sig

	// if the error is not nil, return the error to err channel, end goroutine 2
	err <- c.engine.Stop(ctx)
}
