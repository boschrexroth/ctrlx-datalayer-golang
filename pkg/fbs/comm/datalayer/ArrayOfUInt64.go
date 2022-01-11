// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfUInt64 struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfUInt64(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfUInt64 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfUInt64{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfUInt64(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfUInt64 {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfUInt64{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfUInt64) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfUInt64) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayOfUInt64) Value(j int) uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *ArrayOfUInt64) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayOfUInt64) MutateValue(j int, n uint64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func ArrayOfUInt64Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfUInt64AddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func ArrayOfUInt64StartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func ArrayOfUInt64End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}