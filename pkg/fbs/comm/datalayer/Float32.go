// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Float32T struct {
	Value float32
}

func (t *Float32T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Float32Start(builder)
	Float32AddValue(builder, t.Value)
	return Float32End(builder)
}

func (rcv *Float32) UnPackTo(t *Float32T) {
	t.Value = rcv.Value()
}

func (rcv *Float32) UnPack() *Float32T {
	if rcv == nil { return nil }
	t := &Float32T{}
	rcv.UnPackTo(t)
	return t
}

type Float32 struct {
	_tab flatbuffers.Table
}

func GetRootAsFloat32(buf []byte, offset flatbuffers.UOffsetT) *Float32 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Float32{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFloat32(buf []byte, offset flatbuffers.UOffsetT) *Float32 {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Float32{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Float32) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Float32) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Float32) Value() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Float32) MutateValue(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func Float32Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Float32AddValue(builder *flatbuffers.Builder, value float32) {
	builder.PrependFloat32Slot(0, value, 0.0)
}
func Float32End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
