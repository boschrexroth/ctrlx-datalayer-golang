// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// set of dynamic limits (for commands)
type DynamicLimitsT struct {
	Vel float64
	Acc float64
	Dec float64
	JrkAcc float64
	JrkDec float64
}

func (t *DynamicLimitsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	DynamicLimitsStart(builder)
	DynamicLimitsAddVel(builder, t.Vel)
	DynamicLimitsAddAcc(builder, t.Acc)
	DynamicLimitsAddDec(builder, t.Dec)
	DynamicLimitsAddJrkAcc(builder, t.JrkAcc)
	DynamicLimitsAddJrkDec(builder, t.JrkDec)
	return DynamicLimitsEnd(builder)
}

func (rcv *DynamicLimits) UnPackTo(t *DynamicLimitsT) {
	t.Vel = rcv.Vel()
	t.Acc = rcv.Acc()
	t.Dec = rcv.Dec()
	t.JrkAcc = rcv.JrkAcc()
	t.JrkDec = rcv.JrkDec()
}

func (rcv *DynamicLimits) UnPack() *DynamicLimitsT {
	if rcv == nil { return nil }
	t := &DynamicLimitsT{}
	rcv.UnPackTo(t)
	return t
}

type DynamicLimits struct {
	_tab flatbuffers.Table
}

func GetRootAsDynamicLimits(buf []byte, offset flatbuffers.UOffsetT) *DynamicLimits {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DynamicLimits{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDynamicLimits(buf []byte, offset flatbuffers.UOffsetT) *DynamicLimits {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DynamicLimits{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DynamicLimits) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DynamicLimits) Table() flatbuffers.Table {
	return rcv._tab
}

/// velocity limit (must be greater than zero)
func (rcv *DynamicLimits) Vel() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// velocity limit (must be greater than zero)
func (rcv *DynamicLimits) MutateVel(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// acceleration limit (must be greater than zero)
func (rcv *DynamicLimits) Acc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// acceleration limit (must be greater than zero)
func (rcv *DynamicLimits) MutateAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// deceleration limit (must be greater than zero)
func (rcv *DynamicLimits) Dec() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// deceleration limit (must be greater than zero)
func (rcv *DynamicLimits) MutateDec(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

/// acceleration jerk limit (must be greater than zero OR zero for not jerk limited motion)
func (rcv *DynamicLimits) JrkAcc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// acceleration jerk limit (must be greater than zero OR zero for not jerk limited motion)
func (rcv *DynamicLimits) MutateJrkAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

/// deceleration jerk limit (must be greater than zero OR zero for not jerk limited motion)
func (rcv *DynamicLimits) JrkDec() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// deceleration jerk limit (must be greater than zero OR zero for not jerk limited motion)
func (rcv *DynamicLimits) MutateJrkDec(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

func DynamicLimitsStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func DynamicLimitsAddVel(builder *flatbuffers.Builder, vel float64) {
	builder.PrependFloat64Slot(0, vel, 0.0)
}
func DynamicLimitsAddAcc(builder *flatbuffers.Builder, acc float64) {
	builder.PrependFloat64Slot(1, acc, 0.0)
}
func DynamicLimitsAddDec(builder *flatbuffers.Builder, dec float64) {
	builder.PrependFloat64Slot(2, dec, 0.0)
}
func DynamicLimitsAddJrkAcc(builder *flatbuffers.Builder, jrkAcc float64) {
	builder.PrependFloat64Slot(3, jrkAcc, 0.0)
}
func DynamicLimitsAddJrkDec(builder *flatbuffers.Builder, jrkDec float64) {
	builder.PrependFloat64Slot(4, jrkDec, 0.0)
}
func DynamicLimitsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
