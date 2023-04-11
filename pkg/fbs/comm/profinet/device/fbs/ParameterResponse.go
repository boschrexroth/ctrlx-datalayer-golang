// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ParameterResponseT struct {
	Data []byte
}

func (t *ParameterResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	ParameterResponseStart(builder)
	ParameterResponseAddData(builder, dataOffset)
	return ParameterResponseEnd(builder)
}

func (rcv *ParameterResponse) UnPackTo(t *ParameterResponseT) {
	t.Data = rcv.DataBytes()
}

func (rcv *ParameterResponse) UnPack() *ParameterResponseT {
	if rcv == nil { return nil }
	t := &ParameterResponseT{}
	rcv.UnPackTo(t)
	return t
}

type ParameterResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsParameterResponse(buf []byte, offset flatbuffers.UOffsetT) *ParameterResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ParameterResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsParameterResponse(buf []byte, offset flatbuffers.UOffsetT) *ParameterResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ParameterResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ParameterResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ParameterResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ParameterResponse) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *ParameterResponse) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ParameterResponse) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ParameterResponse) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func ParameterResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ParameterResponseAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(data), 0)
}
func ParameterResponseStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ParameterResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
