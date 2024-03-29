// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configuration of a single parameter of a joint transformation
type KinCfgJntTrafoSingleParamT struct {
	Name string `json:"name"`
	ValueDouble float64 `json:"valueDouble"`
	ValueInt int64 `json:"valueInt"`
	ValueString string `json:"valueString"`
	Unit string `json:"unit"`
}

func (t *KinCfgJntTrafoSingleParamT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	valueStringOffset := flatbuffers.UOffsetT(0)
	if t.ValueString != "" {
		valueStringOffset = builder.CreateString(t.ValueString)
	}
	unitOffset := flatbuffers.UOffsetT(0)
	if t.Unit != "" {
		unitOffset = builder.CreateString(t.Unit)
	}
	KinCfgJntTrafoSingleParamStart(builder)
	KinCfgJntTrafoSingleParamAddName(builder, nameOffset)
	KinCfgJntTrafoSingleParamAddValueDouble(builder, t.ValueDouble)
	KinCfgJntTrafoSingleParamAddValueInt(builder, t.ValueInt)
	KinCfgJntTrafoSingleParamAddValueString(builder, valueStringOffset)
	KinCfgJntTrafoSingleParamAddUnit(builder, unitOffset)
	return KinCfgJntTrafoSingleParamEnd(builder)
}

func (rcv *KinCfgJntTrafoSingleParam) UnPackTo(t *KinCfgJntTrafoSingleParamT) {
	t.Name = string(rcv.Name())
	t.ValueDouble = rcv.ValueDouble()
	t.ValueInt = rcv.ValueInt()
	t.ValueString = string(rcv.ValueString())
	t.Unit = string(rcv.Unit())
}

func (rcv *KinCfgJntTrafoSingleParam) UnPack() *KinCfgJntTrafoSingleParamT {
	if rcv == nil { return nil }
	t := &KinCfgJntTrafoSingleParamT{}
	rcv.UnPackTo(t)
	return t
}

type KinCfgJntTrafoSingleParam struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCfgJntTrafoSingleParam(buf []byte, offset flatbuffers.UOffsetT) *KinCfgJntTrafoSingleParam {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCfgJntTrafoSingleParam{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCfgJntTrafoSingleParam(buf []byte, offset flatbuffers.UOffsetT) *KinCfgJntTrafoSingleParam {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCfgJntTrafoSingleParam{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCfgJntTrafoSingleParam) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCfgJntTrafoSingleParam) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the parameter
func (rcv *KinCfgJntTrafoSingleParam) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the parameter
/// value of the parameter (when it's a double value)
func (rcv *KinCfgJntTrafoSingleParam) ValueDouble() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// value of the parameter (when it's a double value)
func (rcv *KinCfgJntTrafoSingleParam) MutateValueDouble(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// value of the parameter (when it's a integer value)
func (rcv *KinCfgJntTrafoSingleParam) ValueInt() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// value of the parameter (when it's a integer value)
func (rcv *KinCfgJntTrafoSingleParam) MutateValueInt(n int64) bool {
	return rcv._tab.MutateInt64Slot(8, n)
}

/// value of the parameter (when it's a string value)
func (rcv *KinCfgJntTrafoSingleParam) ValueString() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// value of the parameter (when it's a string value)
/// unit of the parameter
func (rcv *KinCfgJntTrafoSingleParam) Unit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the parameter
func KinCfgJntTrafoSingleParamStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func KinCfgJntTrafoSingleParamAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func KinCfgJntTrafoSingleParamAddValueDouble(builder *flatbuffers.Builder, valueDouble float64) {
	builder.PrependFloat64Slot(1, valueDouble, 0.0)
}
func KinCfgJntTrafoSingleParamAddValueInt(builder *flatbuffers.Builder, valueInt int64) {
	builder.PrependInt64Slot(2, valueInt, 0)
}
func KinCfgJntTrafoSingleParamAddValueString(builder *flatbuffers.Builder, valueString flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(valueString), 0)
}
func KinCfgJntTrafoSingleParamAddUnit(builder *flatbuffers.Builder, unit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(unit), 0)
}
func KinCfgJntTrafoSingleParamEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
