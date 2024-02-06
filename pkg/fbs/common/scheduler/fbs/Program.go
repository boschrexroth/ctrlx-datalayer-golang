// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"

	common__scheduler__watchdog__fbs "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/common/scheduler/watchdog/fbs"
)

type ProgramT struct {
	Task *TaskT `json:"task"`
	Callables []*CallableT `json:"callables"`
	Watchdog *common__scheduler__watchdog__fbs.WatchdogT `json:"watchdog"`
}

func (t *ProgramT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	taskOffset := t.Task.Pack(builder)
	callablesOffset := flatbuffers.UOffsetT(0)
	if t.Callables != nil {
		callablesLength := len(t.Callables)
		callablesOffsets := make([]flatbuffers.UOffsetT, callablesLength)
		for j := 0; j < callablesLength; j++ {
			callablesOffsets[j] = t.Callables[j].Pack(builder)
		}
		ProgramStartCallablesVector(builder, callablesLength)
		for j := callablesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(callablesOffsets[j])
		}
		callablesOffset = builder.EndVector(callablesLength)
	}
	watchdogOffset := t.Watchdog.Pack(builder)
	ProgramStart(builder)
	ProgramAddTask(builder, taskOffset)
	ProgramAddCallables(builder, callablesOffset)
	ProgramAddWatchdog(builder, watchdogOffset)
	return ProgramEnd(builder)
}

func (rcv *Program) UnPackTo(t *ProgramT) {
	t.Task = rcv.Task(nil).UnPack()
	callablesLength := rcv.CallablesLength()
	t.Callables = make([]*CallableT, callablesLength)
	for j := 0; j < callablesLength; j++ {
		x := Callable{}
		rcv.Callables(&x, j)
		t.Callables[j] = x.UnPack()
	}
	t.Watchdog = rcv.Watchdog(nil).UnPack()
}

func (rcv *Program) UnPack() *ProgramT {
	if rcv == nil { return nil }
	t := &ProgramT{}
	rcv.UnPackTo(t)
	return t
}

type Program struct {
	_tab flatbuffers.Table
}

func GetRootAsProgram(buf []byte, offset flatbuffers.UOffsetT) *Program {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Program{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsProgram(buf []byte, offset flatbuffers.UOffsetT) *Program {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Program{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Program) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Program) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Program) Task(obj *Task) *Task {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Task)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Program) Callables(obj *Callable, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Program) CallablesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Program) Watchdog(obj *common__scheduler__watchdog__fbs.Watchdog) *common__scheduler__watchdog__fbs.Watchdog {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(common__scheduler__watchdog__fbs.Watchdog)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ProgramStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ProgramAddTask(builder *flatbuffers.Builder, task flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(task), 0)
}
func ProgramAddCallables(builder *flatbuffers.Builder, callables flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(callables), 0)
}
func ProgramStartCallablesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ProgramAddWatchdog(builder *flatbuffers.Builder, watchdog flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(watchdog), 0)
}
func ProgramEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
