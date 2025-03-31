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

// Package main provides the entry point for the Tickex.
package main

import (
	"context"

	"github.com/tickexvn/tickex/internal/edge"
	"github.com/tickexvn/tickex/internal/utils/version"
	"github.com/tickexvn/tickex/pkg/config"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/flag"
	"github.com/tickexvn/tickex/pkg/namespace"
	"github.com/tickexvn/tickex/pkg/txlog"
)

// Build and run the main application with environment variables.
// Remember to inject all layers of the application using the
// core.Inject() function.
//
// Example:
//
//	_ = core.Inject(controllers.New)
//
// This is the Tickex edge application, it will automatically connect to
// other services via gRPC. Start the application along with other services
// in the x/ directory.The application provides APIs for users through a
// single HTTP gateway following the REST API standard. The application
// uses gRPC to connect to other services.Additionally, the system provides
// a Swagger UI interface for users to easily interact with the system
// through a web interface.
//
// Start the application using the Makefile command
//
//	make run.edge // start tickex edge
//	make run.x.<service> // start service
func main() {
	flag.SetDefault(namespace.Edge, "0.0.0.0:9000", "dev")
	flag.SetConsole(version.ASCIIArt)

	_ = flag.ParseEdge()

	app := core.Build(edge.New, config.Default)
	if err := app.Start(context.Background()); err != nil {
		txlog.Fatal(err)
	}
}
