// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///This node provide access to VoE data of an EtherCAT slave device.
///A write request will send data to the slave device.
///A read request read previously received data from a EtherCAT slave device out of master buffer.
///The slave must support the mailbox protocol “Vendor specific protocol over EtherCAT” (VoE).
///Note: the Slave must be in EtherCAT state PreOP, SafeOP or OP for mailbox communication.
type VoeDataRequestT struct {
	AddressType Addresstype `json:"addressType"`
	Address uint16 `json:"address"`
	Data []byte `json:"data"`
	MaxLength uint32 `json:"maxLength"`
}

func (t *VoeDataRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	VoeDataRequestStart(builder)
	VoeDataRequestAddAddressType(builder, t.AddressType)
	VoeDataRequestAddAddress(builder, t.Address)
	VoeDataRequestAddData(builder, dataOffset)
	VoeDataRequestAddMaxLength(builder, t.MaxLength)
	return VoeDataRequestEnd(builder)
}

func (rcv *VoeDataRequest) UnPackTo(t *VoeDataRequestT) {
	t.AddressType = rcv.AddressType()
	t.Address = rcv.Address()
	t.Data = rcv.DataBytes()
	t.MaxLength = rcv.MaxLength()
}

func (rcv *VoeDataRequest) UnPack() *VoeDataRequestT {
	if rcv == nil { return nil }
	t := &VoeDataRequestT{}
	rcv.UnPackTo(t)
	return t
}

type VoeDataRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsVoeDataRequest(buf []byte, offset flatbuffers.UOffsetT) *VoeDataRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &VoeDataRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsVoeDataRequest(buf []byte, offset flatbuffers.UOffsetT) *VoeDataRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &VoeDataRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *VoeDataRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *VoeDataRequest) Table() flatbuffers.Table {
	return rcv._tab
}

///Address type:
///undefined: Undefined - do not use
///autoincrement: Auto increment address
///fixedphysical: EtherCAT address (fixed physical address)
func (rcv *VoeDataRequest) AddressType() Addresstype {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Addresstype(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

///Address type:
///undefined: Undefined - do not use
///autoincrement: Auto increment address
///fixedphysical: EtherCAT address (fixed physical address)
func (rcv *VoeDataRequest) MutateAddressType(n Addresstype) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

///Address depending on addressType.
func (rcv *VoeDataRequest) Address() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Address depending on addressType.
func (rcv *VoeDataRequest) MutateAddress(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

///Data buffer for write data, include VoE header
func (rcv *VoeDataRequest) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *VoeDataRequest) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *VoeDataRequest) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///Data buffer for write data, include VoE header
func (rcv *VoeDataRequest) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

///Size in Bytes of buffer or variable provided for data reception
func (rcv *VoeDataRequest) MaxLength() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Size in Bytes of buffer or variable provided for data reception
func (rcv *VoeDataRequest) MutateMaxLength(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func VoeDataRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func VoeDataRequestAddAddressType(builder *flatbuffers.Builder, addressType Addresstype) {
	builder.PrependByteSlot(0, byte(addressType), 0)
}
func VoeDataRequestAddAddress(builder *flatbuffers.Builder, address uint16) {
	builder.PrependUint16Slot(1, address, 0)
}
func VoeDataRequestAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(data), 0)
}
func VoeDataRequestStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func VoeDataRequestAddMaxLength(builder *flatbuffers.Builder, maxLength uint32) {
	builder.PrependUint32Slot(3, maxLength, 0)
}
func VoeDataRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
