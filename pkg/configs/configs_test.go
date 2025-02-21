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

package configs

import (
	"testing"

	"github.com/tickexvn/tickex/api/gen/go/types/v1"
	"github.com/tickexvn/tickex/pkg/pbtools"
)

func TestConfig(t *testing.T) {
	conf := types.Config{
		ServiceRegistryAddress: "0.0.0.0:8500",
		GatewayAddress:         "0.0.0.0:9000",
		Env:                    "prod",
	}

	if err := pbtools.Validate(&conf); err != nil {
		t.Error(err)
	}
}

func TestConfigEnv(t *testing.T) {
	conf := Default()

	if err := pbtools.Validate(conf); err != nil {
		return
	}

	t.Error("should not validate env")
}
