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

// Package main provides the entry point for the greeter service.
package main

import (
	"context"

	"github.com/tickexvn/tickex/pkg/configs"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/logger"
	_ "github.com/tickexvn/tickex/x/greeter/v1/cmd/init"
	"github.com/tickexvn/tickex/x/greeter/v1/cmd/server"
)

// Build and run main application with environment variable
// Remember to inject all layers of the application by core.Inject() function
//
// # Example:
//
// _ = core.Inject(controllers.New)
func main() {
	app := core.Build(server.New, configs.Default)
	logger.Fatal(app.Start(context.Background()))
}
