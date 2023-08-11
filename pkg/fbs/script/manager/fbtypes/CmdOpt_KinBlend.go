// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CmdOpt_KinBlendT struct {
	Base *CmdOpt_BaseT `json:"base"`
	Dist1 float64 `json:"dist1"`
	Dist2 float64 `json:"dist2"`
}

func (t *CmdOpt_KinBlendT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	CmdOpt_KinBlendStart(builder)
	CmdOpt_KinBlendAddBase(builder, baseOffset)
	CmdOpt_KinBlendAddDist1(builder, t.Dist1)
	CmdOpt_KinBlendAddDist2(builder, t.Dist2)
	return CmdOpt_KinBlendEnd(builder)
}

func (rcv *CmdOpt_KinBlend) UnPackTo(t *CmdOpt_KinBlendT) {
	t.Base = rcv.Base(nil).UnPack()
	t.Dist1 = rcv.Dist1()
	t.Dist2 = rcv.Dist2()
}

func (rcv *CmdOpt_KinBlend) UnPack() *CmdOpt_KinBlendT {
	if rcv == nil { return nil }
	t := &CmdOpt_KinBlendT{}
	rcv.UnPackTo(t)
	return t
}

type CmdOpt_KinBlend struct {
	_tab flatbuffers.Table
}

func GetRootAsCmdOpt_KinBlend(buf []byte, offset flatbuffers.UOffsetT) *CmdOpt_KinBlend {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CmdOpt_KinBlend{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmdOpt_KinBlend(buf []byte, offset flatbuffers.UOffsetT) *CmdOpt_KinBlend {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CmdOpt_KinBlend{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CmdOpt_KinBlend) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CmdOpt_KinBlend) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CmdOpt_KinBlend) Base(obj *CmdOpt_Base) *CmdOpt_Base {
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

func (rcv *CmdOpt_KinBlend) Dist1() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *CmdOpt_KinBlend) MutateDist1(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func (rcv *CmdOpt_KinBlend) Dist2() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *CmdOpt_KinBlend) MutateDist2(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func CmdOpt_KinBlendStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func CmdOpt_KinBlendAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func CmdOpt_KinBlendAddDist1(builder *flatbuffers.Builder, dist1 float64) {
	builder.PrependFloat64Slot(1, dist1, 0.0)
}
func CmdOpt_KinBlendAddDist2(builder *flatbuffers.Builder, dist2 float64) {
	builder.PrependFloat64Slot(2, dist2, 0.0)
}
func CmdOpt_KinBlendEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
