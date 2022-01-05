// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DiagInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsDiagInfo(buf []byte, offset flatbuffers.UOffsetT) *DiagInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DiagInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDiagInfo(buf []byte, offset flatbuffers.UOffsetT) *DiagInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DiagInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DiagInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DiagInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DiagInfo) LastMainDiag() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DiagInfo) MutateLastMainDiag(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *DiagInfo) LastDetailDiag() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DiagInfo) MutateLastDetailDiag(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *DiagInfo) LastErrText() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DiagInfo) LastErrTrace(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *DiagInfo) LastErrTraceLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DiagInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DiagInfoAddLastMainDiag(builder *flatbuffers.Builder, lastMainDiag uint32) {
	builder.PrependUint32Slot(0, lastMainDiag, 0)
}
func DiagInfoAddLastDetailDiag(builder *flatbuffers.Builder, lastDetailDiag uint32) {
	builder.PrependUint32Slot(1, lastDetailDiag, 0)
}
func DiagInfoAddLastErrText(builder *flatbuffers.Builder, lastErrText flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(lastErrText), 0)
}
func DiagInfoAddLastErrTrace(builder *flatbuffers.Builder, lastErrTrace flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(lastErrTrace), 0)
}
func DiagInfoStartLastErrTraceVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DiagInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
