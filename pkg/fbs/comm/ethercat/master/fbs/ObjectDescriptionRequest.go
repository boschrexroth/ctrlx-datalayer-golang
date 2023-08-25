// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ObjectDescriptionRequestT struct {
	AddressType Addresstype `json:"addressType"`
	Address uint16 `json:"address"`
	ObjectIndex uint16 `json:"objectIndex"`
	MaxLength uint32 `json:"maxLength"`
}

func (t *ObjectDescriptionRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ObjectDescriptionRequestStart(builder)
	ObjectDescriptionRequestAddAddressType(builder, t.AddressType)
	ObjectDescriptionRequestAddAddress(builder, t.Address)
	ObjectDescriptionRequestAddObjectIndex(builder, t.ObjectIndex)
	ObjectDescriptionRequestAddMaxLength(builder, t.MaxLength)
	return ObjectDescriptionRequestEnd(builder)
}

func (rcv *ObjectDescriptionRequest) UnPackTo(t *ObjectDescriptionRequestT) {
	t.AddressType = rcv.AddressType()
	t.Address = rcv.Address()
	t.ObjectIndex = rcv.ObjectIndex()
	t.MaxLength = rcv.MaxLength()
}

func (rcv *ObjectDescriptionRequest) UnPack() *ObjectDescriptionRequestT {
	if rcv == nil { return nil }
	t := &ObjectDescriptionRequestT{}
	rcv.UnPackTo(t)
	return t
}

type ObjectDescriptionRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsObjectDescriptionRequest(buf []byte, offset flatbuffers.UOffsetT) *ObjectDescriptionRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ObjectDescriptionRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsObjectDescriptionRequest(buf []byte, offset flatbuffers.UOffsetT) *ObjectDescriptionRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ObjectDescriptionRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ObjectDescriptionRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ObjectDescriptionRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ObjectDescriptionRequest) AddressType() Addresstype {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Addresstype(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ObjectDescriptionRequest) MutateAddressType(n Addresstype) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *ObjectDescriptionRequest) Address() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ObjectDescriptionRequest) MutateAddress(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

func (rcv *ObjectDescriptionRequest) ObjectIndex() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ObjectDescriptionRequest) MutateObjectIndex(n uint16) bool {
	return rcv._tab.MutateUint16Slot(8, n)
}

func (rcv *ObjectDescriptionRequest) MaxLength() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ObjectDescriptionRequest) MutateMaxLength(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func ObjectDescriptionRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ObjectDescriptionRequestAddAddressType(builder *flatbuffers.Builder, addressType Addresstype) {
	builder.PrependByteSlot(0, byte(addressType), 0)
}
func ObjectDescriptionRequestAddAddress(builder *flatbuffers.Builder, address uint16) {
	builder.PrependUint16Slot(1, address, 0)
}
func ObjectDescriptionRequestAddObjectIndex(builder *flatbuffers.Builder, objectIndex uint16) {
	builder.PrependUint16Slot(2, objectIndex, 0)
}
func ObjectDescriptionRequestAddMaxLength(builder *flatbuffers.Builder, maxLength uint32) {
	builder.PrependUint32Slot(3, maxLength, 0)
}
func ObjectDescriptionRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
