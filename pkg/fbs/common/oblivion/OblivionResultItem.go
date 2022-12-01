// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package oblivion

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OblivionResultItemT struct {
	Count uint32
	Bytes uint32
	StackTrace []string
	ThreadName string
	TaskID uint64
}

func (t *OblivionResultItemT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	stackTraceOffset := flatbuffers.UOffsetT(0)
	if t.StackTrace != nil {
		stackTraceLength := len(t.StackTrace)
		stackTraceOffsets := make([]flatbuffers.UOffsetT, stackTraceLength)
		for j := 0; j < stackTraceLength; j++ {
			stackTraceOffsets[j] = builder.CreateString(t.StackTrace[j])
		}
		OblivionResultItemStartStackTraceVector(builder, stackTraceLength)
		for j := stackTraceLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(stackTraceOffsets[j])
		}
		stackTraceOffset = builder.EndVector(stackTraceLength)
	}
	threadNameOffset := builder.CreateString(t.ThreadName)
	OblivionResultItemStart(builder)
	OblivionResultItemAddCount(builder, t.Count)
	OblivionResultItemAddBytes(builder, t.Bytes)
	OblivionResultItemAddStackTrace(builder, stackTraceOffset)
	OblivionResultItemAddThreadName(builder, threadNameOffset)
	OblivionResultItemAddTaskID(builder, t.TaskID)
	return OblivionResultItemEnd(builder)
}

func (rcv *OblivionResultItem) UnPackTo(t *OblivionResultItemT) {
	t.Count = rcv.Count()
	t.Bytes = rcv.Bytes()
	stackTraceLength := rcv.StackTraceLength()
	t.StackTrace = make([]string, stackTraceLength)
	for j := 0; j < stackTraceLength; j++ {
		t.StackTrace[j] = string(rcv.StackTrace(j))
	}
	t.ThreadName = string(rcv.ThreadName())
	t.TaskID = rcv.TaskID()
}

func (rcv *OblivionResultItem) UnPack() *OblivionResultItemT {
	if rcv == nil { return nil }
	t := &OblivionResultItemT{}
	rcv.UnPackTo(t)
	return t
}

type OblivionResultItem struct {
	_tab flatbuffers.Table
}

func GetRootAsOblivionResultItem(buf []byte, offset flatbuffers.UOffsetT) *OblivionResultItem {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OblivionResultItem{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsOblivionResultItem(buf []byte, offset flatbuffers.UOffsetT) *OblivionResultItem {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &OblivionResultItem{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *OblivionResultItem) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OblivionResultItem) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OblivionResultItem) Count() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OblivionResultItem) MutateCount(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *OblivionResultItem) Bytes() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OblivionResultItem) MutateBytes(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *OblivionResultItem) StackTrace(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *OblivionResultItem) StackTraceLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *OblivionResultItem) ThreadName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *OblivionResultItem) TaskID() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OblivionResultItem) MutateTaskID(n uint64) bool {
	return rcv._tab.MutateUint64Slot(12, n)
}

func OblivionResultItemStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func OblivionResultItemAddCount(builder *flatbuffers.Builder, count uint32) {
	builder.PrependUint32Slot(0, count, 0)
}
func OblivionResultItemAddBytes(builder *flatbuffers.Builder, bytes uint32) {
	builder.PrependUint32Slot(1, bytes, 0)
}
func OblivionResultItemAddStackTrace(builder *flatbuffers.Builder, stackTrace flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(stackTrace), 0)
}
func OblivionResultItemStartStackTraceVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func OblivionResultItemAddThreadName(builder *flatbuffers.Builder, threadName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(threadName), 0)
}
func OblivionResultItemAddTaskID(builder *flatbuffers.Builder, taskID uint64) {
	builder.PrependUint64Slot(4, taskID, 0)
}
func OblivionResultItemEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}