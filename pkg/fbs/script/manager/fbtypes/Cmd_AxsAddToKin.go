// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd_AxsAddToKinT struct {
	Base *Cmd_BaseT `json:"base"`
	KinName string `json:"kinName"`
	Buffered bool `json:"buffered"`
}

func (t *Cmd_AxsAddToKinT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	kinNameOffset := flatbuffers.UOffsetT(0)
	if t.KinName != "" {
		kinNameOffset = builder.CreateString(t.KinName)
	}
	Cmd_AxsAddToKinStart(builder)
	Cmd_AxsAddToKinAddBase(builder, baseOffset)
	Cmd_AxsAddToKinAddKinName(builder, kinNameOffset)
	Cmd_AxsAddToKinAddBuffered(builder, t.Buffered)
	return Cmd_AxsAddToKinEnd(builder)
}

func (rcv *Cmd_AxsAddToKin) UnPackTo(t *Cmd_AxsAddToKinT) {
	t.Base = rcv.Base(nil).UnPack()
	t.KinName = string(rcv.KinName())
	t.Buffered = rcv.Buffered()
}

func (rcv *Cmd_AxsAddToKin) UnPack() *Cmd_AxsAddToKinT {
	if rcv == nil { return nil }
	t := &Cmd_AxsAddToKinT{}
	rcv.UnPackTo(t)
	return t
}

type Cmd_AxsAddToKin struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd_AxsAddToKin(buf []byte, offset flatbuffers.UOffsetT) *Cmd_AxsAddToKin {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd_AxsAddToKin{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd_AxsAddToKin(buf []byte, offset flatbuffers.UOffsetT) *Cmd_AxsAddToKin {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd_AxsAddToKin{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd_AxsAddToKin) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd_AxsAddToKin) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd_AxsAddToKin) Base(obj *Cmd_Base) *Cmd_Base {
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

func (rcv *Cmd_AxsAddToKin) KinName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Cmd_AxsAddToKin) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Cmd_AxsAddToKin) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func Cmd_AxsAddToKinStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func Cmd_AxsAddToKinAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func Cmd_AxsAddToKinAddKinName(builder *flatbuffers.Builder, kinName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(kinName), 0)
}
func Cmd_AxsAddToKinAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(2, buffered, false)
}
func Cmd_AxsAddToKinEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
