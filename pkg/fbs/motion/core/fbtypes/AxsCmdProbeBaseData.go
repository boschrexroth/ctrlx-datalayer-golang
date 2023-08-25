// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// base parameters of the axis probe and axis probe abort command
type AxsCmdProbeBaseDataT struct {
	TriggerSrc string `json:"triggerSrc"`
	ProbeIndex string `json:"probeIndex"`
}

func (t *AxsCmdProbeBaseDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	triggerSrcOffset := flatbuffers.UOffsetT(0)
	if t.TriggerSrc != "" {
		triggerSrcOffset = builder.CreateString(t.TriggerSrc)
	}
	probeIndexOffset := flatbuffers.UOffsetT(0)
	if t.ProbeIndex != "" {
		probeIndexOffset = builder.CreateString(t.ProbeIndex)
	}
	AxsCmdProbeBaseDataStart(builder)
	AxsCmdProbeBaseDataAddTriggerSrc(builder, triggerSrcOffset)
	AxsCmdProbeBaseDataAddProbeIndex(builder, probeIndexOffset)
	return AxsCmdProbeBaseDataEnd(builder)
}

func (rcv *AxsCmdProbeBaseData) UnPackTo(t *AxsCmdProbeBaseDataT) {
	t.TriggerSrc = string(rcv.TriggerSrc())
	t.ProbeIndex = string(rcv.ProbeIndex())
}

func (rcv *AxsCmdProbeBaseData) UnPack() *AxsCmdProbeBaseDataT {
	if rcv == nil { return nil }
	t := &AxsCmdProbeBaseDataT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCmdProbeBaseData struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCmdProbeBaseData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdProbeBaseData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCmdProbeBaseData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCmdProbeBaseData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdProbeBaseData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCmdProbeBaseData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCmdProbeBaseData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCmdProbeBaseData) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of trigger source name
func (rcv *AxsCmdProbeBaseData) TriggerSrc() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of trigger source name
/// command of probe index
func (rcv *AxsCmdProbeBaseData) ProbeIndex() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// command of probe index
func AxsCmdProbeBaseDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AxsCmdProbeBaseDataAddTriggerSrc(builder *flatbuffers.Builder, triggerSrc flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(triggerSrc), 0)
}
func AxsCmdProbeBaseDataAddProbeIndex(builder *flatbuffers.Builder, probeIndex flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(probeIndex), 0)
}
func AxsCmdProbeBaseDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
