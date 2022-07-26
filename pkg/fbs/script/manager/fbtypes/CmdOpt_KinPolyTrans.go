// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CmdOpt_KinPolyTransT struct {
	Base *CmdOpt_BaseT
	Dist1 float64
	Dist2 float64
	Eps float64
}

func (t *CmdOpt_KinPolyTransT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	CmdOpt_KinPolyTransStart(builder)
	CmdOpt_KinPolyTransAddBase(builder, baseOffset)
	CmdOpt_KinPolyTransAddDist1(builder, t.Dist1)
	CmdOpt_KinPolyTransAddDist2(builder, t.Dist2)
	CmdOpt_KinPolyTransAddEps(builder, t.Eps)
	return CmdOpt_KinPolyTransEnd(builder)
}

func (rcv *CmdOpt_KinPolyTrans) UnPackTo(t *CmdOpt_KinPolyTransT) {
	t.Base = rcv.Base(nil).UnPack()
	t.Dist1 = rcv.Dist1()
	t.Dist2 = rcv.Dist2()
	t.Eps = rcv.Eps()
}

func (rcv *CmdOpt_KinPolyTrans) UnPack() *CmdOpt_KinPolyTransT {
	if rcv == nil { return nil }
	t := &CmdOpt_KinPolyTransT{}
	rcv.UnPackTo(t)
	return t
}

type CmdOpt_KinPolyTrans struct {
	_tab flatbuffers.Table
}

func GetRootAsCmdOpt_KinPolyTrans(buf []byte, offset flatbuffers.UOffsetT) *CmdOpt_KinPolyTrans {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CmdOpt_KinPolyTrans{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmdOpt_KinPolyTrans(buf []byte, offset flatbuffers.UOffsetT) *CmdOpt_KinPolyTrans {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CmdOpt_KinPolyTrans{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CmdOpt_KinPolyTrans) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CmdOpt_KinPolyTrans) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CmdOpt_KinPolyTrans) Base(obj *CmdOpt_Base) *CmdOpt_Base {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(CmdOpt_Base)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *CmdOpt_KinPolyTrans) Dist1() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *CmdOpt_KinPolyTrans) MutateDist1(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func (rcv *CmdOpt_KinPolyTrans) Dist2() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *CmdOpt_KinPolyTrans) MutateDist2(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func (rcv *CmdOpt_KinPolyTrans) Eps() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *CmdOpt_KinPolyTrans) MutateEps(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

func CmdOpt_KinPolyTransStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func CmdOpt_KinPolyTransAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func CmdOpt_KinPolyTransAddDist1(builder *flatbuffers.Builder, dist1 float64) {
	builder.PrependFloat64Slot(1, dist1, 0.0)
}
func CmdOpt_KinPolyTransAddDist2(builder *flatbuffers.Builder, dist2 float64) {
	builder.PrependFloat64Slot(2, dist2, 0.0)
}
func CmdOpt_KinPolyTransAddEps(builder *flatbuffers.Builder, eps float64) {
	builder.PrependFloat64Slot(3, eps, 0.0)
}
func CmdOpt_KinPolyTransEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
