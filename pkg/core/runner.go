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

	"github.com/tickexvn/tickex/internal/version"
	"go.uber.org/fx"
)

// runner functions called by fx.Invoke.
// when the application starts, it will start the server
func runner(lc fx.Lifecycle, srv Server) {
	fmt.Println(version.ASCIIArt)
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			return srv.ListenAndServe()
		},
		OnStop: func(_ context.Context) error {
			return nil
		},
	})
}
