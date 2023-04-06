// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfFloat32T struct {
	Value []float32
}

func (t *ArrayOfFloat32T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	valueOffset := flatbuffers.UOffsetT(0)
	if t.Value != nil {
		valueLength := len(t.Value)
		ArrayOfFloat32StartValueVector(builder, valueLength)
		for j := valueLength - 1; j >= 0; j-- {
			builder.PrependFloat32(t.Value[j])
		}
		valueOffset = builder.EndVector(valueLength)
	}
	ArrayOfFloat32Start(builder)
	ArrayOfFloat32AddValue(builder, valueOffset)
	return ArrayOfFloat32End(builder)
}

func (rcv *ArrayOfFloat32) UnPackTo(t *ArrayOfFloat32T) {
	valueLength := rcv.ValueLength()
	t.Value = make([]float32, valueLength)
	for j := 0; j < valueLength; j++ {
		t.Value[j] = rcv.Value(j)
	}
}

func (rcv *ArrayOfFloat32) UnPack() *ArrayOfFloat32T {
	if rcv == nil { return nil }
	t := &ArrayOfFloat32T{}
	rcv.UnPackTo(t)
	return t
}

type ArrayOfFloat32 struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfFloat32(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfFloat32 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfFloat32{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfFloat32(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfFloat32 {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfFloat32{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfFloat32) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfFloat32) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayOfFloat32) Value(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *ArrayOfFloat32) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayOfFloat32) MutateValue(j int, n float32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func ArrayOfFloat32Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfFloat32AddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func ArrayOfFloat32StartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ArrayOfFloat32End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
