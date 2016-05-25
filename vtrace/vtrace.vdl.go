// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: vtrace

package vtrace

import (
	"time"
	"v.io/v23/uniqueid"
	"v.io/v23/vdl"
	vdltime "v.io/v23/vdlroot/time"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// An Annotation represents data that is relevant at a specific moment.
// They can be attached to spans to add useful debugging information.
type Annotation struct {
	// When the annotation was added.
	When time.Time
	// The annotation message.
	// TODO(mattr): Allow richer annotations.
	Message string
}

func (Annotation) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vtrace.Annotation"`
}) {
}

func (x Annotation) VDLIsZero() bool {
	if !x.When.IsZero() {
		return false
	}
	if x.Message != "" {
		return false
	}
	return true
}

func (x Annotation) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	if !x.When.IsZero() {
		if err := enc.NextField("When"); err != nil {
			return err
		}
		var wire vdltime.Time
		if err := vdltime.TimeFromNative(&wire, x.When); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.Message != "" {
		if err := enc.NextFieldValueString("Message", vdl.StringType, x.Message); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Annotation) VDLRead(dec vdl.Decoder) error {
	*x = Annotation{}
	if err := dec.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "When":
			var wire vdltime.Time
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.TimeToNative(wire, &x.When); err != nil {
				return err
			}
		case "Message":
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Message = value
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

// A SpanRecord is the wire format for a Span.
type SpanRecord struct {
	Id     uniqueid.Id // The Id of the Span.
	Parent uniqueid.Id // The Id of this Span's parent.
	Name   string      // The Name of this span.
	Start  time.Time   // The start time of this span.
	End    time.Time   // The end time of this span.
	// A series of annotations.
	Annotations []Annotation
}

func (SpanRecord) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vtrace.SpanRecord"`
}) {
}

func (x SpanRecord) VDLIsZero() bool {
	if x.Id != (uniqueid.Id{}) {
		return false
	}
	if x.Parent != (uniqueid.Id{}) {
		return false
	}
	if x.Name != "" {
		return false
	}
	if !x.Start.IsZero() {
		return false
	}
	if !x.End.IsZero() {
		return false
	}
	if len(x.Annotations) != 0 {
		return false
	}
	return true
}

func (x SpanRecord) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_3); err != nil {
		return err
	}
	if x.Id != (uniqueid.Id{}) {
		if err := enc.NextFieldValueBytes("Id", __VDLType_array_4, x.Id[:]); err != nil {
			return err
		}
	}
	if x.Parent != (uniqueid.Id{}) {
		if err := enc.NextFieldValueBytes("Parent", __VDLType_array_4, x.Parent[:]); err != nil {
			return err
		}
	}
	if x.Name != "" {
		if err := enc.NextFieldValueString("Name", vdl.StringType, x.Name); err != nil {
			return err
		}
	}
	if !x.Start.IsZero() {
		if err := enc.NextField("Start"); err != nil {
			return err
		}
		var wire vdltime.Time
		if err := vdltime.TimeFromNative(&wire, x.Start); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if !x.End.IsZero() {
		if err := enc.NextField("End"); err != nil {
			return err
		}
		var wire vdltime.Time
		if err := vdltime.TimeFromNative(&wire, x.End); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if len(x.Annotations) != 0 {
		if err := enc.NextField("Annotations"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_1(enc, x.Annotations); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_1(enc vdl.Encoder, x []Annotation) error {
	if err := enc.StartValue(__VDLType_list_5); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := elem.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *SpanRecord) VDLRead(dec vdl.Decoder) error {
	*x = SpanRecord{}
	if err := dec.StartValue(__VDLType_struct_3); err != nil {
		return err
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Id":
			bytes := x.Id[:]
			if err := dec.ReadValueBytes(16, &bytes); err != nil {
				return err
			}
		case "Parent":
			bytes := x.Parent[:]
			if err := dec.ReadValueBytes(16, &bytes); err != nil {
				return err
			}
		case "Name":
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Name = value
			}
		case "Start":
			var wire vdltime.Time
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.TimeToNative(wire, &x.Start); err != nil {
				return err
			}
		case "End":
			var wire vdltime.Time
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.TimeToNative(wire, &x.End); err != nil {
				return err
			}
		case "Annotations":
			if err := __VDLReadAnon_list_1(dec, &x.Annotations); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_list_1(dec vdl.Decoder, x *[]Annotation) error {
	if err := dec.StartValue(__VDLType_list_5); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]Annotation, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			var elem Annotation
			if err := elem.VDLRead(dec); err != nil {
				return err
			}
			*x = append(*x, elem)
		}
	}
}

type TraceRecord struct {
	Id    uniqueid.Id
	Spans []SpanRecord
}

func (TraceRecord) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vtrace.TraceRecord"`
}) {
}

