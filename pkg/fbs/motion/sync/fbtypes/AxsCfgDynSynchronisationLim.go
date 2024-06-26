// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Configuration of the dynamic synchronisation limits of a single axis
type AxsCfgDynSynchronisationLimT struct {
	Acc float64 `json:"acc"`
	Dec float64 `json:"dec"`
	JrkAcc float64 `json:"jrkAcc"`
	JrkDec float64 `json:"jrkDec"`
	VelNeg float64 `json:"velNeg"`
	VelPos float64 `json:"velPos"`
	AccUnit string `json:"accUnit"`
	DecUnit string `json:"decUnit"`
	JrkAccUnit string `json:"jrkAccUnit"`
	JrkDecUnit string `json:"jrkDecUnit"`
	VelNegUnit string `json:"velNegUnit"`
	VelPosUnit string `json:"velPosUnit"`
}

func (t *AxsCfgDynSynchronisationLimT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	accUnitOffset := flatbuffers.UOffsetT(0)
	if t.AccUnit != "" {
		accUnitOffset = builder.CreateString(t.AccUnit)
	}
	decUnitOffset := flatbuffers.UOffsetT(0)
	if t.DecUnit != "" {
		decUnitOffset = builder.CreateString(t.DecUnit)
	}
	jrkAccUnitOffset := flatbuffers.UOffsetT(0)
	if t.JrkAccUnit != "" {
		jrkAccUnitOffset = builder.CreateString(t.JrkAccUnit)
	}
	jrkDecUnitOffset := flatbuffers.UOffsetT(0)
	if t.JrkDecUnit != "" {
		jrkDecUnitOffset = builder.CreateString(t.JrkDecUnit)
	}
	velNegUnitOffset := flatbuffers.UOffsetT(0)
	if t.VelNegUnit != "" {
		velNegUnitOffset = builder.CreateString(t.VelNegUnit)
	}
	velPosUnitOffset := flatbuffers.UOffsetT(0)
	if t.VelPosUnit != "" {
		velPosUnitOffset = builder.CreateString(t.VelPosUnit)
	}
	AxsCfgDynSynchronisationLimStart(builder)
	AxsCfgDynSynchronisationLimAddAcc(builder, t.Acc)
	AxsCfgDynSynchronisationLimAddDec(builder, t.Dec)
	AxsCfgDynSynchronisationLimAddJrkAcc(builder, t.JrkAcc)
	AxsCfgDynSynchronisationLimAddJrkDec(builder, t.JrkDec)
	AxsCfgDynSynchronisationLimAddVelNeg(builder, t.VelNeg)
	AxsCfgDynSynchronisationLimAddVelPos(builder, t.VelPos)
	AxsCfgDynSynchronisationLimAddAccUnit(builder, accUnitOffset)
	AxsCfgDynSynchronisationLimAddDecUnit(builder, decUnitOffset)
	AxsCfgDynSynchronisationLimAddJrkAccUnit(builder, jrkAccUnitOffset)
	AxsCfgDynSynchronisationLimAddJrkDecUnit(builder, jrkDecUnitOffset)
	AxsCfgDynSynchronisationLimAddVelNegUnit(builder, velNegUnitOffset)
	AxsCfgDynSynchronisationLimAddVelPosUnit(builder, velPosUnitOffset)
	return AxsCfgDynSynchronisationLimEnd(builder)
}

func (rcv *AxsCfgDynSynchronisationLim) UnPackTo(t *AxsCfgDynSynchronisationLimT) {
	t.Acc = rcv.Acc()
	t.Dec = rcv.Dec()
	t.JrkAcc = rcv.JrkAcc()
	t.JrkDec = rcv.JrkDec()
	t.VelNeg = rcv.VelNeg()
	t.VelPos = rcv.VelPos()
	t.AccUnit = string(rcv.AccUnit())
	t.DecUnit = string(rcv.DecUnit())
	t.JrkAccUnit = string(rcv.JrkAccUnit())
	t.JrkDecUnit = string(rcv.JrkDecUnit())
	t.VelNegUnit = string(rcv.VelNegUnit())
	t.VelPosUnit = string(rcv.VelPosUnit())
}

