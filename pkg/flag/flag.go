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
	"github.com/tickexvn/tickex/api/gen/go/tickex/v1"
	"github.com/tickexvn/tickex/internal/utils/version"
)

var (
	once sync.Once
)

var (
	// ascii art use in console with --help option
	asciiConsole = version.ASCIIArt
	isService    = true
)

// flags global variable
var flags = &tickex.Flag{
	Name:    "tickex.server.default",
	Address: "0.0.0.0:9000",
	Mode:    "dev",
}

// EdgeFlags global variable
var edgeFlags = &tickex.FlagEdge{
	IsTurnOnBots: false,
	Secure:       false,
}

// Parse flag args
func Parse() *tickex.Flag {
	once.Do(func() {
		if !isService {
			pflag.BoolVarP(&edgeFlags.IsTurnOnBots, "telegram", "t", edgeFlags.GetIsTurnOnBots(), "turn on telegram system log ?")
			pflag.BoolVarP(&edgeFlags.Secure, "secure", "s", edgeFlags.GetSecure(), "secure api with WAF ?")
			pflag.StringSliceVarP(&edgeFlags.Rules, "rule", "r", edgeFlags.Rules, "owasp crs rules config file")
		}

		pflag.StringVarP(&flags.Name, "name", "n", flags.GetName(), "hostname ?")
		pflag.StringVarP(&flags.Address, "address", "a", flags.GetAddress(), "host address ?")
		pflag.StringVarP(&flags.Mode, "mode", "m", flags.GetMode(), "run mode (dev|prod|sandbox) ?")

		pflag.Usage = func() {
			fmt.Println(asciiConsole)
			fmt.Println("Usage: tickex-<service> [Flags]")
			pflag.PrintDefaults()
			os.Exit(0)
		}

		pflag.Parse()
	})

	return flags
}

// ParseEdge flag args for edge service
func ParseEdge() *tickex.FlagEdge {
	isService = false

	_ = Parse()

	return edgeFlags
}

// SetDefault set default flag values
func SetDefault(name, address, mode string) {
	flags.Name = name
	flags.Address = address
	flags.Mode = mode
}

// SetConsole set console log when --help or -h option
func SetConsole(ascii string) {
	asciiConsole = ascii
}
