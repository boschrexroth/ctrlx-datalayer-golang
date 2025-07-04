// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Frame statistic counters
type FrameStatisticCountersResponseT struct {
	NumTxFrames uint32 `json:"numTxFrames"`
	NumRxFrames uint32 `json:"numRxFrames"`
	NumCyclicFrames uint32 `json:"numCyclicFrames"`
	NumCyclicDatagrams uint32 `json:"numCyclicDatagrams"`
	NumAcyclicFrames uint32 `json:"numAcyclicFrames"`
	NumAcyclicDatagrams uint32 `json:"numAcyclicDatagrams"`
	NumLostFrames uint32 `json:"numLostFrames"`
	NumLostCyclicFrames uint32 `json:"numLostCyclicFrames"`
	NumLostAcyclicFrames uint32 `json:"numLostAcyclicFrames"`
	NumCyclicFramesPerCycle *MinActMaxValuesT `json:"numCyclicFramesPerCycle"`
	NumAcyclicFramesPerCycle *MinActMaxValuesT `json:"numAcyclicFramesPerCycle"`
	NumEoeFramesPerCycle *MinActMaxValuesT `json:"numEoeFramesPerCycle"`
}

func (t *FrameStatisticCountersResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	FrameStatisticCountersResponseStart(builder)
	FrameStatisticCountersResponseAddNumTxFrames(builder, t.NumTxFrames)
	FrameStatisticCountersResponseAddNumRxFrames(builder, t.NumRxFrames)
	FrameStatisticCountersResponseAddNumCyclicFrames(builder, t.NumCyclicFrames)
	FrameStatisticCountersResponseAddNumCyclicDatagrams(builder, t.NumCyclicDatagrams)
	FrameStatisticCountersResponseAddNumAcyclicFrames(builder, t.NumAcyclicFrames)
	FrameStatisticCountersResponseAddNumAcyclicDatagrams(builder, t.NumAcyclicDatagrams)
	FrameStatisticCountersResponseAddNumLostFrames(builder, t.NumLostFrames)
	FrameStatisticCountersResponseAddNumLostCyclicFrames(builder, t.NumLostCyclicFrames)
	FrameStatisticCountersResponseAddNumLostAcyclicFrames(builder, t.NumLostAcyclicFrames)
	numCyclicFramesPerCycleOffset := t.NumCyclicFramesPerCycle.Pack(builder)
	FrameStatisticCountersResponseAddNumCyclicFramesPerCycle(builder, numCyclicFramesPerCycleOffset)
	numAcyclicFramesPerCycleOffset := t.NumAcyclicFramesPerCycle.Pack(builder)
	FrameStatisticCountersResponseAddNumAcyclicFramesPerCycle(builder, numAcyclicFramesPerCycleOffset)
	numEoeFramesPerCycleOffset := t.NumEoeFramesPerCycle.Pack(builder)
	FrameStatisticCountersResponseAddNumEoeFramesPerCycle(builder, numEoeFramesPerCycleOffset)
	return FrameStatisticCountersResponseEnd(builder)
}

func (rcv *FrameStatisticCountersResponse) UnPackTo(t *FrameStatisticCountersResponseT) {
	t.NumTxFrames = rcv.NumTxFrames()
	t.NumRxFrames = rcv.NumRxFrames()
	t.NumCyclicFrames = rcv.NumCyclicFrames()
	t.NumCyclicDatagrams = rcv.NumCyclicDatagrams()
	t.NumAcyclicFrames = rcv.NumAcyclicFrames()
	t.NumAcyclicDatagrams = rcv.NumAcyclicDatagrams()
	t.NumLostFrames = rcv.NumLostFrames()
	t.NumLostCyclicFrames = rcv.NumLostCyclicFrames()
	t.NumLostAcyclicFrames = rcv.NumLostAcyclicFrames()
	t.NumCyclicFramesPerCycle = rcv.NumCyclicFramesPerCycle(nil).UnPack()
	t.NumAcyclicFramesPerCycle = rcv.NumAcyclicFramesPerCycle(nil).UnPack()
	t.NumEoeFramesPerCycle = rcv.NumEoeFramesPerCycle(nil).UnPack()
}

func (rcv *FrameStatisticCountersResponse) UnPack() *FrameStatisticCountersResponseT {
	if rcv == nil { return nil }
	t := &FrameStatisticCountersResponseT{}
	rcv.UnPackTo(t)
	return t
}

type FrameStatisticCountersResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsFrameStatisticCountersResponse(buf []byte, offset flatbuffers.UOffsetT) *FrameStatisticCountersResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FrameStatisticCountersResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFrameStatisticCountersResponse(buf []byte, offset flatbuffers.UOffsetT) *FrameStatisticCountersResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FrameStatisticCountersResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FrameStatisticCountersResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FrameStatisticCountersResponse) Table() flatbuffers.Table {
	return rcv._tab
}

