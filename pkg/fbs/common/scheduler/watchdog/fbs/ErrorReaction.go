// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"

	common__scheduler__watchdog__errorreaction__fbs "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/common/scheduler/watchdog/errorreaction/fbs"
)

type ErrorReactionT struct {
	Class *common__scheduler__watchdog__errorreaction__fbs.ClassT `json:"class"`
	Configuration *common__scheduler__watchdog__errorreaction__fbs.ConfigurationT `json:"configuration"`
	MaxConsecutiveErrors uint32 `json:"maxConsecutiveErrors"`
}

func (t *ErrorReactionT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	classOffset := t.Class.Pack(builder)
	configurationOffset := t.Configuration.Pack(builder)
	ErrorReactionStart(builder)
	ErrorReactionAddClass(builder, classOffset)
	ErrorReactionAddConfiguration(builder, configurationOffset)
	ErrorReactionAddMaxConsecutiveErrors(builder, t.MaxConsecutiveErrors)
	return ErrorReactionEnd(builder)
}

func (rcv *ErrorReaction) UnPackTo(t *ErrorReactionT) {
	t.Class = rcv.Class(nil).UnPack()
	t.Configuration = rcv.Configuration(nil).UnPack()
	t.MaxConsecutiveErrors = rcv.MaxConsecutiveErrors()
}

func (rcv *ErrorReaction) UnPack() *ErrorReactionT {
	if rcv == nil { return nil }
	t := &ErrorReactionT{}
	rcv.UnPackTo(t)
	return t
}

type ErrorReaction struct {
	_tab flatbuffers.Table
}

func GetRootAsErrorReaction(buf []byte, offset flatbuffers.UOffsetT) *ErrorReaction {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ErrorReaction{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsErrorReaction(buf []byte, offset flatbuffers.UOffsetT) *ErrorReaction {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ErrorReaction{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ErrorReaction) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ErrorReaction) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ErrorReaction) Class(obj *common__scheduler__watchdog__errorreaction__fbs.Class) *common__scheduler__watchdog__errorreaction__fbs.Class {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(common__scheduler__watchdog__errorreaction__fbs.Class)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ErrorReaction) Configuration(obj *common__scheduler__watchdog__errorreaction__fbs.Configuration) *common__scheduler__watchdog__errorreaction__fbs.Configuration {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(common__scheduler__watchdog__errorreaction__fbs.Configuration)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ErrorReaction) MaxConsecutiveErrors() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ErrorReaction) MutateMaxConsecutiveErrors(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func ErrorReactionStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ErrorReactionAddClass(builder *flatbuffers.Builder, class flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(class), 0)
}
func ErrorReactionAddConfiguration(builder *flatbuffers.Builder, configuration flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(configuration), 0)
}
func ErrorReactionAddMaxConsecutiveErrors(builder *flatbuffers.Builder, maxConsecutiveErrors uint32) {
	builder.PrependUint32Slot(2, maxConsecutiveErrors, 0)
}
func ErrorReactionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
