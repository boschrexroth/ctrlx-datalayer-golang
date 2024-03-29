// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Response to AoE data Request
type AoeDataResponseT struct {
	Data []byte `json:"data"`
	ErrorCode uint32 `json:"errorCode"`
	CmdResult uint32 `json:"cmdResult"`
}

func (t *AoeDataResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	AoeDataResponseStart(builder)
	AoeDataResponseAddData(builder, dataOffset)
	AoeDataResponseAddErrorCode(builder, t.ErrorCode)
	AoeDataResponseAddCmdResult(builder, t.CmdResult)
	return AoeDataResponseEnd(builder)
}

func (rcv *AoeDataResponse) UnPackTo(t *AoeDataResponseT) {
	t.Data = rcv.DataBytes()
	t.ErrorCode = rcv.ErrorCode()
	t.CmdResult = rcv.CmdResult()
}

func (rcv *AoeDataResponse) UnPack() *AoeDataResponseT {
	if rcv == nil { return nil }
	t := &AoeDataResponseT{}
	rcv.UnPackTo(t)
	return t
}

type AoeDataResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsAoeDataResponse(buf []byte, offset flatbuffers.UOffsetT) *AoeDataResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AoeDataResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAoeDataResponse(buf []byte, offset flatbuffers.UOffsetT) *AoeDataResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AoeDataResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AoeDataResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AoeDataResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AoeDataResponse) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *AoeDataResponse) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *AoeDataResponse) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AoeDataResponse) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

///Error code from AoE header
func (rcv *AoeDataResponse) ErrorCode() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Error code from AoE header
func (rcv *AoeDataResponse) MutateErrorCode(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

///Command response from AoE header
func (rcv *AoeDataResponse) CmdResult() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Command response from AoE header
func (rcv *AoeDataResponse) MutateCmdResult(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func AoeDataResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func AoeDataResponseAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(data), 0)
}
func AoeDataResponseStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func AoeDataResponseAddErrorCode(builder *flatbuffers.Builder, errorCode uint32) {
	builder.PrependUint32Slot(1, errorCode, 0)
}
func AoeDataResponseAddCmdResult(builder *flatbuffers.Builder, cmdResult uint32) {
	builder.PrependUint32Slot(2, cmdResult, 0)
}
func AoeDataResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
