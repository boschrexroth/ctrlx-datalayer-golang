// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LinkLayerOnlineInfoT struct {
	Response *LinkLayerOnlineInfoResponseT
}

func (t *LinkLayerOnlineInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	responseOffset := t.Response.Pack(builder)
	LinkLayerOnlineInfoStart(builder)
	LinkLayerOnlineInfoAddResponse(builder, responseOffset)
	return LinkLayerOnlineInfoEnd(builder)
}

func (rcv *LinkLayerOnlineInfo) UnPackTo(t *LinkLayerOnlineInfoT) {
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *LinkLayerOnlineInfo) UnPack() *LinkLayerOnlineInfoT {
	if rcv == nil { return nil }
	t := &LinkLayerOnlineInfoT{}
	rcv.UnPackTo(t)
	return t
}

type LinkLayerOnlineInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsLinkLayerOnlineInfo(buf []byte, offset flatbuffers.UOffsetT) *LinkLayerOnlineInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LinkLayerOnlineInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLinkLayerOnlineInfo(buf []byte, offset flatbuffers.UOffsetT) *LinkLayerOnlineInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LinkLayerOnlineInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LinkLayerOnlineInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LinkLayerOnlineInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LinkLayerOnlineInfo) Response(obj *LinkLayerOnlineInfoResponse) *LinkLayerOnlineInfoResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(LinkLayerOnlineInfoResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func LinkLayerOnlineInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func LinkLayerOnlineInfoAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(response), 0)
}
func LinkLayerOnlineInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
