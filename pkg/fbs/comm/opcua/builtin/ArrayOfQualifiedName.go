// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package builtin

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfQualifiedNameT struct {
	Array []*QualifiedNameT `json:"array"`
}

func (t *ArrayOfQualifiedNameT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	arrayOffset := flatbuffers.UOffsetT(0)
	if t.Array != nil {
		arrayLength := len(t.Array)
		arrayOffsets := make([]flatbuffers.UOffsetT, arrayLength)
		for j := 0; j < arrayLength; j++ {
			arrayOffsets[j] = t.Array[j].Pack(builder)
		}
		ArrayOfQualifiedNameStartArrayVector(builder, arrayLength)
		for j := arrayLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(arrayOffsets[j])
		}
		arrayOffset = builder.EndVector(arrayLength)
	}
	ArrayOfQualifiedNameStart(builder)
	ArrayOfQualifiedNameAddArray(builder, arrayOffset)
	return ArrayOfQualifiedNameEnd(builder)
}

func (rcv *ArrayOfQualifiedName) UnPackTo(t *ArrayOfQualifiedNameT) {
	arrayLength := rcv.ArrayLength()
	t.Array = make([]*QualifiedNameT, arrayLength)
	for j := 0; j < arrayLength; j++ {
		x := QualifiedName{}
		rcv.Array(&x, j)
		t.Array[j] = x.UnPack()
	}
}

func (rcv *ArrayOfQualifiedName) UnPack() *ArrayOfQualifiedNameT {
	if rcv == nil { return nil }
	t := &ArrayOfQualifiedNameT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayOfQualifiedName struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfQualifiedName(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfQualifiedName {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfQualifiedName{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfQualifiedName(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfQualifiedName {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfQualifiedName{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfQualifiedName) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfQualifiedName) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayOfQualifiedName) Array(obj *QualifiedName, j int) bool {
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

func (rcv *ArrayOfQualifiedName) ArrayLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ArrayOfQualifiedNameStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfQualifiedNameAddArray(builder *flatbuffers.Builder, array flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(array), 0)
}
func ArrayOfQualifiedNameStartArrayVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ArrayOfQualifiedNameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
