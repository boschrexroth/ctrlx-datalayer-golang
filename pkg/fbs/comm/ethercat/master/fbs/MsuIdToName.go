// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Thsi node converts master sync unit id to name.
type MsuIdToNameT struct {
	Request *MsuIdToNameRequestT `json:"request"`
	Response *MsuIdToNameResponseT `json:"response"`
}

func (t *MsuIdToNameT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	requestOffset := t.Request.Pack(builder)
	responseOffset := t.Response.Pack(builder)
	MsuIdToNameStart(builder)
	MsuIdToNameAddRequest(builder, requestOffset)
	MsuIdToNameAddResponse(builder, responseOffset)
	return MsuIdToNameEnd(builder)
}

func (rcv *MsuIdToName) UnPackTo(t *MsuIdToNameT) {
	t.Request = rcv.Request(nil).UnPack()
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *MsuIdToName) UnPack() *MsuIdToNameT {
	if rcv == nil { return nil }
	t := &MsuIdToNameT{}
	rcv.UnPackTo(t)
	return t
}

type MsuIdToName struct {
	_tab flatbuffers.Table
}

func GetRootAsMsuIdToName(buf []byte, offset flatbuffers.UOffsetT) *MsuIdToName {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MsuIdToName{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMsuIdToName(buf []byte, offset flatbuffers.UOffsetT) *MsuIdToName {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MsuIdToName{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MsuIdToName) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MsuIdToName) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MsuIdToName) Request(obj *MsuIdToNameRequest) *MsuIdToNameRequest {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MsuIdToNameRequest)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *MsuIdToName) Response(obj *MsuIdToNameResponse) *MsuIdToNameResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MsuIdToNameResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func MsuIdToNameStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MsuIdToNameAddRequest(builder *flatbuffers.Builder, request flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(request), 0)
}
func MsuIdToNameAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(response), 0)
}
func MsuIdToNameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
