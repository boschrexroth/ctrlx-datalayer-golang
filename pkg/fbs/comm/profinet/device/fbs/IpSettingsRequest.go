// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type IpSettingsRequestT struct {
	Address []byte `json:"address"`
	SubnetMask []byte `json:"subnet_mask"`
	GatewayAddress []byte `json:"gateway_address"`
}

func (t *IpSettingsRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	addressOffset := flatbuffers.UOffsetT(0)
	if t.Address != nil {
		addressOffset = builder.CreateByteString(t.Address)
	}
	subnetMaskOffset := flatbuffers.UOffsetT(0)
	if t.SubnetMask != nil {
		subnetMaskOffset = builder.CreateByteString(t.SubnetMask)
	}
	gatewayAddressOffset := flatbuffers.UOffsetT(0)
	if t.GatewayAddress != nil {
		gatewayAddressOffset = builder.CreateByteString(t.GatewayAddress)
	}
	IpSettingsRequestStart(builder)
	IpSettingsRequestAddAddress(builder, addressOffset)
	IpSettingsRequestAddSubnetMask(builder, subnetMaskOffset)
	IpSettingsRequestAddGatewayAddress(builder, gatewayAddressOffset)
	return IpSettingsRequestEnd(builder)
}

func (rcv *IpSettingsRequest) UnPackTo(t *IpSettingsRequestT) {
	t.Address = rcv.AddressBytes()
	t.SubnetMask = rcv.SubnetMaskBytes()
	t.GatewayAddress = rcv.GatewayAddressBytes()
}

func (rcv *IpSettingsRequest) UnPack() *IpSettingsRequestT {
	if rcv == nil { return nil }
	t := &IpSettingsRequestT{}
	rcv.UnPackTo(t)
	return t
}

type IpSettingsRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsIpSettingsRequest(buf []byte, offset flatbuffers.UOffsetT) *IpSettingsRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IpSettingsRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsIpSettingsRequest(buf []byte, offset flatbuffers.UOffsetT) *IpSettingsRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &IpSettingsRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *IpSettingsRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IpSettingsRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *IpSettingsRequest) Address(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *IpSettingsRequest) AddressLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *IpSettingsRequest) AddressBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *IpSettingsRequest) MutateAddress(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *IpSettingsRequest) SubnetMask(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *IpSettingsRequest) SubnetMaskLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *IpSettingsRequest) SubnetMaskBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *IpSettingsRequest) MutateSubnetMask(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *IpSettingsRequest) GatewayAddress(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *IpSettingsRequest) GatewayAddressLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *IpSettingsRequest) GatewayAddressBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *IpSettingsRequest) MutateGatewayAddress(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func IpSettingsRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func IpSettingsRequestAddAddress(builder *flatbuffers.Builder, address flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(address), 0)
}
func IpSettingsRequestStartAddressVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func IpSettingsRequestAddSubnetMask(builder *flatbuffers.Builder, subnetMask flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(subnetMask), 0)
}
func IpSettingsRequestStartSubnetMaskVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func IpSettingsRequestAddGatewayAddress(builder *flatbuffers.Builder, gatewayAddress flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(gatewayAddress), 0)
}
func IpSettingsRequestStartGatewayAddressVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func IpSettingsRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
