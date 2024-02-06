// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package companion

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Companion model mapping description
type MappingT struct {
	Name string `json:"name"`
	TypeSafety bool `json:"typeSafety"`
	MappingTable []*NodeIdMappingT `json:"mappingTable"`
}

func (t *MappingT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	mappingTableOffset := flatbuffers.UOffsetT(0)
	if t.MappingTable != nil {
		mappingTableLength := len(t.MappingTable)
		mappingTableOffsets := make([]flatbuffers.UOffsetT, mappingTableLength)
		for j := 0; j < mappingTableLength; j++ {
			mappingTableOffsets[j] = t.MappingTable[j].Pack(builder)
		}
		MappingStartMappingTableVector(builder, mappingTableLength)
		for j := mappingTableLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(mappingTableOffsets[j])
		}
		mappingTableOffset = builder.EndVector(mappingTableLength)
	}
	MappingStart(builder)
	MappingAddName(builder, nameOffset)
	MappingAddTypeSafety(builder, t.TypeSafety)
	MappingAddMappingTable(builder, mappingTableOffset)
	return MappingEnd(builder)
}

func (rcv *Mapping) UnPackTo(t *MappingT) {
	t.Name = string(rcv.Name())
	t.TypeSafety = rcv.TypeSafety()
	mappingTableLength := rcv.MappingTableLength()
	t.MappingTable = make([]*NodeIdMappingT, mappingTableLength)
	for j := 0; j < mappingTableLength; j++ {
		x := NodeIdMapping{}
		rcv.MappingTable(&x, j)
		t.MappingTable[j] = x.UnPack()
	}
}

func (rcv *Mapping) UnPack() *MappingT {
	if rcv == nil { return nil }
	t := &MappingT{}
	rcv.UnPackTo(t)
	return t
}

type Mapping struct {
	_tab flatbuffers.Table
}

func GetRootAsMapping(buf []byte, offset flatbuffers.UOffsetT) *Mapping {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Mapping{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMapping(buf []byte, offset flatbuffers.UOffsetT) *Mapping {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Mapping{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Mapping) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Mapping) Table() flatbuffers.Table {
	return rcv._tab
}

/// Name of the mapping
func (rcv *Mapping) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Name of the mapping
/// Check if sourceUaNodeId and targetUaNodeId have same datatype
func (rcv *Mapping) TypeSafety() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

/// Check if sourceUaNodeId and targetUaNodeId have same datatype
func (rcv *Mapping) MutateTypeSafety(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

/// Mapping of the OPC UA NodeIds
func (rcv *Mapping) MappingTable(obj *NodeIdMapping, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Mapping) MappingTableLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Mapping of the OPC UA NodeIds
func MappingStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func MappingAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func MappingAddTypeSafety(builder *flatbuffers.Builder, typeSafety bool) {
	builder.PrependBoolSlot(1, typeSafety, true)
}
func MappingAddMappingTable(builder *flatbuffers.Builder, mappingTable flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(mappingTable), 0)
}
func MappingStartMappingTableVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MappingEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
