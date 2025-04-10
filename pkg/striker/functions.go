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

package striker

import (
	"context"

	"github.com/celestinals/celestinal/pkg/flag"
	"github.com/celestinals/celestinal/pkg/striker/internal"
	"go.uber.org/fx"
)

// Build builds the application.
// The application is built by providing the constructors.
func Build(constructors ...any) Application {
	for _, constructor := range constructors {
		internal.Provide(constructor)
	}

	// disable log: use fx.NopLogger
	if flag.Parse().GetMode() != "dev" {
		return &striker{
			engine: fx.New(
				internal.Option(), fx.Invoke(runner), fx.NopLogger),
		}
	}

	return &striker{
		engine: fx.New(internal.Option(), fx.Invoke(runner)),
	}
}

// InjectLifeCycle injects the given constructor into the application with
// lifecycle hooks.
func InjectLifeCycle[T any](
	constructor func() T, onStart func(T) error, onStop func(T) error) func() T {

	decorateConstructor := func(lc fx.Lifecycle) T {
		ins := constructor()

		lc.Append(fx.Hook{
			OnStart: func(_ context.Context) error {
				return onStart(ins)
			},
			OnStop: func(_ context.Context) error {
				return onStop(ins)
			},
		})

		return ins
	}

	internal.Provide(decorateConstructor)

	return constructor
}

// Inject injects the given constructor into the application.
func Inject[fn any](constructor fn) fn {
	internal.Provide(constructor)
	return constructor
}
