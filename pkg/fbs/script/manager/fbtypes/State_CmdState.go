// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type State_CmdStateT struct {
	ObjName string
	CmdID uint64
}

func (t *State_CmdStateT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	objNameOffset := builder.CreateString(t.ObjName)
	State_CmdStateStart(builder)
	State_CmdStateAddObjName(builder, objNameOffset)
	State_CmdStateAddCmdID(builder, t.CmdID)
	return State_CmdStateEnd(builder)
}

func (rcv *State_CmdState) UnPackTo(t *State_CmdStateT) {
	t.ObjName = string(rcv.ObjName())
	t.CmdID = rcv.CmdID()
}

func (rcv *State_CmdState) UnPack() *State_CmdStateT {
	if rcv == nil { return nil }
	t := &State_CmdStateT{}
	rcv.UnPackTo(t)
	return t
}

type State_CmdState struct {
	_tab flatbuffers.Table
}

func GetRootAsState_CmdState(buf []byte, offset flatbuffers.UOffsetT) *State_CmdState {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &State_CmdState{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsState_CmdState(buf []byte, offset flatbuffers.UOffsetT) *State_CmdState {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &State_CmdState{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *State_CmdState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *State_CmdState) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *State_CmdState) ObjName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *State_CmdState) CmdID() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *State_CmdState) MutateCmdID(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func State_CmdStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func State_CmdStateAddObjName(builder *flatbuffers.Builder, objName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(objName), 0)
}
func State_CmdStateAddCmdID(builder *flatbuffers.Builder, cmdID uint64) {
	builder.PrependUint64Slot(1, cmdID, 0)
}
func State_CmdStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
