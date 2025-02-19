// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: discovery.proto

package discovery

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

// Validate checks the field values on RegisterRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RegisterRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterRequestValidationError{
				field:  "Service",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for StatusPath

	return nil
}

// RegisterRequestValidationError is the validation error returned by
// RegisterRequest.Validate if the designated constraints aren't met.
type RegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterRequestValidationError) ErrorName() string { return "RegisterRequestValidationError" }

// Error satisfies the builtin error interface
func (e RegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterRequestValidationError{}

// Validate checks the field values on HeartbeatRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HeartbeatRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Address

	// no validation rules for Port

	return nil
}

// HeartbeatRequestValidationError is the validation error returned by
// HeartbeatRequest.Validate if the designated constraints aren't met.
type HeartbeatRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeartbeatRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeartbeatRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeartbeatRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeartbeatRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeartbeatRequestValidationError) ErrorName() string { return "HeartbeatRequestValidationError" }

// Error satisfies the builtin error interface
func (e HeartbeatRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeartbeatRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeartbeatRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeartbeatRequestValidationError{}

// Validate checks the field values on DiscoverRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DiscoverRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// DiscoverRequestValidationError is the validation error returned by
// DiscoverRequest.Validate if the designated constraints aren't met.
type DiscoverRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DiscoverRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DiscoverRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DiscoverRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DiscoverRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DiscoverRequestValidationError) ErrorName() string { return "DiscoverRequestValidationError" }

// Error satisfies the builtin error interface
func (e DiscoverRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDiscoverRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DiscoverRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DiscoverRequestValidationError{}

// Validate checks the field values on DiscoverResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DiscoverResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetServices() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DiscoverResponseValidationError{
					field:  fmt.Sprintf("Services[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// DiscoverResponseValidationError is the validation error returned by
// DiscoverResponse.Validate if the designated constraints aren't met.
type DiscoverResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DiscoverResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DiscoverResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DiscoverResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DiscoverResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DiscoverResponseValidationError) ErrorName() string { return "DiscoverResponseValidationError" }

// Error satisfies the builtin error interface
func (e DiscoverResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDiscoverResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DiscoverResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DiscoverResponseValidationError{}

// Validate checks the field values on RegisterResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RegisterResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Message

	return nil
}

// RegisterResponseValidationError is the validation error returned by
// RegisterResponse.Validate if the designated constraints aren't met.
type RegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterResponseValidationError) ErrorName() string { return "RegisterResponseValidationError" }

// Error satisfies the builtin error interface
func (e RegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterResponseValidationError{}

// Validate checks the field values on HeartbeatResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HeartbeatResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	return nil
}

// HeartbeatResponseValidationError is the validation error returned by
// HeartbeatResponse.Validate if the designated constraints aren't met.
type HeartbeatResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeartbeatResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeartbeatResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeartbeatResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeartbeatResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeartbeatResponseValidationError) ErrorName() string {
	return "HeartbeatResponseValidationError"
}

// Error satisfies the builtin error interface
func (e HeartbeatResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeartbeatResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeartbeatResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeartbeatResponseValidationError{}
