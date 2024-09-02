// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DumpRequestT struct {
	Offset uint32 `json:"offset"`
	Datalength uint32 `json:"datalength"`
}

func (t *DumpRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	DumpRequestStart(builder)
	DumpRequestAddOffset(builder, t.Offset)
	DumpRequestAddDatalength(builder, t.Datalength)
	return DumpRequestEnd(builder)
}

func (rcv *DumpRequest) UnPackTo(t *DumpRequestT) {
	t.Offset = rcv.Offset()
	t.Datalength = rcv.Datalength()
}

func (rcv *DumpRequest) UnPack() *DumpRequestT {
	if rcv == nil { return nil }
	t := &DumpRequestT{}
	rcv.UnPackTo(t)
	return t
}

type DumpRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsDumpRequest(buf []byte, offset flatbuffers.UOffsetT) *DumpRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DumpRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDumpRequest(buf []byte, offset flatbuffers.UOffsetT) *DumpRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DumpRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DumpRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DumpRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DumpRequest) Offset() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DumpRequest) MutateOffset(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *DumpRequest) Datalength() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DumpRequest) MutateDatalength(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func DumpRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func DumpRequestAddOffset(builder *flatbuffers.Builder, offset uint32) {
	builder.PrependUint32Slot(0, offset, 0)
}
func DumpRequestAddDatalength(builder *flatbuffers.Builder, datalength uint32) {
	builder.PrependUint32Slot(1, datalength, 0)
}
func DumpRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
