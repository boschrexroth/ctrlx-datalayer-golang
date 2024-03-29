// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters for the jog commands for a kinematics
type KinCmdJogDataT struct {
	JogDir []float64 `json:"jogDir"`
	CoordSys string `json:"coordSys"`
	JogIncrement float64 `json:"jogIncrement"`
	Lim *DynamicLimitsT `json:"lim"`
}

func (t *KinCmdJogDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	jogDirOffset := flatbuffers.UOffsetT(0)
	if t.JogDir != nil {
		jogDirLength := len(t.JogDir)
		KinCmdJogDataStartJogDirVector(builder, jogDirLength)
		for j := jogDirLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.JogDir[j])
		}
		jogDirOffset = builder.EndVector(jogDirLength)
	}
	coordSysOffset := flatbuffers.UOffsetT(0)
	if t.CoordSys != "" {
		coordSysOffset = builder.CreateString(t.CoordSys)
	}
	limOffset := t.Lim.Pack(builder)
	KinCmdJogDataStart(builder)
	KinCmdJogDataAddJogDir(builder, jogDirOffset)
	KinCmdJogDataAddCoordSys(builder, coordSysOffset)
	KinCmdJogDataAddJogIncrement(builder, t.JogIncrement)
	KinCmdJogDataAddLim(builder, limOffset)
	return KinCmdJogDataEnd(builder)
}

func (rcv *KinCmdJogData) UnPackTo(t *KinCmdJogDataT) {
	jogDirLength := rcv.JogDirLength()
	t.JogDir = make([]float64, jogDirLength)
	for j := 0; j < jogDirLength; j++ {
		t.JogDir[j] = rcv.JogDir(j)
	}
	t.CoordSys = string(rcv.CoordSys())
	t.JogIncrement = rcv.JogIncrement()
	t.Lim = rcv.Lim(nil).UnPack()
}

func (rcv *KinCmdJogData) UnPack() *KinCmdJogDataT {
	if rcv == nil { return nil }
	t := &KinCmdJogDataT{}
	rcv.UnPackTo(t)
	return t
}

type KinCmdJogData struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCmdJogData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdJogData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCmdJogData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCmdJogData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdJogData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCmdJogData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCmdJogData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCmdJogData) Table() flatbuffers.Table {
	return rcv._tab
}

/// jog direction as a vector
func (rcv *KinCmdJogData) JogDir(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *KinCmdJogData) JogDirLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// jog direction as a vector
func (rcv *KinCmdJogData) MutateJogDir(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

/// coordSys for jog direction vector
func (rcv *KinCmdJogData) CoordSys() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// coordSys for jog direction vector
/// jog increment (must be zero for continuous jogging, must be a positive value for incremantal jogging)
func (rcv *KinCmdJogData) JogIncrement() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// jog increment (must be zero for continuous jogging, must be a positive value for incremantal jogging)
func (rcv *KinCmdJogData) MutateJogIncrement(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

/// dynamic limits for the motion of this command
func (rcv *KinCmdJogData) Lim(obj *DynamicLimits) *DynamicLimits {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(DynamicLimits)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// dynamic limits for the motion of this command
func KinCmdJogDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func KinCmdJogDataAddJogDir(builder *flatbuffers.Builder, jogDir flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(jogDir), 0)
}
func KinCmdJogDataStartJogDirVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func KinCmdJogDataAddCoordSys(builder *flatbuffers.Builder, coordSys flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(coordSys), 0)
}
func KinCmdJogDataAddJogIncrement(builder *flatbuffers.Builder, jogIncrement float64) {
	builder.PrependFloat64Slot(2, jogIncrement, 0.0)
}
func KinCmdJogDataAddLim(builder *flatbuffers.Builder, lim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(lim), 0)
}
func KinCmdJogDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
