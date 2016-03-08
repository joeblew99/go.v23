// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: application.vdl

// Package application defines types for describing applications.
package application

import (
	"time"
	"v.io/v23/security"
	"v.io/v23/vdl"
	time_2 "v.io/v23/vdlroot/time"
)

// Envelope is a collection of metadata that describes an application.
type Envelope struct {
	// Title is the publisher-assigned application title.  Application
	// installations with the same title are considered as belonging to the
	// same application by the application management system.
	//
	// A change in the title signals a new application.
	Title string
	// Args is an array of command-line arguments to be used when executing
	// the binary.
	Args []string
	// Binary identifies the application binary.
	Binary SignedFile
	// Publisher represents the set of blessings that have been bound to
	// the principal who published this binary.
	Publisher security.Blessings
	// Env is an array that stores the environment variable values to be
	// used when executing the binary.
	Env []string
	// Packages is the set of packages to install on the local filesystem
	// before executing the binary
	Packages Packages
	// Restarts specifies how many times the device manager will attempt
	// to automatically restart an application that has crashed before
	// giving up and marking the application as NotRunning.
	Restarts int32
	// RestartTimeWindow is the time window within which an
	// application exit is considered a crash that counts against the
	// Restarts budget. If the application crashes after less than
	// RestartTimeWindow time for Restarts consecutive times, the
	// application is marked NotRunning and no more restart attempts
	// are made. If the application has run continuously for more
	// than RestartTimeWindow, subsequent crashes will again benefit
	// from up to Restarts restarts (that is, the Restarts budget is
	// reset by a successful run of at least RestartTimeWindow
	// duration).
	RestartTimeWindow time.Duration
}

func (Envelope) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/application.Envelope"`
}) {
}

func (m *Envelope) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	__VDLEnsureNativeBuilt_application()
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Title")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Title), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Args")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget6, err := fieldTarget5.StartList(__VDLTypeapplication1, len(m.Args))
		if err != nil {
			return err
		}
		for i, elem8 := range m.Args {
			elemTarget7, err := listTarget6.StartElem(i)
			if err != nil {
				return err
			}
			if err := elemTarget7.FromString(string(elem8), vdl.StringType); err != nil {
				return err
			}
			if err := listTarget6.FinishElem(elemTarget7); err != nil {
				return err
			}
		}
		if err := fieldTarget5.FinishList(listTarget6); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget9, fieldTarget10, err := fieldsTarget1.StartField("Binary")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Binary.FillVDLTarget(fieldTarget10, __VDLType_application_v_io_v23_services_application_SignedFile); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget9, fieldTarget10); err != nil {
			return err
		}
	}
	var wireValue11 security.WireBlessings
	if err := security.WireBlessingsFromNative(&wireValue11, m.Publisher); err != nil {
		return err
	}

	keyTarget12, fieldTarget13, err := fieldsTarget1.StartField("Publisher")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue11.FillVDLTarget(fieldTarget13, __VDLType_application_v_io_v23_security_WireBlessings); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget12, fieldTarget13); err != nil {
			return err
		}
	}
	keyTarget14, fieldTarget15, err := fieldsTarget1.StartField("Env")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget16, err := fieldTarget15.StartList(__VDLTypeapplication1, len(m.Env))
		if err != nil {
			return err
		}
		for i, elem18 := range m.Env {
			elemTarget17, err := listTarget16.StartElem(i)
			if err != nil {
				return err
			}
			if err := elemTarget17.FromString(string(elem18), vdl.StringType); err != nil {
				return err
			}
			if err := listTarget16.FinishElem(elemTarget17); err != nil {
				return err
			}
		}
		if err := fieldTarget15.FinishList(listTarget16); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget14, fieldTarget15); err != nil {
			return err
		}
	}
	keyTarget19, fieldTarget20, err := fieldsTarget1.StartField("Packages")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Packages.FillVDLTarget(fieldTarget20, __VDLType_application_v_io_v23_services_application_Packages); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget19, fieldTarget20); err != nil {
			return err
		}
	}
	keyTarget21, fieldTarget22, err := fieldsTarget1.StartField("Restarts")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget22.FromInt(int64(m.Restarts), vdl.Int32Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget21, fieldTarget22); err != nil {
			return err
		}
	}
	var wireValue23 time_2.Duration
	if err := time_2.DurationFromNative(&wireValue23, m.RestartTimeWindow); err != nil {
		return err
	}

	keyTarget24, fieldTarget25, err := fieldsTarget1.StartField("RestartTimeWindow")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue23.FillVDLTarget(fieldTarget25, __VDLType_application_time_Duration); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget24, fieldTarget25); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *Envelope) MakeVDLTarget() vdl.Target {
	return nil
}

