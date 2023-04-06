// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Properties of a task
type TaskT struct {
	Name string
	Priority uint32
	Affinity uint32
	Stacksize uint32
	Event string
	Cycletime uint32
}

func (t *TaskT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	eventOffset := builder.CreateString(t.Event)
	TaskStart(builder)
	TaskAddName(builder, nameOffset)
	TaskAddPriority(builder, t.Priority)
	TaskAddAffinity(builder, t.Affinity)
	TaskAddStacksize(builder, t.Stacksize)
	TaskAddEvent(builder, eventOffset)
	TaskAddCycletime(builder, t.Cycletime)
	return TaskEnd(builder)
}

func (rcv *Task) UnPackTo(t *TaskT) {
	t.Name = string(rcv.Name())
	t.Priority = rcv.Priority()
	t.Affinity = rcv.Affinity()
	t.Stacksize = rcv.Stacksize()
	t.Event = string(rcv.Event())
	t.Cycletime = rcv.Cycletime()
}

func (rcv *Task) UnPack() *TaskT {
	if rcv == nil { return nil }
	t := &TaskT{}
	rcv.UnPackTo(t)
	return t
}

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

/// Name of the task [unique]
///   Allowed characters:
///     Any alphanumeric character, beginning with a letter and a maximum length of 15 characters [a-zA-Z][a-zA-Z0-9]{1,15}
func (rcv *Task) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Name of the task [unique]
///   Allowed characters:
///     Any alphanumeric character, beginning with a letter and a maximum length of 15 characters [a-zA-Z][a-zA-Z0-9]{1,15}
/// Priority of the task
///   Priority ranges:
///       0 ..  10  : reserved for the system
///      11 ..  99  : available for real-time user tasks
///     100 .. 139  : available for non real-time user tasks
///   Priority agreements:
///      10         : reserved for Scheduler tick task 'schedMain'
///      11         : highest prior real-time task, use of the policy FIFO policy
///      23         : high prior real-time task, use of the policy FIFO policy
///      29         : mid prior real-time task, use of the policy FIFO policy
///      37         : low prior real-time task, use of the policy FIFO policy
///      99         : lowest prior real-time task, use of the policy round-robin policy
///     100         : highest prior non real-time task, use of the nice value of '-20'
///     120         : common used non real-time task, use of the nice value of '0'
///     139         : lowest prior non real-time task, use of the nice value of '19'
func (rcv *Task) Priority() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 120
}

/// Priority of the task
///   Priority ranges:
///       0 ..  10  : reserved for the system
///      11 ..  99  : available for real-time user tasks
///     100 .. 139  : available for non real-time user tasks
///   Priority agreements:
///      10         : reserved for Scheduler tick task 'schedMain'
///      11         : highest prior real-time task, use of the policy FIFO policy
///      23         : high prior real-time task, use of the policy FIFO policy
///      29         : mid prior real-time task, use of the policy FIFO policy
///      37         : low prior real-time task, use of the policy FIFO policy
///      99         : lowest prior real-time task, use of the policy round-robin policy
///     100         : highest prior non real-time task, use of the nice value of '-20'
///     120         : common used non real-time task, use of the nice value of '0'
///     139         : lowest prior non real-time task, use of the nice value of '19'
func (rcv *Task) MutatePriority(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

/// CPU core affinity of the task, defining on which CPU core it is executed, available cores see 'scheduler/admin/info/cpu-cores'
func (rcv *Task) Affinity() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// CPU core affinity of the task, defining on which CPU core it is executed, available cores see 'scheduler/admin/info/cpu-cores'
func (rcv *Task) MutateAffinity(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

/// Stack size of the task in [byte]
func (rcv *Task) Stacksize() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 131072
}

/// Stack size of the task in [byte]
func (rcv *Task) MutateStacksize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

/// Execution event of the task ["cyclic"]
func (rcv *Task) Event() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Execution event of the task ["cyclic"]
/// Cycle time of the task in [µs]
func (rcv *Task) Cycletime() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 20000
}

/// Cycle time of the task in [µs]
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
	builder.PrependUint32Slot(1, priority, 120)
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
