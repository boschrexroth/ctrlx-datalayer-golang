// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayBool_tT struct {
	Data []bool `json:"data"`
}

func (t *ArrayBool_tT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataLength := len(t.Data)
		ArrayBool_tStartDataVector(builder, dataLength)
		for j := dataLength - 1; j >= 0; j-- {
			builder.PrependBool(t.Data[j])
		}
		dataOffset = builder.EndVector(dataLength)
	}
	ArrayBool_tStart(builder)
	ArrayBool_tAddData(builder, dataOffset)
	return ArrayBool_tEnd(builder)
}

func (rcv *ArrayBool_t) UnPackTo(t *ArrayBool_tT) {
	dataLength := rcv.DataLength()
	t.Data = make([]bool, dataLength)
	for j := 0; j < dataLength; j++ {
		t.Data[j] = rcv.Data(j)
	}
}

func (rcv *ArrayBool_t) UnPack() *ArrayBool_tT {
	if rcv == nil { return nil }
	t := &ArrayBool_tT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayBool_t struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayBool_t(buf []byte, offset flatbuffers.UOffsetT) *ArrayBool_t {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayBool_t{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayBool_t(buf []byte, offset flatbuffers.UOffsetT) *ArrayBool_t {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayBool_t{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayBool_t) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayBool_t) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayBool_t) Data(j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetBool(a + flatbuffers.UOffsetT(j*1))
	}
	return false
}

func (rcv *ArrayBool_t) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ArrayBool_t) MutateData(j int, n bool) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateBool(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func ArrayBool_tStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayBool_tAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(data), 0)
}
func ArrayBool_tStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ArrayBool_tEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
