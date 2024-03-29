// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// data of a single configuration parameter of an joint transformation when reading all data of an implemented joint transformation
type KinStateJntTrafoDataAllParamT struct {
	ReqAxes []*KinStateJntTrafoDataReqAxisT `json:"reqAxes"`
	Parameter []*KinStateJntTrafoDataParamT `json:"parameter"`
}

func (t *KinStateJntTrafoDataAllParamT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	reqAxesOffset := flatbuffers.UOffsetT(0)
	if t.ReqAxes != nil {
		reqAxesLength := len(t.ReqAxes)
		reqAxesOffsets := make([]flatbuffers.UOffsetT, reqAxesLength)
		for j := 0; j < reqAxesLength; j++ {
			reqAxesOffsets[j] = t.ReqAxes[j].Pack(builder)
		}
		KinStateJntTrafoDataAllParamStartReqAxesVector(builder, reqAxesLength)
		for j := reqAxesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(reqAxesOffsets[j])
		}
		reqAxesOffset = builder.EndVector(reqAxesLength)
	}
	parameterOffset := flatbuffers.UOffsetT(0)
	if t.Parameter != nil {
		parameterLength := len(t.Parameter)
		parameterOffsets := make([]flatbuffers.UOffsetT, parameterLength)
		for j := 0; j < parameterLength; j++ {
			parameterOffsets[j] = t.Parameter[j].Pack(builder)
		}
		KinStateJntTrafoDataAllParamStartParameterVector(builder, parameterLength)
		for j := parameterLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(parameterOffsets[j])
		}
		parameterOffset = builder.EndVector(parameterLength)
	}
	KinStateJntTrafoDataAllParamStart(builder)
	KinStateJntTrafoDataAllParamAddReqAxes(builder, reqAxesOffset)
	KinStateJntTrafoDataAllParamAddParameter(builder, parameterOffset)
	return KinStateJntTrafoDataAllParamEnd(builder)
}

func (rcv *KinStateJntTrafoDataAllParam) UnPackTo(t *KinStateJntTrafoDataAllParamT) {
	reqAxesLength := rcv.ReqAxesLength()
	t.ReqAxes = make([]*KinStateJntTrafoDataReqAxisT, reqAxesLength)
	for j := 0; j < reqAxesLength; j++ {
		x := KinStateJntTrafoDataReqAxis{}
		rcv.ReqAxes(&x, j)
		t.ReqAxes[j] = x.UnPack()
	}
	parameterLength := rcv.ParameterLength()
	t.Parameter = make([]*KinStateJntTrafoDataParamT, parameterLength)
	for j := 0; j < parameterLength; j++ {
		x := KinStateJntTrafoDataParam{}
		rcv.Parameter(&x, j)
		t.Parameter[j] = x.UnPack()
	}
}

func (rcv *KinStateJntTrafoDataAllParam) UnPack() *KinStateJntTrafoDataAllParamT {
	if rcv == nil { return nil }
	t := &KinStateJntTrafoDataAllParamT{}
	rcv.UnPackTo(t)
	return t
}

type KinStateJntTrafoDataAllParam struct {
	_tab flatbuffers.Table
}

func GetRootAsKinStateJntTrafoDataAllParam(buf []byte, offset flatbuffers.UOffsetT) *KinStateJntTrafoDataAllParam {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinStateJntTrafoDataAllParam{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinStateJntTrafoDataAllParam(buf []byte, offset flatbuffers.UOffsetT) *KinStateJntTrafoDataAllParam {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinStateJntTrafoDataAllParam{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinStateJntTrafoDataAllParam) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinStateJntTrafoDataAllParam) Table() flatbuffers.Table {
	return rcv._tab
}

/// Information of all required axes of this joint transformation
func (rcv *KinStateJntTrafoDataAllParam) ReqAxes(obj *KinStateJntTrafoDataReqAxis, j int) bool {
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

func (rcv *KinStateJntTrafoDataAllParam) ReqAxesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Information of all required axes of this joint transformation
/// Vector of all configuration parameters of the joint transformation
func (rcv *KinStateJntTrafoDataAllParam) Parameter(obj *KinStateJntTrafoDataParam, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *KinStateJntTrafoDataAllParam) ParameterLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Vector of all configuration parameters of the joint transformation
func KinStateJntTrafoDataAllParamStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func KinStateJntTrafoDataAllParamAddReqAxes(builder *flatbuffers.Builder, reqAxes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(reqAxes), 0)
}
func KinStateJntTrafoDataAllParamStartReqAxesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KinStateJntTrafoDataAllParamAddParameter(builder *flatbuffers.Builder, parameter flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(parameter), 0)
}
func KinStateJntTrafoDataAllParamStartParameterVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KinStateJntTrafoDataAllParamEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
