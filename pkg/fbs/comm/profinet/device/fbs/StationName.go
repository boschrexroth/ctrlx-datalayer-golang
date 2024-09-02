// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StationNameT struct {
	Request *StationNameRequestT `json:"request"`
	Response *StationNameResponseT `json:"response"`
}

func (t *StationNameT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	requestOffset := t.Request.Pack(builder)
	responseOffset := t.Response.Pack(builder)
	StationNameStart(builder)
	StationNameAddRequest(builder, requestOffset)
	StationNameAddResponse(builder, responseOffset)
	return StationNameEnd(builder)
}

func (rcv *StationName) UnPackTo(t *StationNameT) {
	t.Request = rcv.Request(nil).UnPack()
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *StationName) UnPack() *StationNameT {
	if rcv == nil { return nil }
	t := &StationNameT{}
	rcv.UnPackTo(t)
	return t
}

type StationName struct {
	_tab flatbuffers.Table
}

func GetRootAsStationName(buf []byte, offset flatbuffers.UOffsetT) *StationName {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StationName{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStationName(buf []byte, offset flatbuffers.UOffsetT) *StationName {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StationName{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StationName) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StationName) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StationName) Request(obj *StationNameRequest) *StationNameRequest {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(StationNameRequest)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *StationName) Response(obj *StationNameResponse) *StationNameResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(StationNameResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func StationNameStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func StationNameAddRequest(builder *flatbuffers.Builder, request flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(request), 0)
}
func StationNameAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(response), 0)
}
func StationNameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}