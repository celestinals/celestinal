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

// Package init provides the initialization logic for the greeter service.
package init

import (
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/x/greeter/internal/controllers"
	"github.com/tickexvn/tickex/x/greeter/internal/domain"
)

var (
	// handlers/controllers layer
	_ = core.Inject(controllers.New)

	// domain layer
	_ = core.Inject(domain.New)

	// repo layer
	//_ = core.Inject(repos.NewAuthor)

	// data layer
	//_ = core.Inject(authors.New)
)
