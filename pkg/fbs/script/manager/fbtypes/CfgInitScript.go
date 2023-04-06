// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CfgInitScriptT struct {
	File string
	Language string
	Parameter []string
}

func (t *CfgInitScriptT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	fileOffset := builder.CreateString(t.File)
	languageOffset := builder.CreateString(t.Language)
	parameterOffset := flatbuffers.UOffsetT(0)
	if t.Parameter != nil {
		parameterLength := len(t.Parameter)
		parameterOffsets := make([]flatbuffers.UOffsetT, parameterLength)
		for j := 0; j < parameterLength; j++ {
			parameterOffsets[j] = builder.CreateString(t.Parameter[j])
		}
		CfgInitScriptStartParameterVector(builder, parameterLength)
		for j := parameterLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(parameterOffsets[j])
		}
		parameterOffset = builder.EndVector(parameterLength)
	}
	CfgInitScriptStart(builder)
	CfgInitScriptAddFile(builder, fileOffset)
	CfgInitScriptAddLanguage(builder, languageOffset)
	CfgInitScriptAddParameter(builder, parameterOffset)
	return CfgInitScriptEnd(builder)
}

func (rcv *CfgInitScript) UnPackTo(t *CfgInitScriptT) {
	t.File = string(rcv.File())
	t.Language = string(rcv.Language())
	parameterLength := rcv.ParameterLength()
	t.Parameter = make([]string, parameterLength)
	for j := 0; j < parameterLength; j++ {
		t.Parameter[j] = string(rcv.Parameter(j))
	}
}

func (rcv *CfgInitScript) UnPack() *CfgInitScriptT {
	if rcv == nil { return nil }
	t := &CfgInitScriptT{}
	rcv.UnPackTo(t)
	return t
}

type CfgInitScript struct {
	_tab flatbuffers.Table
}

func GetRootAsCfgInitScript(buf []byte, offset flatbuffers.UOffsetT) *CfgInitScript {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CfgInitScript{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCfgInitScript(buf []byte, offset flatbuffers.UOffsetT) *CfgInitScript {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CfgInitScript{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CfgInitScript) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CfgInitScript) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CfgInitScript) File() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CfgInitScript) Language() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CfgInitScript) Parameter(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *CfgInitScript) ParameterLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func CfgInitScriptStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func CfgInitScriptAddFile(builder *flatbuffers.Builder, file flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(file), 0)
}
func CfgInitScriptAddLanguage(builder *flatbuffers.Builder, language flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(language), 0)
}
func CfgInitScriptAddParameter(builder *flatbuffers.Builder, parameter flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(parameter), 0)
}
func CfgInitScriptStartParameterVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CfgInitScriptEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
