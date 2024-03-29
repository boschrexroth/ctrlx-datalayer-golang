// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfInt32T struct {
	Value []int32 `json:"value"`
}

func (t *ArrayOfInt32T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	valueOffset := flatbuffers.UOffsetT(0)
	if t.Value != nil {
		valueLength := len(t.Value)
		ArrayOfInt32StartValueVector(builder, valueLength)
		for j := valueLength - 1; j >= 0; j-- {
			builder.PrependInt32(t.Value[j])
		}
		valueOffset = builder.EndVector(valueLength)
	}
	ArrayOfInt32Start(builder)
	ArrayOfInt32AddValue(builder, valueOffset)
	return ArrayOfInt32End(builder)
}

func (rcv *ArrayOfInt32) UnPackTo(t *ArrayOfInt32T) {
	valueLength := rcv.ValueLength()
	t.Value = make([]int32, valueLength)
	for j := 0; j < valueLength; j++ {
		t.Value[j] = rcv.Value(j)
	}
}

func (rcv *ArrayOfInt32) UnPack() *ArrayOfInt32T {
	if rcv == nil { return nil }
	t := &ArrayOfInt32T{}
	rcv.UnPackTo(t)
	return t
}

type ArrayOfInt32 struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfInt32(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfInt32 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfInt32{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfInt32(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfInt32 {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfInt32{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfInt32) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfInt32) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayOfInt32) Value(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *ArrayOfInt32) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayOfInt32) MutateValue(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func ArrayOfInt32Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfInt32AddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func ArrayOfInt32StartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ArrayOfInt32End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
