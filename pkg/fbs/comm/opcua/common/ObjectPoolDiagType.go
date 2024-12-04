// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package common

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ObjectPoolDiagTypeT struct {
	Name string `json:"name"`
	UsedObjects uint32 `json:"usedObjects"`
	MaxUsedObjects uint32 `json:"maxUsedObjects"`
	NumObjects uint32 `json:"numObjects"`
}

func (t *ObjectPoolDiagTypeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	ObjectPoolDiagTypeStart(builder)
	ObjectPoolDiagTypeAddName(builder, nameOffset)
	ObjectPoolDiagTypeAddUsedObjects(builder, t.UsedObjects)
	ObjectPoolDiagTypeAddMaxUsedObjects(builder, t.MaxUsedObjects)
	ObjectPoolDiagTypeAddNumObjects(builder, t.NumObjects)
	return ObjectPoolDiagTypeEnd(builder)
}

func (rcv *ObjectPoolDiagType) UnPackTo(t *ObjectPoolDiagTypeT) {
	t.Name = string(rcv.Name())
	t.UsedObjects = rcv.UsedObjects()
	t.MaxUsedObjects = rcv.MaxUsedObjects()
	t.NumObjects = rcv.NumObjects()
}

func (rcv *ObjectPoolDiagType) UnPack() *ObjectPoolDiagTypeT {
	if rcv == nil { return nil }
	t := &ObjectPoolDiagTypeT{}
	rcv.UnPackTo(t)
	return t
}

type ObjectPoolDiagType struct {
	_tab flatbuffers.Table
}

func GetRootAsObjectPoolDiagType(buf []byte, offset flatbuffers.UOffsetT) *ObjectPoolDiagType {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ObjectPoolDiagType{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsObjectPoolDiagType(buf []byte, offset flatbuffers.UOffsetT) *ObjectPoolDiagType {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ObjectPoolDiagType{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ObjectPoolDiagType) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ObjectPoolDiagType) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ObjectPoolDiagType) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ObjectPoolDiagType) UsedObjects() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ObjectPoolDiagType) MutateUsedObjects(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *ObjectPoolDiagType) MaxUsedObjects() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ObjectPoolDiagType) MutateMaxUsedObjects(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *ObjectPoolDiagType) NumObjects() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ObjectPoolDiagType) MutateNumObjects(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func ObjectPoolDiagTypeStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ObjectPoolDiagTypeAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func ObjectPoolDiagTypeAddUsedObjects(builder *flatbuffers.Builder, usedObjects uint32) {
	builder.PrependUint32Slot(1, usedObjects, 0)
}
func ObjectPoolDiagTypeAddMaxUsedObjects(builder *flatbuffers.Builder, maxUsedObjects uint32) {
	builder.PrependUint32Slot(2, maxUsedObjects, 0)
}
func ObjectPoolDiagTypeAddNumObjects(builder *flatbuffers.Builder, numObjects uint32) {
	builder.PrependUint32Slot(3, numObjects, 0)
}
func ObjectPoolDiagTypeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}