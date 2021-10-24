package rules

import "fmt"

type Type int

const (
	// INT
	Nil Type = iota
	Int
	Int64
	Int32
	Int16
	Int8
	SliceInt
	SliceInt64
	SliceInt32
	SliceInt16
	SliceInt8

	// Uint
	Uint
	Uintptr
	Uint64
	Uint32
	Uint16
	Uint8
	SliceUint
	SliceUintptr
	SliceUint64
	SliceUint32
	SliceUint16
	SliceUint8

	// Float
	Float64
	Float32
	SliceFloat64
	SliceFloat32

	// Complex
	Complex64
	Complex128
	SliceComplex64
	SliceComplex128

	// String
	String
	SliceString

	// Interface
	Interface
	SliceInterface
)

const (
	typeDontMatch = `is not type "%s"`
)

func IsType(v interface{}, t Type) string {
	switch v.(type) {
	case int:
		if t != Int {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case int64:
		if t != Int64 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case int32:
		if t != Int32 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case int16:
		if t != Int16 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case int8:
		if t != Int8 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []int:
		if t != SliceInt {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []int64:
		if t != SliceInt64 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []int32:
		if t != SliceInt32 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []int16:
		if t != SliceInt16 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []int8:
		if t != SliceInt8 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case uint:
		if t != Uint {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case uint64:
		if t != Uint64 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case uint32:
		if t != Uint32 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case uint16:
		if t != Uint16 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case uint8:
		if t != Uint8 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []uint:
		if t != SliceUint {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []uint64:
		if t != SliceUint64 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []uint32:
		if t != SliceUint32 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []uint16:
		if t != SliceUint16 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []uint8:
		if t != SliceUint8 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case float64:
		if t != Float64 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case float32:
		if t != Float32 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []float64:
		if t != SliceFloat64 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []float32:
		if t != SliceFloat32 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case complex64:
		if t != Complex64 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case complex128:
		if t != Complex128 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []complex64:
		if t != SliceComplex64 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []complex128:
		if t != SliceComplex128 {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case string:
		if t != String {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	case []string:
		if t != SliceString {
			return fmt.Sprintf(typeDontMatch, t.toString())
		}
	default:
		return "unrecognized type"
	}
	return ""
}

func (t Type) toString() string {
	switch t {
	case Int:
		return "int"
	case Int64:
		return "int64"
	case Int32:
		return "int32"
	case Int16:
		return "int16"
	case Int8:
		return "int8"
	case SliceInt:
		return "[]int8"
	case SliceInt64:
		return "[]int8"
	case SliceInt32:
		return "[]int8"
	case SliceInt16:
		return "[]int8"
	case SliceInt8:
		return "[]int8"
	case Uint:
		return "uint"
	case Uint64:
		return "uint64"
	case Uint32:
		return "uint32"
	case Uint16:
		return "uint16"
	case Uint8:
		return "uint8"
	case SliceUint:
		return "[]uint8"
	case SliceUint64:
		return "[]uint8"
	case SliceUint32:
		return "[]uint8"
	case SliceUint16:
		return "[]uint8"
	case SliceUint8:
		return "[]uint8"
	case Float64:
		return "float64"
	case Float32:
		return "float32"
	case SliceFloat64:
		return "[]float64"
	case SliceFloat32:
		return "[]float32"
	case Complex64:
		return "complex64"
	case Complex128:
		return "complex128"
	case SliceComplex64:
		return "[]complex64"
	case SliceComplex128:
		return "[]complex128"
	case String:
		return "string"
	case SliceString:
		return "[]string"
	case Interface:
		return "interface"
	case SliceInterface:
		return "[]interface"
	default:
		return "unrecognized type"
	}
}
