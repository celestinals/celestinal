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

// Package main provides the entry point for the greeter service.
package main

import (
	"context"

	greeterpb "github.com/celestinals/celestinal/api/gen/go/celestinal/greeter/v1"
	"github.com/celestinals/celestinal/internal/greeter/v1"
	"github.com/celestinals/celestinal/pkg/config"
	"github.com/celestinals/celestinal/pkg/flag"
	"github.com/celestinals/celestinal/pkg/logger"
	"github.com/celestinals/celestinal/pkg/names"
	"github.com/celestinals/celestinal/pkg/striker"
)

// Build and run main application with environment variable
// Remember to inject all layers of the application by
// striker.Inject() function
//
// Example:
//
// _ = striker.Inject(controllers.New)
func main() {
	flag.SetDefault(names.GreeterV1, "127.0.0.1:0", "dev")
	flag.SetConsole(greeterpb.ASCII)

	_ = flag.Parse()

	app := striker.Build(greeter.New, config.Default)
	if err := app.Run(context.Background()); err != nil {
		logger.Fatal(err)
	}
}
