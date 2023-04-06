// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configuration parameter of a single axis transformation
type KinCfgAxsTrafoAllParamT struct {
	AxisAssignment *KinCfgAxsTrafoAxisAssignmentT
	Groups []*KinCfgAxsTrafoParamGroupT
	General *KinCfgAxsTrafoParamGroupT
}

func (t *KinCfgAxsTrafoAllParamT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	axisAssignmentOffset := t.AxisAssignment.Pack(builder)
	groupsOffset := flatbuffers.UOffsetT(0)
	if t.Groups != nil {
		groupsLength := len(t.Groups)
		groupsOffsets := make([]flatbuffers.UOffsetT, groupsLength)
		for j := 0; j < groupsLength; j++ {
			groupsOffsets[j] = t.Groups[j].Pack(builder)
		}
		KinCfgAxsTrafoAllParamStartGroupsVector(builder, groupsLength)
		for j := groupsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(groupsOffsets[j])
		}
		groupsOffset = builder.EndVector(groupsLength)
	}
	generalOffset := t.General.Pack(builder)
	KinCfgAxsTrafoAllParamStart(builder)
	KinCfgAxsTrafoAllParamAddAxisAssignment(builder, axisAssignmentOffset)
	KinCfgAxsTrafoAllParamAddGroups(builder, groupsOffset)
	KinCfgAxsTrafoAllParamAddGeneral(builder, generalOffset)
	return KinCfgAxsTrafoAllParamEnd(builder)
}

func (rcv *KinCfgAxsTrafoAllParam) UnPackTo(t *KinCfgAxsTrafoAllParamT) {
	t.AxisAssignment = rcv.AxisAssignment(nil).UnPack()
	groupsLength := rcv.GroupsLength()
	t.Groups = make([]*KinCfgAxsTrafoParamGroupT, groupsLength)
	for j := 0; j < groupsLength; j++ {
		x := KinCfgAxsTrafoParamGroup{}
		rcv.Groups(&x, j)
		t.Groups[j] = x.UnPack()
	}
	t.General = rcv.General(nil).UnPack()
}

func (rcv *KinCfgAxsTrafoAllParam) UnPack() *KinCfgAxsTrafoAllParamT {
	if rcv == nil { return nil }
	t := &KinCfgAxsTrafoAllParamT{}
	rcv.UnPackTo(t)
	return t
}

type KinCfgAxsTrafoAllParam struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCfgAxsTrafoAllParam(buf []byte, offset flatbuffers.UOffsetT) *KinCfgAxsTrafoAllParam {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCfgAxsTrafoAllParam{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCfgAxsTrafoAllParam(buf []byte, offset flatbuffers.UOffsetT) *KinCfgAxsTrafoAllParam {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCfgAxsTrafoAllParam{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCfgAxsTrafoAllParam) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCfgAxsTrafoAllParam) Table() flatbuffers.Table {
	return rcv._tab
}

/// axis assignment 
func (rcv *KinCfgAxsTrafoAllParam) AxisAssignment(obj *KinCfgAxsTrafoAxisAssignment) *KinCfgAxsTrafoAxisAssignment {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(KinCfgAxsTrafoAxisAssignment)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// axis assignment 
/// all configuration parameter groups of this axis transformation
func (rcv *KinCfgAxsTrafoAllParam) Groups(obj *KinCfgAxsTrafoParamGroup, j int) bool {
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

func (rcv *KinCfgAxsTrafoAllParam) GroupsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// all configuration parameter groups of this axis transformation
/// all general parameter of this axis transformation
func (rcv *KinCfgAxsTrafoAllParam) General(obj *KinCfgAxsTrafoParamGroup) *KinCfgAxsTrafoParamGroup {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(KinCfgAxsTrafoParamGroup)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// all general parameter of this axis transformation
func KinCfgAxsTrafoAllParamStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func KinCfgAxsTrafoAllParamAddAxisAssignment(builder *flatbuffers.Builder, axisAssignment flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(axisAssignment), 0)
}
func KinCfgAxsTrafoAllParamAddGroups(builder *flatbuffers.Builder, groups flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(groups), 0)
}
func KinCfgAxsTrafoAllParamStartGroupsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KinCfgAxsTrafoAllParamAddGeneral(builder *flatbuffers.Builder, general flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(general), 0)
}
func KinCfgAxsTrafoAllParamEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
