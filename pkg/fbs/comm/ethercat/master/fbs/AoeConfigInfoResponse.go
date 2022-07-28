// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AoeConfigInfoResponseT struct {
	NetId []byte
}

func (t *AoeConfigInfoResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	netIdOffset := flatbuffers.UOffsetT(0)
	if t.NetId != nil {
		netIdOffset = builder.CreateByteString(t.NetId)
	}
	AoeConfigInfoResponseStart(builder)
	AoeConfigInfoResponseAddNetId(builder, netIdOffset)
	return AoeConfigInfoResponseEnd(builder)
}

func (rcv *AoeConfigInfoResponse) UnPackTo(t *AoeConfigInfoResponseT) {
	t.NetId = rcv.NetIdBytes()
}

func (rcv *AoeConfigInfoResponse) UnPack() *AoeConfigInfoResponseT {
	if rcv == nil { return nil }
	t := &AoeConfigInfoResponseT{}
	rcv.UnPackTo(t)
	return t
}

type AoeConfigInfoResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsAoeConfigInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *AoeConfigInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AoeConfigInfoResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAoeConfigInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *AoeConfigInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AoeConfigInfoResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AoeConfigInfoResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AoeConfigInfoResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AoeConfigInfoResponse) NetId(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *AoeConfigInfoResponse) NetIdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *AoeConfigInfoResponse) NetIdBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AoeConfigInfoResponse) MutateNetId(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func AoeConfigInfoResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func AoeConfigInfoResponseAddNetId(builder *flatbuffers.Builder, netId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(netId), 0)
}
func AoeConfigInfoResponseStartNetIdVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func AoeConfigInfoResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
