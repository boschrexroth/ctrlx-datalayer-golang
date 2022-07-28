// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type EnablingStatesT struct {
	EnablingStates []*EnablingStateT
	MachineIdentification string
}

func (t *EnablingStatesT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	enablingStatesOffset := flatbuffers.UOffsetT(0)
	if t.EnablingStates != nil {
		enablingStatesLength := len(t.EnablingStates)
		enablingStatesOffsets := make([]flatbuffers.UOffsetT, enablingStatesLength)
		for j := 0; j < enablingStatesLength; j++ {
			enablingStatesOffsets[j] = t.EnablingStates[j].Pack(builder)
		}
		EnablingStatesStartEnablingStatesVector(builder, enablingStatesLength)
		for j := enablingStatesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(enablingStatesOffsets[j])
		}
		enablingStatesOffset = builder.EndVector(enablingStatesLength)
	}
	machineIdentificationOffset := builder.CreateString(t.MachineIdentification)
	EnablingStatesStart(builder)
	EnablingStatesAddEnablingStates(builder, enablingStatesOffset)
	EnablingStatesAddMachineIdentification(builder, machineIdentificationOffset)
	return EnablingStatesEnd(builder)
}

func (rcv *EnablingStates) UnPackTo(t *EnablingStatesT) {
	enablingStatesLength := rcv.EnablingStatesLength()
	t.EnablingStates = make([]*EnablingStateT, enablingStatesLength)
	for j := 0; j < enablingStatesLength; j++ {
		x := EnablingState{}
		rcv.EnablingStates(&x, j)
		t.EnablingStates[j] = x.UnPack()
	}
	t.MachineIdentification = string(rcv.MachineIdentification())
}

func (rcv *EnablingStates) UnPack() *EnablingStatesT {
	if rcv == nil { return nil }
	t := &EnablingStatesT{}
	rcv.UnPackTo(t)
	return t
}

type EnablingStates struct {
	_tab flatbuffers.Table
}

func GetRootAsEnablingStates(buf []byte, offset flatbuffers.UOffsetT) *EnablingStates {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EnablingStates{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEnablingStates(buf []byte, offset flatbuffers.UOffsetT) *EnablingStates {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &EnablingStates{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *EnablingStates) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EnablingStates) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *EnablingStates) EnablingStates(obj *EnablingState, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *EnablingStates) EnablingStatesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *EnablingStates) MachineIdentification() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func EnablingStatesStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func EnablingStatesAddEnablingStates(builder *flatbuffers.Builder, enablingStates flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(enablingStates), 0)
}
func EnablingStatesStartEnablingStatesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EnablingStatesAddMachineIdentification(builder *flatbuffers.Builder, machineIdentification flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(machineIdentification), 0)
}
func EnablingStatesEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
