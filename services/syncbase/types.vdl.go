// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package syncbase

import (
	"time"
	"v.io/v23/vdl"
	time_2 "v.io/v23/vdlroot/time"
)

// DevModeUpdateVClockOpts specifies what DevModeUpdateVClock should do, as
// described below.
type DevModeUpdateVClockOpts struct {
	// If specified, sets the NTP host to talk to for subsequent NTP requests.
	NtpHost string
	// If Now is specified, the fake system clock is updated to the given values
	// of Now and ElapsedTime. If Now is not specified (i.e. takes the zero
	// value), the system clock is not touched by DevModeUpdateVClock.
	Now         time.Time
	ElapsedTime time.Duration
	// If specified, the clock daemon's local and/or NTP update code is triggered
	// after applying the updates specified by the fields above. (Helpful because
	// otherwise these only run periodically.) These functions work even if the
	// clock daemon hasn't been started.
	DoNtpUpdate   bool
	DoLocalUpdate bool
}

func (DevModeUpdateVClockOpts) __VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/syncbase.DevModeUpdateVClockOpts"`
}) {
}

func (m *DevModeUpdateVClockOpts) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {

	__VDLEnsureNativeBuilt_types()
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("NtpHost")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.NtpHost), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	var wireValue4 time_2.Time
	if err := time_2.TimeFromNative(&wireValue4, m.Now); err != nil {
		return err
	}

	keyTarget5, fieldTarget6, err := fieldsTarget1.StartField("Now")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue4.FillVDLTarget(fieldTarget6, __VDLType_types_time_Time); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget5, fieldTarget6); err != nil {
			return err
		}
	}
	var wireValue7 time_2.Duration
	if err := time_2.DurationFromNative(&wireValue7, m.ElapsedTime); err != nil {
		return err
	}

	keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("ElapsedTime")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue7.FillVDLTarget(fieldTarget9, __VDLType_types_time_Duration); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
			return err
		}
	}
	keyTarget10, fieldTarget11, err := fieldsTarget1.StartField("DoNtpUpdate")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget11.FromBool(bool(m.DoNtpUpdate), vdl.BoolType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget10, fieldTarget11); err != nil {
			return err
		}
	}
	keyTarget12, fieldTarget13, err := fieldsTarget1.StartField("DoLocalUpdate")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget13.FromBool(bool(m.DoLocalUpdate), vdl.BoolType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget12, fieldTarget13); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *DevModeUpdateVClockOpts) MakeVDLTarget() vdl.Target {
	return nil
}

func init() {
	vdl.Register((*DevModeUpdateVClockOpts)(nil))
}

var __VDLTypetypes0 *vdl.Type

func __VDLTypetypes0_gen() *vdl.Type {
	__VDLTypetypes0Builder := vdl.TypeBuilder{}

	__VDLTypetypes01 := __VDLTypetypes0Builder.Optional()
	__VDLTypetypes02 := __VDLTypetypes0Builder.Struct()
	__VDLTypetypes03 := __VDLTypetypes0Builder.Named("v.io/v23/services/syncbase.DevModeUpdateVClockOpts").AssignBase(__VDLTypetypes02)
	__VDLTypetypes04 := vdl.StringType
	__VDLTypetypes02.AppendField("NtpHost", __VDLTypetypes04)
	__VDLTypetypes05 := __VDLTypetypes0Builder.Struct()
	__VDLTypetypes06 := __VDLTypetypes0Builder.Named("time.Time").AssignBase(__VDLTypetypes05)
	__VDLTypetypes07 := vdl.Int64Type
	__VDLTypetypes05.AppendField("Seconds", __VDLTypetypes07)
	__VDLTypetypes08 := vdl.Int32Type
	__VDLTypetypes05.AppendField("Nanos", __VDLTypetypes08)
	__VDLTypetypes02.AppendField("Now", __VDLTypetypes06)
	__VDLTypetypes09 := __VDLTypetypes0Builder.Struct()
	__VDLTypetypes010 := __VDLTypetypes0Builder.Named("time.Duration").AssignBase(__VDLTypetypes09)
	__VDLTypetypes09.AppendField("Seconds", __VDLTypetypes07)
	__VDLTypetypes09.AppendField("Nanos", __VDLTypetypes08)
	__VDLTypetypes02.AppendField("ElapsedTime", __VDLTypetypes010)
	__VDLTypetypes011 := vdl.BoolType
	__VDLTypetypes02.AppendField("DoNtpUpdate", __VDLTypetypes011)
	__VDLTypetypes02.AppendField("DoLocalUpdate", __VDLTypetypes011)
	__VDLTypetypes01.AssignElem(__VDLTypetypes03)
	__VDLTypetypes0Builder.Build()
	__VDLTypetypes0v, err := __VDLTypetypes01.Built()
	if err != nil {
		panic(err)
	}
	return __VDLTypetypes0v
}
func init() {
	__VDLTypetypes0 = __VDLTypetypes0_gen()
}

