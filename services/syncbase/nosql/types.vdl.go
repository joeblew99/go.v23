// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package nosql

import (
	// VDL system imports
	"fmt"
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/v23/security/access"
)

// BatchOptions configures a batch.
// TODO(sadovsky): Add more options, e.g. to configure isolation, timeouts,
// whether to track the read set and/or write set, etc.
// TODO(sadovsky): Maybe add a DefaultBatchOptions() function that initializes
// BatchOptions with our desired defaults. Clients would be encouraged to
// initialize their BatchOptions object using that function and then modify it
// to their liking.
type BatchOptions struct {
	// Arbitrary string, typically used to describe the intent behind a batch.
	// Hints are surfaced to clients during conflict resolution.
	// TODO(sadovsky): Use "any" here?
	Hint string
	// ReadOnly specifies whether the batch should allow writes.
	// If ReadOnly is set to true, Abort() should be used to release any resources
	// associated with this batch (though it is not strictly required), and
	// Commit() will always fail.
	ReadOnly bool
}

func (BatchOptions) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.BatchOptions"`
}) {
}

// PrefixPermissions represents a pair of (prefix, perms).
type PrefixPermissions struct {
	Prefix string
	Perms  access.Permissions
}

func (PrefixPermissions) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.PrefixPermissions"`
}) {
}

// KeyValue is a key-value pair.
type KeyValue struct {
	Key   string
	Value []byte
}

func (KeyValue) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.KeyValue"`
}) {
}

// SyncGroupSpec contains the specification for a SyncGroup.
type SyncGroupSpec struct {
	// Human readable description.
	Description string
	// Permissions for the SyncGroup.
	Perms access.Permissions
	// SyncGroup prefixes (relative to the database).  Prefixes
	// must take the form "<tableName>:<rowKeyPrefix>" where
	// tableName is non-empty.
	Prefixes []string
	// Mount tables at which to advertise this SyncGroup. These
	// are the mount tables used for rendezvous in addition to the
	// one in the neighborhood. Typically, we will have only one
	// entry.  However, an array allows mount tables to be changed
	// over time.
	//
	// TODO(hpucha): Figure out a convention for
	// advertising SyncGroups in the mount table.
	MountTables []string
	// Option to change the privacy of the SyncGroup. Configures
	// whether blobs in a SyncGroup can be served to clients
	// holding blobrefs obtained from other SyncGroups.
	IsPrivate bool
}

func (SyncGroupSpec) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.SyncGroupSpec"`
}) {
}

// SyncGroupMemberInfo contains per-member metadata.
type SyncGroupMemberInfo struct {
	SyncPriority byte
}

func (SyncGroupMemberInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.SyncGroupMemberInfo"`
}) {
}

// ResolverType defines the possible conflict resolution policies.
// A Conflict is defined as presence of two independent sets of updates
// originating from the same version of an object. Syncbase
// uses version vectors to determine sequence of changes to a given row. Hence
// if device A updates a row with key "foo" from version V3 to V4, then syncs
// with device B which further updates the same row from version V4 to V5 and
// then V5 is synced back to device A, device A will see V5 as a forward
// progression of "foo" and not a conflict with V3 of "foo". But in the
// meantime if device A had already updated "foo" again from version V4 to
// version V6 then there is a conflict between V5 and V6 with V4 being the
// common ancestor.
type ResolverType int

const (
	ResolverTypeLastWins ResolverType = iota
	ResolverTypeAppResolves
	ResolverTypeDefer
)

// ResolverTypeAll holds all labels for ResolverType.
var ResolverTypeAll = [...]ResolverType{ResolverTypeLastWins, ResolverTypeAppResolves, ResolverTypeDefer}

// ResolverTypeFromString creates a ResolverType from a string label.
func ResolverTypeFromString(label string) (x ResolverType, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *ResolverType) Set(label string) error {
	switch label {
	case "LastWins", "lastwins":
		*x = ResolverTypeLastWins
		return nil
	case "AppResolves", "appresolves":
		*x = ResolverTypeAppResolves
		return nil
	case "Defer", "defer":
		*x = ResolverTypeDefer
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in nosql.ResolverType", label)
}

// String returns the string label of x.
func (x ResolverType) String() string {
	switch x {
	case ResolverTypeLastWins:
		return "LastWins"
	case ResolverTypeAppResolves:
		return "AppResolves"
	case ResolverTypeDefer:
		return "Defer"
	}
	return ""
}

func (ResolverType) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.ResolverType"`
	Enum struct{ LastWins, AppResolves, Defer string }
}) {
}

