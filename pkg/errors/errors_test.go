package errors

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		code     codes.Code
		expected string
	}{
		{
			name:     "ErrUnspecified",
			err:      ErrUnspecified,
			code:     codes.Unknown,
			expected: unspecified,
		},
		{
			name:     "ErrInternalError",
			err:      ErrInternalError,
			code:     codes.Internal,
			expected: internalError,
		},
		{
			name:     "ErrNotFound",
			err:      ErrNotFound,
			code:     codes.NotFound,
			expected: notFound,
		},
		{
			name:     "ErrUnauthorized",
			err:      ErrUnauthorized,
			code:     codes.Unauthenticated,
			expected: unauthorized,
		},
		{
			name:     "ErrForbidden",
			err:      ErrForbidden,
			code:     codes.PermissionDenied,
			expected: forbidden,
		},
		{
			name:     "ErrInvalidData",
			err:      ErrInvalidData,
			code:     codes.InvalidArgument,
			expected: invalidData,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if status.Code(tt.err) != tt.code {
				t.Errorf("expected code %v, got %v", tt.code, status.Code(tt.err))
			}
			if status.Convert(tt.err).Message() != tt.expected {
				t.Errorf("expected message %v, got %v", tt.expected, status.Convert(tt.err).Message())
			}
		})
	}
}
