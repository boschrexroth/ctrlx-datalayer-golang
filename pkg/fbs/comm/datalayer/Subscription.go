// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SubscriptionT struct {
	Properties *SubscriptionPropertiesT `json:"properties"`
	Nodes []string `json:"nodes"`
}

func (t *SubscriptionT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	propertiesOffset := t.Properties.Pack(builder)
	nodesOffset := flatbuffers.UOffsetT(0)
	if t.Nodes != nil {
		nodesLength := len(t.Nodes)
		nodesOffsets := make([]flatbuffers.UOffsetT, nodesLength)
		for j := 0; j < nodesLength; j++ {
			nodesOffsets[j] = builder.CreateString(t.Nodes[j])
		}
		SubscriptionStartNodesVector(builder, nodesLength)
		for j := nodesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(nodesOffsets[j])
		}
		nodesOffset = builder.EndVector(nodesLength)
	}
	SubscriptionStart(builder)
	SubscriptionAddProperties(builder, propertiesOffset)
	SubscriptionAddNodes(builder, nodesOffset)
	return SubscriptionEnd(builder)
}

func (rcv *Subscription) UnPackTo(t *SubscriptionT) {
	t.Properties = rcv.Properties(nil).UnPack()
	nodesLength := rcv.NodesLength()
	t.Nodes = make([]string, nodesLength)
	for j := 0; j < nodesLength; j++ {
		t.Nodes[j] = string(rcv.Nodes(j))
	}
}

func (rcv *Subscription) UnPack() *SubscriptionT {
	if rcv == nil { return nil }
	t := &SubscriptionT{}
	rcv.UnPackTo(t)
	return t
}

type Subscription struct {
	_tab flatbuffers.Table
}

func GetRootAsSubscription(buf []byte, offset flatbuffers.UOffsetT) *Subscription {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Subscription{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSubscription(buf []byte, offset flatbuffers.UOffsetT) *Subscription {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Subscription{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Subscription) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Subscription) Table() flatbuffers.Table {
	return rcv._tab
}

/// properties of subscription
func (rcv *Subscription) Properties(obj *SubscriptionProperties) *SubscriptionProperties {
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

/// properties of subscription
/// subscribed nodes
func (rcv *Subscription) Nodes(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Subscription) NodesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// subscribed nodes
func SubscriptionStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SubscriptionAddProperties(builder *flatbuffers.Builder, properties flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(properties), 0)
}
func SubscriptionAddNodes(builder *flatbuffers.Builder, nodes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(nodes), 0)
}
func SubscriptionStartNodesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SubscriptionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
