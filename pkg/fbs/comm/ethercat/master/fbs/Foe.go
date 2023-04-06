// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FoeT struct {
	Request *FoeRequestT
	Response *FoeResponseT
}

func (t *FoeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	requestOffset := t.Request.Pack(builder)
	responseOffset := t.Response.Pack(builder)
	FoeStart(builder)
	FoeAddRequest(builder, requestOffset)
	FoeAddResponse(builder, responseOffset)
	return FoeEnd(builder)
}

func (rcv *Foe) UnPackTo(t *FoeT) {
	t.Request = rcv.Request(nil).UnPack()
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *Foe) UnPack() *FoeT {
	if rcv == nil { return nil }
	t := &FoeT{}
	rcv.UnPackTo(t)
	return t
}

type Foe struct {
	_tab flatbuffers.Table
}

func GetRootAsFoe(buf []byte, offset flatbuffers.UOffsetT) *Foe {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Foe{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFoe(buf []byte, offset flatbuffers.UOffsetT) *Foe {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Foe{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Foe) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Foe) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Foe) Request(obj *FoeRequest) *FoeRequest {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FoeRequest)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Foe) Response(obj *FoeResponse) *FoeResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FoeResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func FoeStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func FoeAddRequest(builder *flatbuffers.Builder, request flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(request), 0)
}
func FoeAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(response), 0)
}
func FoeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
