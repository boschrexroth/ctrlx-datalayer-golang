// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Slave register response
type SlaveRegisterResponseT struct {
	Data []byte `json:"data"`
}

func (t *SlaveRegisterResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	SlaveRegisterResponseStart(builder)
	SlaveRegisterResponseAddData(builder, dataOffset)
	return SlaveRegisterResponseEnd(builder)
}

func (rcv *SlaveRegisterResponse) UnPackTo(t *SlaveRegisterResponseT) {
	t.Data = rcv.DataBytes()
}

func (rcv *SlaveRegisterResponse) UnPack() *SlaveRegisterResponseT {
	if rcv == nil { return nil }
	t := &SlaveRegisterResponseT{}
	rcv.UnPackTo(t)
	return t
}

type SlaveRegisterResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsSlaveRegisterResponse(buf []byte, offset flatbuffers.UOffsetT) *SlaveRegisterResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SlaveRegisterResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSlaveRegisterResponse(buf []byte, offset flatbuffers.UOffsetT) *SlaveRegisterResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SlaveRegisterResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SlaveRegisterResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SlaveRegisterResponse) Table() flatbuffers.Table {
	return rcv._tab
}

///Data buffer
func (rcv *SlaveRegisterResponse) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *SlaveRegisterResponse) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SlaveRegisterResponse) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///Data buffer
func (rcv *SlaveRegisterResponse) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func SlaveRegisterResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SlaveRegisterResponseAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(data), 0)
}
func SlaveRegisterResponseStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func SlaveRegisterResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
