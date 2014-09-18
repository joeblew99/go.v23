// This file was auto-generated by the veyron vdl tool.
// Source: types.vdl

// Package types contains the types used by the logreader interface.
package types

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_verror "veyron.io/veyron/veyron2/verror"
)

// LogLine is a log entry from a log file.
type LogEntry struct {
	// The offset (in bytes) where this entry starts.
	Position int64
	// The content of the log entry.
	Line string
}

// A special NumEntries value that indicates that all entries should be
// returned by ReadLog.
const AllEntries = int32(-1)

// This error indicates that the end of the file was reached.
const EOF = _gen_verror.ID("veyron.io/veyron/veyron2/services/mgmt/logreader/types.EOF")
