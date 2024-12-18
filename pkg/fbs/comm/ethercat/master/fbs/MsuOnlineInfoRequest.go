// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Master syc unit online request
type MsuOnlineInfoRequestT struct {
	Id uint16 `json:"id"`
}

func (t *MsuOnlineInfoRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	MsuOnlineInfoRequestStart(builder)
	MsuOnlineInfoRequestAddId(builder, t.Id)
	return MsuOnlineInfoRequestEnd(builder)
}

func (rcv *MsuOnlineInfoRequest) UnPackTo(t *MsuOnlineInfoRequestT) {
	t.Id = rcv.Id()
}

func (rcv *MsuOnlineInfoRequest) UnPack() *MsuOnlineInfoRequestT {
	if rcv == nil { return nil }
	t := &MsuOnlineInfoRequestT{}
	rcv.UnPackTo(t)
	return t
}

type MsuOnlineInfoRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsMsuOnlineInfoRequest(buf []byte, offset flatbuffers.UOffsetT) *MsuOnlineInfoRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MsuOnlineInfoRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMsuOnlineInfoRequest(buf []byte, offset flatbuffers.UOffsetT) *MsuOnlineInfoRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MsuOnlineInfoRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MsuOnlineInfoRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MsuOnlineInfoRequest) Table() flatbuffers.Table {
	return rcv._tab
}

///Msu Id
func (rcv *MsuOnlineInfoRequest) Id() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Msu Id
func (rcv *MsuOnlineInfoRequest) MutateId(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func MsuOnlineInfoRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MsuOnlineInfoRequestAddId(builder *flatbuffers.Builder, id uint16) {
	builder.PrependUint16Slot(0, id, 0)
}
func MsuOnlineInfoRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
