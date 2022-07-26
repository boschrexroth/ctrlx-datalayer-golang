// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ProfileChannelInfoRequestT struct {
	AddressType Addresstype
	Address uint16
	Channel uint32
}

func (t *ProfileChannelInfoRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ProfileChannelInfoRequestStart(builder)
	ProfileChannelInfoRequestAddAddressType(builder, t.AddressType)
	ProfileChannelInfoRequestAddAddress(builder, t.Address)
	ProfileChannelInfoRequestAddChannel(builder, t.Channel)
	return ProfileChannelInfoRequestEnd(builder)
}

func (rcv *ProfileChannelInfoRequest) UnPackTo(t *ProfileChannelInfoRequestT) {
	t.AddressType = rcv.AddressType()
	t.Address = rcv.Address()
	t.Channel = rcv.Channel()
}

func (rcv *ProfileChannelInfoRequest) UnPack() *ProfileChannelInfoRequestT {
	if rcv == nil { return nil }
	t := &ProfileChannelInfoRequestT{}
	rcv.UnPackTo(t)
	return t
}

type ProfileChannelInfoRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsProfileChannelInfoRequest(buf []byte, offset flatbuffers.UOffsetT) *ProfileChannelInfoRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ProfileChannelInfoRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsProfileChannelInfoRequest(buf []byte, offset flatbuffers.UOffsetT) *ProfileChannelInfoRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ProfileChannelInfoRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ProfileChannelInfoRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ProfileChannelInfoRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ProfileChannelInfoRequest) AddressType() Addresstype {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Addresstype(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ProfileChannelInfoRequest) MutateAddressType(n Addresstype) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *ProfileChannelInfoRequest) Address() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ProfileChannelInfoRequest) MutateAddress(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

func (rcv *ProfileChannelInfoRequest) Channel() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ProfileChannelInfoRequest) MutateChannel(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func ProfileChannelInfoRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ProfileChannelInfoRequestAddAddressType(builder *flatbuffers.Builder, addressType Addresstype) {
	builder.PrependByteSlot(0, byte(addressType), 0)
}
func ProfileChannelInfoRequestAddAddress(builder *flatbuffers.Builder, address uint16) {
	builder.PrependUint16Slot(1, address, 0)
}
func ProfileChannelInfoRequestAddChannel(builder *flatbuffers.Builder, channel uint32) {
	builder.PrependUint32Slot(2, channel, 0)
}
func ProfileChannelInfoRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
