// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configuration of the product coordinate system of this kinematics
type SysCfgPcsAll struct {
	_tab flatbuffers.Table
}

func GetRootAsSysCfgPcsAll(buf []byte, offset flatbuffers.UOffsetT) *SysCfgPcsAll {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SysCfgPcsAll{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSysCfgPcsAll(buf []byte, offset flatbuffers.UOffsetT) *SysCfgPcsAll {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SysCfgPcsAll{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SysCfgPcsAll) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SysCfgPcsAll) Table() flatbuffers.Table {
	return rcv._tab
}

/// vector of all configured PCS sets
func (rcv *SysCfgPcsAll) Sets(obj *SysCfgPcsSet, j int) bool {
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

func (rcv *SysCfgPcsAll) SetsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of all configured PCS sets
/// vector of all configured PCS groups
func (rcv *SysCfgPcsAll) Groups(obj *SysCfgPcsGroup, j int) bool {
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

func (rcv *SysCfgPcsAll) GroupsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of all configured PCS groups
func SysCfgPcsAllStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SysCfgPcsAllAddSets(builder *flatbuffers.Builder, sets flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(sets), 0)
}
func SysCfgPcsAllStartSetsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SysCfgPcsAllAddGroups(builder *flatbuffers.Builder, groups flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(groups), 0)
}
func SysCfgPcsAllStartGroupsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SysCfgPcsAllEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}