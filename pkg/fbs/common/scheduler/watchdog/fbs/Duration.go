// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DurationT struct {
	Time uint32
	ErrorCount uint32
	Reset bool
	ErrorReaction *ErrorReactionT
}

func (t *DurationT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	errorReactionOffset := t.ErrorReaction.Pack(builder)
	DurationStart(builder)
	DurationAddTime(builder, t.Time)
	DurationAddErrorCount(builder, t.ErrorCount)
	DurationAddReset(builder, t.Reset)
	DurationAddErrorReaction(builder, errorReactionOffset)
	return DurationEnd(builder)
}

func (rcv *Duration) UnPackTo(t *DurationT) {
	t.Time = rcv.Time()
	t.ErrorCount = rcv.ErrorCount()
	t.Reset = rcv.Reset()
	t.ErrorReaction = rcv.ErrorReaction(nil).UnPack()
}

func (rcv *Duration) UnPack() *DurationT {
	if rcv == nil { return nil }
	t := &DurationT{}
	rcv.UnPackTo(t)
	return t
}

type Duration struct {
	_tab flatbuffers.Table
}

func GetRootAsDuration(buf []byte, offset flatbuffers.UOffsetT) *Duration {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Duration{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDuration(buf []byte, offset flatbuffers.UOffsetT) *Duration {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Duration{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Duration) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Duration) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Duration) Time() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Duration) MutateTime(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Duration) ErrorCount() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Duration) MutateErrorCount(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *Duration) Reset() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Duration) MutateReset(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func (rcv *Duration) ErrorReaction(obj *ErrorReaction) *ErrorReaction {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ErrorReaction)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func DurationStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DurationAddTime(builder *flatbuffers.Builder, time uint32) {
	builder.PrependUint32Slot(0, time, 0)
}
func DurationAddErrorCount(builder *flatbuffers.Builder, errorCount uint32) {
	builder.PrependUint32Slot(1, errorCount, 0)
}
func DurationAddReset(builder *flatbuffers.Builder, reset bool) {
	builder.PrependBoolSlot(2, reset, false)
}
func DurationAddErrorReaction(builder *flatbuffers.Builder, errorReaction flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(errorReaction), 0)
}
func DurationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
