// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TypeUlongT struct {
	Value uint64 `json:"value"`
}

func (t *TypeUlongT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	TypeUlongStart(builder)
	TypeUlongAddValue(builder, t.Value)
	return TypeUlongEnd(builder)
}

func (rcv *TypeUlong) UnPackTo(t *TypeUlongT) {
	t.Value = rcv.Value()
}

func (rcv *TypeUlong) UnPack() *TypeUlongT {
	if rcv == nil { return nil }
	t := &TypeUlongT{}
	rcv.UnPackTo(t)
	return t
}

type TypeUlong struct {
	_tab flatbuffers.Table
}

func GetRootAsTypeUlong(buf []byte, offset flatbuffers.UOffsetT) *TypeUlong {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TypeUlong{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTypeUlong(buf []byte, offset flatbuffers.UOffsetT) *TypeUlong {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TypeUlong{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TypeUlong) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TypeUlong) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TypeUlong) Value() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TypeUlong) MutateValue(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func TypeUlongStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TypeUlongAddValue(builder *flatbuffers.Builder, value uint64) {
	builder.PrependUint64Slot(0, value, 0)
}
func TypeUlongEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
