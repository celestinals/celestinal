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
	"fmt"

	"github.com/tickexvn/tickex/api/gen/go/types/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	forbidden     = fmt.Sprintf("TICKEX-%d: %s", types.Errors_ERRORS_FORBIDDEN, types.Errors_ERRORS_FORBIDDEN.String())
	unspecified   = fmt.Sprintf("TICKEX-%d: %s", types.Errors_ERRORS_UNSPECIFIED, types.Errors_ERRORS_UNSPECIFIED.String())
	internalError = fmt.Sprintf("TICKEX-%d: %s", types.Errors_ERRORS_INTERNAL_ERROR, types.Errors_ERRORS_INTERNAL_ERROR.String())
	notFound      = fmt.Sprintf("TICKEX-%d: %s", types.Errors_ERRORS_NOT_FOUND, types.Errors_ERRORS_NOT_FOUND.String())
	unauthorized  = fmt.Sprintf("TICKEX-%d: %s", types.Errors_ERRORS_UNAUTHORIZED, types.Errors_ERRORS_UNAUTHORIZED.String())
	invalidData   = fmt.Sprintf("TICKEX-%d: %s", types.Errors_ERRORS_INVALID_DATA, types.Errors_ERRORS_INVALID_DATA.String())
)

var (
	// ErrUnspecified is a generic error
	ErrUnspecified = status.Error(codes.Unknown, unspecified)

	// ErrInternalError is an internal error
	ErrInternalError = status.Error(codes.Internal, internalError)

	// ErrNotFound is a not found error
	ErrNotFound = status.Error(codes.NotFound, notFound)

	// ErrUnauthorized is an unauthorized error
	ErrUnauthorized = status.Error(codes.Unauthenticated, unauthorized)

	// ErrForbidden is a forbidden error
	ErrForbidden = status.Error(codes.PermissionDenied, forbidden)

	// ErrInvalidData is an invalid data error
	ErrInvalidData = status.Error(codes.InvalidArgument, invalidData)
)
