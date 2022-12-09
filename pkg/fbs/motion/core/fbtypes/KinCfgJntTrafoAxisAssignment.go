// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// data of a all axis assignments for a joint transformation
type KinCfgJntTrafoAxisAssignmentT struct {
	Assignment []*KinCfgJntTrafoSingleAxisAssignmentT
}

func (t *KinCfgJntTrafoAxisAssignmentT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	assignmentOffset := flatbuffers.UOffsetT(0)
	if t.Assignment != nil {
		assignmentLength := len(t.Assignment)
		assignmentOffsets := make([]flatbuffers.UOffsetT, assignmentLength)
		for j := 0; j < assignmentLength; j++ {
			assignmentOffsets[j] = t.Assignment[j].Pack(builder)
		}
		KinCfgJntTrafoAxisAssignmentStartAssignmentVector(builder, assignmentLength)
		for j := assignmentLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(assignmentOffsets[j])
		}
		assignmentOffset = builder.EndVector(assignmentLength)
	}
	KinCfgJntTrafoAxisAssignmentStart(builder)
	KinCfgJntTrafoAxisAssignmentAddAssignment(builder, assignmentOffset)
	return KinCfgJntTrafoAxisAssignmentEnd(builder)
}

func (rcv *KinCfgJntTrafoAxisAssignment) UnPackTo(t *KinCfgJntTrafoAxisAssignmentT) {
	assignmentLength := rcv.AssignmentLength()
	t.Assignment = make([]*KinCfgJntTrafoSingleAxisAssignmentT, assignmentLength)
	for j := 0; j < assignmentLength; j++ {
		x := KinCfgJntTrafoSingleAxisAssignment{}
		rcv.Assignment(&x, j)
		t.Assignment[j] = x.UnPack()
	}
}

func (rcv *KinCfgJntTrafoAxisAssignment) UnPack() *KinCfgJntTrafoAxisAssignmentT {
	if rcv == nil { return nil }
	t := &KinCfgJntTrafoAxisAssignmentT{}
	rcv.UnPackTo(t)
	return t
}

type KinCfgJntTrafoAxisAssignment struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCfgJntTrafoAxisAssignment(buf []byte, offset flatbuffers.UOffsetT) *KinCfgJntTrafoAxisAssignment {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCfgJntTrafoAxisAssignment{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCfgJntTrafoAxisAssignment(buf []byte, offset flatbuffers.UOffsetT) *KinCfgJntTrafoAxisAssignment {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCfgJntTrafoAxisAssignment{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCfgJntTrafoAxisAssignment) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCfgJntTrafoAxisAssignment) Table() flatbuffers.Table {
	return rcv._tab
}

/// assignment as pairs of <axis name; ACS index>
func (rcv *KinCfgJntTrafoAxisAssignment) Assignment(obj *KinCfgJntTrafoSingleAxisAssignment, j int) bool {
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

func (rcv *KinCfgJntTrafoAxisAssignment) AssignmentLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// assignment as pairs of <axis name; ACS index>
func KinCfgJntTrafoAxisAssignmentStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func KinCfgJntTrafoAxisAssignmentAddAssignment(builder *flatbuffers.Builder, assignment flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(assignment), 0)
}
func KinCfgJntTrafoAxisAssignmentStartAssignmentVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KinCfgJntTrafoAxisAssignmentEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
