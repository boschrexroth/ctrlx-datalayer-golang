// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Memory struct {
	_tab flatbuffers.Table
}

func GetRootAsMemory(buf []byte, offset flatbuffers.UOffsetT) *Memory {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Memory{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMemory(buf []byte, offset flatbuffers.UOffsetT) *Memory {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Memory{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Memory) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Memory) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Memory) Type() MemoryType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return MemoryType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Memory) MutateType(n MemoryType) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func (rcv *Memory) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Memory) SizeBytes() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Memory) MutateSizeBytes(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func MemoryStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func MemoryAddType(builder *flatbuffers.Builder, type_ MemoryType) {
	builder.PrependInt8Slot(0, int8(type_), 0)
}
func MemoryAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(id), 0)
}
func MemoryAddSizeBytes(builder *flatbuffers.Builder, sizeBytes uint32) {
	builder.PrependUint32Slot(2, sizeBytes, 0)
}
func MemoryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}