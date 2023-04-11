// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type type_ubyteT struct {
	Value byte
}

func (t *type_ubyteT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	type_ubyteStart(builder)
	type_ubyteAddValue(builder, t.Value)
	return type_ubyteEnd(builder)
}

func (rcv *type_ubyte) UnPackTo(t *type_ubyteT) {
	t.Value = rcv.Value()
}

func (rcv *type_ubyte) UnPack() *type_ubyteT {
	if rcv == nil { return nil }
	t := &type_ubyteT{}
	rcv.UnPackTo(t)
	return t
}

type type_ubyte struct {
	_tab flatbuffers.Table
}

func GetRootAstype_ubyte(buf []byte, offset flatbuffers.UOffsetT) *type_ubyte {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &type_ubyte{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAstype_ubyte(buf []byte, offset flatbuffers.UOffsetT) *type_ubyte {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &type_ubyte{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *type_ubyte) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *type_ubyte) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *type_ubyte) Value() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *type_ubyte) MutateValue(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func type_ubyteStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func type_ubyteAddValue(builder *flatbuffers.Builder, value byte) {
	builder.PrependByteSlot(0, value, 0)
}
func type_ubyteEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
