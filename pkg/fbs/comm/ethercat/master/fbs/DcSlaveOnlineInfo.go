// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///This node delivers slave info data of a configured slave from the EtherCAT Master
type DcSlaveOnlineInfoT struct {
	Request *AddressedRequestT `json:"request"`
	Response *DcSlaveOnlineInfoResponseT `json:"response"`
}

func (t *DcSlaveOnlineInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	requestOffset := t.Request.Pack(builder)
	responseOffset := t.Response.Pack(builder)
	DcSlaveOnlineInfoStart(builder)
	DcSlaveOnlineInfoAddRequest(builder, requestOffset)
	DcSlaveOnlineInfoAddResponse(builder, responseOffset)
	return DcSlaveOnlineInfoEnd(builder)
}

func (rcv *DcSlaveOnlineInfo) UnPackTo(t *DcSlaveOnlineInfoT) {
	t.Request = rcv.Request(nil).UnPack()
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *DcSlaveOnlineInfo) UnPack() *DcSlaveOnlineInfoT {
	if rcv == nil { return nil }
	t := &DcSlaveOnlineInfoT{}
	rcv.UnPackTo(t)
	return t
}

type DcSlaveOnlineInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsDcSlaveOnlineInfo(buf []byte, offset flatbuffers.UOffsetT) *DcSlaveOnlineInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DcSlaveOnlineInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDcSlaveOnlineInfo(buf []byte, offset flatbuffers.UOffsetT) *DcSlaveOnlineInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DcSlaveOnlineInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DcSlaveOnlineInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DcSlaveOnlineInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DcSlaveOnlineInfo) Request(obj *AddressedRequest) *AddressedRequest {
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

func (rcv *DcSlaveOnlineInfo) Response(obj *DcSlaveOnlineInfoResponse) *DcSlaveOnlineInfoResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(DcSlaveOnlineInfoResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func DcSlaveOnlineInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func DcSlaveOnlineInfoAddRequest(builder *flatbuffers.Builder, request flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(request), 0)
}
func DcSlaveOnlineInfoAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(response), 0)
}
func DcSlaveOnlineInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
