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

// Package flag provide flag variable props
package flag

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/pflag"
	"github.com/tickexvn/tickex/api/gen/go/stdx/v1"
	"github.com/tickexvn/tickex/internal/utils/version"
)

var once sync.Once
var isService = true

// Flags global variable
var Flags = &stdx.Flag{
	Name:    "Tickex mesh server",
	Address: "0.0.0.0:9000",
	Mode:    "dev",
}

// EdgeFlags global variable
var EdgeFlags = &stdx.FlagEdge{
	IsTurnOnBots: false,
	Secure:       false,
}

// Parse flag args
func Parse() *stdx.Flag {
	once.Do(func() {
		if !isService {
			pflag.BoolVarP(&EdgeFlags.IsTurnOnBots, "telegram", "t", EdgeFlags.GetIsTurnOnBots(), "turn on telegram system log ?")
			pflag.BoolVarP(&EdgeFlags.Secure, "secure", "s", EdgeFlags.GetSecure(), "secure api with WAF ?")
			pflag.StringSliceVarP(&EdgeFlags.Rules, "rule", "r", EdgeFlags.Rules, "owasp crs rules config file")
		}

		pflag.StringVarP(&Flags.Name, "name", "n", Flags.GetName(), "hostname ?")
		pflag.StringVarP(&Flags.Address, "address", "a", Flags.GetAddress(), "host address ?")
		pflag.StringVarP(&Flags.Mode, "mode", "m", Flags.GetMode(), "run mode (dev|prod|sandbox) ?")

		pflag.Usage = func() {
			fmt.Println(version.ASCIIArt)
			fmt.Println("Usage: tickex-<service> [Flags]")
			pflag.PrintDefaults()
			os.Exit(0)
		}

		pflag.Parse()
	})

	return Flags
}

// ParseEdge flag args for edge service
func ParseEdge() *stdx.FlagEdge {
	isService = false

	_ = Parse()

	return EdgeFlags
}
