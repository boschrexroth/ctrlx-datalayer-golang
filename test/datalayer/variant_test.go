package datalayer_test

import (
	"runtime"
	"testing"

	"github.com/boschrexroth/ctrlx-datalayer-golang/pkg/datalayer"
	a "github.com/stretchr/testify/assert"
)

var v *datalayer.Variant

func init() {
	v = datalayer.NewVariant()
	runtime.SetFinalizer(v, func(v *datalayer.Variant) {
		datalayer.DeleteVariant(v)
	})
}

// TestBool8 ...
func TestBool8(t *testing.T) {
	value := bool(true)
	a.NotPanics(t, func() { v.SetBool8(value) })
	a.Equal(t, value, v.GetBool8())
	a.Equal(t, datalayer.VariantTypeBool8, v.GetType())
}

// TestInt8 ...
func TestInt8(t *testing.T) {
	value := int8(-128)
	a.NotPanics(t, func() { v.SetInt8(value) })
	a.Equal(t, value, v.GetInt8())
	a.Equal(t, datalayer.VariantTypeInt8, v.GetType())
}

// TestUint8 ...
func TestUint8(t *testing.T) {
	value := uint8(255)
	a.NotPanics(t, func() { v.SetUint8(value) })
	a.Equal(t, value, v.GetUint8())
	a.Equal(t, datalayer.VariantTypeUint8, v.GetType())
}

// TestInt16 ...
func TestInt16(t *testing.T) {
	value := int16(-32768)
	a.NotPanics(t, func() { v.SetInt16(value) })
	a.Equal(t, value, v.GetInt16())
	a.Equal(t, datalayer.VariantTypeInt16, v.GetType())
}

// TestUint16 ...
func TestUint16(t *testing.T) {
	value := uint16(65535)
	a.NotPanics(t, func() { v.SetUint16(value) })
	a.Equal(t, value, v.GetUint16())
	a.Equal(t, datalayer.VariantTypeUint16, v.GetType())
}

// TestInt32 ...
func TestInt32(t *testing.T) {
	value := int32(-2147483648)
	a.NotPanics(t, func() { v.SetInt32(value) })
	a.Equal(t, value, v.GetInt32())
	a.Equal(t, datalayer.VariantTypeInt32, v.GetType())
}

// TestUint32 ...
func TestUint32(t *testing.T) {
	value := uint32(4294967295)
	a.NotPanics(t, func() { v.SetUint32(value) })
	a.Equal(t, value, v.GetUint32())
	a.Equal(t, datalayer.VariantTypeUint32, v.GetType())
}

// TestInt64 ...
func TestInt64(t *testing.T) {
	value := int64(-1)
	a.NotPanics(t, func() { v.SetInt64(value) })
	a.Equal(t, value, v.GetInt64())
	a.Equal(t, datalayer.VariantTypeInt64, v.GetType())
}

// TestUint64 ...
func TestUint64(t *testing.T) {
	value := uint64(1)
	a.NotPanics(t, func() { v.SetUint64(value) })
	a.Equal(t, value, v.GetUint64())
	a.Equal(t, datalayer.VariantTypeUint64, v.GetType())
}

// TestFloat32 ...
func TestFloat32(t *testing.T) {
	value := float32(1.5)
	a.NotPanics(t, func() { v.SetFloat32(value) })
	a.Equal(t, value, v.GetFloat32())
	a.Equal(t, datalayer.VariantTypeFloat32, v.GetType())
}

// TestFloat64 ...
func TestFloat64(t *testing.T) {
	value := float64(1.5)
	a.NotPanics(t, func() { v.SetFloat64(value) })
	a.Equal(t, value, v.GetFloat64())
	a.Equal(t, datalayer.VariantTypeFloat64, v.GetType())
}

// TestString ...
func TestString(t *testing.T) {
	value := string("Test123")
	a.NotPanics(t, func() { v.SetString(value) })
	a.Equal(t, value, v.GetString())
	a.Equal(t, datalayer.VariantTypeString, v.GetType())
}

