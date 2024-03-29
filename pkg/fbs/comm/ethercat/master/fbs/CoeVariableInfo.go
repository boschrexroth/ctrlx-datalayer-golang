// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///CoE variable info
type CoeVariableInfoT struct {
	PdoIndex uint16 `json:"pdoIndex"`
	ObjectIndex uint16 `json:"objectIndex"`
	SubIndex byte `json:"subIndex"`
}

func (t *CoeVariableInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	CoeVariableInfoStart(builder)
	CoeVariableInfoAddPdoIndex(builder, t.PdoIndex)
	CoeVariableInfoAddObjectIndex(builder, t.ObjectIndex)
	CoeVariableInfoAddSubIndex(builder, t.SubIndex)
	return CoeVariableInfoEnd(builder)
}

func (rcv *CoeVariableInfo) UnPackTo(t *CoeVariableInfoT) {
	t.PdoIndex = rcv.PdoIndex()
	t.ObjectIndex = rcv.ObjectIndex()
	t.SubIndex = rcv.SubIndex()
}

func (rcv *CoeVariableInfo) UnPack() *CoeVariableInfoT {
	if rcv == nil { return nil }
	t := &CoeVariableInfoT{}
	rcv.UnPackTo(t)
	return t
}

type CoeVariableInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsCoeVariableInfo(buf []byte, offset flatbuffers.UOffsetT) *CoeVariableInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CoeVariableInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCoeVariableInfo(buf []byte, offset flatbuffers.UOffsetT) *CoeVariableInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CoeVariableInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CoeVariableInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CoeVariableInfo) Table() flatbuffers.Table {
	return rcv._tab
}

///PDO index, e.g. 0x1A00
func (rcv *CoeVariableInfo) PdoIndex() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///PDO index, e.g. 0x1A00
func (rcv *CoeVariableInfo) MutatePdoIndex(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

///Object index, e.g. 0x3000
func (rcv *CoeVariableInfo) ObjectIndex() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Object index, e.g. 0x3000
func (rcv *CoeVariableInfo) MutateObjectIndex(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

///Sub index, e.g. 0x01
func (rcv *CoeVariableInfo) SubIndex() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

///Sub index, e.g. 0x01
func (rcv *CoeVariableInfo) MutateSubIndex(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func CoeVariableInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func CoeVariableInfoAddPdoIndex(builder *flatbuffers.Builder, pdoIndex uint16) {
	builder.PrependUint16Slot(0, pdoIndex, 0)
}
func CoeVariableInfoAddObjectIndex(builder *flatbuffers.Builder, objectIndex uint16) {
	builder.PrependUint16Slot(1, objectIndex, 0)
}
func CoeVariableInfoAddSubIndex(builder *flatbuffers.Builder, subIndex byte) {
	builder.PrependByteSlot(2, subIndex, 0)
}
func CoeVariableInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
