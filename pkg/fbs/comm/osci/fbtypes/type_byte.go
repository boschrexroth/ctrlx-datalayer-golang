// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type type_byteT struct {
	Value int8
}

func (t *type_byteT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	type_byteStart(builder)
	type_byteAddValue(builder, t.Value)
	return type_byteEnd(builder)
}

func (rcv *type_byte) UnPackTo(t *type_byteT) {
	t.Value = rcv.Value()
}

func (rcv *type_byte) UnPack() *type_byteT {
	if rcv == nil { return nil }
	t := &type_byteT{}
	rcv.UnPackTo(t)
	return t
}

type type_byte struct {
	_tab flatbuffers.Table
}

func GetRootAstype_byte(buf []byte, offset flatbuffers.UOffsetT) *type_byte {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &type_byte{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAstype_byte(buf []byte, offset flatbuffers.UOffsetT) *type_byte {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &type_byte{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *type_byte) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *type_byte) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *type_byte) Value() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *type_byte) MutateValue(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func type_byteStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func type_byteAddValue(builder *flatbuffers.Builder, value int8) {
	builder.PrependInt8Slot(0, value, 0)
}
func type_byteEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
