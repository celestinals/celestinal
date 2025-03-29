// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: tickex/stdx/v1/options.proto

package stdx

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on TickexMethodOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TickexMethodOptions) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Ignore

	for idx, item := range m.GetRequire() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TickexMethodOptionsValidationError{
					field:  fmt.Sprintf("Require[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TickexMethodOptionsValidationError is the validation error returned by
// TickexMethodOptions.Validate if the designated constraints aren't met.
type TickexMethodOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TickexMethodOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TickexMethodOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TickexMethodOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TickexMethodOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TickexMethodOptionsValidationError) ErrorName() string {
	return "TickexMethodOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e TickexMethodOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTickexMethodOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TickexMethodOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TickexMethodOptionsValidationError{}

// Validate checks the field values on Require with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Require) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Role

	// no validation rules for Permission

	return nil
}

// RequireValidationError is the validation error returned by Require.Validate
// if the designated constraints aren't met.
type RequireValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequireValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequireValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequireValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequireValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequireValidationError) ErrorName() string { return "RequireValidationError" }

// Error satisfies the builtin error interface
func (e RequireValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequire.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequireValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequireValidationError{}
