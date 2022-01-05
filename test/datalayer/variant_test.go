/**
 * MIT License
 *
 * Copyright (c) 2021 Bosch Rexroth AG
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package datalayer_test

import (
	"runtime"
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	a "github.com/stretchr/testify/assert"
)

var vtest *datalayer.Variant

func init() {
	vtest = datalayer.NewVariant()
	runtime.SetFinalizer(vtest, func(vtest *datalayer.Variant) {
		datalayer.DeleteVariant(vtest)
	})
}

func TestDeleteVariant(t *testing.T) {
	a.NotPanics(t, func() { datalayer.DeleteVariant(nil) })
}

// TestBool8 ...
func TestBool8(t *testing.T) {
	value := bool(true)
	a.NotPanics(t, func() { vtest.SetBool8(value) })
	a.Equal(t, value, vtest.GetBool8())
	a.Equal(t, datalayer.VariantTypeBool8, vtest.GetType())

	s := vtest.GetSize()
	a.Equal(t, s, uint64(1))
	c := vtest.GetCount()
	a.Equal(t, c, uint64(1))

	p := vtest.GetData()
	a.NotNil(t, p)

	b := vtest.CheckConvert(datalayer.VariantTypeBool8)
	a.True(t, b == datalayer.Result(0))
	b = vtest.CheckConvert(datalayer.VariantTypeArrayBool8)
	a.True(t, b != datalayer.Result(0))
}

// TestInt8 ...
func TestInt8(t *testing.T) {
	value := int8(-128)
	a.NotPanics(t, func() { vtest.SetInt8(value) })
	a.Equal(t, value, vtest.GetInt8())
	a.Equal(t, datalayer.VariantTypeInt8, vtest.GetType())
	s := vtest.GetSize()
	a.Equal(t, s, uint64(1))
	c := vtest.GetCount()
	a.Equal(t, c, uint64(1))
}

// TestUint8 ...
func TestUint8(t *testing.T) {
	value := uint8(255)
	a.NotPanics(t, func() { vtest.SetUint8(value) })
	a.Equal(t, value, vtest.GetUint8())
	a.Equal(t, datalayer.VariantTypeUint8, vtest.GetType())
}

// TestInt16 ...
func TestInt16(t *testing.T) {
	value := int16(-32768)
	a.NotPanics(t, func() { vtest.SetInt16(value) })
	a.Equal(t, value, vtest.GetInt16())
	a.Equal(t, datalayer.VariantTypeInt16, vtest.GetType())
	s := vtest.GetSize()
	a.Equal(t, s, uint64(2))
	c := vtest.GetCount()
	a.Equal(t, c, uint64(1))
}

// TestUint16 ...
func TestUint16(t *testing.T) {
	value := uint16(65535)
	a.NotPanics(t, func() { vtest.SetUint16(value) })
	a.Equal(t, value, vtest.GetUint16())
	a.Equal(t, datalayer.VariantTypeUint16, vtest.GetType())
}

// TestInt32 ...
func TestInt32(t *testing.T) {
	value := int32(-2147483648)
	a.NotPanics(t, func() { vtest.SetInt32(value) })
	a.Equal(t, value, vtest.GetInt32())
	a.Equal(t, datalayer.VariantTypeInt32, vtest.GetType())
	s := vtest.GetSize()
	a.Equal(t, s, uint64(4))
	c := vtest.GetCount()
	a.Equal(t, c, uint64(1))
}

// TestUint32 ...
func TestUint32(t *testing.T) {
	value := uint32(4294967295)
	a.NotPanics(t, func() { vtest.SetUint32(value) })
	a.Equal(t, value, vtest.GetUint32())
	a.Equal(t, datalayer.VariantTypeUint32, vtest.GetType())
}

// TestInt64 ...
func TestInt64(t *testing.T) {
	value := int64(-1)
	a.NotPanics(t, func() { vtest.SetInt64(value) })
	a.Equal(t, value, vtest.GetInt64())
	a.Equal(t, datalayer.VariantTypeInt64, vtest.GetType())
	s := vtest.GetSize()
	a.Equal(t, s, uint64(8))
	c := vtest.GetCount()
	a.Equal(t, c, uint64(1))
}

// TestUint64 ...
func TestUint64(t *testing.T) {
	value := uint64(1)
	a.NotPanics(t, func() { vtest.SetUint64(value) })
	a.Equal(t, value, vtest.GetUint64())
	a.Equal(t, datalayer.VariantTypeUint64, vtest.GetType())
}

// TestFloat32 ...
func TestFloat32(t *testing.T) {
	value := float32(1.5)
	a.NotPanics(t, func() { vtest.SetFloat32(value) })
	a.Equal(t, value, vtest.GetFloat32())
	a.Equal(t, datalayer.VariantTypeFloat32, vtest.GetType())
}

// TestFloat64 ...
func TestFloat64(t *testing.T) {
	value := float64(1.5)
	a.NotPanics(t, func() { vtest.SetFloat64(value) })
	a.Equal(t, value, vtest.GetFloat64())
	a.Equal(t, datalayer.VariantTypeFloat64, vtest.GetType())
}

// TestString ...
func TestString(t *testing.T) {
	value := string("Test123")
	a.NotPanics(t, func() { vtest.SetString(value) })
	a.Equal(t, value, vtest.GetString())
	a.Equal(t, datalayer.VariantTypeString, vtest.GetType())
}

// TestArrayBool8 ...
func TestArrayBool8(t *testing.T) {
	arr := []bool{true, false}
	a.NotPanics(t, func() { vtest.SetArrayBool8(arr) })
	a.Equal(t, arr, vtest.GetArrayBool8())
	a.Equal(t, datalayer.VariantTypeArrayBool8, vtest.GetType())
	s := vtest.GetSize()
	a.Equal(t, s, uint64(2))
	c := vtest.GetCount()
	a.Equal(t, c, uint64(2))

	a.NotPanics(t, func() { vtest.SetArrayBool8([]bool{}) })
	e := vtest.GetArrayBool8()
	a.True(t, len(e) == 0)
}

// TestArrayInt8 ...
func TestArrayInt8(t *testing.T) {
	arr := []int8{1}
	a.NotPanics(t, func() { vtest.SetArrayInt8(arr) })
	a.Equal(t, arr, vtest.GetArrayInt8())
	a.Equal(t, datalayer.VariantTypeArrayInt8, vtest.GetType())

	a.NotPanics(t, func() { vtest.SetArrayInt8([]int8{}) })
	e := vtest.GetArrayInt8()
	a.True(t, len(e) == 0)
}

// TestArrayUint8 ...
func TestArrayUint8(t *testing.T) {
	arr := []uint8{1}
	a.NotPanics(t, func() { vtest.SetArrayUint8(arr) })
	a.Equal(t, arr, vtest.GetArrayUint8())
	a.Equal(t, datalayer.VariantTypeArrayUint8, vtest.GetType())

	a.NotPanics(t, func() { vtest.SetArrayUint8([]uint8{}) })
	e := vtest.GetArrayUint8()
	a.True(t, len(e) == 0)
}

// TestArrayInt16 ...
func TestArrayInt16(t *testing.T) {
	arr := []int16{1, 2, 3}
	a.NotPanics(t, func() { vtest.SetArrayInt16(arr) })
	a.Equal(t, arr, vtest.GetArrayInt16())
	a.Equal(t, datalayer.VariantTypeArrayInt16, vtest.GetType())
	s := vtest.GetSize()
	a.Equal(t, s, uint64(6))
	c := vtest.GetCount()
	a.Equal(t, c, uint64(3))

	a.NotPanics(t, func() { vtest.SetArrayInt16([]int16{}) })
	e := vtest.GetArrayInt16()
	a.True(t, len(e) == 0)
}

// TestArrayUint16 ...
func TestArrayUint16(t *testing.T) {
	arr := []uint16{1}
	a.NotPanics(t, func() { vtest.SetArrayUint16(arr) })
	a.Equal(t, arr, vtest.GetArrayUint16())
	a.Equal(t, datalayer.VariantTypeArrayUint16, vtest.GetType())

	a.NotPanics(t, func() { vtest.SetArrayUint16([]uint16{}) })
	e := vtest.GetArrayUint16()
	a.True(t, len(e) == 0)
}

// TestArrayInt32 ...
func TestArrayInt32(t *testing.T) {
	arr := []int32{1}
	a.NotPanics(t, func() { vtest.SetArrayInt32(arr) })
	a.Equal(t, arr, vtest.GetArrayInt32())
	a.Equal(t, datalayer.VariantTypeArrayInt32, vtest.GetType())

	a.NotPanics(t, func() { vtest.SetArrayInt32([]int32{}) })
	e := vtest.GetArrayInt32()
	a.True(t, len(e) == 0)
}

// TestArrayUint32 ...
func TestArrayUint32(t *testing.T) {
	arr := []uint32{1}
	a.NotPanics(t, func() { vtest.SetArrayUint32(arr) })
	a.Equal(t, arr, vtest.GetArrayUint32())
	a.Equal(t, datalayer.VariantTypeArrayUint32, vtest.GetType())

	a.NotPanics(t, func() { vtest.SetArrayUint32([]uint32{}) })
	e := vtest.GetArrayUint32()
	a.True(t, len(e) == 0)
}

// TestArrayInt64 ...
func TestArrayInt64(t *testing.T) {
	arr := []int64{1}
	a.NotPanics(t, func() { vtest.SetArrayInt64(arr) })
	a.Equal(t, arr, vtest.GetArrayInt64())
	a.Equal(t, datalayer.VariantTypeArrayInt64, vtest.GetType())

	a.NotPanics(t, func() { vtest.SetArrayInt64([]int64{}) })
	e := vtest.GetArrayInt64()
	a.True(t, len(e) == 0)
}

// TestArrayUint64 ...
func TestArrayUint64(t *testing.T) {
	arr := []uint64{1}
	a.NotPanics(t, func() { vtest.SetArrayUint64(arr) })
	a.Equal(t, arr, vtest.GetArrayUint64())
	a.Equal(t, datalayer.VariantTypeArrayUint64, vtest.GetType())

	a.NotPanics(t, func() { vtest.SetArrayUint64([]uint64{}) })
	e := vtest.GetArrayUint64()
	a.True(t, len(e) == 0)
}

// TestArrayFloat32 ...
func TestArrayFloat32(t *testing.T) {
	arr := []float32{1}
	a.NotPanics(t, func() { vtest.SetArrayFloat32(arr) })
	a.Equal(t, arr, vtest.GetArrayFloat32())
	a.Equal(t, datalayer.VariantTypeArrayFloat32, vtest.GetType())

	a.NotPanics(t, func() { vtest.SetArrayFloat32([]float32{}) })
	e := vtest.GetArrayFloat32()
	a.True(t, len(e) == 0)
}

// TestArrayFloat64 ...
func TestArrayFloat64(t *testing.T) {
	arr := []float64{1}
	a.NotPanics(t, func() { vtest.SetArrayFloat64(arr) })
	a.Equal(t, arr, vtest.GetArrayFloat64())
	a.Equal(t, datalayer.VariantTypeArrayFloat64, vtest.GetType())

	a.NotPanics(t, func() { vtest.SetArrayFloat64([]float64{}) })
	e := vtest.GetArrayFloat64()
	a.True(t, len(e) == 0)
}

// TestArrayString ...
func TestArrayString(t *testing.T) {
	arr := []string{"Test", "123"}
	a.NotPanics(t, func() { vtest.SetArrayString(arr) })
	a.Equal(t, arr, vtest.GetArrayString())
	a.Equal(t, datalayer.VariantTypeArrayString, vtest.GetType())

	a.NotPanics(t, func() { vtest.SetArrayString([]string{}) })
	e := vtest.GetArrayString()
	a.True(t, len(e) == 0)
}

func TestFlatbuffers(t *testing.T) {
	a.NotPanics(t, func() { vtest.SetFlatbuffers([]byte{}) })
	e := vtest.GetFlatbuffers()
	a.NotNil(t, e)
	a.True(t, len(e) == 0)
}
func TestCopy(t *testing.T) {
	a.NotPanics(t, func() { vtest.SetBool8(true) })
	val := datalayer.NewVariant()
	defer datalayer.DeleteVariant(val)
	a.NotPanics(t, func() { vtest.Copy(val) })
	a.Equal(t, datalayer.VariantTypeBool8, val.GetType())
}
