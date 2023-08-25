// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayFloat64_tT struct {
	Data []float64 `json:"data"`
}

func (t *ArrayFloat64_tT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataLength := len(t.Data)
		ArrayFloat64_tStartDataVector(builder, dataLength)
		for j := dataLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.Data[j])
		}
		dataOffset = builder.EndVector(dataLength)
	}
	ArrayFloat64_tStart(builder)
	ArrayFloat64_tAddData(builder, dataOffset)
	return ArrayFloat64_tEnd(builder)
}

func (rcv *ArrayFloat64_t) UnPackTo(t *ArrayFloat64_tT) {
	dataLength := rcv.DataLength()
	t.Data = make([]float64, dataLength)
	for j := 0; j < dataLength; j++ {
		t.Data[j] = rcv.Data(j)
	}
}

func (rcv *ArrayFloat64_t) UnPack() *ArrayFloat64_tT {
	if rcv == nil { return nil }
	t := &ArrayFloat64_tT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayFloat64_t struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayFloat64_t(buf []byte, offset flatbuffers.UOffsetT) *ArrayFloat64_t {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayFloat64_t{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayFloat64_t(buf []byte, offset flatbuffers.UOffsetT) *ArrayFloat64_t {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayFloat64_t{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayFloat64_t) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayFloat64_t) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayFloat64_t) Data(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *ArrayFloat64_t) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayFloat64_t) MutateData(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func ArrayFloat64_tStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayFloat64_tAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(data), 0)
}
func ArrayFloat64_tStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func ArrayFloat64_tEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
