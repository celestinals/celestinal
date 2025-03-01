// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: transfermarkt_shared.proto

package shared

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

// Validate checks the field values on SayHelloResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SayHelloResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Message

	return nil
}

// SayHelloResponseValidationError is the validation error returned by
// SayHelloResponse.Validate if the designated constraints aren't met.
type SayHelloResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayHelloResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayHelloResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayHelloResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayHelloResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayHelloResponseValidationError) ErrorName() string { return "SayHelloResponseValidationError" }

// Error satisfies the builtin error interface
func (e SayHelloResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayHelloResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayHelloResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayHelloResponseValidationError{}
