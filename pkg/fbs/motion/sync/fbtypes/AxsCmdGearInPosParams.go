// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// GearInPos calculation pipeline parameters for the axis GearInPos command
type AxsCmdGearInPosParamsT struct {
	MasterOffset float64 `json:"masterOffset"`
	SlaveOffset float64 `json:"slaveOffset"`
	RatioNumerator int32 `json:"ratioNumerator"`
	RatioDenominator int32 `json:"ratioDenominator"`
	FineAdjust float64 `json:"fineAdjust"`
}

func (t *AxsCmdGearInPosParamsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	AxsCmdGearInPosParamsStart(builder)
	AxsCmdGearInPosParamsAddMasterOffset(builder, t.MasterOffset)
	AxsCmdGearInPosParamsAddSlaveOffset(builder, t.SlaveOffset)
	AxsCmdGearInPosParamsAddRatioNumerator(builder, t.RatioNumerator)
	AxsCmdGearInPosParamsAddRatioDenominator(builder, t.RatioDenominator)
	AxsCmdGearInPosParamsAddFineAdjust(builder, t.FineAdjust)
	return AxsCmdGearInPosParamsEnd(builder)
}

func (rcv *AxsCmdGearInPosParams) UnPackTo(t *AxsCmdGearInPosParamsT) {
	t.MasterOffset = rcv.MasterOffset()
	t.SlaveOffset = rcv.SlaveOffset()
	t.RatioNumerator = rcv.RatioNumerator()
	t.RatioDenominator = rcv.RatioDenominator()
	t.FineAdjust = rcv.FineAdjust()
}

func (rcv *AxsCmdGearInPosParams) UnPack() *AxsCmdGearInPosParamsT {
	if rcv == nil { return nil }
	t := &AxsCmdGearInPosParamsT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCmdGearInPosParams struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCmdGearInPosParams(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdGearInPosParams {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCmdGearInPosParams{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCmdGearInPosParams(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdGearInPosParams {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCmdGearInPosParams{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCmdGearInPosParams) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCmdGearInPosParams) Table() flatbuffers.Table {
	return rcv._tab
}

/// master offset value
func (rcv *AxsCmdGearInPosParams) MasterOffset() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// master offset value
func (rcv *AxsCmdGearInPosParams) MutateMasterOffset(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// slave offset value
func (rcv *AxsCmdGearInPosParams) SlaveOffset() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// slave offset value
func (rcv *AxsCmdGearInPosParams) MutateSlaveOffset(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// ratio numerator value
func (rcv *AxsCmdGearInPosParams) RatioNumerator() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 1
}

/// ratio numerator value
func (rcv *AxsCmdGearInPosParams) MutateRatioNumerator(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

/// ratio denominator value
func (rcv *AxsCmdGearInPosParams) RatioDenominator() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 1
}

/// ratio denominator value
func (rcv *AxsCmdGearInPosParams) MutateRatioDenominator(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

/// fine adjust parameter
func (rcv *AxsCmdGearInPosParams) FineAdjust() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// fine adjust parameter
func (rcv *AxsCmdGearInPosParams) MutateFineAdjust(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

func AxsCmdGearInPosParamsStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func AxsCmdGearInPosParamsAddMasterOffset(builder *flatbuffers.Builder, masterOffset float64) {
	builder.PrependFloat64Slot(0, masterOffset, 0.0)
}
func AxsCmdGearInPosParamsAddSlaveOffset(builder *flatbuffers.Builder, slaveOffset float64) {
	builder.PrependFloat64Slot(1, slaveOffset, 0.0)
}
func AxsCmdGearInPosParamsAddRatioNumerator(builder *flatbuffers.Builder, ratioNumerator int32) {
	builder.PrependInt32Slot(2, ratioNumerator, 1)
}
func AxsCmdGearInPosParamsAddRatioDenominator(builder *flatbuffers.Builder, ratioDenominator int32) {
	builder.PrependInt32Slot(3, ratioDenominator, 1)
}
func AxsCmdGearInPosParamsAddFineAdjust(builder *flatbuffers.Builder, fineAdjust float64) {
	builder.PrependFloat64Slot(4, fineAdjust, 0.0)
}
func AxsCmdGearInPosParamsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
