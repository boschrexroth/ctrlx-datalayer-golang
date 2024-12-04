// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LosslessIntervalT struct {
	SamplingInterval uint64 `json:"samplingInterval"`
	Tolerance uint64 `json:"tolerance"`
}

func (t *LosslessIntervalT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	LosslessIntervalStart(builder)
	LosslessIntervalAddSamplingInterval(builder, t.SamplingInterval)
	LosslessIntervalAddTolerance(builder, t.Tolerance)
	return LosslessIntervalEnd(builder)
}

func (rcv *LosslessInterval) UnPackTo(t *LosslessIntervalT) {
	t.SamplingInterval = rcv.SamplingInterval()
	t.Tolerance = rcv.Tolerance()
}

func (rcv *LosslessInterval) UnPack() *LosslessIntervalT {
	if rcv == nil { return nil }
	t := &LosslessIntervalT{}
	rcv.UnPackTo(t)
	return t
}

type LosslessInterval struct {
	_tab flatbuffers.Table
}

func GetRootAsLosslessInterval(buf []byte, offset flatbuffers.UOffsetT) *LosslessInterval {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LosslessInterval{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLosslessInterval(buf []byte, offset flatbuffers.UOffsetT) *LosslessInterval {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LosslessInterval{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LosslessInterval) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LosslessInterval) Table() flatbuffers.Table {
	return rcv._tab
}

/// Time in µs to wait for new data
func (rcv *LosslessInterval) SamplingInterval() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Time in µs to wait for new data
func (rcv *LosslessInterval) MutateSamplingInterval(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

/// Time in µs as tolerance for samplingInterval
func (rcv *LosslessInterval) Tolerance() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Time in µs as tolerance for samplingInterval
func (rcv *LosslessInterval) MutateTolerance(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func LosslessIntervalStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func LosslessIntervalAddSamplingInterval(builder *flatbuffers.Builder, samplingInterval uint64) {
	builder.PrependUint64Slot(0, samplingInterval, 0)
}
func LosslessIntervalAddTolerance(builder *flatbuffers.Builder, tolerance uint64) {
	builder.PrependUint64Slot(1, tolerance, 0)
}
func LosslessIntervalEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
