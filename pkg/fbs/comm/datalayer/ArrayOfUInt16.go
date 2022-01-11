// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfUInt16 struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfUInt16(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfUInt16 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfUInt16{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfUInt16(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfUInt16 {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfUInt16{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfUInt16) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfUInt16) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayOfUInt16) Value(j int) uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint16(a + flatbuffers.UOffsetT(j*2))
	}
	return 0
}

func (rcv *ArrayOfUInt16) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayOfUInt16) MutateValue(j int, n uint16) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint16(a+flatbuffers.UOffsetT(j*2), n)
	}
	return false
}

func ArrayOfUInt16Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfUInt16AddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func ArrayOfUInt16StartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(2, numElems, 2)
}
func ArrayOfUInt16End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}