///Number of total transmitted frames
func (rcv *FrameStatisticCountersResponse) NumTxFrames() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Number of total transmitted frames
func (rcv *FrameStatisticCountersResponse) MutateNumTxFrames(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

///Number of total received frames
func (rcv *FrameStatisticCountersResponse) NumRxFrames() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Number of total received frames
func (rcv *FrameStatisticCountersResponse) MutateNumRxFrames(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

///Number of cyclic frames
func (rcv *FrameStatisticCountersResponse) NumCyclicFrames() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Number of cyclic frames
func (rcv *FrameStatisticCountersResponse) MutateNumCyclicFrames(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

///Number of cyclic datagrams
func (rcv *FrameStatisticCountersResponse) NumCyclicDatagrams() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Number of cyclic datagrams
func (rcv *FrameStatisticCountersResponse) MutateNumCyclicDatagrams(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

///Number of acyclic frames
func (rcv *FrameStatisticCountersResponse) NumAcyclicFrames() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Number of acyclic frames
func (rcv *FrameStatisticCountersResponse) MutateNumAcyclicFrames(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

///Number of acyclic datagrams
func (rcv *FrameStatisticCountersResponse) NumAcyclicDatagrams() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Number of acyclic datagrams
func (rcv *FrameStatisticCountersResponse) MutateNumAcyclicDatagrams(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

///Number of total lost frames
func (rcv *FrameStatisticCountersResponse) NumLostFrames() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Number of total lost frames
func (rcv *FrameStatisticCountersResponse) MutateNumLostFrames(n uint32) bool {
	return rcv._tab.MutateUint32Slot(16, n)
}

///Number of lost cyclic frames
func (rcv *FrameStatisticCountersResponse) NumLostCyclicFrames() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Number of lost cyclic frames
func (rcv *FrameStatisticCountersResponse) MutateNumLostCyclicFrames(n uint32) bool {
	return rcv._tab.MutateUint32Slot(18, n)
}

///Number of lost acyclic frames
func (rcv *FrameStatisticCountersResponse) NumLostAcyclicFrames() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Number of lost acyclic frames
func (rcv *FrameStatisticCountersResponse) MutateNumLostAcyclicFrames(n uint32) bool {
	return rcv._tab.MutateUint32Slot(20, n)
}

///Number of cyclic frames sent per cylce
func (rcv *FrameStatisticCountersResponse) NumCyclicFramesPerCycle(obj *MinActMaxValues) *MinActMaxValues {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(MinActMaxValues)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

///Number of cyclic frames sent per cylce
///Number of acyclic frames sent per cylce
func (rcv *FrameStatisticCountersResponse) NumAcyclicFramesPerCycle(obj *MinActMaxValues) *MinActMaxValues {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(MinActMaxValues)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

///Number of acyclic frames sent per cylce
///Number of EoE frames processed (switched) per cyclie
func (rcv *FrameStatisticCountersResponse) NumEoeFramesPerCycle(obj *MinActMaxValues) *MinActMaxValues {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(MinActMaxValues)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

///Number of EoE frames processed (switched) per cyclie
func FrameStatisticCountersResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(12)
}
func FrameStatisticCountersResponseAddNumTxFrames(builder *flatbuffers.Builder, numTxFrames uint32) {
	builder.PrependUint32Slot(0, numTxFrames, 0)
}
func FrameStatisticCountersResponseAddNumRxFrames(builder *flatbuffers.Builder, numRxFrames uint32) {
	builder.PrependUint32Slot(1, numRxFrames, 0)
}
func FrameStatisticCountersResponseAddNumCyclicFrames(builder *flatbuffers.Builder, numCyclicFrames uint32) {
	builder.PrependUint32Slot(2, numCyclicFrames, 0)
}
func FrameStatisticCountersResponseAddNumCyclicDatagrams(builder *flatbuffers.Builder, numCyclicDatagrams uint32) {
	builder.PrependUint32Slot(3, numCyclicDatagrams, 0)
}
func FrameStatisticCountersResponseAddNumAcyclicFrames(builder *flatbuffers.Builder, numAcyclicFrames uint32) {
	builder.PrependUint32Slot(4, numAcyclicFrames, 0)
}
func FrameStatisticCountersResponseAddNumAcyclicDatagrams(builder *flatbuffers.Builder, numAcyclicDatagrams uint32) {
	builder.PrependUint32Slot(5, numAcyclicDatagrams, 0)
}
func FrameStatisticCountersResponseAddNumLostFrames(builder *flatbuffers.Builder, numLostFrames uint32) {
	builder.PrependUint32Slot(6, numLostFrames, 0)
}
func FrameStatisticCountersResponseAddNumLostCyclicFrames(builder *flatbuffers.Builder, numLostCyclicFrames uint32) {
	builder.PrependUint32Slot(7, numLostCyclicFrames, 0)
}
func FrameStatisticCountersResponseAddNumLostAcyclicFrames(builder *flatbuffers.Builder, numLostAcyclicFrames uint32) {
	builder.PrependUint32Slot(8, numLostAcyclicFrames, 0)
}
func FrameStatisticCountersResponseAddNumCyclicFramesPerCycle(builder *flatbuffers.Builder, numCyclicFramesPerCycle flatbuffers.UOffsetT) {
	builder.PrependStructSlot(9, flatbuffers.UOffsetT(numCyclicFramesPerCycle), 0)
}
func FrameStatisticCountersResponseAddNumAcyclicFramesPerCycle(builder *flatbuffers.Builder, numAcyclicFramesPerCycle flatbuffers.UOffsetT) {
	builder.PrependStructSlot(10, flatbuffers.UOffsetT(numAcyclicFramesPerCycle), 0)
}
func FrameStatisticCountersResponseAddNumEoeFramesPerCycle(builder *flatbuffers.Builder, numEoeFramesPerCycle flatbuffers.UOffsetT) {
	builder.PrependStructSlot(11, flatbuffers.UOffsetT(numEoeFramesPerCycle), 0)
}
func FrameStatisticCountersResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
