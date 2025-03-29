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

// Package config provides the configs for the service.
package config

import (
	"os"
	"strconv"
	"sync"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/tickexvn/tickex/api/gen/go/tickex/v1"
)

var conf *tickex.Config
var once sync.Once

// Default returns the environment.
func Default() *tickex.Config {
	once.Do(func() {
		sChatID := os.Getenv(tickex.TickexPublic_TICKEX_PUBLIC_CHAT_ID.String())
		ID, _ := strconv.ParseInt(sChatID, 10, 64)

		conf = &tickex.Config{
			ServiceRegistryAddr: os.Getenv(tickex.TickexPublic_TICKEX_PUBLIC_SERVICE_REGISTRY_ADDR.String()),
			ApiAddr:             os.Getenv(tickex.TickexPublic_TICKEX_PUBLIC_API_ADDR.String()),
			Env:                 os.Getenv(tickex.TickexPublic_TICKEX_PUBLIC_ENV.String()),
			BotToken:            os.Getenv(tickex.TickexPublic_TICKEX_PUBLIC_BOT_TOKEN.String()),
			ChatId:              ID,
		}
	})

	return conf
}
