package datalayer_test

import (
	"fmt"

	"github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/datalayer"
)

func ExampleVariant_CheckConvert() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()

	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	// Convert the variant to another type.
	r := v.CheckConvert(datalayer.VariantTypeBool8)
	if r != datalayer.ResultOk {
		return
	}
}

func ExampleVariant_Copy() {
	// Initialize two variants and set values
	v := datalayer.NewVariant()
	val := datalayer.NewVariant()

	// Destroy variants instance.
	defer datalayer.DeleteVariant(v)
	defer datalayer.DeleteVariant(val)

	// Copy the content of a variant to another variant.
	v.SetBool8(true)
	v.Copy(val)
}

// Get Methods ------------------------------------
func ExampleVariant_GetBool8() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// DeleteVariant destroys the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetBool8(true)

	// GetBool8 returns the value of the variant as a bool
	r := v.GetBool8()
	fmt.Print(r)
}

func ExampleVariant_GetInt8() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()

	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetInt8(-128)

	// Return the value of the variant as a int8
	r := v.GetInt8()
	fmt.Print(r)
}

func ExampleVariant_GetInt16() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetInt16(-32768)

	// Return the value of the variant as a int16
	r := v.GetInt16()
	fmt.Print(r)
}

func ExampleVariant_GetUint16() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetUint16(65535)

	// Return the value of the variant as a uint16
	r := v.GetUint16()
	fmt.Print(r)
}

func ExampleVariant_GetInt32() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetInt32(-2147483648)

	// Return the value of the variant as a int32
	r := v.GetInt32()
	fmt.Print(r)
}

func ExampleVariant_GetUint32() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetUint32(4294967295)

	// Return the value of the variant as a Uint32
	r := v.GetUint32()
	fmt.Print(r)
}

func ExampleVariant_GetInt64() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetInt64(-1)

	// Return the value of the variant as a int64
	r := v.GetInt64()
	fmt.Print(r)
}

func ExampleVariant_GetUint64() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetUint64(1)

	// Return the value of the variant as a uint64
	r := v.GetUint64()
	fmt.Print(r)
}

func ExampleVariant_GetFloat32() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetFloat32(1.5)

	// Return the value of the variant as a float
	r := v.GetFloat32()
	fmt.Print(r)
}

func ExampleVariant_GetFloat64() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetFloat64(1.5)

	// Return the value of the variant as a double
	r := v.GetFloat64()
	fmt.Print(r)
}

func ExampleVariant_GetString() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetString("someString")

	// Return the value of the variant as a
	r := v.GetString()
	fmt.Print(r)
}

func ExampleVariant_GetArrayBool8() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetArrayBool8([]bool{true, false})

	// Return the array of int8 if the type is a array of int8 otherwise null.
	r := v.GetArrayBool8()
	fmt.Print(r)
}

func ExampleVariant_GetArrayInt8() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetArrayInt8([]int8{1})

	// Return the array of int8 if the type is a array of int8 otherwise null.
	r := v.GetArrayInt8()
	fmt.Print(r)
}

func ExampleVariant_GetArrayUint8() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetArrayUint8([]uint8{1})

	// Return the array of uint8 if the type is a array of uint8 otherwise null.
	r := v.GetArrayUint8()
	fmt.Print(r)
}

func ExampleVariant_GetArrayInt16() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetArrayInt16([]int16{1, 2, 3})

	// Return the array of int16 if the type is a array of int16 otherwise null.
	r := v.GetArrayInt16()
	fmt.Print(r)
}

func ExampleVariant_GetArrayUint16() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetArrayUint16([]uint16{1})

	// Return the array of uint16 if the type is a array of uint16 otherwise null.
	r := v.GetArrayUint16()
	fmt.Print(r)
}

func ExampleVariant_GetArrayInt32() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetArrayInt32([]int32{1})

	// Return the array of int32 if the type is a array of int32 otherwise null.
	r := v.GetArrayInt32()
	fmt.Print(r)
}

func ExampleVariant_GetArrayUint32() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetArrayUint32([]uint32{1})

	// Return the array of uint32 if the type is a array of uint32 otherwise null.
	r := v.GetArrayUint32()
	fmt.Print(r)
}

