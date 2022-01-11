// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Task struct {
	_tab flatbuffers.Table
}

func GetRootAsTask(buf []byte, offset flatbuffers.UOffsetT) *Task {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Task{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTask(buf []byte, offset flatbuffers.UOffsetT) *Task {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Task{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Task) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Task) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Task) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Task) Priority() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 119
}

func (rcv *Task) MutatePriority(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *Task) Affinity() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Task) MutateAffinity(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *Task) Stacksize() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 131072
}

func (rcv *Task) MutateStacksize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *Task) Event() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Task) Cycletime() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 20000
}

func (rcv *Task) MutateCycletime(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

func TaskStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func TaskAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func TaskAddPriority(builder *flatbuffers.Builder, priority uint32) {
	builder.PrependUint32Slot(1, priority, 119)
}
func TaskAddAffinity(builder *flatbuffers.Builder, affinity uint32) {
	builder.PrependUint32Slot(2, affinity, 0)
}
func TaskAddStacksize(builder *flatbuffers.Builder, stacksize uint32) {
	builder.PrependUint32Slot(3, stacksize, 131072)
}
func TaskAddEvent(builder *flatbuffers.Builder, event flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(event), 0)
}
func TaskAddCycletime(builder *flatbuffers.Builder, cycletime uint32) {
	builder.PrependUint32Slot(5, cycletime, 20000)
}
func TaskEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}