// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type type_ulongT struct {
	Value uint64 `json:"value"`
}

func (t *type_ulongT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	type_ulongStart(builder)
	type_ulongAddValue(builder, t.Value)
	return type_ulongEnd(builder)
}

func (rcv *type_ulong) UnPackTo(t *type_ulongT) {
	t.Value = rcv.Value()
}

func (rcv *type_ulong) UnPack() *type_ulongT {
	if rcv == nil { return nil }
	t := &type_ulongT{}
	rcv.UnPackTo(t)
	return t
}

type type_ulong struct {
	_tab flatbuffers.Table
}

func GetRootAstype_ulong(buf []byte, offset flatbuffers.UOffsetT) *type_ulong {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &type_ulong{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAstype_ulong(buf []byte, offset flatbuffers.UOffsetT) *type_ulong {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &type_ulong{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *type_ulong) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *type_ulong) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *type_ulong) Value() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *type_ulong) MutateValue(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func type_ulongStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func type_ulongAddValue(builder *flatbuffers.Builder, value uint64) {
	builder.PrependUint64Slot(0, value, 0)
}
func type_ulongEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
