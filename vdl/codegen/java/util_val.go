package java

import (
	"fmt"
	"strconv"
	"strings"

	"veyron.io/veyron/veyron2/vdl"
	"veyron.io/veyron/veyron2/vdl/compile"
)

// javaConstVal returns the value string for the provided constant value.
func javaConstVal(v *vdl.Value, env *compile.Env) (ret string) {
	ret = javaVal(v, env)
	if def := env.FindTypeDef(v.Type()); def != nil && def.File != compile.BuiltInFile { // User-defined type.
		ret = "new " + javaType(v.Type(), false, env) + "(" + ret + ")"
	}
	return
}

// javaVal returns the value string for the provided Value.
func javaVal(v *vdl.Value, env *compile.Env) string {
	const longSuffix = "L"
	const floatSuffix = "f"

	switch v.Kind() {
	case vdl.Bool:
		if v.Bool() {
			return "true"
		} else {
			return "false"
		}
	case vdl.Byte:
		return "(byte)" + strconv.FormatUint(uint64(v.Byte()), 10)
	case vdl.Uint16:
		return "(short)" + strconv.FormatUint(v.Uint(), 10)
	case vdl.Int16:
		return "(short)" + strconv.FormatInt(v.Int(), 10)
	case vdl.Uint32:
		return strconv.FormatUint(v.Uint(), 10)
	case vdl.Int32:
		return strconv.FormatInt(v.Int(), 10)
	case vdl.Uint64:
		return strconv.FormatUint(v.Uint(), 10) + longSuffix
	case vdl.Int64:
		return strconv.FormatInt(v.Int(), 10) + longSuffix
	case vdl.Float32, vdl.Float64:
		c := strconv.FormatFloat(v.Float(), 'g', -1, bitlen(v.Kind()))
		if strings.Index(c, ".") == -1 {
			c += ".0"
		}
		if v.Kind() == vdl.Float32 {
			return c + floatSuffix
		}
		return c
	case vdl.Complex64, vdl.Complex128:
		r := strconv.FormatFloat(real(v.Complex()), 'g', -1, bitlen(v.Kind()))
		i := strconv.FormatFloat(imag(v.Complex()), 'g', -1, bitlen(v.Kind()))
		return fmt.Sprintf("new %s(%s, %s)", javaType(v.Type(), true, env), r, i)
	case vdl.String:
		return strconv.Quote(v.RawString())
	}
	if v.Type().IsBytes() {
		return strconv.Quote(string(v.Bytes()))
	}
	// TODO(spetrovic): Handle Enum, List, Map, Struct, OneOf, Any
	panic(fmt.Errorf("vdl: javaVal unhandled type %v %v", v.Kind(), v.Type()))
}
