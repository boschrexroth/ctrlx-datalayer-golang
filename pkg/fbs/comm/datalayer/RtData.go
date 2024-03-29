// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RtDataT struct {
	RevisionNumber uint32 `json:"revisionNumber"`
	Areas []*RtDataAreaT `json:"areas"`
}

func (t *RtDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	areasOffset := flatbuffers.UOffsetT(0)
	if t.Areas != nil {
		areasLength := len(t.Areas)
		areasOffsets := make([]flatbuffers.UOffsetT, areasLength)
		for j := 0; j < areasLength; j++ {
			areasOffsets[j] = t.Areas[j].Pack(builder)
		}
		RtDataStartAreasVector(builder, areasLength)
		for j := areasLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(areasOffsets[j])
		}
		areasOffset = builder.EndVector(areasLength)
	}
	RtDataStart(builder)
	RtDataAddRevisionNumber(builder, t.RevisionNumber)
	RtDataAddAreas(builder, areasOffset)
	return RtDataEnd(builder)
}

func (rcv *RtData) UnPackTo(t *RtDataT) {
	t.RevisionNumber = rcv.RevisionNumber()
	areasLength := rcv.AreasLength()
	t.Areas = make([]*RtDataAreaT, areasLength)
	for j := 0; j < areasLength; j++ {
		x := RtDataArea{}
		rcv.Areas(&x, j)
		t.Areas[j] = x.UnPack()
	}
}

func (rcv *RtData) UnPack() *RtDataT {
	if rcv == nil { return nil }
	t := &RtDataT{}
	rcv.UnPackTo(t)
	return t
}

type RtData struct {
	_tab flatbuffers.Table
}

func GetRootAsRtData(buf []byte, offset flatbuffers.UOffsetT) *RtData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RtData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRtData(buf []byte, offset flatbuffers.UOffsetT) *RtData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RtData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RtData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RtData) Table() flatbuffers.Table {
	return rcv._tab
}

/// current revision number of RT memory
func (rcv *RtData) RevisionNumber() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// current revision number of RT memory
func (rcv *RtData) MutateRevisionNumber(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

/// areas to read/write
func (rcv *RtData) Areas(obj *RtDataArea, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *RtData) AreasLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// areas to read/write
func RtDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func RtDataAddRevisionNumber(builder *flatbuffers.Builder, revisionNumber uint32) {
	builder.PrependUint32Slot(0, revisionNumber, 0)
}
func RtDataAddAreas(builder *flatbuffers.Builder, areas flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(areas), 0)
}
func RtDataStartAreasVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func RtDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
