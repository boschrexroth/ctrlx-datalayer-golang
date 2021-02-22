package datalayer

//#include <stdbool.h>
//#include <stdlib.h>
//#include "system.h"
import "C"
import "unsafe"

// Result ulong
type Result C.DLR_RESULT

// VariantType enum
type VariantType C.enum_DLR_VARIANT_TYPE

// VariantType enum definition
const (
	VariantTypeUnknown VariantType = C.DLR_VARIANT_TYPE_UNKNOWN
	VariantTypeBool8   VariantType = C.DLR_VARIANT_TYPE_BOOL8
	VariantTypeInt8    VariantType = C.DLR_VARIANT_TYPE_INT8
	VariantTypeUint8   VariantType = C.DLR_VARIANT_TYPE_UINT8
	VariantTypeInt16   VariantType = C.DLR_VARIANT_TYPE_INT16
	VariantTypeUint16  VariantType = C.DLR_VARIANT_TYPE_UINT16
	VariantTypeInt32   VariantType = C.DLR_VARIANT_TYPE_INT32
	VariantTypeUint32  VariantType = C.DLR_VARIANT_TYPE_UINT32
	VariantTypeInt64   VariantType = C.DLR_VARIANT_TYPE_INT64
	VariantTypeUint64  VariantType = C.DLR_VARIANT_TYPE_UINT64
	VariantTypeFloat32 VariantType = C.DLR_VARIANT_TYPE_FLOAT32
	VariantTypeFloat64 VariantType = C.DLR_VARIANT_TYPE_FLOAT64
	VariantTypeString  VariantType = C.DLR_VARIANT_TYPE_STRING

	VariantTypeArrayBool8   VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_BOOL8
	VariantTypeArrayInt8    VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_INT8
	VariantTypeArrayUint8   VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_UINT8
	VariantTypeArrayInt16   VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_INT16
	VariantTypeArrayUint16  VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_UINT16
	VariantTypeArrayInt32   VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_INT32
	VariantTypeArrayUint32  VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_UINT32
	VariantTypeArrayInt64   VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_INT64
	VariantTypeArrayUint64  VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_UINT64
	VariantTypeArrayFloat32 VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_FLOAT32
	VariantTypeArrayFloat64 VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_FLOAT64
	VariantTypeArrayString  VariantType = C.DLR_VARIANT_TYPE_ARRAY_OF_STRING

	VariantTypeRaw         VariantType = C.DLR_VARIANT_TYPE_RAW
	VariantTypeFlatbuffers VariantType = C.DLR_VARIANT_TYPE_FLATBUFFERS
)

// Variant class
type Variant struct {
	this C.DLR_VARIANT
}

// NewVariant constructor
func NewVariant() *Variant {
	ptr := C.DLR_variantCreate()
	return &Variant{this: ptr}
}

// DeleteVariant destructor
func DeleteVariant(v *Variant) {
	C.DLR_variantDelete(v.this)
}

// GetType function
func (v *Variant) GetType() VariantType {
	t := C.DLR_variantGetType(v.this)
	ty := VariantType(t)
	return ty
}

// GetData function
func (v *Variant) GetData() unsafe.Pointer {
	return unsafe.Pointer(C.DLR_variantGetData(v.this))
}

// GetSize function
func (v *Variant) GetSize() uint64 {
	return uint64(C.DLR_variantGetSize(v.this))
}

// GetCount function
func (v *Variant) GetCount() uint64 {
	return uint64(C.DLR_variantGetCount(v.this))
}

// CheckConvert function
func (v *Variant) CheckConvert(ty VariantType) Result {
	return Result(C.DLR_variantCheckConvert(v.this, C.DLR_VARIANT_TYPE(ty)))
}

// Copy function
func (v *Variant) Copy(dest *Variant) Result {
	return Result(C.DLR_variantCopy(dest.this, v.this))
}

// --------------------------------------------------------------------------- getVar

// GetBool8 function
func (v *Variant) GetBool8() bool {
	return bool(C.DLR_variantGetBOOL8(v.this))
}

// GetInt8 function
func (v *Variant) GetInt8() int8 {
	return int8(C.DLR_variantGetINT8(v.this))
}

// GetUint8 function
func (v *Variant) GetUint8() uint8 {
	return uint8(C.DLR_variantGetUINT8(v.this))
}

// GetInt16 function
func (v *Variant) GetInt16() int16 {
	return int16(C.DLR_variantGetINT16(v.this))
}

// GetUint16 function
func (v *Variant) GetUint16() uint16 {
	return uint16(C.DLR_variantGetUINT16(v.this))
}

// GetInt32 function
func (v *Variant) GetInt32() int32 {
	return int32(C.DLR_variantGetINT32(v.this))
}

// GetUint32 function
func (v *Variant) GetUint32() uint32 {
	return uint32(C.DLR_variantGetUINT32(v.this))
}

// GetInt64 function
func (v *Variant) GetInt64() int64 {
	return int64(C.DLR_variantGetINT64(v.this))
}

// GetUint64 function
func (v *Variant) GetUint64() uint64 {
	return uint64(C.DLR_variantGetUINT64(v.this))
}

