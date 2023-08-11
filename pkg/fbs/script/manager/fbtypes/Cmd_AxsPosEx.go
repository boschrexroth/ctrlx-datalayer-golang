// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd_AxsPosExT struct {
	Base *Cmd_BaseT `json:"base"`
	Pos float64 `json:"pos"`
	Buffered bool `json:"buffered"`
	Lim *Cmd_DynLimitsT `json:"lim"`
	Direction CmdPosAbsDir `json:"direction"`
}

func (t *Cmd_AxsPosExT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	limOffset := t.Lim.Pack(builder)
	Cmd_AxsPosExStart(builder)
	Cmd_AxsPosExAddBase(builder, baseOffset)
	Cmd_AxsPosExAddPos(builder, t.Pos)
	Cmd_AxsPosExAddBuffered(builder, t.Buffered)
	Cmd_AxsPosExAddLim(builder, limOffset)
	Cmd_AxsPosExAddDirection(builder, t.Direction)
	return Cmd_AxsPosExEnd(builder)
}

func (rcv *Cmd_AxsPosEx) UnPackTo(t *Cmd_AxsPosExT) {
	t.Base = rcv.Base(nil).UnPack()
	t.Pos = rcv.Pos()
	t.Buffered = rcv.Buffered()
	t.Lim = rcv.Lim(nil).UnPack()
	t.Direction = rcv.Direction()
}

func (rcv *Cmd_AxsPosEx) UnPack() *Cmd_AxsPosExT {
	if rcv == nil { return nil }
	t := &Cmd_AxsPosExT{}
	rcv.UnPackTo(t)
	return t
}

type Cmd_AxsPosEx struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd_AxsPosEx(buf []byte, offset flatbuffers.UOffsetT) *Cmd_AxsPosEx {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd_AxsPosEx{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd_AxsPosEx(buf []byte, offset flatbuffers.UOffsetT) *Cmd_AxsPosEx {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd_AxsPosEx{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd_AxsPosEx) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd_AxsPosEx) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd_AxsPosEx) Base(obj *Cmd_Base) *Cmd_Base {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Cmd_Base)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Cmd_AxsPosEx) Pos() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Cmd_AxsPosEx) MutatePos(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func (rcv *Cmd_AxsPosEx) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Cmd_AxsPosEx) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func (rcv *Cmd_AxsPosEx) Lim(obj *Cmd_DynLimits) *Cmd_DynLimits {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
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

/// Selected direction for PosAbs with modulo axis (ignored in all other cases)
func (rcv *Cmd_AxsPosEx) Direction() CmdPosAbsDir {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return CmdPosAbsDir(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Selected direction for PosAbs with modulo axis (ignored in all other cases)
func (rcv *Cmd_AxsPosEx) MutateDirection(n CmdPosAbsDir) bool {
	return rcv._tab.MutateInt8Slot(12, int8(n))
}

func Cmd_AxsPosExStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func Cmd_AxsPosExAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func Cmd_AxsPosExAddPos(builder *flatbuffers.Builder, pos float64) {
	builder.PrependFloat64Slot(1, pos, 0.0)
}
func Cmd_AxsPosExAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(2, buffered, false)
}
func Cmd_AxsPosExAddLim(builder *flatbuffers.Builder, lim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(lim), 0)
}
func Cmd_AxsPosExAddDirection(builder *flatbuffers.Builder, direction CmdPosAbsDir) {
	builder.PrependInt8Slot(4, int8(direction), 0)
}
func Cmd_AxsPosExEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
