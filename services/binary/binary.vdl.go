// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: binary

// Package binary defines types for describing executable binaries.
package binary

import (
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// Description describes a binary. Binaries are named and have been
// determined to run on some set of profiles. The mechanism for
// determing profiles is specifically not specified and left to the
// implementation of the interface that generates the description.
type Description struct {
	// Name is the Object name of the application binary that can
	// be used to fetch the actual binary from a content server.
	Name string
	// Profiles is a set of names of compatible profiles.  Each
	// name can either be an Object name that resolves to a
	// Profile, or can be the profile's label, e.g.:
	//
	//   "profiles/google/cluster/diskfull"
	//   "linux-media"
	//
	// Application developers can specify compatible profiles by
	// hand, but we also want to be able to automatically derive
	// the matching profiles from examining the binary itself
	// (e.g. that's what Build.Describe() does).
	Profiles map[string]bool
}

func (Description) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/binary.Description"`
}) {
}

func (x Description) VDLIsZero() bool {
	if x.Name != "" {
		return false
	}
	if len(x.Profiles) != 0 {
		return false
	}
	return true
}

func (x Description) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	if x.Name != "" {
		if err := enc.NextFieldValueString("Name", vdl.StringType, x.Name); err != nil {
			return err
		}
	}
	if len(x.Profiles) != 0 {
		if err := enc.NextField("Profiles"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_map_1(enc, x.Profiles); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_map_1(enc vdl.Encoder, x map[string]bool) error {
	if err := enc.StartValue(__VDLType_map_2); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key, elem := range x {
		if err := enc.NextEntryValueString(vdl.StringType, key); err != nil {
			return err
		}
		if err := enc.WriteValueBool(vdl.BoolType, elem); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Description) VDLRead(dec vdl.Decoder) error {
	*x = Description{}
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
		case "Name":
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Name = value
			}
		case "Profiles":
			if err := __VDLReadAnon_map_1(dec, &x.Profiles); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_map_1(dec vdl.Decoder, x *map[string]bool) error {
	if err := dec.StartValue(__VDLType_map_2); err != nil {
		return err
	}
	var tmpMap map[string]bool
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(map[string]bool, len)
	}
	for {
		switch done, key, err := dec.NextEntryValueString(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		default:
			var elem bool
			switch value, err := dec.ReadValueBool(); {
			case err != nil:
				return err
			default:
				elem = value
			}
			if tmpMap == nil {
				tmpMap = make(map[string]bool)
			}
			tmpMap[key] = elem
		}
	}
}

// PartInfo holds information describing a binary part.
type PartInfo struct {
	// Checksum holds the hex-encoded MD5 checksum of the binary part.
	Checksum string
	// Size holds the binary part size in bytes.
	Size int64
}

func (PartInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/binary.PartInfo"`
}) {
}

func (x PartInfo) VDLIsZero() bool {
	return x == PartInfo{}
}

func (x PartInfo) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_3); err != nil {
		return err
	}
	if x.Checksum != "" {
		if err := enc.NextFieldValueString("Checksum", vdl.StringType, x.Checksum); err != nil {
			return err
		}
	}
	if x.Size != 0 {
		if err := enc.NextFieldValueInt("Size", vdl.Int64Type, x.Size); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *PartInfo) VDLRead(dec vdl.Decoder) error {
	*x = PartInfo{}
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
		case "Checksum":
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Checksum = value
			}
		case "Size":
			switch value, err := dec.ReadValueInt(64); {
			case err != nil:
				return err
			default:
				x.Size = value
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

const MissingChecksum = ""
const MissingSize = int64(-1)

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_struct_1 *vdl.Type
	__VDLType_map_2    *vdl.Type
	__VDLType_struct_3 *vdl.Type
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
	vdl.Register((*Description)(nil))
	vdl.Register((*PartInfo)(nil))

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*Description)(nil)).Elem()
	__VDLType_map_2 = vdl.TypeOf((*map[string]bool)(nil))
	__VDLType_struct_3 = vdl.TypeOf((*PartInfo)(nil)).Elem()

	return struct{}{}
}
