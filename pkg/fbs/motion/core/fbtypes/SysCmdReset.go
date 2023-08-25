// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// request structure for the ResetAllMotionObjects nodes
type SysCmdResetT struct {
	Type SysResetType `json:"type"`
}

func (t *SysCmdResetT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	SysCmdResetStart(builder)
	SysCmdResetAddType(builder, t.Type)
	return SysCmdResetEnd(builder)
}

func (rcv *SysCmdReset) UnPackTo(t *SysCmdResetT) {
	t.Type = rcv.Type()
}

func (rcv *SysCmdReset) UnPack() *SysCmdResetT {
	if rcv == nil { return nil }
	t := &SysCmdResetT{}
	rcv.UnPackTo(t)
	return t
}

type SysCmdReset struct {
	_tab flatbuffers.Table
}

func GetRootAsSysCmdReset(buf []byte, offset flatbuffers.UOffsetT) *SysCmdReset {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SysCmdReset{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSysCmdReset(buf []byte, offset flatbuffers.UOffsetT) *SysCmdReset {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SysCmdReset{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SysCmdReset) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SysCmdReset) Table() flatbuffers.Table {
	return rcv._tab
}

/// type of the system reset request
func (rcv *SysCmdReset) Type() SysResetType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return SysResetType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// type of the system reset request
func (rcv *SysCmdReset) MutateType(n SysResetType) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func SysCmdResetStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SysCmdResetAddType(builder *flatbuffers.Builder, type_ SysResetType) {
	builder.PrependInt8Slot(0, int8(type_), 0)
}
func SysCmdResetEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
