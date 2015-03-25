// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: vomtype.vdl

package testdata

import (
	// VDL system imports
	"fmt"
	"v.io/v23/vdl"
)

// vomdata config types
type ConvertGroup struct {
	Name        string
	PrimaryType *vdl.Type
	Values      []*vdl.Value
}

func (ConvertGroup) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.ConvertGroup"
}) {
}

type VomdataStruct struct {
	EncodeDecodeData []*vdl.Value
	CompatData       map[string][]*vdl.Type
	ConvertData      map[string][]ConvertGroup
}

func (VomdataStruct) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.VomdataStruct"
}) {
}

// Named Types
type NBool bool

func (NBool) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NBool"
}) {
}

type NString string

func (NString) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NString"
}) {
}

type NByteSlice []byte

func (NByteSlice) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NByteSlice"
}) {
}

type NByteArray [4]byte

func (NByteArray) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NByteArray"
}) {
}

type NByte byte

func (NByte) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NByte"
}) {
}

type NUint16 uint16

func (NUint16) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NUint16"
}) {
}

type NUint32 uint32

func (NUint32) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NUint32"
}) {
}

type NUint64 uint64

func (NUint64) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NUint64"
}) {
}

type NInt16 int16

func (NInt16) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NInt16"
}) {
}

type NInt32 int32

func (NInt32) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NInt32"
}) {
}

type NInt64 int64

func (NInt64) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NInt64"
}) {
}

type NFloat32 float32

func (NFloat32) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NFloat32"
}) {
}

type NFloat64 float64

func (NFloat64) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NFloat64"
}) {
}

type NComplex64 complex64

func (NComplex64) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NComplex64"
}) {
}

type NComplex128 complex128

func (NComplex128) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NComplex128"
}) {
}

type NArray2Uint64 [2]uint64

func (NArray2Uint64) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NArray2Uint64"
}) {
}

type NListUint64 []uint64

func (NListUint64) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NListUint64"
}) {
}

type NSetUint64 map[uint64]struct{}

func (NSetUint64) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NSetUint64"
}) {
}

type NMapUint64String map[uint64]string

func (NMapUint64String) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NMapUint64String"
}) {
}

type NStruct struct {
	A bool
	B string
	C int64
}

func (NStruct) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NStruct"
}) {
}

type NEnum int

const (
	NEnumA NEnum = iota
	NEnumB
	NEnumC
)

// NEnumAll holds all labels for NEnum.
var NEnumAll = [...]NEnum{NEnumA, NEnumB, NEnumC}

// NEnumFromString creates a NEnum from a string label.
func NEnumFromString(label string) (x NEnum, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *NEnum) Set(label string) error {
	switch label {
	case "A", "a":
		*x = NEnumA
		return nil
	case "B", "b":
		*x = NEnumB
		return nil
	case "C", "c":
		*x = NEnumC
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in testdata.NEnum", label)
}

// String returns the string label of x.
func (x NEnum) String() string {
	switch x {
	case NEnumA:
		return "A"
	case NEnumB:
		return "B"
	case NEnumC:
		return "C"
	}
	return ""
}

func (NEnum) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.NEnum"
	Enum struct{ A, B, C string }
}) {
}

type (
	// NUnion represents any single field of the NUnion union type.
	NUnion interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the NUnion union type.
		__VDLReflect(__NUnionReflect)
	}
	// NUnionA represents field A of the NUnion union type.
	NUnionA struct{ Value bool }
	// NUnionB represents field B of the NUnion union type.
	NUnionB struct{ Value string }
	// NUnionC represents field C of the NUnion union type.
	NUnionC struct{ Value int64 }
	// __NUnionReflect describes the NUnion union type.
	__NUnionReflect struct {
		Name  string "v.io/v23/vom/testdata.NUnion"
		Type  NUnion
		Union struct {
			A NUnionA
			B NUnionB
			C NUnionC
		}
	}
)

func (x NUnionA) Index() int                   { return 0 }
func (x NUnionA) Interface() interface{}       { return x.Value }
func (x NUnionA) Name() string                 { return "A" }
func (x NUnionA) __VDLReflect(__NUnionReflect) {}

