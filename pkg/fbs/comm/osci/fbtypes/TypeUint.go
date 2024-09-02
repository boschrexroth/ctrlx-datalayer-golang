// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TypeUintT struct {
	Value uint32 `json:"value"`
}

func (t *TypeUintT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	TypeUintStart(builder)
	TypeUintAddValue(builder, t.Value)
	return TypeUintEnd(builder)
}

func (rcv *TypeUint) UnPackTo(t *TypeUintT) {
	t.Value = rcv.Value()
}

func (rcv *TypeUint) UnPack() *TypeUintT {
	if rcv == nil { return nil }
	t := &TypeUintT{}
	rcv.UnPackTo(t)
	return t
}

type TypeUint struct {
	_tab flatbuffers.Table
}

func GetRootAsTypeUint(buf []byte, offset flatbuffers.UOffsetT) *TypeUint {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TypeUint{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTypeUint(buf []byte, offset flatbuffers.UOffsetT) *TypeUint {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TypeUint{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TypeUint) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TypeUint) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TypeUint) Value() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TypeUint) MutateValue(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func TypeUintStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TypeUintAddValue(builder *flatbuffers.Builder, value uint32) {
	builder.PrependUint32Slot(0, value, 0)
}
func TypeUintEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}