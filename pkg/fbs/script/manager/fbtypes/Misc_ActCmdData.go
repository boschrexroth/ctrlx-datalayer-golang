// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Misc_ActCmdDataT struct {
	SrcName string `json:"srcName"`
	SrcType string `json:"srcType"`
	SrcLine uint64 `json:"srcLine"`
}

func (t *Misc_ActCmdDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	srcNameOffset := flatbuffers.UOffsetT(0)
	if t.SrcName != "" {
		srcNameOffset = builder.CreateString(t.SrcName)
	}
	srcTypeOffset := flatbuffers.UOffsetT(0)
	if t.SrcType != "" {
		srcTypeOffset = builder.CreateString(t.SrcType)
	}
	Misc_ActCmdDataStart(builder)
	Misc_ActCmdDataAddSrcName(builder, srcNameOffset)
	Misc_ActCmdDataAddSrcType(builder, srcTypeOffset)
	Misc_ActCmdDataAddSrcLine(builder, t.SrcLine)
	return Misc_ActCmdDataEnd(builder)
}

func (rcv *Misc_ActCmdData) UnPackTo(t *Misc_ActCmdDataT) {
	t.SrcName = string(rcv.SrcName())
	t.SrcType = string(rcv.SrcType())
	t.SrcLine = rcv.SrcLine()
}

func (rcv *Misc_ActCmdData) UnPack() *Misc_ActCmdDataT {
	if rcv == nil { return nil }
	t := &Misc_ActCmdDataT{}
	rcv.UnPackTo(t)
	return t
}

type Misc_ActCmdData struct {
	_tab flatbuffers.Table
}

func GetRootAsMisc_ActCmdData(buf []byte, offset flatbuffers.UOffsetT) *Misc_ActCmdData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Misc_ActCmdData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMisc_ActCmdData(buf []byte, offset flatbuffers.UOffsetT) *Misc_ActCmdData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Misc_ActCmdData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Misc_ActCmdData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Misc_ActCmdData) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Misc_ActCmdData) SrcName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Misc_ActCmdData) SrcType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Misc_ActCmdData) SrcLine() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Misc_ActCmdData) MutateSrcLine(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

func Misc_ActCmdDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func Misc_ActCmdDataAddSrcName(builder *flatbuffers.Builder, srcName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(srcName), 0)
}
func Misc_ActCmdDataAddSrcType(builder *flatbuffers.Builder, srcType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(srcType), 0)
}
func Misc_ActCmdDataAddSrcLine(builder *flatbuffers.Builder, srcLine uint64) {
	builder.PrependUint64Slot(2, srcLine, 0)
}
func Misc_ActCmdDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
