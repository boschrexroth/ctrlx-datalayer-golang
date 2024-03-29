// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters for the extended move direct asynchronous command for a kinematics
type KinCmdMoveDirectAsyncExDataT struct {
	CmdKinPose []*KinCmdPosePairT `json:"cmdKinPose"`
	CoordSys CoordSys `json:"coordSys"`
	DynLimFactors *DynamicLimitsT `json:"dynLimFactors"`
	Buffered bool `json:"buffered"`
}

func (t *KinCmdMoveDirectAsyncExDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	cmdKinPoseOffset := flatbuffers.UOffsetT(0)
	if t.CmdKinPose != nil {
		cmdKinPoseLength := len(t.CmdKinPose)
		cmdKinPoseOffsets := make([]flatbuffers.UOffsetT, cmdKinPoseLength)
		for j := 0; j < cmdKinPoseLength; j++ {
			cmdKinPoseOffsets[j] = t.CmdKinPose[j].Pack(builder)
		}
		KinCmdMoveDirectAsyncExDataStartCmdKinPoseVector(builder, cmdKinPoseLength)
		for j := cmdKinPoseLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(cmdKinPoseOffsets[j])
		}
		cmdKinPoseOffset = builder.EndVector(cmdKinPoseLength)
	}
	dynLimFactorsOffset := t.DynLimFactors.Pack(builder)
	KinCmdMoveDirectAsyncExDataStart(builder)
	KinCmdMoveDirectAsyncExDataAddCmdKinPose(builder, cmdKinPoseOffset)
	KinCmdMoveDirectAsyncExDataAddCoordSys(builder, t.CoordSys)
	KinCmdMoveDirectAsyncExDataAddDynLimFactors(builder, dynLimFactorsOffset)
	KinCmdMoveDirectAsyncExDataAddBuffered(builder, t.Buffered)
	return KinCmdMoveDirectAsyncExDataEnd(builder)
}

func (rcv *KinCmdMoveDirectAsyncExData) UnPackTo(t *KinCmdMoveDirectAsyncExDataT) {
	cmdKinPoseLength := rcv.CmdKinPoseLength()
	t.CmdKinPose = make([]*KinCmdPosePairT, cmdKinPoseLength)
	for j := 0; j < cmdKinPoseLength; j++ {
		x := KinCmdPosePair{}
		rcv.CmdKinPose(&x, j)
		t.CmdKinPose[j] = x.UnPack()
	}
	t.CoordSys = rcv.CoordSys()
	t.DynLimFactors = rcv.DynLimFactors(nil).UnPack()
	t.Buffered = rcv.Buffered()
}

func (rcv *KinCmdMoveDirectAsyncExData) UnPack() *KinCmdMoveDirectAsyncExDataT {
	if rcv == nil { return nil }
	t := &KinCmdMoveDirectAsyncExDataT{}
	rcv.UnPackTo(t)
	return t
}

type KinCmdMoveDirectAsyncExData struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCmdMoveDirectAsyncExData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdMoveDirectAsyncExData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCmdMoveDirectAsyncExData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCmdMoveDirectAsyncExData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdMoveDirectAsyncExData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCmdMoveDirectAsyncExData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCmdMoveDirectAsyncExData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCmdMoveDirectAsyncExData) Table() flatbuffers.Table {
	return rcv._tab
}

/// commanded target position with meanings
func (rcv *KinCmdMoveDirectAsyncExData) CmdKinPose(obj *KinCmdPosePair, j int) bool {
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

func (rcv *KinCmdMoveDirectAsyncExData) CmdKinPoseLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// commanded target position with meanings
/// coordSys for commanded target position
func (rcv *KinCmdMoveDirectAsyncExData) CoordSys() CoordSys {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return CoordSys(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// coordSys for commanded target position
func (rcv *KinCmdMoveDirectAsyncExData) MutateCoordSys(n CoordSys) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

/// factors to reduce the individual dynamic limits for the motion of this command
func (rcv *KinCmdMoveDirectAsyncExData) DynLimFactors(obj *DynamicLimits) *DynamicLimits {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(DynamicLimits)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// factors to reduce the individual dynamic limits for the motion of this command
/// should this be a buffered command?
func (rcv *KinCmdMoveDirectAsyncExData) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

/// should this be a buffered command?
func (rcv *KinCmdMoveDirectAsyncExData) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func KinCmdMoveDirectAsyncExDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func KinCmdMoveDirectAsyncExDataAddCmdKinPose(builder *flatbuffers.Builder, cmdKinPose flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(cmdKinPose), 0)
}
func KinCmdMoveDirectAsyncExDataStartCmdKinPoseVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KinCmdMoveDirectAsyncExDataAddCoordSys(builder *flatbuffers.Builder, coordSys CoordSys) {
	builder.PrependInt8Slot(1, int8(coordSys), 0)
}
func KinCmdMoveDirectAsyncExDataAddDynLimFactors(builder *flatbuffers.Builder, dynLimFactors flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(dynLimFactors), 0)
}
func KinCmdMoveDirectAsyncExDataAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(3, buffered, true)
}
func KinCmdMoveDirectAsyncExDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