// ConflictInfo contains information to fully specify a conflict
// for a key, providing the (local, remote, ancestor) tuple.
// A key under conflict can be a part of a batch in local, remote or both
// updates. Since the batches can have more than one key, all ConflictInfos
// for the keys within the batches are grouped together into a single conflict
// batch and sent as a stream with the Continued field representing conflict
// batch boundaries.
type ConflictInfo struct {
	// Data is a unit chunk of ConflictInfo which can be sent over the conflict
	// stream.
	Data ConflictData
	// Continued represents whether the batch of ConflictInfos has ended.
	Continued bool
}

func (ConflictInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.ConflictInfo"`
}) {
}

type (
	// ConflictData represents any single field of the ConflictData union type.
	//
	// ConflictData represents a unit of conflict data sent over the stream. It
	// can either contain information about a Batch or about an operation done
	// on a row.
	ConflictData interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the ConflictData union type.
		__VDLReflect(__ConflictDataReflect)
	}
	// ConflictDataBatch represents field Batch of the ConflictData union type.
	ConflictDataBatch struct{ Value BatchInfo }
	// ConflictDataRow represents field Row of the ConflictData union type.
	ConflictDataRow struct{ Value RowInfo }
	// __ConflictDataReflect describes the ConflictData union type.
	__ConflictDataReflect struct {
		Name  string `vdl:"v.io/v23/services/syncbase/nosql.ConflictData"`
		Type  ConflictData
		Union struct {
			Batch ConflictDataBatch
			Row   ConflictDataRow
		}
	}
)

func (x ConflictDataBatch) Index() int                         { return 0 }
func (x ConflictDataBatch) Interface() interface{}             { return x.Value }
func (x ConflictDataBatch) Name() string                       { return "Batch" }
func (x ConflictDataBatch) __VDLReflect(__ConflictDataReflect) {}

func (x ConflictDataRow) Index() int                         { return 1 }
func (x ConflictDataRow) Interface() interface{}             { return x.Value }
func (x ConflictDataRow) Name() string                       { return "Row" }
func (x ConflictDataRow) __VDLReflect(__ConflictDataReflect) {}

type BatchInfo struct {
	// Id is an identifier for a batch contained in a conflict. It is
	// unique only in the context of a given conflict. Its purpose is solely to
	// group one or more RowInfo objects together to represent a batch that
	// was committed by the client.
	Id uint16
	// Hint is the hint provided by the client when this batch was committed.
	Hint string
	// Source states where the batch comes from.
	Source BatchSource
}

func (BatchInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.BatchInfo"`
}) {
}

// BatchSource represents where the batch was committed.
type BatchSource int

const (
	BatchSourceLocal BatchSource = iota
	BatchSourceRemote
)

// BatchSourceAll holds all labels for BatchSource.
var BatchSourceAll = [...]BatchSource{BatchSourceLocal, BatchSourceRemote}

// BatchSourceFromString creates a BatchSource from a string label.
func BatchSourceFromString(label string) (x BatchSource, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *BatchSource) Set(label string) error {
	switch label {
	case "Local", "local":
		*x = BatchSourceLocal
		return nil
	case "Remote", "remote":
		*x = BatchSourceRemote
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in nosql.BatchSource", label)
}

// String returns the string label of x.
func (x BatchSource) String() string {
	switch x {
	case BatchSourceLocal:
		return "Local"
	case BatchSourceRemote:
		return "Remote"
	}
	return ""
}

func (BatchSource) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.BatchSource"`
	Enum struct{ Local, Remote string }
}) {
}

// RowInfo contains a single operation performed on a row (in case of read or
// write) or a range or rows (in case of scan) along with a mapping to each
// of the batches that this operation belongs to.
// For example, if Row1 was updated on local syncbase conflicting with a write
// on remote syncbase as part of two separate batches, then it will be
// represented by a single RowInfo with Write Operation containing the
// respective local and remote values along with the batch id for both batches
// stored in the BatchIds field.
type RowInfo struct {
	// Op is a specific operation represented by RowInfo
	Op Operation
	// BatchIds contains ids of all batches that this RowInfo is a part of.
	BatchIds []uint16
}

func (RowInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.RowInfo"`
}) {
}

type (
	// Operation represents any single field of the Operation union type.
	//
	// Operation represents a specific operation on a row or a set of rows that is
	// a part of the conflict.
	Operation interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the Operation union type.
		__VDLReflect(__OperationReflect)
	}
	// OperationRead represents field Read of the Operation union type.
	//
	// Read represents a read operation performed on a specific row. For a given
	// row key there can only be at max one Read operation within a conflict.
	OperationRead struct{ Value RowOp }
	// OperationWrite represents field Write of the Operation union type.
	//
	// Write represents a write operation performed on a specific row. For a
	// given row key there can only be at max one Write operation within a
	// conflict.
	OperationWrite struct{ Value RowOp }
	// OperationScan represents field Scan of the Operation union type.
	//
	// Scan represents a scan operation performed over a specific range of keys.
	// For a given key range there can be at max one ScanOp within the Conflict.
	OperationScan struct{ Value ScanOp }
	// __OperationReflect describes the Operation union type.
	__OperationReflect struct {
		Name  string `vdl:"v.io/v23/services/syncbase/nosql.Operation"`
		Type  Operation
		Union struct {
			Read  OperationRead
			Write OperationWrite
			Scan  OperationScan
		}
	}
)

