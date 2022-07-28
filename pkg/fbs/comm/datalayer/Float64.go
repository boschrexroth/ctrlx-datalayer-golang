// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Float64T struct {
	Value float64
}

func (t *Float64T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Float64Start(builder)
	Float64AddValue(builder, t.Value)
	return Float64End(builder)
}

func (rcv *Float64) UnPackTo(t *Float64T) {
	t.Value = rcv.Value()
}

func (rcv *Float64) UnPack() *Float64T {
	if rcv == nil { return nil }
	t := &Float64T{}
	rcv.UnPackTo(t)
	return t
}

type Float64 struct {
	_tab flatbuffers.Table
}

func GetRootAsFloat64(buf []byte, offset flatbuffers.UOffsetT) *Float64 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Float64{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFloat64(buf []byte, offset flatbuffers.UOffsetT) *Float64 {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Float64{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Float64) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Float64) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Float64) Value() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Float64) MutateValue(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

func Float64Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Float64AddValue(builder *flatbuffers.Builder, value float64) {
	builder.PrependFloat64Slot(0, value, 0.0)
}
func Float64End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
