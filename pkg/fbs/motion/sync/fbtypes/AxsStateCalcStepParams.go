// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// all supported calculation steps
type AxsStateCalcStepParamsT struct {
	Params []*AxsStateCalcStepSingleParamT
}

func (t *AxsStateCalcStepParamsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	paramsOffset := flatbuffers.UOffsetT(0)
	if t.Params != nil {
		paramsLength := len(t.Params)
		paramsOffsets := make([]flatbuffers.UOffsetT, paramsLength)
		for j := 0; j < paramsLength; j++ {
			paramsOffsets[j] = t.Params[j].Pack(builder)
		}
		AxsStateCalcStepParamsStartParamsVector(builder, paramsLength)
		for j := paramsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(paramsOffsets[j])
		}
		paramsOffset = builder.EndVector(paramsLength)
	}
	AxsStateCalcStepParamsStart(builder)
	AxsStateCalcStepParamsAddParams(builder, paramsOffset)
	return AxsStateCalcStepParamsEnd(builder)
}

func (rcv *AxsStateCalcStepParams) UnPackTo(t *AxsStateCalcStepParamsT) {
	paramsLength := rcv.ParamsLength()
	t.Params = make([]*AxsStateCalcStepSingleParamT, paramsLength)
	for j := 0; j < paramsLength; j++ {
		x := AxsStateCalcStepSingleParam{}
		rcv.Params(&x, j)
		t.Params[j] = x.UnPack()
	}
}

func (rcv *AxsStateCalcStepParams) UnPack() *AxsStateCalcStepParamsT {
	if rcv == nil { return nil }
	t := &AxsStateCalcStepParamsT{}
	rcv.UnPackTo(t)
	return t
}

type AxsStateCalcStepParams struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsStateCalcStepParams(buf []byte, offset flatbuffers.UOffsetT) *AxsStateCalcStepParams {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsStateCalcStepParams{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsStateCalcStepParams(buf []byte, offset flatbuffers.UOffsetT) *AxsStateCalcStepParams {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsStateCalcStepParams{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsStateCalcStepParams) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsStateCalcStepParams) Table() flatbuffers.Table {
	return rcv._tab
}

/// vector of all params
func (rcv *AxsStateCalcStepParams) Params(obj *AxsStateCalcStepSingleParam, j int) bool {
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

func (rcv *AxsStateCalcStepParams) ParamsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of all params
func AxsStateCalcStepParamsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func AxsStateCalcStepParamsAddParams(builder *flatbuffers.Builder, params flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(params), 0)
}
func AxsStateCalcStepParamsStartParamsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func AxsStateCalcStepParamsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}