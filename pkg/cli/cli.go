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

// Package cli provide flag variable props
package cli

import (
	"flag"
	"fmt"
	"sync"

	"github.com/tickexvn/tickex/api/gen/go/types/v1"
)

var once sync.Once
var flags = &types.Flags{
	TurnOnBots: false,
	Name:       "Edge Server",
	Address:    "0.0.0.0:9000",
}

// Parse flag args
func Parse() *types.Flags {
	once.Do(func() {
		flag.BoolVar(&flags.TurnOnBots, "bot", flags.GetTurnOnBots(), "turn on bots?")
		flag.StringVar(&flags.Name, "name", flags.GetName(), "hostname?")
		flag.StringVar(&flags.Address, "address", flags.GetAddress(), "host address?")

		flag.Usage = func() {
			fmt.Println("Usage: tickex [flags]")
			flag.PrintDefaults()
		}
		flag.Parse()
	})

	return flags
}