func (x NUnionB) Index() int                   { return 1 }
func (x NUnionB) Interface() interface{}       { return x.Value }
func (x NUnionB) Name() string                 { return "B" }
func (x NUnionB) __VDLReflect(__NUnionReflect) {}

func (x NUnionC) Index() int                   { return 2 }
func (x NUnionC) Interface() interface{}       { return x.Value }
func (x NUnionC) Name() string                 { return "C" }
func (x NUnionC) __VDLReflect(__NUnionReflect) {}

// Nested Custom Types
type MBool NBool

func (MBool) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.MBool"
}) {
}

type MStruct struct {
	A bool
	B NBool
	C MBool
	D *NStruct
	E *vdl.Type
	F *vdl.Value
}

func (MStruct) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.MStruct"
}) {
}

type MList []NListUint64

func (MList) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.MList"
}) {
}

type MMap map[NFloat32]NListUint64

func (MMap) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.MMap"
}) {
}

// Recursive Type Definitions
type RecA []RecA

func (RecA) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.RecA"
}) {
}

type RecX []RecY

func (RecX) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.RecX"
}) {
}

type RecY []RecX

func (RecY) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.RecY"
}) {
}

type RecStruct struct {
	A *RecStruct
}

func (RecStruct) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.RecStruct"
}) {
}

// Additional types for compatibility and conversion checks
type ListString []string

func (ListString) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.ListString"
}) {
}

type Array3String [3]string

func (Array3String) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.Array3String"
}) {
}

type Array4String [4]string

func (Array4String) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.Array4String"
}) {
}

type AbcStruct struct {
	A bool
	B string
	C int64
}

func (AbcStruct) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.AbcStruct"
}) {
}

type AdeStruct struct {
	A bool
	D *vdl.Value
	E *vdl.Type
}

func (AdeStruct) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.AdeStruct"
}) {
}

type XyzStruct struct {
	X bool
	Y MBool
	Z string
}

func (XyzStruct) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.XyzStruct"
}) {
}

type YzStruct struct {
	Y NBool
	Z NString
}

func (YzStruct) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.YzStruct"
}) {
}

type ZStruct struct {
	Z string
}

func (ZStruct) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.ZStruct"
}) {
}

type MapOnlyStruct struct {
	Key1 int64
	Key2 uint32
	Key3 complex128
}

func (MapOnlyStruct) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.MapOnlyStruct"
}) {
}

type StructOnlyMap map[string]uint64

func (StructOnlyMap) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.StructOnlyMap"
}) {
}

type MapSetStruct struct {
	Feat bool
	Tire bool
	Eel  bool
}

func (MapSetStruct) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.MapSetStruct"
}) {
}

type SetStructMap map[string]bool

func (SetStructMap) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.SetStructMap"
}) {
}

type MapStructSet map[string]struct{}

func (MapStructSet) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.MapStructSet"
}) {
}

type SetOnlyMap map[float64]bool

func (SetOnlyMap) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.SetOnlyMap"
}) {
}

type SometimesSetMap map[float64]*vdl.Value

func (SometimesSetMap) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.SometimesSetMap"
}) {
}

type MapOnlySet map[float64]struct{}

func (MapOnlySet) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.MapOnlySet"
}) {
}

type SetOnlyA map[bool]struct{}

func (SetOnlyA) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.SetOnlyA"
}) {
}

type SetOnlyA2 map[NBool]struct{}

func (SetOnlyA2) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.SetOnlyA2"
}) {
}

type SetOnlyB map[int16]struct{}

func (SetOnlyB) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.SetOnlyB"
}) {
}

type SetOnlyB2 map[NInt16]struct{}

func (SetOnlyB2) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.SetOnlyB2"
}) {
}

type MapOnlyA map[uint32]uint32

func (MapOnlyA) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.MapOnlyA"
}) {
}

type MapOnlyA2 map[int64]float64

func (MapOnlyA2) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.MapOnlyA2"
}) {
}

