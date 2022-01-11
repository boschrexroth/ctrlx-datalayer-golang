// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Data of all supported units
type UnitDataAll struct {
	_tab flatbuffers.Table
}

func GetRootAsUnitDataAll(buf []byte, offset flatbuffers.UOffsetT) *UnitDataAll {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UnitDataAll{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUnitDataAll(buf []byte, offset flatbuffers.UOffsetT) *UnitDataAll {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &UnitDataAll{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *UnitDataAll) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UnitDataAll) Table() flatbuffers.Table {
	return rcv._tab
}

/// vector of all supported units
func (rcv *UnitDataAll) Supported(obj *UnitDataSingle, j int) bool {
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

func (rcv *UnitDataAll) SupportedLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of all supported units
func UnitDataAllStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func UnitDataAllAddSupported(builder *flatbuffers.Builder, supported flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(supported), 0)
}
func UnitDataAllStartSupportedVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func UnitDataAllEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}