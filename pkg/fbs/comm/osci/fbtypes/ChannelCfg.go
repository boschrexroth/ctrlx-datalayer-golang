// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ChannelCfgT struct {
	Name string `json:"name"`
	Alias string `json:"alias"`
	Source string `json:"source"`
	DataType string `json:"dataType"`
	Type ChannelTypeEnumFb `json:"type"`
	Unit string `json:"unit"`
}

func (t *ChannelCfgT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	aliasOffset := flatbuffers.UOffsetT(0)
	if t.Alias != "" {
		aliasOffset = builder.CreateString(t.Alias)
	}
	sourceOffset := flatbuffers.UOffsetT(0)
	if t.Source != "" {
		sourceOffset = builder.CreateString(t.Source)
	}
	dataTypeOffset := flatbuffers.UOffsetT(0)
	if t.DataType != "" {
		dataTypeOffset = builder.CreateString(t.DataType)
	}
	unitOffset := flatbuffers.UOffsetT(0)
	if t.Unit != "" {
		unitOffset = builder.CreateString(t.Unit)
	}
	ChannelCfgStart(builder)
	ChannelCfgAddName(builder, nameOffset)
	ChannelCfgAddAlias(builder, aliasOffset)
	ChannelCfgAddSource(builder, sourceOffset)
	ChannelCfgAddDataType(builder, dataTypeOffset)
	ChannelCfgAddType(builder, t.Type)
	ChannelCfgAddUnit(builder, unitOffset)
	return ChannelCfgEnd(builder)
}

func (rcv *ChannelCfg) UnPackTo(t *ChannelCfgT) {
	t.Name = string(rcv.Name())
	t.Alias = string(rcv.Alias())
	t.Source = string(rcv.Source())
	t.DataType = string(rcv.DataType())
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

func (rcv *ChannelCfg) Alias() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ChannelCfg) Source() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ChannelCfg) DataType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ChannelCfg) Type() ChannelTypeEnumFb {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return ChannelTypeEnumFb(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ChannelCfg) MutateType(n ChannelTypeEnumFb) bool {
	return rcv._tab.MutateInt8Slot(12, int8(n))
}

func (rcv *ChannelCfg) Unit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ChannelCfgStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func ChannelCfgAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func ChannelCfgAddAlias(builder *flatbuffers.Builder, alias flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(alias), 0)
}
func ChannelCfgAddSource(builder *flatbuffers.Builder, source flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(source), 0)
}
func ChannelCfgAddDataType(builder *flatbuffers.Builder, dataType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(dataType), 0)
}
func ChannelCfgAddType(builder *flatbuffers.Builder, type_ ChannelTypeEnumFb) {
	builder.PrependInt8Slot(4, int8(type_), 0)
}
func ChannelCfgAddUnit(builder *flatbuffers.Builder, unit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(unit), 0)
}
func ChannelCfgEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
