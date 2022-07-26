// MIT License
//
// Copyright (c) 2021-2022 Bosch Rexroth AG
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package datalayer

//#include <stdbool.h>
//#include <stdlib.h>
//#include <system.h>
import "C"
import (
	"math"
	"time"
	"unsafe"
)

// VariantType enum
type VariantType C.enum_DLR_VARIANT_TYPE

// VariantType enum definition
const (
	VariantTypeUnknown   VariantType = C.DLR_VARIANT_TYPE_UNKNOWN
	VariantTypeBool8     VariantType = C.DLR_VARIANT_TYPE_BOOL8
	VariantTypeInt8      VariantType = C.DLR_VARIANT_TYPE_INT8
	VariantTypeUint8     VariantType = C.DLR_VARIANT_TYPE_UINT8
	VariantTypeInt16     VariantType = C.DLR_VARIANT_TYPE_INT16
	VariantTypeUint16    VariantType = C.DLR_VARIANT_TYPE_UINT16
	VariantTypeInt32     VariantType = C.DLR_VARIANT_TYPE_INT32
	VariantTypeUint32    VariantType = C.DLR_VARIANT_TYPE_UINT32
	VariantTypeInt64     VariantType = C.DLR_VARIANT_TYPE_INT64
	VariantTypeUint64    VariantType = C.DLR_VARIANT_TYPE_UINT64
	VariantTypeFloat32   VariantType = C.DLR_VARIANT_TYPE_FLOAT32
	VariantTypeFloat64   VariantType = C.DLR_VARIANT_TYPE_FLOAT64
	VariantTypeString    VariantType = C.DLR_VARIANT_TYPE_STRING
	VariantTypeTimestamp VariantType = C.DLR_VARIANT_TIMESTAMP

	VariantTypeArrayBool8     VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_BOOL8
	VariantTypeArrayInt8      VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_INT8
	VariantTypeArrayUint8     VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_UINT8
	VariantTypeArrayInt16     VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_INT16
	VariantTypeArrayUint16    VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_UINT16
	VariantTypeArrayInt32     VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_INT32
	VariantTypeArrayUint32    VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_UINT32
	VariantTypeArrayInt64     VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_INT64
	VariantTypeArrayUint64    VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_UINT64
	VariantTypeArrayFloat32   VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_FLOAT32
	VariantTypeArrayFloat64   VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_FLOAT64
	VariantTypeArrayString    VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_STRING
	VariantTypeArrayTimestamp VariantType = C.DLR_VARIANT_ARRAY_OF_TIMESTAMP

	VariantTypeRaw         VariantType = C.DLR_VARIANT_TYPE_RAW
	VariantTypeFlatbuffers VariantType = C.DLR_VARIANT_TYPE_FLATBUFFERS
)

// Variant is a container for a many types of data.
type Variant struct {
	this C.DLR_VARIANT
}

// NewVariant generates the variant instance.
func NewVariant() *Variant {
	ptr := C.DLR_variantCreate()
	return &Variant{this: ptr}
}

// DeleteVariant destroys the variant instance.
func DeleteVariant(v *Variant) {
	if v == nil {
		return
	}
	if v.this == nil {
		return
	}
	C.DLR_variantDelete(v.this)
}

// GetType returns the type of the variant instance.
// It returns the variant type.
func (v *Variant) GetType() VariantType {
	if v.this == nil {
		return VariantTypeUnknown
	}
	t := C.DLR_variantGetType(v.this)
	ty := VariantType(t)
	return ty
}

// GetData takes the pointer to the data of the variant instance.
// It returns the array of bytes.
func (v *Variant) GetData() unsafe.Pointer {
	if v.this == nil {
		return unsafe.Pointer(nil)
	}
	return unsafe.Pointer(C.DLR_variantGetData(v.this))
}

// GetSize gets the size of the type in bytes.
// It returns the size of the type in bytes.
func (v *Variant) GetSize() uint64 {
	if v.this == nil {
		return uint64(0)
	}
	return uint64(C.DLR_variantGetSize(v.this))
}

// GetCount returns the count of elements in the variant (scalar data types = 1, array = count of elements in array).
// It returns the count of a type.
func (v *Variant) GetCount() uint64 {
	if v.this == nil {
		return uint64(0)
	}
	return uint64(C.DLR_variantGetCount(v.this))
}

