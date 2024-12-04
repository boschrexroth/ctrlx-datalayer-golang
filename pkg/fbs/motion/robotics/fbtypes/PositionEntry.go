// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"

	motion__core__fbtypes "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/motion/core/fbtypes"
)

type PositionEntryT struct {
	Meaning motion__core__fbtypes.Meaning `json:"meaning"`
	Value float64 `json:"value"`
	Unit string `json:"unit"`
}

func (t *PositionEntryT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	unitOffset := flatbuffers.UOffsetT(0)
	if t.Unit != "" {
		unitOffset = builder.CreateString(t.Unit)
	}
	PositionEntryStart(builder)
	PositionEntryAddMeaning(builder, t.Meaning)
	PositionEntryAddValue(builder, t.Value)
	PositionEntryAddUnit(builder, unitOffset)
	return PositionEntryEnd(builder)
}

func (rcv *PositionEntry) UnPackTo(t *PositionEntryT) {
	t.Meaning = rcv.Meaning()
	t.Value = rcv.Value()
	t.Unit = string(rcv.Unit())
}

func (rcv *PositionEntry) UnPack() *PositionEntryT {
	if rcv == nil { return nil }
	t := &PositionEntryT{}
	rcv.UnPackTo(t)
	return t
}

type PositionEntry struct {
	_tab flatbuffers.Table
}

func GetRootAsPositionEntry(buf []byte, offset flatbuffers.UOffsetT) *PositionEntry {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PositionEntry{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPositionEntry(buf []byte, offset flatbuffers.UOffsetT) *PositionEntry {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PositionEntry{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *PositionEntry) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PositionEntry) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PositionEntry) Meaning() motion__core__fbtypes.Meaning {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return motion__core__fbtypes.Meaning(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *PositionEntry) MutateMeaning(n motion__core__fbtypes.Meaning) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func PositionEntryKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &PositionEntry{}
	obj2 := &PositionEntry{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return obj1.Meaning() < obj2.Meaning()
}

func (rcv *PositionEntry) LookupByKey(key motion__core__fbtypes.Meaning, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &PositionEntry{}
		obj.Init(buf, tableOffset)
		val := obj.Meaning()
		comp := 0
		if val > key {
			comp = 1
		} else if val < key {
			comp = -1
		}
		if comp > 0 {
			span = middle
		} else if comp < 0 {
			middle += 1
			start += middle
			span -= middle
		} else {
			rcv.Init(buf, tableOffset)
			return true
		}
	}
	return false
}

func (rcv *PositionEntry) Value() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *PositionEntry) MutateValue(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func (rcv *PositionEntry) Unit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func PositionEntryStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func PositionEntryAddMeaning(builder *flatbuffers.Builder, meaning motion__core__fbtypes.Meaning) {
	builder.PrependInt8Slot(0, int8(meaning), 0)
}
func PositionEntryAddValue(builder *flatbuffers.Builder, value float64) {
	builder.PrependFloat64Slot(1, value, 0.0)
}
func PositionEntryAddUnit(builder *flatbuffers.Builder, unit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(unit), 0)
}
func PositionEntryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}