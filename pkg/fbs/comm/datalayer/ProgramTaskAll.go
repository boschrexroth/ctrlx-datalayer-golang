// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

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