// CheckConvert checks whether the variant can be converted to another type.
// It returns the status of function call.
func (v *Variant) CheckConvert(ty VariantType) Result {
	if v.this == nil {
		return ResultNotInitialized
	}
	return Result(C.DLR_variantCheckConvert(v.this, C.DLR_VARIANT_TYPE(ty)))
}

// Copy copies the content of a variant to another variant.
// It returns the status of function call or the copy of variant or a tuple.
func (v *Variant) Copy(dest *Variant) Result {
	if dest == nil {
		return ResultNotInitialized
	}
	if dest.this == nil {
		return ResultNotInitialized
	}
	if v.this == nil {
		return ResultNotInitialized
	}
	return Result(C.DLR_variantCopy(dest.this, v.this))
}

// --------------------------------------------------------------------------- getVar

// GetBool8 returns the value of the variant as a bool (auto convert if possible) otherwise 0.
// It returns [True, False]
func (v *Variant) GetBool8() bool {
	return bool(C.DLR_variantGetBOOL8(v.this))
}

// GetInt8 returns the value of the variant as a int8 (auto convert if possible) otherwise 0.
// It returns [-128, 127]
func (v *Variant) GetInt8() int8 {
	return int8(C.DLR_variantGetINT8(v.this))
}

// GetUint8 returns the value of the variant as a uint8 (auto convert if possible) otherwise 0.
// It returns [0, 255]
func (v *Variant) GetUint8() uint8 {
	return uint8(C.DLR_variantGetUINT8(v.this))
}

// GetInt16 returns the value of the variant as a int16 (auto convert if possible) otherwise 0.
// It returns [-32768, 32767]
func (v *Variant) GetInt16() int16 {
	return int16(C.DLR_variantGetINT16(v.this))
}

// GetUint16 returns the value of the variant as a uint16 (auto convert if possible) otherwise 0.
// It returns [0, 65.535]
func (v *Variant) GetUint16() uint16 {
	return uint16(C.DLR_variantGetUINT16(v.this))
}

// GetInt32 returns the value of the variant as a int32 (auto convert if possible) otherwise 0.
// It returns [-2.147.483.648, 2.147.483.647]
func (v *Variant) GetInt32() int32 {
	return int32(C.DLR_variantGetINT32(v.this))
}

// GetUint32 returns the value of the variant as a Uint32 (auto convert if possible) otherwise 0.
// It returns [0, 4.294.967.295]
func (v *Variant) GetUint32() uint32 {
	return uint32(C.DLR_variantGetUINT32(v.this))
}

// GetInt64 returns the value of the variant as a int64 (auto convert if possible) otherwise 0.
// It returns [-9.223.372.036.854.775.808, 9.223.372.036.854.775.807]
func (v *Variant) GetInt64() int64 {
	return int64(C.DLR_variantGetINT64(v.this))
}

// GetUint64 returns the value of the variant as a uint64 (auto convert if possible) otherwise 0.
// It returns [0, 18446744073709551615]
func (v *Variant) GetUint64() uint64 {
	return uint64(C.DLR_variantGetUINT64(v.this))
}

// GetFloat32 returns the value of the variant as a float (auto convert if possible) otherwise 0.
// It returns [1.2E-38, 3.4E+3]
func (v *Variant) GetFloat32() float32 {
	return float32(C.DLR_variantGetFLOAT32(v.this))
}

// GetFloat64 returns the value of the variant as a double (auto convert if possible) otherwise 0.
// It returns [2.3E-308, 1.7E+308]
func (v *Variant) GetFloat64() float64 {
	return float64(C.DLR_variantGetFLOAT64(v.this))
}

// GetString returns the array of bool8 if the type is a array of bool otherwise null.
// It returns the string.
func (v *Variant) GetString() string {
	p := C.DLR_variantGetSTRING(v.this)
	s := C.GoString(p)
	return s
}

