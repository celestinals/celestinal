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

	"github.com/tickexvn/tickex/api/gen/go/tickex/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	forbidden = fmt.Sprintf("TICKEX-%d: %s",
		tickex.Errors_ERRORS_FORBIDDEN, tickex.Errors_ERRORS_FORBIDDEN.String())

	unspecified = fmt.Sprintf("TICKEX-%d: %s",
		tickex.Errors_ERRORS_UNSPECIFIED, tickex.Errors_ERRORS_UNSPECIFIED.String())

	internalError = fmt.Sprintf("TICKEX-%d: %s",
		tickex.Errors_ERRORS_INTERNAL_ERROR, tickex.Errors_ERRORS_INTERNAL_ERROR.String())

	notFound = fmt.Sprintf("TICKEX-%d: %s",
		tickex.Errors_ERRORS_NOT_FOUND, tickex.Errors_ERRORS_NOT_FOUND.String())

	unauthorized = fmt.Sprintf("TICKEX-%d: %s",
		tickex.Errors_ERRORS_UNAUTHORIZED, tickex.Errors_ERRORS_UNAUTHORIZED.String())

	invalidData = fmt.Sprintf("TICKEX-%d: %s",
		tickex.Errors_ERRORS_INVALID_DATA, tickex.Errors_ERRORS_INVALID_DATA.String())

	unimplemented = fmt.Sprintf("TICKEX-%d: %s",
		tickex.Errors_ERRORS_UNIMPLEMENTED, tickex.Errors_ERRORS_UNIMPLEMENTED.String())
)

var (
	// StatusUnspecified is a generic error
	StatusUnspecified = status.Error(codes.Unknown, unspecified)

	// StatusInternalError is an internal error
	StatusInternalError = status.Error(codes.Internal, internalError)

	// StatusNotFound is a not found error
	StatusNotFound = status.Error(codes.NotFound, notFound)

	// StatusUnauthorized is an unauthorized error
	StatusUnauthorized = status.Error(codes.Unauthenticated, unauthorized)

	// StatusForbidden is a forbidden error
	StatusForbidden = status.Error(codes.PermissionDenied, forbidden)

	// StatusInvalidData is an invalid data error
	StatusInvalidData = status.Error(codes.InvalidArgument, invalidData)

	// StatusUnimplemented is a not implemented error
	StatusUnimplemented = status.Error(codes.Unimplemented, unimplemented)
)

var (
	// ErrUnimplemented is a generic error
	ErrUnimplemented = errors.New(unimplemented)
)

// F wrapped error with format template
func F(template string, args ...any) error {
	return fmt.Errorf(template, args...)
}
