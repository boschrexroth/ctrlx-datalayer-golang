// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"

	motion__core__fbtypes "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/motion/core/fbtypes"
)

type PositionT struct {
	Entries []*PositionEntryT `json:"entries"`
}

func (t *PositionT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	entriesOffset := flatbuffers.UOffsetT(0)
	if t.Entries != nil {
		entriesLength := len(t.Entries)
		entriesOffsets := make([]flatbuffers.UOffsetT, entriesLength)
		for j := 0; j < entriesLength; j++ {
			entriesOffsets[j] = t.Entries[j].Pack(builder)
		}
		PositionStartEntriesVector(builder, entriesLength)
		for j := entriesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(entriesOffsets[j])
		}
		entriesOffset = builder.EndVector(entriesLength)
	}
	PositionStart(builder)
	PositionAddEntries(builder, entriesOffset)
	return PositionEnd(builder)
}

func (rcv *Position) UnPackTo(t *PositionT) {
	entriesLength := rcv.EntriesLength()
	t.Entries = make([]*PositionEntryT, entriesLength)
	for j := 0; j < entriesLength; j++ {
		x := PositionEntry{}
		rcv.Entries(&x, j)
		t.Entries[j] = x.UnPack()
	}
}

func (rcv *Position) UnPack() *PositionT {
	if rcv == nil { return nil }
	t := &PositionT{}
	rcv.UnPackTo(t)
	return t
}

type Position struct {
	_tab flatbuffers.Table
}

func GetRootAsPosition(buf []byte, offset flatbuffers.UOffsetT) *Position {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Position{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPosition(buf []byte, offset flatbuffers.UOffsetT) *Position {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Position{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Position) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Position) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Position) Entries(obj *PositionEntry, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Position) EntriesByKey(obj *PositionEntry, key motion__core__fbtypes.Meaning) bool{
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *Position) EntriesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func PositionStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func PositionAddEntries(builder *flatbuffers.Builder, entries flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(entries), 0)
}
func PositionStartEntriesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PositionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}