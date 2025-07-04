// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package companion

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Companion model mapping description
type MirrorMappingT struct {
	Name string `json:"name"`
	MirrorMappingTable []*NodeIdDlAddressMappingT `json:"mirrorMappingTable"`
}

func (t *MirrorMappingT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	mirrorMappingTableOffset := flatbuffers.UOffsetT(0)
	if t.MirrorMappingTable != nil {
		mirrorMappingTableLength := len(t.MirrorMappingTable)
		mirrorMappingTableOffsets := make([]flatbuffers.UOffsetT, mirrorMappingTableLength)
		for j := 0; j < mirrorMappingTableLength; j++ {
			mirrorMappingTableOffsets[j] = t.MirrorMappingTable[j].Pack(builder)
		}
		MirrorMappingStartMirrorMappingTableVector(builder, mirrorMappingTableLength)
		for j := mirrorMappingTableLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(mirrorMappingTableOffsets[j])
		}
		mirrorMappingTableOffset = builder.EndVector(mirrorMappingTableLength)
	}
	MirrorMappingStart(builder)
	MirrorMappingAddName(builder, nameOffset)
	MirrorMappingAddMirrorMappingTable(builder, mirrorMappingTableOffset)
	return MirrorMappingEnd(builder)
}

func (rcv *MirrorMapping) UnPackTo(t *MirrorMappingT) {
	t.Name = string(rcv.Name())
	mirrorMappingTableLength := rcv.MirrorMappingTableLength()
	t.MirrorMappingTable = make([]*NodeIdDlAddressMappingT, mirrorMappingTableLength)
	for j := 0; j < mirrorMappingTableLength; j++ {
		x := NodeIdDlAddressMapping{}
		rcv.MirrorMappingTable(&x, j)
		t.MirrorMappingTable[j] = x.UnPack()
	}
}

func (rcv *MirrorMapping) UnPack() *MirrorMappingT {
	if rcv == nil { return nil }
	t := &MirrorMappingT{}
	rcv.UnPackTo(t)
	return t
}

type MirrorMapping struct {
	_tab flatbuffers.Table
}

func GetRootAsMirrorMapping(buf []byte, offset flatbuffers.UOffsetT) *MirrorMapping {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MirrorMapping{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMirrorMapping(buf []byte, offset flatbuffers.UOffsetT) *MirrorMapping {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MirrorMapping{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MirrorMapping) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MirrorMapping) Table() flatbuffers.Table {
	return rcv._tab
}

/// Name of the mapping
func (rcv *MirrorMapping) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Name of the mapping
/// Mapping with mirror of the source NodeIds into datalayer address space
func (rcv *MirrorMapping) MirrorMappingTable(obj *NodeIdDlAddressMapping, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *MirrorMapping) MirrorMappingTableLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Mapping with mirror of the source NodeIds into datalayer address space
func MirrorMappingStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MirrorMappingAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func MirrorMappingAddMirrorMappingTable(builder *flatbuffers.Builder, mirrorMappingTable flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(mirrorMappingTable), 0)
}
func MirrorMappingStartMirrorMappingTableVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MirrorMappingEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
