// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ChannelCfgT struct {
	Name string
	Source string
	Type ChannelTypeEnumFb
	Unit string
}

func (t *ChannelCfgT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	sourceOffset := builder.CreateString(t.Source)
	unitOffset := builder.CreateString(t.Unit)
	ChannelCfgStart(builder)
	ChannelCfgAddName(builder, nameOffset)
	ChannelCfgAddSource(builder, sourceOffset)
	ChannelCfgAddType(builder, t.Type)
	ChannelCfgAddUnit(builder, unitOffset)
	return ChannelCfgEnd(builder)
}

func (rcv *ChannelCfg) UnPackTo(t *ChannelCfgT) {
	t.Name = string(rcv.Name())
	t.Source = string(rcv.Source())
	t.Type = rcv.Type()
	t.Unit = string(rcv.Unit())
}

func (rcv *ChannelCfg) UnPack() *ChannelCfgT {
	if rcv == nil { return nil }
	t := &ChannelCfgT{}
	rcv.UnPackTo(t)
	return t
}

type ChannelCfg struct {
	_tab flatbuffers.Table
}

func GetRootAsChannelCfg(buf []byte, offset flatbuffers.UOffsetT) *ChannelCfg {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ChannelCfg{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsChannelCfg(buf []byte, offset flatbuffers.UOffsetT) *ChannelCfg {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ChannelCfg{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ChannelCfg) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ChannelCfg) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ChannelCfg) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ChannelCfg) Source() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ChannelCfg) Type() ChannelTypeEnumFb {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return ChannelTypeEnumFb(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ChannelCfg) MutateType(n ChannelTypeEnumFb) bool {
	return rcv._tab.MutateInt8Slot(8, int8(n))
}

func (rcv *ChannelCfg) Unit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ChannelCfgStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ChannelCfgAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func ChannelCfgAddSource(builder *flatbuffers.Builder, source flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(source), 0)
}
func ChannelCfgAddType(builder *flatbuffers.Builder, type_ ChannelTypeEnumFb) {
	builder.PrependInt8Slot(2, int8(type_), 0)
}
func ChannelCfgAddUnit(builder *flatbuffers.Builder, unit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(unit), 0)
}
func ChannelCfgEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
