// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfInt8 struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfInt8(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfInt8 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfInt8{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfInt8(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfInt8 {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfInt8{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfInt8) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfInt8) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayOfInt8) Value(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *ArrayOfInt8) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayOfInt8) MutateValue(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func ArrayOfInt8Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfInt8AddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func ArrayOfInt8StartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ArrayOfInt8End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}