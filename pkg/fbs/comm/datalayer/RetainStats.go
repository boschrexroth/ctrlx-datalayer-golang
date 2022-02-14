// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

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

func (rcv *RetainStats) Total() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RetainStats) MutateTotal(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *RetainStats) Free() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RetainStats) MutateFree(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *RetainStats) Used() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RetainStats) MutateUsed(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *RetainStats) BiggestFree() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RetainStats) MutateBiggestFree(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *RetainStats) SyncCounter() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RetainStats) MutateSyncCounter(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func (rcv *RetainStats) LastUsed() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *RetainStats) MutateLastUsed(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

func (rcv *RetainStats) Info() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func RetainStatsStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
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
func RetainStatsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
