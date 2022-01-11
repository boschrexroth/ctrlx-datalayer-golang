// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configured limits of a single axis
type AxsCfgLimits struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgLimits(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgLimits {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgLimits{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgLimits(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgLimits {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgLimits{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgLimits) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgLimits) Table() flatbuffers.Table {
	return rcv._tab
}

/// minimum position limit (should be less than maximum position limit)
func (rcv *AxsCfgLimits) PosMin() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// minimum position limit (should be less than maximum position limit)
func (rcv *AxsCfgLimits) MutatePosMin(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// maximum position limit (should be greater than minimum position limit)
func (rcv *AxsCfgLimits) PosMax() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// maximum position limit (should be greater than minimum position limit)
func (rcv *AxsCfgLimits) MutatePosMax(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// absolute minimum velocity limit (should be greater than zero)
func (rcv *AxsCfgLimits) VelPos() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// absolute minimum velocity limit (should be greater than zero)
func (rcv *AxsCfgLimits) MutateVelPos(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

/// absolute maximum velocity limit (should be greater than zero)
func (rcv *AxsCfgLimits) VelNeg() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// absolute maximum velocity limit (should be greater than zero)
func (rcv *AxsCfgLimits) MutateVelNeg(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

/// absolute acceleration limit (should be greater than zero)
func (rcv *AxsCfgLimits) Acc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// absolute acceleration limit (should be greater than zero)
func (rcv *AxsCfgLimits) MutateAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

/// absolute deceleration limit (should be greater than zero)
func (rcv *AxsCfgLimits) Dec() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// absolute deceleration limit (should be greater than zero)
func (rcv *AxsCfgLimits) MutateDec(n float64) bool {
	return rcv._tab.MutateFloat64Slot(14, n)
}

/// absolute acceleration jerk limit (should be greater than zero OR zero for not jerk limited motion)
func (rcv *AxsCfgLimits) JrkAcc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// absolute acceleration jerk limit (should be greater than zero OR zero for not jerk limited motion)
func (rcv *AxsCfgLimits) MutateJrkAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(16, n)
}

/// absolute deceleration jerk limit (should be greater than zero OR zero for not jerk limited motion)
func (rcv *AxsCfgLimits) JrkDec() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// absolute deceleration jerk limit (should be greater than zero OR zero for not jerk limited motion)
func (rcv *AxsCfgLimits) MutateJrkDec(n float64) bool {
	return rcv._tab.MutateFloat64Slot(18, n)
}

/// unit of posMin
func (rcv *AxsCfgLimits) PosMinUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of posMin
/// unit of posMax
func (rcv *AxsCfgLimits) PosMaxUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of posMax
/// unit of velPos
func (rcv *AxsCfgLimits) VelPosUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of velPos
/// unit of velNeg
func (rcv *AxsCfgLimits) VelNegUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of velNeg
/// unit of acc
func (rcv *AxsCfgLimits) AccUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of acc
/// unit of dec
func (rcv *AxsCfgLimits) DecUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of dec
/// unit of jrkAcc
func (rcv *AxsCfgLimits) JrkAccUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of jrkAcc
/// unit of jrkDec
func (rcv *AxsCfgLimits) JrkDecUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of jrkDec
func AxsCfgLimitsStart(builder *flatbuffers.Builder) {
	builder.StartObject(16)
}
func AxsCfgLimitsAddPosMin(builder *flatbuffers.Builder, posMin float64) {
	builder.PrependFloat64Slot(0, posMin, 0.0)
}
func AxsCfgLimitsAddPosMax(builder *flatbuffers.Builder, posMax float64) {
	builder.PrependFloat64Slot(1, posMax, 0.0)
}
func AxsCfgLimitsAddVelPos(builder *flatbuffers.Builder, velPos float64) {
	builder.PrependFloat64Slot(2, velPos, 0.0)
}
func AxsCfgLimitsAddVelNeg(builder *flatbuffers.Builder, velNeg float64) {
	builder.PrependFloat64Slot(3, velNeg, 0.0)
}
func AxsCfgLimitsAddAcc(builder *flatbuffers.Builder, acc float64) {
	builder.PrependFloat64Slot(4, acc, 0.0)
}
func AxsCfgLimitsAddDec(builder *flatbuffers.Builder, dec float64) {
	builder.PrependFloat64Slot(5, dec, 0.0)
}
func AxsCfgLimitsAddJrkAcc(builder *flatbuffers.Builder, jrkAcc float64) {
	builder.PrependFloat64Slot(6, jrkAcc, 0.0)
}
func AxsCfgLimitsAddJrkDec(builder *flatbuffers.Builder, jrkDec float64) {
	builder.PrependFloat64Slot(7, jrkDec, 0.0)
}
func AxsCfgLimitsAddPosMinUnit(builder *flatbuffers.Builder, posMinUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(posMinUnit), 0)
}
func AxsCfgLimitsAddPosMaxUnit(builder *flatbuffers.Builder, posMaxUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(posMaxUnit), 0)
}
func AxsCfgLimitsAddVelPosUnit(builder *flatbuffers.Builder, velPosUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(velPosUnit), 0)
}
func AxsCfgLimitsAddVelNegUnit(builder *flatbuffers.Builder, velNegUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(velNegUnit), 0)
}
func AxsCfgLimitsAddAccUnit(builder *flatbuffers.Builder, accUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(accUnit), 0)
}
func AxsCfgLimitsAddDecUnit(builder *flatbuffers.Builder, decUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(decUnit), 0)
}
func AxsCfgLimitsAddJrkAccUnit(builder *flatbuffers.Builder, jrkAccUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(14, flatbuffers.UOffsetT(jrkAccUnit), 0)
}
func AxsCfgLimitsAddJrkDecUnit(builder *flatbuffers.Builder, jrkDecUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(15, flatbuffers.UOffsetT(jrkDecUnit), 0)
}
func AxsCfgLimitsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}