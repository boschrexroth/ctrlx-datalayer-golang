// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Duration measurement of execution of the task (related to real-time usage)
type DurationRtT struct {
	Total uint64 `json:"total"`
	Task uint64 `json:"task"`
	Other uint64 `json:"other"`
	Equidistance uint64 `json:"equidistance"`
	Deviation uint64 `json:"deviation"`
	Counter uint64 `json:"counter"`
	Samplerate uint64 `json:"samplerate"`
	Remaining uint64 `json:"remaining"`
}

func (t *DurationRtT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	DurationRtStart(builder)
	DurationRtAddTotal(builder, t.Total)
	DurationRtAddTask(builder, t.Task)
	DurationRtAddOther(builder, t.Other)
	DurationRtAddEquidistance(builder, t.Equidistance)
	DurationRtAddDeviation(builder, t.Deviation)
	DurationRtAddCounter(builder, t.Counter)
	DurationRtAddSamplerate(builder, t.Samplerate)
	DurationRtAddRemaining(builder, t.Remaining)
	return DurationRtEnd(builder)
}

func (rcv *DurationRt) UnPackTo(t *DurationRtT) {
	t.Total = rcv.Total()
	t.Task = rcv.Task()
	t.Other = rcv.Other()
	t.Equidistance = rcv.Equidistance()
	t.Deviation = rcv.Deviation()
	t.Counter = rcv.Counter()
	t.Samplerate = rcv.Samplerate()
	t.Remaining = rcv.Remaining()
}

func (rcv *DurationRt) UnPack() *DurationRtT {
	if rcv == nil { return nil }
	t := &DurationRtT{}
	rcv.UnPackTo(t)
	return t
}

type DurationRt struct {
	_tab flatbuffers.Table
}

func GetRootAsDurationRt(buf []byte, offset flatbuffers.UOffsetT) *DurationRt {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DurationRt{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDurationRt(buf []byte, offset flatbuffers.UOffsetT) *DurationRt {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DurationRt{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DurationRt) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DurationRt) Table() flatbuffers.Table {
	return rcv._tab
}

/// Total duration of the task including interruptions by other tasks in [µs]
func (rcv *DurationRt) Total() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Total duration of the task including interruptions by other tasks in [µs]
func (rcv *DurationRt) MutateTotal(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

/// Task duration of the task without interruptions by other tasks in [µs]
func (rcv *DurationRt) Task() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Task duration of the task without interruptions by other tasks in [µs]
func (rcv *DurationRt) MutateTask(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

/// Duration of interruptions of the task by other tasks in [µs]
func (rcv *DurationRt) Other() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Duration of interruptions of the task by other tasks in [µs]
func (rcv *DurationRt) MutateOther(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

/// Equidistance of the task from one cycle to an other in [µs]
func (rcv *DurationRt) Equidistance() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Equidistance of the task from one cycle to an other in [µs]
func (rcv *DurationRt) MutateEquidistance(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

/// Deviation of the task to the expected start time in [µs]
func (rcv *DurationRt) Deviation() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Deviation of the task to the expected start time in [µs]
func (rcv *DurationRt) MutateDeviation(n uint64) bool {
	return rcv._tab.MutateUint64Slot(12, n)
}

/// Count of execution of the task
func (rcv *DurationRt) Counter() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Count of execution of the task
func (rcv *DurationRt) MutateCounter(n uint64) bool {
	return rcv._tab.MutateUint64Slot(14, n)
}

/// Sample rate resp. cycle time of the task in [µs]
func (rcv *DurationRt) Samplerate() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Sample rate resp. cycle time of the task in [µs]
func (rcv *DurationRt) MutateSamplerate(n uint64) bool {
	return rcv._tab.MutateUint64Slot(16, n)
}

/// Remaining time from the end of the task to the begin of the next cycle of it in [µs]
func (rcv *DurationRt) Remaining() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Remaining time from the end of the task to the begin of the next cycle of it in [µs]
func (rcv *DurationRt) MutateRemaining(n uint64) bool {
	return rcv._tab.MutateUint64Slot(18, n)
}

func DurationRtStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func DurationRtAddTotal(builder *flatbuffers.Builder, total uint64) {
	builder.PrependUint64Slot(0, total, 0)
}
func DurationRtAddTask(builder *flatbuffers.Builder, task uint64) {
	builder.PrependUint64Slot(1, task, 0)
}
func DurationRtAddOther(builder *flatbuffers.Builder, other uint64) {
	builder.PrependUint64Slot(2, other, 0)
}
func DurationRtAddEquidistance(builder *flatbuffers.Builder, equidistance uint64) {
	builder.PrependUint64Slot(3, equidistance, 0)
}
func DurationRtAddDeviation(builder *flatbuffers.Builder, deviation uint64) {
	builder.PrependUint64Slot(4, deviation, 0)
}
func DurationRtAddCounter(builder *flatbuffers.Builder, counter uint64) {
	builder.PrependUint64Slot(5, counter, 0)
}
func DurationRtAddSamplerate(builder *flatbuffers.Builder, samplerate uint64) {
	builder.PrependUint64Slot(6, samplerate, 0)
}
func DurationRtAddRemaining(builder *flatbuffers.Builder, remaining uint64) {
	builder.PrependUint64Slot(7, remaining, 0)
}
func DurationRtEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
