// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// common configuration of orientation
type KinOriCfgT struct {
	EffectiveRadius *KinOriRadiusT `json:"effective_radius"`
	Lim *KinCfgLimitsT `json:"lim"`
	Units *UnitCfgObjT `json:"units"`
	RotToLinConversion *KinCfgRotToLinConversionT `json:"rotToLinConversion"`
}

func (t *KinOriCfgT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	effectiveRadiusOffset := t.EffectiveRadius.Pack(builder)
	limOffset := t.Lim.Pack(builder)
	unitsOffset := t.Units.Pack(builder)
	rotToLinConversionOffset := t.RotToLinConversion.Pack(builder)
	KinOriCfgStart(builder)
	KinOriCfgAddEffectiveRadius(builder, effectiveRadiusOffset)
	KinOriCfgAddLim(builder, limOffset)
	KinOriCfgAddUnits(builder, unitsOffset)
	KinOriCfgAddRotToLinConversion(builder, rotToLinConversionOffset)
	return KinOriCfgEnd(builder)
}

func (rcv *KinOriCfg) UnPackTo(t *KinOriCfgT) {
	t.EffectiveRadius = rcv.EffectiveRadius(nil).UnPack()
	t.Lim = rcv.Lim(nil).UnPack()
	t.Units = rcv.Units(nil).UnPack()
	t.RotToLinConversion = rcv.RotToLinConversion(nil).UnPack()
}

func (rcv *KinOriCfg) UnPack() *KinOriCfgT {
	if rcv == nil { return nil }
	t := &KinOriCfgT{}
	rcv.UnPackTo(t)
	return t
}

type KinOriCfg struct {
	_tab flatbuffers.Table
}

func GetRootAsKinOriCfg(buf []byte, offset flatbuffers.UOffsetT) *KinOriCfg {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinOriCfg{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinOriCfg(buf []byte, offset flatbuffers.UOffsetT) *KinOriCfg {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinOriCfg{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinOriCfg) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinOriCfg) Table() flatbuffers.Table {
	return rcv._tab
}

/// DEPRECATED; Do not use! Values in this structure are ignored.
func (rcv *KinOriCfg) EffectiveRadius(obj *KinOriRadius) *KinOriRadius {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(KinOriRadius)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// DEPRECATED; Do not use! Values in this structure are ignored.
/// max orientation vel
func (rcv *KinOriCfg) Lim(obj *KinCfgLimits) *KinCfgLimits {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(KinCfgLimits)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// max orientation vel
/// general orientation unit configuration of this kinematics
func (rcv *KinOriCfg) Units(obj *UnitCfgObj) *UnitCfgObj {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(UnitCfgObj)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// general orientation unit configuration of this kinematics
/// Rotation to Linear conversion factor
func (rcv *KinOriCfg) RotToLinConversion(obj *KinCfgRotToLinConversion) *KinCfgRotToLinConversion {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(KinCfgRotToLinConversion)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Rotation to Linear conversion factor
func KinOriCfgStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func KinOriCfgAddEffectiveRadius(builder *flatbuffers.Builder, effectiveRadius flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(effectiveRadius), 0)
}
func KinOriCfgAddLim(builder *flatbuffers.Builder, lim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(lim), 0)
}
func KinOriCfgAddUnits(builder *flatbuffers.Builder, units flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(units), 0)
}
func KinOriCfgAddRotToLinConversion(builder *flatbuffers.Builder, rotToLinConversion flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(rotToLinConversion), 0)
}
func KinOriCfgEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