func (x TraceRecord) VDLIsZero() bool {
	if x.Id != (uniqueid.Id{}) {
		return false
	}
	if len(x.Spans) != 0 {
		return false
	}
	return true
}

func (x TraceRecord) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_6); err != nil {
		return err
	}
	if x.Id != (uniqueid.Id{}) {
		if err := enc.NextFieldValueBytes("Id", __VDLType_array_4, x.Id[:]); err != nil {
			return err
		}
	}
	if len(x.Spans) != 0 {
		if err := enc.NextField("Spans"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_2(enc, x.Spans); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_2(enc vdl.Encoder, x []SpanRecord) error {
	if err := enc.StartValue(__VDLType_list_7); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := elem.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *TraceRecord) VDLRead(dec vdl.Decoder) error {
	*x = TraceRecord{}
	if err := dec.StartValue(__VDLType_struct_6); err != nil {
		return err
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Id":
			bytes := x.Id[:]
			if err := dec.ReadValueBytes(16, &bytes); err != nil {
				return err
			}
		case "Spans":
			if err := __VDLReadAnon_list_2(dec, &x.Spans); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_list_2(dec vdl.Decoder, x *[]SpanRecord) error {
	if err := dec.StartValue(__VDLType_list_7); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]SpanRecord, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			var elem SpanRecord
			if err := elem.VDLRead(dec); err != nil {
				return err
			}
			*x = append(*x, elem)
		}
	}
}

type TraceFlags int32

func (TraceFlags) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vtrace.TraceFlags"`
}) {
}

func (x TraceFlags) VDLIsZero() bool {
	return x == 0
}

func (x TraceFlags) VDLWrite(enc vdl.Encoder) error {
	if err := enc.WriteValueInt(__VDLType_int32_8, int64(x)); err != nil {
		return err
	}
	return nil
}

func (x *TraceFlags) VDLRead(dec vdl.Decoder) error {
	switch value, err := dec.ReadValueInt(32); {
	case err != nil:
		return err
	default:
		*x = TraceFlags(value)
	}
	return nil
}

// Request is the object that carries trace informtion between processes.
type Request struct {
	SpanId   uniqueid.Id // The Id of the span that originated the RPC call.
	TraceId  uniqueid.Id // The Id of the trace this call is a part of.
	Flags    TraceFlags
	LogLevel int32
}

func (Request) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vtrace.Request"`
}) {
}

func (x Request) VDLIsZero() bool {
	return x == Request{}
}

