// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Fieldbus configuration in binary format
type FieldbusConfigResponseT struct {
	Ini []byte `json:"ini"`
	Eni []byte `json:"eni"`
}

func (t *FieldbusConfigResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	iniOffset := flatbuffers.UOffsetT(0)
	if t.Ini != nil {
		iniOffset = builder.CreateByteString(t.Ini)
	}
	eniOffset := flatbuffers.UOffsetT(0)
	if t.Eni != nil {
		eniOffset = builder.CreateByteString(t.Eni)
	}
	FieldbusConfigResponseStart(builder)
	FieldbusConfigResponseAddIni(builder, iniOffset)
	FieldbusConfigResponseAddEni(builder, eniOffset)
	return FieldbusConfigResponseEnd(builder)
}

func (rcv *FieldbusConfigResponse) UnPackTo(t *FieldbusConfigResponseT) {
	t.Ini = rcv.IniBytes()
	t.Eni = rcv.EniBytes()
}

func (rcv *FieldbusConfigResponse) UnPack() *FieldbusConfigResponseT {
	if rcv == nil { return nil }
	t := &FieldbusConfigResponseT{}
	rcv.UnPackTo(t)
	return t
}

type FieldbusConfigResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsFieldbusConfigResponse(buf []byte, offset flatbuffers.UOffsetT) *FieldbusConfigResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FieldbusConfigResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFieldbusConfigResponse(buf []byte, offset flatbuffers.UOffsetT) *FieldbusConfigResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FieldbusConfigResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FieldbusConfigResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FieldbusConfigResponse) Table() flatbuffers.Table {
	return rcv._tab
}

///Master initialization in binary format
func (rcv *FieldbusConfigResponse) Ini(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *FieldbusConfigResponse) IniLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FieldbusConfigResponse) IniBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///Master initialization in binary format
func (rcv *FieldbusConfigResponse) MutateIni(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

///EtherCAT network information in binary format
func (rcv *FieldbusConfigResponse) Eni(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *FieldbusConfigResponse) EniLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FieldbusConfigResponse) EniBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///EtherCAT network information in binary format
func (rcv *FieldbusConfigResponse) MutateEni(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func FieldbusConfigResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func FieldbusConfigResponseAddIni(builder *flatbuffers.Builder, ini flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ini), 0)
}
func FieldbusConfigResponseStartIniVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FieldbusConfigResponseAddEni(builder *flatbuffers.Builder, eni flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(eni), 0)
}
func FieldbusConfigResponseStartEniVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FieldbusConfigResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
