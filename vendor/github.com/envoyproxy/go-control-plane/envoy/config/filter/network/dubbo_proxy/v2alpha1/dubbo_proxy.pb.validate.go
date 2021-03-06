// Code generated by protoc-gen-validate
// source: envoy/config/filter/network/dubbo_proxy/v2alpha1/dubbo_proxy.proto
// DO NOT EDIT!!!

package v2

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

// Validate checks the field values on DubboProxy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *DubboProxy) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetStatPrefix()) < 1 {
		return DubboProxyValidationError{
			Field:  "StatPrefix",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if _, ok := DubboProxy_ProtocolType_name[int32(m.GetProtocolType())]; !ok {
		return DubboProxyValidationError{
			Field:  "ProtocolType",
			Reason: "value must be one of the defined enum values",
		}
	}

	if _, ok := DubboProxy_SerializationType_name[int32(m.GetSerializationType())]; !ok {
		return DubboProxyValidationError{
			Field:  "SerializationType",
			Reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// DubboProxyValidationError is the validation error returned by
// DubboProxy.Validate if the designated constraints aren't met.
type DubboProxyValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e DubboProxyValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDubboProxy.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = DubboProxyValidationError{}
