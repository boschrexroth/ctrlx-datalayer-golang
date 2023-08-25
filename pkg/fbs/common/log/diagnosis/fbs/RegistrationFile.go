// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// This table defines the structure of the registration json file that is needed to register diagnostics.
type RegistrationFileT struct {
	Language string `json:"language"`
	Product string `json:"product"`
	Component string `json:"component"`
	MainDiagnostics []*MainDiagnosticT `json:"mainDiagnostics"`
}

func (t *RegistrationFileT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	languageOffset := flatbuffers.UOffsetT(0)
	if t.Language != "" {
		languageOffset = builder.CreateString(t.Language)
	}
	productOffset := flatbuffers.UOffsetT(0)
	if t.Product != "" {
		productOffset = builder.CreateString(t.Product)
	}
	componentOffset := flatbuffers.UOffsetT(0)
	if t.Component != "" {
		componentOffset = builder.CreateString(t.Component)
	}
	mainDiagnosticsOffset := flatbuffers.UOffsetT(0)
	if t.MainDiagnostics != nil {
		mainDiagnosticsLength := len(t.MainDiagnostics)
		mainDiagnosticsOffsets := make([]flatbuffers.UOffsetT, mainDiagnosticsLength)
		for j := 0; j < mainDiagnosticsLength; j++ {
			mainDiagnosticsOffsets[j] = t.MainDiagnostics[j].Pack(builder)
		}
		RegistrationFileStartMainDiagnosticsVector(builder, mainDiagnosticsLength)
		for j := mainDiagnosticsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(mainDiagnosticsOffsets[j])
		}
		mainDiagnosticsOffset = builder.EndVector(mainDiagnosticsLength)
	}
	RegistrationFileStart(builder)
	RegistrationFileAddLanguage(builder, languageOffset)
	RegistrationFileAddProduct(builder, productOffset)
	RegistrationFileAddComponent(builder, componentOffset)
	RegistrationFileAddMainDiagnostics(builder, mainDiagnosticsOffset)
	return RegistrationFileEnd(builder)
}

func (rcv *RegistrationFile) UnPackTo(t *RegistrationFileT) {
	t.Language = string(rcv.Language())
	t.Product = string(rcv.Product())
	t.Component = string(rcv.Component())
	mainDiagnosticsLength := rcv.MainDiagnosticsLength()
	t.MainDiagnostics = make([]*MainDiagnosticT, mainDiagnosticsLength)
	for j := 0; j < mainDiagnosticsLength; j++ {
		x := MainDiagnostic{}
		rcv.MainDiagnostics(&x, j)
		t.MainDiagnostics[j] = x.UnPack()
	}
}

func (rcv *RegistrationFile) UnPack() *RegistrationFileT {
	if rcv == nil { return nil }
	t := &RegistrationFileT{}
	rcv.UnPackTo(t)
	return t
}

type RegistrationFile struct {
	_tab flatbuffers.Table
}

func GetRootAsRegistrationFile(buf []byte, offset flatbuffers.UOffsetT) *RegistrationFile {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RegistrationFile{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRegistrationFile(buf []byte, offset flatbuffers.UOffsetT) *RegistrationFile {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RegistrationFile{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RegistrationFile) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RegistrationFile) Table() flatbuffers.Table {
	return rcv._tab
}

/// Language of the default texts (optional).
func (rcv *RegistrationFile) Language() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Language of the default texts (optional).
/// Product associated to the diagnostics (optional).
func (rcv *RegistrationFile) Product() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Product associated to the diagnostics (optional).
/// Component associated to the diagnostics (optional).
func (rcv *RegistrationFile) Component() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Component associated to the diagnostics (optional).
/// All main diagnostics (including their related detailed diagnostics) that should be registered.
func (rcv *RegistrationFile) MainDiagnostics(obj *MainDiagnostic, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *RegistrationFile) MainDiagnosticsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// All main diagnostics (including their related detailed diagnostics) that should be registered.
func RegistrationFileStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func RegistrationFileAddLanguage(builder *flatbuffers.Builder, language flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(language), 0)
}
func RegistrationFileAddProduct(builder *flatbuffers.Builder, product flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(product), 0)
}
func RegistrationFileAddComponent(builder *flatbuffers.Builder, component flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(component), 0)
}
func RegistrationFileAddMainDiagnostics(builder *flatbuffers.Builder, mainDiagnostics flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(mainDiagnostics), 0)
}
func RegistrationFileStartMainDiagnosticsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func RegistrationFileEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
