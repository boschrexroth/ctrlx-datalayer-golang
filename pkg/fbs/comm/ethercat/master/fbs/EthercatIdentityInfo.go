// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type EthercatIdentityInfoT struct {
	VendorId uint32
	ProductCode uint32
	RevisionNumber uint32
	SerialNumber uint32
}

func (t *EthercatIdentityInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	return CreateEthercatIdentityInfo(builder, t.VendorId, t.ProductCode, t.RevisionNumber, t.SerialNumber)
}
func (rcv *EthercatIdentityInfo) UnPackTo(t *EthercatIdentityInfoT) {
	t.VendorId = rcv.VendorId()
	t.ProductCode = rcv.ProductCode()
	t.RevisionNumber = rcv.RevisionNumber()
	t.SerialNumber = rcv.SerialNumber()
}

func (rcv *EthercatIdentityInfo) UnPack() *EthercatIdentityInfoT {
	if rcv == nil { return nil }
	t := &EthercatIdentityInfoT{}
	rcv.UnPackTo(t)
	return t
}

type EthercatIdentityInfo struct {
	_tab flatbuffers.Struct
}

func (rcv *EthercatIdentityInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EthercatIdentityInfo) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *EthercatIdentityInfo) VendorId() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *EthercatIdentityInfo) MutateVendorId(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *EthercatIdentityInfo) ProductCode() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *EthercatIdentityInfo) MutateProductCode(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func (rcv *EthercatIdentityInfo) RevisionNumber() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *EthercatIdentityInfo) MutateRevisionNumber(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func (rcv *EthercatIdentityInfo) SerialNumber() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(12))
}
func (rcv *EthercatIdentityInfo) MutateSerialNumber(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(12), n)
}

func CreateEthercatIdentityInfo(builder *flatbuffers.Builder, vendorId uint32, productCode uint32, revisionNumber uint32, serialNumber uint32) flatbuffers.UOffsetT {
	builder.Prep(4, 16)
	builder.PrependUint32(serialNumber)
	builder.PrependUint32(revisionNumber)
	builder.PrependUint32(productCode)
	builder.PrependUint32(vendorId)
	return builder.Offset()
}