// Packages represents a set of packages. The map key is the local
// file/directory name, relative to the instance's packages directory, where the
// package should be installed. For archives, this name represents a directory
// into which the archive is to be extracted, and for regular files it
// represents the name for the file.  The map value is the package
// specification.
//
// Each object's media type determines how to install it.
//
// For example, with key=pkg1,value=SignedFile{File:binaryrepo/configfiles} (an
// archive), the "configfiles" package will be installed under the "pkg1"
// directory. With key=pkg2,value=SignedFile{File:binaryrepo/binfile} (a
// binary), the "binfile" file will be installed as the "pkg2" file.
//
// The keys must be valid file/directory names, without path separators.
//
// Any number of packages may be specified.
type Packages map[string]SignedFile

func (Packages) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/application.Packages"`
}) {
}

func (m Packages) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	mapTarget1, err := t.StartMap(__VDLType_application_v_io_v23_services_application_Packages, len(m))
	if err != nil {
		return err
	}
	for key3, value5 := range m {
		keyTarget2, err := mapTarget1.StartKey()
		if err != nil {
			return err
		}
		if err := keyTarget2.FromString(string(key3), vdl.StringType); err != nil {
			return err
		}
		valueTarget4, err := mapTarget1.FinishKeyStartField(keyTarget2)
		if err != nil {
			return err
		}

		if err := value5.FillVDLTarget(valueTarget4, __VDLType_application_v_io_v23_services_application_SignedFile); err != nil {
			return err
		}
		if err := mapTarget1.FinishField(keyTarget2, valueTarget4); err != nil {
			return err
		}
	}
	if err := t.FinishMap(mapTarget1); err != nil {
		return err
	}
	return nil
}

func (m Packages) MakeVDLTarget() vdl.Target {
	return nil
}

// SignedFile represents a file accompanied by a signature of its contents.
type SignedFile struct {
	//  File is the object name of the file.
	File string
	// Signature represents a signature on the sha256 hash of the file
	// contents by the publisher principal.
	Signature security.Signature
}

func (SignedFile) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/application.SignedFile"`
}) {
}

func (m *SignedFile) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	if __VDLType_application_v_io_v23_services_application_SignedFile == nil || __VDLTypeapplication2 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("File")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.File), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Signature")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Signature.FillVDLTarget(fieldTarget5, __VDLType_application_v_io_v23_security_Signature); err != nil {
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

func (m *SignedFile) MakeVDLTarget() vdl.Target {
	return nil
}

func init() {
	vdl.Register((*Envelope)(nil))
	vdl.Register((*Packages)(nil))
	vdl.Register((*SignedFile)(nil))
}

var __VDLTypeapplication0 *vdl.Type

