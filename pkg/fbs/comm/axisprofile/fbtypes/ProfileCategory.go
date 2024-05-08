// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ProfileCategoryT struct {
	Category CategoryOfProfile `json:"category"`
}

func (t *ProfileCategoryT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ProfileCategoryStart(builder)
	ProfileCategoryAddCategory(builder, t.Category)
	return ProfileCategoryEnd(builder)
}

func (rcv *ProfileCategory) UnPackTo(t *ProfileCategoryT) {
	t.Category = rcv.Category()
}

func (rcv *ProfileCategory) UnPack() *ProfileCategoryT {
	if rcv == nil { return nil }
	t := &ProfileCategoryT{}
	rcv.UnPackTo(t)
	return t
}

type ProfileCategory struct {
	_tab flatbuffers.Table
}

func GetRootAsProfileCategory(buf []byte, offset flatbuffers.UOffsetT) *ProfileCategory {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ProfileCategory{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsProfileCategory(buf []byte, offset flatbuffers.UOffsetT) *ProfileCategory {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ProfileCategory{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ProfileCategory) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ProfileCategory) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ProfileCategory) Category() CategoryOfProfile {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return CategoryOfProfile(rcv._tab.GetUint32(o + rcv._tab.Pos))
	}
	return 1
}

func (rcv *ProfileCategory) MutateCategory(n CategoryOfProfile) bool {
	return rcv._tab.MutateUint32Slot(4, uint32(n))
}

func ProfileCategoryStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ProfileCategoryAddCategory(builder *flatbuffers.Builder, category CategoryOfProfile) {
	builder.PrependUint32Slot(0, uint32(category), 1)
}
func ProfileCategoryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}