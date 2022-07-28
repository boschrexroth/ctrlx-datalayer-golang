// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LocaleTextT struct {
	Id string
	Text string
}

func (t *LocaleTextT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	idOffset := builder.CreateString(t.Id)
	textOffset := builder.CreateString(t.Text)
	LocaleTextStart(builder)
	LocaleTextAddId(builder, idOffset)
	LocaleTextAddText(builder, textOffset)
	return LocaleTextEnd(builder)
}

func (rcv *LocaleText) UnPackTo(t *LocaleTextT) {
	t.Id = string(rcv.Id())
	t.Text = string(rcv.Text())
}

func (rcv *LocaleText) UnPack() *LocaleTextT {
	if rcv == nil { return nil }
	t := &LocaleTextT{}
	rcv.UnPackTo(t)
	return t
}

type LocaleText struct {
	_tab flatbuffers.Table
}

func GetRootAsLocaleText(buf []byte, offset flatbuffers.UOffsetT) *LocaleText {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LocaleText{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLocaleText(buf []byte, offset flatbuffers.UOffsetT) *LocaleText {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LocaleText{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LocaleText) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LocaleText) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LocaleText) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LocaleText) Text() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func LocaleTextStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func LocaleTextAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func LocaleTextAddText(builder *flatbuffers.Builder, text flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(text), 0)
}
func LocaleTextEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
