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

package errors

import (
	"errors"
	"testing"
	
	"github.com/tickexvn/tickex/api/gen/go/types/v1"
	"github.com/tickexvn/tickex/internal/version"
)

func TestNew(t *testing.T) {
	err := New(types.Errors_NOT_FOUND, "Resource not found", nil)

	if err.Code != types.Errors_NOT_FOUND {
		t.Errorf("expected code %s, got %s", types.Errors_NOT_FOUND.String(), err.Code)
	}

	if err.Message != "Resource not found" {
		t.Errorf("expected message 'Resource not found', got '%s'", err.Message)
	}

	if err.Cause != nil {
		t.Errorf("expected nil cause, got %v", err.Cause)
	}
}

func TestErrorFormat(t *testing.T) {
	err := New(types.Errors_INVALID_DATA, "Invalid input", nil)
	expected := version.Header(types.Status_E) + " [INVALID_DATA] Invalid input"

	if err.Error() != expected {
		t.Errorf("expected error string '%s', got '%s'", expected, err.Error())
	}

	t.Log(err.Error())

	wrappedErr := New(types.Errors_INVALID_DATA, "Invalid input", errors.New("missing field"))
	expectedWrapped := version.Header(types.Status_E) + " [INVALID_DATA] Invalid input: missing field"

	if wrappedErr.Error() != expectedWrapped {
		t.Errorf("expected wrapped error string '%s', got '%s'", expectedWrapped, wrappedErr.Error())
	}
}

func TestIs(t *testing.T) {
	err := New(types.Errors_UNAUTHORIZED, "Unauthorized access", nil)

	if !Is(err, types.Errors_UNAUTHORIZED) {
		t.Errorf("expected Is(err, types.Errors_UNAUTHORIZED) to be true, got false")
	}

	if Is(err, types.Errors_NOT_FOUND) {
		t.Errorf("expected Is(err, types.Errors_NOT_FOUND) to be false, got true")
	}
}

func TestUnwrap(t *testing.T) {
	rootErr := errors.New("root cause")
	err := New(types.Errors_INTERNAL_ERROR, "Something went wrong", rootErr)

	unwrappedErr := errors.Unwrap(err)
	if !errors.Is(unwrappedErr, rootErr) {
		t.Errorf("expected unwrapped error to be '%v', got '%v'", rootErr, unwrappedErr)
	}
}

func TestErrorsAs(t *testing.T) {
	var appErr *Error
	err := New(types.Errors_FORBIDDEN, "Access denied", nil)

	if !errors.As(err, &appErr) {
		t.Errorf("expected errors.As to return true, got false")
	}

	if appErr.Code != types.Errors_FORBIDDEN {
		t.Errorf("expected extracted error code %s, got %s", types.Errors_FORBIDDEN.String(), appErr.Code)
	}
}
