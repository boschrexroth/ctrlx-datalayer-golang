// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd_AxsAbortT struct {
	Base *Cmd_BaseT `json:"base"`
	Dec float64 `json:"dec"`
	JrkDec float64 `json:"jrkDec"`
}

func (t *Cmd_AxsAbortT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	Cmd_AxsAbortStart(builder)
	Cmd_AxsAbortAddBase(builder, baseOffset)
	Cmd_AxsAbortAddDec(builder, t.Dec)
	Cmd_AxsAbortAddJrkDec(builder, t.JrkDec)
	return Cmd_AxsAbortEnd(builder)
}

func (rcv *Cmd_AxsAbort) UnPackTo(t *Cmd_AxsAbortT) {
	t.Base = rcv.Base(nil).UnPack()
	t.Dec = rcv.Dec()
	t.JrkDec = rcv.JrkDec()
}

func (rcv *Cmd_AxsAbort) UnPack() *Cmd_AxsAbortT {
	if rcv == nil { return nil }
	t := &Cmd_AxsAbortT{}
	rcv.UnPackTo(t)
	return t
}

type Cmd_AxsAbort struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd_AxsAbort(buf []byte, offset flatbuffers.UOffsetT) *Cmd_AxsAbort {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd_AxsAbort{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd_AxsAbort(buf []byte, offset flatbuffers.UOffsetT) *Cmd_AxsAbort {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd_AxsAbort{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd_AxsAbort) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd_AxsAbort) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd_AxsAbort) Base(obj *Cmd_Base) *Cmd_Base {
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

func (rcv *Cmd_AxsAbort) Dec() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Cmd_AxsAbort) MutateDec(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func (rcv *Cmd_AxsAbort) JrkDec() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Cmd_AxsAbort) MutateJrkDec(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func Cmd_AxsAbortStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func Cmd_AxsAbortAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func Cmd_AxsAbortAddDec(builder *flatbuffers.Builder, dec float64) {
	builder.PrependFloat64Slot(1, dec, 0.0)
}
func Cmd_AxsAbortAddJrkDec(builder *flatbuffers.Builder, jrkDec float64) {
	builder.PrependFloat64Slot(2, jrkDec, 0.0)
}
func Cmd_AxsAbortEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