// TestArrayBool8 ...
func TestArrayBool8(t *testing.T) {
	arr := []bool{true}
	a.NotPanics(t, func() { v.SetArrayBool8(arr) })
	a.Equal(t, arr, v.GetArrayBool8())
	a.Equal(t, datalayer.VariantTypeArrayBool8, v.GetType())
}

// TestArrayInt8 ...
func TestArrayInt8(t *testing.T) {
	arr := []int8{1}
	a.NotPanics(t, func() { v.SetArrayInt8(arr) })
	a.Equal(t, arr, v.GetArrayInt8())
	a.Equal(t, datalayer.VariantTypeArrayInt8, v.GetType())
}

// TestArrayUint8 ...
func TestArrayUint8(t *testing.T) {
	arr := []uint8{1}
	a.NotPanics(t, func() { v.SetArrayUint8(arr) })
	a.Equal(t, arr, v.GetArrayUint8())
	a.Equal(t, datalayer.VariantTypeArrayUint8, v.GetType())
}

// TestArrayInt16 ...
func TestArrayInt16(t *testing.T) {
	arr := []int16{1}
	a.NotPanics(t, func() { v.SetArrayInt16(arr) })
	a.Equal(t, arr, v.GetArrayInt16())
	a.Equal(t, datalayer.VariantTypeArrayInt16, v.GetType())
}

// TestArrayUint16 ...
func TestArrayUint16(t *testing.T) {
	arr := []uint16{1}
	a.NotPanics(t, func() { v.SetArrayUint16(arr) })
	a.Equal(t, arr, v.GetArrayUint16())
	a.Equal(t, datalayer.VariantTypeArrayUint16, v.GetType())
}

// TestArrayInt32 ...
func TestArrayInt32(t *testing.T) {
	arr := []int32{1}
	a.NotPanics(t, func() { v.SetArrayInt32(arr) })
	a.Equal(t, arr, v.GetArrayInt32())
	a.Equal(t, datalayer.VariantTypeArrayInt32, v.GetType())
}

// TestArrayUint32 ...
func TestArrayUint32(t *testing.T) {
	arr := []uint32{1}
	a.NotPanics(t, func() { v.SetArrayUint32(arr) })
	a.Equal(t, arr, v.GetArrayUint32())
	a.Equal(t, datalayer.VariantTypeArrayUint32, v.GetType())
}

// TestArrayInt64 ...
func TestArrayInt64(t *testing.T) {
	arr := []int64{1}
	a.NotPanics(t, func() { v.SetArrayInt64(arr) })
	a.Equal(t, arr, v.GetArrayInt64())
	a.Equal(t, datalayer.VariantTypeArrayInt64, v.GetType())
}

// TestArrayUint64 ...
func TestArrayUint64(t *testing.T) {
	arr := []uint64{1}
	a.NotPanics(t, func() { v.SetArrayUint64(arr) })
	a.Equal(t, arr, v.GetArrayUint64())
	a.Equal(t, datalayer.VariantTypeArrayUint64, v.GetType())
}

// TestArrayFloat32 ...
func TestArrayFloat32(t *testing.T) {
	arr := []float32{1}
	a.NotPanics(t, func() { v.SetArrayFloat32(arr) })
	a.Equal(t, arr, v.GetArrayFloat32())
	a.Equal(t, datalayer.VariantTypeArrayFloat32, v.GetType())
}

// TestArrayFloat64 ...
func TestArrayFloat64(t *testing.T) {
	arr := []float64{1}
	a.NotPanics(t, func() { v.SetArrayFloat64(arr) })
	a.Equal(t, arr, v.GetArrayFloat64())
	a.Equal(t, datalayer.VariantTypeArrayFloat64, v.GetType())
}

// TestArrayString ...
func TestArrayString(t *testing.T) {
	arr := []string{"Test", "123"}
	a.NotPanics(t, func() { v.SetArrayString(arr) })
	a.Equal(t, arr, v.GetArrayString())
	a.Equal(t, datalayer.VariantTypeArrayString, v.GetType())
}
