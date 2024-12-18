// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// position and orientation of a single joint
type PoseExternT struct {
	JointPose []float64 `json:"jointPose"`
}

func (t *PoseExternT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	jointPoseOffset := flatbuffers.UOffsetT(0)
	if t.JointPose != nil {
		jointPoseLength := len(t.JointPose)
		PoseExternStartJointPoseVector(builder, jointPoseLength)
		for j := jointPoseLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.JointPose[j])
		}
		jointPoseOffset = builder.EndVector(jointPoseLength)
	}
	PoseExternStart(builder)
	PoseExternAddJointPose(builder, jointPoseOffset)
	return PoseExternEnd(builder)
}

func (rcv *PoseExtern) UnPackTo(t *PoseExternT) {
	jointPoseLength := rcv.JointPoseLength()
	t.JointPose = make([]float64, jointPoseLength)
	for j := 0; j < jointPoseLength; j++ {
		t.JointPose[j] = rcv.JointPose(j)
	}
}

func (rcv *PoseExtern) UnPack() *PoseExternT {
	if rcv == nil { return nil }
	t := &PoseExternT{}
	rcv.UnPackTo(t)
	return t
}

type PoseExtern struct {
	_tab flatbuffers.Table
}

func GetRootAsPoseExtern(buf []byte, offset flatbuffers.UOffsetT) *PoseExtern {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PoseExtern{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPoseExtern(buf []byte, offset flatbuffers.UOffsetT) *PoseExtern {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PoseExtern{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *PoseExtern) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PoseExtern) Table() flatbuffers.Table {
	return rcv._tab
}

/// array of 6 double values consisting of 3 values for the xyz position and 3 values for the orientation (euler xyz extrinsic)
func (rcv *PoseExtern) JointPose(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *PoseExtern) JointPoseLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// array of 6 double values consisting of 3 values for the xyz position and 3 values for the orientation (euler xyz extrinsic)
func (rcv *PoseExtern) MutateJointPose(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func PoseExternStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func PoseExternAddJointPose(builder *flatbuffers.Builder, jointPose flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(jointPose), 0)
}
func PoseExternStartJointPoseVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func PoseExternEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