func (x OperationRead) Index() int                      { return 0 }
func (x OperationRead) Interface() interface{}          { return x.Value }
func (x OperationRead) Name() string                    { return "Read" }
func (x OperationRead) __VDLReflect(__OperationReflect) {}

func (x OperationWrite) Index() int                      { return 1 }
func (x OperationWrite) Interface() interface{}          { return x.Value }
func (x OperationWrite) Name() string                    { return "Write" }
func (x OperationWrite) __VDLReflect(__OperationReflect) {}

func (x OperationScan) Index() int                      { return 2 }
func (x OperationScan) Interface() interface{}          { return x.Value }
func (x OperationScan) Name() string                    { return "Scan" }
func (x OperationScan) __VDLReflect(__OperationReflect) {}

// RowOp represents a read or write operation on a row corresponding to the
// given key.
type RowOp struct {
	// The key under conflict.
	Key string
	// LocalValue contains the value read or written by local syncbase or nil.
	LocalValue *Value
	// RemoteValue contains the value read or written by remote syncbase or nil.
	RemoteValue *Value
	// AncestorValue contains the value for the key which is the lowest common
	// ancestor of the two values represented by LocalValue and RemoteValue or
	// nil if no ancestor exists or if the operation was read.
	AncestorValue *Value
}

func (RowOp) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.RowOp"`
}) {
}

// ScanOp provides details of a scan operation.
type ScanOp struct {
	// Start contains the starting key for a range scan.
	Start string
	// Limit contains the end key for a range scan.
	Limit string
}

func (ScanOp) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.ScanOp"`
}) {
}

// Value contains the encoded bytes for a row's value stored in syncbase.
type Value struct {
	// VOM encoded bytes for a row's value
	Bytes []byte
	// Write timestamp for this value in nanoseconds.
	// TODO(jlodhia): change the timestamp to vdl.time here, in commit timestamp
	// and clock data.
	WriteTs int64
}

func (Value) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.Value"`
}) {
}

// ValueSelection represents the value that was selected as the final resolution
// for a conflict.
type ValueSelection int

const (
	ValueSelectionLocal ValueSelection = iota
	ValueSelectionRemote
	ValueSelectionOther
)

// ValueSelectionAll holds all labels for ValueSelection.
var ValueSelectionAll = [...]ValueSelection{ValueSelectionLocal, ValueSelectionRemote, ValueSelectionOther}

// ValueSelectionFromString creates a ValueSelection from a string label.
func ValueSelectionFromString(label string) (x ValueSelection, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *ValueSelection) Set(label string) error {
	switch label {
	case "Local", "local":
		*x = ValueSelectionLocal
		return nil
	case "Remote", "remote":
		*x = ValueSelectionRemote
		return nil
	case "Other", "other":
		*x = ValueSelectionOther
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in nosql.ValueSelection", label)
}

// String returns the string label of x.
func (x ValueSelection) String() string {
	switch x {
	case ValueSelectionLocal:
		return "Local"
	case ValueSelectionRemote:
		return "Remote"
	case ValueSelectionOther:
		return "Other"
	}
	return ""
}

func (ValueSelection) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.ValueSelection"`
	Enum struct{ Local, Remote, Other string }
}) {
}

// ResolutionInfo contains the application’s reply to a conflict for a key,
// providing the resolution value. The resolution may be over a group of keys
// in which case the application must send a stream of ResolutionInfos with
// the Continued field for the last ResolutionInfo representing the end of the
// batch with a value false. ResolutionInfos sent as part of a batch will be
// committed as a batch. If the commit fails, the Conflict will be re-sent.
type ResolutionInfo struct {
	// Key is the key under conflict.
	Key string
	// Selection represents the value that was selected as resolution.
	Selection ValueSelection
	// Result is the resolved value for the key. This field should be used only
	// if value of Selection field is 'Other'
	Result *Value
	// Continued represents whether the batch of ResolutionInfos has ended.
	Continued bool
}

func (ResolutionInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.ResolutionInfo"`
}) {
}

// SchemaMetadata maintains metadata related to the schema of a given database.
// There is one SchemaMetadata per database.
type SchemaMetadata struct {
	// Non negative Schema version number. Should be increased with every schema change
	// (e.g. adding fields to structs) that cannot be handled by previous
	// versions of the app.
	Version int32
	Policy  CrPolicy
}

