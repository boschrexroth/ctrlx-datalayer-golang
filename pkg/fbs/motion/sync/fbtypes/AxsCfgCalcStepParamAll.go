// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configuration of all parameters of a single calculation step
type AxsCfgCalcStepParamAllT struct {
	Params []*AxsCfgCalcStepParamT `json:"params"`
}

func (t *AxsCfgCalcStepParamAllT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	paramsOffset := flatbuffers.UOffsetT(0)
	if t.Params != nil {
		paramsLength := len(t.Params)
		paramsOffsets := make([]flatbuffers.UOffsetT, paramsLength)
		for j := 0; j < paramsLength; j++ {
			paramsOffsets[j] = t.Params[j].Pack(builder)
		}
		AxsCfgCalcStepParamAllStartParamsVector(builder, paramsLength)
		for j := paramsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(paramsOffsets[j])
		}
		paramsOffset = builder.EndVector(paramsLength)
	}
	AxsCfgCalcStepParamAllStart(builder)
	AxsCfgCalcStepParamAllAddParams(builder, paramsOffset)
	return AxsCfgCalcStepParamAllEnd(builder)
}

func (rcv *AxsCfgCalcStepParamAll) UnPackTo(t *AxsCfgCalcStepParamAllT) {
	paramsLength := rcv.ParamsLength()
	t.Params = make([]*AxsCfgCalcStepParamT, paramsLength)
	for j := 0; j < paramsLength; j++ {
		x := AxsCfgCalcStepParam{}
		rcv.Params(&x, j)
		t.Params[j] = x.UnPack()
	}
}

func (rcv *AxsCfgCalcStepParamAll) UnPack() *AxsCfgCalcStepParamAllT {
	if rcv == nil { return nil }
	t := &AxsCfgCalcStepParamAllT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCfgCalcStepParamAll struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgCalcStepParamAll(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgCalcStepParamAll {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgCalcStepParamAll{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgCalcStepParamAll(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgCalcStepParamAll {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgCalcStepParamAll{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgCalcStepParamAll) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgCalcStepParamAll) Table() flatbuffers.Table {
	return rcv._tab
}

/// vector of all params
func (rcv *AxsCfgCalcStepParamAll) Params(obj *AxsCfgCalcStepParam, j int) bool {
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

func (rcv *AxsCfgCalcStepParamAll) ParamsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of all params
func AxsCfgCalcStepParamAllStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func AxsCfgCalcStepParamAllAddParams(builder *flatbuffers.Builder, params flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(params), 0)
}
func AxsCfgCalcStepParamAllStartParamsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func AxsCfgCalcStepParamAllEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
