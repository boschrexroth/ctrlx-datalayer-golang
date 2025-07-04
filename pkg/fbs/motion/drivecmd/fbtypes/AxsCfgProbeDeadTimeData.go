// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters of dead time compensation (drive)
type AxsCfgProbeDeadTimeDataT struct {
	PosEdgeNs float64 `json:"posEdgeNS"`
	NegEdgeNs float64 `json:"negEdgeNS"`
}

func (t *AxsCfgProbeDeadTimeDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	AxsCfgProbeDeadTimeDataStart(builder)
	AxsCfgProbeDeadTimeDataAddPosEdgeNs(builder, t.PosEdgeNs)
	AxsCfgProbeDeadTimeDataAddNegEdgeNs(builder, t.NegEdgeNs)
	return AxsCfgProbeDeadTimeDataEnd(builder)
}

func (rcv *AxsCfgProbeDeadTimeData) UnPackTo(t *AxsCfgProbeDeadTimeDataT) {
	t.PosEdgeNs = rcv.PosEdgeNs()
	t.NegEdgeNs = rcv.NegEdgeNs()
}

func (rcv *AxsCfgProbeDeadTimeData) UnPack() *AxsCfgProbeDeadTimeDataT {
	if rcv == nil { return nil }
	t := &AxsCfgProbeDeadTimeDataT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCfgProbeDeadTimeData struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgProbeDeadTimeData(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgProbeDeadTimeData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgProbeDeadTimeData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgProbeDeadTimeData(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgProbeDeadTimeData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgProbeDeadTimeData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgProbeDeadTimeData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgProbeDeadTimeData) Table() flatbuffers.Table {
	return rcv._tab
}

/// positive edge dead time compensation to drive in Nano seconds
func (rcv *AxsCfgProbeDeadTimeData) PosEdgeNs() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// positive edge dead time compensation to drive in Nano seconds
func (rcv *AxsCfgProbeDeadTimeData) MutatePosEdgeNs(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// negative edge dead time compensation to drive in Nano seconds
func (rcv *AxsCfgProbeDeadTimeData) NegEdgeNs() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// negative edge dead time compensation to drive in Nano seconds
func (rcv *AxsCfgProbeDeadTimeData) MutateNegEdgeNs(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func AxsCfgProbeDeadTimeDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AxsCfgProbeDeadTimeDataAddPosEdgeNs(builder *flatbuffers.Builder, posEdgeNs float64) {
	builder.PrependFloat64Slot(0, posEdgeNs, 0.0)
}
func AxsCfgProbeDeadTimeDataAddNegEdgeNs(builder *flatbuffers.Builder, negEdgeNs float64) {
	builder.PrependFloat64Slot(1, negEdgeNs, 0.0)
}
func AxsCfgProbeDeadTimeDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
