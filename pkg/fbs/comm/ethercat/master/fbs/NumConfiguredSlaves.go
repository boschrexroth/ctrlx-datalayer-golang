// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NumConfiguredSlavesT struct {
	Response *NumConfiguredSlavesResponseT
}

func (t *NumConfiguredSlavesT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	responseOffset := t.Response.Pack(builder)
	NumConfiguredSlavesStart(builder)
	NumConfiguredSlavesAddResponse(builder, responseOffset)
	return NumConfiguredSlavesEnd(builder)
}

func (rcv *NumConfiguredSlaves) UnPackTo(t *NumConfiguredSlavesT) {
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *NumConfiguredSlaves) UnPack() *NumConfiguredSlavesT {
	if rcv == nil { return nil }
	t := &NumConfiguredSlavesT{}
	rcv.UnPackTo(t)
	return t
}

type NumConfiguredSlaves struct {
	_tab flatbuffers.Table
}

func GetRootAsNumConfiguredSlaves(buf []byte, offset flatbuffers.UOffsetT) *NumConfiguredSlaves {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NumConfiguredSlaves{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNumConfiguredSlaves(buf []byte, offset flatbuffers.UOffsetT) *NumConfiguredSlaves {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &NumConfiguredSlaves{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *NumConfiguredSlaves) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NumConfiguredSlaves) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NumConfiguredSlaves) Response(obj *NumConfiguredSlavesResponse) *NumConfiguredSlavesResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(NumConfiguredSlavesResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func NumConfiguredSlavesStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func NumConfiguredSlavesAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(response), 0)
}
func NumConfiguredSlavesEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
