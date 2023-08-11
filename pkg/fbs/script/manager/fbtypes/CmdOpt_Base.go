// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CmdOpt_BaseT struct {
	Base *Cmd_BaseT `json:"base"`
	PermType string `json:"permType"`
}

func (t *CmdOpt_BaseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	permTypeOffset := flatbuffers.UOffsetT(0)
	if t.PermType != "" {
		permTypeOffset = builder.CreateString(t.PermType)
	}
	CmdOpt_BaseStart(builder)
	CmdOpt_BaseAddBase(builder, baseOffset)
	CmdOpt_BaseAddPermType(builder, permTypeOffset)
	return CmdOpt_BaseEnd(builder)
}

func (rcv *CmdOpt_Base) UnPackTo(t *CmdOpt_BaseT) {
	t.Base = rcv.Base(nil).UnPack()
	t.PermType = string(rcv.PermType())
}

func (rcv *CmdOpt_Base) UnPack() *CmdOpt_BaseT {
	if rcv == nil { return nil }
	t := &CmdOpt_BaseT{}
	rcv.UnPackTo(t)
	return t
}

type CmdOpt_Base struct {
	_tab flatbuffers.Table
}

func GetRootAsCmdOpt_Base(buf []byte, offset flatbuffers.UOffsetT) *CmdOpt_Base {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CmdOpt_Base{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmdOpt_Base(buf []byte, offset flatbuffers.UOffsetT) *CmdOpt_Base {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CmdOpt_Base{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CmdOpt_Base) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CmdOpt_Base) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CmdOpt_Base) Base(obj *Cmd_Base) *Cmd_Base {
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

func (rcv *CmdOpt_Base) PermType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func CmdOpt_BaseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func CmdOpt_BaseAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func CmdOpt_BaseAddPermType(builder *flatbuffers.Builder, permType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(permType), 0)
}
func CmdOpt_BaseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
