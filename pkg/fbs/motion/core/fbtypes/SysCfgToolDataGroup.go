// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	"bytes"
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configuration of a single group of sets for tool data
type SysCfgToolDataGroupT struct {
	GroupName string `json:"groupName"`
	Sets []string `json:"sets"`
}

func (t *SysCfgToolDataGroupT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	groupNameOffset := flatbuffers.UOffsetT(0)
	if t.GroupName != "" {
		groupNameOffset = builder.CreateString(t.GroupName)
	}
	setsOffset := flatbuffers.UOffsetT(0)
	if t.Sets != nil {
		setsLength := len(t.Sets)
		setsOffsets := make([]flatbuffers.UOffsetT, setsLength)
		for j := 0; j < setsLength; j++ {
			setsOffsets[j] = builder.CreateString(t.Sets[j])
		}
		SysCfgToolDataGroupStartSetsVector(builder, setsLength)
		for j := setsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(setsOffsets[j])
		}
		setsOffset = builder.EndVector(setsLength)
	}
	SysCfgToolDataGroupStart(builder)
	SysCfgToolDataGroupAddGroupName(builder, groupNameOffset)
	SysCfgToolDataGroupAddSets(builder, setsOffset)
	return SysCfgToolDataGroupEnd(builder)
}

func (rcv *SysCfgToolDataGroup) UnPackTo(t *SysCfgToolDataGroupT) {
	t.GroupName = string(rcv.GroupName())
	setsLength := rcv.SetsLength()
	t.Sets = make([]string, setsLength)
	for j := 0; j < setsLength; j++ {
		t.Sets[j] = string(rcv.Sets(j))
	}
}

func (rcv *SysCfgToolDataGroup) UnPack() *SysCfgToolDataGroupT {
	if rcv == nil { return nil }
	t := &SysCfgToolDataGroupT{}
	rcv.UnPackTo(t)
	return t
}

type SysCfgToolDataGroup struct {
	_tab flatbuffers.Table
}

func GetRootAsSysCfgToolDataGroup(buf []byte, offset flatbuffers.UOffsetT) *SysCfgToolDataGroup {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SysCfgToolDataGroup{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSysCfgToolDataGroup(buf []byte, offset flatbuffers.UOffsetT) *SysCfgToolDataGroup {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SysCfgToolDataGroup{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SysCfgToolDataGroup) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SysCfgToolDataGroup) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the group (required for load/save)
func (rcv *SysCfgToolDataGroup) GroupName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the group (required for load/save)
func SysCfgToolDataGroupKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &SysCfgToolDataGroup{}
	obj2 := &SysCfgToolDataGroup{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return string(obj1.GroupName()) < string(obj2.GroupName())
}

func (rcv *SysCfgToolDataGroup) LookupByKey(key string, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	bKey := []byte(key)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &SysCfgToolDataGroup{}
		obj.Init(buf, tableOffset)
		comp := bytes.Compare(obj.GroupName(), bKey)
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

/// vector of tool data sets in this group (sequence matters!)
func (rcv *SysCfgToolDataGroup) Sets(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *SysCfgToolDataGroup) SetsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of tool data sets in this group (sequence matters!)
func SysCfgToolDataGroupStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SysCfgToolDataGroupAddGroupName(builder *flatbuffers.Builder, groupName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(groupName), 0)
}
func SysCfgToolDataGroupAddSets(builder *flatbuffers.Builder, sets flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(sets), 0)
}
func SysCfgToolDataGroupStartSetsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SysCfgToolDataGroupEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}