// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// common axis properties for a single axis
type AxsCfgPropertiesT struct {
	AxsType string `json:"axsType"`
	Modulo bool `json:"modulo"`
	ModuloValue float64 `json:"moduloValue"`
	ModuloValueUnit string `json:"moduloValueUnit"`
	AxsCategory *AxsCategoryT `json:"axsCategory"`
}

func (t *AxsCfgPropertiesT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	axsTypeOffset := flatbuffers.UOffsetT(0)
	if t.AxsType != "" {
		axsTypeOffset = builder.CreateString(t.AxsType)
	}
	moduloValueUnitOffset := flatbuffers.UOffsetT(0)
	if t.ModuloValueUnit != "" {
		moduloValueUnitOffset = builder.CreateString(t.ModuloValueUnit)
	}
	axsCategoryOffset := t.AxsCategory.Pack(builder)
	AxsCfgPropertiesStart(builder)
	AxsCfgPropertiesAddAxsType(builder, axsTypeOffset)
	AxsCfgPropertiesAddModulo(builder, t.Modulo)
	AxsCfgPropertiesAddModuloValue(builder, t.ModuloValue)
	AxsCfgPropertiesAddModuloValueUnit(builder, moduloValueUnitOffset)
	AxsCfgPropertiesAddAxsCategory(builder, axsCategoryOffset)
	return AxsCfgPropertiesEnd(builder)
}

func (rcv *AxsCfgProperties) UnPackTo(t *AxsCfgPropertiesT) {
	t.AxsType = string(rcv.AxsType())
	t.Modulo = rcv.Modulo()
	t.ModuloValue = rcv.ModuloValue()
	t.ModuloValueUnit = string(rcv.ModuloValueUnit())
	t.AxsCategory = rcv.AxsCategory(nil).UnPack()
}

func (rcv *AxsCfgProperties) UnPack() *AxsCfgPropertiesT {
	if rcv == nil { return nil }
	t := &AxsCfgPropertiesT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCfgProperties struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgProperties(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgProperties {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgProperties{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgProperties(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgProperties {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgProperties{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgProperties) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgProperties) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the axis type (e.g. "LINEAR")
func (rcv *AxsCfgProperties) AxsType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the axis type (e.g. "LINEAR")
/// is this a modulo axis?
func (rcv *AxsCfgProperties) Modulo() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// is this a modulo axis?
func (rcv *AxsCfgProperties) MutateModulo(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

/// the modulo value, if the axis is a modulo axis
func (rcv *AxsCfgProperties) ModuloValue() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// the modulo value, if the axis is a modulo axis
func (rcv *AxsCfgProperties) MutateModuloValue(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

/// unit of moduloValue
func (rcv *AxsCfgProperties) ModuloValueUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of moduloValue
/// axis category (e.g. "ENCODERAXS")
func (rcv *AxsCfgProperties) AxsCategory(obj *AxsCategory) *AxsCategory {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AxsCategory)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// axis category (e.g. "ENCODERAXS")
func AxsCfgPropertiesStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func AxsCfgPropertiesAddAxsType(builder *flatbuffers.Builder, axsType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(axsType), 0)
}
func AxsCfgPropertiesAddModulo(builder *flatbuffers.Builder, modulo bool) {
	builder.PrependBoolSlot(1, modulo, false)
}
func AxsCfgPropertiesAddModuloValue(builder *flatbuffers.Builder, moduloValue float64) {
	builder.PrependFloat64Slot(2, moduloValue, 0.0)
}
func AxsCfgPropertiesAddModuloValueUnit(builder *flatbuffers.Builder, moduloValueUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(moduloValueUnit), 0)
}
func AxsCfgPropertiesAddAxsCategory(builder *flatbuffers.Builder, axsCategory flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(axsCategory), 0)
}
func AxsCfgPropertiesEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
