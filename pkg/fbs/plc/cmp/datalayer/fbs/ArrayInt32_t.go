// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayInt32_tT struct {
	Data []int32
}

func (t *ArrayInt32_tT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataLength := len(t.Data)
		ArrayInt32_tStartDataVector(builder, dataLength)
		for j := dataLength - 1; j >= 0; j-- {
			builder.PrependInt32(t.Data[j])
		}
		dataOffset = builder.EndVector(dataLength)
	}
	ArrayInt32_tStart(builder)
	ArrayInt32_tAddData(builder, dataOffset)
	return ArrayInt32_tEnd(builder)
}

func (rcv *ArrayInt32_t) UnPackTo(t *ArrayInt32_tT) {
	dataLength := rcv.DataLength()
	t.Data = make([]int32, dataLength)
	for j := 0; j < dataLength; j++ {
		t.Data[j] = rcv.Data(j)
	}
}

func (rcv *ArrayInt32_t) UnPack() *ArrayInt32_tT {
	if rcv == nil { return nil }
	t := &ArrayInt32_tT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayInt32_t struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayInt32_t(buf []byte, offset flatbuffers.UOffsetT) *ArrayInt32_t {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayInt32_t{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayInt32_t(buf []byte, offset flatbuffers.UOffsetT) *ArrayInt32_t {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayInt32_t{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayInt32_t) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayInt32_t) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayInt32_t) Data(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *ArrayInt32_t) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayInt32_t) MutateData(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func ArrayInt32_tStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayInt32_tAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(data), 0)
}
func ArrayInt32_tStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ArrayInt32_tEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
