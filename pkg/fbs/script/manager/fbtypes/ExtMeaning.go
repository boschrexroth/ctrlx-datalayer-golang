// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// One pair with value and meaning for the internal kinematics pose (used for commanding)
type ExtMeaningT struct {
	Meaning Meaning `json:"meaning"`
	AddAttributes AddAttributes `json:"addAttributes"`
}

func (t *ExtMeaningT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ExtMeaningStart(builder)
	ExtMeaningAddMeaning(builder, t.Meaning)
	ExtMeaningAddAddAttributes(builder, t.AddAttributes)
	return ExtMeaningEnd(builder)
}

func (rcv *ExtMeaning) UnPackTo(t *ExtMeaningT) {
	t.Meaning = rcv.Meaning()
	t.AddAttributes = rcv.AddAttributes()
}

func (rcv *ExtMeaning) UnPack() *ExtMeaningT {
	if rcv == nil { return nil }
	t := &ExtMeaningT{}
	rcv.UnPackTo(t)
	return t
}

type ExtMeaning struct {
	_tab flatbuffers.Table
}

func GetRootAsExtMeaning(buf []byte, offset flatbuffers.UOffsetT) *ExtMeaning {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ExtMeaning{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsExtMeaning(buf []byte, offset flatbuffers.UOffsetT) *ExtMeaning {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ExtMeaning{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ExtMeaning) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ExtMeaning) Table() flatbuffers.Table {
	return rcv._tab
}

/// related meaning of the value
/// possible meanings are: "X", "Y", "Z", "AX1", "AX2", "AX3", "AX4", "AX5", "AX6", "AX7", "AX8", "AX9", "AX10", "AX11" "AX12", "AX13", "AX14", "AX15", "AX16", "EUL1_ZYZ_I", "EUL2_ZYZ_I", "EUL3_ZYZ_I", "EUL1_XYZ_E", "EUL2_XYZ_E", "EUL3_XYZ_E", "EUL1_ZYX_E", "EUL2_ZYX_E", "EUL3_ZYX_E"
func (rcv *ExtMeaning) Meaning() Meaning {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Meaning(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// related meaning of the value
/// possible meanings are: "X", "Y", "Z", "AX1", "AX2", "AX3", "AX4", "AX5", "AX6", "AX7", "AX8", "AX9", "AX10", "AX11" "AX12", "AX13", "AX14", "AX15", "AX16", "EUL1_ZYZ_I", "EUL2_ZYZ_I", "EUL3_ZYZ_I", "EUL1_XYZ_E", "EUL2_XYZ_E", "EUL3_XYZ_E", "EUL1_ZYX_E", "EUL2_ZYX_E", "EUL3_ZYX_E"
func (rcv *ExtMeaning) MutateMeaning(n Meaning) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

///  possible attributes are: "UNDEF", "INCR_VALUE", "ABS_VALUE", "SHORT_WAY", "POS_ROT_DIR", "NEG_ROT_DIR"
func (rcv *ExtMeaning) AddAttributes() AddAttributes {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return AddAttributes(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

///  possible attributes are: "UNDEF", "INCR_VALUE", "ABS_VALUE", "SHORT_WAY", "POS_ROT_DIR", "NEG_ROT_DIR"
func (rcv *ExtMeaning) MutateAddAttributes(n AddAttributes) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

func ExtMeaningStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ExtMeaningAddMeaning(builder *flatbuffers.Builder, meaning Meaning) {
	builder.PrependInt8Slot(0, int8(meaning), 0)
}
func ExtMeaningAddAddAttributes(builder *flatbuffers.Builder, addAttributes AddAttributes) {
	builder.PrependInt8Slot(1, int8(addAttributes), 0)
}
func ExtMeaningEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
