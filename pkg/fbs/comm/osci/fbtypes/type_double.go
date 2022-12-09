// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type type_doubleT struct {
	Value float64
}

func (t *type_doubleT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	type_doubleStart(builder)
	type_doubleAddValue(builder, t.Value)
	return type_doubleEnd(builder)
}

func (rcv *type_double) UnPackTo(t *type_doubleT) {
	t.Value = rcv.Value()
}

func (rcv *type_double) UnPack() *type_doubleT {
	if rcv == nil { return nil }
	t := &type_doubleT{}
	rcv.UnPackTo(t)
	return t
}

type type_double struct {
	_tab flatbuffers.Table
}

func GetRootAstype_double(buf []byte, offset flatbuffers.UOffsetT) *type_double {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &type_double{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAstype_double(buf []byte, offset flatbuffers.UOffsetT) *type_double {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &type_double{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *type_double) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *type_double) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *type_double) Value() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *type_double) MutateValue(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

func type_doubleStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func type_doubleAddValue(builder *flatbuffers.Builder, value float64) {
	builder.PrependFloat64Slot(0, value, 0.0)
}
func type_doubleEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
