// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StartupErrorReaction struct {
	_tab flatbuffers.Table
}

func GetRootAsStartupErrorReaction(buf []byte, offset flatbuffers.UOffsetT) *StartupErrorReaction {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StartupErrorReaction{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStartupErrorReaction(buf []byte, offset flatbuffers.UOffsetT) *StartupErrorReaction {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StartupErrorReaction{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StartupErrorReaction) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StartupErrorReaction) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StartupErrorReaction) ErrorReaction() CurrentErrorReaction {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return CurrentErrorReaction(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *StartupErrorReaction) MutateErrorReaction(n CurrentErrorReaction) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func StartupErrorReactionStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func StartupErrorReactionAddErrorReaction(builder *flatbuffers.Builder, errorReaction CurrentErrorReaction) {
	builder.PrependInt8Slot(0, int8(errorReaction), 0)
}
func StartupErrorReactionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
