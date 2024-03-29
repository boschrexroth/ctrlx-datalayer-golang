// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configuration of encoder axis
type AxsCfgEncoderT struct {
	OutputRevLoadGear uint32 `json:"outputRevLoadGear"`
	InputRevLoadGear uint32 `json:"inputRevLoadGear"`
	Resolution uint32 `json:"resolution"`
	FeedConst float64 `json:"feedConst"`
	FilterType FilterType `json:"filterType"`
	CutoffFrq float64 `json:"cutoffFrq"`
	FeedConstUnit string `json:"feedConstUnit"`
	CutoffFrqUnit string `json:"cutoffFrqUnit"`
}

func (t *AxsCfgEncoderT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	feedConstUnitOffset := flatbuffers.UOffsetT(0)
	if t.FeedConstUnit != "" {
		feedConstUnitOffset = builder.CreateString(t.FeedConstUnit)
	}
	cutoffFrqUnitOffset := flatbuffers.UOffsetT(0)
	if t.CutoffFrqUnit != "" {
		cutoffFrqUnitOffset = builder.CreateString(t.CutoffFrqUnit)
	}
	AxsCfgEncoderStart(builder)
	AxsCfgEncoderAddOutputRevLoadGear(builder, t.OutputRevLoadGear)
	AxsCfgEncoderAddInputRevLoadGear(builder, t.InputRevLoadGear)
	AxsCfgEncoderAddResolution(builder, t.Resolution)
	AxsCfgEncoderAddFeedConst(builder, t.FeedConst)
	AxsCfgEncoderAddFilterType(builder, t.FilterType)
	AxsCfgEncoderAddCutoffFrq(builder, t.CutoffFrq)
	AxsCfgEncoderAddFeedConstUnit(builder, feedConstUnitOffset)
	AxsCfgEncoderAddCutoffFrqUnit(builder, cutoffFrqUnitOffset)
	return AxsCfgEncoderEnd(builder)
}

func (rcv *AxsCfgEncoder) UnPackTo(t *AxsCfgEncoderT) {
	t.OutputRevLoadGear = rcv.OutputRevLoadGear()
	t.InputRevLoadGear = rcv.InputRevLoadGear()
	t.Resolution = rcv.Resolution()
	t.FeedConst = rcv.FeedConst()
	t.FilterType = rcv.FilterType()
	t.CutoffFrq = rcv.CutoffFrq()
	t.FeedConstUnit = string(rcv.FeedConstUnit())
	t.CutoffFrqUnit = string(rcv.CutoffFrqUnit())
}

func (rcv *AxsCfgEncoder) UnPack() *AxsCfgEncoderT {
	if rcv == nil { return nil }
	t := &AxsCfgEncoderT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCfgEncoder struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgEncoder(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgEncoder {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgEncoder{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgEncoder(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgEncoder {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgEncoder{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgEncoder) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgEncoder) Table() flatbuffers.Table {
	return rcv._tab
}

/// output revolution of load gear
func (rcv *AxsCfgEncoder) OutputRevLoadGear() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// output revolution of load gear
func (rcv *AxsCfgEncoder) MutateOutputRevLoadGear(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

/// output revolution of load gear
func (rcv *AxsCfgEncoder) InputRevLoadGear() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// output revolution of load gear
func (rcv *AxsCfgEncoder) MutateInputRevLoadGear(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

/// encoder resolution
func (rcv *AxsCfgEncoder) Resolution() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// encoder resolution
func (rcv *AxsCfgEncoder) MutateResolution(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

/// feed constant which only been used as linear axis
func (rcv *AxsCfgEncoder) FeedConst() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// feed constant which only been used as linear axis
func (rcv *AxsCfgEncoder) MutateFeedConst(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

/// filter type
func (rcv *AxsCfgEncoder) FilterType() FilterType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return FilterType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// filter type
func (rcv *AxsCfgEncoder) MutateFilterType(n FilterType) bool {
	return rcv._tab.MutateInt8Slot(12, int8(n))
}

/// cutoff frequency of the filter
func (rcv *AxsCfgEncoder) CutoffFrq() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// cutoff frequency of the filter
func (rcv *AxsCfgEncoder) MutateCutoffFrq(n float64) bool {
	return rcv._tab.MutateFloat64Slot(14, n)
}

/// unit of feed constant
func (rcv *AxsCfgEncoder) FeedConstUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of feed constant
/// unit of cutoff frequency
func (rcv *AxsCfgEncoder) CutoffFrqUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of cutoff frequency
func AxsCfgEncoderStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func AxsCfgEncoderAddOutputRevLoadGear(builder *flatbuffers.Builder, outputRevLoadGear uint32) {
	builder.PrependUint32Slot(0, outputRevLoadGear, 0)
}
func AxsCfgEncoderAddInputRevLoadGear(builder *flatbuffers.Builder, inputRevLoadGear uint32) {
	builder.PrependUint32Slot(1, inputRevLoadGear, 0)
}
func AxsCfgEncoderAddResolution(builder *flatbuffers.Builder, resolution uint32) {
	builder.PrependUint32Slot(2, resolution, 0)
}
func AxsCfgEncoderAddFeedConst(builder *flatbuffers.Builder, feedConst float64) {
	builder.PrependFloat64Slot(3, feedConst, 0.0)
}
func AxsCfgEncoderAddFilterType(builder *flatbuffers.Builder, filterType FilterType) {
	builder.PrependInt8Slot(4, int8(filterType), 0)
}
func AxsCfgEncoderAddCutoffFrq(builder *flatbuffers.Builder, cutoffFrq float64) {
	builder.PrependFloat64Slot(5, cutoffFrq, 0.0)
}
func AxsCfgEncoderAddFeedConstUnit(builder *flatbuffers.Builder, feedConstUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(feedConstUnit), 0)
}
func AxsCfgEncoderAddCutoffFrqUnit(builder *flatbuffers.Builder, cutoffFrqUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(cutoffFrqUnit), 0)
}
func AxsCfgEncoderEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
