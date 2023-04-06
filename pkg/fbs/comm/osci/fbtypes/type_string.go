// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type type_stringT struct {
	Value string
}

func (t *type_stringT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	valueOffset := builder.CreateString(t.Value)
	type_stringStart(builder)
	type_stringAddValue(builder, valueOffset)
	return type_stringEnd(builder)
}

func (rcv *type_string) UnPackTo(t *type_stringT) {
	t.Value = string(rcv.Value())
}

func (rcv *type_string) UnPack() *type_stringT {
	if rcv == nil { return nil }
	t := &type_stringT{}
	rcv.UnPackTo(t)
	return t
}

type type_string struct {
	_tab flatbuffers.Table
}

func GetRootAstype_string(buf []byte, offset flatbuffers.UOffsetT) *type_string {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &type_string{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAstype_string(buf []byte, offset flatbuffers.UOffsetT) *type_string {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &type_string{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *type_string) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *type_string) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *type_string) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func type_stringStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func type_stringAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func type_stringEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
