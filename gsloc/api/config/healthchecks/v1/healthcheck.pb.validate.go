// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gsloc/api/config/healthchecks/v1/healthcheck.proto

package hcconf

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

	gsloctype "github.com/orange-cloudfoundry/gsloc-go-sdk/gsloc/type/v1"
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

	_ = gsloctype.CodecClientType(0)
)

// Validate checks the field values on HealthCheck with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HealthCheck) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheck with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HealthCheckMultiError, or
// nil if none found.
func (m *HealthCheck) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheck) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTimeout() == nil {
		err := HealthCheckValidationError{
			field:  "Timeout",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetTimeout(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = HealthCheckValidationError{
				field:  "Timeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := HealthCheckValidationError{
					field:  "Timeout",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if m.GetInterval() == nil {
		err := HealthCheckValidationError{
			field:  "Interval",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetInterval(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = HealthCheckValidationError{
				field:  "Interval",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := HealthCheckValidationError{
					field:  "Interval",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if m.GetPort() <= 0 {
		err := HealthCheckValidationError{
			field:  "Port",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTlsConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HealthCheckValidationError{
					field:  "TlsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HealthCheckValidationError{
					field:  "TlsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTlsConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckValidationError{
				field:  "TlsConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.HealthChecker.(type) {

	case *HealthCheck_HttpHealthCheck:

		if all {
			switch v := interface{}(m.GetHttpHealthCheck()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "HttpHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "HttpHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHttpHealthCheck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "HttpHealthCheck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheck_TcpHealthCheck:

		if all {
			switch v := interface{}(m.GetTcpHealthCheck()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "TcpHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "TcpHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTcpHealthCheck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "TcpHealthCheck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheck_GrpcHealthCheck:

		if all {
			switch v := interface{}(m.GetGrpcHealthCheck()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "GrpcHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "GrpcHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetGrpcHealthCheck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "GrpcHealthCheck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheck_NoHealthCheck:

		if all {
			switch v := interface{}(m.GetNoHealthCheck()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "NoHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckValidationError{
						field:  "NoHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetNoHealthCheck()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "NoHealthCheck",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := HealthCheckValidationError{
			field:  "HealthChecker",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return HealthCheckMultiError(errors)
	}

	return nil
}

// HealthCheckMultiError is an error wrapping multiple validation errors
// returned by HealthCheck.ValidateAll() if the designated constraints aren't met.
type HealthCheckMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckMultiError) AllErrors() []error { return m }

// HealthCheckValidationError is the validation error returned by
// HealthCheck.Validate if the designated constraints aren't met.
type HealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckValidationError) ErrorName() string { return "HealthCheckValidationError" }

// Error satisfies the builtin error interface
func (e HealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckValidationError{}

// Validate checks the field values on TlsConfig with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TlsConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TlsConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TlsConfigMultiError, or nil
// if none found.
func (m *TlsConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *TlsConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Enable

	// no validation rules for Ca

	// no validation rules for ServerName

	if len(errors) > 0 {
		return TlsConfigMultiError(errors)
	}

	return nil
}

// TlsConfigMultiError is an error wrapping multiple validation errors returned
// by TlsConfig.ValidateAll() if the designated constraints aren't met.
type TlsConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TlsConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TlsConfigMultiError) AllErrors() []error { return m }

// TlsConfigValidationError is the validation error returned by
// TlsConfig.Validate if the designated constraints aren't met.
type TlsConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TlsConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TlsConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TlsConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TlsConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TlsConfigValidationError) ErrorName() string { return "TlsConfigValidationError" }

// Error satisfies the builtin error interface
func (e TlsConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTlsConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TlsConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TlsConfigValidationError{}

// Validate checks the field values on HealthCheckPayload with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheckPayload) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheckPayload with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheckPayloadMultiError, or nil if none found.
func (m *HealthCheckPayload) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckPayload) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Payload.(type) {

	case *HealthCheckPayload_Text:

		if utf8.RuneCountInString(m.GetText()) < 1 {
			err := HealthCheckPayloadValidationError{
				field:  "Text",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *HealthCheckPayload_Binary:
		// no validation rules for Binary

	default:
		err := HealthCheckPayloadValidationError{
			field:  "Payload",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return HealthCheckPayloadMultiError(errors)
	}

	return nil
}

// HealthCheckPayloadMultiError is an error wrapping multiple validation errors
// returned by HealthCheckPayload.ValidateAll() if the designated constraints
// aren't met.
type HealthCheckPayloadMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckPayloadMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckPayloadMultiError) AllErrors() []error { return m }

// HealthCheckPayloadValidationError is the validation error returned by
// HealthCheckPayload.Validate if the designated constraints aren't met.
type HealthCheckPayloadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckPayloadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckPayloadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckPayloadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckPayloadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckPayloadValidationError) ErrorName() string {
	return "HealthCheckPayloadValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckPayloadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckPayload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckPayloadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckPayloadValidationError{}

// Validate checks the field values on HttpHealthCheck with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *HttpHealthCheck) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HttpHealthCheck with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HttpHealthCheckMultiError, or nil if none found.
func (m *HttpHealthCheck) ValidateAll() error {
	return m.validate(true)
}

func (m *HttpHealthCheck) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetHost() != "" {

		if !_HttpHealthCheck_Host_Pattern.MatchString(m.GetHost()) {
			err := HttpHealthCheckValidationError{
				field:  "Host",
				reason: "value does not match regex pattern \"^[^\\x00\\n\\r]*$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if utf8.RuneCountInString(m.GetPath()) < 1 {
		err := HttpHealthCheckValidationError{
			field:  "Path",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_HttpHealthCheck_Path_Pattern.MatchString(m.GetPath()) {
		err := HttpHealthCheckValidationError{
			field:  "Path",
			reason: "value does not match regex pattern \"^[^\\x00\\n\\r]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetSend()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpHealthCheckValidationError{
					field:  "Send",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpHealthCheckValidationError{
					field:  "Send",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSend()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpHealthCheckValidationError{
				field:  "Send",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetReceive()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpHealthCheckValidationError{
					field:  "Receive",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpHealthCheckValidationError{
					field:  "Receive",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReceive()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpHealthCheckValidationError{
				field:  "Receive",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetRequestHeadersToAdd()) > 1000 {
		err := HttpHealthCheckValidationError{
			field:  "RequestHeadersToAdd",
			reason: "value must contain no more than 1000 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetRequestHeadersToAdd() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HttpHealthCheckValidationError{
						field:  fmt.Sprintf("RequestHeadersToAdd[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HttpHealthCheckValidationError{
						field:  fmt.Sprintf("RequestHeadersToAdd[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpHealthCheckValidationError{
					field:  fmt.Sprintf("RequestHeadersToAdd[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetExpectedStatuses()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpHealthCheckValidationError{
					field:  "ExpectedStatuses",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpHealthCheckValidationError{
					field:  "ExpectedStatuses",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpectedStatuses()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpHealthCheckValidationError{
				field:  "ExpectedStatuses",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := gsloctype.CodecClientType_name[int32(m.GetCodecClientType())]; !ok {
		err := HttpHealthCheckValidationError{
			field:  "CodecClientType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _HttpHealthCheck_Method_NotInLookup[m.GetMethod()]; ok {
		err := HttpHealthCheckValidationError{
			field:  "Method",
			reason: "value must not be in list [6]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := RequestMethod_name[int32(m.GetMethod())]; !ok {
		err := HttpHealthCheckValidationError{
			field:  "Method",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HttpHealthCheckMultiError(errors)
	}

	return nil
}

// HttpHealthCheckMultiError is an error wrapping multiple validation errors
// returned by HttpHealthCheck.ValidateAll() if the designated constraints
// aren't met.
type HttpHealthCheckMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HttpHealthCheckMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HttpHealthCheckMultiError) AllErrors() []error { return m }

// HttpHealthCheckValidationError is the validation error returned by
// HttpHealthCheck.Validate if the designated constraints aren't met.
type HttpHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpHealthCheckValidationError) ErrorName() string { return "HttpHealthCheckValidationError" }

// Error satisfies the builtin error interface
func (e HttpHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpHealthCheckValidationError{}

var _HttpHealthCheck_Host_Pattern = regexp.MustCompile("^[^\x00\n\r]*$")

var _HttpHealthCheck_Path_Pattern = regexp.MustCompile("^[^\x00\n\r]*$")

var _HttpHealthCheck_Method_NotInLookup = map[RequestMethod]struct{}{
	6: {},
}

// Validate checks the field values on TcpHealthCheck with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TcpHealthCheck) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TcpHealthCheck with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TcpHealthCheckMultiError,
// or nil if none found.
func (m *TcpHealthCheck) ValidateAll() error {
	return m.validate(true)
}

func (m *TcpHealthCheck) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSend()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpHealthCheckValidationError{
					field:  "Send",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpHealthCheckValidationError{
					field:  "Send",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSend()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpHealthCheckValidationError{
				field:  "Send",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetReceive() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TcpHealthCheckValidationError{
						field:  fmt.Sprintf("Receive[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TcpHealthCheckValidationError{
						field:  fmt.Sprintf("Receive[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TcpHealthCheckValidationError{
					field:  fmt.Sprintf("Receive[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TcpHealthCheckMultiError(errors)
	}

	return nil
}

// TcpHealthCheckMultiError is an error wrapping multiple validation errors
// returned by TcpHealthCheck.ValidateAll() if the designated constraints
// aren't met.
type TcpHealthCheckMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TcpHealthCheckMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TcpHealthCheckMultiError) AllErrors() []error { return m }

// TcpHealthCheckValidationError is the validation error returned by
// TcpHealthCheck.Validate if the designated constraints aren't met.
type TcpHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TcpHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TcpHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TcpHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TcpHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TcpHealthCheckValidationError) ErrorName() string { return "TcpHealthCheckValidationError" }

// Error satisfies the builtin error interface
func (e TcpHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TcpHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TcpHealthCheckValidationError{}

// Validate checks the field values on GrpcHealthCheck with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GrpcHealthCheck) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrpcHealthCheck with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GrpcHealthCheckMultiError, or nil if none found.
func (m *GrpcHealthCheck) ValidateAll() error {
	return m.validate(true)
}

func (m *GrpcHealthCheck) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ServiceName

	if !_GrpcHealthCheck_Authority_Pattern.MatchString(m.GetAuthority()) {
		err := GrpcHealthCheckValidationError{
			field:  "Authority",
			reason: "value does not match regex pattern \"^[^\\x00\\n\\r]*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GrpcHealthCheckMultiError(errors)
	}

	return nil
}

// GrpcHealthCheckMultiError is an error wrapping multiple validation errors
// returned by GrpcHealthCheck.ValidateAll() if the designated constraints
// aren't met.
type GrpcHealthCheckMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrpcHealthCheckMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrpcHealthCheckMultiError) AllErrors() []error { return m }

// GrpcHealthCheckValidationError is the validation error returned by
// GrpcHealthCheck.Validate if the designated constraints aren't met.
type GrpcHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcHealthCheckValidationError) ErrorName() string { return "GrpcHealthCheckValidationError" }

// Error satisfies the builtin error interface
func (e GrpcHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcHealthCheckValidationError{}

var _GrpcHealthCheck_Authority_Pattern = regexp.MustCompile("^[^\x00\n\r]*$")

// Validate checks the field values on NoHealthCheck with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NoHealthCheck) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NoHealthCheck with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NoHealthCheckMultiError, or
// nil if none found.
func (m *NoHealthCheck) ValidateAll() error {
	return m.validate(true)
}

func (m *NoHealthCheck) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return NoHealthCheckMultiError(errors)
	}

	return nil
}

// NoHealthCheckMultiError is an error wrapping multiple validation errors
// returned by NoHealthCheck.ValidateAll() if the designated constraints
// aren't met.
type NoHealthCheckMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NoHealthCheckMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NoHealthCheckMultiError) AllErrors() []error { return m }

// NoHealthCheckValidationError is the validation error returned by
// NoHealthCheck.Validate if the designated constraints aren't met.
type NoHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NoHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NoHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NoHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NoHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NoHealthCheckValidationError) ErrorName() string { return "NoHealthCheckValidationError" }

// Error satisfies the builtin error interface
func (e NoHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNoHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NoHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NoHealthCheckValidationError{}
