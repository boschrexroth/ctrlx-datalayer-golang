// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Packed pair of target position meaning and attribute
type Ext3dCircleMeaningT struct {
	Meaning CircleMeaning3d `json:"meaning"`
	AddAttributes AddAttributes `json:"addAttributes"`
}

func (t *Ext3dCircleMeaningT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Ext3dCircleMeaningStart(builder)
	Ext3dCircleMeaningAddMeaning(builder, t.Meaning)
	Ext3dCircleMeaningAddAddAttributes(builder, t.AddAttributes)
	return Ext3dCircleMeaningEnd(builder)
}

func (rcv *Ext3dCircleMeaning) UnPackTo(t *Ext3dCircleMeaningT) {
	t.Meaning = rcv.Meaning()
	t.AddAttributes = rcv.AddAttributes()
}

func (rcv *Ext3dCircleMeaning) UnPack() *Ext3dCircleMeaningT {
	if rcv == nil { return nil }
	t := &Ext3dCircleMeaningT{}
	rcv.UnPackTo(t)
	return t
}

type Ext3dCircleMeaning struct {
	_tab flatbuffers.Table
}

func GetRootAsExt3dCircleMeaning(buf []byte, offset flatbuffers.UOffsetT) *Ext3dCircleMeaning {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Ext3dCircleMeaning{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsExt3dCircleMeaning(buf []byte, offset flatbuffers.UOffsetT) *Ext3dCircleMeaning {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Ext3dCircleMeaning{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Ext3dCircleMeaning) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Ext3dCircleMeaning) Table() flatbuffers.Table {
	return rcv._tab
}

/// related meaning of the value
/// possible meanings are: "IP_X", "IP_Y", "IP_Z"
func (rcv *Ext3dCircleMeaning) Meaning() CircleMeaning3d {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return CircleMeaning3d(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// related meaning of the value
/// possible meanings are: "IP_X", "IP_Y", "IP_Z"
func (rcv *Ext3dCircleMeaning) MutateMeaning(n CircleMeaning3d) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

/// attributes for the target position meaning
/// possible meanings are: "UNDEF", "INCR_VALUE", "ABS_VALUE", "SHORT_WAY", "POS_ROT_DIR", "NEG_ROT_DIR"
func (rcv *Ext3dCircleMeaning) AddAttributes() AddAttributes {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return AddAttributes(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// attributes for the target position meaning
/// possible meanings are: "UNDEF", "INCR_VALUE", "ABS_VALUE", "SHORT_WAY", "POS_ROT_DIR", "NEG_ROT_DIR"
func (rcv *Ext3dCircleMeaning) MutateAddAttributes(n AddAttributes) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

func Ext3dCircleMeaningStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Ext3dCircleMeaningAddMeaning(builder *flatbuffers.Builder, meaning CircleMeaning3d) {
	builder.PrependInt8Slot(0, int8(meaning), 0)
}
func Ext3dCircleMeaningAddAddAttributes(builder *flatbuffers.Builder, addAttributes AddAttributes) {
	builder.PrependInt8Slot(1, int8(addAttributes), 0)
}
func Ext3dCircleMeaningEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
