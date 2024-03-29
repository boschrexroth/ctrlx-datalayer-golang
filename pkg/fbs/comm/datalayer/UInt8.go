// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UInt8T struct {
	Value byte `json:"value"`
}

func (t *UInt8T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	UInt8Start(builder)
	UInt8AddValue(builder, t.Value)
	return UInt8End(builder)
}

func (rcv *UInt8) UnPackTo(t *UInt8T) {
	t.Value = rcv.Value()
}

func (rcv *UInt8) UnPack() *UInt8T {
	if rcv == nil { return nil }
	t := &UInt8T{}
	rcv.UnPackTo(t)
	return t
}

type UInt8 struct {
	_tab flatbuffers.Table
}

func GetRootAsUInt8(buf []byte, offset flatbuffers.UOffsetT) *UInt8 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UInt8{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUInt8(buf []byte, offset flatbuffers.UOffsetT) *UInt8 {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &UInt8{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *UInt8) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UInt8) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *UInt8) Value() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *UInt8) MutateValue(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func UInt8Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func UInt8AddValue(builder *flatbuffers.Builder, value byte) {
	builder.PrependByteSlot(0, value, 0)
}
func UInt8End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
