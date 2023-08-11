// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// This table defines the elements of a main or detailed diagnostics that should be registered.
type DiagnosisElementT struct {
	DiagnosisNumber uint32 `json:"diagnosisNumber"`
	Version byte `json:"version"`
	TextEnglish string `json:"textEnglish"`
}

func (t *DiagnosisElementT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	textEnglishOffset := flatbuffers.UOffsetT(0)
	if t.TextEnglish != "" {
		textEnglishOffset = builder.CreateString(t.TextEnglish)
	}
	DiagnosisElementStart(builder)
	DiagnosisElementAddDiagnosisNumber(builder, t.DiagnosisNumber)
	DiagnosisElementAddVersion(builder, t.Version)
	DiagnosisElementAddTextEnglish(builder, textEnglishOffset)
	return DiagnosisElementEnd(builder)
}

func (rcv *DiagnosisElement) UnPackTo(t *DiagnosisElementT) {
	t.DiagnosisNumber = rcv.DiagnosisNumber()
	t.Version = rcv.Version()
	t.TextEnglish = string(rcv.TextEnglish())
}

func (rcv *DiagnosisElement) UnPack() *DiagnosisElementT {
	if rcv == nil { return nil }
	t := &DiagnosisElementT{}
	rcv.UnPackTo(t)
	return t
}

type DiagnosisElement struct {
	_tab flatbuffers.Table
}

func GetRootAsDiagnosisElement(buf []byte, offset flatbuffers.UOffsetT) *DiagnosisElement {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DiagnosisElement{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDiagnosisElement(buf []byte, offset flatbuffers.UOffsetT) *DiagnosisElement {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DiagnosisElement{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DiagnosisElement) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DiagnosisElement) Table() flatbuffers.Table {
	return rcv._tab
}

/// Main or detailed diagnostic number.
func (rcv *DiagnosisElement) DiagnosisNumber() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// Main or detailed diagnostic number.
func (rcv *DiagnosisElement) MutateDiagnosisNumber(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

/// Version of the diagnostics starting with 1.
func (rcv *DiagnosisElement) Version() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

/// Version of the diagnostics starting with 1.
func (rcv *DiagnosisElement) MutateVersion(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

/// English text of the diagnostics.
func (rcv *DiagnosisElement) TextEnglish() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// English text of the diagnostics.
func DiagnosisElementStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func DiagnosisElementAddDiagnosisNumber(builder *flatbuffers.Builder, diagnosisNumber uint32) {
	builder.PrependUint32Slot(0, diagnosisNumber, 0)
}
func DiagnosisElementAddVersion(builder *flatbuffers.Builder, version byte) {
	builder.PrependByteSlot(1, version, 0)
}
func DiagnosisElementAddTextEnglish(builder *flatbuffers.Builder, textEnglish flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(textEnglish), 0)
}
func DiagnosisElementEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
