// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type State_KinGetPos struct {
	_tab flatbuffers.Table
}

func GetRootAsState_KinGetPos(buf []byte, offset flatbuffers.UOffsetT) *State_KinGetPos {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &State_KinGetPos{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsState_KinGetPos(buf []byte, offset flatbuffers.UOffsetT) *State_KinGetPos {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &State_KinGetPos{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *State_KinGetPos) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *State_KinGetPos) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *State_KinGetPos) ObjName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *State_KinGetPos) CoordSysOut() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func State_KinGetPosStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func State_KinGetPosAddObjName(builder *flatbuffers.Builder, objName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(objName), 0)
}
func State_KinGetPosAddCoordSysOut(builder *flatbuffers.Builder, coordSysOut flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(coordSysOut), 0)
}
func State_KinGetPosEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
