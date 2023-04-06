// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package licenseproxy

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LicenseCapabilityT struct {
	Name string
	Version string
	Count int32
	IsPermanent bool
	StartDate string
	FinalExpirationDate string
}

func (t *LicenseCapabilityT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	versionOffset := builder.CreateString(t.Version)
	startDateOffset := builder.CreateString(t.StartDate)
	finalExpirationDateOffset := builder.CreateString(t.FinalExpirationDate)
	LicenseCapabilityStart(builder)
	LicenseCapabilityAddName(builder, nameOffset)
	LicenseCapabilityAddVersion(builder, versionOffset)
	LicenseCapabilityAddCount(builder, t.Count)
	LicenseCapabilityAddIsPermanent(builder, t.IsPermanent)
	LicenseCapabilityAddStartDate(builder, startDateOffset)
	LicenseCapabilityAddFinalExpirationDate(builder, finalExpirationDateOffset)
	return LicenseCapabilityEnd(builder)
}

func (rcv *LicenseCapability) UnPackTo(t *LicenseCapabilityT) {
	t.Name = string(rcv.Name())
	t.Version = string(rcv.Version())
	t.Count = rcv.Count()
	t.IsPermanent = rcv.IsPermanent()
	t.StartDate = string(rcv.StartDate())
	t.FinalExpirationDate = string(rcv.FinalExpirationDate())
}

func (rcv *LicenseCapability) UnPack() *LicenseCapabilityT {
	if rcv == nil { return nil }
	t := &LicenseCapabilityT{}
	rcv.UnPackTo(t)
	return t
}

type LicenseCapability struct {
	_tab flatbuffers.Table
}

func GetRootAsLicenseCapability(buf []byte, offset flatbuffers.UOffsetT) *LicenseCapability {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LicenseCapability{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLicenseCapability(buf []byte, offset flatbuffers.UOffsetT) *LicenseCapability {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LicenseCapability{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LicenseCapability) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LicenseCapability) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LicenseCapability) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LicenseCapability) Version() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LicenseCapability) Count() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LicenseCapability) MutateCount(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *LicenseCapability) IsPermanent() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *LicenseCapability) MutateIsPermanent(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func (rcv *LicenseCapability) StartDate() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LicenseCapability) FinalExpirationDate() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func LicenseCapabilityStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func LicenseCapabilityAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func LicenseCapabilityAddVersion(builder *flatbuffers.Builder, version flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(version), 0)
}
func LicenseCapabilityAddCount(builder *flatbuffers.Builder, count int32) {
	builder.PrependInt32Slot(2, count, 0)
}
func LicenseCapabilityAddIsPermanent(builder *flatbuffers.Builder, isPermanent bool) {
	builder.PrependBoolSlot(3, isPermanent, false)
}
func LicenseCapabilityAddStartDate(builder *flatbuffers.Builder, startDate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(startDate), 0)
}
func LicenseCapabilityAddFinalExpirationDate(builder *flatbuffers.Builder, finalExpirationDate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(finalExpirationDate), 0)
}
func LicenseCapabilityEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
