// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SoeVariableInfoT struct {
	ConfigurationListIdn uint16 `json:"configurationListIdn"`
	Idn uint16 `json:"idn"`
}

func (t *SoeVariableInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	SoeVariableInfoStart(builder)
	SoeVariableInfoAddConfigurationListIdn(builder, t.ConfigurationListIdn)
	SoeVariableInfoAddIdn(builder, t.Idn)
	return SoeVariableInfoEnd(builder)
}

func (rcv *SoeVariableInfo) UnPackTo(t *SoeVariableInfoT) {
	t.ConfigurationListIdn = rcv.ConfigurationListIdn()
	t.Idn = rcv.Idn()
}

func (rcv *SoeVariableInfo) UnPack() *SoeVariableInfoT {
	if rcv == nil { return nil }
	t := &SoeVariableInfoT{}
	rcv.UnPackTo(t)
	return t
}

type SoeVariableInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsSoeVariableInfo(buf []byte, offset flatbuffers.UOffsetT) *SoeVariableInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SoeVariableInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSoeVariableInfo(buf []byte, offset flatbuffers.UOffsetT) *SoeVariableInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SoeVariableInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SoeVariableInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SoeVariableInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SoeVariableInfo) ConfigurationListIdn() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SoeVariableInfo) MutateConfigurationListIdn(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func (rcv *SoeVariableInfo) Idn() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SoeVariableInfo) MutateIdn(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

func SoeVariableInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SoeVariableInfoAddConfigurationListIdn(builder *flatbuffers.Builder, configurationListIdn uint16) {
	builder.PrependUint16Slot(0, configurationListIdn, 0)
}
func SoeVariableInfoAddIdn(builder *flatbuffers.Builder, idn uint16) {
	builder.PrependUint16Slot(1, idn, 0)
}
func SoeVariableInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}