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

// Package cestflag provide cestflag variable props
package cestflag

import (
	"fmt"
	"os"
	"sync"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	"github.com/celestinals/celestinal/internal/utils/version"

	cestns "github.com/celestinals/celestinal/pkg/names"

	"github.com/spf13/pflag"
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
var flags = &celestinal.Flag{
	Name:    "celestinal.server.default",
	Address: "0.0.0.0:9000",
	Mode:    "dev",
}

// EdgeFlags global variable
var edgeFlags = &celestinal.FlagEdge{
	Telegram: false,
}

// Parse cestflag args
func Parse() *celestinal.Flag {
	once.Do(func() {
		if !isService {
			pflag.BoolVarP(&edgeFlags.Telegram, "telegram", "t", edgeFlags.GetTelegram(), "turn on telegram system log ?")
		}

		pflag.StringVarP(&flags.Name, "name", "n", flags.GetName(), "hostname ?")
		pflag.StringVarP(&flags.Address, "address", "a", flags.GetAddress(), "host address ?")
		pflag.StringVarP(&flags.Mode, "mode", "m", flags.GetMode(), "run mode (dev|prod|sandbox) ?")

		pflag.Usage = func() {
			fmt.Println(asciiConsole)
			fmt.Println("Usage: <service> [Flags]")
			pflag.PrintDefaults()
			os.Exit(0)
		}

		pflag.Parse()
	})

	return flags
}

// ParseEdge cestflag args for apigateway service
func ParseEdge() *celestinal.FlagEdge {
	isService = false

	_ = Parse()

	return edgeFlags
}

// SetDefault set default cestflag values
func SetDefault(name cestns.Namespace, address, mode string) {
	flags.Name = name.String()
	flags.Address = address
	flags.Mode = mode
}

// SetConsole set console log when --help or -h option
func SetConsole(ascii string) {
	asciiConsole = ascii
}
