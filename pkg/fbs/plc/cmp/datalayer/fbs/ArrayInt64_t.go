// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayInt64_tT struct {
	Data []int64
}

func (t *ArrayInt64_tT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataLength := len(t.Data)
		ArrayInt64_tStartDataVector(builder, dataLength)
		for j := dataLength - 1; j >= 0; j-- {
			builder.PrependInt64(t.Data[j])
		}
		dataOffset = builder.EndVector(dataLength)
	}
	ArrayInt64_tStart(builder)
	ArrayInt64_tAddData(builder, dataOffset)
	return ArrayInt64_tEnd(builder)
}

func (rcv *ArrayInt64_t) UnPackTo(t *ArrayInt64_tT) {
	dataLength := rcv.DataLength()
	t.Data = make([]int64, dataLength)
	for j := 0; j < dataLength; j++ {
		t.Data[j] = rcv.Data(j)
	}
}

func (rcv *ArrayInt64_t) UnPack() *ArrayInt64_tT {
	if rcv == nil { return nil }
	t := &ArrayInt64_tT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayInt64_t struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayInt64_t(buf []byte, offset flatbuffers.UOffsetT) *ArrayInt64_t {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayInt64_t{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayInt64_t(buf []byte, offset flatbuffers.UOffsetT) *ArrayInt64_t {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayInt64_t{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayInt64_t) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayInt64_t) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayInt64_t) Data(j int) int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *ArrayInt64_t) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayInt64_t) MutateData(j int, n int64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func ArrayInt64_tStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayInt64_tAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(data), 0)
}
func ArrayInt64_tStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func ArrayInt64_tEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
