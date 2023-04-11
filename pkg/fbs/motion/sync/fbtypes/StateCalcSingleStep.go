// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// a single supported calculation step
type StateCalcSingleStepT struct {
	Name string
	Description string
	DocuRef string
	Inputs []DataID
	Outputs []DataID
	Parameter *StateCalcStepParamsT
	MutexParameter *StateCalcStepMutexParamT
}

func (t *StateCalcSingleStepT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	descriptionOffset := builder.CreateString(t.Description)
	docuRefOffset := builder.CreateString(t.DocuRef)
	inputsOffset := flatbuffers.UOffsetT(0)
	if t.Inputs != nil {
		inputsLength := len(t.Inputs)
		StateCalcSingleStepStartInputsVector(builder, inputsLength)
		for j := inputsLength - 1; j >= 0; j-- {
			builder.PrependInt8(int8(t.Inputs[j]))
		}
		inputsOffset = builder.EndVector(inputsLength)
	}
	outputsOffset := flatbuffers.UOffsetT(0)
	if t.Outputs != nil {
		outputsLength := len(t.Outputs)
		StateCalcSingleStepStartOutputsVector(builder, outputsLength)
		for j := outputsLength - 1; j >= 0; j-- {
			builder.PrependInt8(int8(t.Outputs[j]))
		}
		outputsOffset = builder.EndVector(outputsLength)
	}
	parameterOffset := t.Parameter.Pack(builder)
	mutexParameterOffset := t.MutexParameter.Pack(builder)
	StateCalcSingleStepStart(builder)
	StateCalcSingleStepAddName(builder, nameOffset)
	StateCalcSingleStepAddDescription(builder, descriptionOffset)
	StateCalcSingleStepAddDocuRef(builder, docuRefOffset)
	StateCalcSingleStepAddInputs(builder, inputsOffset)
	StateCalcSingleStepAddOutputs(builder, outputsOffset)
	StateCalcSingleStepAddParameter(builder, parameterOffset)
	StateCalcSingleStepAddMutexParameter(builder, mutexParameterOffset)
	return StateCalcSingleStepEnd(builder)
}

func (rcv *StateCalcSingleStep) UnPackTo(t *StateCalcSingleStepT) {
	t.Name = string(rcv.Name())
	t.Description = string(rcv.Description())
	t.DocuRef = string(rcv.DocuRef())
	inputsLength := rcv.InputsLength()
	t.Inputs = make([]DataID, inputsLength)
	for j := 0; j < inputsLength; j++ {
		t.Inputs[j] = rcv.Inputs(j)
	}
	outputsLength := rcv.OutputsLength()
	t.Outputs = make([]DataID, outputsLength)
	for j := 0; j < outputsLength; j++ {
		t.Outputs[j] = rcv.Outputs(j)
	}
	t.Parameter = rcv.Parameter(nil).UnPack()
	t.MutexParameter = rcv.MutexParameter(nil).UnPack()
}

func (rcv *StateCalcSingleStep) UnPack() *StateCalcSingleStepT {
	if rcv == nil { return nil }
	t := &StateCalcSingleStepT{}
	rcv.UnPackTo(t)
	return t
}

type StateCalcSingleStep struct {
	_tab flatbuffers.Table
}

func GetRootAsStateCalcSingleStep(buf []byte, offset flatbuffers.UOffsetT) *StateCalcSingleStep {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StateCalcSingleStep{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStateCalcSingleStep(buf []byte, offset flatbuffers.UOffsetT) *StateCalcSingleStep {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StateCalcSingleStep{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StateCalcSingleStep) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StateCalcSingleStep) Table() flatbuffers.Table {
	return rcv._tab
}

/// Name of the calculation step (indicated by its type ID)
func (rcv *StateCalcSingleStep) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Name of the calculation step (indicated by its type ID)
/// what the calculation step does
func (rcv *StateCalcSingleStep) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// what the calculation step does
/// reference to user manual
func (rcv *StateCalcSingleStep) DocuRef() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// reference to user manual
/// vector of required input types
func (rcv *StateCalcSingleStep) Inputs(j int) DataID {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return DataID(rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1)))
	}
	return 0
}

func (rcv *StateCalcSingleStep) InputsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of required input types
func (rcv *StateCalcSingleStep) MutateInputs(j int, n DataID) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), int8(n))
	}
	return false
}

/// vector of required output types
func (rcv *StateCalcSingleStep) Outputs(j int) DataID {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return DataID(rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1)))
	}
	return 0
}

func (rcv *StateCalcSingleStep) OutputsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of required output types
func (rcv *StateCalcSingleStep) MutateOutputs(j int, n DataID) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), int8(n))
	}
	return false
}

/// parameters of the calculation step
func (rcv *StateCalcSingleStep) Parameter(obj *StateCalcStepParams) *StateCalcStepParams {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(StateCalcStepParams)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// parameters of the calculation step
/// mutex groups of parameters (identified by the names), that are optional and exclude each other
func (rcv *StateCalcSingleStep) MutexParameter(obj *StateCalcStepMutexParam) *StateCalcStepMutexParam {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(StateCalcStepMutexParam)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// mutex groups of parameters (identified by the names), that are optional and exclude each other
func StateCalcSingleStepStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func StateCalcSingleStepAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func StateCalcSingleStepAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(description), 0)
}
func StateCalcSingleStepAddDocuRef(builder *flatbuffers.Builder, docuRef flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(docuRef), 0)
}
func StateCalcSingleStepAddInputs(builder *flatbuffers.Builder, inputs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(inputs), 0)
}
func StateCalcSingleStepStartInputsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func StateCalcSingleStepAddOutputs(builder *flatbuffers.Builder, outputs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(outputs), 0)
}
func StateCalcSingleStepStartOutputsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func StateCalcSingleStepAddParameter(builder *flatbuffers.Builder, parameter flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(parameter), 0)
}
func StateCalcSingleStepAddMutexParameter(builder *flatbuffers.Builder, mutexParameter flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(mutexParameter), 0)
}
func StateCalcSingleStepEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
