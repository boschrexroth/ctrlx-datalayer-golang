// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// currently the actual values of the probe
type state_AxsProbeActualValuesReqT struct {
	ObjName string `json:"objName"`
	ProbeIndex string `json:"probeIndex"`
}

func (t *state_AxsProbeActualValuesReqT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	objNameOffset := flatbuffers.UOffsetT(0)
	if t.ObjName != "" {
		objNameOffset = builder.CreateString(t.ObjName)
	}
	probeIndexOffset := flatbuffers.UOffsetT(0)
	if t.ProbeIndex != "" {
		probeIndexOffset = builder.CreateString(t.ProbeIndex)
	}
	state_AxsProbeActualValuesReqStart(builder)
	state_AxsProbeActualValuesReqAddObjName(builder, objNameOffset)
	state_AxsProbeActualValuesReqAddProbeIndex(builder, probeIndexOffset)
	return state_AxsProbeActualValuesReqEnd(builder)
}

func (rcv *state_AxsProbeActualValuesReq) UnPackTo(t *state_AxsProbeActualValuesReqT) {
	t.ObjName = string(rcv.ObjName())
	t.ProbeIndex = string(rcv.ProbeIndex())
}

func (rcv *state_AxsProbeActualValuesReq) UnPack() *state_AxsProbeActualValuesReqT {
	if rcv == nil { return nil }
	t := &state_AxsProbeActualValuesReqT{}
	rcv.UnPackTo(t)
	return t
}

type state_AxsProbeActualValuesReq struct {
	_tab flatbuffers.Table
}

func GetRootAsstate_AxsProbeActualValuesReq(buf []byte, offset flatbuffers.UOffsetT) *state_AxsProbeActualValuesReq {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &state_AxsProbeActualValuesReq{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsstate_AxsProbeActualValuesReq(buf []byte, offset flatbuffers.UOffsetT) *state_AxsProbeActualValuesReq {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &state_AxsProbeActualValuesReq{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *state_AxsProbeActualValuesReq) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *state_AxsProbeActualValuesReq) Table() flatbuffers.Table {
	return rcv._tab
}

/// probe object name
func (rcv *state_AxsProbeActualValuesReq) ObjName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// probe object name
/// probe index
func (rcv *state_AxsProbeActualValuesReq) ProbeIndex() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// probe index
func state_AxsProbeActualValuesReqStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func state_AxsProbeActualValuesReqAddObjName(builder *flatbuffers.Builder, objName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(objName), 0)
}
func state_AxsProbeActualValuesReqAddProbeIndex(builder *flatbuffers.Builder, probeIndex flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(probeIndex), 0)
}
func state_AxsProbeActualValuesReqEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