func __VDLTypeapplication0_gen() *vdl.Type {
	__VDLTypeapplication0Builder := vdl.TypeBuilder{}

	__VDLTypeapplication01 := __VDLTypeapplication0Builder.Optional()
	__VDLTypeapplication02 := __VDLTypeapplication0Builder.Struct()
	__VDLTypeapplication03 := __VDLTypeapplication0Builder.Named("v.io/v23/services/application.Envelope").AssignBase(__VDLTypeapplication02)
	__VDLTypeapplication04 := vdl.StringType
	__VDLTypeapplication02.AppendField("Title", __VDLTypeapplication04)
	__VDLTypeapplication05 := __VDLTypeapplication0Builder.List()
	__VDLTypeapplication05.AssignElem(__VDLTypeapplication04)
	__VDLTypeapplication02.AppendField("Args", __VDLTypeapplication05)
	__VDLTypeapplication06 := __VDLTypeapplication0Builder.Struct()
	__VDLTypeapplication07 := __VDLTypeapplication0Builder.Named("v.io/v23/services/application.SignedFile").AssignBase(__VDLTypeapplication06)
	__VDLTypeapplication06.AppendField("File", __VDLTypeapplication04)
	__VDLTypeapplication08 := __VDLTypeapplication0Builder.Struct()
	__VDLTypeapplication09 := __VDLTypeapplication0Builder.Named("v.io/v23/security.Signature").AssignBase(__VDLTypeapplication08)
	__VDLTypeapplication010 := __VDLTypeapplication0Builder.List()
	__VDLTypeapplication011 := vdl.ByteType
	__VDLTypeapplication010.AssignElem(__VDLTypeapplication011)
	__VDLTypeapplication08.AppendField("Purpose", __VDLTypeapplication010)
	__VDLTypeapplication012 := vdl.StringType
	__VDLTypeapplication013 := __VDLTypeapplication0Builder.Named("v.io/v23/security.Hash").AssignBase(__VDLTypeapplication012)
	__VDLTypeapplication08.AppendField("Hash", __VDLTypeapplication013)
	__VDLTypeapplication08.AppendField("R", __VDLTypeapplication010)
	__VDLTypeapplication08.AppendField("S", __VDLTypeapplication010)
	__VDLTypeapplication06.AppendField("Signature", __VDLTypeapplication09)
	__VDLTypeapplication02.AppendField("Binary", __VDLTypeapplication07)
	__VDLTypeapplication014 := __VDLTypeapplication0Builder.Struct()
	__VDLTypeapplication015 := __VDLTypeapplication0Builder.Named("v.io/v23/security.WireBlessings").AssignBase(__VDLTypeapplication014)
	__VDLTypeapplication016 := __VDLTypeapplication0Builder.List()
	__VDLTypeapplication017 := __VDLTypeapplication0Builder.List()
	__VDLTypeapplication018 := __VDLTypeapplication0Builder.Struct()
	__VDLTypeapplication019 := __VDLTypeapplication0Builder.Named("v.io/v23/security.Certificate").AssignBase(__VDLTypeapplication018)
	__VDLTypeapplication018.AppendField("Extension", __VDLTypeapplication04)
	__VDLTypeapplication018.AppendField("PublicKey", __VDLTypeapplication010)
	__VDLTypeapplication020 := __VDLTypeapplication0Builder.List()
	__VDLTypeapplication021 := __VDLTypeapplication0Builder.Struct()
	__VDLTypeapplication022 := __VDLTypeapplication0Builder.Named("v.io/v23/security.Caveat").AssignBase(__VDLTypeapplication021)
	__VDLTypeapplication023 := __VDLTypeapplication0Builder.Array()
	__VDLTypeapplication024 := __VDLTypeapplication0Builder.Named("v.io/v23/uniqueid.Id").AssignBase(__VDLTypeapplication023)
	__VDLTypeapplication023.AssignElem(__VDLTypeapplication011)
	__VDLTypeapplication023.AssignLen(16)
	__VDLTypeapplication021.AppendField("Id", __VDLTypeapplication024)
	__VDLTypeapplication021.AppendField("ParamVom", __VDLTypeapplication010)
	__VDLTypeapplication020.AssignElem(__VDLTypeapplication022)
	__VDLTypeapplication018.AppendField("Caveats", __VDLTypeapplication020)
	__VDLTypeapplication018.AppendField("Signature", __VDLTypeapplication09)
	__VDLTypeapplication017.AssignElem(__VDLTypeapplication019)
	__VDLTypeapplication016.AssignElem(__VDLTypeapplication017)
	__VDLTypeapplication014.AppendField("CertificateChains", __VDLTypeapplication016)
	__VDLTypeapplication02.AppendField("Publisher", __VDLTypeapplication015)
	__VDLTypeapplication02.AppendField("Env", __VDLTypeapplication05)
	__VDLTypeapplication025 := __VDLTypeapplication0Builder.Map()
	__VDLTypeapplication026 := __VDLTypeapplication0Builder.Named("v.io/v23/services/application.Packages").AssignBase(__VDLTypeapplication025)
	__VDLTypeapplication025.AssignKey(__VDLTypeapplication04)
	__VDLTypeapplication025.AssignElem(__VDLTypeapplication07)
	__VDLTypeapplication02.AppendField("Packages", __VDLTypeapplication026)
	__VDLTypeapplication027 := vdl.Int32Type
	__VDLTypeapplication02.AppendField("Restarts", __VDLTypeapplication027)
	__VDLTypeapplication028 := __VDLTypeapplication0Builder.Struct()
	__VDLTypeapplication029 := __VDLTypeapplication0Builder.Named("time.Duration").AssignBase(__VDLTypeapplication028)
	__VDLTypeapplication030 := vdl.Int64Type
	__VDLTypeapplication028.AppendField("Seconds", __VDLTypeapplication030)
	__VDLTypeapplication028.AppendField("Nanos", __VDLTypeapplication027)
	__VDLTypeapplication02.AppendField("RestartTimeWindow", __VDLTypeapplication029)
	__VDLTypeapplication01.AssignElem(__VDLTypeapplication03)
	__VDLTypeapplication0Builder.Build()
	__VDLTypeapplication0v, err := __VDLTypeapplication01.Built()
	if err != nil {
		panic(err)
	}
	return __VDLTypeapplication0v
}
func init() {
	__VDLTypeapplication0 = __VDLTypeapplication0_gen()
}

