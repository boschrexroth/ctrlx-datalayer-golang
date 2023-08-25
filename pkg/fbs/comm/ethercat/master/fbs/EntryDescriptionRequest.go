// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type EntryDescriptionRequestT struct {
	AddressType Addresstype `json:"addressType"`
	Address uint16 `json:"address"`
	ObjectIndex uint16 `json:"objectIndex"`
	SubIndex byte `json:"subIndex"`
	ValueInfo byte `json:"valueInfo"`
	MaxLength uint32 `json:"maxLength"`
}

func (t *EntryDescriptionRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	EntryDescriptionRequestStart(builder)
	EntryDescriptionRequestAddAddressType(builder, t.AddressType)
	EntryDescriptionRequestAddAddress(builder, t.Address)
	EntryDescriptionRequestAddObjectIndex(builder, t.ObjectIndex)
	EntryDescriptionRequestAddSubIndex(builder, t.SubIndex)
	EntryDescriptionRequestAddValueInfo(builder, t.ValueInfo)
	EntryDescriptionRequestAddMaxLength(builder, t.MaxLength)
	return EntryDescriptionRequestEnd(builder)
}

func (rcv *EntryDescriptionRequest) UnPackTo(t *EntryDescriptionRequestT) {
	t.AddressType = rcv.AddressType()
	t.Address = rcv.Address()
	t.ObjectIndex = rcv.ObjectIndex()
	t.SubIndex = rcv.SubIndex()
	t.ValueInfo = rcv.ValueInfo()
	t.MaxLength = rcv.MaxLength()
}

func (rcv *EntryDescriptionRequest) UnPack() *EntryDescriptionRequestT {
	if rcv == nil { return nil }
	t := &EntryDescriptionRequestT{}
	rcv.UnPackTo(t)
	return t
}

type EntryDescriptionRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsEntryDescriptionRequest(buf []byte, offset flatbuffers.UOffsetT) *EntryDescriptionRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EntryDescriptionRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEntryDescriptionRequest(buf []byte, offset flatbuffers.UOffsetT) *EntryDescriptionRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &EntryDescriptionRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *EntryDescriptionRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EntryDescriptionRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *EntryDescriptionRequest) AddressType() Addresstype {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Addresstype(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *EntryDescriptionRequest) MutateAddressType(n Addresstype) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *EntryDescriptionRequest) Address() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *EntryDescriptionRequest) MutateAddress(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

func (rcv *EntryDescriptionRequest) ObjectIndex() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *EntryDescriptionRequest) MutateObjectIndex(n uint16) bool {
	return rcv._tab.MutateUint16Slot(8, n)
}

func (rcv *EntryDescriptionRequest) SubIndex() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *EntryDescriptionRequest) MutateSubIndex(n byte) bool {
	return rcv._tab.MutateByteSlot(10, n)
}

func (rcv *EntryDescriptionRequest) ValueInfo() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *EntryDescriptionRequest) MutateValueInfo(n byte) bool {
	return rcv._tab.MutateByteSlot(12, n)
}

func (rcv *EntryDescriptionRequest) MaxLength() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *EntryDescriptionRequest) MutateMaxLength(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

func EntryDescriptionRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func EntryDescriptionRequestAddAddressType(builder *flatbuffers.Builder, addressType Addresstype) {
	builder.PrependByteSlot(0, byte(addressType), 0)
}
func EntryDescriptionRequestAddAddress(builder *flatbuffers.Builder, address uint16) {
	builder.PrependUint16Slot(1, address, 0)
}
func EntryDescriptionRequestAddObjectIndex(builder *flatbuffers.Builder, objectIndex uint16) {
	builder.PrependUint16Slot(2, objectIndex, 0)
}
func EntryDescriptionRequestAddSubIndex(builder *flatbuffers.Builder, subIndex byte) {
	builder.PrependByteSlot(3, subIndex, 0)
}
func EntryDescriptionRequestAddValueInfo(builder *flatbuffers.Builder, valueInfo byte) {
	builder.PrependByteSlot(4, valueInfo, 0)
}
func EntryDescriptionRequestAddMaxLength(builder *flatbuffers.Builder, maxLength uint32) {
	builder.PrependUint32Slot(5, maxLength, 0)
}
func EntryDescriptionRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
