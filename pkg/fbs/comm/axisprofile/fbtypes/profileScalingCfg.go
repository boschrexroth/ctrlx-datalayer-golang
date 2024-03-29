// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type profileScalingCfgT struct {
	ScalingStrategy *profileScalingStrategyT `json:"scalingStrategy"`
	ScalingType *profileDeviceScalingTypeT `json:"scalingType"`
	Numerator uint32 `json:"numerator"`
	Denominator uint32 `json:"denominator"`
	Resolution uint32 `json:"resolution"`
}

func (t *profileScalingCfgT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	scalingStrategyOffset := t.ScalingStrategy.Pack(builder)
	scalingTypeOffset := t.ScalingType.Pack(builder)
	profileScalingCfgStart(builder)
	profileScalingCfgAddScalingStrategy(builder, scalingStrategyOffset)
	profileScalingCfgAddScalingType(builder, scalingTypeOffset)
	profileScalingCfgAddNumerator(builder, t.Numerator)
	profileScalingCfgAddDenominator(builder, t.Denominator)
	profileScalingCfgAddResolution(builder, t.Resolution)
	return profileScalingCfgEnd(builder)
}

func (rcv *profileScalingCfg) UnPackTo(t *profileScalingCfgT) {
	t.ScalingStrategy = rcv.ScalingStrategy(nil).UnPack()
	t.ScalingType = rcv.ScalingType(nil).UnPack()
	t.Numerator = rcv.Numerator()
	t.Denominator = rcv.Denominator()
	t.Resolution = rcv.Resolution()
}

func (rcv *profileScalingCfg) UnPack() *profileScalingCfgT {
	if rcv == nil { return nil }
	t := &profileScalingCfgT{}
	rcv.UnPackTo(t)
	return t
}

type profileScalingCfg struct {
	_tab flatbuffers.Table
}

func GetRootAsprofileScalingCfg(buf []byte, offset flatbuffers.UOffsetT) *profileScalingCfg {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &profileScalingCfg{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsprofileScalingCfg(buf []byte, offset flatbuffers.UOffsetT) *profileScalingCfg {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &profileScalingCfg{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *profileScalingCfg) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *profileScalingCfg) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *profileScalingCfg) ScalingStrategy(obj *profileScalingStrategy) *profileScalingStrategy {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(profileScalingStrategy)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *profileScalingCfg) ScalingType(obj *profileDeviceScalingType) *profileDeviceScalingType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(profileDeviceScalingType)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *profileScalingCfg) Numerator() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *profileScalingCfg) MutateNumerator(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *profileScalingCfg) Denominator() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *profileScalingCfg) MutateDenominator(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *profileScalingCfg) Resolution() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *profileScalingCfg) MutateResolution(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func profileScalingCfgStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func profileScalingCfgAddScalingStrategy(builder *flatbuffers.Builder, scalingStrategy flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(scalingStrategy), 0)
}
func profileScalingCfgAddScalingType(builder *flatbuffers.Builder, scalingType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(scalingType), 0)
}
func profileScalingCfgAddNumerator(builder *flatbuffers.Builder, numerator uint32) {
	builder.PrependUint32Slot(2, numerator, 0)
}
func profileScalingCfgAddDenominator(builder *flatbuffers.Builder, denominator uint32) {
	builder.PrependUint32Slot(3, denominator, 0)
}
func profileScalingCfgAddResolution(builder *flatbuffers.Builder, resolution uint32) {
	builder.PrependUint32Slot(4, resolution, 0)
}
func profileScalingCfgEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
