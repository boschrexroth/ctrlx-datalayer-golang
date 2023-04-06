// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NumConnectedSlavesResponseT struct {
	NumSlaves uint32
}

func (t *NumConnectedSlavesResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	NumConnectedSlavesResponseStart(builder)
	NumConnectedSlavesResponseAddNumSlaves(builder, t.NumSlaves)
	return NumConnectedSlavesResponseEnd(builder)
}

func (rcv *NumConnectedSlavesResponse) UnPackTo(t *NumConnectedSlavesResponseT) {
	t.NumSlaves = rcv.NumSlaves()
}

func (rcv *NumConnectedSlavesResponse) UnPack() *NumConnectedSlavesResponseT {
	if rcv == nil { return nil }
	t := &NumConnectedSlavesResponseT{}
	rcv.UnPackTo(t)
	return t
}

type NumConnectedSlavesResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsNumConnectedSlavesResponse(buf []byte, offset flatbuffers.UOffsetT) *NumConnectedSlavesResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NumConnectedSlavesResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNumConnectedSlavesResponse(buf []byte, offset flatbuffers.UOffsetT) *NumConnectedSlavesResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &NumConnectedSlavesResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *NumConnectedSlavesResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NumConnectedSlavesResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NumConnectedSlavesResponse) NumSlaves() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NumConnectedSlavesResponse) MutateNumSlaves(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func NumConnectedSlavesResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func NumConnectedSlavesResponseAddNumSlaves(builder *flatbuffers.Builder, numSlaves uint32) {
	builder.PrependUint32Slot(0, numSlaves, 0)
}
func NumConnectedSlavesResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
