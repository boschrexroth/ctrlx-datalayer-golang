// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MasterInfoT struct {
	Response *MasterInfoResponseT
}

func (t *MasterInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	responseOffset := t.Response.Pack(builder)
	MasterInfoStart(builder)
	MasterInfoAddResponse(builder, responseOffset)
	return MasterInfoEnd(builder)
}

func (rcv *MasterInfo) UnPackTo(t *MasterInfoT) {
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *MasterInfo) UnPack() *MasterInfoT {
	if rcv == nil { return nil }
	t := &MasterInfoT{}
	rcv.UnPackTo(t)
	return t
}

type MasterInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsMasterInfo(buf []byte, offset flatbuffers.UOffsetT) *MasterInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MasterInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMasterInfo(buf []byte, offset flatbuffers.UOffsetT) *MasterInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MasterInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MasterInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MasterInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MasterInfo) Response(obj *MasterInfoResponse) *MasterInfoResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MasterInfoResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func MasterInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MasterInfoAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(response), 0)
}
func MasterInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
