// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CmdOpt_KinSetPCS struct {
	_tab flatbuffers.Table
}

func GetRootAsCmdOpt_KinSetPCS(buf []byte, offset flatbuffers.UOffsetT) *CmdOpt_KinSetPCS {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CmdOpt_KinSetPCS{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmdOpt_KinSetPCS(buf []byte, offset flatbuffers.UOffsetT) *CmdOpt_KinSetPCS {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CmdOpt_KinSetPCS{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CmdOpt_KinSetPCS) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CmdOpt_KinSetPCS) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CmdOpt_KinSetPCS) Base(obj *CmdOpt_Base) *CmdOpt_Base {
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

func (rcv *CmdOpt_KinSetPCS) SetName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func CmdOpt_KinSetPCSStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func CmdOpt_KinSetPCSAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func CmdOpt_KinSetPCSAddSetName(builder *flatbuffers.Builder, setName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(setName), 0)
}
func CmdOpt_KinSetPCSEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
