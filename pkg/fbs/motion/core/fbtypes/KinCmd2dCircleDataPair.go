// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// One pair with value and meaning for the internal additional circle data (used for commanding)
type KinCmd2dCircleDataPairT struct {
	Value float64 `json:"value"`
	ExtMeaning *Ext2dCircleMeaningT `json:"extMeaning"`
}

func (t *KinCmd2dCircleDataPairT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	extMeaningOffset := t.ExtMeaning.Pack(builder)
	KinCmd2dCircleDataPairStart(builder)
	KinCmd2dCircleDataPairAddValue(builder, t.Value)
	KinCmd2dCircleDataPairAddExtMeaning(builder, extMeaningOffset)
	return KinCmd2dCircleDataPairEnd(builder)
}

func (rcv *KinCmd2dCircleDataPair) UnPackTo(t *KinCmd2dCircleDataPairT) {
	t.Value = rcv.Value()
	t.ExtMeaning = rcv.ExtMeaning(nil).UnPack()
}

func (rcv *KinCmd2dCircleDataPair) UnPack() *KinCmd2dCircleDataPairT {
	if rcv == nil { return nil }
	t := &KinCmd2dCircleDataPairT{}
	rcv.UnPackTo(t)
	return t
}

type KinCmd2dCircleDataPair struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCmd2dCircleDataPair(buf []byte, offset flatbuffers.UOffsetT) *KinCmd2dCircleDataPair {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCmd2dCircleDataPair{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCmd2dCircleDataPair(buf []byte, offset flatbuffers.UOffsetT) *KinCmd2dCircleDataPair {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCmd2dCircleDataPair{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCmd2dCircleDataPair) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCmd2dCircleDataPair) Table() flatbuffers.Table {
	return rcv._tab
}

/// one value of a circle data
func (rcv *KinCmd2dCircleDataPair) Value() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// one value of a circle data
func (rcv *KinCmd2dCircleDataPair) MutateValue(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// related meaning of the value
/// possible meanings are: "MP_X", "MP_Y", "MP_Z", "RADIUS", "REV"
func (rcv *KinCmd2dCircleDataPair) ExtMeaning(obj *Ext2dCircleMeaning) *Ext2dCircleMeaning {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Ext2dCircleMeaning)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// related meaning of the value
/// possible meanings are: "MP_X", "MP_Y", "MP_Z", "RADIUS", "REV"
func KinCmd2dCircleDataPairStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func KinCmd2dCircleDataPairAddValue(builder *flatbuffers.Builder, value float64) {
	builder.PrependFloat64Slot(0, value, 0.0)
}
func KinCmd2dCircleDataPairAddExtMeaning(builder *flatbuffers.Builder, extMeaning flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(extMeaning), 0)
}
func KinCmd2dCircleDataPairEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
