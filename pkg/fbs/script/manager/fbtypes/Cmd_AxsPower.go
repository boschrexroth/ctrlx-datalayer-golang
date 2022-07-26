// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd_AxsPowerT struct {
	Base *Cmd_BaseT
	SwitchOn bool
}

func (t *Cmd_AxsPowerT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	Cmd_AxsPowerStart(builder)
	Cmd_AxsPowerAddBase(builder, baseOffset)
	Cmd_AxsPowerAddSwitchOn(builder, t.SwitchOn)
	return Cmd_AxsPowerEnd(builder)
}

func (rcv *Cmd_AxsPower) UnPackTo(t *Cmd_AxsPowerT) {
	t.Base = rcv.Base(nil).UnPack()
	t.SwitchOn = rcv.SwitchOn()
}

func (rcv *Cmd_AxsPower) UnPack() *Cmd_AxsPowerT {
	if rcv == nil { return nil }
	t := &Cmd_AxsPowerT{}
	rcv.UnPackTo(t)
	return t
}

type Cmd_AxsPower struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd_AxsPower(buf []byte, offset flatbuffers.UOffsetT) *Cmd_AxsPower {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd_AxsPower{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd_AxsPower(buf []byte, offset flatbuffers.UOffsetT) *Cmd_AxsPower {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd_AxsPower{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd_AxsPower) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd_AxsPower) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd_AxsPower) Base(obj *Cmd_Base) *Cmd_Base {
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

func (rcv *Cmd_AxsPower) SwitchOn() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Cmd_AxsPower) MutateSwitchOn(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func Cmd_AxsPowerStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Cmd_AxsPowerAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func Cmd_AxsPowerAddSwitchOn(builder *flatbuffers.Builder, switchOn bool) {
	builder.PrependBoolSlot(1, switchOn, false)
}
func Cmd_AxsPowerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
