// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// single result of the validation of a calculation pipeline
type SinglePipelineValidationT struct {
	MainDiag uint32
	DetailDiag uint32
	Uri string
	AddInfo string
}

func (t *SinglePipelineValidationT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	uriOffset := builder.CreateString(t.Uri)
	addInfoOffset := builder.CreateString(t.AddInfo)
	SinglePipelineValidationStart(builder)
	SinglePipelineValidationAddMainDiag(builder, t.MainDiag)
	SinglePipelineValidationAddDetailDiag(builder, t.DetailDiag)
	SinglePipelineValidationAddUri(builder, uriOffset)
	SinglePipelineValidationAddAddInfo(builder, addInfoOffset)
	return SinglePipelineValidationEnd(builder)
}

func (rcv *SinglePipelineValidation) UnPackTo(t *SinglePipelineValidationT) {
	t.MainDiag = rcv.MainDiag()
	t.DetailDiag = rcv.DetailDiag()
	t.Uri = string(rcv.Uri())
	t.AddInfo = string(rcv.AddInfo())
}

func (rcv *SinglePipelineValidation) UnPack() *SinglePipelineValidationT {
	if rcv == nil { return nil }
	t := &SinglePipelineValidationT{}
	rcv.UnPackTo(t)
	return t
}

type SinglePipelineValidation struct {
	_tab flatbuffers.Table
}

func GetRootAsSinglePipelineValidation(buf []byte, offset flatbuffers.UOffsetT) *SinglePipelineValidation {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SinglePipelineValidation{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSinglePipelineValidation(buf []byte, offset flatbuffers.UOffsetT) *SinglePipelineValidation {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SinglePipelineValidation{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SinglePipelineValidation) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SinglePipelineValidation) Table() flatbuffers.Table {
	return rcv._tab
}

/// main diagnosis code
func (rcv *SinglePipelineValidation) MainDiag() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// main diagnosis code
func (rcv *SinglePipelineValidation) MutateMainDiag(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

/// detail diagnosis code
func (rcv *SinglePipelineValidation) DetailDiag() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// detail diagnosis code
func (rcv *SinglePipelineValidation) MutateDetailDiag(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

/// URI of the faulty instance
func (rcv *SinglePipelineValidation) Uri() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// URI of the faulty instance
/// additional infomation
func (rcv *SinglePipelineValidation) AddInfo() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// additional infomation
func SinglePipelineValidationStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func SinglePipelineValidationAddMainDiag(builder *flatbuffers.Builder, mainDiag uint32) {
	builder.PrependUint32Slot(0, mainDiag, 0)
}
func SinglePipelineValidationAddDetailDiag(builder *flatbuffers.Builder, detailDiag uint32) {
	builder.PrependUint32Slot(1, detailDiag, 0)
}
func SinglePipelineValidationAddUri(builder *flatbuffers.Builder, uri flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(uri), 0)
}
func SinglePipelineValidationAddAddInfo(builder *flatbuffers.Builder, addInfo flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(addInfo), 0)
}
func SinglePipelineValidationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