type MapOnlyB map[bool]string

func (MapOnlyB) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.MapOnlyB"
}) {
}

type MapOnlyB2 map[NBool]NString

func (MapOnlyB2) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.MapOnlyB2"
}) {
}

type (
	// BdeUnion represents any single field of the BdeUnion union type.
	BdeUnion interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the BdeUnion union type.
		__VDLReflect(__BdeUnionReflect)
	}
	// BdeUnionB represents field B of the BdeUnion union type.
	BdeUnionB struct{ Value string }
	// BdeUnionD represents field D of the BdeUnion union type.
	BdeUnionD struct{ Value *vdl.Value }
	// BdeUnionE represents field E of the BdeUnion union type.
	BdeUnionE struct{ Value *vdl.Type }
	// __BdeUnionReflect describes the BdeUnion union type.
	__BdeUnionReflect struct {
		Name  string "v.io/v23/vom/testdata.BdeUnion"
		Type  BdeUnion
		Union struct {
			B BdeUnionB
			D BdeUnionD
			E BdeUnionE
		}
	}
)

func (x BdeUnionB) Index() int                     { return 0 }
func (x BdeUnionB) Interface() interface{}         { return x.Value }
func (x BdeUnionB) Name() string                   { return "B" }
func (x BdeUnionB) __VDLReflect(__BdeUnionReflect) {}

func (x BdeUnionD) Index() int                     { return 1 }
func (x BdeUnionD) Interface() interface{}         { return x.Value }
func (x BdeUnionD) Name() string                   { return "D" }
func (x BdeUnionD) __VDLReflect(__BdeUnionReflect) {}

func (x BdeUnionE) Index() int                     { return 2 }
func (x BdeUnionE) Interface() interface{}         { return x.Value }
func (x BdeUnionE) Name() string                   { return "E" }
func (x BdeUnionE) __VDLReflect(__BdeUnionReflect) {}

type BrieEnum int

const (
	BrieEnumGlee BrieEnum = iota
	BrieEnumBrie
	BrieEnumThree
)

// BrieEnumAll holds all labels for BrieEnum.
var BrieEnumAll = [...]BrieEnum{BrieEnumGlee, BrieEnumBrie, BrieEnumThree}

// BrieEnumFromString creates a BrieEnum from a string label.
func BrieEnumFromString(label string) (x BrieEnum, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *BrieEnum) Set(label string) error {
	switch label {
	case "Glee", "glee":
		*x = BrieEnumGlee
		return nil
	case "Brie", "brie":
		*x = BrieEnumBrie
		return nil
	case "Three", "three":
		*x = BrieEnumThree
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in testdata.BrieEnum", label)
}

// String returns the string label of x.
func (x BrieEnum) String() string {
	switch x {
	case BrieEnumGlee:
		return "Glee"
	case BrieEnumBrie:
		return "Brie"
	case BrieEnumThree:
		return "Three"
	}
	return ""
}

func (BrieEnum) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.BrieEnum"
	Enum struct{ Glee, Brie, Three string }
}) {
}

type BeanEnum int

const (
	BeanEnumBean BeanEnum = iota
)

// BeanEnumAll holds all labels for BeanEnum.
var BeanEnumAll = [...]BeanEnum{BeanEnumBean}

// BeanEnumFromString creates a BeanEnum from a string label.
func BeanEnumFromString(label string) (x BeanEnum, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *BeanEnum) Set(label string) error {
	switch label {
	case "Bean", "bean":
		*x = BeanEnumBean
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in testdata.BeanEnum", label)
}

// String returns the string label of x.
func (x BeanEnum) String() string {
	switch x {
	case BeanEnumBean:
		return "Bean"
	}
	return ""
}

func (BeanEnum) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.BeanEnum"
	Enum struct{ Bean string }
}) {
}

type FoodEnum int

const (
	FoodEnumBean FoodEnum = iota
	FoodEnumBrie
	FoodEnumCherry
)

// FoodEnumAll holds all labels for FoodEnum.
var FoodEnumAll = [...]FoodEnum{FoodEnumBean, FoodEnumBrie, FoodEnumCherry}

