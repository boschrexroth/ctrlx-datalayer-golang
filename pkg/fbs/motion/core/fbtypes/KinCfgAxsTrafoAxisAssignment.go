// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// data of a all axis assignments for an axis transformation
type KinCfgAxsTrafoAxisAssignment struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCfgAxsTrafoAxisAssignment(buf []byte, offset flatbuffers.UOffsetT) *KinCfgAxsTrafoAxisAssignment {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCfgAxsTrafoAxisAssignment{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCfgAxsTrafoAxisAssignment(buf []byte, offset flatbuffers.UOffsetT) *KinCfgAxsTrafoAxisAssignment {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCfgAxsTrafoAxisAssignment{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCfgAxsTrafoAxisAssignment) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCfgAxsTrafoAxisAssignment) Table() flatbuffers.Table {
	return rcv._tab
}

/// assignment as pairs of <axis name; ACS index>
func (rcv *KinCfgAxsTrafoAxisAssignment) Assignment(obj *KinCfgAxsTrafoSingleAxisAssignment, j int) bool {
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

func (rcv *KinCfgAxsTrafoAxisAssignment) AssignmentLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// assignment as pairs of <axis name; ACS index>
func KinCfgAxsTrafoAxisAssignmentStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func KinCfgAxsTrafoAxisAssignmentAddAssignment(builder *flatbuffers.Builder, assignment flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(assignment), 0)
}
func KinCfgAxsTrafoAxisAssignmentStartAssignmentVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KinCfgAxsTrafoAxisAssignmentEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}