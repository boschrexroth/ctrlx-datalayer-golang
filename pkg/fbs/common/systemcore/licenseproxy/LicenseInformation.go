// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package licenseproxy

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LicenseInformationT struct {
	Mode string `json:"mode"`
	RemainingTime int32 `json:"remainingTime"`
}

func (t *LicenseInformationT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	modeOffset := flatbuffers.UOffsetT(0)
	if t.Mode != "" {
		modeOffset = builder.CreateString(t.Mode)
	}
	LicenseInformationStart(builder)
	LicenseInformationAddMode(builder, modeOffset)
	LicenseInformationAddRemainingTime(builder, t.RemainingTime)
	return LicenseInformationEnd(builder)
}

func (rcv *LicenseInformation) UnPackTo(t *LicenseInformationT) {
	t.Mode = string(rcv.Mode())
	t.RemainingTime = rcv.RemainingTime()
}

func (rcv *LicenseInformation) UnPack() *LicenseInformationT {
	if rcv == nil { return nil }
	t := &LicenseInformationT{}
	rcv.UnPackTo(t)
	return t
}

type LicenseInformation struct {
	_tab flatbuffers.Table
}

func GetRootAsLicenseInformation(buf []byte, offset flatbuffers.UOffsetT) *LicenseInformation {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LicenseInformation{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLicenseInformation(buf []byte, offset flatbuffers.UOffsetT) *LicenseInformation {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LicenseInformation{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LicenseInformation) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LicenseInformation) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LicenseInformation) Mode() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LicenseInformation) RemainingTime() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LicenseInformation) MutateRemainingTime(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func LicenseInformationStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func LicenseInformationAddMode(builder *flatbuffers.Builder, mode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(mode), 0)
}
func LicenseInformationAddRemainingTime(builder *flatbuffers.Builder, remainingTime int32) {
	builder.PrependInt32Slot(1, remainingTime, 0)
}
func LicenseInformationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
