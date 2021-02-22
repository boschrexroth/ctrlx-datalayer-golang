// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cyclic struct {
	_tab flatbuffers.Table
}

func GetRootAsCyclic(buf []byte, offset flatbuffers.UOffsetT) *Cyclic {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cyclic{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCyclic(buf []byte, offset flatbuffers.UOffsetT) *Cyclic {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cyclic{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cyclic) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cyclic) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cyclic) ErrorCount() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Cyclic) MutateErrorCount(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Cyclic) Reset() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Cyclic) MutateReset(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *Cyclic) ErrorReaction(obj *ErrorReaction) *ErrorReaction {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
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

func CyclicStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func CyclicAddErrorCount(builder *flatbuffers.Builder, errorCount uint32) {
	builder.PrependUint32Slot(0, errorCount, 0)
}
func CyclicAddReset(builder *flatbuffers.Builder, reset bool) {
	builder.PrependBoolSlot(1, reset, false)
}
func CyclicAddErrorReaction(builder *flatbuffers.Builder, errorReaction flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(errorReaction), 0)
}
func CyclicEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
