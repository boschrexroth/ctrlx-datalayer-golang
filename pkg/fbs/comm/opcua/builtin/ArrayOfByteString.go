// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package builtin

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfByteStringT struct {
	Array []*ByteStringT `json:"array"`
}

func (t *ArrayOfByteStringT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	arrayOffset := flatbuffers.UOffsetT(0)
	if t.Array != nil {
		arrayLength := len(t.Array)
		arrayOffsets := make([]flatbuffers.UOffsetT, arrayLength)
		for j := 0; j < arrayLength; j++ {
			arrayOffsets[j] = t.Array[j].Pack(builder)
		}
		ArrayOfByteStringStartArrayVector(builder, arrayLength)
		for j := arrayLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(arrayOffsets[j])
		}
		arrayOffset = builder.EndVector(arrayLength)
	}
	ArrayOfByteStringStart(builder)
	ArrayOfByteStringAddArray(builder, arrayOffset)
	return ArrayOfByteStringEnd(builder)
}

func (rcv *ArrayOfByteString) UnPackTo(t *ArrayOfByteStringT) {
	arrayLength := rcv.ArrayLength()
	t.Array = make([]*ByteStringT, arrayLength)
	for j := 0; j < arrayLength; j++ {
		x := ByteString{}
		rcv.Array(&x, j)
		t.Array[j] = x.UnPack()
	}
}

func (rcv *ArrayOfByteString) UnPack() *ArrayOfByteStringT {
	if rcv == nil { return nil }
	t := &ArrayOfByteStringT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayOfByteString struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfByteString(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfByteString {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfByteString{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfByteString(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfByteString {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfByteString{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfByteString) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfByteString) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayOfByteString) Array(obj *ByteString, j int) bool {
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

func (rcv *ArrayOfByteString) ArrayLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ArrayOfByteStringStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfByteStringAddArray(builder *flatbuffers.Builder, array flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(array), 0)
}
func ArrayOfByteStringStartArrayVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ArrayOfByteStringEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
