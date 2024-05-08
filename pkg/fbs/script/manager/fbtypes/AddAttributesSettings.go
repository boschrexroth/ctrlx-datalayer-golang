// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AddAttributesSettingsT struct {
	Type AddAttributes `json:"type"`
}

func (t *AddAttributesSettingsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	AddAttributesSettingsStart(builder)
	AddAttributesSettingsAddType(builder, t.Type)
	return AddAttributesSettingsEnd(builder)
}

func (rcv *AddAttributesSettings) UnPackTo(t *AddAttributesSettingsT) {
	t.Type = rcv.Type()
}

func (rcv *AddAttributesSettings) UnPack() *AddAttributesSettingsT {
	if rcv == nil { return nil }
	t := &AddAttributesSettingsT{}
	rcv.UnPackTo(t)
	return t
}

type AddAttributesSettings struct {
	_tab flatbuffers.Table
}

func GetRootAsAddAttributesSettings(buf []byte, offset flatbuffers.UOffsetT) *AddAttributesSettings {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AddAttributesSettings{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAddAttributesSettings(buf []byte, offset flatbuffers.UOffsetT) *AddAttributesSettings {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AddAttributesSettings{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AddAttributesSettings) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AddAttributesSettings) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AddAttributesSettings) Type() AddAttributes {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return AddAttributes(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *AddAttributesSettings) MutateType(n AddAttributes) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func AddAttributesSettingsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func AddAttributesSettingsAddType(builder *flatbuffers.Builder, type_ AddAttributes) {
	builder.PrependInt8Slot(0, int8(type_), 0)
}
func AddAttributesSettingsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}