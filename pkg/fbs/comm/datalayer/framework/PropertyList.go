// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package framework

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PropertyListT struct {
	Properties []*PropertyT
}

func (t *PropertyListT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	propertiesOffset := flatbuffers.UOffsetT(0)
	if t.Properties != nil {
		propertiesLength := len(t.Properties)
		propertiesOffsets := make([]flatbuffers.UOffsetT, propertiesLength)
		for j := 0; j < propertiesLength; j++ {
			propertiesOffsets[j] = t.Properties[j].Pack(builder)
		}
		PropertyListStartPropertiesVector(builder, propertiesLength)
		for j := propertiesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(propertiesOffsets[j])
		}
		propertiesOffset = builder.EndVector(propertiesLength)
	}
	PropertyListStart(builder)
	PropertyListAddProperties(builder, propertiesOffset)
	return PropertyListEnd(builder)
}

func (rcv *PropertyList) UnPackTo(t *PropertyListT) {
	propertiesLength := rcv.PropertiesLength()
	t.Properties = make([]*PropertyT, propertiesLength)
	for j := 0; j < propertiesLength; j++ {
		x := Property{}
		rcv.Properties(&x, j)
		t.Properties[j] = x.UnPack()
	}
}

func (rcv *PropertyList) UnPack() *PropertyListT {
	if rcv == nil { return nil }
	t := &PropertyListT{}
	rcv.UnPackTo(t)
	return t
}

type PropertyList struct {
	_tab flatbuffers.Table
}

func GetRootAsPropertyList(buf []byte, offset flatbuffers.UOffsetT) *PropertyList {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PropertyList{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPropertyList(buf []byte, offset flatbuffers.UOffsetT) *PropertyList {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PropertyList{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *PropertyList) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PropertyList) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PropertyList) Properties(obj *Property, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *PropertyList) PropertiesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func PropertyListStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func PropertyListAddProperties(builder *flatbuffers.Builder, properties flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(properties), 0)
}
func PropertyListStartPropertiesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PropertyListEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
