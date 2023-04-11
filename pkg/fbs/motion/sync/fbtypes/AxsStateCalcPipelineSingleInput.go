// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// single input values for reading the pipeline
type AxsStateCalcPipelineSingleInputT struct {
	StepID uint32
	DataID DataID
	Value float64
}

func (t *AxsStateCalcPipelineSingleInputT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	AxsStateCalcPipelineSingleInputStart(builder)
	AxsStateCalcPipelineSingleInputAddStepID(builder, t.StepID)
	AxsStateCalcPipelineSingleInputAddDataID(builder, t.DataID)
	AxsStateCalcPipelineSingleInputAddValue(builder, t.Value)
	return AxsStateCalcPipelineSingleInputEnd(builder)
}

func (rcv *AxsStateCalcPipelineSingleInput) UnPackTo(t *AxsStateCalcPipelineSingleInputT) {
	t.StepID = rcv.StepID()
	t.DataID = rcv.DataID()
	t.Value = rcv.Value()
}

func (rcv *AxsStateCalcPipelineSingleInput) UnPack() *AxsStateCalcPipelineSingleInputT {
	if rcv == nil { return nil }
	t := &AxsStateCalcPipelineSingleInputT{}
	rcv.UnPackTo(t)
	return t
}

type AxsStateCalcPipelineSingleInput struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsStateCalcPipelineSingleInput(buf []byte, offset flatbuffers.UOffsetT) *AxsStateCalcPipelineSingleInput {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsStateCalcPipelineSingleInput{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsStateCalcPipelineSingleInput(buf []byte, offset flatbuffers.UOffsetT) *AxsStateCalcPipelineSingleInput {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsStateCalcPipelineSingleInput{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsStateCalcPipelineSingleInput) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsStateCalcPipelineSingleInput) Table() flatbuffers.Table {
	return rcv._tab
}

/// master input id
func (rcv *AxsStateCalcPipelineSingleInput) StepID() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// master input id
func (rcv *AxsStateCalcPipelineSingleInput) MutateStepID(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

/// id of the requested data
func (rcv *AxsStateCalcPipelineSingleInput) DataID() DataID {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return DataID(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// id of the requested data
func (rcv *AxsStateCalcPipelineSingleInput) MutateDataID(n DataID) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

/// the value itself
func (rcv *AxsStateCalcPipelineSingleInput) Value() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// the value itself
func (rcv *AxsStateCalcPipelineSingleInput) MutateValue(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func AxsStateCalcPipelineSingleInputStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func AxsStateCalcPipelineSingleInputAddStepID(builder *flatbuffers.Builder, stepID uint32) {
	builder.PrependUint32Slot(0, stepID, 0)
}
func AxsStateCalcPipelineSingleInputAddDataID(builder *flatbuffers.Builder, dataID DataID) {
	builder.PrependInt8Slot(1, int8(dataID), 0)
}
func AxsStateCalcPipelineSingleInputAddValue(builder *flatbuffers.Builder, value float64) {
	builder.PrependFloat64Slot(2, value, 0.0)
}
func AxsStateCalcPipelineSingleInputEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
