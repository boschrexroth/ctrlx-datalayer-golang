// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ToolDataEdgeOriSettingsT struct {
	Type ToolDataEdgeOri `json:"type"`
}

func (t *ToolDataEdgeOriSettingsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ToolDataEdgeOriSettingsStart(builder)
	ToolDataEdgeOriSettingsAddType(builder, t.Type)
	return ToolDataEdgeOriSettingsEnd(builder)
}

func (rcv *ToolDataEdgeOriSettings) UnPackTo(t *ToolDataEdgeOriSettingsT) {
	t.Type = rcv.Type()
}

func (rcv *ToolDataEdgeOriSettings) UnPack() *ToolDataEdgeOriSettingsT {
	if rcv == nil { return nil }
	t := &ToolDataEdgeOriSettingsT{}
	rcv.UnPackTo(t)
	return t
}

type ToolDataEdgeOriSettings struct {
	_tab flatbuffers.Table
}

func GetRootAsToolDataEdgeOriSettings(buf []byte, offset flatbuffers.UOffsetT) *ToolDataEdgeOriSettings {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ToolDataEdgeOriSettings{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsToolDataEdgeOriSettings(buf []byte, offset flatbuffers.UOffsetT) *ToolDataEdgeOriSettings {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ToolDataEdgeOriSettings{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ToolDataEdgeOriSettings) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ToolDataEdgeOriSettings) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ToolDataEdgeOriSettings) Type() ToolDataEdgeOri {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return ToolDataEdgeOri(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ToolDataEdgeOriSettings) MutateType(n ToolDataEdgeOri) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func ToolDataEdgeOriSettingsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ToolDataEdgeOriSettingsAddType(builder *flatbuffers.Builder, type_ ToolDataEdgeOri) {
	builder.PrependInt8Slot(0, int8(type_), 0)
}
func ToolDataEdgeOriSettingsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
