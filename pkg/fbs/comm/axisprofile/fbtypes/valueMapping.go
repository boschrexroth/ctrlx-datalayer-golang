// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type valueMappingT struct {
	Mapping []*mappingEntryT
}

func (t *valueMappingT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	mappingOffset := flatbuffers.UOffsetT(0)
	if t.Mapping != nil {
		mappingLength := len(t.Mapping)
		mappingOffsets := make([]flatbuffers.UOffsetT, mappingLength)
		for j := 0; j < mappingLength; j++ {
			mappingOffsets[j] = t.Mapping[j].Pack(builder)
		}
		valueMappingStartMappingVector(builder, mappingLength)
		for j := mappingLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(mappingOffsets[j])
		}
		mappingOffset = builder.EndVector(mappingLength)
	}
	valueMappingStart(builder)
	valueMappingAddMapping(builder, mappingOffset)
	return valueMappingEnd(builder)
}

func (rcv *valueMapping) UnPackTo(t *valueMappingT) {
	mappingLength := rcv.MappingLength()
	t.Mapping = make([]*mappingEntryT, mappingLength)
	for j := 0; j < mappingLength; j++ {
		x := mappingEntry{}
		rcv.Mapping(&x, j)
		t.Mapping[j] = x.UnPack()
	}
}

func (rcv *valueMapping) UnPack() *valueMappingT {
	if rcv == nil { return nil }
	t := &valueMappingT{}
	rcv.UnPackTo(t)
	return t
}

type valueMapping struct {
	_tab flatbuffers.Table
}

func GetRootAsvalueMapping(buf []byte, offset flatbuffers.UOffsetT) *valueMapping {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &valueMapping{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsvalueMapping(buf []byte, offset flatbuffers.UOffsetT) *valueMapping {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &valueMapping{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *valueMapping) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *valueMapping) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *valueMapping) Mapping(obj *mappingEntry, j int) bool {
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

func (rcv *valueMapping) MappingLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func valueMappingStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func valueMappingAddMapping(builder *flatbuffers.Builder, mapping flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(mapping), 0)
}
func valueMappingStartMappingVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func valueMappingEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