var __VDLType_types_time_Duration *vdl.Type

func __VDLType_types_time_Duration_gen() *vdl.Type {
	__VDLType_types_time_DurationBuilder := vdl.TypeBuilder{}

	__VDLType_types_time_Duration1 := __VDLType_types_time_DurationBuilder.Struct()
	__VDLType_types_time_Duration2 := __VDLType_types_time_DurationBuilder.Named("time.Duration").AssignBase(__VDLType_types_time_Duration1)
	__VDLType_types_time_Duration3 := vdl.Int64Type
	__VDLType_types_time_Duration1.AppendField("Seconds", __VDLType_types_time_Duration3)
	__VDLType_types_time_Duration4 := vdl.Int32Type
	__VDLType_types_time_Duration1.AppendField("Nanos", __VDLType_types_time_Duration4)
	__VDLType_types_time_DurationBuilder.Build()
	__VDLType_types_time_Durationv, err := __VDLType_types_time_Duration2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_types_time_Durationv
}
func init() {
	__VDLType_types_time_Duration = __VDLType_types_time_Duration_gen()
}

var __VDLType_types_time_Time *vdl.Type

func __VDLType_types_time_Time_gen() *vdl.Type {
	__VDLType_types_time_TimeBuilder := vdl.TypeBuilder{}

	__VDLType_types_time_Time1 := __VDLType_types_time_TimeBuilder.Struct()
	__VDLType_types_time_Time2 := __VDLType_types_time_TimeBuilder.Named("time.Time").AssignBase(__VDLType_types_time_Time1)
	__VDLType_types_time_Time3 := vdl.Int64Type
	__VDLType_types_time_Time1.AppendField("Seconds", __VDLType_types_time_Time3)
	__VDLType_types_time_Time4 := vdl.Int32Type
	__VDLType_types_time_Time1.AppendField("Nanos", __VDLType_types_time_Time4)
	__VDLType_types_time_TimeBuilder.Build()
	__VDLType_types_time_Timev, err := __VDLType_types_time_Time2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_types_time_Timev
}
func init() {
	__VDLType_types_time_Time = __VDLType_types_time_Time_gen()
}

var __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts *vdl.Type

func __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts_gen() *vdl.Type {
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOptsBuilder := vdl.TypeBuilder{}

	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts1 := __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOptsBuilder.Struct()
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts2 := __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOptsBuilder.Named("v.io/v23/services/syncbase.DevModeUpdateVClockOpts").AssignBase(__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts1)
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts3 := vdl.StringType
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts1.AppendField("NtpHost", __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts3)
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts4 := __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOptsBuilder.Struct()
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts5 := __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOptsBuilder.Named("time.Time").AssignBase(__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts4)
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts6 := vdl.Int64Type
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts4.AppendField("Seconds", __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts6)
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts7 := vdl.Int32Type
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts4.AppendField("Nanos", __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts7)
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts1.AppendField("Now", __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts5)
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts8 := __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOptsBuilder.Struct()
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts9 := __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOptsBuilder.Named("time.Duration").AssignBase(__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts8)
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts8.AppendField("Seconds", __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts6)
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts8.AppendField("Nanos", __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts7)
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts1.AppendField("ElapsedTime", __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts9)
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts10 := vdl.BoolType
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts1.AppendField("DoNtpUpdate", __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts10)
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts1.AppendField("DoLocalUpdate", __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts10)
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOptsBuilder.Build()
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOptsv, err := __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOptsv
}
func init() {
	__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts = __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts_gen()
}
func __VDLEnsureNativeBuilt_types() {
	if __VDLTypetypes0 == nil {
		__VDLTypetypes0 = __VDLTypetypes0_gen()
	}
	if __VDLType_types_time_Duration == nil {
		__VDLType_types_time_Duration = __VDLType_types_time_Duration_gen()
	}
	if __VDLType_types_time_Time == nil {
		__VDLType_types_time_Time = __VDLType_types_time_Time_gen()
	}
	if __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts == nil {
		__VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts = __VDLType_types_v_io_v23_services_syncbase_DevModeUpdateVClockOpts_gen()
	}
}
