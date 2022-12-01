// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NumConnectedDcSlavesT struct {
	Response *NumConnectedDcSlavesResponseT
}

func (t *NumConnectedDcSlavesT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	responseOffset := t.Response.Pack(builder)
	NumConnectedDcSlavesStart(builder)
	NumConnectedDcSlavesAddResponse(builder, responseOffset)
	return NumConnectedDcSlavesEnd(builder)
}

func (rcv *NumConnectedDcSlaves) UnPackTo(t *NumConnectedDcSlavesT) {
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *NumConnectedDcSlaves) UnPack() *NumConnectedDcSlavesT {
	if rcv == nil { return nil }
	t := &NumConnectedDcSlavesT{}
	rcv.UnPackTo(t)
	return t
}

type NumConnectedDcSlaves struct {
	_tab flatbuffers.Table
}

func GetRootAsNumConnectedDcSlaves(buf []byte, offset flatbuffers.UOffsetT) *NumConnectedDcSlaves {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NumConnectedDcSlaves{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNumConnectedDcSlaves(buf []byte, offset flatbuffers.UOffsetT) *NumConnectedDcSlaves {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &NumConnectedDcSlaves{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *NumConnectedDcSlaves) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NumConnectedDcSlaves) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NumConnectedDcSlaves) Response(obj *NumConnectedDcSlavesResponse) *NumConnectedDcSlavesResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(NumConnectedDcSlavesResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func NumConnectedDcSlavesStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func NumConnectedDcSlavesAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(response), 0)
}
func NumConnectedDcSlavesEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
