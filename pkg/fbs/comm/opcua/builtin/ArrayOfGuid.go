// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package builtin

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfGuidT struct {
	Array []*GuidT `json:"array"`
}

func (t *ArrayOfGuidT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	arrayOffset := flatbuffers.UOffsetT(0)
	if t.Array != nil {
		arrayLength := len(t.Array)
		arrayOffsets := make([]flatbuffers.UOffsetT, arrayLength)
		for j := 0; j < arrayLength; j++ {
			arrayOffsets[j] = t.Array[j].Pack(builder)
		}
		ArrayOfGuidStartArrayVector(builder, arrayLength)
		for j := arrayLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(arrayOffsets[j])
		}
		arrayOffset = builder.EndVector(arrayLength)
	}
	ArrayOfGuidStart(builder)
	ArrayOfGuidAddArray(builder, arrayOffset)
	return ArrayOfGuidEnd(builder)
}

func (rcv *ArrayOfGuid) UnPackTo(t *ArrayOfGuidT) {
	arrayLength := rcv.ArrayLength()
	t.Array = make([]*GuidT, arrayLength)
	for j := 0; j < arrayLength; j++ {
		x := Guid{}
		rcv.Array(&x, j)
		t.Array[j] = x.UnPack()
	}
}

func (rcv *ArrayOfGuid) UnPack() *ArrayOfGuidT {
	if rcv == nil { return nil }
	t := &ArrayOfGuidT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayOfGuid struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfGuid(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfGuid {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfGuid{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfGuid(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfGuid {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfGuid{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfGuid) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfGuid) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayOfGuid) Array(obj *Guid, j int) bool {
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

func (rcv *ArrayOfGuid) ArrayLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ArrayOfGuidStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfGuidAddArray(builder *flatbuffers.Builder, array flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(array), 0)
}
func ArrayOfGuidStartArrayVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ArrayOfGuidEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
