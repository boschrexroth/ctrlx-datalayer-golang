// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd_DynLimitsT struct {
	Vel float64 `json:"vel"`
	Acc float64 `json:"acc"`
	Dec float64 `json:"dec"`
	JrkAcc float64 `json:"jrkAcc"`
	JrkDec float64 `json:"jrkDec"`
}

func (t *Cmd_DynLimitsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Cmd_DynLimitsStart(builder)
	Cmd_DynLimitsAddVel(builder, t.Vel)
	Cmd_DynLimitsAddAcc(builder, t.Acc)
	Cmd_DynLimitsAddDec(builder, t.Dec)
	Cmd_DynLimitsAddJrkAcc(builder, t.JrkAcc)
	Cmd_DynLimitsAddJrkDec(builder, t.JrkDec)
	return Cmd_DynLimitsEnd(builder)
}

func (rcv *Cmd_DynLimits) UnPackTo(t *Cmd_DynLimitsT) {
	t.Vel = rcv.Vel()
	t.Acc = rcv.Acc()
	t.Dec = rcv.Dec()
	t.JrkAcc = rcv.JrkAcc()
	t.JrkDec = rcv.JrkDec()
}

func (rcv *Cmd_DynLimits) UnPack() *Cmd_DynLimitsT {
	if rcv == nil { return nil }
	t := &Cmd_DynLimitsT{}
	rcv.UnPackTo(t)
	return t
}

type Cmd_DynLimits struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd_DynLimits(buf []byte, offset flatbuffers.UOffsetT) *Cmd_DynLimits {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd_DynLimits{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd_DynLimits(buf []byte, offset flatbuffers.UOffsetT) *Cmd_DynLimits {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd_DynLimits{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd_DynLimits) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd_DynLimits) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd_DynLimits) Vel() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Cmd_DynLimits) MutateVel(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

func (rcv *Cmd_DynLimits) Acc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Cmd_DynLimits) MutateAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func (rcv *Cmd_DynLimits) Dec() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Cmd_DynLimits) MutateDec(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func (rcv *Cmd_DynLimits) JrkAcc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Cmd_DynLimits) MutateJrkAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

func (rcv *Cmd_DynLimits) JrkDec() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Cmd_DynLimits) MutateJrkDec(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

func Cmd_DynLimitsStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func Cmd_DynLimitsAddVel(builder *flatbuffers.Builder, vel float64) {
	builder.PrependFloat64Slot(0, vel, 0.0)
}
func Cmd_DynLimitsAddAcc(builder *flatbuffers.Builder, acc float64) {
	builder.PrependFloat64Slot(1, acc, 0.0)
}
func Cmd_DynLimitsAddDec(builder *flatbuffers.Builder, dec float64) {
	builder.PrependFloat64Slot(2, dec, 0.0)
}
func Cmd_DynLimitsAddJrkAcc(builder *flatbuffers.Builder, jrkAcc float64) {
	builder.PrependFloat64Slot(3, jrkAcc, 0.0)
}
func Cmd_DynLimitsAddJrkDec(builder *flatbuffers.Builder, jrkDec float64) {
	builder.PrependFloat64Slot(4, jrkDec, 0.0)
}
func Cmd_DynLimitsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
