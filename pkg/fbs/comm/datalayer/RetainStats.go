// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RetainStatsT struct {
	Total uint32 `json:"total"`
	Free uint32 `json:"free"`
	Used uint32 `json:"used"`
	BiggestFree uint32 `json:"biggestFree"`
	SyncCounter uint32 `json:"syncCounter"`
	LastUsed uint32 `json:"lastUsed"`
	Info string `json:"info"`
	Diagnosis *RetainDiagnosisT `json:"diagnosis"`
}

func (t *RetainStatsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	infoOffset := flatbuffers.UOffsetT(0)
	if t.Info != "" {
		infoOffset = builder.CreateString(t.Info)
	}
	diagnosisOffset := t.Diagnosis.Pack(builder)
	RetainStatsStart(builder)
	RetainStatsAddTotal(builder, t.Total)
	RetainStatsAddFree(builder, t.Free)
	RetainStatsAddUsed(builder, t.Used)
	RetainStatsAddBiggestFree(builder, t.BiggestFree)
	RetainStatsAddSyncCounter(builder, t.SyncCounter)
	RetainStatsAddLastUsed(builder, t.LastUsed)
	RetainStatsAddInfo(builder, infoOffset)
	RetainStatsAddDiagnosis(builder, diagnosisOffset)
	return RetainStatsEnd(builder)
}

func (rcv *RetainStats) UnPackTo(t *RetainStatsT) {
	t.Total = rcv.Total()
	t.Free = rcv.Free()
	t.Used = rcv.Used()
	t.BiggestFree = rcv.BiggestFree()
	t.SyncCounter = rcv.SyncCounter()
	t.LastUsed = rcv.LastUsed()
	t.Info = string(rcv.Info())
	t.Diagnosis = rcv.Diagnosis(nil).UnPack()
}

func (rcv *RetainStats) UnPack() *RetainStatsT {
	if rcv == nil { return nil }
	t := &RetainStatsT{}
	rcv.UnPackTo(t)
	return t
}

type RetainStats struct {
	_tab flatbuffers.Table
}

func GetRootAsRetainStats(buf []byte, offset flatbuffers.UOffsetT) *RetainStats {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RetainStats{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRetainStats(buf []byte, offset flatbuffers.UOffsetT) *RetainStats {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RetainStats{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RetainStats) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RetainStats) Table() flatbuffers.Table {
	return rcv._tab
}

/// total size of memory in bytes
func (rcv *RetainStats) Total() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// total size of memory in bytes
func (rcv *RetainStats) MutateTotal(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

/// free size of memory in bytes
func (rcv *RetainStats) Free() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// free size of memory in bytes
func (rcv *RetainStats) MutateFree(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

/// used size of memory in bytes
func (rcv *RetainStats) Used() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// used size of memory in bytes
func (rcv *RetainStats) MutateUsed(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

/// biggest free chunk of memory in bytes
func (rcv *RetainStats) BiggestFree() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// biggest free chunk of memory in bytes
func (rcv *RetainStats) MutateBiggestFree(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

/// now often was the nvram synced
func (rcv *RetainStats) SyncCounter() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// now often was the nvram synced
func (rcv *RetainStats) MutateSyncCounter(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

/// last used offset
func (rcv *RetainStats) LastUsed() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// last used offset
func (rcv *RetainStats) MutateLastUsed(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

/// debug information of shared memory
func (rcv *RetainStats) Info() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// debug information of shared memory
/// diagnosis of nvram
func (rcv *RetainStats) Diagnosis(obj *RetainDiagnosis) *RetainDiagnosis {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(RetainDiagnosis)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// diagnosis of nvram
func RetainStatsStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func RetainStatsAddTotal(builder *flatbuffers.Builder, total uint32) {
	builder.PrependUint32Slot(0, total, 0)
}
func RetainStatsAddFree(builder *flatbuffers.Builder, free uint32) {
	builder.PrependUint32Slot(1, free, 0)
}
func RetainStatsAddUsed(builder *flatbuffers.Builder, used uint32) {
	builder.PrependUint32Slot(2, used, 0)
}
func RetainStatsAddBiggestFree(builder *flatbuffers.Builder, biggestFree uint32) {
	builder.PrependUint32Slot(3, biggestFree, 0)
}
func RetainStatsAddSyncCounter(builder *flatbuffers.Builder, syncCounter uint32) {
	builder.PrependUint32Slot(4, syncCounter, 0)
}
func RetainStatsAddLastUsed(builder *flatbuffers.Builder, lastUsed uint32) {
	builder.PrependUint32Slot(5, lastUsed, 0)
}
func RetainStatsAddInfo(builder *flatbuffers.Builder, info flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(info), 0)
}
func RetainStatsAddDiagnosis(builder *flatbuffers.Builder, diagnosis flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(diagnosis), 0)
}
func RetainStatsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