// GetFloat32 function
func (v *Variant) GetFloat32() float32 {
	return float32(C.DLR_variantGetFLOAT32(v.this))
}

// GetFloat64 function
func (v *Variant) GetFloat64() float64 {
	return float64(C.DLR_variantGetFLOAT64(v.this))
}

// GetString function
func (v *Variant) GetString() string {
	p := C.DLR_variantGetSTRING(v.this)
	s := C.GoString(p)
	return s
}

// GetArrayBool8 function
func (v *Variant) GetArrayBool8() []bool {
	var cdata *C.bool = C.DLR_variantGetArrayOfBOOL8(v.this)
	length := C.DLR_variantGetCount(v.this)
	return (*[1 << 28]bool)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayInt8 function
func (v *Variant) GetArrayInt8() []int8 {
	var cdata *C.schar = C.DLR_variantGetArrayOfINT8(v.this)
	length := C.DLR_variantGetCount(v.this)
	return (*[1 << 28]int8)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayUint8 function
func (v *Variant) GetArrayUint8() []uint8 {
	var cdata *C.uchar = C.DLR_variantGetArrayOfUINT8(v.this)
	length := C.DLR_variantGetCount(v.this)
	return (*[1 << 28]uint8)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayInt16 function
func (v *Variant) GetArrayInt16() []int16 {
	var cdata *C.short = C.DLR_variantGetArrayOfINT16(v.this)
	length := C.DLR_variantGetCount(v.this)
	return (*[1 << 28]int16)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayUint16 function
func (v *Variant) GetArrayUint16() []uint16 {
	var cdata *C.ushort = C.DLR_variantGetArrayOfUINT16(v.this)
	length := C.DLR_variantGetCount(v.this)
	return (*[1 << 28]uint16)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayInt32 function
func (v *Variant) GetArrayInt32() []int32 {
	var cdata *C.int = C.DLR_variantGetArrayOfINT32(v.this)
	length := C.DLR_variantGetCount(v.this)
	return (*[1 << 28]int32)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayUint32 function
func (v *Variant) GetArrayUint32() []uint32 {
	var cdata *C.uint = C.DLR_variantGetArrayOfUINT32(v.this)
	length := C.DLR_variantGetCount(v.this)
	return (*[1 << 28]uint32)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayFloat32 function
func (v *Variant) GetArrayFloat32() []float32 {
	var cdata *C.float = C.DLR_variantGetArrayOfFLOAT32(v.this)
	length := C.DLR_variantGetCount(v.this)
	return (*[1 << 28]float32)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayFloat64 function
func (v *Variant) GetArrayFloat64() []float64 {
	var cdata *C.double = C.DLR_variantGetArrayOfFLOAT64(v.this)
	length := C.DLR_variantGetCount(v.this)
	return (*[1 << 28]float64)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayString function
func (v *Variant) GetArrayString() []string {
	var cdata **C.char = C.DLR_variantGetArrayOfSTRING(v.this)
	length := C.DLR_variantGetCount(v.this)
	tmp := (*[1 << 28]*C.char)(unsafe.Pointer(cdata))[:length:length]
	slice := make([]string, length)
	for i, s := range tmp {
		slice[i] = C.GoString(s)
	}
	return slice
}

// GetFlatbuffers function
func (v *Variant) GetFlatbuffers() []byte {
	length := C.DLR_variantGetSize(v.this)
	return (*[1 << 28]byte)(unsafe.Pointer(C.DLR_variantGetData(v.this)))[:length:length]
}

// --------------------------------------------------------------------------- setVar

// SetBool8 function
func (v *Variant) SetBool8(data bool) {
	C.DLR_variantSetBOOL8(v.this, C.bool(data))
}

// SetInt8 function
func (v *Variant) SetInt8(data int8) {
	C.DLR_variantSetINT8(v.this, C.schar(data))
}

// SetUint8 function
func (v *Variant) SetUint8(data uint8) {
	C.DLR_variantSetUINT8(v.this, C.uchar(data))
}

// SetInt16 function
func (v *Variant) SetInt16(data int16) {
	C.DLR_variantSetINT16(v.this, C.short(data))
}

// SetUint16 function
func (v *Variant) SetUint16(data uint16) {
	C.DLR_variantSetUINT16(v.this, C.ushort(data))
}

// SetInt32 function
func (v *Variant) SetInt32(data int32) {
	C.DLR_variantSetINT32(v.this, C.int(data))
}

// SetUint32 function
func (v *Variant) SetUint32(data uint32) {
	C.DLR_variantSetUINT32(v.this, C.uint(data))
}

// SetFloat32 function
func (v *Variant) SetFloat32(data float32) {
	C.DLR_variantSetFLOAT32(v.this, C.float(data))
}

// SetFloat64 function
func (v *Variant) SetFloat64(data float64) {
	C.DLR_variantSetFLOAT64(v.this, C.double(data))
}

// SetString function
func (v *Variant) SetString(data string) {
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cdata))
	C.DLR_variantSetSTRING(v.this, cdata)
}
