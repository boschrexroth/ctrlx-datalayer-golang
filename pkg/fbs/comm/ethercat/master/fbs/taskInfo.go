// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"

	common__scheduler__fbs "github.com/boschrexroth/ctrlx-datalayer-golang/pkg/fbs/common/scheduler/fbs"
)

type taskInfoT struct {
	Task *common__scheduler__fbs.TaskT
	Counter uint64
}

func (t *taskInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	taskOffset := t.Task.Pack(builder)
	taskInfoStart(builder)
	taskInfoAddTask(builder, taskOffset)
	taskInfoAddCounter(builder, t.Counter)
	return taskInfoEnd(builder)
}

func (rcv *taskInfo) UnPackTo(t *taskInfoT) {
	t.Task = rcv.Task(nil).UnPack()
	t.Counter = rcv.Counter()
}

func (rcv *taskInfo) UnPack() *taskInfoT {
	if rcv == nil { return nil }
	t := &taskInfoT{}
	rcv.UnPackTo(t)
	return t
}

type taskInfo struct {
	_tab flatbuffers.Table
}

func GetRootAstaskInfo(buf []byte, offset flatbuffers.UOffsetT) *taskInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &taskInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAstaskInfo(buf []byte, offset flatbuffers.UOffsetT) *taskInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &taskInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *taskInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *taskInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *taskInfo) Task(obj *common__scheduler__fbs.Task) *common__scheduler__fbs.Task {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(common__scheduler__fbs.Task)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *taskInfo) Counter() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *taskInfo) MutateCounter(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func taskInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func taskInfoAddTask(builder *flatbuffers.Builder, task flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(task), 0)
}
func taskInfoAddCounter(builder *flatbuffers.Builder, counter uint64) {
	builder.PrependUint64Slot(1, counter, 0)
}
func taskInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
