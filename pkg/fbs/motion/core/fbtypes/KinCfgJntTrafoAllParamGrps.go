// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configuration parameter of a joint transformation
type KinCfgJntTrafoAllParamGrpsT struct {
	Groups []*KinCfgJntTrafoParamGroupT
}

func (t *KinCfgJntTrafoAllParamGrpsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	groupsOffset := flatbuffers.UOffsetT(0)
	if t.Groups != nil {
		groupsLength := len(t.Groups)
		groupsOffsets := make([]flatbuffers.UOffsetT, groupsLength)
		for j := 0; j < groupsLength; j++ {
			groupsOffsets[j] = t.Groups[j].Pack(builder)
		}
		KinCfgJntTrafoAllParamGrpsStartGroupsVector(builder, groupsLength)
		for j := groupsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(groupsOffsets[j])
		}
		groupsOffset = builder.EndVector(groupsLength)
	}
	KinCfgJntTrafoAllParamGrpsStart(builder)
	KinCfgJntTrafoAllParamGrpsAddGroups(builder, groupsOffset)
	return KinCfgJntTrafoAllParamGrpsEnd(builder)
}

func (rcv *KinCfgJntTrafoAllParamGrps) UnPackTo(t *KinCfgJntTrafoAllParamGrpsT) {
	groupsLength := rcv.GroupsLength()
	t.Groups = make([]*KinCfgJntTrafoParamGroupT, groupsLength)
	for j := 0; j < groupsLength; j++ {
		x := KinCfgJntTrafoParamGroup{}
		rcv.Groups(&x, j)
		t.Groups[j] = x.UnPack()
	}
}

func (rcv *KinCfgJntTrafoAllParamGrps) UnPack() *KinCfgJntTrafoAllParamGrpsT {
	if rcv == nil { return nil }
	t := &KinCfgJntTrafoAllParamGrpsT{}
	rcv.UnPackTo(t)
	return t
}

type KinCfgJntTrafoAllParamGrps struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCfgJntTrafoAllParamGrps(buf []byte, offset flatbuffers.UOffsetT) *KinCfgJntTrafoAllParamGrps {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCfgJntTrafoAllParamGrps{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCfgJntTrafoAllParamGrps(buf []byte, offset flatbuffers.UOffsetT) *KinCfgJntTrafoAllParamGrps {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCfgJntTrafoAllParamGrps{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCfgJntTrafoAllParamGrps) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCfgJntTrafoAllParamGrps) Table() flatbuffers.Table {
	return rcv._tab
}

/// all configuration parameter groups of this joint transformation
func (rcv *KinCfgJntTrafoAllParamGrps) Groups(obj *KinCfgJntTrafoParamGroup, j int) bool {
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

func (rcv *KinCfgJntTrafoAllParamGrps) GroupsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// all configuration parameter groups of this joint transformation
func KinCfgJntTrafoAllParamGrpsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func KinCfgJntTrafoAllParamGrpsAddGroups(builder *flatbuffers.Builder, groups flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(groups), 0)
}
func KinCfgJntTrafoAllParamGrpsStartGroupsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KinCfgJntTrafoAllParamGrpsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
