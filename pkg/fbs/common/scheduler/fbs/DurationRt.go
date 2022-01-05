// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

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

func (rcv *DurationRt) Total() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DurationRt) MutateTotal(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *DurationRt) Task() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DurationRt) MutateTask(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *DurationRt) Other() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DurationRt) MutateOther(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

func (rcv *DurationRt) Equidistance() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DurationRt) MutateEquidistance(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func (rcv *DurationRt) Deviation() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DurationRt) MutateDeviation(n uint64) bool {
	return rcv._tab.MutateUint64Slot(12, n)
}

func (rcv *DurationRt) Counter() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DurationRt) MutateCounter(n uint64) bool {
	return rcv._tab.MutateUint64Slot(14, n)
}

func (rcv *DurationRt) Samplerate() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DurationRt) MutateSamplerate(n uint64) bool {
	return rcv._tab.MutateUint64Slot(16, n)
}

func DurationRtStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
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
func DurationRtEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
