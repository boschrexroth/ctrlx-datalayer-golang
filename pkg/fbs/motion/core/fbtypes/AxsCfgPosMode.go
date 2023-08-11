// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configuration for position mode function parameters for a single axis
type AxsCfgPosModeT struct {
	TargetPosWindowModulo float64 `json:"targetPosWindowModulo"`
	TargetPosWindowModuloUnit string `json:"targetPosWindowModuloUnit"`
}

func (t *AxsCfgPosModeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	targetPosWindowModuloUnitOffset := flatbuffers.UOffsetT(0)
	if t.TargetPosWindowModuloUnit != "" {
		targetPosWindowModuloUnitOffset = builder.CreateString(t.TargetPosWindowModuloUnit)
	}
	AxsCfgPosModeStart(builder)
	AxsCfgPosModeAddTargetPosWindowModulo(builder, t.TargetPosWindowModulo)
	AxsCfgPosModeAddTargetPosWindowModuloUnit(builder, targetPosWindowModuloUnitOffset)
	return AxsCfgPosModeEnd(builder)
}

func (rcv *AxsCfgPosMode) UnPackTo(t *AxsCfgPosModeT) {
	t.TargetPosWindowModulo = rcv.TargetPosWindowModulo()
	t.TargetPosWindowModuloUnit = string(rcv.TargetPosWindowModuloUnit())
}

func (rcv *AxsCfgPosMode) UnPack() *AxsCfgPosModeT {
	if rcv == nil { return nil }
	t := &AxsCfgPosModeT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCfgPosMode struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgPosMode(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgPosMode {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgPosMode{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgPosMode(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgPosMode {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgPosMode{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgPosMode) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgPosMode) Table() flatbuffers.Table {
	return rcv._tab
}

/// the target position window for modulo axes (modulo axes always use shortest way, when inside this window)
func (rcv *AxsCfgPosMode) TargetPosWindowModulo() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// the target position window for modulo axes (modulo axes always use shortest way, when inside this window)
func (rcv *AxsCfgPosMode) MutateTargetPosWindowModulo(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// unit of the target position window for modulo axes
func (rcv *AxsCfgPosMode) TargetPosWindowModuloUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the target position window for modulo axes
func AxsCfgPosModeStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AxsCfgPosModeAddTargetPosWindowModulo(builder *flatbuffers.Builder, targetPosWindowModulo float64) {
	builder.PrependFloat64Slot(0, targetPosWindowModulo, 0.0)
}
func AxsCfgPosModeAddTargetPosWindowModuloUnit(builder *flatbuffers.Builder, targetPosWindowModuloUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(targetPosWindowModuloUnit), 0)
}
func AxsCfgPosModeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}