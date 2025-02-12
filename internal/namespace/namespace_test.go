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

package namespace

import (
	"context"
	"testing"
)

func TestCreateNamespace(t *testing.T) {
	ns := &Namespace{
		spaces: make(map[string]struct{}),
	}

	_, err := ns.Create("1dev", "tkx", "gateway")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	_, err = ns.Create("dev", "tkx", "gateway")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = ns.Create("dev", "tkx", "gateway")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	for _, _ns := range ns.List() {
		t.Logf("%s\n", _ns)
	}
}

func TestValidateNamespace(t *testing.T) {
	ns := &Namespace{
		spaces: make(map[string]struct{}),
	}

	err := ns.Validate("dev", "tkx", "gateway")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = ns.Validate("", "tkx", "gateway")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	err = ns.Validate("dev", "tkx", "")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	err = ns.Validate("dev", "tkx", "gateway.")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}

func TestDeleteNamespace(t *testing.T) {
	ns := &Namespace{
		spaces: make(map[string]struct{}),
	}

	_, err := ns.Create("dev", "tkx", "gateway")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = ns.Delete("dev", "tkx", "gateway")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = ns.Delete("dev", "tkx", "gateway")
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}

func TestExistsNamespace(t *testing.T) {
	ns := &Namespace{
		spaces: make(map[string]struct{}),
	}

	_, err := ns.Create("dev", "tkx", "gateway")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exists := ns.Exists("dev", "tkx", "gateway")
	if !exists {
		t.Fatalf("expected namespace to exist, got false")
	}

	exists = ns.Exists("dev", "tkx", "gateway1")
	if exists {
		t.Fatalf("expected namespace to not exist, got true")
	}
}

func TestListNamespace(t *testing.T) {
	ns := &Namespace{
		spaces: make(map[string]struct{}),
	}

	_, err := ns.Create("dev", "tkx", "gateway")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	namespaces := ns.List()
	if len(namespaces) != 1 {
		t.Fatalf("expected 1 namespace, got %d", len(namespaces))
	}

}

func TestWithContext(t *testing.T) {
	ctx := WithNamespace(context.Background(), "dev.tx.gateway")
	ns, err := FromContext(ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	t.Logf("%s\n", ns)
}

func TestFromContextWithNil(t *testing.T) {
	_, err := FromContext(context.Background())
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	t.Logf("%v\n", err)
}
