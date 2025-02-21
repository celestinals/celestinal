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

	"github.com/tickexvn/tickex/internal/version"
	"github.com/tickexvn/tickex/pkg/logger"
	"go.uber.org/fx"
)

type container struct {
	engine *fx.App
}

// Start implements IContainer.
func (c *container) Start(ctx context.Context) error {
	go c.stop(ctx)
	if err := c.engine.Start(ctx); err != nil {
		return err
	}

	return nil
}

func (c *container) stop(ctx context.Context) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan

	logger.Infof("%s Received signal: %v. Stopping application...\n",
		version.Header(), sig)

	if err := c.engine.Stop(ctx); err != nil {
		logger.Error(err)

		os.Exit(1)
		return
	}

	os.Exit(0)
}
