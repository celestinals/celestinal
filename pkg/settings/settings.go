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

// Package settings provides the settings for the service.
package settings

import (
	"os"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/tickexvn/tickex/api/gen/go/types/v1"
)

// DefaultConfig returns the environment.
func DefaultConfig() *types.Config {
	return &types.Config{
		ServiceRegistryAddress: os.Getenv("SERVICE_REGISTRY_ADDRESS"),
		GatewayAddress:         os.Getenv("GATEWAY_ADDRESS"),
	}
}
