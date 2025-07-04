// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// a single parameter of a calculation step, when writing it, only name value and unit should be set.
type AxsCfgCalcStepParamT struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Unit string `json:"unit"`
	Description string `json:"description"`
	Mandatory bool `json:"mandatory"`
	Type ParameterType `json:"type"`
}

func (t *AxsCfgCalcStepParamT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	valueOffset := flatbuffers.UOffsetT(0)
	if t.Value != "" {
		valueOffset = builder.CreateString(t.Value)
	}
	unitOffset := flatbuffers.UOffsetT(0)
	if t.Unit != "" {
		unitOffset = builder.CreateString(t.Unit)
	}
	descriptionOffset := flatbuffers.UOffsetT(0)
	if t.Description != "" {
		descriptionOffset = builder.CreateString(t.Description)
	}
	AxsCfgCalcStepParamStart(builder)
	AxsCfgCalcStepParamAddName(builder, nameOffset)
	AxsCfgCalcStepParamAddValue(builder, valueOffset)
	AxsCfgCalcStepParamAddUnit(builder, unitOffset)
	AxsCfgCalcStepParamAddDescription(builder, descriptionOffset)
	AxsCfgCalcStepParamAddMandatory(builder, t.Mandatory)
	AxsCfgCalcStepParamAddType(builder, t.Type)
	return AxsCfgCalcStepParamEnd(builder)
}

func (rcv *AxsCfgCalcStepParam) UnPackTo(t *AxsCfgCalcStepParamT) {
	t.Name = string(rcv.Name())
	t.Value = string(rcv.Value())
	t.Unit = string(rcv.Unit())
	t.Description = string(rcv.Description())
	t.Mandatory = rcv.Mandatory()
	t.Type = rcv.Type()
}

func (rcv *AxsCfgCalcStepParam) UnPack() *AxsCfgCalcStepParamT {
	if rcv == nil { return nil }
	t := &AxsCfgCalcStepParamT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCfgCalcStepParam struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgCalcStepParam(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgCalcStepParam {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgCalcStepParam{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgCalcStepParam(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgCalcStepParam {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgCalcStepParam{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgCalcStepParam) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgCalcStepParam) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the parameter
func (rcv *AxsCfgCalcStepParam) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the parameter
/// value of the parameter
func (rcv *AxsCfgCalcStepParam) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// value of the parameter
/// unit of the parameter
func (rcv *AxsCfgCalcStepParam) Unit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the parameter
/// what the parameter of the calculation step does, only for reading
func (rcv *AxsCfgCalcStepParam) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// what the parameter of the calculation step does, only for reading
/// is this parameter mandatory, only for reading
func (rcv *AxsCfgCalcStepParam) Mandatory() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// is this parameter mandatory, only for reading
func (rcv *AxsCfgCalcStepParam) MutateMandatory(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

/// type of the parameter, only for reading
func (rcv *AxsCfgCalcStepParam) Type() ParameterType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return ParameterType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// type of the parameter, only for reading
func (rcv *AxsCfgCalcStepParam) MutateType(n ParameterType) bool {
	return rcv._tab.MutateInt8Slot(14, int8(n))
}

func AxsCfgCalcStepParamStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func AxsCfgCalcStepParamAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func AxsCfgCalcStepParamAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func AxsCfgCalcStepParamAddUnit(builder *flatbuffers.Builder, unit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(unit), 0)
}
func AxsCfgCalcStepParamAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(description), 0)
}
func AxsCfgCalcStepParamAddMandatory(builder *flatbuffers.Builder, mandatory bool) {
	builder.PrependBoolSlot(4, mandatory, false)
}
func AxsCfgCalcStepParamAddType(builder *flatbuffers.Builder, type_ ParameterType) {
	builder.PrependInt8Slot(5, int8(type_), 0)
}
func AxsCfgCalcStepParamEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
