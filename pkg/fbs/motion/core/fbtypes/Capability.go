// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Response for a single capability
type CapabilityT struct {
	Count uint32 `json:"count"`
	MainDiag uint32 `json:"mainDiag"`
	DetailDiag uint32 `json:"detailDiag"`
	AddInfo string `json:"addInfo"`
}

func (t *CapabilityT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	addInfoOffset := flatbuffers.UOffsetT(0)
	if t.AddInfo != "" {
		addInfoOffset = builder.CreateString(t.AddInfo)
	}
	CapabilityStart(builder)
	CapabilityAddCount(builder, t.Count)
	CapabilityAddMainDiag(builder, t.MainDiag)
	CapabilityAddDetailDiag(builder, t.DetailDiag)
	CapabilityAddAddInfo(builder, addInfoOffset)
	return CapabilityEnd(builder)
}

func (rcv *Capability) UnPackTo(t *CapabilityT) {
	t.Count = rcv.Count()
	t.MainDiag = rcv.MainDiag()
	t.DetailDiag = rcv.DetailDiag()
	t.AddInfo = string(rcv.AddInfo())
}

func (rcv *Capability) UnPack() *CapabilityT {
	if rcv == nil { return nil }
	t := &CapabilityT{}
	rcv.UnPackTo(t)
	return t
}

type Capability struct {
	_tab flatbuffers.Table
}

func GetRootAsCapability(buf []byte, offset flatbuffers.UOffsetT) *Capability {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Capability{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCapability(buf []byte, offset flatbuffers.UOffsetT) *Capability {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Capability{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Capability) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Capability) Table() flatbuffers.Table {
	return rcv._tab
}

/// how many items are allowed by the system or the ressource? (when true/false then 1 represents true and 0 represents false)
func (rcv *Capability) Count() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// how many items are allowed by the system or the ressource? (when true/false then 1 represents true and 0 represents false)
func (rcv *Capability) MutateCount(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

/// in case of count==0: get here the main diagnosis code, why the capability is missing
func (rcv *Capability) MainDiag() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// in case of count==0: get here the main diagnosis code, why the capability is missing
func (rcv *Capability) MutateMainDiag(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

/// in case of count==0: get here the detail diagnosis code, why the capability is missing
func (rcv *Capability) DetailDiag() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// in case of count==0: get here the detail diagnosis code, why the capability is missing
func (rcv *Capability) MutateDetailDiag(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

/// in case of count==0: get here additional information text, why the capability is missing
func (rcv *Capability) AddInfo() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// in case of count==0: get here additional information text, why the capability is missing
func CapabilityStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func CapabilityAddCount(builder *flatbuffers.Builder, count uint32) {
	builder.PrependUint32Slot(0, count, 0)
}
func CapabilityAddMainDiag(builder *flatbuffers.Builder, mainDiag uint32) {
	builder.PrependUint32Slot(1, mainDiag, 0)
}
func CapabilityAddDetailDiag(builder *flatbuffers.Builder, detailDiag uint32) {
	builder.PrependUint32Slot(2, detailDiag, 0)
}
func CapabilityAddAddInfo(builder *flatbuffers.Builder, addInfo flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(addInfo), 0)
}
func CapabilityEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