func (rcv *AxsCfgDynSynchronisationLim) UnPack() *AxsCfgDynSynchronisationLimT {
	if rcv == nil { return nil }
	t := &AxsCfgDynSynchronisationLimT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCfgDynSynchronisationLim struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgDynSynchronisationLim(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgDynSynchronisationLim {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgDynSynchronisationLim{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgDynSynchronisationLim(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgDynSynchronisationLim {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgDynSynchronisationLim{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgDynSynchronisationLim) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgDynSynchronisationLim) Table() flatbuffers.Table {
	return rcv._tab
}

/// Acceleration limit
func (rcv *AxsCfgDynSynchronisationLim) Acc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Acceleration limit
func (rcv *AxsCfgDynSynchronisationLim) MutateAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// Deceleration limit
func (rcv *AxsCfgDynSynchronisationLim) Dec() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Deceleration limit
func (rcv *AxsCfgDynSynchronisationLim) MutateDec(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// Jerk limit, when accelerating
func (rcv *AxsCfgDynSynchronisationLim) JrkAcc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Jerk limit, when accelerating
func (rcv *AxsCfgDynSynchronisationLim) MutateJrkAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

/// Jerk limit, when decelerating
func (rcv *AxsCfgDynSynchronisationLim) JrkDec() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Jerk limit, when decelerating
func (rcv *AxsCfgDynSynchronisationLim) MutateJrkDec(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

/// Velocity limit, when moving the slave axis in negative direction
func (rcv *AxsCfgDynSynchronisationLim) VelNeg() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Velocity limit, when moving the slave axis in negative direction
func (rcv *AxsCfgDynSynchronisationLim) MutateVelNeg(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

/// Velocity limit, when moving the slave axis in positive direction
func (rcv *AxsCfgDynSynchronisationLim) VelPos() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Velocity limit, when moving the slave axis in positive direction
func (rcv *AxsCfgDynSynchronisationLim) MutateVelPos(n float64) bool {
	return rcv._tab.MutateFloat64Slot(14, n)
}

/// unit of acc
func (rcv *AxsCfgDynSynchronisationLim) AccUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of acc
/// unit of dec
func (rcv *AxsCfgDynSynchronisationLim) DecUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of dec
/// unit of jrkAcc
func (rcv *AxsCfgDynSynchronisationLim) JrkAccUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of jrkAcc
/// unit of jrkDec
func (rcv *AxsCfgDynSynchronisationLim) JrkDecUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of jrkDec
/// unit of velNeg
func (rcv *AxsCfgDynSynchronisationLim) VelNegUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of velNeg
/// unit of velPos
func (rcv *AxsCfgDynSynchronisationLim) VelPosUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of velPos
func AxsCfgDynSynchronisationLimStart(builder *flatbuffers.Builder) {
	builder.StartObject(12)
}
func AxsCfgDynSynchronisationLimAddAcc(builder *flatbuffers.Builder, acc float64) {
	builder.PrependFloat64Slot(0, acc, 0.0)
}
func AxsCfgDynSynchronisationLimAddDec(builder *flatbuffers.Builder, dec float64) {
	builder.PrependFloat64Slot(1, dec, 0.0)
}
func AxsCfgDynSynchronisationLimAddJrkAcc(builder *flatbuffers.Builder, jrkAcc float64) {
	builder.PrependFloat64Slot(2, jrkAcc, 0.0)
}
func AxsCfgDynSynchronisationLimAddJrkDec(builder *flatbuffers.Builder, jrkDec float64) {
	builder.PrependFloat64Slot(3, jrkDec, 0.0)
}
func AxsCfgDynSynchronisationLimAddVelNeg(builder *flatbuffers.Builder, velNeg float64) {
	builder.PrependFloat64Slot(4, velNeg, 0.0)
}
func AxsCfgDynSynchronisationLimAddVelPos(builder *flatbuffers.Builder, velPos float64) {
	builder.PrependFloat64Slot(5, velPos, 0.0)
}
func AxsCfgDynSynchronisationLimAddAccUnit(builder *flatbuffers.Builder, accUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(accUnit), 0)
}
func AxsCfgDynSynchronisationLimAddDecUnit(builder *flatbuffers.Builder, decUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(decUnit), 0)
}
func AxsCfgDynSynchronisationLimAddJrkAccUnit(builder *flatbuffers.Builder, jrkAccUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(jrkAccUnit), 0)
}
func AxsCfgDynSynchronisationLimAddJrkDecUnit(builder *flatbuffers.Builder, jrkDecUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(jrkDecUnit), 0)
}
func AxsCfgDynSynchronisationLimAddVelNegUnit(builder *flatbuffers.Builder, velNegUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(velNegUnit), 0)
}
func AxsCfgDynSynchronisationLimAddVelPosUnit(builder *flatbuffers.Builder, velPosUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(velPosUnit), 0)
}
func AxsCfgDynSynchronisationLimEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
