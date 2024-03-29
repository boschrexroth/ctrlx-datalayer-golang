// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package importParameter

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ImportParameterT struct {
	FileName string `json:"fileName"`
	SetIndex int8 `json:"setIndex"`
}

func (t *ImportParameterT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	fileNameOffset := flatbuffers.UOffsetT(0)
	if t.FileName != "" {
		fileNameOffset = builder.CreateString(t.FileName)
	}
	ImportParameterStart(builder)
	ImportParameterAddFileName(builder, fileNameOffset)
	ImportParameterAddSetIndex(builder, t.SetIndex)
	return ImportParameterEnd(builder)
}

func (rcv *ImportParameter) UnPackTo(t *ImportParameterT) {
	t.FileName = string(rcv.FileName())
	t.SetIndex = rcv.SetIndex()
}

func (rcv *ImportParameter) UnPack() *ImportParameterT {
	if rcv == nil { return nil }
	t := &ImportParameterT{}
	rcv.UnPackTo(t)
	return t
}

type ImportParameter struct {
	_tab flatbuffers.Table
}

func GetRootAsImportParameter(buf []byte, offset flatbuffers.UOffsetT) *ImportParameter {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ImportParameter{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsImportParameter(buf []byte, offset flatbuffers.UOffsetT) *ImportParameter {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ImportParameter{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ImportParameter) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ImportParameter) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ImportParameter) FileName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ImportParameter) SetIndex() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ImportParameter) MutateSetIndex(n int8) bool {
	return rcv._tab.MutateInt8Slot(6, n)
}

func ImportParameterStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ImportParameterAddFileName(builder *flatbuffers.Builder, fileName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(fileName), 0)
}
func ImportParameterAddSetIndex(builder *flatbuffers.Builder, setIndex int8) {
	builder.PrependInt8Slot(1, setIndex, 0)
}
func ImportParameterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