var __VDLTypeapplication2 *vdl.Type = vdl.TypeOf((*SignedFile)(nil))
var __VDLTypeapplication1 *vdl.Type = vdl.TypeOf([]string(nil))
var __VDLType_application_time_Duration *vdl.Type

func __VDLType_application_time_Duration_gen() *vdl.Type {
	__VDLType_application_time_DurationBuilder := vdl.TypeBuilder{}

	__VDLType_application_time_Duration1 := __VDLType_application_time_DurationBuilder.Struct()
	__VDLType_application_time_Duration2 := __VDLType_application_time_DurationBuilder.Named("time.Duration").AssignBase(__VDLType_application_time_Duration1)
	__VDLType_application_time_Duration3 := vdl.Int64Type
	__VDLType_application_time_Duration1.AppendField("Seconds", __VDLType_application_time_Duration3)
	__VDLType_application_time_Duration4 := vdl.Int32Type
	__VDLType_application_time_Duration1.AppendField("Nanos", __VDLType_application_time_Duration4)
	__VDLType_application_time_DurationBuilder.Build()
	__VDLType_application_time_Durationv, err := __VDLType_application_time_Duration2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_application_time_Durationv
}
func init() {
	__VDLType_application_time_Duration = __VDLType_application_time_Duration_gen()
}

var __VDLType_application_v_io_v23_security_Signature *vdl.Type = vdl.TypeOf(security.Signature{})
var __VDLType_application_v_io_v23_security_WireBlessings *vdl.Type

