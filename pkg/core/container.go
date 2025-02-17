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
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/tickexvn/tickex/api/gen/go/types/v1"
	"github.com/tickexvn/tickex/internal/version"
	"github.com/tickexvn/tickex/pkg/errors"
	"github.com/tickexvn/tickex/pkg/logger"
	"go.uber.org/fx"
)

type container struct {
	engine *fx.App
}

// Start implements IContainer.
func (c *container) Start() error {
	go c.stop()
	if err := c.engine.Start(context.Background()); err != nil {
		errs := errors.New(types.Errors_ERRORS_INTERNAL_ERROR, "START ERROR", err)

		return fmt.Errorf("%s", errs.Error())
	}

	return nil
}

func (c *container) stop() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	logger.Infof("%s Received signal: %v. Stopping application...\n",
		version.Header(types.Status_STATUS_I), sig)

	if err := c.engine.Stop(context.Background()); err != nil {
		logger.Error(errors.New(types.Errors_ERRORS_INTERNAL_ERROR,
			"While stopping app", err).Error())

		os.Exit(1)
		return
	}

	os.Exit(0)
}
