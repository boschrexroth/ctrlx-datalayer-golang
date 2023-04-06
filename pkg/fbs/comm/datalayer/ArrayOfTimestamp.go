// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfTimestampT struct {
	Value []uint64
}

func (t *ArrayOfTimestampT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	valueOffset := flatbuffers.UOffsetT(0)
	if t.Value != nil {
		valueLength := len(t.Value)
		ArrayOfTimestampStartValueVector(builder, valueLength)
		for j := valueLength - 1; j >= 0; j-- {
			builder.PrependUint64(t.Value[j])
		}
		valueOffset = builder.EndVector(valueLength)
	}
	ArrayOfTimestampStart(builder)
	ArrayOfTimestampAddValue(builder, valueOffset)
	return ArrayOfTimestampEnd(builder)
}

func (rcv *ArrayOfTimestamp) UnPackTo(t *ArrayOfTimestampT) {
	valueLength := rcv.ValueLength()
	t.Value = make([]uint64, valueLength)
	for j := 0; j < valueLength; j++ {
		t.Value[j] = rcv.Value(j)
	}
}

func (rcv *ArrayOfTimestamp) UnPack() *ArrayOfTimestampT {
	if rcv == nil { return nil }
	t := &ArrayOfTimestampT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayOfTimestamp struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfTimestamp(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfTimestamp {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfTimestamp{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfTimestamp(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfTimestamp {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfTimestamp{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfTimestamp) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfTimestamp) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayOfTimestamp) Value(j int) uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *ArrayOfTimestamp) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayOfTimestamp) MutateValue(j int, n uint64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func ArrayOfTimestampStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfTimestampAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func ArrayOfTimestampStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func ArrayOfTimestampEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