func (x Request) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_9); err != nil {
		return err
	}
	if x.SpanId != (uniqueid.Id{}) {
		if err := enc.NextFieldValueBytes("SpanId", __VDLType_array_4, x.SpanId[:]); err != nil {
			return err
		}
	}
	if x.TraceId != (uniqueid.Id{}) {
		if err := enc.NextFieldValueBytes("TraceId", __VDLType_array_4, x.TraceId[:]); err != nil {
			return err
		}
	}
	if x.Flags != 0 {
		if err := enc.NextFieldValueInt("Flags", __VDLType_int32_8, int64(x.Flags)); err != nil {
			return err
		}
	}
	if x.LogLevel != 0 {
		if err := enc.NextFieldValueInt("LogLevel", vdl.Int32Type, int64(x.LogLevel)); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Request) VDLRead(dec vdl.Decoder) error {
	*x = Request{}
	if err := dec.StartValue(__VDLType_struct_9); err != nil {
		return err
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "SpanId":
			bytes := x.SpanId[:]
			if err := dec.ReadValueBytes(16, &bytes); err != nil {
				return err
			}
		case "TraceId":
			bytes := x.TraceId[:]
			if err := dec.ReadValueBytes(16, &bytes); err != nil {
				return err
			}
		case "Flags":
			switch value, err := dec.ReadValueInt(32); {
			case err != nil:
				return err
			default:
				x.Flags = TraceFlags(value)
			}
		case "LogLevel":
			switch value, err := dec.ReadValueInt(32); {
			case err != nil:
				return err
			default:
				x.LogLevel = int32(value)
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

type Response struct {
	// Flags give options for trace collection, the client should alter its
	// collection for this trace according to the flags sent back from the
	// server.
	Flags TraceFlags
	// Trace is collected trace data.  This may be empty.
	Trace TraceRecord
}

func (Response) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/vtrace.Response"`
}) {
}

func (x Response) VDLIsZero() bool {
	if x.Flags != 0 {
		return false
	}
	if !x.Trace.VDLIsZero() {
		return false
	}
	return true
}

func (x Response) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_10); err != nil {
		return err
	}
	if x.Flags != 0 {
		if err := enc.NextFieldValueInt("Flags", __VDLType_int32_8, int64(x.Flags)); err != nil {
			return err
		}
	}
	if !x.Trace.VDLIsZero() {
		if err := enc.NextField("Trace"); err != nil {
			return err
		}
		if err := x.Trace.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Response) VDLRead(dec vdl.Decoder) error {
	*x = Response{}
	if err := dec.StartValue(__VDLType_struct_10); err != nil {
		return err
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Flags":
			switch value, err := dec.ReadValueInt(32); {
			case err != nil:
				return err
			default:
				x.Flags = TraceFlags(value)
			}
		case "Trace":
			if err := x.Trace.VDLRead(dec); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

//////////////////////////////////////////////////
// Const definitions

const Empty = TraceFlags(0)
const CollectInMemory = TraceFlags(1)

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_struct_1  *vdl.Type
	__VDLType_struct_2  *vdl.Type
	__VDLType_struct_3  *vdl.Type
	__VDLType_array_4   *vdl.Type
	__VDLType_list_5    *vdl.Type
	__VDLType_struct_6  *vdl.Type
	__VDLType_list_7    *vdl.Type
	__VDLType_int32_8   *vdl.Type
	__VDLType_struct_9  *vdl.Type
	__VDLType_struct_10 *vdl.Type
)

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*Annotation)(nil))
	vdl.Register((*SpanRecord)(nil))
	vdl.Register((*TraceRecord)(nil))
	vdl.Register((*TraceFlags)(nil))
	vdl.Register((*Request)(nil))
	vdl.Register((*Response)(nil))

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*Annotation)(nil)).Elem()
	__VDLType_struct_2 = vdl.TypeOf((*vdltime.Time)(nil)).Elem()
	__VDLType_struct_3 = vdl.TypeOf((*SpanRecord)(nil)).Elem()
	__VDLType_array_4 = vdl.TypeOf((*uniqueid.Id)(nil))
	__VDLType_list_5 = vdl.TypeOf((*[]Annotation)(nil))
	__VDLType_struct_6 = vdl.TypeOf((*TraceRecord)(nil)).Elem()
	__VDLType_list_7 = vdl.TypeOf((*[]SpanRecord)(nil))
	__VDLType_int32_8 = vdl.TypeOf((*TraceFlags)(nil))
	__VDLType_struct_9 = vdl.TypeOf((*Request)(nil)).Elem()
	__VDLType_struct_10 = vdl.TypeOf((*Response)(nil)).Elem()

	return struct{}{}
}
