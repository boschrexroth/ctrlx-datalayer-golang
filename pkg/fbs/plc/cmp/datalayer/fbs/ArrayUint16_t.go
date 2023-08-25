// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayUint16_tT struct {
	Data []uint16 `json:"data"`
}

func (t *ArrayUint16_tT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataLength := len(t.Data)
		ArrayUint16_tStartDataVector(builder, dataLength)
		for j := dataLength - 1; j >= 0; j-- {
			builder.PrependUint16(t.Data[j])
		}
		dataOffset = builder.EndVector(dataLength)
	}
	ArrayUint16_tStart(builder)
	ArrayUint16_tAddData(builder, dataOffset)
	return ArrayUint16_tEnd(builder)
}

func (rcv *ArrayUint16_t) UnPackTo(t *ArrayUint16_tT) {
	dataLength := rcv.DataLength()
	t.Data = make([]uint16, dataLength)
	for j := 0; j < dataLength; j++ {
		t.Data[j] = rcv.Data(j)
	}
}

func (rcv *ArrayUint16_t) UnPack() *ArrayUint16_tT {
	if rcv == nil { return nil }
	t := &ArrayUint16_tT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayUint16_t struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayUint16_t(buf []byte, offset flatbuffers.UOffsetT) *ArrayUint16_t {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayUint16_t{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayUint16_t(buf []byte, offset flatbuffers.UOffsetT) *ArrayUint16_t {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayUint16_t{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayUint16_t) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayUint16_t) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayUint16_t) Data(j int) uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint16(a + flatbuffers.UOffsetT(j*2))
	}
	return 0
}

func (rcv *ArrayUint16_t) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayUint16_t) MutateData(j int, n uint16) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint16(a+flatbuffers.UOffsetT(j*2), n)
	}
	return false
}

func ArrayUint16_tStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayUint16_tAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(data), 0)
}
func ArrayUint16_tStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(2, numElems, 2)
}
func ArrayUint16_tEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