func ExampleVariant_GetArrayFloat32() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetArrayFloat32([]float32{1})

	// Return the array of float if the type is a array of float otherwise null.
	r := v.GetArrayFloat32()
	fmt.Print(r)

}

func ExampleVariant_GetArrayFloat64() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetArrayFloat64([]float64{1})

	// Return the array of double if the type is a array of double otherwise null.
	r := v.GetArrayFloat64()
	fmt.Print(r)
}

func ExampleVariant_GetArrayString() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetArrayString([]string{"some", "string"})

	// Return the type of the variant.
	r := v.GetArrayString()
	fmt.Print(r)
}

func ExampleVariant_GetFlatbuffers() {
	// Initialize the variant and set the value
	v := datalayer.NewVariant()
	// Destroy the variant instance.
	defer datalayer.DeleteVariant(v)

	v.SetFlatbuffers([]byte{})

	// Return the flatbuffers if the type is a flatbuffers otherwise null.
	r := v.GetFlatbuffers()
	fmt.Print(r)
}

// Set Methods ------------------------------------
func ExampleVariant_SetBool8() {
	// Initialize the variant
	v := datalayer.NewVariant()

	// DeleteVariant destroys the variant instance.
	defer datalayer.DeleteVariant(v)

	// SetBool8 sets a bool[True, False] value.
	v.SetBool8(true)
}

func ExampleVariant_SetInt8() {
	// Initialize the variant
	v := datalayer.NewVariant()

	// DeleteVariant destroys the variant instance.
	defer datalayer.DeleteVariant(v)

	// SetInt8 sets a int8[-128, 127] value.
	v.SetInt8(-128)
}

func ExampleVariant_SetUint8() {
	// Initialize the variant
	v := datalayer.NewVariant()

	// DeleteVariant destroys the variant instance.
	defer datalayer.DeleteVariant(v)

	// SetUint8 sets a uint8[0, 255] value.
	v.SetUint8(145)
}

func ExampleVariant_SetInt16() {
	// Initialize the variant
	v := datalayer.NewVariant()

	// DeleteVariant destroys the variant instance.
	defer datalayer.DeleteVariant(v)

	// SetUint8 sets a uint8[0, 65.535] value.
	v.SetInt16(-24239)
}

func ExampleVariant_SetUint16() {
	// Initialize the variant
	v := datalayer.NewVariant()

	// DeleteVariant destroys the variant instance.
	defer datalayer.DeleteVariant(v)

	// SetUint16 sets a uint16[0, 65.535] value.
	v.SetUint16(65535)
}

func ExampleVariant_SetInt32() {
	// Initialize the variant
	v := datalayer.NewVariant()

	// DeleteVariant destroys the variant instance.
	defer datalayer.DeleteVariant(v)

	// SetInt32 sets a int32[-2.147.483.648, 2.147.483.647] value.
	v.SetInt32(-2147483648)
}

func ExampleVariant_SetUint32() {
	// Initialize the variant
	v := datalayer.NewVariant()

	// DeleteVariant destroys the variant instance.
	defer datalayer.DeleteVariant(v)

	// SetUint32 sets a uint32[0, 4.294.967.295] value.
	v.SetUint32(4294967295)
}

func ExampleVariant_SetFloat32() {
	// Initialize the variant
	v := datalayer.NewVariant()

	// DeleteVariant destroys the variant instance.
	defer datalayer.DeleteVariant(v)

	// SetFloat32 sets a float[1.2E-38, 3.4E+3] value.
	v.SetFloat32(1.5)
}

func ExampleVariant_SetFloat64() {
	// Initialize the variant
	v := datalayer.NewVariant()

	// DeleteVariant destroys the variant instance.
	defer datalayer.DeleteVariant(v)

	// SetFloat64 sets a double[2.3E-308, 1.7E+308] value.
	v.SetFloat64(1.5)
}

func ExampleVariant_SetString() {
	// Initialize the variant
	v := datalayer.NewVariant()

	// DeleteVariant destroys the variant instance.
	defer datalayer.DeleteVariant(v)

	//SetString sets a string.
	v.SetString("someString")
}
