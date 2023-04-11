// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AddressMappingEntryT struct {
	Address *AddressedRequestT
	SlaveName string
}

func (t *AddressMappingEntryT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	addressOffset := t.Address.Pack(builder)
	slaveNameOffset := builder.CreateString(t.SlaveName)
	AddressMappingEntryStart(builder)
	AddressMappingEntryAddAddress(builder, addressOffset)
	AddressMappingEntryAddSlaveName(builder, slaveNameOffset)
	return AddressMappingEntryEnd(builder)
}

func (rcv *AddressMappingEntry) UnPackTo(t *AddressMappingEntryT) {
	t.Address = rcv.Address(nil).UnPack()
	t.SlaveName = string(rcv.SlaveName())
}

func (rcv *AddressMappingEntry) UnPack() *AddressMappingEntryT {
	if rcv == nil { return nil }
	t := &AddressMappingEntryT{}
	rcv.UnPackTo(t)
	return t
}

type AddressMappingEntry struct {
	_tab flatbuffers.Table
}

func GetRootAsAddressMappingEntry(buf []byte, offset flatbuffers.UOffsetT) *AddressMappingEntry {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AddressMappingEntry{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAddressMappingEntry(buf []byte, offset flatbuffers.UOffsetT) *AddressMappingEntry {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AddressMappingEntry{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AddressMappingEntry) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AddressMappingEntry) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AddressMappingEntry) Address(obj *AddressedRequest) *AddressedRequest {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AddressedRequest)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *AddressMappingEntry) SlaveName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func AddressMappingEntryStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AddressMappingEntryAddAddress(builder *flatbuffers.Builder, address flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(address), 0)
}
func AddressMappingEntryAddSlaveName(builder *flatbuffers.Builder, slaveName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(slaveName), 0)
}
func AddressMappingEntryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
