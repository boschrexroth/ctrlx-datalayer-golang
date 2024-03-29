// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters for the extended move direct command for a kinematics
type KinCmdMoveDirectExDataT struct {
	CmdKinPose []*KinCmdPosePairT `json:"cmdKinPose"`
	CoordSys CoordSys `json:"coordSys"`
	Lim *DynamicLimitsT `json:"lim"`
	Buffered bool `json:"buffered"`
}

func (t *KinCmdMoveDirectExDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	cmdKinPoseOffset := flatbuffers.UOffsetT(0)
	if t.CmdKinPose != nil {
		cmdKinPoseLength := len(t.CmdKinPose)
		cmdKinPoseOffsets := make([]flatbuffers.UOffsetT, cmdKinPoseLength)
		for j := 0; j < cmdKinPoseLength; j++ {
			cmdKinPoseOffsets[j] = t.CmdKinPose[j].Pack(builder)
		}
		KinCmdMoveDirectExDataStartCmdKinPoseVector(builder, cmdKinPoseLength)
		for j := cmdKinPoseLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(cmdKinPoseOffsets[j])
		}
		cmdKinPoseOffset = builder.EndVector(cmdKinPoseLength)
	}
	limOffset := t.Lim.Pack(builder)
	KinCmdMoveDirectExDataStart(builder)
	KinCmdMoveDirectExDataAddCmdKinPose(builder, cmdKinPoseOffset)
	KinCmdMoveDirectExDataAddCoordSys(builder, t.CoordSys)
	KinCmdMoveDirectExDataAddLim(builder, limOffset)
	KinCmdMoveDirectExDataAddBuffered(builder, t.Buffered)
	return KinCmdMoveDirectExDataEnd(builder)
}

func (rcv *KinCmdMoveDirectExData) UnPackTo(t *KinCmdMoveDirectExDataT) {
	cmdKinPoseLength := rcv.CmdKinPoseLength()
	t.CmdKinPose = make([]*KinCmdPosePairT, cmdKinPoseLength)
	for j := 0; j < cmdKinPoseLength; j++ {
		x := KinCmdPosePair{}
		rcv.CmdKinPose(&x, j)
		t.CmdKinPose[j] = x.UnPack()
	}
	t.CoordSys = rcv.CoordSys()
	t.Lim = rcv.Lim(nil).UnPack()
	t.Buffered = rcv.Buffered()
}

func (rcv *KinCmdMoveDirectExData) UnPack() *KinCmdMoveDirectExDataT {
	if rcv == nil { return nil }
	t := &KinCmdMoveDirectExDataT{}
	rcv.UnPackTo(t)
	return t
}

type KinCmdMoveDirectExData struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCmdMoveDirectExData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdMoveDirectExData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCmdMoveDirectExData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCmdMoveDirectExData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdMoveDirectExData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCmdMoveDirectExData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCmdMoveDirectExData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCmdMoveDirectExData) Table() flatbuffers.Table {
	return rcv._tab
}

/// commanded target position with meanings
func (rcv *KinCmdMoveDirectExData) CmdKinPose(obj *KinCmdPosePair, j int) bool {
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

func (rcv *KinCmdMoveDirectExData) CmdKinPoseLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// commanded target position with meanings
/// coordSys for commanded target position
func (rcv *KinCmdMoveDirectExData) CoordSys() CoordSys {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return CoordSys(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// coordSys for commanded target position
func (rcv *KinCmdMoveDirectExData) MutateCoordSys(n CoordSys) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

/// dynamic limits for the motion of this command
func (rcv *KinCmdMoveDirectExData) Lim(obj *DynamicLimits) *DynamicLimits {
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

/// dynamic limits for the motion of this command
/// should this be a buffered command?
func (rcv *KinCmdMoveDirectExData) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

/// should this be a buffered command?
func (rcv *KinCmdMoveDirectExData) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func KinCmdMoveDirectExDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func KinCmdMoveDirectExDataAddCmdKinPose(builder *flatbuffers.Builder, cmdKinPose flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(cmdKinPose), 0)
}
func KinCmdMoveDirectExDataStartCmdKinPoseVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KinCmdMoveDirectExDataAddCoordSys(builder *flatbuffers.Builder, coordSys CoordSys) {
	builder.PrependInt8Slot(1, int8(coordSys), 0)
}
func KinCmdMoveDirectExDataAddLim(builder *flatbuffers.Builder, lim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(lim), 0)
}
func KinCmdMoveDirectExDataAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(3, buffered, true)
}
func KinCmdMoveDirectExDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
