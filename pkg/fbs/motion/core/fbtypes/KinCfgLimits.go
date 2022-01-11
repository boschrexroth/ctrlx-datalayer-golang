// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// all configured limits of this kinematics
type KinCfgLimits struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCfgLimits(buf []byte, offset flatbuffers.UOffsetT) *KinCfgLimits {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCfgLimits{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCfgLimits(buf []byte, offset flatbuffers.UOffsetT) *KinCfgLimits {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCfgLimits{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCfgLimits) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCfgLimits) Table() flatbuffers.Table {
	return rcv._tab
}

/// path velocity limit (should be greater than zero)
func (rcv *KinCfgLimits) Vel() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// path velocity limit (should be greater than zero)
func (rcv *KinCfgLimits) MutateVel(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// path acceleration limit (should be greater than zero)
func (rcv *KinCfgLimits) Acc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// path acceleration limit (should be greater than zero)
func (rcv *KinCfgLimits) MutateAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// path deceleration limit (should be greater than zero)
func (rcv *KinCfgLimits) Dec() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// path deceleration limit (should be greater than zero)
func (rcv *KinCfgLimits) MutateDec(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

/// path acceleration jerk limit (should be greater than zero OR zero for not jerk limited motion)
func (rcv *KinCfgLimits) JrkAcc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// path acceleration jerk limit (should be greater than zero OR zero for not jerk limited motion)
func (rcv *KinCfgLimits) MutateJrkAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

/// path deceleration jerk limit (should be greater than zero OR zero for not jerk limited motion)
func (rcv *KinCfgLimits) JrkDec() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// path deceleration jerk limit (should be greater than zero OR zero for not jerk limited motion)
func (rcv *KinCfgLimits) MutateJrkDec(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

/// unit of vel
func (rcv *KinCfgLimits) VelUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of vel
/// unit of acc
func (rcv *KinCfgLimits) AccUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of acc
/// unit of dec
func (rcv *KinCfgLimits) DecUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of dec
/// unit of jrkAcc
func (rcv *KinCfgLimits) JrkAccUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of jrkAcc
/// unit of jrkDec
func (rcv *KinCfgLimits) JrkDecUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of jrkDec
func KinCfgLimitsStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func KinCfgLimitsAddVel(builder *flatbuffers.Builder, vel float64) {
	builder.PrependFloat64Slot(0, vel, 0.0)
}
func KinCfgLimitsAddAcc(builder *flatbuffers.Builder, acc float64) {
	builder.PrependFloat64Slot(1, acc, 0.0)
}
func KinCfgLimitsAddDec(builder *flatbuffers.Builder, dec float64) {
	builder.PrependFloat64Slot(2, dec, 0.0)
}
func KinCfgLimitsAddJrkAcc(builder *flatbuffers.Builder, jrkAcc float64) {
	builder.PrependFloat64Slot(3, jrkAcc, 0.0)
}
func KinCfgLimitsAddJrkDec(builder *flatbuffers.Builder, jrkDec float64) {
	builder.PrependFloat64Slot(4, jrkDec, 0.0)
}
func KinCfgLimitsAddVelUnit(builder *flatbuffers.Builder, velUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(velUnit), 0)
}
func KinCfgLimitsAddAccUnit(builder *flatbuffers.Builder, accUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(accUnit), 0)
}
func KinCfgLimitsAddDecUnit(builder *flatbuffers.Builder, decUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(decUnit), 0)
}
func KinCfgLimitsAddJrkAccUnit(builder *flatbuffers.Builder, jrkAccUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(jrkAccUnit), 0)
}
func KinCfgLimitsAddJrkDecUnit(builder *flatbuffers.Builder, jrkDecUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(jrkDecUnit), 0)
}
func KinCfgLimitsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}