func __VDLType_application_v_io_v23_security_WireBlessings_gen() *vdl.Type {
	__VDLType_application_v_io_v23_security_WireBlessingsBuilder := vdl.TypeBuilder{}

	__VDLType_application_v_io_v23_security_WireBlessings1 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.Struct()
	__VDLType_application_v_io_v23_security_WireBlessings2 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.Named("v.io/v23/security.WireBlessings").AssignBase(__VDLType_application_v_io_v23_security_WireBlessings1)
	__VDLType_application_v_io_v23_security_WireBlessings3 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.List()
	__VDLType_application_v_io_v23_security_WireBlessings4 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.List()
	__VDLType_application_v_io_v23_security_WireBlessings5 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.Struct()
	__VDLType_application_v_io_v23_security_WireBlessings6 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.Named("v.io/v23/security.Certificate").AssignBase(__VDLType_application_v_io_v23_security_WireBlessings5)
	__VDLType_application_v_io_v23_security_WireBlessings7 := vdl.StringType
	__VDLType_application_v_io_v23_security_WireBlessings5.AppendField("Extension", __VDLType_application_v_io_v23_security_WireBlessings7)
	__VDLType_application_v_io_v23_security_WireBlessings8 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.List()
	__VDLType_application_v_io_v23_security_WireBlessings9 := vdl.ByteType
	__VDLType_application_v_io_v23_security_WireBlessings8.AssignElem(__VDLType_application_v_io_v23_security_WireBlessings9)
	__VDLType_application_v_io_v23_security_WireBlessings5.AppendField("PublicKey", __VDLType_application_v_io_v23_security_WireBlessings8)
	__VDLType_application_v_io_v23_security_WireBlessings10 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.List()
	__VDLType_application_v_io_v23_security_WireBlessings11 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.Struct()
	__VDLType_application_v_io_v23_security_WireBlessings12 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.Named("v.io/v23/security.Caveat").AssignBase(__VDLType_application_v_io_v23_security_WireBlessings11)
	__VDLType_application_v_io_v23_security_WireBlessings13 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.Array()
	__VDLType_application_v_io_v23_security_WireBlessings14 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.Named("v.io/v23/uniqueid.Id").AssignBase(__VDLType_application_v_io_v23_security_WireBlessings13)
	__VDLType_application_v_io_v23_security_WireBlessings13.AssignElem(__VDLType_application_v_io_v23_security_WireBlessings9)
	__VDLType_application_v_io_v23_security_WireBlessings13.AssignLen(16)
	__VDLType_application_v_io_v23_security_WireBlessings11.AppendField("Id", __VDLType_application_v_io_v23_security_WireBlessings14)
	__VDLType_application_v_io_v23_security_WireBlessings11.AppendField("ParamVom", __VDLType_application_v_io_v23_security_WireBlessings8)
	__VDLType_application_v_io_v23_security_WireBlessings10.AssignElem(__VDLType_application_v_io_v23_security_WireBlessings12)
	__VDLType_application_v_io_v23_security_WireBlessings5.AppendField("Caveats", __VDLType_application_v_io_v23_security_WireBlessings10)
	__VDLType_application_v_io_v23_security_WireBlessings15 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.Struct()
	__VDLType_application_v_io_v23_security_WireBlessings16 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.Named("v.io/v23/security.Signature").AssignBase(__VDLType_application_v_io_v23_security_WireBlessings15)
	__VDLType_application_v_io_v23_security_WireBlessings15.AppendField("Purpose", __VDLType_application_v_io_v23_security_WireBlessings8)
	__VDLType_application_v_io_v23_security_WireBlessings17 := vdl.StringType
	__VDLType_application_v_io_v23_security_WireBlessings18 := __VDLType_application_v_io_v23_security_WireBlessingsBuilder.Named("v.io/v23/security.Hash").AssignBase(__VDLType_application_v_io_v23_security_WireBlessings17)
	__VDLType_application_v_io_v23_security_WireBlessings15.AppendField("Hash", __VDLType_application_v_io_v23_security_WireBlessings18)
	__VDLType_application_v_io_v23_security_WireBlessings15.AppendField("R", __VDLType_application_v_io_v23_security_WireBlessings8)
	__VDLType_application_v_io_v23_security_WireBlessings15.AppendField("S", __VDLType_application_v_io_v23_security_WireBlessings8)
	__VDLType_application_v_io_v23_security_WireBlessings5.AppendField("Signature", __VDLType_application_v_io_v23_security_WireBlessings16)
	__VDLType_application_v_io_v23_security_WireBlessings4.AssignElem(__VDLType_application_v_io_v23_security_WireBlessings6)
	__VDLType_application_v_io_v23_security_WireBlessings3.AssignElem(__VDLType_application_v_io_v23_security_WireBlessings4)
	__VDLType_application_v_io_v23_security_WireBlessings1.AppendField("CertificateChains", __VDLType_application_v_io_v23_security_WireBlessings3)
	__VDLType_application_v_io_v23_security_WireBlessingsBuilder.Build()
	__VDLType_application_v_io_v23_security_WireBlessingsv, err := __VDLType_application_v_io_v23_security_WireBlessings2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_application_v_io_v23_security_WireBlessingsv
}
func init() {
	__VDLType_application_v_io_v23_security_WireBlessings = __VDLType_application_v_io_v23_security_WireBlessings_gen()
}

var __VDLType_application_v_io_v23_services_application_Envelope *vdl.Type

