// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Information about the current cycle of the task (related to real-time usage)
type DataRtT struct {
	StartTime uint64 `json:"startTime"`
	Counter uint64 `json:"counter"`
	CalculatedStartTime uint64 `json:"calculatedStartTime"`
}

func (t *DataRtT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	DataRtStart(builder)
	DataRtAddStartTime(builder, t.StartTime)
	DataRtAddCounter(builder, t.Counter)
	DataRtAddCalculatedStartTime(builder, t.CalculatedStartTime)
	return DataRtEnd(builder)
}

func (rcv *DataRt) UnPackTo(t *DataRtT) {
	t.StartTime = rcv.StartTime()
	t.Counter = rcv.Counter()
	t.CalculatedStartTime = rcv.CalculatedStartTime()
}

func (rcv *DataRt) UnPack() *DataRtT {
	if rcv == nil { return nil }
	t := &DataRtT{}
	rcv.UnPackTo(t)
	return t
}

type DataRt struct {
	_tab flatbuffers.Table
}

func GetRootAsDataRt(buf []byte, offset flatbuffers.UOffsetT) *DataRt {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DataRt{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDataRt(buf []byte, offset flatbuffers.UOffsetT) *DataRt {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DataRt{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DataRt) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DataRt) Table() flatbuffers.Table {
	return rcv._tab
}

/// Start time of the task in [µs] 
func (rcv *DataRt) StartTime() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Start time of the task in [µs] 
func (rcv *DataRt) MutateStartTime(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

/// Count of execution of the task
func (rcv *DataRt) Counter() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Count of execution of the task
func (rcv *DataRt) MutateCounter(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

/// Calculated start time of the task in [µs]
func (rcv *DataRt) CalculatedStartTime() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Calculated start time of the task in [µs]
func (rcv *DataRt) MutateCalculatedStartTime(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

func DataRtStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func DataRtAddStartTime(builder *flatbuffers.Builder, startTime uint64) {
	builder.PrependUint64Slot(0, startTime, 0)
}
func DataRtAddCounter(builder *flatbuffers.Builder, counter uint64) {
	builder.PrependUint64Slot(1, counter, 0)
}
func DataRtAddCalculatedStartTime(builder *flatbuffers.Builder, calculatedStartTime uint64) {
	builder.PrependUint64Slot(2, calculatedStartTime, 0)
}
func DataRtEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
