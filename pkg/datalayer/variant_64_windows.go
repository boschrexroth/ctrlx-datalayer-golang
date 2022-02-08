//go:build windows
// +build windows

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

/*
#include <stdbool.h>
#include <stdlib.h>
#include <system.h>
*/
import "C"
import "unsafe"

// GetArrayInt64 function
func (v *Variant) GetArrayInt64() []int64 {
	var cdata *C.longlong = C.DLR_variantGetArrayOfINT64(v.this)
	length := C.DLR_variantGetCount(v.this)
	if length == 0 {
		return []int64{}
	}
	return (*[1 << 28]int64)(unsafe.Pointer(cdata))[:length:length]
}

// GetArrayUint64 function
func (v *Variant) GetArrayUint64() []uint64 {
	var cdata *C.ulonglong = C.DLR_variantGetArrayOfUINT64(v.this)
	length := C.DLR_variantGetCount(v.this)
	if length == 0 {
		return []uint64{}
	}
	return (*[1 << 28]uint64)(unsafe.Pointer(cdata))[:length:length]
}

// SetInt64 function
func (v *Variant) SetInt64(data int64) {
	C.DLR_variantSetINT64(v.this, C.longlong(data))
}

// SetUint64 function
func (v *Variant) SetUint64(data uint64) {
	C.DLR_variantSetUINT64(v.this, C.ulonglong(data))
}

// SetArrayBool8 function
func (v *Variant) SetArrayBool8(data []bool) {
	if len(data) == 0 {
		C.DLR_variantSetARRAY_OF_BOOL8(v.this, nil, C.ulonglong(0))
	} else {
		C.DLR_variantSetARRAY_OF_BOOL8(v.this, (*C.bool)(unsafe.Pointer(&data[0])), C.ulonglong(len(data)))
	}
}

// SetArrayInt8 function
func (v *Variant) SetArrayInt8(data []int8) {
	if len(data) == 0 {
		C.DLR_variantSetARRAY_OF_INT8(v.this, nil, C.ulonglong(0))
	} else {
		C.DLR_variantSetARRAY_OF_INT8(v.this, (*C.schar)(unsafe.Pointer(&data[0])), C.ulonglong(len(data)))
	}
}

// SetArrayUint8 function
func (v *Variant) SetArrayUint8(data []uint8) {
	if len(data) == 0 {
		C.DLR_variantSetARRAY_OF_UINT8(v.this, nil, C.ulonglong(0))
	} else {
		C.DLR_variantSetARRAY_OF_UINT8(v.this, (*C.uchar)(unsafe.Pointer(&data[0])), C.ulonglong(len(data)))
	}
}

// SetArrayInt16 function
func (v *Variant) SetArrayInt16(data []int16) {
	if len(data) == 0 {
		C.DLR_variantSetARRAY_OF_INT16(v.this, nil, C.ulonglong(0))
	} else {
		C.DLR_variantSetARRAY_OF_INT16(v.this, (*C.short)(unsafe.Pointer(&data[0])), C.ulonglong(len(data)))
	}
}

// SetArrayUint16 function
func (v *Variant) SetArrayUint16(data []uint16) {
	if len(data) == 0 {
		C.DLR_variantSetARRAY_OF_UINT16(v.this, nil, C.ulonglong(0))
	} else {
		C.DLR_variantSetARRAY_OF_UINT16(v.this, (*C.ushort)(unsafe.Pointer(&data[0])), C.ulonglong(len(data)))
	}
}

// SetArrayInt32 function
func (v *Variant) SetArrayInt32(data []int32) {
	if len(data) == 0 {
		C.DLR_variantSetARRAY_OF_INT32(v.this, nil, C.ulonglong(0))
	} else {
		C.DLR_variantSetARRAY_OF_INT32(v.this, (*C.int)(unsafe.Pointer(&data[0])), C.ulonglong(len(data)))
	}
}

// SetArrayUint32 function
func (v *Variant) SetArrayUint32(data []uint32) {
	if len(data) == 0 {
		C.DLR_variantSetARRAY_OF_UINT32(v.this, nil, C.ulonglong(0))
	} else {
		C.DLR_variantSetARRAY_OF_UINT32(v.this, (*C.uint)(unsafe.Pointer(&data[0])), C.ulonglong(len(data)))
	}
}

// SetArrayInt64 function
func (v *Variant) SetArrayInt64(data []int64) {
	if len(data) == 0 {
		C.DLR_variantSetARRAY_OF_INT64(v.this, nil, C.ulonglong(0))
	} else {
		C.DLR_variantSetARRAY_OF_INT64(v.this, (*C.longlong)(unsafe.Pointer(&data[0])), C.ulonglong(len(data)))
	}
}

// SetArrayUint64 function
func (v *Variant) SetArrayUint64(data []uint64) {
	if len(data) == 0 {
		C.DLR_variantSetARRAY_OF_UINT64(v.this, nil, C.ulonglong(0))
	} else {
		C.DLR_variantSetARRAY_OF_UINT64(v.this, (*C.ulonglong)(unsafe.Pointer(&data[0])), C.ulonglong(len(data)))
	}
}

// SetArrayFloat32 function
func (v *Variant) SetArrayFloat32(data []float32) {
	if len(data) == 0 {
		C.DLR_variantSetARRAY_OF_FLOAT32(v.this, nil, C.ulonglong(0))
	} else {
		C.DLR_variantSetARRAY_OF_FLOAT32(v.this, (*C.float)(unsafe.Pointer(&data[0])), C.ulonglong(len(data)))
	}
}

// SetArrayFloat64 function
func (v *Variant) SetArrayFloat64(data []float64) {
	if len(data) == 0 {
		C.DLR_variantSetARRAY_OF_FLOAT64(v.this, nil, C.ulonglong(0))
	} else {
		C.DLR_variantSetARRAY_OF_FLOAT64(v.this, (*C.double)(unsafe.Pointer(&data[0])), C.ulonglong(len(data)))
	}
}

// SetArrayString function
func (v *Variant) SetArrayString(data []string) {
	if len(data) == 0 {
		C.DLR_variantSetARRAY_OF_STRING(v.this, nil, C.ulonglong(0))
	} else {
		arr := make([]*C.char, len(data))
		for i, s := range data {
			cs := C.CString(s)
			defer C.free(unsafe.Pointer(cs))
			arr[i] = cs
		}
		C.DLR_variantSetARRAY_OF_STRING(v.this, &arr[0], C.ulonglong(len(data)))
	}
}

// SetFlatbuffers function
func (v *Variant) SetFlatbuffers(data []byte) {
	if len(data) == 0 {
		C.DLR_variantSetFlatbuffers(v.this, nil, C.ulonglong(0))
	} else {
		C.DLR_variantSetFlatbuffers(v.this, (*C.schar)(unsafe.Pointer(&data[0])), C.ulonglong(len(data)))
	}
}
