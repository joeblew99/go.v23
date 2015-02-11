// This file was auto-generated by the veyron vdl tool.
// Source: types.vdl

// Package types contains the types used by the logreader interface.
package types

import (
	// VDL system imports
	"v.io/core/veyron2/context"
	"v.io/core/veyron2/i18n"
	"v.io/core/veyron2/vdl"
	"v.io/core/veyron2/verror"
)

// LogLine is a log entry from a log file.
type LogEntry struct {
	// The offset (in bytes) where this entry starts.
	Position int64
	// The content of the log entry.
	Line string
}

func (LogEntry) __VDLReflect(struct {
	Name string "v.io/core/veyron2/services/mgmt/logreader/types.LogEntry"
}) {
}

func init() {
	vdl.Register((*LogEntry)(nil))
}

// A special NumEntries value that indicates that all entries should be
// returned by ReadLog.
const AllEntries = int32(-1)

var (
	ErrEOF = verror.Register("v.io/core/veyron2/services/mgmt/logreader/types.EOF", verror.NoRetry, "{1:}{2:} EOF")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrEOF.ID), "{1:}{2:} EOF")
}

// NewErrEOF returns an error with the ErrEOF ID.
func NewErrEOF(ctx *context.T) error {
	return verror.New(ErrEOF, ctx)
}
