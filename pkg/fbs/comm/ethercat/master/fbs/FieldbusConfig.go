// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///This node provide an interface to read and write the EtherCAT fieldbus configuration to the EtherCAT master
type FieldbusConfigT struct {
	Request *FieldbusConfigRequestT `json:"request"`
	Response *FieldbusConfigResponseT `json:"response"`
}

func (t *FieldbusConfigT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	requestOffset := t.Request.Pack(builder)
	responseOffset := t.Response.Pack(builder)
	FieldbusConfigStart(builder)
	FieldbusConfigAddRequest(builder, requestOffset)
	FieldbusConfigAddResponse(builder, responseOffset)
	return FieldbusConfigEnd(builder)
}

func (rcv *FieldbusConfig) UnPackTo(t *FieldbusConfigT) {
	t.Request = rcv.Request(nil).UnPack()
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *FieldbusConfig) UnPack() *FieldbusConfigT {
	if rcv == nil { return nil }
	t := &FieldbusConfigT{}
	rcv.UnPackTo(t)
	return t
}

type FieldbusConfig struct {
	_tab flatbuffers.Table
}

func GetRootAsFieldbusConfig(buf []byte, offset flatbuffers.UOffsetT) *FieldbusConfig {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FieldbusConfig{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFieldbusConfig(buf []byte, offset flatbuffers.UOffsetT) *FieldbusConfig {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FieldbusConfig{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FieldbusConfig) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FieldbusConfig) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FieldbusConfig) Request(obj *FieldbusConfigRequest) *FieldbusConfigRequest {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FieldbusConfigRequest)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *FieldbusConfig) Response(obj *FieldbusConfigResponse) *FieldbusConfigResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FieldbusConfigResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func FieldbusConfigStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func FieldbusConfigAddRequest(builder *flatbuffers.Builder, request flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(request), 0)
}
func FieldbusConfigAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(response), 0)
}
func FieldbusConfigEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
