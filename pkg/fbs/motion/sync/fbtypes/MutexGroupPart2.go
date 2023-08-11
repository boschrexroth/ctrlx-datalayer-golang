// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// The second group of optional parameters, which mutually exclude each other
type MutexGroupPart2T struct {
	GroupPart2 []*MutexGroupPart1T `json:"groupPart2"`
}

func (t *MutexGroupPart2T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	groupPart2Offset := flatbuffers.UOffsetT(0)
	if t.GroupPart2 != nil {
		groupPart2Length := len(t.GroupPart2)
		groupPart2Offsets := make([]flatbuffers.UOffsetT, groupPart2Length)
		for j := 0; j < groupPart2Length; j++ {
			groupPart2Offsets[j] = t.GroupPart2[j].Pack(builder)
		}
		MutexGroupPart2StartGroupPart2Vector(builder, groupPart2Length)
		for j := groupPart2Length - 1; j >= 0; j-- {
			builder.PrependUOffsetT(groupPart2Offsets[j])
		}
		groupPart2Offset = builder.EndVector(groupPart2Length)
	}
	MutexGroupPart2Start(builder)
	MutexGroupPart2AddGroupPart2(builder, groupPart2Offset)
	return MutexGroupPart2End(builder)
}

func (rcv *MutexGroupPart2) UnPackTo(t *MutexGroupPart2T) {
	groupPart2Length := rcv.GroupPart2Length()
	t.GroupPart2 = make([]*MutexGroupPart1T, groupPart2Length)
	for j := 0; j < groupPart2Length; j++ {
		x := MutexGroupPart1{}
		rcv.GroupPart2(&x, j)
		t.GroupPart2[j] = x.UnPack()
	}
}

func (rcv *MutexGroupPart2) UnPack() *MutexGroupPart2T {
	if rcv == nil { return nil }
	t := &MutexGroupPart2T{}
	rcv.UnPackTo(t)
	return t
}

type MutexGroupPart2 struct {
	_tab flatbuffers.Table
}

func GetRootAsMutexGroupPart2(buf []byte, offset flatbuffers.UOffsetT) *MutexGroupPart2 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MutexGroupPart2{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMutexGroupPart2(buf []byte, offset flatbuffers.UOffsetT) *MutexGroupPart2 {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MutexGroupPart2{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MutexGroupPart2) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MutexGroupPart2) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MutexGroupPart2) GroupPart2(obj *MutexGroupPart1, j int) bool {
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

func (rcv *MutexGroupPart2) GroupPart2Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MutexGroupPart2Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MutexGroupPart2AddGroupPart2(builder *flatbuffers.Builder, groupPart2 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(groupPart2), 0)
}
func MutexGroupPart2StartGroupPart2Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MutexGroupPart2End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
