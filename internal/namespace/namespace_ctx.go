/*
 * Copyright 2024 The Tickex Authors.
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

// Package namespace provides a namespace.
package namespace

import (
	"context"
	"errors"
)

// Key represents a context key.
type Key string

const tkxNamespace Key = "tx.namespace"

// WithNamespace returns a new context with the namespace.
func WithNamespace(ctx context.Context, ns string) context.Context {
	return context.WithValue(ctx, tkxNamespace, ns)
}

// FromContext returns the namespace from the context.
func FromContext(ctx context.Context) (string, error) {
	ns := ctx.Value(tkxNamespace)
	if ns == nil {
		return "", errors.New("namespace not found in context")
	}
	return ns.(string), nil
}
