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

// Package errors provide all type of error in tickex universal
package errors

import (
	"errors"
	"fmt"

	"github.com/tickexvn/tickex/api/gen/go/types/v1"
	"github.com/tickexvn/tickex/internal/version"
)

// Error represents an error
type Error struct {
	Code    types.Errors
	Message string
	Cause   error
}

// Error returns the error message
func (e *Error) Error() string {
	return e.format()
}

// Unwrap returns the cause of the error
func (e *Error) Unwrap() error {
	return e.Cause
}

// Combine combines the error and its cause
func (e *Error) Combine() error {
	if e.Cause == nil {
		return nil
	}

	return errors.New(e.format())
}

func (e *Error) format() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s [%s] %s: %v", version.Header(types.Status_STATUS_E), e.Code.String(), e.Message, e.Cause)
	}

	return fmt.Sprintf("%s [%s] %s", version.Header(types.Status_STATUS_E), e.Code.String(), e.Message)
}

// New creates a new error
func New(code types.Errors, message string, cause error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

// Is checks if the error is of the target code
func Is(err error, targetCode types.Errors) bool {
	var appErr *Error
	if errors.As(err, &appErr) {
		return appErr.Code == targetCode
	}

	return false
}
