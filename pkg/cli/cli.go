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
	"fmt"
	"os"
	"sync"

	"github.com/spf13/pflag"
	"github.com/tickexvn/tickex/api/gen/go/common/flags/v1"
	"github.com/tickexvn/tickex/internal/utils/version"
)

var once sync.Once
var isService = true

// ServiceFlags global variable
var ServiceFlags = &flags.ServiceFlags{
	Name:    "Tickex mesh server",
	Address: "0.0.0.0:9000",
}

// EdgeFlags global variable
var EdgeFlags = &flags.EdgeFlags{
	IsTurnOnBots: false,
	Secure:       false,
}

// Parse flag args
func Parse() *flags.ServiceFlags {
	once.Do(func() {
		if !isService {
			pflag.BoolVarP(&EdgeFlags.IsTurnOnBots, "telegram", "t", EdgeFlags.GetIsTurnOnBots(), "turn on telegram system log?")
			pflag.BoolVarP(&EdgeFlags.Secure, "secure", "s", EdgeFlags.GetSecure(), "secure api with WAF?")
			pflag.StringSliceVarP(&EdgeFlags.Rules, "rule", "r", EdgeFlags.Rules, "OWASP CRS rules config file")
		}

		pflag.StringVarP(&ServiceFlags.Name, "name", "n", ServiceFlags.GetName(), "hostname?")
		pflag.StringVarP(&ServiceFlags.Address, "address", "a", ServiceFlags.GetAddress(), "host address?")

		pflag.Usage = func() {
			fmt.Println(version.ASCIIArt)
			fmt.Println("Usage: tickex [Flags]")
			pflag.PrintDefaults()
			os.Exit(0)
		}

		pflag.Parse()
	})

	return ServiceFlags
}

// ParseEdge flag args for edge service
func ParseEdge() *flags.EdgeFlags {
	isService = false

	_ = Parse()

	return EdgeFlags
}
