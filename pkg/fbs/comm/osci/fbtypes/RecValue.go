// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RecValueT struct {
	Time *TimeT
	ValuesBool bool
	ValuesByte int8
	ValuesShort int16
	ValuesInt int32
	ValuesLong int64
	ValuesUbyte byte
	ValuesUshort uint16
	ValuesUint uint32
	ValuesUlong uint64
	ValuesFloat float32
	ValuesDouble float64
	ValuesString string
}

func (t *RecValueT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	timeOffset := t.Time.Pack(builder)
	valuesStringOffset := builder.CreateString(t.ValuesString)
	RecValueStart(builder)
	RecValueAddTime(builder, timeOffset)
	RecValueAddValuesBool(builder, t.ValuesBool)
	RecValueAddValuesByte(builder, t.ValuesByte)
	RecValueAddValuesShort(builder, t.ValuesShort)
	RecValueAddValuesInt(builder, t.ValuesInt)
	RecValueAddValuesLong(builder, t.ValuesLong)
	RecValueAddValuesUbyte(builder, t.ValuesUbyte)
	RecValueAddValuesUshort(builder, t.ValuesUshort)
	RecValueAddValuesUint(builder, t.ValuesUint)
	RecValueAddValuesUlong(builder, t.ValuesUlong)
	RecValueAddValuesFloat(builder, t.ValuesFloat)
	RecValueAddValuesDouble(builder, t.ValuesDouble)
	RecValueAddValuesString(builder, valuesStringOffset)
	return RecValueEnd(builder)
}

func (rcv *RecValue) UnPackTo(t *RecValueT) {
	t.Time = rcv.Time(nil).UnPack()
	t.ValuesBool = rcv.ValuesBool()
	t.ValuesByte = rcv.ValuesByte()
	t.ValuesShort = rcv.ValuesShort()
	t.ValuesInt = rcv.ValuesInt()
	t.ValuesLong = rcv.ValuesLong()
	t.ValuesUbyte = rcv.ValuesUbyte()
	t.ValuesUshort = rcv.ValuesUshort()
	t.ValuesUint = rcv.ValuesUint()
	t.ValuesUlong = rcv.ValuesUlong()
	t.ValuesFloat = rcv.ValuesFloat()
	t.ValuesDouble = rcv.ValuesDouble()
	t.ValuesString = string(rcv.ValuesString())
}

func (rcv *RecValue) UnPack() *RecValueT {
	if rcv == nil { return nil }
	t := &RecValueT{}
	rcv.UnPackTo(t)
	return t
}

type RecValue struct {
	_tab flatbuffers.Table
}

func GetRootAsRecValue(buf []byte, offset flatbuffers.UOffsetT) *RecValue {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RecValue{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRecValue(buf []byte, offset flatbuffers.UOffsetT) *RecValue {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RecValue{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RecValue) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RecValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *RecValue) Time(obj *Time) *Time {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Time)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *RecValue) ValuesBool() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *RecValue) MutateValuesBool(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *RecValue) ValuesByte() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RecValue) MutateValuesByte(n int8) bool {
	return rcv._tab.MutateInt8Slot(8, n)
}

func (rcv *RecValue) ValuesShort() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RecValue) MutateValuesShort(n int16) bool {
	return rcv._tab.MutateInt16Slot(10, n)
}

func (rcv *RecValue) ValuesInt() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RecValue) MutateValuesInt(n int32) bool {
	return rcv._tab.MutateInt32Slot(12, n)
}

func (rcv *RecValue) ValuesLong() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RecValue) MutateValuesLong(n int64) bool {
	return rcv._tab.MutateInt64Slot(14, n)
}

func (rcv *RecValue) ValuesUbyte() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RecValue) MutateValuesUbyte(n byte) bool {
	return rcv._tab.MutateByteSlot(16, n)
}

func (rcv *RecValue) ValuesUshort() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RecValue) MutateValuesUshort(n uint16) bool {
	return rcv._tab.MutateUint16Slot(18, n)
}

func (rcv *RecValue) ValuesUint() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RecValue) MutateValuesUint(n uint32) bool {
	return rcv._tab.MutateUint32Slot(20, n)
}

func (rcv *RecValue) ValuesUlong() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RecValue) MutateValuesUlong(n uint64) bool {
	return rcv._tab.MutateUint64Slot(22, n)
}

func (rcv *RecValue) ValuesFloat() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *RecValue) MutateValuesFloat(n float32) bool {
	return rcv._tab.MutateFloat32Slot(24, n)
}

func (rcv *RecValue) ValuesDouble() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *RecValue) MutateValuesDouble(n float64) bool {
	return rcv._tab.MutateFloat64Slot(26, n)
}

func (rcv *RecValue) ValuesString() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func RecValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(13)
}
func RecValueAddTime(builder *flatbuffers.Builder, time flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(time), 0)
}
func RecValueAddValuesBool(builder *flatbuffers.Builder, valuesBool bool) {
	builder.PrependBoolSlot(1, valuesBool, false)
}
func RecValueAddValuesByte(builder *flatbuffers.Builder, valuesByte int8) {
	builder.PrependInt8Slot(2, valuesByte, 0)
}
func RecValueAddValuesShort(builder *flatbuffers.Builder, valuesShort int16) {
	builder.PrependInt16Slot(3, valuesShort, 0)
}
func RecValueAddValuesInt(builder *flatbuffers.Builder, valuesInt int32) {
	builder.PrependInt32Slot(4, valuesInt, 0)
}
func RecValueAddValuesLong(builder *flatbuffers.Builder, valuesLong int64) {
	builder.PrependInt64Slot(5, valuesLong, 0)
}
func RecValueAddValuesUbyte(builder *flatbuffers.Builder, valuesUbyte byte) {
	builder.PrependByteSlot(6, valuesUbyte, 0)
}
func RecValueAddValuesUshort(builder *flatbuffers.Builder, valuesUshort uint16) {
	builder.PrependUint16Slot(7, valuesUshort, 0)
}
func RecValueAddValuesUint(builder *flatbuffers.Builder, valuesUint uint32) {
	builder.PrependUint32Slot(8, valuesUint, 0)
}
func RecValueAddValuesUlong(builder *flatbuffers.Builder, valuesUlong uint64) {
	builder.PrependUint64Slot(9, valuesUlong, 0)
}
func RecValueAddValuesFloat(builder *flatbuffers.Builder, valuesFloat float32) {
	builder.PrependFloat32Slot(10, valuesFloat, 0.0)
}
func RecValueAddValuesDouble(builder *flatbuffers.Builder, valuesDouble float64) {
	builder.PrependFloat64Slot(11, valuesDouble, 0.0)
}
func RecValueAddValuesString(builder *flatbuffers.Builder, valuesString flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(valuesString), 0)
}
func RecValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}