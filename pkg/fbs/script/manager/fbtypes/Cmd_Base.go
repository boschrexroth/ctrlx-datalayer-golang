// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd_BaseT struct {
	Name string
	Source string
	Line uint64
}

func (t *Cmd_BaseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	sourceOffset := builder.CreateString(t.Source)
	Cmd_BaseStart(builder)
	Cmd_BaseAddName(builder, nameOffset)
	Cmd_BaseAddSource(builder, sourceOffset)
	Cmd_BaseAddLine(builder, t.Line)
	return Cmd_BaseEnd(builder)
}

func (rcv *Cmd_Base) UnPackTo(t *Cmd_BaseT) {
	t.Name = string(rcv.Name())
	t.Source = string(rcv.Source())
	t.Line = rcv.Line()
}

func (rcv *Cmd_Base) UnPack() *Cmd_BaseT {
	if rcv == nil { return nil }
	t := &Cmd_BaseT{}
	rcv.UnPackTo(t)
	return t
}

type Cmd_Base struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd_Base(buf []byte, offset flatbuffers.UOffsetT) *Cmd_Base {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd_Base{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd_Base(buf []byte, offset flatbuffers.UOffsetT) *Cmd_Base {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd_Base{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd_Base) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd_Base) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd_Base) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Cmd_Base) Source() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Cmd_Base) Line() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Cmd_Base) MutateLine(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

func Cmd_BaseStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func Cmd_BaseAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func Cmd_BaseAddSource(builder *flatbuffers.Builder, source flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(source), 0)
}
func Cmd_BaseAddLine(builder *flatbuffers.Builder, line uint64) {
	builder.PrependUint64Slot(2, line, 0)
}
func Cmd_BaseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
