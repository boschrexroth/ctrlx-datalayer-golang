// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type IpSettingsT struct {
	Request *IpSettingsRequestT `json:"request"`
	Response *IpSettingsResponseT `json:"response"`
}

func (t *IpSettingsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	requestOffset := t.Request.Pack(builder)
	responseOffset := t.Response.Pack(builder)
	IpSettingsStart(builder)
	IpSettingsAddRequest(builder, requestOffset)
	IpSettingsAddResponse(builder, responseOffset)
	return IpSettingsEnd(builder)
}

func (rcv *IpSettings) UnPackTo(t *IpSettingsT) {
	t.Request = rcv.Request(nil).UnPack()
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *IpSettings) UnPack() *IpSettingsT {
	if rcv == nil { return nil }
	t := &IpSettingsT{}
	rcv.UnPackTo(t)
	return t
}

type IpSettings struct {
	_tab flatbuffers.Table
}

func GetRootAsIpSettings(buf []byte, offset flatbuffers.UOffsetT) *IpSettings {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IpSettings{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsIpSettings(buf []byte, offset flatbuffers.UOffsetT) *IpSettings {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &IpSettings{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *IpSettings) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IpSettings) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *IpSettings) Request(obj *IpSettingsRequest) *IpSettingsRequest {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(IpSettingsRequest)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *IpSettings) Response(obj *IpSettingsResponse) *IpSettingsResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(IpSettingsResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func IpSettingsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func IpSettingsAddRequest(builder *flatbuffers.Builder, request flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(request), 0)
}
func IpSettingsAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(response), 0)
}
func IpSettingsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}