// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd_AxsPos struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd_AxsPos(buf []byte, offset flatbuffers.UOffsetT) *Cmd_AxsPos {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd_AxsPos{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd_AxsPos(buf []byte, offset flatbuffers.UOffsetT) *Cmd_AxsPos {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd_AxsPos{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd_AxsPos) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd_AxsPos) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd_AxsPos) Base(obj *Cmd_Base) *Cmd_Base {
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

func (rcv *Cmd_AxsPos) Pos() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Cmd_AxsPos) MutatePos(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func (rcv *Cmd_AxsPos) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Cmd_AxsPos) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func (rcv *Cmd_AxsPos) Lim(obj *Cmd_DynLimits) *Cmd_DynLimits {
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

func Cmd_AxsPosStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func Cmd_AxsPosAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func Cmd_AxsPosAddPos(builder *flatbuffers.Builder, pos float64) {
	builder.PrependFloat64Slot(1, pos, 0.0)
}
func Cmd_AxsPosAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(2, buffered, false)
}
func Cmd_AxsPosAddLim(builder *flatbuffers.Builder, lim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(lim), 0)
}
func Cmd_AxsPosEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
