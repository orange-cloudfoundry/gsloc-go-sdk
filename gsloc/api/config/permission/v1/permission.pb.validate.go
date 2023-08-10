// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gsloc/api/config/permission/v1/permission.proto

package permission

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on Element with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Element) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Element with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ElementMultiError, or nil if none found.
func (m *Element) ValidateAll() error {
	return m.validate(true)
}

func (m *Element) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ElementType

	// no validation rules for ElementName

	if len(errors) > 0 {
		return ElementMultiError(errors)
	}

	return nil
}

// ElementMultiError is an error wrapping multiple validation errors returned
// by Element.ValidateAll() if the designated constraints aren't met.
type ElementMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ElementMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ElementMultiError) AllErrors() []error { return m }

// ElementValidationError is the validation error returned by Element.Validate
// if the designated constraints aren't met.
type ElementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ElementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ElementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ElementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ElementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ElementValidationError) ErrorName() string { return "ElementValidationError" }

// Error satisfies the builtin error interface
func (e ElementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sElement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ElementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ElementValidationError{}

// Validate checks the field values on ElementPermission with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ElementPermission) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ElementPermission with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ElementPermissionMultiError, or nil if none found.
func (m *ElementPermission) ValidateAll() error {
	return m.validate(true)
}

func (m *ElementPermission) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Role

	if all {
		switch v := interface{}(m.GetElement()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ElementPermissionValidationError{
					field:  "Element",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ElementPermissionValidationError{
					field:  "Element",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetElement()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ElementPermissionValidationError{
				field:  "Element",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ElementPermissionMultiError(errors)
	}

	return nil
}

// ElementPermissionMultiError is an error wrapping multiple validation errors
// returned by ElementPermission.ValidateAll() if the designated constraints
// aren't met.
type ElementPermissionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ElementPermissionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ElementPermissionMultiError) AllErrors() []error { return m }

// ElementPermissionValidationError is the validation error returned by
// ElementPermission.Validate if the designated constraints aren't met.
type ElementPermissionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ElementPermissionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ElementPermissionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ElementPermissionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ElementPermissionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ElementPermissionValidationError) ErrorName() string {
	return "ElementPermissionValidationError"
}

// Error satisfies the builtin error interface
func (e ElementPermissionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sElementPermission.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ElementPermissionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ElementPermissionValidationError{}