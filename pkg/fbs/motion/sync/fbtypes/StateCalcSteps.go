// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// all supported calculation steps
type StateCalcStepsT struct {
	Steps []*StateCalcSingleStepT
}

func (t *StateCalcStepsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	stepsOffset := flatbuffers.UOffsetT(0)
	if t.Steps != nil {
		stepsLength := len(t.Steps)
		stepsOffsets := make([]flatbuffers.UOffsetT, stepsLength)
		for j := 0; j < stepsLength; j++ {
			stepsOffsets[j] = t.Steps[j].Pack(builder)
		}
		StateCalcStepsStartStepsVector(builder, stepsLength)
		for j := stepsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(stepsOffsets[j])
		}
		stepsOffset = builder.EndVector(stepsLength)
	}
	StateCalcStepsStart(builder)
	StateCalcStepsAddSteps(builder, stepsOffset)
	return StateCalcStepsEnd(builder)
}

func (rcv *StateCalcSteps) UnPackTo(t *StateCalcStepsT) {
	stepsLength := rcv.StepsLength()
	t.Steps = make([]*StateCalcSingleStepT, stepsLength)
	for j := 0; j < stepsLength; j++ {
		x := StateCalcSingleStep{}
		rcv.Steps(&x, j)
		t.Steps[j] = x.UnPack()
	}
}

func (rcv *StateCalcSteps) UnPack() *StateCalcStepsT {
	if rcv == nil { return nil }
	t := &StateCalcStepsT{}
	rcv.UnPackTo(t)
	return t
}

type StateCalcSteps struct {
	_tab flatbuffers.Table
}

func GetRootAsStateCalcSteps(buf []byte, offset flatbuffers.UOffsetT) *StateCalcSteps {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StateCalcSteps{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStateCalcSteps(buf []byte, offset flatbuffers.UOffsetT) *StateCalcSteps {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StateCalcSteps{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StateCalcSteps) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StateCalcSteps) Table() flatbuffers.Table {
	return rcv._tab
}

/// vector of all supported calculation steps
func (rcv *StateCalcSteps) Steps(obj *StateCalcSingleStep, j int) bool {
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

func (rcv *StateCalcSteps) StepsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of all supported calculation steps
func StateCalcStepsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func StateCalcStepsAddSteps(builder *flatbuffers.Builder, steps flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(steps), 0)
}
func StateCalcStepsStartStepsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func StateCalcStepsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}