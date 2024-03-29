// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// currently interpolated values
type IpoAddDataT struct {
	DistFromStart float64 `json:"distFromStart"`
	DistToTarget float64 `json:"distToTarget"`
	TimeFromStart float64 `json:"timeFromStart"`
	TimeToTarget float64 `json:"timeToTarget"`
	DistFromStartUnit string `json:"distFromStartUnit"`
	DistToTargetUnit string `json:"distToTargetUnit"`
	TimeFromStartUnit string `json:"timeFromStartUnit"`
	TimeToTargetUnit string `json:"timeToTargetUnit"`
}

func (t *IpoAddDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	distFromStartUnitOffset := flatbuffers.UOffsetT(0)
	if t.DistFromStartUnit != "" {
		distFromStartUnitOffset = builder.CreateString(t.DistFromStartUnit)
	}
	distToTargetUnitOffset := flatbuffers.UOffsetT(0)
	if t.DistToTargetUnit != "" {
		distToTargetUnitOffset = builder.CreateString(t.DistToTargetUnit)
	}
	timeFromStartUnitOffset := flatbuffers.UOffsetT(0)
	if t.TimeFromStartUnit != "" {
		timeFromStartUnitOffset = builder.CreateString(t.TimeFromStartUnit)
	}
	timeToTargetUnitOffset := flatbuffers.UOffsetT(0)
	if t.TimeToTargetUnit != "" {
		timeToTargetUnitOffset = builder.CreateString(t.TimeToTargetUnit)
	}
	IpoAddDataStart(builder)
	IpoAddDataAddDistFromStart(builder, t.DistFromStart)
	IpoAddDataAddDistToTarget(builder, t.DistToTarget)
	IpoAddDataAddTimeFromStart(builder, t.TimeFromStart)
	IpoAddDataAddTimeToTarget(builder, t.TimeToTarget)
	IpoAddDataAddDistFromStartUnit(builder, distFromStartUnitOffset)
	IpoAddDataAddDistToTargetUnit(builder, distToTargetUnitOffset)
	IpoAddDataAddTimeFromStartUnit(builder, timeFromStartUnitOffset)
	IpoAddDataAddTimeToTargetUnit(builder, timeToTargetUnitOffset)
	return IpoAddDataEnd(builder)
}

func (rcv *IpoAddData) UnPackTo(t *IpoAddDataT) {
	t.DistFromStart = rcv.DistFromStart()
	t.DistToTarget = rcv.DistToTarget()
	t.TimeFromStart = rcv.TimeFromStart()
	t.TimeToTarget = rcv.TimeToTarget()
	t.DistFromStartUnit = string(rcv.DistFromStartUnit())
	t.DistToTargetUnit = string(rcv.DistToTargetUnit())
	t.TimeFromStartUnit = string(rcv.TimeFromStartUnit())
	t.TimeToTargetUnit = string(rcv.TimeToTargetUnit())
}

func (rcv *IpoAddData) UnPack() *IpoAddDataT {
	if rcv == nil { return nil }
	t := &IpoAddDataT{}
	rcv.UnPackTo(t)
	return t
}

type IpoAddData struct {
	_tab flatbuffers.Table
}

func GetRootAsIpoAddData(buf []byte, offset flatbuffers.UOffsetT) *IpoAddData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IpoAddData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsIpoAddData(buf []byte, offset flatbuffers.UOffsetT) *IpoAddData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &IpoAddData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *IpoAddData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IpoAddData) Table() flatbuffers.Table {
	return rcv._tab
}

/// distance from start position
func (rcv *IpoAddData) DistFromStart() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// distance from start position
func (rcv *IpoAddData) MutateDistFromStart(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// distance to target position
func (rcv *IpoAddData) DistToTarget() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// distance to target position
func (rcv *IpoAddData) MutateDistToTarget(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// time from start position
func (rcv *IpoAddData) TimeFromStart() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// time from start position
func (rcv *IpoAddData) MutateTimeFromStart(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

/// time to target position
func (rcv *IpoAddData) TimeToTarget() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// time to target position
func (rcv *IpoAddData) MutateTimeToTarget(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

/// unit of the distance from start position
func (rcv *IpoAddData) DistFromStartUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the distance from start position
/// unit of the distance to target position
func (rcv *IpoAddData) DistToTargetUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the distance to target position
/// unit of the time from start position
func (rcv *IpoAddData) TimeFromStartUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the time from start position
/// unit of the time to target position
func (rcv *IpoAddData) TimeToTargetUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the time to target position
func IpoAddDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func IpoAddDataAddDistFromStart(builder *flatbuffers.Builder, distFromStart float64) {
	builder.PrependFloat64Slot(0, distFromStart, 0.0)
}
func IpoAddDataAddDistToTarget(builder *flatbuffers.Builder, distToTarget float64) {
	builder.PrependFloat64Slot(1, distToTarget, 0.0)
}
func IpoAddDataAddTimeFromStart(builder *flatbuffers.Builder, timeFromStart float64) {
	builder.PrependFloat64Slot(2, timeFromStart, 0.0)
}
func IpoAddDataAddTimeToTarget(builder *flatbuffers.Builder, timeToTarget float64) {
	builder.PrependFloat64Slot(3, timeToTarget, 0.0)
}
func IpoAddDataAddDistFromStartUnit(builder *flatbuffers.Builder, distFromStartUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(distFromStartUnit), 0)
}
func IpoAddDataAddDistToTargetUnit(builder *flatbuffers.Builder, distToTargetUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(distToTargetUnit), 0)
}
func IpoAddDataAddTimeFromStartUnit(builder *flatbuffers.Builder, timeFromStartUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(timeFromStartUnit), 0)
}
func IpoAddDataAddTimeToTargetUnit(builder *flatbuffers.Builder, timeToTargetUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(timeToTargetUnit), 0)
}
func IpoAddDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
