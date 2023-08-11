// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type type_uintT struct {
	Value uint32 `json:"value"`
}

func (t *type_uintT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	type_uintStart(builder)
	type_uintAddValue(builder, t.Value)
	return type_uintEnd(builder)
}

func (rcv *type_uint) UnPackTo(t *type_uintT) {
	t.Value = rcv.Value()
}

func (rcv *type_uint) UnPack() *type_uintT {
	if rcv == nil { return nil }
	t := &type_uintT{}
	rcv.UnPackTo(t)
	return t
}

type type_uint struct {
	_tab flatbuffers.Table
}

func GetRootAstype_uint(buf []byte, offset flatbuffers.UOffsetT) *type_uint {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &type_uint{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAstype_uint(buf []byte, offset flatbuffers.UOffsetT) *type_uint {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &type_uint{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *type_uint) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *type_uint) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *type_uint) Value() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *type_uint) MutateValue(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func type_uintStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func type_uintAddValue(builder *flatbuffers.Builder, value uint32) {
	builder.PrependUint32Slot(0, value, 0)
}
func type_uintEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
