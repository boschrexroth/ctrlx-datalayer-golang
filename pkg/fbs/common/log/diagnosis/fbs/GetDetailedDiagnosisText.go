// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// This table defines the required elements to request the diagnostic text of a diagnostic log.
type GetDetailedDiagnosisTextT struct {
	DetailedDiagnosisNumber string `json:"detailedDiagnosisNumber"`
	RelatedMainDiagnosisNumber string `json:"relatedMainDiagnosisNumber"`
}

func (t *GetDetailedDiagnosisTextT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	detailedDiagnosisNumberOffset := flatbuffers.UOffsetT(0)
	if t.DetailedDiagnosisNumber != "" {
		detailedDiagnosisNumberOffset = builder.CreateString(t.DetailedDiagnosisNumber)
	}
	relatedMainDiagnosisNumberOffset := flatbuffers.UOffsetT(0)
	if t.RelatedMainDiagnosisNumber != "" {
		relatedMainDiagnosisNumberOffset = builder.CreateString(t.RelatedMainDiagnosisNumber)
	}
	GetDetailedDiagnosisTextStart(builder)
	GetDetailedDiagnosisTextAddDetailedDiagnosisNumber(builder, detailedDiagnosisNumberOffset)
	GetDetailedDiagnosisTextAddRelatedMainDiagnosisNumber(builder, relatedMainDiagnosisNumberOffset)
	return GetDetailedDiagnosisTextEnd(builder)
}

func (rcv *GetDetailedDiagnosisText) UnPackTo(t *GetDetailedDiagnosisTextT) {
	t.DetailedDiagnosisNumber = string(rcv.DetailedDiagnosisNumber())
	t.RelatedMainDiagnosisNumber = string(rcv.RelatedMainDiagnosisNumber())
}

func (rcv *GetDetailedDiagnosisText) UnPack() *GetDetailedDiagnosisTextT {
	if rcv == nil { return nil }
	t := &GetDetailedDiagnosisTextT{}
	rcv.UnPackTo(t)
	return t
}

type GetDetailedDiagnosisText struct {
	_tab flatbuffers.Table
}

func GetRootAsGetDetailedDiagnosisText(buf []byte, offset flatbuffers.UOffsetT) *GetDetailedDiagnosisText {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GetDetailedDiagnosisText{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsGetDetailedDiagnosisText(buf []byte, offset flatbuffers.UOffsetT) *GetDetailedDiagnosisText {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &GetDetailedDiagnosisText{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *GetDetailedDiagnosisText) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GetDetailedDiagnosisText) Table() flatbuffers.Table {
	return rcv._tab
}

/// Detailed diagnostic number of the diagnostic log.
func (rcv *GetDetailedDiagnosisText) DetailedDiagnosisNumber() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Detailed diagnostic number of the diagnostic log.
/// Related main diagnostic number of the diagnostic log.
func (rcv *GetDetailedDiagnosisText) RelatedMainDiagnosisNumber() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Related main diagnostic number of the diagnostic log.
func GetDetailedDiagnosisTextStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func GetDetailedDiagnosisTextAddDetailedDiagnosisNumber(builder *flatbuffers.Builder, detailedDiagnosisNumber flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(detailedDiagnosisNumber), 0)
}
func GetDetailedDiagnosisTextAddRelatedMainDiagnosisNumber(builder *flatbuffers.Builder, relatedMainDiagnosisNumber flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(relatedMainDiagnosisNumber), 0)
}
func GetDetailedDiagnosisTextEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
