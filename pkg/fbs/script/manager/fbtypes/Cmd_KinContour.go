// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd_KinContourT struct {
	Base *Cmd_BaseT `json:"base"`
	IsStart bool `json:"isStart"`
	PrepCmds uint32 `json:"prepCmds"`
}

func (t *Cmd_KinContourT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	Cmd_KinContourStart(builder)
	Cmd_KinContourAddBase(builder, baseOffset)
	Cmd_KinContourAddIsStart(builder, t.IsStart)
	Cmd_KinContourAddPrepCmds(builder, t.PrepCmds)
	return Cmd_KinContourEnd(builder)
}

func (rcv *Cmd_KinContour) UnPackTo(t *Cmd_KinContourT) {
	t.Base = rcv.Base(nil).UnPack()
	t.IsStart = rcv.IsStart()
	t.PrepCmds = rcv.PrepCmds()
}

func (rcv *Cmd_KinContour) UnPack() *Cmd_KinContourT {
	if rcv == nil { return nil }
	t := &Cmd_KinContourT{}
	rcv.UnPackTo(t)
	return t
}

type Cmd_KinContour struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd_KinContour(buf []byte, offset flatbuffers.UOffsetT) *Cmd_KinContour {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd_KinContour{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd_KinContour(buf []byte, offset flatbuffers.UOffsetT) *Cmd_KinContour {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd_KinContour{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd_KinContour) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd_KinContour) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd_KinContour) Base(obj *Cmd_Base) *Cmd_Base {
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

func (rcv *Cmd_KinContour) IsStart() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *Cmd_KinContour) MutateIsStart(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *Cmd_KinContour) PrepCmds() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Cmd_KinContour) MutatePrepCmds(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func Cmd_KinContourStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func Cmd_KinContourAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func Cmd_KinContourAddIsStart(builder *flatbuffers.Builder, isStart bool) {
	builder.PrependBoolSlot(1, isStart, true)
}
func Cmd_KinContourAddPrepCmds(builder *flatbuffers.Builder, prepCmds uint32) {
	builder.PrependUint32Slot(2, prepCmds, 0)
}
func Cmd_KinContourEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
