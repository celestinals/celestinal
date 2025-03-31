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

package core

import "github.com/tickexvn/tickex/api/gen/go/tickex/v1"

// service is register properties
type service struct {
	Host string
	Port uint32
	Name string
	Tags []string
}

// ServiceInfo is Serve method properties
type ServiceInfo struct {
	Config *tickex.Config
	Addr   string
	Tags   []string
	Name   string
}
