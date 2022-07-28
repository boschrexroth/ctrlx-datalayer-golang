// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StringT struct {
	Value string
}

func (t *StringT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	valueOffset := builder.CreateString(t.Value)
	StringStart(builder)
	StringAddValue(builder, valueOffset)
	return StringEnd(builder)
}

func (rcv *String) UnPackTo(t *StringT) {
	t.Value = string(rcv.Value())
}

func (rcv *String) UnPack() *StringT {
	if rcv == nil { return nil }
	t := &StringT{}
	rcv.UnPackTo(t)
	return t
}

type String struct {
	_tab flatbuffers.Table
}

func GetRootAsString(buf []byte, offset flatbuffers.UOffsetT) *String {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &String{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsString(buf []byte, offset flatbuffers.UOffsetT) *String {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &String{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *String) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *String) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *String) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func StringStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func StringAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func StringEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
