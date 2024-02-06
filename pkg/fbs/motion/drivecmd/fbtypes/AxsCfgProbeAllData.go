// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters of all probe data
type AxsCfgProbeAllDataT struct {
	Probe1 *AxsCfgProbeDataT `json:"probe1"`
	Probe2 *AxsCfgProbeDataT `json:"probe2"`
}

func (t *AxsCfgProbeAllDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	probe1Offset := t.Probe1.Pack(builder)
	probe2Offset := t.Probe2.Pack(builder)
	AxsCfgProbeAllDataStart(builder)
	AxsCfgProbeAllDataAddProbe1(builder, probe1Offset)
	AxsCfgProbeAllDataAddProbe2(builder, probe2Offset)
	return AxsCfgProbeAllDataEnd(builder)
}

func (rcv *AxsCfgProbeAllData) UnPackTo(t *AxsCfgProbeAllDataT) {
	t.Probe1 = rcv.Probe1(nil).UnPack()
	t.Probe2 = rcv.Probe2(nil).UnPack()
}

func (rcv *AxsCfgProbeAllData) UnPack() *AxsCfgProbeAllDataT {
	if rcv == nil { return nil }
	t := &AxsCfgProbeAllDataT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCfgProbeAllData struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgProbeAllData(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgProbeAllData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgProbeAllData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgProbeAllData(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgProbeAllData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgProbeAllData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgProbeAllData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgProbeAllData) Table() flatbuffers.Table {
	return rcv._tab
}

/// all probes
func (rcv *AxsCfgProbeAllData) Probe1(obj *AxsCfgProbeData) *AxsCfgProbeData {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AxsCfgProbeData)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// all probes
func (rcv *AxsCfgProbeAllData) Probe2(obj *AxsCfgProbeData) *AxsCfgProbeData {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AxsCfgProbeData)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func AxsCfgProbeAllDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AxsCfgProbeAllDataAddProbe1(builder *flatbuffers.Builder, probe1 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(probe1), 0)
}
func AxsCfgProbeAllDataAddProbe2(builder *flatbuffers.Builder, probe2 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(probe2), 0)
}
func AxsCfgProbeAllDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
