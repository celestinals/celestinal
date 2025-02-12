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
	"errors"
	"fmt"
	"regexp"
	"sync"
)

// Namespace represents a namespace.
type Namespace struct {
	spaces map[string]struct{}
	mu     sync.RWMutex
}

// Validate validates a namespace.
func (ns *Namespace) Validate(env, domain, service string) error {
	if env == "" || domain == "" || service == "" {
		return errors.New("ENVIRONMENT, DOMAIN, and SERVICE cannot be empty")
	}

	// Format [ENVIRONMENT].[DOMAIN].[SERVICE]
	re := regexp.MustCompile(`^(?:[a-zA-Z_][a-zA-Z0-9_-]*)\.(?:[a-zA-Z_][a-zA-Z0-9_-]*)\.(?:[a-zA-Z_][a-zA-Z0-9_-]*)$`)
	fullNamespace := fmt.Sprintf("%s.%s.%s", env, domain, service)
	if !re.MatchString(fullNamespace) {
		return errors.New("namespace does not match the format [ENVIRONMENT].[DOMAIN].[SERVICE]")
	}

	return nil
}

// Create creates a new namespace.
func (ns *Namespace) Create(env, domain, service string) (string, error) {
	if err := ns.Validate(env, domain, service); err != nil {
		return "", err
	}

	ns.mu.Lock()
	defer ns.mu.Unlock()

	fullNamespace := fmt.Sprintf("%s.%s.%s", env, domain, service)
	if _, exists := ns.spaces[fullNamespace]; exists {
		return "", errors.New("namespace already exists")
	}

	ns.spaces[fullNamespace] = struct{}{}
	return fullNamespace, nil
}

// Delete deletes a namespace.
func (ns *Namespace) Delete(env, domain, service string) error {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	fullNamespace := fmt.Sprintf("%s.%s.%s", env, domain, service)
	if _, exists := ns.spaces[fullNamespace]; !exists {
		return errors.New("namespace does not exist")
	}

	delete(ns.spaces, fullNamespace)
	return nil
}

// List returns a list of all namespaces.
func (ns *Namespace) List() []string {
	ns.mu.RLock()
	defer ns.mu.RUnlock()

	namespaces := make([]string, 0, len(ns.spaces))
	for ns := range ns.spaces {
		namespaces = append(namespaces, ns)
	}

	return namespaces
}

// Exists checks if a namespace exists.
func (ns *Namespace) Exists(env, domain, service string) bool {
	ns.mu.RLock()
	defer ns.mu.RUnlock()

	fullNamespace := fmt.Sprintf("%s.%s.%s", env, domain, service)
	_, exists := ns.spaces[fullNamespace]
	return exists
}
