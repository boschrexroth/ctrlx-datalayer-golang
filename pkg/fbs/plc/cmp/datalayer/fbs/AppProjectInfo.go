// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AppProjectInfoT struct {
	Name string
	Title string
	Version string
	Author string
	Description string
}

func (t *AppProjectInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	titleOffset := builder.CreateString(t.Title)
	versionOffset := builder.CreateString(t.Version)
	authorOffset := builder.CreateString(t.Author)
	descriptionOffset := builder.CreateString(t.Description)
	AppProjectInfoStart(builder)
	AppProjectInfoAddName(builder, nameOffset)
	AppProjectInfoAddTitle(builder, titleOffset)
	AppProjectInfoAddVersion(builder, versionOffset)
	AppProjectInfoAddAuthor(builder, authorOffset)
	AppProjectInfoAddDescription(builder, descriptionOffset)
	return AppProjectInfoEnd(builder)
}

func (rcv *AppProjectInfo) UnPackTo(t *AppProjectInfoT) {
	t.Name = string(rcv.Name())
	t.Title = string(rcv.Title())
	t.Version = string(rcv.Version())
	t.Author = string(rcv.Author())
	t.Description = string(rcv.Description())
}

func (rcv *AppProjectInfo) UnPack() *AppProjectInfoT {
	if rcv == nil { return nil }
	t := &AppProjectInfoT{}
	rcv.UnPackTo(t)
	return t
}

type AppProjectInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsAppProjectInfo(buf []byte, offset flatbuffers.UOffsetT) *AppProjectInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AppProjectInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAppProjectInfo(buf []byte, offset flatbuffers.UOffsetT) *AppProjectInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AppProjectInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AppProjectInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AppProjectInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AppProjectInfo) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AppProjectInfo) Title() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AppProjectInfo) Version() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AppProjectInfo) Author() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AppProjectInfo) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func AppProjectInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func AppProjectInfoAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func AppProjectInfoAddTitle(builder *flatbuffers.Builder, title flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(title), 0)
}
func AppProjectInfoAddVersion(builder *flatbuffers.Builder, version flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(version), 0)
}
func AppProjectInfoAddAuthor(builder *flatbuffers.Builder, author flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(author), 0)
}
func AppProjectInfoAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(description), 0)
}
func AppProjectInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}