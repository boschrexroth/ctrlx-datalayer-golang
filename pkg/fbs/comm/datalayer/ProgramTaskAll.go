// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ProgramTaskAllT struct {
	Tasks []*ProgramTaskT
}

func (t *ProgramTaskAllT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	TasksOffset := flatbuffers.UOffsetT(0)
	if t.Tasks != nil {
		TasksLength := len(t.Tasks)
		TasksOffsets := make([]flatbuffers.UOffsetT, TasksLength)
		for j := 0; j < TasksLength; j++ {
			TasksOffsets[j] = t.Tasks[j].Pack(builder)
		}
		ProgramTaskAllStartTasksVector(builder, TasksLength)
		for j := TasksLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(TasksOffsets[j])
		}
		TasksOffset = builder.EndVector(TasksLength)
	}
	ProgramTaskAllStart(builder)
	ProgramTaskAllAddTasks(builder, TasksOffset)
	return ProgramTaskAllEnd(builder)
}

func (rcv *ProgramTaskAll) UnPackTo(t *ProgramTaskAllT) {
	TasksLength := rcv.TasksLength()
	t.Tasks = make([]*ProgramTaskT, TasksLength)
	for j := 0; j < TasksLength; j++ {
		x := ProgramTask{}
		rcv.Tasks(&x, j)
		t.Tasks[j] = x.UnPack()
	}
}

func (rcv *ProgramTaskAll) UnPack() *ProgramTaskAllT {
	if rcv == nil { return nil }
	t := &ProgramTaskAllT{}
	rcv.UnPackTo(t)
	return t
}

type ProgramTaskAll struct {
	_tab flatbuffers.Table
}

func GetRootAsProgramTaskAll(buf []byte, offset flatbuffers.UOffsetT) *ProgramTaskAll {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ProgramTaskAll{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsProgramTaskAll(buf []byte, offset flatbuffers.UOffsetT) *ProgramTaskAll {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ProgramTaskAll{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ProgramTaskAll) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ProgramTaskAll) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ProgramTaskAll) Tasks(obj *ProgramTask, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ProgramTaskAll) TasksLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ProgramTaskAllStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ProgramTaskAllAddTasks(builder *flatbuffers.Builder, Tasks flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Tasks), 0)
}
func ProgramTaskAllStartTasksVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ProgramTaskAllEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
