// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v3alpha/listener/quic_config.proto

package envoy_api_v3alpha_listener

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on QuicProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *QuicProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetMaxConcurrentStreams()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return QuicProtocolOptionsValidationError{
					field:  "MaxConcurrentStreams",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetIdleTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return QuicProtocolOptionsValidationError{
					field:  "IdleTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetCryptoHandshakeTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return QuicProtocolOptionsValidationError{
					field:  "CryptoHandshakeTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// QuicProtocolOptionsValidationError is the validation error returned by
// QuicProtocolOptions.Validate if the designated constraints aren't met.
type QuicProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuicProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuicProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuicProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuicProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuicProtocolOptionsValidationError) ErrorName() string {
	return "QuicProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e QuicProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuicProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuicProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuicProtocolOptionsValidationError{}
