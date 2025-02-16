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
	"strings"

	"github.com/tickexvn/tickex/api/gen/go/types/v1"
)

var (
	// Package is filled at linking time
	Package = "github.com/tickexvn/tickex"

	// Version holds the complete version number. Filled in at linking time.
	Version = "0.0.1"

	// GoVersion is Go tree's version.
	GoVersion = strings.ToUpper(runtime.Version())

	// FullName is the full name of the project.
	FullName = "TICKEX"

	// Code is the code of the project.
	Code = "TKX"
)

// Header returns the header info string.
func Header(status types.Status) string {
	return fmt.Sprintf("%s<%s<<%s<<%s>>>>>", getStatusSuffix(status), FullName, Version, GoVersion)
}

// getStatusSuffix returns the suffix of the status.
func getStatusSuffix(status types.Status) string {
	arr := strings.Split(status.String(), "_")
	if len(arr) < 2 {
		return ""
	}

	return arr[len(arr)-1]
}
