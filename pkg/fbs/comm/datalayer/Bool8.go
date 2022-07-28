// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Bool8T struct {
	Value bool
}

func (t *Bool8T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Bool8Start(builder)
	Bool8AddValue(builder, t.Value)
	return Bool8End(builder)
}

func (rcv *Bool8) UnPackTo(t *Bool8T) {
	t.Value = rcv.Value()
}

func (rcv *Bool8) UnPack() *Bool8T {
	if rcv == nil { return nil }
	t := &Bool8T{}
	rcv.UnPackTo(t)
	return t
}

type Bool8 struct {
	_tab flatbuffers.Table
}

func GetRootAsBool8(buf []byte, offset flatbuffers.UOffsetT) *Bool8 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Bool8{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBool8(buf []byte, offset flatbuffers.UOffsetT) *Bool8 {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Bool8{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Bool8) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Bool8) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Bool8) Value() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Bool8) MutateValue(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func Bool8Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Bool8AddValue(builder *flatbuffers.Builder, value bool) {
	builder.PrependBoolSlot(0, value, false)
}
func Bool8End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
