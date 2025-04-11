// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: celestinal/v1/flags.proto

package celestinal

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

// Validate checks the field values on Flag with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Flag) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Address

	// no validation rules for Mode

	// no validation rules for LogLevel

	return nil
}

// FlagValidationError is the validation error returned by Flag.Validate if the
// designated constraints aren't met.
type FlagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FlagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FlagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FlagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FlagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FlagValidationError) ErrorName() string { return "FlagValidationError" }

// Error satisfies the builtin error interface
func (e FlagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFlag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FlagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FlagValidationError{}

// Validate checks the field values on FlagAPIServer with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FlagAPIServer) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Telegram

	// no validation rules for SwaggerPath

	// no validation rules for ApiSpecsPath

	return nil
}

// FlagAPIServerValidationError is the validation error returned by
// FlagAPIServer.Validate if the designated constraints aren't met.
type FlagAPIServerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FlagAPIServerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FlagAPIServerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FlagAPIServerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FlagAPIServerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FlagAPIServerValidationError) ErrorName() string { return "FlagAPIServerValidationError" }

// Error satisfies the builtin error interface
func (e FlagAPIServerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFlagAPIServer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FlagAPIServerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FlagAPIServerValidationError{}
