// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package server

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ServerStateTypeT struct {
	State EnumServerState `json:"state"`
}

func (t *ServerStateTypeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ServerStateTypeStart(builder)
	ServerStateTypeAddState(builder, t.State)
	return ServerStateTypeEnd(builder)
}

func (rcv *ServerStateType) UnPackTo(t *ServerStateTypeT) {
	t.State = rcv.State()
}

func (rcv *ServerStateType) UnPack() *ServerStateTypeT {
	if rcv == nil { return nil }
	t := &ServerStateTypeT{}
	rcv.UnPackTo(t)
	return t
}

type ServerStateType struct {
	_tab flatbuffers.Table
}

func GetRootAsServerStateType(buf []byte, offset flatbuffers.UOffsetT) *ServerStateType {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ServerStateType{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsServerStateType(buf []byte, offset flatbuffers.UOffsetT) *ServerStateType {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ServerStateType{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ServerStateType) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ServerStateType) Table() flatbuffers.Table {
	return rcv._tab
}

/// Shows the current state of the OPC UA Server
func (rcv *ServerStateType) State() EnumServerState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return EnumServerState(rcv._tab.GetInt32(o + rcv._tab.Pos))
	}
	return 0
}

/// Shows the current state of the OPC UA Server
func (rcv *ServerStateType) MutateState(n EnumServerState) bool {
	return rcv._tab.MutateInt32Slot(4, int32(n))
}

func ServerStateTypeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ServerStateTypeAddState(builder *flatbuffers.Builder, state EnumServerState) {
	builder.PrependInt32Slot(0, int32(state), 0)
}
func ServerStateTypeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
