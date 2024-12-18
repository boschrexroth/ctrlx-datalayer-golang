// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"

	motion__core__fbtypes "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/motion/core/fbtypes"
)

/// parameters for the move to point commands for a kinematic
type CmdKinMoveToPointDataT struct {
	PointName string `json:"pointName"`
	Lim *motion__core__fbtypes.DynamicLimitsT `json:"lim"`
	Buffered bool `json:"buffered"`
}

func (t *CmdKinMoveToPointDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	pointNameOffset := flatbuffers.UOffsetT(0)
	if t.PointName != "" {
		pointNameOffset = builder.CreateString(t.PointName)
	}
	limOffset := t.Lim.Pack(builder)
	CmdKinMoveToPointDataStart(builder)
	CmdKinMoveToPointDataAddPointName(builder, pointNameOffset)
	CmdKinMoveToPointDataAddLim(builder, limOffset)
	CmdKinMoveToPointDataAddBuffered(builder, t.Buffered)
	return CmdKinMoveToPointDataEnd(builder)
}

func (rcv *CmdKinMoveToPointData) UnPackTo(t *CmdKinMoveToPointDataT) {
	t.PointName = string(rcv.PointName())
	t.Lim = rcv.Lim(nil).UnPack()
	t.Buffered = rcv.Buffered()
}

func (rcv *CmdKinMoveToPointData) UnPack() *CmdKinMoveToPointDataT {
	if rcv == nil { return nil }
	t := &CmdKinMoveToPointDataT{}
	rcv.UnPackTo(t)
	return t
}

type CmdKinMoveToPointData struct {
	_tab flatbuffers.Table
}

func GetRootAsCmdKinMoveToPointData(buf []byte, offset flatbuffers.UOffsetT) *CmdKinMoveToPointData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CmdKinMoveToPointData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmdKinMoveToPointData(buf []byte, offset flatbuffers.UOffsetT) *CmdKinMoveToPointData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CmdKinMoveToPointData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CmdKinMoveToPointData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CmdKinMoveToPointData) Table() flatbuffers.Table {
	return rcv._tab
}

/// pointName
func (rcv *CmdKinMoveToPointData) PointName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// pointName
/// dynamic limits for the motion of this command
func (rcv *CmdKinMoveToPointData) Lim(obj *motion__core__fbtypes.DynamicLimits) *motion__core__fbtypes.DynamicLimits {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(motion__core__fbtypes.DynamicLimits)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// dynamic limits for the motion of this command
/// should this be a buffered command?
func (rcv *CmdKinMoveToPointData) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

/// should this be a buffered command?
func (rcv *CmdKinMoveToPointData) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func CmdKinMoveToPointDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func CmdKinMoveToPointDataAddPointName(builder *flatbuffers.Builder, pointName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(pointName), 0)
}
func CmdKinMoveToPointDataAddLim(builder *flatbuffers.Builder, lim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(lim), 0)
}
func CmdKinMoveToPointDataAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(2, buffered, true)
}
func CmdKinMoveToPointDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