// FoodEnumFromString creates a FoodEnum from a string label.
func FoodEnumFromString(label string) (x FoodEnum, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *FoodEnum) Set(label string) error {
	switch label {
	case "Bean", "bean":
		*x = FoodEnumBean
		return nil
	case "Brie", "brie":
		*x = FoodEnumBrie
		return nil
	case "Cherry", "cherry":
		*x = FoodEnumCherry
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in testdata.FoodEnum", label)
}

// String returns the string label of x.
func (x FoodEnum) String() string {
	switch x {
	case FoodEnumBean:
		return "Bean"
	case FoodEnumBrie:
		return "Brie"
	case FoodEnumCherry:
		return "Cherry"
	}
	return ""
}

func (FoodEnum) __VDLReflect(struct {
	Name string "v.io/v23/vom/testdata.FoodEnum"
	Enum struct{ Bean, Brie, Cherry string }
}) {
}

func init() {
	vdl.Register((*ConvertGroup)(nil))
	vdl.Register((*VomdataStruct)(nil))
	vdl.Register((*NBool)(nil))
	vdl.Register((*NString)(nil))
	vdl.Register((*NByteSlice)(nil))
	vdl.Register((*NByteArray)(nil))
	vdl.Register((*NByte)(nil))
	vdl.Register((*NUint16)(nil))
	vdl.Register((*NUint32)(nil))
	vdl.Register((*NUint64)(nil))
	vdl.Register((*NInt16)(nil))
	vdl.Register((*NInt32)(nil))
	vdl.Register((*NInt64)(nil))
	vdl.Register((*NFloat32)(nil))
	vdl.Register((*NFloat64)(nil))
	vdl.Register((*NComplex64)(nil))
	vdl.Register((*NComplex128)(nil))
	vdl.Register((*NArray2Uint64)(nil))
	vdl.Register((*NListUint64)(nil))
	vdl.Register((*NSetUint64)(nil))
	vdl.Register((*NMapUint64String)(nil))
	vdl.Register((*NStruct)(nil))
	vdl.Register((*NEnum)(nil))
	vdl.Register((*NUnion)(nil))
	vdl.Register((*MBool)(nil))
	vdl.Register((*MStruct)(nil))
	vdl.Register((*MList)(nil))
	vdl.Register((*MMap)(nil))
	vdl.Register((*RecA)(nil))
	vdl.Register((*RecX)(nil))
	vdl.Register((*RecY)(nil))
	vdl.Register((*RecStruct)(nil))
	vdl.Register((*ListString)(nil))
	vdl.Register((*Array3String)(nil))
	vdl.Register((*Array4String)(nil))
	vdl.Register((*AbcStruct)(nil))
	vdl.Register((*AdeStruct)(nil))
	vdl.Register((*XyzStruct)(nil))
	vdl.Register((*YzStruct)(nil))
	vdl.Register((*ZStruct)(nil))
	vdl.Register((*MapOnlyStruct)(nil))
	vdl.Register((*StructOnlyMap)(nil))
	vdl.Register((*MapSetStruct)(nil))
	vdl.Register((*SetStructMap)(nil))
	vdl.Register((*MapStructSet)(nil))
	vdl.Register((*SetOnlyMap)(nil))
	vdl.Register((*SometimesSetMap)(nil))
	vdl.Register((*MapOnlySet)(nil))
	vdl.Register((*SetOnlyA)(nil))
	vdl.Register((*SetOnlyA2)(nil))
	vdl.Register((*SetOnlyB)(nil))
	vdl.Register((*SetOnlyB2)(nil))
	vdl.Register((*MapOnlyA)(nil))
	vdl.Register((*MapOnlyA2)(nil))
	vdl.Register((*MapOnlyB)(nil))
	vdl.Register((*MapOnlyB2)(nil))
	vdl.Register((*BdeUnion)(nil))
	vdl.Register((*BrieEnum)(nil))
	vdl.Register((*BeanEnum)(nil))
	vdl.Register((*FoodEnum)(nil))
}
