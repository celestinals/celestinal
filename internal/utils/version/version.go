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

// Package version provides the version of the package.
package version

import (
	"fmt"
	"runtime"
)

const asciiArt = `
 _______     __          
/_  __(_)___/ /_______ __	%s
 / / / / __/  '_/ -_) \ /	------------
/_/ /_/\__/_/\_\\__/_\_\	%s
`

var (

	// Package is filled at linking time
	Package = "github.com/tickexvn/tickex"

	// Version holds the complete version number. Filled in at linking time.
	Version = "v0.0.1-beta"

	// GoVersion is Go tree's version.
	GoVersion = runtime.Version()

	// Name is the full name of the project.
	Name = "TICKEX"

	// BrandName is the brand name of the project.s
	BrandName = "TICKEX // TX"

	// Code is the code of the project.
	Code = "TX"

	// ASCIIArt using in console
	ASCIIArt = fmt.Sprintf(asciiArt,
		BrandName,
		Version,
	)
)

// ASCII prints the ASCII art of the project.
func ASCII() {
	fmt.Println(ASCIIArt)
}
