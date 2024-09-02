// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package builtin

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ByteStringT struct {
	Value []int8 `json:"value"`
}

func (t *ByteStringT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	valueOffset := flatbuffers.UOffsetT(0)
	if t.Value != nil {
		valueLength := len(t.Value)
		ByteStringStartValueVector(builder, valueLength)
		for j := valueLength - 1; j >= 0; j-- {
			builder.PrependInt8(t.Value[j])
		}
		valueOffset = builder.EndVector(valueLength)
	}
	ByteStringStart(builder)
	ByteStringAddValue(builder, valueOffset)
	return ByteStringEnd(builder)
}

func (rcv *ByteString) UnPackTo(t *ByteStringT) {
	valueLength := rcv.ValueLength()
	t.Value = make([]int8, valueLength)
	for j := 0; j < valueLength; j++ {
		t.Value[j] = rcv.Value(j)
	}
}

func (rcv *ByteString) UnPack() *ByteStringT {
	if rcv == nil { return nil }
	t := &ByteStringT{}
	rcv.UnPackTo(t)
	return t
}

type ByteString struct {
	_tab flatbuffers.Table
}

func GetRootAsByteString(buf []byte, offset flatbuffers.UOffsetT) *ByteString {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ByteString{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsByteString(buf []byte, offset flatbuffers.UOffsetT) *ByteString {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ByteString{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ByteString) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ByteString) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ByteString) Value(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *ByteString) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ByteString) MutateValue(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func ByteStringStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ByteStringAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func ByteStringStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ByteStringEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
