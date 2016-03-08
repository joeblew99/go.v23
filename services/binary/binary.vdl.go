// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: binary.vdl

// Package binary defines types for describing executable binaries.
package binary

import (
	"v.io/v23/vdl"
)

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

func (m *Description) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_binary_v_io_v23_services_binary_Description == nil || __VDLTypebinary0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Name")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Name), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Profiles")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		mapTarget6, err := fieldTarget5.StartMap(__VDLTypebinary1, len(m.Profiles))
		if err != nil {
			return err
		}
		for key8, value10 := range m.Profiles {
			keyTarget7, err := mapTarget6.StartKey()
			if err != nil {
				return err
			}
			if err := keyTarget7.FromString(string(key8), vdl.StringType); err != nil {
				return err
			}
			valueTarget9, err := mapTarget6.FinishKeyStartField(keyTarget7)
			if err != nil {
				return err
			}
			if err := valueTarget9.FromBool(bool(value10), vdl.BoolType); err != nil {
				return err
			}
			if err := mapTarget6.FinishField(keyTarget7, valueTarget9); err != nil {
				return err
			}
		}
		if err := fieldTarget5.FinishMap(mapTarget6); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Description) MakeVDLTarget() vdl.Target {
	return nil
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

func (m *PartInfo) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_binary_v_io_v23_services_binary_PartInfo == nil || __VDLTypebinary2 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Checksum")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Checksum), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Size")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromInt(int64(m.Size), vdl.Int64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *PartInfo) MakeVDLTarget() vdl.Target {
	return nil
}

func init() {
	vdl.Register((*Description)(nil))
	vdl.Register((*PartInfo)(nil))
}

var __VDLTypebinary0 *vdl.Type = vdl.TypeOf((*Description)(nil))
var __VDLTypebinary2 *vdl.Type = vdl.TypeOf((*PartInfo)(nil))
var __VDLTypebinary1 *vdl.Type = vdl.TypeOf(map[string]bool(nil))
var __VDLType_binary_v_io_v23_services_binary_Description *vdl.Type = vdl.TypeOf(Description{})
var __VDLType_binary_v_io_v23_services_binary_PartInfo *vdl.Type = vdl.TypeOf(PartInfo{})

func __VDLEnsureNativeBuilt_binary() {
}

const MissingChecksum = ""

const MissingSize = int64(-1)
