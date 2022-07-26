// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SlaveOnlineInfoT struct {
	Request *AddressedRequestT
	Response *SlaveOnlineInfoResponseT
}

func (t *SlaveOnlineInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	requestOffset := t.Request.Pack(builder)
	responseOffset := t.Response.Pack(builder)
	SlaveOnlineInfoStart(builder)
	SlaveOnlineInfoAddRequest(builder, requestOffset)
	SlaveOnlineInfoAddResponse(builder, responseOffset)
	return SlaveOnlineInfoEnd(builder)
}

func (rcv *SlaveOnlineInfo) UnPackTo(t *SlaveOnlineInfoT) {
	t.Request = rcv.Request(nil).UnPack()
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *SlaveOnlineInfo) UnPack() *SlaveOnlineInfoT {
	if rcv == nil { return nil }
	t := &SlaveOnlineInfoT{}
	rcv.UnPackTo(t)
	return t
}

type SlaveOnlineInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsSlaveOnlineInfo(buf []byte, offset flatbuffers.UOffsetT) *SlaveOnlineInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SlaveOnlineInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSlaveOnlineInfo(buf []byte, offset flatbuffers.UOffsetT) *SlaveOnlineInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SlaveOnlineInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SlaveOnlineInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SlaveOnlineInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SlaveOnlineInfo) Request(obj *AddressedRequest) *AddressedRequest {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AddressedRequest)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SlaveOnlineInfo) Response(obj *SlaveOnlineInfoResponse) *SlaveOnlineInfoResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SlaveOnlineInfoResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func SlaveOnlineInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SlaveOnlineInfoAddRequest(builder *flatbuffers.Builder, request flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(request), 0)
}
func SlaveOnlineInfoAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(response), 0)
}
func SlaveOnlineInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
