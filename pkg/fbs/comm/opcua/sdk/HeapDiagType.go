// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package sdk

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type HeapDiagTypeT struct {
	UsedMemory uint32 `json:"usedMemory"`
	MaxUsedMemory uint32 `json:"maxUsedMemory"`
}

func (t *HeapDiagTypeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	HeapDiagTypeStart(builder)
	HeapDiagTypeAddUsedMemory(builder, t.UsedMemory)
	HeapDiagTypeAddMaxUsedMemory(builder, t.MaxUsedMemory)
	return HeapDiagTypeEnd(builder)
}

func (rcv *HeapDiagType) UnPackTo(t *HeapDiagTypeT) {
	t.UsedMemory = rcv.UsedMemory()
	t.MaxUsedMemory = rcv.MaxUsedMemory()
}

func (rcv *HeapDiagType) UnPack() *HeapDiagTypeT {
	if rcv == nil { return nil }
	t := &HeapDiagTypeT{}
	rcv.UnPackTo(t)
	return t
}

type HeapDiagType struct {
	_tab flatbuffers.Table
}

func GetRootAsHeapDiagType(buf []byte, offset flatbuffers.UOffsetT) *HeapDiagType {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &HeapDiagType{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsHeapDiagType(buf []byte, offset flatbuffers.UOffsetT) *HeapDiagType {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &HeapDiagType{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *HeapDiagType) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *HeapDiagType) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *HeapDiagType) UsedMemory() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *HeapDiagType) MutateUsedMemory(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *HeapDiagType) MaxUsedMemory() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *HeapDiagType) MutateMaxUsedMemory(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func HeapDiagTypeStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func HeapDiagTypeAddUsedMemory(builder *flatbuffers.Builder, usedMemory uint32) {
	builder.PrependUint32Slot(0, usedMemory, 0)
}
func HeapDiagTypeAddMaxUsedMemory(builder *flatbuffers.Builder, maxUsedMemory uint32) {
	builder.PrependUint32Slot(1, maxUsedMemory, 0)
}
func HeapDiagTypeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
