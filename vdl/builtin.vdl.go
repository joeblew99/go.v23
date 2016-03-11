// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: builtin.vdl

package vdl

import (
	"fmt"
	"reflect"
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
	Name string `vdl:"v.io/v23/vdl.WireError"`
}) {
}

func (m *WireError) FillVDLTarget(t Target, tt *Type) error {
	if __VDLType_builtin_v_io_v23_vdl_WireError == nil || __VDLTypebuiltin0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Id")
	if err != ErrFieldNoExist && err != nil {
		return err
	}
	if err != ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Id), StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("RetryCode")
	if err != ErrFieldNoExist && err != nil {
		return err
	}
	if err != ErrFieldNoExist {

		if err := m.RetryCode.FillVDLTarget(fieldTarget5, __VDLType_builtin_v_io_v23_vdl_WireRetryCode); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("Msg")
	if err != ErrFieldNoExist && err != nil {
		return err
	}
	if err != ErrFieldNoExist {
		if err := fieldTarget7.FromString(string(m.Msg), StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	keyTarget8, fieldTarget9, err := fieldsTarget1.StartField("ParamList")
	if err != ErrFieldNoExist && err != nil {
		return err
	}
	if err != ErrFieldNoExist {

		listTarget10, err := fieldTarget9.StartList(__VDLTypebuiltin1, len(m.ParamList))
		if err != nil {
			return err
		}
		for i, elem12 := range m.ParamList {
			elemTarget11, err := listTarget10.StartElem(i)
			if err != nil {
				return err
			}

			if elem12 == nil {
				if err := elemTarget11.FromNil(AnyType); err != nil {
					return err
				}
			} else {
				if err := FromValue(elemTarget11, elem12); err != nil {
					return err
				}
			}
			if err := listTarget10.FinishElem(elemTarget11); err != nil {
				return err
			}
		}
		if err := fieldTarget9.FinishList(listTarget10); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget8, fieldTarget9); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *WireError) MakeVDLTarget() Target {
	return &WireErrorTarget{Value: m}
}

type WireErrorTarget struct {
	Value *WireError
	TargetBase
	FieldsTargetBase
}

func (t *WireErrorTarget) StartFields(tt *Type) (FieldsTarget, error) {
	if !Compatible(tt, __VDLType_builtin_v_io_v23_vdl_WireError) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_builtin_v_io_v23_vdl_WireError)
	}
	return t, nil
}
func (t *WireErrorTarget) StartField(name string) (key, field Target, _ error) {
	switch name {
	case "Id":
		val, err := &StringTarget{Value: &t.Value.Id}, error(nil)
		return nil, val, err
	case "RetryCode":
		val, err := &WireRetryCodeTarget{Value: &t.Value.RetryCode}, error(nil)
		return nil, val, err
	case "Msg":
		val, err := &StringTarget{Value: &t.Value.Msg}, error(nil)
		return nil, val, err
	case "ParamList":
		val, err := &builtin5b5d616e79Target{Value: &t.Value.ParamList}, error(nil)
		return nil, val, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_builtin_v_io_v23_vdl_WireError)
	}
}
func (t *WireErrorTarget) FinishField(_, _ Target) error {
	return nil
}
func (t *WireErrorTarget) FinishFields(_ FieldsTarget) error {

	return nil
}

type WireRetryCodeTarget struct {
	Value *WireRetryCode
	TargetBase
}

func (t *WireRetryCodeTarget) FromEnumLabel(src string, tt *Type) error {
	if !Compatible(tt, __VDLType_builtin_v_io_v23_vdl_WireRetryCode) {
		return fmt.Errorf("type %v incompatible with %v", tt, __VDLType_builtin_v_io_v23_vdl_WireRetryCode)
	}
	switch src {
	case "NoRetry":
		*t.Value = 0
	case "RetryConnection":
		*t.Value = 1
	case "RetryRefetch":
		*t.Value = 2
	case "RetryBackoff":
		*t.Value = 3
	default:
		return fmt.Errorf("label %s not in enum %v", src, __VDLType_builtin_v_io_v23_vdl_WireRetryCode)
	}

	return nil
}

type builtin5b5d616e79Target struct {
	Value *[]*Value
	TargetBase
	ListTargetBase
}

func (t *builtin5b5d616e79Target) StartList(tt *Type, len int) (ListTarget, error) {
	if !Compatible(tt, __VDLTypebuiltin1) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLTypebuiltin1)
	}
	if cap(*t.Value) < len {
		*t.Value = make([]*Value, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *builtin5b5d616e79Target) StartElem(index int) (elem Target, _ error) {
	return ReflectTarget(reflect.ValueOf(&(*t.Value)[index]))
}
func (t *builtin5b5d616e79Target) FinishElem(elem Target) error {
	return nil
}
func (t *builtin5b5d616e79Target) FinishList(elem ListTarget) error {

	return nil
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
var WireRetryCodeAll = [...]WireRetryCode{WireRetryCodeNoRetry, WireRetryCodeRetryConnection, WireRetryCodeRetryRefetch, WireRetryCodeRetryBackoff}

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
	Name string `vdl:"v.io/v23/vdl.WireRetryCode"`
	Enum struct{ NoRetry, RetryConnection, RetryRefetch, RetryBackoff string }
}) {
}

func (m *WireRetryCode) FillVDLTarget(t Target, tt *Type) error {
	if err := t.FromEnumLabel((*m).String(), __VDLType_builtin_v_io_v23_vdl_WireRetryCode); err != nil {
		return err
	}
	return nil
}

func (m *WireRetryCode) MakeVDLTarget() Target {
	return &WireRetryCodeTarget{Value: m}
}

func init() {
	Register((*WireError)(nil))
	Register((*WireRetryCode)(nil))
}

var __VDLTypebuiltin0 *Type = TypeOf((*WireError)(nil))
var __VDLTypebuiltin1 *Type = TypeOf([]*Value(nil))
var __VDLType_builtin_v_io_v23_vdl_WireError *Type = TypeOf(WireError{})
var __VDLType_builtin_v_io_v23_vdl_WireRetryCode *Type = TypeOf(WireRetryCodeNoRetry)

func __VDLEnsureNativeBuilt_builtin() {
}
