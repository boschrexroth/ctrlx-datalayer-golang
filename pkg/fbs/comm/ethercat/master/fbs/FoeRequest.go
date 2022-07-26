// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FoeRequestT struct {
	AddressType Addresstype
	Address uint16
	Data []byte
	MaxLength uint32
	Filename string
	Password uint32
	Timeout uint32
}

func (t *FoeRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	filenameOffset := builder.CreateString(t.Filename)
	FoeRequestStart(builder)
	FoeRequestAddAddressType(builder, t.AddressType)
	FoeRequestAddAddress(builder, t.Address)
	FoeRequestAddData(builder, dataOffset)
	FoeRequestAddMaxLength(builder, t.MaxLength)
	FoeRequestAddFilename(builder, filenameOffset)
	FoeRequestAddPassword(builder, t.Password)
	FoeRequestAddTimeout(builder, t.Timeout)
	return FoeRequestEnd(builder)
}

func (rcv *FoeRequest) UnPackTo(t *FoeRequestT) {
	t.AddressType = rcv.AddressType()
	t.Address = rcv.Address()
	t.Data = rcv.DataBytes()
	t.MaxLength = rcv.MaxLength()
	t.Filename = string(rcv.Filename())
	t.Password = rcv.Password()
	t.Timeout = rcv.Timeout()
}

func (rcv *FoeRequest) UnPack() *FoeRequestT {
	if rcv == nil { return nil }
	t := &FoeRequestT{}
	rcv.UnPackTo(t)
	return t
}

type FoeRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsFoeRequest(buf []byte, offset flatbuffers.UOffsetT) *FoeRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FoeRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFoeRequest(buf []byte, offset flatbuffers.UOffsetT) *FoeRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FoeRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FoeRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FoeRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FoeRequest) AddressType() Addresstype {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Addresstype(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *FoeRequest) MutateAddressType(n Addresstype) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *FoeRequest) Address() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FoeRequest) MutateAddress(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

func (rcv *FoeRequest) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *FoeRequest) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FoeRequest) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FoeRequest) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *FoeRequest) MaxLength() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FoeRequest) MutateMaxLength(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *FoeRequest) Filename() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FoeRequest) Password() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FoeRequest) MutatePassword(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

func (rcv *FoeRequest) Timeout() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FoeRequest) MutateTimeout(n uint32) bool {
	return rcv._tab.MutateUint32Slot(16, n)
}

func FoeRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func FoeRequestAddAddressType(builder *flatbuffers.Builder, addressType Addresstype) {
	builder.PrependByteSlot(0, byte(addressType), 0)
}
func FoeRequestAddAddress(builder *flatbuffers.Builder, address uint16) {
	builder.PrependUint16Slot(1, address, 0)
}
func FoeRequestAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(data), 0)
}
func FoeRequestStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FoeRequestAddMaxLength(builder *flatbuffers.Builder, maxLength uint32) {
	builder.PrependUint32Slot(3, maxLength, 0)
}
func FoeRequestAddFilename(builder *flatbuffers.Builder, filename flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(filename), 0)
}
func FoeRequestAddPassword(builder *flatbuffers.Builder, password uint32) {
	builder.PrependUint32Slot(5, password, 0)
}
func FoeRequestAddTimeout(builder *flatbuffers.Builder, timeout uint32) {
	builder.PrependUint32Slot(6, timeout, 0)
}
func FoeRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
