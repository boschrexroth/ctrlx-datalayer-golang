// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AoeControlRequestT struct {
	AddressType Addresstype `json:"addressType"`
	Address uint16 `json:"address"`
	TargetNetId []byte `json:"targetNetId"`
	TargetPort uint16 `json:"targetPort"`
	AoeState uint16 `json:"aoeState"`
	DeviceState uint16 `json:"deviceState"`
	Data []byte `json:"data"`
}

func (t *AoeControlRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	targetNetIdOffset := flatbuffers.UOffsetT(0)
	if t.TargetNetId != nil {
		targetNetIdOffset = builder.CreateByteString(t.TargetNetId)
	}
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	AoeControlRequestStart(builder)
	AoeControlRequestAddAddressType(builder, t.AddressType)
	AoeControlRequestAddAddress(builder, t.Address)
	AoeControlRequestAddTargetNetId(builder, targetNetIdOffset)
	AoeControlRequestAddTargetPort(builder, t.TargetPort)
	AoeControlRequestAddAoeState(builder, t.AoeState)
	AoeControlRequestAddDeviceState(builder, t.DeviceState)
	AoeControlRequestAddData(builder, dataOffset)
	return AoeControlRequestEnd(builder)
}

func (rcv *AoeControlRequest) UnPackTo(t *AoeControlRequestT) {
	t.AddressType = rcv.AddressType()
	t.Address = rcv.Address()
	t.TargetNetId = rcv.TargetNetIdBytes()
	t.TargetPort = rcv.TargetPort()
	t.AoeState = rcv.AoeState()
	t.DeviceState = rcv.DeviceState()
	t.Data = rcv.DataBytes()
}

func (rcv *AoeControlRequest) UnPack() *AoeControlRequestT {
	if rcv == nil { return nil }
	t := &AoeControlRequestT{}
	rcv.UnPackTo(t)
	return t
}

type AoeControlRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsAoeControlRequest(buf []byte, offset flatbuffers.UOffsetT) *AoeControlRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AoeControlRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAoeControlRequest(buf []byte, offset flatbuffers.UOffsetT) *AoeControlRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AoeControlRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AoeControlRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AoeControlRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AoeControlRequest) AddressType() Addresstype {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Addresstype(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *AoeControlRequest) MutateAddressType(n Addresstype) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *AoeControlRequest) Address() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AoeControlRequest) MutateAddress(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

func (rcv *AoeControlRequest) TargetNetId(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *AoeControlRequest) TargetNetIdLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *AoeControlRequest) TargetNetIdBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AoeControlRequest) MutateTargetNetId(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *AoeControlRequest) TargetPort() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AoeControlRequest) MutateTargetPort(n uint16) bool {
	return rcv._tab.MutateUint16Slot(10, n)
}

func (rcv *AoeControlRequest) AoeState() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AoeControlRequest) MutateAoeState(n uint16) bool {
	return rcv._tab.MutateUint16Slot(12, n)
}

func (rcv *AoeControlRequest) DeviceState() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AoeControlRequest) MutateDeviceState(n uint16) bool {
	return rcv._tab.MutateUint16Slot(14, n)
}

func (rcv *AoeControlRequest) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *AoeControlRequest) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *AoeControlRequest) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AoeControlRequest) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func AoeControlRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func AoeControlRequestAddAddressType(builder *flatbuffers.Builder, addressType Addresstype) {
	builder.PrependByteSlot(0, byte(addressType), 0)
}
func AoeControlRequestAddAddress(builder *flatbuffers.Builder, address uint16) {
	builder.PrependUint16Slot(1, address, 0)
}
func AoeControlRequestAddTargetNetId(builder *flatbuffers.Builder, targetNetId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(targetNetId), 0)
}
func AoeControlRequestStartTargetNetIdVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func AoeControlRequestAddTargetPort(builder *flatbuffers.Builder, targetPort uint16) {
	builder.PrependUint16Slot(3, targetPort, 0)
}
func AoeControlRequestAddAoeState(builder *flatbuffers.Builder, aoeState uint16) {
	builder.PrependUint16Slot(4, aoeState, 0)
}
func AoeControlRequestAddDeviceState(builder *flatbuffers.Builder, deviceState uint16) {
	builder.PrependUint16Slot(5, deviceState, 0)
}
func AoeControlRequestAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(data), 0)
}
func AoeControlRequestStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func AoeControlRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
