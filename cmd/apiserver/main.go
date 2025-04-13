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

// Package main provides the entry point for the Celestinal.
package main

import (
	"context"

	"github.com/celestinals/celestinal/internal/apiserver"
	"github.com/celestinals/celestinal/internal/pkg/version"

	"github.com/celestinals/celestinal/pkg/config"
	"github.com/celestinals/celestinal/pkg/flag"
	"github.com/celestinals/celestinal/pkg/logger"
	"github.com/celestinals/celestinal/pkg/names"
	"github.com/celestinals/celestinal/pkg/striker"
)

// Build and run the main application with environment variables.
// Remember to inject all layers of the application using the
// striker.Inject() function.
//
// Example:
//
//	_ = striker.Inject(controllers.New)
//
// This is the celestinal apiserver application, it will automatically connect to
// other services via gRPC. Run the application along with other services
// in the x/ directory.The application provides APIs for users through a
// single HTTP gateway following the REST API standard. The application
// uses gRPC to connect to other services.Additionally, the system provides
// a Swagger UI interface for users to easily interact with the system
// through a web interface.
//
// Run the application using the Makefile command
//
//	make run.apiserver // start celestinal apiserver
//	make run.<service> // start service
func main() {
	flag.SetDefault(names.APIServer, "dev")
	flag.SetConsole(version.ASCIIArt)

	if err := flag.Validate(flag.ParseAPIServer()); err != nil {
		logger.Fatal(err)
	}

	app := striker.Build(apiserver.New, config.Default)
	if err := app.Run(context.Background()); err != nil {
		logger.Fatal(err)
	}
}
