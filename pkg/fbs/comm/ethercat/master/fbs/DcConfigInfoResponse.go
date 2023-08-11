// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DcConfigInfoResponseT struct {
	SyncMode SyncMode `json:"syncMode"`
	CycleTime uint32 `json:"cycleTime"`
	SyncShiftTime uint32 `json:"syncShiftTime"`
	SyncWindowMonitoring bool `json:"syncWindowMonitoring"`
	DeviationLimit uint32 `json:"deviationLimit"`
	ContinuousDelayCompensation bool `json:"continuousDelayCompensation"`
}

func (t *DcConfigInfoResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	DcConfigInfoResponseStart(builder)
	DcConfigInfoResponseAddSyncMode(builder, t.SyncMode)
	DcConfigInfoResponseAddCycleTime(builder, t.CycleTime)
	DcConfigInfoResponseAddSyncShiftTime(builder, t.SyncShiftTime)
	DcConfigInfoResponseAddSyncWindowMonitoring(builder, t.SyncWindowMonitoring)
	DcConfigInfoResponseAddDeviationLimit(builder, t.DeviationLimit)
	DcConfigInfoResponseAddContinuousDelayCompensation(builder, t.ContinuousDelayCompensation)
	return DcConfigInfoResponseEnd(builder)
}

func (rcv *DcConfigInfoResponse) UnPackTo(t *DcConfigInfoResponseT) {
	t.SyncMode = rcv.SyncMode()
	t.CycleTime = rcv.CycleTime()
	t.SyncShiftTime = rcv.SyncShiftTime()
	t.SyncWindowMonitoring = rcv.SyncWindowMonitoring()
	t.DeviationLimit = rcv.DeviationLimit()
	t.ContinuousDelayCompensation = rcv.ContinuousDelayCompensation()
}

func (rcv *DcConfigInfoResponse) UnPack() *DcConfigInfoResponseT {
	if rcv == nil { return nil }
	t := &DcConfigInfoResponseT{}
	rcv.UnPackTo(t)
	return t
}

type DcConfigInfoResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsDcConfigInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *DcConfigInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DcConfigInfoResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDcConfigInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *DcConfigInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DcConfigInfoResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DcConfigInfoResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DcConfigInfoResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DcConfigInfoResponse) SyncMode() SyncMode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return SyncMode(rcv._tab.GetUint32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *DcConfigInfoResponse) MutateSyncMode(n SyncMode) bool {
	return rcv._tab.MutateUint32Slot(4, uint32(n))
}

func (rcv *DcConfigInfoResponse) CycleTime() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DcConfigInfoResponse) MutateCycleTime(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *DcConfigInfoResponse) SyncShiftTime() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DcConfigInfoResponse) MutateSyncShiftTime(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *DcConfigInfoResponse) SyncWindowMonitoring() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *DcConfigInfoResponse) MutateSyncWindowMonitoring(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func (rcv *DcConfigInfoResponse) DeviationLimit() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DcConfigInfoResponse) MutateDeviationLimit(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func (rcv *DcConfigInfoResponse) ContinuousDelayCompensation() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *DcConfigInfoResponse) MutateContinuousDelayCompensation(n bool) bool {
	return rcv._tab.MutateBoolSlot(14, n)
}

func DcConfigInfoResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func DcConfigInfoResponseAddSyncMode(builder *flatbuffers.Builder, syncMode SyncMode) {
	builder.PrependUint32Slot(0, uint32(syncMode), 0)
}
func DcConfigInfoResponseAddCycleTime(builder *flatbuffers.Builder, cycleTime uint32) {
	builder.PrependUint32Slot(1, cycleTime, 0)
}
func DcConfigInfoResponseAddSyncShiftTime(builder *flatbuffers.Builder, syncShiftTime uint32) {
	builder.PrependUint32Slot(2, syncShiftTime, 0)
}
func DcConfigInfoResponseAddSyncWindowMonitoring(builder *flatbuffers.Builder, syncWindowMonitoring bool) {
	builder.PrependBoolSlot(3, syncWindowMonitoring, false)
}
func DcConfigInfoResponseAddDeviationLimit(builder *flatbuffers.Builder, deviationLimit uint32) {
	builder.PrependUint32Slot(4, deviationLimit, 0)
}
func DcConfigInfoResponseAddContinuousDelayCompensation(builder *flatbuffers.Builder, continuousDelayCompensation bool) {
	builder.PrependBoolSlot(5, continuousDelayCompensation, false)
}
func DcConfigInfoResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
