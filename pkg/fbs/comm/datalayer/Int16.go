// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Int16T struct {
	Value int16 `json:"value"`
}

func (t *Int16T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Int16Start(builder)
	Int16AddValue(builder, t.Value)
	return Int16End(builder)
}

func (rcv *Int16) UnPackTo(t *Int16T) {
	t.Value = rcv.Value()
}

func (rcv *Int16) UnPack() *Int16T {
	if rcv == nil { return nil }
	t := &Int16T{}
	rcv.UnPackTo(t)
	return t
}

type Int16 struct {
	_tab flatbuffers.Table
}

func GetRootAsInt16(buf []byte, offset flatbuffers.UOffsetT) *Int16 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Int16{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInt16(buf []byte, offset flatbuffers.UOffsetT) *Int16 {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Int16{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Int16) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Int16) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Int16) Value() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Int16) MutateValue(n int16) bool {
	return rcv._tab.MutateInt16Slot(4, n)
}

func Int16Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Int16AddValue(builder *flatbuffers.Builder, value int16) {
	builder.PrependInt16Slot(0, value, 0)
}
func Int16End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
