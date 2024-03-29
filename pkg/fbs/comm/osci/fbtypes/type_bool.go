// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type type_boolT struct {
	Value bool `json:"value"`
}

func (t *type_boolT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	type_boolStart(builder)
	type_boolAddValue(builder, t.Value)
	return type_boolEnd(builder)
}

func (rcv *type_bool) UnPackTo(t *type_boolT) {
	t.Value = rcv.Value()
}

func (rcv *type_bool) UnPack() *type_boolT {
	if rcv == nil { return nil }
	t := &type_boolT{}
	rcv.UnPackTo(t)
	return t
}

type type_bool struct {
	_tab flatbuffers.Table
}

func GetRootAstype_bool(buf []byte, offset flatbuffers.UOffsetT) *type_bool {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &type_bool{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAstype_bool(buf []byte, offset flatbuffers.UOffsetT) *type_bool {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &type_bool{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *type_bool) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *type_bool) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *type_bool) Value() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *type_bool) MutateValue(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func type_boolStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func type_boolAddValue(builder *flatbuffers.Builder, value bool) {
	builder.PrependBoolSlot(0, value, false)
}
func type_boolEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