func (SchemaMetadata) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.SchemaMetadata"`
}) {
}

// For a given row with a conflict, all rules are matched against the row.
// If no rules match the row, we default to "LastWins". If multiple
// rules match the row, ties are broken as follows:
//  1. If one match has a longer prefix than the other, take that one.
//  2. Else, if only one match specifies a type, take that one.
//  3. Else, the two matches are identical; take the last one in the Rules array.
type CrPolicy struct {
	Rules []CrRule
}

func (CrPolicy) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.CrPolicy"`
}) {
}

// CrRule provides a filter and the type of resolution to perform for a row
// under conflict that passes the filter.
type CrRule struct {
	// TableName is the name of the table that this rule applies to.
	TableName string
	// KeyPrefix represents the set of keys within the given table for which
	// this policy applies. TableName must not be empty if this field is set.
	KeyPrefix string
	// Type includes the full package path for the value type for which this
	// policy applies.
	Type string
	// Policy for resolving conflict.
	Resolver ResolverType
}

func (CrRule) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.CrRule"`
}) {
}

// BlobRef is a reference to a blob.
type BlobRef string

func (BlobRef) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.BlobRef"`
}) {
}

// BlobFetchState represents the state transitions of a blob fetch.
type BlobFetchState int

const (
	BlobFetchStatePending BlobFetchState = iota
	BlobFetchStateLocating
	BlobFetchStateFetching
	BlobFetchStateDone
)

// BlobFetchStateAll holds all labels for BlobFetchState.
var BlobFetchStateAll = [...]BlobFetchState{BlobFetchStatePending, BlobFetchStateLocating, BlobFetchStateFetching, BlobFetchStateDone}

// BlobFetchStateFromString creates a BlobFetchState from a string label.
func BlobFetchStateFromString(label string) (x BlobFetchState, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *BlobFetchState) Set(label string) error {
	switch label {
	case "Pending", "pending":
		*x = BlobFetchStatePending
		return nil
	case "Locating", "locating":
		*x = BlobFetchStateLocating
		return nil
	case "Fetching", "fetching":
		*x = BlobFetchStateFetching
		return nil
	case "Done", "done":
		*x = BlobFetchStateDone
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in nosql.BlobFetchState", label)
}

// String returns the string label of x.
func (x BlobFetchState) String() string {
	switch x {
	case BlobFetchStatePending:
		return "Pending"
	case BlobFetchStateLocating:
		return "Locating"
	case BlobFetchStateFetching:
		return "Fetching"
	case BlobFetchStateDone:
		return "Done"
	}
	return ""
}

func (BlobFetchState) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.BlobFetchState"`
	Enum struct{ Pending, Locating, Fetching, Done string }
}) {
}

// BlobFetchStatus describes the progress of an asynchronous blob fetch.
type BlobFetchStatus struct {
	State    BlobFetchState // State of the blob fetch request.
	Received int64          // Total number of bytes received.
	Total    int64          // Blob size.
}

func (BlobFetchStatus) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.BlobFetchStatus"`
}) {
}

// StoreChange is the new value for a watched entity.
// TODO(rogulenko): Consider adding the Shell state.
type StoreChange struct {
	// Value is the new value for the row if the Change state equals to Exists,
	// otherwise the Value is nil.
	Value []byte
	// FromSync indicates whether the change came from sync. If FromSync is
	// false, then the change originated from the local device.
	FromSync bool
}

func (StoreChange) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase/nosql.StoreChange"`
}) {
}

func init() {
	vdl.Register((*BatchOptions)(nil))
	vdl.Register((*PrefixPermissions)(nil))
	vdl.Register((*KeyValue)(nil))
	vdl.Register((*SyncGroupSpec)(nil))
	vdl.Register((*SyncGroupMemberInfo)(nil))
	vdl.Register((*ResolverType)(nil))
	vdl.Register((*ConflictInfo)(nil))
	vdl.Register((*ConflictData)(nil))
	vdl.Register((*BatchInfo)(nil))
	vdl.Register((*BatchSource)(nil))
	vdl.Register((*RowInfo)(nil))
	vdl.Register((*Operation)(nil))
	vdl.Register((*RowOp)(nil))
	vdl.Register((*ScanOp)(nil))
	vdl.Register((*Value)(nil))
	vdl.Register((*ValueSelection)(nil))
	vdl.Register((*ResolutionInfo)(nil))
	vdl.Register((*SchemaMetadata)(nil))
	vdl.Register((*CrPolicy)(nil))
	vdl.Register((*CrRule)(nil))
	vdl.Register((*BlobRef)(nil))
	vdl.Register((*BlobFetchState)(nil))
	vdl.Register((*BlobFetchStatus)(nil))
	vdl.Register((*StoreChange)(nil))
}

const NullBlobRef = BlobRef("")
