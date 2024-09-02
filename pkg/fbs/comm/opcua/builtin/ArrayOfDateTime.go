// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package builtin

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfDateTimeT struct {
	Array []*DateTimeT `json:"array"`
}

func (t *ArrayOfDateTimeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	arrayOffset := flatbuffers.UOffsetT(0)
	if t.Array != nil {
		arrayLength := len(t.Array)
		arrayOffsets := make([]flatbuffers.UOffsetT, arrayLength)
		for j := 0; j < arrayLength; j++ {
			arrayOffsets[j] = t.Array[j].Pack(builder)
		}
		ArrayOfDateTimeStartArrayVector(builder, arrayLength)
		for j := arrayLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(arrayOffsets[j])
		}
		arrayOffset = builder.EndVector(arrayLength)
	}
	ArrayOfDateTimeStart(builder)
	ArrayOfDateTimeAddArray(builder, arrayOffset)
	return ArrayOfDateTimeEnd(builder)
}

func (rcv *ArrayOfDateTime) UnPackTo(t *ArrayOfDateTimeT) {
	arrayLength := rcv.ArrayLength()
	t.Array = make([]*DateTimeT, arrayLength)
	for j := 0; j < arrayLength; j++ {
		x := DateTime{}
		rcv.Array(&x, j)
		t.Array[j] = x.UnPack()
	}
}

func (rcv *ArrayOfDateTime) UnPack() *ArrayOfDateTimeT {
	if rcv == nil { return nil }
	t := &ArrayOfDateTimeT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayOfDateTime struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfDateTime(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfDateTime {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfDateTime{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfDateTime(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfDateTime {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfDateTime{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfDateTime) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfDateTime) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayOfDateTime) Array(obj *DateTime, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ArrayOfDateTime) ArrayLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ArrayOfDateTimeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfDateTimeAddArray(builder *flatbuffers.Builder, array flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(array), 0)
}
func ArrayOfDateTimeStartArrayVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ArrayOfDateTimeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}