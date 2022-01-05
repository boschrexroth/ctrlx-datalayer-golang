// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type EnablingState struct {
	_tab flatbuffers.Table
}

func GetRootAsEnablingState(buf []byte, offset flatbuffers.UOffsetT) *EnablingState {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EnablingState{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEnablingState(buf []byte, offset flatbuffers.UOffsetT) *EnablingState {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &EnablingState{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *EnablingState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EnablingState) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *EnablingState) UnitName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *EnablingState) Messages() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *EnablingState) MutateMessages(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *EnablingState) Warnings() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *EnablingState) MutateWarnings(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func (rcv *EnablingState) Errors() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *EnablingState) MutateErrors(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func EnablingStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func EnablingStateAddUnitName(builder *flatbuffers.Builder, unitName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(unitName), 0)
}
func EnablingStateAddMessages(builder *flatbuffers.Builder, messages bool) {
	builder.PrependBoolSlot(1, messages, false)
}
func EnablingStateAddWarnings(builder *flatbuffers.Builder, warnings bool) {
	builder.PrependBoolSlot(2, warnings, false)
}
func EnablingStateAddErrors(builder *flatbuffers.Builder, errors bool) {
	builder.PrependBoolSlot(3, errors, false)
}
func EnablingStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
