// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// the describe parameters of filter
type DescribeParamT struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Mandatory bool `json:"mandatory"`
	Datatype string `json:"datatype"`
	Unit string `json:"unit"`
}

func (t *DescribeParamT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	descriptionOffset := flatbuffers.UOffsetT(0)
	if t.Description != "" {
		descriptionOffset = builder.CreateString(t.Description)
	}
	datatypeOffset := flatbuffers.UOffsetT(0)
	if t.Datatype != "" {
		datatypeOffset = builder.CreateString(t.Datatype)
	}
	unitOffset := flatbuffers.UOffsetT(0)
	if t.Unit != "" {
		unitOffset = builder.CreateString(t.Unit)
	}
	DescribeParamStart(builder)
	DescribeParamAddName(builder, nameOffset)
	DescribeParamAddDescription(builder, descriptionOffset)
	DescribeParamAddMandatory(builder, t.Mandatory)
	DescribeParamAddDatatype(builder, datatypeOffset)
	DescribeParamAddUnit(builder, unitOffset)
	return DescribeParamEnd(builder)
}

func (rcv *DescribeParam) UnPackTo(t *DescribeParamT) {
	t.Name = string(rcv.Name())
	t.Description = string(rcv.Description())
	t.Mandatory = rcv.Mandatory()
	t.Datatype = string(rcv.Datatype())
	t.Unit = string(rcv.Unit())
}

func (rcv *DescribeParam) UnPack() *DescribeParamT {
	if rcv == nil { return nil }
	t := &DescribeParamT{}
	rcv.UnPackTo(t)
	return t
}

type DescribeParam struct {
	_tab flatbuffers.Table
}

func GetRootAsDescribeParam(buf []byte, offset flatbuffers.UOffsetT) *DescribeParam {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DescribeParam{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDescribeParam(buf []byte, offset flatbuffers.UOffsetT) *DescribeParam {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DescribeParam{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DescribeParam) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DescribeParam) Table() flatbuffers.Table {
	return rcv._tab
}

/// Name of current parameter
func (rcv *DescribeParam) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Name of current parameter
/// Description of current filter parameter
func (rcv *DescribeParam) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Description of current filter parameter
/// Is this parameter mandatory
func (rcv *DescribeParam) Mandatory() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Is this parameter mandatory
func (rcv *DescribeParam) MutateMandatory(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

/// Type of the variable
func (rcv *DescribeParam) Datatype() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Type of the variable
/// unit
func (rcv *DescribeParam) Unit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit
func DescribeParamStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func DescribeParamAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func DescribeParamAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(description), 0)
}
func DescribeParamAddMandatory(builder *flatbuffers.Builder, mandatory bool) {
	builder.PrependBoolSlot(2, mandatory, false)
}
func DescribeParamAddDatatype(builder *flatbuffers.Builder, datatype flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(datatype), 0)
}
func DescribeParamAddUnit(builder *flatbuffers.Builder, unit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(unit), 0)
}
func DescribeParamEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
