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

	"github.com/tickexvn/tickex/api/gen/go/common/flags/v1"
	"github.com/tickexvn/tickex/internal/version"
)

var once sync.Once
var isService = true

// ServiceFlags global variable
var ServiceFlags = &flags.ServiceFlags{
	Name:    "Tickex mesh server",
	Address: "0.0.0.0:9000",
}

var EdgeFlags = &flags.EdgeFlags{
	IsTurnOnBots: false,
	Secure:       false,
}

// Parse flag args
func Parse() *flags.ServiceFlags {
	once.Do(func() {
		if !isService {
			flag.BoolVar(&EdgeFlags.IsTurnOnBots, "bot", EdgeFlags.GetIsTurnOnBots(), "turn on bots?")
			flag.BoolVar(&EdgeFlags.Secure, "secure", EdgeFlags.GetSecure(), "secure api with WAF?")
		}

		flag.StringVar(&ServiceFlags.Name, "name", ServiceFlags.GetName(), "hostname?")
		flag.StringVar(&ServiceFlags.Address, "address", ServiceFlags.GetAddress(), "host address?")

		flag.Usage = func() {
			fmt.Println(version.AsciiArt)
			fmt.Println("Usage: tickex [Flags]")
			flag.PrintDefaults()
		}
		flag.Parse()
	})

	return ServiceFlags
}

func ParseEdge() *flags.EdgeFlags {
	isService = false

	_ = Parse()

	return EdgeFlags
}
