// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters of touch probe
type AxsCmdProbeExDataT struct {
	ProbeIndex ProbeIndexType `json:"probeIndex"`
}

func (t *AxsCmdProbeExDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	AxsCmdProbeExDataStart(builder)
	AxsCmdProbeExDataAddProbeIndex(builder, t.ProbeIndex)
	return AxsCmdProbeExDataEnd(builder)
}

func (rcv *AxsCmdProbeExData) UnPackTo(t *AxsCmdProbeExDataT) {
	t.ProbeIndex = rcv.ProbeIndex()
}

func (rcv *AxsCmdProbeExData) UnPack() *AxsCmdProbeExDataT {
	if rcv == nil { return nil }
	t := &AxsCmdProbeExDataT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCmdProbeExData struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCmdProbeExData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdProbeExData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCmdProbeExData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCmdProbeExData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdProbeExData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCmdProbeExData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCmdProbeExData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCmdProbeExData) Table() flatbuffers.Table {
	return rcv._tab
}

/// probe index
func (rcv *AxsCmdProbeExData) ProbeIndex() ProbeIndexType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return ProbeIndexType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// probe index
func (rcv *AxsCmdProbeExData) MutateProbeIndex(n ProbeIndexType) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func AxsCmdProbeExDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func AxsCmdProbeExDataAddProbeIndex(builder *flatbuffers.Builder, probeIndex ProbeIndexType) {
	builder.PrependInt8Slot(0, int8(probeIndex), 0)
}
func AxsCmdProbeExDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}