// GetArrayBool8 returns the array of int8 if the type is a array of int8 otherwise null.
// It returns the array of bool8.
func (v *Variant) GetArrayBool8() []bool {
	var cdata *C.bool = C.DLR_variantGetArrayOfBOOL8(v.this)
	length := C.DLR_variantGetCount(v.this)
	if length == 0 {
		return []bool{}
	}
	return (*[1 << 28]bool)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayInt8 returns the array of int8 if the type is a array of int8 otherwise null.
// It returns the array of int8.
func (v *Variant) GetArrayInt8() []int8 {
	var cdata *C.schar = C.DLR_variantGetArrayOfINT8(v.this)
	length := C.DLR_variantGetCount(v.this)
	if length == 0 {
		return []int8{}
	}
	return (*[1 << 28]int8)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayUint8 returns the array of uint8 if the type is a array of uint8 otherwise null.
// It returns the array of uint8.
func (v *Variant) GetArrayUint8() []uint8 {
	var cdata *C.uchar = C.DLR_variantGetArrayOfUINT8(v.this)
	length := C.DLR_variantGetCount(v.this)
	if length == 0 {
		return []uint8{}
	}
	return (*[1 << 28]uint8)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayInt16 returns the array of int16 if the type is a array of int16 otherwise null.
// It returns the array of int16.
func (v *Variant) GetArrayInt16() []int16 {
	var cdata *C.short = C.DLR_variantGetArrayOfINT16(v.this)
	length := C.DLR_variantGetCount(v.this)
	if length == 0 {
		return []int16{}
	}
	return (*[1 << 28]int16)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayUint16 returns the array of uint16 if the type is a array of uint16 otherwise null.
// It returns the array of uint16.
func (v *Variant) GetArrayUint16() []uint16 {
	var cdata *C.ushort = C.DLR_variantGetArrayOfUINT16(v.this)
	length := C.DLR_variantGetCount(v.this)
	if length == 0 {
		return []uint16{}
	}
	return (*[1 << 28]uint16)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayInt32 returns the array of int32 if the type is a array of int32 otherwise null.
// It returns the array of int32.
func (v *Variant) GetArrayInt32() []int32 {
	var cdata *C.int = C.DLR_variantGetArrayOfINT32(v.this)
	length := C.DLR_variantGetCount(v.this)
	if length == 0 {
		return []int32{}
	}
	return (*[1 << 28]int32)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayUint32 returns the array of uint32 if the type is a array of uint32 otherwise null.
// It returns the array of uint32.
func (v *Variant) GetArrayUint32() []uint32 {
	var cdata *C.uint = C.DLR_variantGetArrayOfUINT32(v.this)
	length := C.DLR_variantGetCount(v.this)
	if length == 0 {
		return []uint32{}
	}
	return (*[1 << 28]uint32)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayFloat32 returns the array of float if the type is a array of float otherwise null.
// It returns the array of float32.
func (v *Variant) GetArrayFloat32() []float32 {
	var cdata *C.float = C.DLR_variantGetArrayOfFLOAT32(v.this)
	length := C.DLR_variantGetCount(v.this)
	if length == 0 {
		return []float32{}
	}
	return (*[1 << 28]float32)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayFloat64 returns the array of double if the type is a array of double otherwise null.
// It returns the array of float64.
func (v *Variant) GetArrayFloat64() []float64 {
	var cdata *C.double = C.DLR_variantGetArrayOfFLOAT64(v.this)
	length := C.DLR_variantGetCount(v.this)
	if length == 0 {
		return []float64{}
	}
	return (*[1 << 28]float64)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayString returns the the type of the variant.
// It returns the array of strings.
func (v *Variant) GetArrayString() []string {
	var cdata **C.char = C.DLR_variantGetArrayOfSTRING(v.this)
	length := C.DLR_variantGetCount(v.this)
	if length == 0 {
		return []string{}
	}
	tmp := (*[1 << 28]*C.char)(unsafe.Pointer(cdata))[:length:length]
	slice := make([]string, length)
	for i, s := range tmp {
		slice[i] = C.GoString(s)
	}
	return slice
}

// GetFlatbuffers returns the flatbuffers if the type is a flatbuffers otherwise null.
// It returns the flatbuffer (bytearray).
func (v *Variant) GetFlatbuffers() []byte {
	length := C.DLR_variantGetSize(v.this)
	if length == 0 {
		return []byte{}
	}
	return (*[1 << 28]byte)(unsafe.Pointer(C.DLR_variantGetData(v.this)))[:length:length]
}

// --------------------------------------------------------------------------- setVar

// SetBool8 sets a bool value.
// It returns the status of function call.
func (v *Variant) SetBool8(data bool) {
	C.DLR_variantSetBOOL8(v.this, C.bool(data))
}

// SetInt8 sets a int8 value.
// It returns the status of function call.
func (v *Variant) SetInt8(data int8) {
	C.DLR_variantSetINT8(v.this, C.schar(data))
}

// SetUint8 sets a uint8 value.
// It returns the status of function call.
func (v *Variant) SetUint8(data uint8) {
	C.DLR_variantSetUINT8(v.this, C.uchar(data))
}

// SetInt16 sets a int16 value.
// It returns the status of function call.
func (v *Variant) SetInt16(data int16) {
	C.DLR_variantSetINT16(v.this, C.short(data))
}

// SetUint16 sets a uint16 value.
// It returns the status of function call.
func (v *Variant) SetUint16(data uint16) {
	C.DLR_variantSetUINT16(v.this, C.ushort(data))
}

// SetInt32 sets a int32 value.
// It returns the status of function call.
func (v *Variant) SetInt32(data int32) {
	C.DLR_variantSetINT32(v.this, C.int(data))
}

// SetUint32 sets a uint32 value.
// It returns the status of function call.
func (v *Variant) SetUint32(data uint32) {
	C.DLR_variantSetUINT32(v.this, C.uint(data))
}

// SetFloat32 sets a float value.
// It returns the status of function call.
func (v *Variant) SetFloat32(data float32) {
	C.DLR_variantSetFLOAT32(v.this, C.float(data))
}

// SetFloat64 sets a double value.
// It returns the status of function call.
func (v *Variant) SetFloat64(data float64) {
	C.DLR_variantSetFLOAT64(v.this, C.double(data))
}

// SetString sets a string.
// It returns the status of function call.
func (v *Variant) SetString(data string) {
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cdata))
	C.DLR_variantSetSTRING(v.this, cdata)
}

// SetTimestamp sets a timestamp value as (FILETIME) 64 bit 100ns since 1.1.1601 (UTC)
// It returns the status of function call.
func (v *Variant) SetTimestamp(ft uint64) {
	C.DLR_variantSetTimestamp(v.this, C.ulong(ft))
}

func convertUnixTime2FileTime(input int64) int64 {
	temp := input / 100
	temp = temp + 116444736000000000
	return temp
}

// GetTime returns the value of the variant as a time.Time since January 1, 1970 UTC.
// It returns >= January 1, 1970 UTC
func (v *Variant) GetTime() time.Time {
	t := v.GetUint64()
	return getTime(int64(t))
}

// SetTime sets a time.Time value since January 1, 1970 UTC
// It returns the status of function call.
func (v *Variant) SetTime(df time.Time) {
	ns := df.UTC().UnixNano()
	ft := convertUnixTime2FileTime(ns)
	v.SetTimestamp(uint64(ft))
}

// see. https://stackoverflow.com/questions/57901280/calculate-time-time-from-timestamp-starting-from-1601-01-01-in-go
func getTime(input int64) time.Time {
	maxd := time.Duration(math.MaxInt64).Truncate(100 * time.Nanosecond)
	maxdUnits := int64(maxd / 100) // number of 100-ns units

	t := time.Date(1601, 1, 1, 0, 0, 0, 0, time.UTC)
	for input > maxdUnits {
		t = t.Add(maxd)
		input -= maxdUnits
	}
	if input != 0 {
		t = t.Add(time.Duration(input * 100))
	}
	return t
}

// GetArrayTime returns the value of the variant as an array of time.Time since January 1, 1970 UTC.
// It returns >= January 1, 1970 UTC
func (v *Variant) GetArrayTime() []time.Time {
	t := v.GetArrayUint64()
	vals := []time.Time{}
	for _, ft := range t {
		vals = append(vals, getTime(int64(ft)))
	}
	return vals
}

// SetArrayTime sets an array of time.Time value since January 1, 1970 UTC
// It returns the status of function call.
func (v *Variant) SetArrayTime(dfs []time.Time) {
	fts := []uint64{}
	for _, df := range dfs {
		ns := df.UTC().UnixNano()
		fts = append(fts, uint64(convertUnixTime2FileTime(ns)))
	}
	v.SetArrayTimestamp(fts)
}
