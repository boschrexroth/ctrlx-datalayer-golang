// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Fieldbus configuration in binary format
type FieldbusConfigRequestT struct {
	Ini []byte `json:"ini"`
	Eni []byte `json:"eni"`
}

func (t *FieldbusConfigRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	iniOffset := flatbuffers.UOffsetT(0)
	if t.Ini != nil {
		iniOffset = builder.CreateByteString(t.Ini)
	}
	eniOffset := flatbuffers.UOffsetT(0)
	if t.Eni != nil {
		eniOffset = builder.CreateByteString(t.Eni)
	}
	FieldbusConfigRequestStart(builder)
	FieldbusConfigRequestAddIni(builder, iniOffset)
	FieldbusConfigRequestAddEni(builder, eniOffset)
	return FieldbusConfigRequestEnd(builder)
}

func (rcv *FieldbusConfigRequest) UnPackTo(t *FieldbusConfigRequestT) {
	t.Ini = rcv.IniBytes()
	t.Eni = rcv.EniBytes()
}

func (rcv *FieldbusConfigRequest) UnPack() *FieldbusConfigRequestT {
	if rcv == nil { return nil }
	t := &FieldbusConfigRequestT{}
	rcv.UnPackTo(t)
	return t
}

type FieldbusConfigRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsFieldbusConfigRequest(buf []byte, offset flatbuffers.UOffsetT) *FieldbusConfigRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FieldbusConfigRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFieldbusConfigRequest(buf []byte, offset flatbuffers.UOffsetT) *FieldbusConfigRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FieldbusConfigRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FieldbusConfigRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FieldbusConfigRequest) Table() flatbuffers.Table {
	return rcv._tab
}

///Master initialization in binary format
func (rcv *FieldbusConfigRequest) Ini(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *FieldbusConfigRequest) IniLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FieldbusConfigRequest) IniBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///Master initialization in binary format
func (rcv *FieldbusConfigRequest) MutateIni(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

///EtherCAT network information in binary format
func (rcv *FieldbusConfigRequest) Eni(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *FieldbusConfigRequest) EniLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FieldbusConfigRequest) EniBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///EtherCAT network information in binary format
func (rcv *FieldbusConfigRequest) MutateEni(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func FieldbusConfigRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func FieldbusConfigRequestAddIni(builder *flatbuffers.Builder, ini flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ini), 0)
}
func FieldbusConfigRequestStartIniVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FieldbusConfigRequestAddEni(builder *flatbuffers.Builder, eni flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(eni), 0)
}
func FieldbusConfigRequestStartEniVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FieldbusConfigRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
