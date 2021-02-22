// +build linux

package datalayer

/*
#include <stdbool.h>
#include <stdlib.h>
#include "system.h"
*/
import "C"
import "unsafe"

// GetArrayInt64 function
func (v *Variant) GetArrayInt64() []int64 {
	var cdata *C.long = C.DLR_variantGetArrayOfINT64(v.this)
	length := C.DLR_variantGetCount(v.this)
	return (*[1 << 28]int64)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayUint64 function
func (v *Variant) GetArrayUint64() []uint64 {
	var cdata *C.ulong = C.DLR_variantGetArrayOfUINT64(v.this)
	length := C.DLR_variantGetCount(v.this)
	return (*[1 << 28]uint64)(unsafe.Pointer(cdata))[:length:length]
}

// SetInt64 function
func (v *Variant) SetInt64(data int64) {
	C.DLR_variantSetINT64(v.this, C.long(data))
}

// SetUint64 function
func (v *Variant) SetUint64(data uint64) {
	C.DLR_variantSetUINT64(v.this, C.ulong(data))
}

// SetArrayBool8 function
func (v *Variant) SetArrayBool8(data []bool) {
	C.DLR_variantSetARRAY_OF_BOOL8(v.this, (*C.bool)(unsafe.Pointer(&data[0])), C.ulong(len(data)))
}

// SetArrayInt8 function
func (v *Variant) SetArrayInt8(data []int8) {
	C.DLR_variantSetARRAY_OF_INT8(v.this, (*C.schar)(unsafe.Pointer(&data[0])), C.ulong(len(data)))
}

// SetArrayUint8 function
func (v *Variant) SetArrayUint8(data []uint8) {
	C.DLR_variantSetARRAY_OF_UINT8(v.this, (*C.uchar)(unsafe.Pointer(&data[0])), C.ulong(len(data)))
}

// SetArrayInt16 function
func (v *Variant) SetArrayInt16(data []int16) {
	C.DLR_variantSetARRAY_OF_INT16(v.this, (*C.short)(unsafe.Pointer(&data[0])), C.ulong(len(data)))
}

// SetArrayUint16 function
func (v *Variant) SetArrayUint16(data []uint16) {
	C.DLR_variantSetARRAY_OF_UINT16(v.this, (*C.ushort)(unsafe.Pointer(&data[0])), C.ulong(len(data)))
}

// SetArrayInt32 function
func (v *Variant) SetArrayInt32(data []int32) {
	C.DLR_variantSetARRAY_OF_INT32(v.this, (*C.int)(unsafe.Pointer(&data[0])), C.ulong(len(data)))
}

// SetArrayUint32 function
func (v *Variant) SetArrayUint32(data []uint32) {
	C.DLR_variantSetARRAY_OF_UINT32(v.this, (*C.uint)(unsafe.Pointer(&data[0])), C.ulong(len(data)))
}

// SetArrayInt64 function
func (v *Variant) SetArrayInt64(data []int64) {
	C.DLR_variantSetARRAY_OF_INT64(v.this, (*C.long)(unsafe.Pointer(&data[0])), C.ulong(len(data)))
}

// SetArrayUint64 function
func (v *Variant) SetArrayUint64(data []uint64) {
	C.DLR_variantSetARRAY_OF_UINT64(v.this, (*C.ulong)(unsafe.Pointer(&data[0])), C.ulong(len(data)))
}

// SetArrayFloat32 function
func (v *Variant) SetArrayFloat32(data []float32) {
	C.DLR_variantSetARRAY_OF_FLOAT32(v.this, (*C.float)(unsafe.Pointer(&data[0])), C.ulong(len(data)))
}

// SetArrayFloat64 function
func (v *Variant) SetArrayFloat64(data []float64) {
	C.DLR_variantSetARRAY_OF_FLOAT64(v.this, (*C.double)(unsafe.Pointer(&data[0])), C.ulong(len(data)))
}

// SetArrayString function
func (v *Variant) SetArrayString(data []string) {
	arr := make([]*C.char, len(data))
	for i, s := range data {
		cs := C.CString(s)
		defer C.free(unsafe.Pointer(cs))
		arr[i] = cs
	}
	C.DLR_variantSetARRAY_OF_STRING(v.this, &arr[0], C.ulong(len(data)))
}

// SetFlatbuffers function
func (v *Variant) SetFlatbuffers(data []byte) {
	C.DLR_variantSetFlatbuffers(v.this, (*C.schar)(unsafe.Pointer(&data[0])), C.ulong(len(data)))
}
