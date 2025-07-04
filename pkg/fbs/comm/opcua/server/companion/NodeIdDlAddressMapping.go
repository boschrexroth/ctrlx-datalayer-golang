// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package companion

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Mapping of one sourceUaNodeId to targetDlAddress
type NodeIdDlAddressMappingT struct {
	SourceUaNodeId string `json:"sourceUaNodeId"`
	TargetDlAddress string `json:"targetDlAddress"`
	MappingResult []MirrorMappingResult `json:"mappingResult"`
}

func (t *NodeIdDlAddressMappingT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	sourceUaNodeIdOffset := flatbuffers.UOffsetT(0)
	if t.SourceUaNodeId != "" {
		sourceUaNodeIdOffset = builder.CreateString(t.SourceUaNodeId)
	}
	targetDlAddressOffset := flatbuffers.UOffsetT(0)
	if t.TargetDlAddress != "" {
		targetDlAddressOffset = builder.CreateString(t.TargetDlAddress)
	}
	mappingResultOffset := flatbuffers.UOffsetT(0)
	if t.MappingResult != nil {
		mappingResultLength := len(t.MappingResult)
		NodeIdDlAddressMappingStartMappingResultVector(builder, mappingResultLength)
		for j := mappingResultLength - 1; j >= 0; j-- {
			builder.PrependInt8(int8(t.MappingResult[j]))
		}
		mappingResultOffset = builder.EndVector(mappingResultLength)
	}
	NodeIdDlAddressMappingStart(builder)
	NodeIdDlAddressMappingAddSourceUaNodeId(builder, sourceUaNodeIdOffset)
	NodeIdDlAddressMappingAddTargetDlAddress(builder, targetDlAddressOffset)
	NodeIdDlAddressMappingAddMappingResult(builder, mappingResultOffset)
	return NodeIdDlAddressMappingEnd(builder)
}

func (rcv *NodeIdDlAddressMapping) UnPackTo(t *NodeIdDlAddressMappingT) {
	t.SourceUaNodeId = string(rcv.SourceUaNodeId())
	t.TargetDlAddress = string(rcv.TargetDlAddress())
	mappingResultLength := rcv.MappingResultLength()
	t.MappingResult = make([]MirrorMappingResult, mappingResultLength)
	for j := 0; j < mappingResultLength; j++ {
		t.MappingResult[j] = rcv.MappingResult(j)
	}
}

func (rcv *NodeIdDlAddressMapping) UnPack() *NodeIdDlAddressMappingT {
	if rcv == nil { return nil }
	t := &NodeIdDlAddressMappingT{}
	rcv.UnPackTo(t)
	return t
}

type NodeIdDlAddressMapping struct {
	_tab flatbuffers.Table
}

func GetRootAsNodeIdDlAddressMapping(buf []byte, offset flatbuffers.UOffsetT) *NodeIdDlAddressMapping {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NodeIdDlAddressMapping{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNodeIdDlAddressMapping(buf []byte, offset flatbuffers.UOffsetT) *NodeIdDlAddressMapping {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &NodeIdDlAddressMapping{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *NodeIdDlAddressMapping) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NodeIdDlAddressMapping) Table() flatbuffers.Table {
	return rcv._tab
}

/// The original OPC UA NodeId which will be mapped into the datalayer address space, e.g. ns=2;s=plc/app/Appl/...
func (rcv *NodeIdDlAddressMapping) SourceUaNodeId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The original OPC UA NodeId which will be mapped into the datalayer address space, e.g. ns=2;s=plc/app/Appl/...
/// The target datalayer address in the datalayer address space, e.g. mydlnode/foo/bar/...
func (rcv *NodeIdDlAddressMapping) TargetDlAddress() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The target datalayer address in the datalayer address space, e.g. mydlnode/foo/bar/...
/// Verification of the mapping, used as output parameter
func (rcv *NodeIdDlAddressMapping) MappingResult(j int) MirrorMappingResult {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return MirrorMappingResult(rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1)))
	}
	return 0
}

func (rcv *NodeIdDlAddressMapping) MappingResultLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Verification of the mapping, used as output parameter
func (rcv *NodeIdDlAddressMapping) MutateMappingResult(j int, n MirrorMappingResult) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), int8(n))
	}
	return false
}

func NodeIdDlAddressMappingStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func NodeIdDlAddressMappingAddSourceUaNodeId(builder *flatbuffers.Builder, sourceUaNodeId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(sourceUaNodeId), 0)
}
func NodeIdDlAddressMappingAddTargetDlAddress(builder *flatbuffers.Builder, targetDlAddress flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(targetDlAddress), 0)
}
func NodeIdDlAddressMappingAddMappingResult(builder *flatbuffers.Builder, mappingResult flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(mappingResult), 0)
}
func NodeIdDlAddressMappingStartMappingResultVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func NodeIdDlAddressMappingEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
