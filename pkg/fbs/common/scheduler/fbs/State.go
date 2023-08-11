// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StateT struct {
	State CurrentState `json:"state"`
}

func (t *StateT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	StateStart(builder)
	StateAddState(builder, t.State)
	return StateEnd(builder)
}

func (rcv *State) UnPackTo(t *StateT) {
	t.State = rcv.State()
}

func (rcv *State) UnPack() *StateT {
	if rcv == nil { return nil }
	t := &StateT{}
	rcv.UnPackTo(t)
	return t
}

type State struct {
	_tab flatbuffers.Table
}

func GetRootAsState(buf []byte, offset flatbuffers.UOffsetT) *State {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &State{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsState(buf []byte, offset flatbuffers.UOffsetT) *State {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &State{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *State) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *State) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *State) State() CurrentState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return CurrentState(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 1
}

func (rcv *State) MutateState(n CurrentState) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func StateStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func StateAddState(builder *flatbuffers.Builder, state CurrentState) {
	builder.PrependInt8Slot(0, int8(state), 1)
}
func StateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
