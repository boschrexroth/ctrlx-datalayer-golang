// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters of single probe data
type AxsCfgProbeSignalDataT struct {
	Select SignalSelect `json:"select"`
	ShotType ShotType `json:"shotType"`
	Source string `json:"source"`
	SourceInfo string `json:"sourceInfo"`
}

func (t *AxsCfgProbeSignalDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	sourceOffset := flatbuffers.UOffsetT(0)
	if t.Source != "" {
		sourceOffset = builder.CreateString(t.Source)
	}
	sourceInfoOffset := flatbuffers.UOffsetT(0)
	if t.SourceInfo != "" {
		sourceInfoOffset = builder.CreateString(t.SourceInfo)
	}
	AxsCfgProbeSignalDataStart(builder)
	AxsCfgProbeSignalDataAddSelect(builder, t.Select)
	AxsCfgProbeSignalDataAddShotType(builder, t.ShotType)
	AxsCfgProbeSignalDataAddSource(builder, sourceOffset)
	AxsCfgProbeSignalDataAddSourceInfo(builder, sourceInfoOffset)
	return AxsCfgProbeSignalDataEnd(builder)
}

func (rcv *AxsCfgProbeSignalData) UnPackTo(t *AxsCfgProbeSignalDataT) {
	t.Select = rcv.Select()
	t.ShotType = rcv.ShotType()
	t.Source = string(rcv.Source())
	t.SourceInfo = string(rcv.SourceInfo())
}

func (rcv *AxsCfgProbeSignalData) UnPack() *AxsCfgProbeSignalDataT {
	if rcv == nil { return nil }
	t := &AxsCfgProbeSignalDataT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCfgProbeSignalData struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgProbeSignalData(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgProbeSignalData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgProbeSignalData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgProbeSignalData(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgProbeSignalData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgProbeSignalData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgProbeSignalData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgProbeSignalData) Table() flatbuffers.Table {
	return rcv._tab
}

/// encoder1, encoder2, fine time, ipo-pos, ipo-vel
func (rcv *AxsCfgProbeSignalData) Select() SignalSelect {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return SignalSelect(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 1
}

/// encoder1, encoder2, fine time, ipo-pos, ipo-vel
func (rcv *AxsCfgProbeSignalData) MutateSelect(n SignalSelect) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

/// single shot or continuous running
func (rcv *AxsCfgProbeSignalData) ShotType() ShotType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return ShotType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 2
}

/// single shot or continuous running
func (rcv *AxsCfgProbeSignalData) MutateShotType(n ShotType) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

/// self, axis_1, IO
func (rcv *AxsCfgProbeSignalData) Source() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// self, axis_1, IO
/// uri and other addition information
func (rcv *AxsCfgProbeSignalData) SourceInfo() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// uri and other addition information
func AxsCfgProbeSignalDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func AxsCfgProbeSignalDataAddSelect(builder *flatbuffers.Builder, select_ SignalSelect) {
	builder.PrependInt8Slot(0, int8(select_), 1)
}
func AxsCfgProbeSignalDataAddShotType(builder *flatbuffers.Builder, shotType ShotType) {
	builder.PrependInt8Slot(1, int8(shotType), 2)
}
func AxsCfgProbeSignalDataAddSource(builder *flatbuffers.Builder, source flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(source), 0)
}
func AxsCfgProbeSignalDataAddSourceInfo(builder *flatbuffers.Builder, sourceInfo flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(sourceInfo), 0)
}
func AxsCfgProbeSignalDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
