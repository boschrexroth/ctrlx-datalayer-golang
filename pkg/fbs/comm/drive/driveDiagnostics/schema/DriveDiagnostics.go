// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package schema

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DriveDiagnosticsT struct {
	MainDiagnosticsNumber string `json:"main_diagnostics_number"`
	DetailedDiagnosticsNumber string `json:"detailed_diagnostics_number"`
	Language string `json:"language"`
	MainDescription string `json:"main_description"`
	DetailedDescription string `json:"detailed_description"`
}

func (t *DriveDiagnosticsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	mainDiagnosticsNumberOffset := flatbuffers.UOffsetT(0)
	if t.MainDiagnosticsNumber != "" {
		mainDiagnosticsNumberOffset = builder.CreateString(t.MainDiagnosticsNumber)
	}
	detailedDiagnosticsNumberOffset := flatbuffers.UOffsetT(0)
	if t.DetailedDiagnosticsNumber != "" {
		detailedDiagnosticsNumberOffset = builder.CreateString(t.DetailedDiagnosticsNumber)
	}
	languageOffset := flatbuffers.UOffsetT(0)
	if t.Language != "" {
		languageOffset = builder.CreateString(t.Language)
	}
	mainDescriptionOffset := flatbuffers.UOffsetT(0)
	if t.MainDescription != "" {
		mainDescriptionOffset = builder.CreateString(t.MainDescription)
	}
	detailedDescriptionOffset := flatbuffers.UOffsetT(0)
	if t.DetailedDescription != "" {
		detailedDescriptionOffset = builder.CreateString(t.DetailedDescription)
	}
	DriveDiagnosticsStart(builder)
	DriveDiagnosticsAddMainDiagnosticsNumber(builder, mainDiagnosticsNumberOffset)
	DriveDiagnosticsAddDetailedDiagnosticsNumber(builder, detailedDiagnosticsNumberOffset)
	DriveDiagnosticsAddLanguage(builder, languageOffset)
	DriveDiagnosticsAddMainDescription(builder, mainDescriptionOffset)
	DriveDiagnosticsAddDetailedDescription(builder, detailedDescriptionOffset)
	return DriveDiagnosticsEnd(builder)
}

func (rcv *DriveDiagnostics) UnPackTo(t *DriveDiagnosticsT) {
	t.MainDiagnosticsNumber = string(rcv.MainDiagnosticsNumber())
	t.DetailedDiagnosticsNumber = string(rcv.DetailedDiagnosticsNumber())
	t.Language = string(rcv.Language())
	t.MainDescription = string(rcv.MainDescription())
	t.DetailedDescription = string(rcv.DetailedDescription())
}

func (rcv *DriveDiagnostics) UnPack() *DriveDiagnosticsT {
	if rcv == nil { return nil }
	t := &DriveDiagnosticsT{}
	rcv.UnPackTo(t)
	return t
}

type DriveDiagnostics struct {
	_tab flatbuffers.Table
}

func GetRootAsDriveDiagnostics(buf []byte, offset flatbuffers.UOffsetT) *DriveDiagnostics {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DriveDiagnostics{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDriveDiagnostics(buf []byte, offset flatbuffers.UOffsetT) *DriveDiagnostics {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DriveDiagnostics{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DriveDiagnostics) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DriveDiagnostics) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DriveDiagnostics) MainDiagnosticsNumber() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DriveDiagnostics) DetailedDiagnosticsNumber() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DriveDiagnostics) Language() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DriveDiagnostics) MainDescription() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DriveDiagnostics) DetailedDescription() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DriveDiagnosticsStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func DriveDiagnosticsAddMainDiagnosticsNumber(builder *flatbuffers.Builder, mainDiagnosticsNumber flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(mainDiagnosticsNumber), 0)
}
func DriveDiagnosticsAddDetailedDiagnosticsNumber(builder *flatbuffers.Builder, detailedDiagnosticsNumber flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(detailedDiagnosticsNumber), 0)
}
func DriveDiagnosticsAddLanguage(builder *flatbuffers.Builder, language flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(language), 0)
}
func DriveDiagnosticsAddMainDescription(builder *flatbuffers.Builder, mainDescription flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(mainDescription), 0)
}
func DriveDiagnosticsAddDetailedDescription(builder *flatbuffers.Builder, detailedDescription flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(detailedDescription), 0)
}
func DriveDiagnosticsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
