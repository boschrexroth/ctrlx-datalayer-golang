// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AoeDataRequestT struct {
	AddressType Addresstype `json:"addressType"`
	Address uint16 `json:"address"`
	TargetNetId []byte `json:"targetNetId"`
	TargetPort uint16 `json:"targetPort"`
	IndexGroup uint32 `json:"indexGroup"`
	IndexOffset uint32 `json:"indexOffset"`
	Data []byte `json:"data"`
	MaxLength uint32 `json:"maxLength"`
}

func (t *AoeDataRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	targetNetIdOffset := flatbuffers.UOffsetT(0)
	if t.TargetNetId != nil {
		targetNetIdOffset = builder.CreateByteString(t.TargetNetId)
	}
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	AoeDataRequestStart(builder)
	AoeDataRequestAddAddressType(builder, t.AddressType)
	AoeDataRequestAddAddress(builder, t.Address)
	AoeDataRequestAddTargetNetId(builder, targetNetIdOffset)
	AoeDataRequestAddTargetPort(builder, t.TargetPort)
	AoeDataRequestAddIndexGroup(builder, t.IndexGroup)
	AoeDataRequestAddIndexOffset(builder, t.IndexOffset)
	AoeDataRequestAddData(builder, dataOffset)
	AoeDataRequestAddMaxLength(builder, t.MaxLength)
	return AoeDataRequestEnd(builder)
}

func (rcv *AoeDataRequest) UnPackTo(t *AoeDataRequestT) {
	t.AddressType = rcv.AddressType()
	t.Address = rcv.Address()
	t.TargetNetId = rcv.TargetNetIdBytes()
	t.TargetPort = rcv.TargetPort()
	t.IndexGroup = rcv.IndexGroup()
	t.IndexOffset = rcv.IndexOffset()
	t.Data = rcv.DataBytes()
	t.MaxLength = rcv.MaxLength()
}

func (rcv *AoeDataRequest) UnPack() *AoeDataRequestT {
	if rcv == nil { return nil }
	t := &AoeDataRequestT{}
	rcv.UnPackTo(t)
	return t
}

type AoeDataRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsAoeDataRequest(buf []byte, offset flatbuffers.UOffsetT) *AoeDataRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AoeDataRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAoeDataRequest(buf []byte, offset flatbuffers.UOffsetT) *AoeDataRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AoeDataRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AoeDataRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AoeDataRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AoeDataRequest) AddressType() Addresstype {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Addresstype(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *AoeDataRequest) MutateAddressType(n Addresstype) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *AoeDataRequest) Address() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AoeDataRequest) MutateAddress(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

func (rcv *AoeDataRequest) TargetNetId(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *AoeDataRequest) TargetNetIdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *AoeDataRequest) TargetNetIdBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AoeDataRequest) MutateTargetNetId(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *AoeDataRequest) TargetPort() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AoeDataRequest) MutateTargetPort(n uint16) bool {
	return rcv._tab.MutateUint16Slot(10, n)
}

func (rcv *AoeDataRequest) IndexGroup() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AoeDataRequest) MutateIndexGroup(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func (rcv *AoeDataRequest) IndexOffset() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AoeDataRequest) MutateIndexOffset(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

func (rcv *AoeDataRequest) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *AoeDataRequest) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *AoeDataRequest) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AoeDataRequest) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *AoeDataRequest) MaxLength() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AoeDataRequest) MutateMaxLength(n uint32) bool {
	return rcv._tab.MutateUint32Slot(18, n)
}

func AoeDataRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func AoeDataRequestAddAddressType(builder *flatbuffers.Builder, addressType Addresstype) {
	builder.PrependByteSlot(0, byte(addressType), 0)
}
func AoeDataRequestAddAddress(builder *flatbuffers.Builder, address uint16) {
	builder.PrependUint16Slot(1, address, 0)
}
func AoeDataRequestAddTargetNetId(builder *flatbuffers.Builder, targetNetId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(targetNetId), 0)
}
func AoeDataRequestStartTargetNetIdVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func AoeDataRequestAddTargetPort(builder *flatbuffers.Builder, targetPort uint16) {
	builder.PrependUint16Slot(3, targetPort, 0)
}
func AoeDataRequestAddIndexGroup(builder *flatbuffers.Builder, indexGroup uint32) {
	builder.PrependUint32Slot(4, indexGroup, 0)
}
func AoeDataRequestAddIndexOffset(builder *flatbuffers.Builder, indexOffset uint32) {
	builder.PrependUint32Slot(5, indexOffset, 0)
}
func AoeDataRequestAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(data), 0)
}
func AoeDataRequestStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func AoeDataRequestAddMaxLength(builder *flatbuffers.Builder, maxLength uint32) {
	builder.PrependUint32Slot(7, maxLength, 0)
}
func AoeDataRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
