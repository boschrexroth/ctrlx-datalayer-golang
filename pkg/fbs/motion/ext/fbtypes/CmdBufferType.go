// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameter defines the buffer type of generic command
type CmdBufferTypeT struct {
	Buffered bool
}

func (t *CmdBufferTypeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	CmdBufferTypeStart(builder)
	CmdBufferTypeAddBuffered(builder, t.Buffered)
	return CmdBufferTypeEnd(builder)
}

func (rcv *CmdBufferType) UnPackTo(t *CmdBufferTypeT) {
	t.Buffered = rcv.Buffered()
}

func (rcv *CmdBufferType) UnPack() *CmdBufferTypeT {
	if rcv == nil { return nil }
	t := &CmdBufferTypeT{}
	rcv.UnPackTo(t)
	return t
}

type CmdBufferType struct {
	_tab flatbuffers.Table
}

func GetRootAsCmdBufferType(buf []byte, offset flatbuffers.UOffsetT) *CmdBufferType {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CmdBufferType{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmdBufferType(buf []byte, offset flatbuffers.UOffsetT) *CmdBufferType {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CmdBufferType{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CmdBufferType) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CmdBufferType) Table() flatbuffers.Table {
	return rcv._tab
}

/// buffered type for generic command
func (rcv *CmdBufferType) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// buffered type for generic command
func (rcv *CmdBufferType) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func CmdBufferTypeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func CmdBufferTypeAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(0, buffered, false)
}
func CmdBufferTypeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
