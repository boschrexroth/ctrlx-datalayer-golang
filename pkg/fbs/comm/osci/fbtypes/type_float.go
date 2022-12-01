// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type type_floatT struct {
	Value float32
}

func (t *type_floatT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	type_floatStart(builder)
	type_floatAddValue(builder, t.Value)
	return type_floatEnd(builder)
}

func (rcv *type_float) UnPackTo(t *type_floatT) {
	t.Value = rcv.Value()
}

func (rcv *type_float) UnPack() *type_floatT {
	if rcv == nil { return nil }
	t := &type_floatT{}
	rcv.UnPackTo(t)
	return t
}

type type_float struct {
	_tab flatbuffers.Table
}

func GetRootAstype_float(buf []byte, offset flatbuffers.UOffsetT) *type_float {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &type_float{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAstype_float(buf []byte, offset flatbuffers.UOffsetT) *type_float {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &type_float{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *type_float) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *type_float) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *type_float) Value() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *type_float) MutateValue(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func type_floatStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func type_floatAddValue(builder *flatbuffers.Builder, value float32) {
	builder.PrependFloat32Slot(0, value, 0.0)
}
func type_floatEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}