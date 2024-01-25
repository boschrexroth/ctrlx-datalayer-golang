// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///This node reads/writes a Service-Data-Object (SDO) from an EtherCAT slave device. 
///The slave must support the mailbox protocol “CAN application protocol over EtherCAT” (CoE). 
///Note: the Slave must be in EtherCAT state PreOP, SafeOP or OP for mailbox communication.
///The object is addressed with index and subindex. 
type SDORequestT struct {
	AddressType Addresstype `json:"addressType"`
	Address uint16 `json:"address"`
	ObjectIndex uint16 `json:"objectIndex"`
	SubIndex byte `json:"subIndex"`
	Flags SDOFlags `json:"flags"`
	Data []byte `json:"data"`
	MaxLength uint32 `json:"maxLength"`
}

func (t *SDORequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	SDORequestStart(builder)
	SDORequestAddAddressType(builder, t.AddressType)
	SDORequestAddAddress(builder, t.Address)
	SDORequestAddObjectIndex(builder, t.ObjectIndex)
	SDORequestAddSubIndex(builder, t.SubIndex)
	SDORequestAddFlags(builder, t.Flags)
	SDORequestAddData(builder, dataOffset)
	SDORequestAddMaxLength(builder, t.MaxLength)
	return SDORequestEnd(builder)
}

func (rcv *SDORequest) UnPackTo(t *SDORequestT) {
	t.AddressType = rcv.AddressType()
	t.Address = rcv.Address()
	t.ObjectIndex = rcv.ObjectIndex()
	t.SubIndex = rcv.SubIndex()
	t.Flags = rcv.Flags()
	t.Data = rcv.DataBytes()
	t.MaxLength = rcv.MaxLength()
}

func (rcv *SDORequest) UnPack() *SDORequestT {
	if rcv == nil { return nil }
	t := &SDORequestT{}
	rcv.UnPackTo(t)
	return t
}

type SDORequest struct {
	_tab flatbuffers.Table
}

func GetRootAsSDORequest(buf []byte, offset flatbuffers.UOffsetT) *SDORequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SDORequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSDORequest(buf []byte, offset flatbuffers.UOffsetT) *SDORequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SDORequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SDORequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SDORequest) Table() flatbuffers.Table {
	return rcv._tab
}

///Address type: 
///undefined: Undefined - do not use
///autoincrement: Auto increment address
///fixedphysical: EtherCAT address (fixed physical address)
func (rcv *SDORequest) AddressType() Addresstype {
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
func (rcv *SDORequest) MutateAddressType(n Addresstype) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

///Address depending on addressType.
func (rcv *SDORequest) Address() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Address depending on addressType.
func (rcv *SDORequest) MutateAddress(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

///Object index
func (rcv *SDORequest) ObjectIndex() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Object index
func (rcv *SDORequest) MutateObjectIndex(n uint16) bool {
	return rcv._tab.MutateUint16Slot(8, n)
}

///Sub-Index of the object
func (rcv *SDORequest) SubIndex() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

///Sub-Index of the object
func (rcv *SDORequest) MutateSubIndex(n byte) bool {
	return rcv._tab.MutateByteSlot(10, n)
}

///With completeAccess all subindices of the object can be read in one upload call. 
///Hereby the subindex input is the start subindex:
///noFlags = all subindices including subindex 0 (start subindex 0)
///completeAccess = all subindices without subindex 0 (start subindex 1)
///Note: completeAccess is not supported by all slaves or can also be object-specific.
func (rcv *SDORequest) Flags() SDOFlags {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return SDOFlags(rcv._tab.GetUint32(o + rcv._tab.Pos))
	}
	return 0
}

///With completeAccess all subindices of the object can be read in one upload call. 
///Hereby the subindex input is the start subindex:
///noFlags = all subindices including subindex 0 (start subindex 0)
///completeAccess = all subindices without subindex 0 (start subindex 1)
///Note: completeAccess is not supported by all slaves or can also be object-specific.
func (rcv *SDORequest) MutateFlags(n SDOFlags) bool {
	return rcv._tab.MutateUint32Slot(12, uint32(n))
}

///Data buffer for writing data
func (rcv *SDORequest) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *SDORequest) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SDORequest) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///Data buffer for writing data
func (rcv *SDORequest) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

///Size in Bytes of buffer or variable provided for data reception
func (rcv *SDORequest) MaxLength() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Size in Bytes of buffer or variable provided for data reception
func (rcv *SDORequest) MutateMaxLength(n uint32) bool {
	return rcv._tab.MutateUint32Slot(16, n)
}

func SDORequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func SDORequestAddAddressType(builder *flatbuffers.Builder, addressType Addresstype) {
	builder.PrependByteSlot(0, byte(addressType), 0)
}
func SDORequestAddAddress(builder *flatbuffers.Builder, address uint16) {
	builder.PrependUint16Slot(1, address, 0)
}
func SDORequestAddObjectIndex(builder *flatbuffers.Builder, objectIndex uint16) {
	builder.PrependUint16Slot(2, objectIndex, 0)
}
func SDORequestAddSubIndex(builder *flatbuffers.Builder, subIndex byte) {
	builder.PrependByteSlot(3, subIndex, 0)
}
func SDORequestAddFlags(builder *flatbuffers.Builder, flags SDOFlags) {
	builder.PrependUint32Slot(4, uint32(flags), 0)
}
func SDORequestAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(data), 0)
}
func SDORequestStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func SDORequestAddMaxLength(builder *flatbuffers.Builder, maxLength uint32) {
	builder.PrependUint32Slot(6, maxLength, 0)
}
func SDORequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
