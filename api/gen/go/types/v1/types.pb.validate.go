// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: types.proto

package types

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

// Validate checks the field values on Pages with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Pages) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Index

	// no validation rules for Size

	// no validation rules for Total

	return nil
}

// PagesValidationError is the validation error returned by Pages.Validate if
// the designated constraints aren't met.
type PagesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PagesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PagesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PagesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PagesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PagesValidationError) ErrorName() string { return "PagesValidationError" }

// Error satisfies the builtin error interface
func (e PagesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPages.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PagesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PagesValidationError{}

// Validate checks the field values on Metadata with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Metadata) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetadataValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Author

	return nil
}

// MetadataValidationError is the validation error returned by
// Metadata.Validate if the designated constraints aren't met.
type MetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataValidationError) ErrorName() string { return "MetadataValidationError" }

// Error satisfies the builtin error interface
func (e MetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataValidationError{}

// Validate checks the field values on RobotMessage with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RobotMessage) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RobotMessageValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Header

	// no validation rules for Body

	// no validation rules for Footer

	return nil
}

// RobotMessageValidationError is the validation error returned by
// RobotMessage.Validate if the designated constraints aren't met.
type RobotMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RobotMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RobotMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RobotMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RobotMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RobotMessageValidationError) ErrorName() string { return "RobotMessageValidationError" }

// Error satisfies the builtin error interface
func (e RobotMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRobotMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RobotMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RobotMessageValidationError{}

// Validate checks the field values on Context with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Context) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TraceId

	// no validation rules for RequestId

	// no validation rules for UserId

	// no validation rules for Authorization

	// no validation rules for Locale

	// no validation rules for CorrelationId

	// no validation rules for ServiceName

	// no validation rules for RetryCount

	// no validation rules for SpanId

	// no validation rules for Ip

	// no validation rules for Environment

	return nil
}

// ContextValidationError is the validation error returned by Context.Validate
// if the designated constraints aren't met.
type ContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContextValidationError) ErrorName() string { return "ContextValidationError" }

// Error satisfies the builtin error interface
func (e ContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContextValidationError{}

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ServiceRegistryAddress

	// no validation rules for GatewayAddress

	// no validation rules for Env

	// no validation rules for BotToken

	// no validation rules for ChatId

	return nil
}

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}

// Validate checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Service) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Address

	// no validation rules for Port

	return nil
}

// ServiceValidationError is the validation error returned by Service.Validate
// if the designated constraints aren't met.
type ServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceValidationError) ErrorName() string { return "ServiceValidationError" }

// Error satisfies the builtin error interface
func (e ServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceValidationError{}

// Validate checks the field values on Flags with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Flags) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TurnOnBots

	// no validation rules for Name

	// no validation rules for Address

	return nil
}

// FlagsValidationError is the validation error returned by Flags.Validate if
// the designated constraints aren't met.
type FlagsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FlagsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FlagsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FlagsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FlagsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FlagsValidationError) ErrorName() string { return "FlagsValidationError" }

// Error satisfies the builtin error interface
func (e FlagsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFlags.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FlagsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FlagsValidationError{}
