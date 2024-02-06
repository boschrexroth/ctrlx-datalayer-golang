// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///This node returns the current online DC (Distributed Clock) information
type DcOnlineInfoT struct {
	Response *DcOnlineInfoResponseT `json:"response"`
}

func (t *DcOnlineInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	responseOffset := t.Response.Pack(builder)
	DcOnlineInfoStart(builder)
	DcOnlineInfoAddResponse(builder, responseOffset)
	return DcOnlineInfoEnd(builder)
}

func (rcv *DcOnlineInfo) UnPackTo(t *DcOnlineInfoT) {
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *DcOnlineInfo) UnPack() *DcOnlineInfoT {
	if rcv == nil { return nil }
	t := &DcOnlineInfoT{}
	rcv.UnPackTo(t)
	return t
}

type DcOnlineInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsDcOnlineInfo(buf []byte, offset flatbuffers.UOffsetT) *DcOnlineInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DcOnlineInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDcOnlineInfo(buf []byte, offset flatbuffers.UOffsetT) *DcOnlineInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DcOnlineInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DcOnlineInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DcOnlineInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DcOnlineInfo) Response(obj *DcOnlineInfoResponse) *DcOnlineInfoResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(DcOnlineInfoResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func DcOnlineInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DcOnlineInfoAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(response), 0)
}
func DcOnlineInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
