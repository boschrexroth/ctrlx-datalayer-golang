// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CmdOpt_KinAxsDynLimT struct {
	Base *CmdOpt_BaseT
	AxsName string
	Lim *Cmd_DynLimitsT
}

func (t *CmdOpt_KinAxsDynLimT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	axsNameOffset := builder.CreateString(t.AxsName)
	limOffset := t.Lim.Pack(builder)
	CmdOpt_KinAxsDynLimStart(builder)
	CmdOpt_KinAxsDynLimAddBase(builder, baseOffset)
	CmdOpt_KinAxsDynLimAddAxsName(builder, axsNameOffset)
	CmdOpt_KinAxsDynLimAddLim(builder, limOffset)
	return CmdOpt_KinAxsDynLimEnd(builder)
}

func (rcv *CmdOpt_KinAxsDynLim) UnPackTo(t *CmdOpt_KinAxsDynLimT) {
	t.Base = rcv.Base(nil).UnPack()
	t.AxsName = string(rcv.AxsName())
	t.Lim = rcv.Lim(nil).UnPack()
}

func (rcv *CmdOpt_KinAxsDynLim) UnPack() *CmdOpt_KinAxsDynLimT {
	if rcv == nil { return nil }
	t := &CmdOpt_KinAxsDynLimT{}
	rcv.UnPackTo(t)
	return t
}

type CmdOpt_KinAxsDynLim struct {
	_tab flatbuffers.Table
}

func GetRootAsCmdOpt_KinAxsDynLim(buf []byte, offset flatbuffers.UOffsetT) *CmdOpt_KinAxsDynLim {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CmdOpt_KinAxsDynLim{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmdOpt_KinAxsDynLim(buf []byte, offset flatbuffers.UOffsetT) *CmdOpt_KinAxsDynLim {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CmdOpt_KinAxsDynLim{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CmdOpt_KinAxsDynLim) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CmdOpt_KinAxsDynLim) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CmdOpt_KinAxsDynLim) Base(obj *CmdOpt_Base) *CmdOpt_Base {
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

func (rcv *CmdOpt_KinAxsDynLim) AxsName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CmdOpt_KinAxsDynLim) Lim(obj *Cmd_DynLimits) *Cmd_DynLimits {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Cmd_DynLimits)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func CmdOpt_KinAxsDynLimStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func CmdOpt_KinAxsDynLimAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func CmdOpt_KinAxsDynLimAddAxsName(builder *flatbuffers.Builder, axsName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(axsName), 0)
}
func CmdOpt_KinAxsDynLimAddLim(builder *flatbuffers.Builder, lim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(lim), 0)
}
func CmdOpt_KinAxsDynLimEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
