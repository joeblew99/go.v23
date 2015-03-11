// This file was auto-generated by the veyron vdl tool.
// Source: builtin.vdl

package vdl

import (
	// VDL system imports
	"fmt"
)

// WireError is the wire representation for the built-in error type.  Errors and
// exceptions in each programming environment are converted to this type to
// ensure wire compatibility.  Generated code for each environment provides
// automatic conversions into idiomatic native representations.
type WireError struct {
	Id        string        // Error Id, used to uniquely identify each error.
	RetryCode WireRetryCode // Retry behavior suggested for the receiver.
	Msg       string        // Error message, may be empty.
	ParamList []*Value      // Variadic parameters contained in the error.
}

func (WireError) __VDLReflect(struct {
	Name string "v.io/v23/vdl.WireError"
}) {
}

// WireRetryCode is the suggested retry behavior for the receiver of an error.
// If the receiver doesn't know how to handle the specific error, it should
// attempt the suggested retry behavior.
type WireRetryCode int

const (
	WireRetryCodeNoRetry WireRetryCode = iota
	WireRetryCodeRetryConnection
	WireRetryCodeRetryRefetch
	WireRetryCodeRetryBackoff
)

// WireRetryCodeAll holds all labels for WireRetryCode.
var WireRetryCodeAll = []WireRetryCode{WireRetryCodeNoRetry, WireRetryCodeRetryConnection, WireRetryCodeRetryRefetch, WireRetryCodeRetryBackoff}

// WireRetryCodeFromString creates a WireRetryCode from a string label.
func WireRetryCodeFromString(label string) (x WireRetryCode, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *WireRetryCode) Set(label string) error {
	switch label {
	case "NoRetry", "noretry":
		*x = WireRetryCodeNoRetry
		return nil
	case "RetryConnection", "retryconnection":
		*x = WireRetryCodeRetryConnection
		return nil
	case "RetryRefetch", "retryrefetch":
		*x = WireRetryCodeRetryRefetch
		return nil
	case "RetryBackoff", "retrybackoff":
		*x = WireRetryCodeRetryBackoff
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in vdl.WireRetryCode", label)
}

// String returns the string label of x.
func (x WireRetryCode) String() string {
	switch x {
	case WireRetryCodeNoRetry:
		return "NoRetry"
	case WireRetryCodeRetryConnection:
		return "RetryConnection"
	case WireRetryCodeRetryRefetch:
		return "RetryRefetch"
	case WireRetryCodeRetryBackoff:
		return "RetryBackoff"
	}
	return ""
}

func (WireRetryCode) __VDLReflect(struct {
	Name string "v.io/v23/vdl.WireRetryCode"
	Enum struct{ NoRetry, RetryConnection, RetryRefetch, RetryBackoff string }
}) {
}

func init() {
	Register((*WireError)(nil))
	Register((*WireRetryCode)(nil))
}
