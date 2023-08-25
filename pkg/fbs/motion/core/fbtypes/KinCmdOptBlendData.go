// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters for the command option blending for kinematics
type KinCmdOptBlendDataT struct {
	PermType string `json:"permType"`
	Dist1 float64 `json:"dist1"`
	Dist2 float64 `json:"dist2"`
}

func (t *KinCmdOptBlendDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	permTypeOffset := flatbuffers.UOffsetT(0)
	if t.PermType != "" {
		permTypeOffset = builder.CreateString(t.PermType)
	}
	KinCmdOptBlendDataStart(builder)
	KinCmdOptBlendDataAddPermType(builder, permTypeOffset)
	KinCmdOptBlendDataAddDist1(builder, t.Dist1)
	KinCmdOptBlendDataAddDist2(builder, t.Dist2)
	return KinCmdOptBlendDataEnd(builder)
}

func (rcv *KinCmdOptBlendData) UnPackTo(t *KinCmdOptBlendDataT) {
	t.PermType = string(rcv.PermType())
	t.Dist1 = rcv.Dist1()
	t.Dist2 = rcv.Dist2()
}

func (rcv *KinCmdOptBlendData) UnPack() *KinCmdOptBlendDataT {
	if rcv == nil { return nil }
	t := &KinCmdOptBlendDataT{}
	rcv.UnPackTo(t)
	return t
}

type KinCmdOptBlendData struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCmdOptBlendData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdOptBlendData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCmdOptBlendData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCmdOptBlendData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdOptBlendData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCmdOptBlendData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCmdOptBlendData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCmdOptBlendData) Table() flatbuffers.Table {
	return rcv._tab
}

/// permanent type (e.g. "Once")
func (rcv *KinCmdOptBlendData) PermType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// permanent type (e.g. "Once")
/// distance D1 (refer to the manual, should be greater than zero)
func (rcv *KinCmdOptBlendData) Dist1() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 1.0
}

/// distance D1 (refer to the manual, should be greater than zero)
func (rcv *KinCmdOptBlendData) MutateDist1(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// distance D2 (refer to the manual, should be greater than zero)
func (rcv *KinCmdOptBlendData) Dist2() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 1.0
}

/// distance D2 (refer to the manual, should be greater than zero)
func (rcv *KinCmdOptBlendData) MutateDist2(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func KinCmdOptBlendDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func KinCmdOptBlendDataAddPermType(builder *flatbuffers.Builder, permType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(permType), 0)
}
func KinCmdOptBlendDataAddDist1(builder *flatbuffers.Builder, dist1 float64) {
	builder.PrependFloat64Slot(1, dist1, 1.0)
}
func KinCmdOptBlendDataAddDist2(builder *flatbuffers.Builder, dist2 float64) {
	builder.PrependFloat64Slot(2, dist2, 1.0)
}
func KinCmdOptBlendDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
