// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs2

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Utilization of the cycle time of Scheduler tick task causes by itself
type UtilizationThresholdT struct {
	Warning byte `json:"warning"`
	Error byte `json:"error"`
}

func (t *UtilizationThresholdT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	UtilizationThresholdStart(builder)
	UtilizationThresholdAddWarning(builder, t.Warning)
	UtilizationThresholdAddError(builder, t.Error)
	return UtilizationThresholdEnd(builder)
}

func (rcv *UtilizationThreshold) UnPackTo(t *UtilizationThresholdT) {
	t.Warning = rcv.Warning()
	t.Error = rcv.Error()
}

func (rcv *UtilizationThreshold) UnPack() *UtilizationThresholdT {
	if rcv == nil { return nil }
	t := &UtilizationThresholdT{}
	rcv.UnPackTo(t)
	return t
}

type UtilizationThreshold struct {
	_tab flatbuffers.Table
}

func GetRootAsUtilizationThreshold(buf []byte, offset flatbuffers.UOffsetT) *UtilizationThreshold {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UtilizationThreshold{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUtilizationThreshold(buf []byte, offset flatbuffers.UOffsetT) *UtilizationThreshold {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &UtilizationThreshold{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *UtilizationThreshold) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UtilizationThreshold) Table() flatbuffers.Table {
	return rcv._tab
}

/// Threshold of utilization of the cycle time of Scheduler tick task causes by itself for reporting a warning
func (rcv *UtilizationThreshold) Warning() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 50
}

/// Threshold of utilization of the cycle time of Scheduler tick task causes by itself for reporting a warning
func (rcv *UtilizationThreshold) MutateWarning(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

/// Threshold of utilization of the cycle time of Scheduler tick task causes by itself for reporting an error
func (rcv *UtilizationThreshold) Error() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 75
}

/// Threshold of utilization of the cycle time of Scheduler tick task causes by itself for reporting an error
func (rcv *UtilizationThreshold) MutateError(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func UtilizationThresholdStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func UtilizationThresholdAddWarning(builder *flatbuffers.Builder, warning byte) {
	builder.PrependByteSlot(0, warning, 50)
}
func UtilizationThresholdAddError(builder *flatbuffers.Builder, error byte) {
	builder.PrependByteSlot(1, error, 75)
}
func UtilizationThresholdEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
