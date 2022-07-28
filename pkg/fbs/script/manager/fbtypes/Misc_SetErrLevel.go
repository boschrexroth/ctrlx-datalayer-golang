// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Misc_SetErrLevelT struct {
	Base *Cmd_BaseT
	ErrLvl string
}

func (t *Misc_SetErrLevelT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	errLvlOffset := builder.CreateString(t.ErrLvl)
	Misc_SetErrLevelStart(builder)
	Misc_SetErrLevelAddBase(builder, baseOffset)
	Misc_SetErrLevelAddErrLvl(builder, errLvlOffset)
	return Misc_SetErrLevelEnd(builder)
}

func (rcv *Misc_SetErrLevel) UnPackTo(t *Misc_SetErrLevelT) {
	t.Base = rcv.Base(nil).UnPack()
	t.ErrLvl = string(rcv.ErrLvl())
}

func (rcv *Misc_SetErrLevel) UnPack() *Misc_SetErrLevelT {
	if rcv == nil { return nil }
	t := &Misc_SetErrLevelT{}
	rcv.UnPackTo(t)
	return t
}

type Misc_SetErrLevel struct {
	_tab flatbuffers.Table
}

func GetRootAsMisc_SetErrLevel(buf []byte, offset flatbuffers.UOffsetT) *Misc_SetErrLevel {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Misc_SetErrLevel{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMisc_SetErrLevel(buf []byte, offset flatbuffers.UOffsetT) *Misc_SetErrLevel {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Misc_SetErrLevel{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Misc_SetErrLevel) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Misc_SetErrLevel) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Misc_SetErrLevel) Base(obj *Cmd_Base) *Cmd_Base {
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

func (rcv *Misc_SetErrLevel) ErrLvl() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func Misc_SetErrLevelStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Misc_SetErrLevelAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func Misc_SetErrLevelAddErrLvl(builder *flatbuffers.Builder, errLvl flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(errLvl), 0)
}
func Misc_SetErrLevelEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
