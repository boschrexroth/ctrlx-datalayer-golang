// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SafeAreasExecErrReactionSettingsT struct {
	Type SafeAreasExecErrReaction `json:"type"`
}

func (t *SafeAreasExecErrReactionSettingsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	SafeAreasExecErrReactionSettingsStart(builder)
	SafeAreasExecErrReactionSettingsAddType(builder, t.Type)
	return SafeAreasExecErrReactionSettingsEnd(builder)
}

func (rcv *SafeAreasExecErrReactionSettings) UnPackTo(t *SafeAreasExecErrReactionSettingsT) {
	t.Type = rcv.Type()
}

func (rcv *SafeAreasExecErrReactionSettings) UnPack() *SafeAreasExecErrReactionSettingsT {
	if rcv == nil { return nil }
	t := &SafeAreasExecErrReactionSettingsT{}
	rcv.UnPackTo(t)
	return t
}

type SafeAreasExecErrReactionSettings struct {
	_tab flatbuffers.Table
}

func GetRootAsSafeAreasExecErrReactionSettings(buf []byte, offset flatbuffers.UOffsetT) *SafeAreasExecErrReactionSettings {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SafeAreasExecErrReactionSettings{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSafeAreasExecErrReactionSettings(buf []byte, offset flatbuffers.UOffsetT) *SafeAreasExecErrReactionSettings {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SafeAreasExecErrReactionSettings{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SafeAreasExecErrReactionSettings) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SafeAreasExecErrReactionSettings) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SafeAreasExecErrReactionSettings) Type() SafeAreasExecErrReaction {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return SafeAreasExecErrReaction(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 1
}

func (rcv *SafeAreasExecErrReactionSettings) MutateType(n SafeAreasExecErrReaction) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func SafeAreasExecErrReactionSettingsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SafeAreasExecErrReactionSettingsAddType(builder *flatbuffers.Builder, type_ SafeAreasExecErrReaction) {
	builder.PrependInt8Slot(0, int8(type_), 1)
}
func SafeAreasExecErrReactionSettingsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
