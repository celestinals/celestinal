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
	"time"

	"github.com/tickexvn/tickex/pkg/flag"
	"github.com/tickexvn/tickex/pkg/protobuf"
	"github.com/tickexvn/tickex/pkg/txlog"
	"go.uber.org/fx"
)

const timeout = 500 * time.Millisecond // 500 milliseconds

// runner functions called by fx.Invoke.
// when the application starts, it will start the server
func runner(lc fx.Lifecycle, srv Server) {
	// init logger
	logger := txlog.NewTxSystemLog()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			errChan := make(chan error, 1)
			go func() {
				if err := srv.ListenAndServe(ctx); err != nil {
					txlog.Warnf("[runner] %+v", err)
					errChan <- err
				}
			}()

			select {
			case err := <-errChan:
				return err
			case <-time.After(timeout):
				return protobuf.Validate(flag.Parse())
			}

		},
		OnStop: func(ctx context.Context) error {
			_ = logger.Sync()

			return srv.Shutdown(ctx)
		},
	})
}
