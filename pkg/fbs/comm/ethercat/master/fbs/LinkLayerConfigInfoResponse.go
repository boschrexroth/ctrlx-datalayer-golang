// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LinkLayerConfigInfoResponseT struct {
	Port string
	LinkLayer string
	Arguments string
}

func (t *LinkLayerConfigInfoResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	portOffset := builder.CreateString(t.Port)
	linkLayerOffset := builder.CreateString(t.LinkLayer)
	argumentsOffset := builder.CreateString(t.Arguments)
	LinkLayerConfigInfoResponseStart(builder)
	LinkLayerConfigInfoResponseAddPort(builder, portOffset)
	LinkLayerConfigInfoResponseAddLinkLayer(builder, linkLayerOffset)
	LinkLayerConfigInfoResponseAddArguments(builder, argumentsOffset)
	return LinkLayerConfigInfoResponseEnd(builder)
}

func (rcv *LinkLayerConfigInfoResponse) UnPackTo(t *LinkLayerConfigInfoResponseT) {
	t.Port = string(rcv.Port())
	t.LinkLayer = string(rcv.LinkLayer())
	t.Arguments = string(rcv.Arguments())
}

func (rcv *LinkLayerConfigInfoResponse) UnPack() *LinkLayerConfigInfoResponseT {
	if rcv == nil { return nil }
	t := &LinkLayerConfigInfoResponseT{}
	rcv.UnPackTo(t)
	return t
}

type LinkLayerConfigInfoResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsLinkLayerConfigInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *LinkLayerConfigInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LinkLayerConfigInfoResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLinkLayerConfigInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *LinkLayerConfigInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LinkLayerConfigInfoResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LinkLayerConfigInfoResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LinkLayerConfigInfoResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LinkLayerConfigInfoResponse) Port() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LinkLayerConfigInfoResponse) LinkLayer() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LinkLayerConfigInfoResponse) Arguments() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func LinkLayerConfigInfoResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func LinkLayerConfigInfoResponseAddPort(builder *flatbuffers.Builder, port flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(port), 0)
}
func LinkLayerConfigInfoResponseAddLinkLayer(builder *flatbuffers.Builder, linkLayer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(linkLayer), 0)
}
func LinkLayerConfigInfoResponseAddArguments(builder *flatbuffers.Builder, arguments flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(arguments), 0)
}
func LinkLayerConfigInfoResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
