// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Profile channel info response
type ProfileChannelInfoResponseT struct {
	ProfileNo uint16 `json:"ProfileNo"`
	AddInfo uint16 `json:"AddInfo"`
	DisplayName string `json:"DisplayName"`
}

func (t *ProfileChannelInfoResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	displayNameOffset := flatbuffers.UOffsetT(0)
	if t.DisplayName != "" {
		displayNameOffset = builder.CreateString(t.DisplayName)
	}
	ProfileChannelInfoResponseStart(builder)
	ProfileChannelInfoResponseAddProfileNo(builder, t.ProfileNo)
	ProfileChannelInfoResponseAddAddInfo(builder, t.AddInfo)
	ProfileChannelInfoResponseAddDisplayName(builder, displayNameOffset)
	return ProfileChannelInfoResponseEnd(builder)
}

func (rcv *ProfileChannelInfoResponse) UnPackTo(t *ProfileChannelInfoResponseT) {
	t.ProfileNo = rcv.ProfileNo()
	t.AddInfo = rcv.AddInfo()
	t.DisplayName = string(rcv.DisplayName())
}

func (rcv *ProfileChannelInfoResponse) UnPack() *ProfileChannelInfoResponseT {
	if rcv == nil { return nil }
	t := &ProfileChannelInfoResponseT{}
	rcv.UnPackTo(t)
	return t
}

type ProfileChannelInfoResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsProfileChannelInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *ProfileChannelInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ProfileChannelInfoResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsProfileChannelInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *ProfileChannelInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ProfileChannelInfoResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ProfileChannelInfoResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ProfileChannelInfoResponse) Table() flatbuffers.Table {
	return rcv._tab
}

///Profile number of this device, e.g. 5001 for Modular Device Profile
func (rcv *ProfileChannelInfoResponse) ProfileNo() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Profile number of this device, e.g. 5001 for Modular Device Profile
func (rcv *ProfileChannelInfoResponse) MutateProfileNo(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

///Additional profile info number of this channel, e.g. Channel Profile for Encoders
func (rcv *ProfileChannelInfoResponse) AddInfo() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Additional profile info number of this channel, e.g. Channel Profile for Encoders
func (rcv *ProfileChannelInfoResponse) MutateAddInfo(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

///Displayed name of this channel, e.g. Encoder 1
func (rcv *ProfileChannelInfoResponse) DisplayName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///Displayed name of this channel, e.g. Encoder 1
func ProfileChannelInfoResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ProfileChannelInfoResponseAddProfileNo(builder *flatbuffers.Builder, profileNo uint16) {
	builder.PrependUint16Slot(0, profileNo, 0)
}
func ProfileChannelInfoResponseAddAddInfo(builder *flatbuffers.Builder, addInfo uint16) {
	builder.PrependUint16Slot(1, addInfo, 0)
}
func ProfileChannelInfoResponseAddDisplayName(builder *flatbuffers.Builder, displayName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(displayName), 0)
}
func ProfileChannelInfoResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
