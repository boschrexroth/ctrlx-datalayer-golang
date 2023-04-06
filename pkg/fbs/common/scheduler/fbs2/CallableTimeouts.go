// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs2

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// General settings for timeouts when switching callable operation states
type CallableTimeoutsT struct {
	Begin uint32
	Execute uint32
	End uint32
}

func (t *CallableTimeoutsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	CallableTimeoutsStart(builder)
	CallableTimeoutsAddBegin(builder, t.Begin)
	CallableTimeoutsAddExecute(builder, t.Execute)
	CallableTimeoutsAddEnd(builder, t.End)
	return CallableTimeoutsEnd(builder)
}

func (rcv *CallableTimeouts) UnPackTo(t *CallableTimeoutsT) {
	t.Begin = rcv.Begin()
	t.Execute = rcv.Execute()
	t.End = rcv.End()
}

func (rcv *CallableTimeouts) UnPack() *CallableTimeoutsT {
	if rcv == nil { return nil }
	t := &CallableTimeoutsT{}
	rcv.UnPackTo(t)
	return t
}

type CallableTimeouts struct {
	_tab flatbuffers.Table
}

func GetRootAsCallableTimeouts(buf []byte, offset flatbuffers.UOffsetT) *CallableTimeouts {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CallableTimeouts{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCallableTimeouts(buf []byte, offset flatbuffers.UOffsetT) *CallableTimeouts {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CallableTimeouts{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CallableTimeouts) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CallableTimeouts) Table() flatbuffers.Table {
	return rcv._tab
}

/// General settings for timeouts when switching callable operation states - phase BEGIN
func (rcv *CallableTimeouts) Begin() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 1000
}

/// General settings for timeouts when switching callable operation states - phase BEGIN
func (rcv *CallableTimeouts) MutateBegin(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

/// General settings for timeouts when switching callable operation states - phase EXECUTE
func (rcv *CallableTimeouts) Execute() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 60000
}

/// General settings for timeouts when switching callable operation states - phase EXECUTE
func (rcv *CallableTimeouts) MutateExecute(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

/// General settings for timeouts when switching callable operation states - phase END
func (rcv *CallableTimeouts) End() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 1000
}

/// General settings for timeouts when switching callable operation states - phase END
func (rcv *CallableTimeouts) MutateEnd(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func CallableTimeoutsStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func CallableTimeoutsAddBegin(builder *flatbuffers.Builder, begin uint32) {
	builder.PrependUint32Slot(0, begin, 1000)
}
func CallableTimeoutsAddExecute(builder *flatbuffers.Builder, execute uint32) {
	builder.PrependUint32Slot(1, execute, 60000)
}
func CallableTimeoutsAddEnd(builder *flatbuffers.Builder, end uint32) {
	builder.PrependUint32Slot(2, end, 1000)
}
func CallableTimeoutsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
