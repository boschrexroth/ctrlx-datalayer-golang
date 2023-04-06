// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SubscriptionDataT struct {
	Properties *SubscriptionPropertiesT
	Nodes []string
}

func (t *SubscriptionDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	propertiesOffset := t.Properties.Pack(builder)
	nodesOffset := flatbuffers.UOffsetT(0)
	if t.Nodes != nil {
		nodesLength := len(t.Nodes)
		nodesOffsets := make([]flatbuffers.UOffsetT, nodesLength)
		for j := 0; j < nodesLength; j++ {
			nodesOffsets[j] = builder.CreateString(t.Nodes[j])
		}
		SubscriptionDataStartNodesVector(builder, nodesLength)
		for j := nodesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(nodesOffsets[j])
		}
		nodesOffset = builder.EndVector(nodesLength)
	}
	SubscriptionDataStart(builder)
	SubscriptionDataAddProperties(builder, propertiesOffset)
	SubscriptionDataAddNodes(builder, nodesOffset)
	return SubscriptionDataEnd(builder)
}

func (rcv *SubscriptionData) UnPackTo(t *SubscriptionDataT) {
	t.Properties = rcv.Properties(nil).UnPack()
	nodesLength := rcv.NodesLength()
	t.Nodes = make([]string, nodesLength)
	for j := 0; j < nodesLength; j++ {
		t.Nodes[j] = string(rcv.Nodes(j))
	}
}

func (rcv *SubscriptionData) UnPack() *SubscriptionDataT {
	if rcv == nil { return nil }
	t := &SubscriptionDataT{}
	rcv.UnPackTo(t)
	return t
}

type SubscriptionData struct {
	_tab flatbuffers.Table
}

func GetRootAsSubscriptionData(buf []byte, offset flatbuffers.UOffsetT) *SubscriptionData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SubscriptionData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSubscriptionData(buf []byte, offset flatbuffers.UOffsetT) *SubscriptionData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SubscriptionData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SubscriptionData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SubscriptionData) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SubscriptionData) Properties(obj *SubscriptionProperties) *SubscriptionProperties {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SubscriptionProperties)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SubscriptionData) Nodes(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *SubscriptionData) NodesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func SubscriptionDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SubscriptionDataAddProperties(builder *flatbuffers.Builder, properties flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(properties), 0)
}
func SubscriptionDataAddNodes(builder *flatbuffers.Builder, nodes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(nodes), 0)
}
func SubscriptionDataStartNodesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SubscriptionDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