func __VDLType_application_v_io_v23_services_application_Envelope_gen() *vdl.Type {
	__VDLType_application_v_io_v23_services_application_EnvelopeBuilder := vdl.TypeBuilder{}

	__VDLType_application_v_io_v23_services_application_Envelope1 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Struct()
	__VDLType_application_v_io_v23_services_application_Envelope2 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Named("v.io/v23/services/application.Envelope").AssignBase(__VDLType_application_v_io_v23_services_application_Envelope1)
	__VDLType_application_v_io_v23_services_application_Envelope3 := vdl.StringType
	__VDLType_application_v_io_v23_services_application_Envelope1.AppendField("Title", __VDLType_application_v_io_v23_services_application_Envelope3)
	__VDLType_application_v_io_v23_services_application_Envelope4 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.List()
	__VDLType_application_v_io_v23_services_application_Envelope4.AssignElem(__VDLType_application_v_io_v23_services_application_Envelope3)
	__VDLType_application_v_io_v23_services_application_Envelope1.AppendField("Args", __VDLType_application_v_io_v23_services_application_Envelope4)
	__VDLType_application_v_io_v23_services_application_Envelope5 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Struct()
	__VDLType_application_v_io_v23_services_application_Envelope6 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Named("v.io/v23/services/application.SignedFile").AssignBase(__VDLType_application_v_io_v23_services_application_Envelope5)
	__VDLType_application_v_io_v23_services_application_Envelope5.AppendField("File", __VDLType_application_v_io_v23_services_application_Envelope3)
	__VDLType_application_v_io_v23_services_application_Envelope7 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Struct()
	__VDLType_application_v_io_v23_services_application_Envelope8 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Named("v.io/v23/security.Signature").AssignBase(__VDLType_application_v_io_v23_services_application_Envelope7)
	__VDLType_application_v_io_v23_services_application_Envelope9 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.List()
	__VDLType_application_v_io_v23_services_application_Envelope10 := vdl.ByteType
	__VDLType_application_v_io_v23_services_application_Envelope9.AssignElem(__VDLType_application_v_io_v23_services_application_Envelope10)
	__VDLType_application_v_io_v23_services_application_Envelope7.AppendField("Purpose", __VDLType_application_v_io_v23_services_application_Envelope9)
	__VDLType_application_v_io_v23_services_application_Envelope11 := vdl.StringType
	__VDLType_application_v_io_v23_services_application_Envelope12 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Named("v.io/v23/security.Hash").AssignBase(__VDLType_application_v_io_v23_services_application_Envelope11)
	__VDLType_application_v_io_v23_services_application_Envelope7.AppendField("Hash", __VDLType_application_v_io_v23_services_application_Envelope12)
	__VDLType_application_v_io_v23_services_application_Envelope7.AppendField("R", __VDLType_application_v_io_v23_services_application_Envelope9)
	__VDLType_application_v_io_v23_services_application_Envelope7.AppendField("S", __VDLType_application_v_io_v23_services_application_Envelope9)
	__VDLType_application_v_io_v23_services_application_Envelope5.AppendField("Signature", __VDLType_application_v_io_v23_services_application_Envelope8)
	__VDLType_application_v_io_v23_services_application_Envelope1.AppendField("Binary", __VDLType_application_v_io_v23_services_application_Envelope6)
	__VDLType_application_v_io_v23_services_application_Envelope13 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Struct()
	__VDLType_application_v_io_v23_services_application_Envelope14 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Named("v.io/v23/security.WireBlessings").AssignBase(__VDLType_application_v_io_v23_services_application_Envelope13)
	__VDLType_application_v_io_v23_services_application_Envelope15 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.List()
	__VDLType_application_v_io_v23_services_application_Envelope16 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.List()
	__VDLType_application_v_io_v23_services_application_Envelope17 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Struct()
	__VDLType_application_v_io_v23_services_application_Envelope18 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Named("v.io/v23/security.Certificate").AssignBase(__VDLType_application_v_io_v23_services_application_Envelope17)
	__VDLType_application_v_io_v23_services_application_Envelope17.AppendField("Extension", __VDLType_application_v_io_v23_services_application_Envelope3)
	__VDLType_application_v_io_v23_services_application_Envelope17.AppendField("PublicKey", __VDLType_application_v_io_v23_services_application_Envelope9)
	__VDLType_application_v_io_v23_services_application_Envelope19 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.List()
	__VDLType_application_v_io_v23_services_application_Envelope20 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Struct()
	__VDLType_application_v_io_v23_services_application_Envelope21 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Named("v.io/v23/security.Caveat").AssignBase(__VDLType_application_v_io_v23_services_application_Envelope20)
	__VDLType_application_v_io_v23_services_application_Envelope22 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Array()
	__VDLType_application_v_io_v23_services_application_Envelope23 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Named("v.io/v23/uniqueid.Id").AssignBase(__VDLType_application_v_io_v23_services_application_Envelope22)
	__VDLType_application_v_io_v23_services_application_Envelope22.AssignElem(__VDLType_application_v_io_v23_services_application_Envelope10)
	__VDLType_application_v_io_v23_services_application_Envelope22.AssignLen(16)
	__VDLType_application_v_io_v23_services_application_Envelope20.AppendField("Id", __VDLType_application_v_io_v23_services_application_Envelope23)
	__VDLType_application_v_io_v23_services_application_Envelope20.AppendField("ParamVom", __VDLType_application_v_io_v23_services_application_Envelope9)
	__VDLType_application_v_io_v23_services_application_Envelope19.AssignElem(__VDLType_application_v_io_v23_services_application_Envelope21)
	__VDLType_application_v_io_v23_services_application_Envelope17.AppendField("Caveats", __VDLType_application_v_io_v23_services_application_Envelope19)
	__VDLType_application_v_io_v23_services_application_Envelope17.AppendField("Signature", __VDLType_application_v_io_v23_services_application_Envelope8)
	__VDLType_application_v_io_v23_services_application_Envelope16.AssignElem(__VDLType_application_v_io_v23_services_application_Envelope18)
	__VDLType_application_v_io_v23_services_application_Envelope15.AssignElem(__VDLType_application_v_io_v23_services_application_Envelope16)
	__VDLType_application_v_io_v23_services_application_Envelope13.AppendField("CertificateChains", __VDLType_application_v_io_v23_services_application_Envelope15)
	__VDLType_application_v_io_v23_services_application_Envelope1.AppendField("Publisher", __VDLType_application_v_io_v23_services_application_Envelope14)
	__VDLType_application_v_io_v23_services_application_Envelope1.AppendField("Env", __VDLType_application_v_io_v23_services_application_Envelope4)
	__VDLType_application_v_io_v23_services_application_Envelope24 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Map()
	__VDLType_application_v_io_v23_services_application_Envelope25 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Named("v.io/v23/services/application.Packages").AssignBase(__VDLType_application_v_io_v23_services_application_Envelope24)
	__VDLType_application_v_io_v23_services_application_Envelope24.AssignKey(__VDLType_application_v_io_v23_services_application_Envelope3)
	__VDLType_application_v_io_v23_services_application_Envelope24.AssignElem(__VDLType_application_v_io_v23_services_application_Envelope6)
	__VDLType_application_v_io_v23_services_application_Envelope1.AppendField("Packages", __VDLType_application_v_io_v23_services_application_Envelope25)
	__VDLType_application_v_io_v23_services_application_Envelope26 := vdl.Int32Type
	__VDLType_application_v_io_v23_services_application_Envelope1.AppendField("Restarts", __VDLType_application_v_io_v23_services_application_Envelope26)
	__VDLType_application_v_io_v23_services_application_Envelope27 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Struct()
	__VDLType_application_v_io_v23_services_application_Envelope28 := __VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Named("time.Duration").AssignBase(__VDLType_application_v_io_v23_services_application_Envelope27)
	__VDLType_application_v_io_v23_services_application_Envelope29 := vdl.Int64Type
	__VDLType_application_v_io_v23_services_application_Envelope27.AppendField("Seconds", __VDLType_application_v_io_v23_services_application_Envelope29)
	__VDLType_application_v_io_v23_services_application_Envelope27.AppendField("Nanos", __VDLType_application_v_io_v23_services_application_Envelope26)
	__VDLType_application_v_io_v23_services_application_Envelope1.AppendField("RestartTimeWindow", __VDLType_application_v_io_v23_services_application_Envelope28)
	__VDLType_application_v_io_v23_services_application_EnvelopeBuilder.Build()
	__VDLType_application_v_io_v23_services_application_Envelopev, err := __VDLType_application_v_io_v23_services_application_Envelope2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_application_v_io_v23_services_application_Envelopev
}
func init() {
	__VDLType_application_v_io_v23_services_application_Envelope = __VDLType_application_v_io_v23_services_application_Envelope_gen()
}

var __VDLType_application_v_io_v23_services_application_Packages *vdl.Type = vdl.TypeOf(Packages(nil))
var __VDLType_application_v_io_v23_services_application_SignedFile *vdl.Type = vdl.TypeOf(SignedFile{})

func __VDLEnsureNativeBuilt_application() {
	if __VDLTypeapplication0 == nil {
		__VDLTypeapplication0 = __VDLTypeapplication0_gen()
	}
	if __VDLType_application_time_Duration == nil {
		__VDLType_application_time_Duration = __VDLType_application_time_Duration_gen()
	}
	if __VDLType_application_v_io_v23_security_WireBlessings == nil {
		__VDLType_application_v_io_v23_security_WireBlessings = __VDLType_application_v_io_v23_security_WireBlessings_gen()
	}
	if __VDLType_application_v_io_v23_services_application_Envelope == nil {
		__VDLType_application_v_io_v23_services_application_Envelope = __VDLType_application_v_io_v23_services_application_Envelope_gen()
	}
}

// Device manager application envelopes must present this title.
const DeviceManagerTitle = "device manager"
