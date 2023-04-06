// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters of the get collision guard state function
type GetCollisionGuardStateParamsT struct {
	EmergencyDist float64
	MinDist float64
	SlowDownDist float64
	SyncDist float64
	ActPos float64
	TargetPos float64
	ActBrakingDist float64
	NeighborPos float64
	NeighborTargetPos float64
	NeighborBrakingDist float64
	HysteresisDist float64
	PrevState CollisionGuardState
}

func (t *GetCollisionGuardStateParamsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	GetCollisionGuardStateParamsStart(builder)
	GetCollisionGuardStateParamsAddEmergencyDist(builder, t.EmergencyDist)
	GetCollisionGuardStateParamsAddMinDist(builder, t.MinDist)
	GetCollisionGuardStateParamsAddSlowDownDist(builder, t.SlowDownDist)
	GetCollisionGuardStateParamsAddSyncDist(builder, t.SyncDist)
	GetCollisionGuardStateParamsAddActPos(builder, t.ActPos)
	GetCollisionGuardStateParamsAddTargetPos(builder, t.TargetPos)
	GetCollisionGuardStateParamsAddActBrakingDist(builder, t.ActBrakingDist)
	GetCollisionGuardStateParamsAddNeighborPos(builder, t.NeighborPos)
	GetCollisionGuardStateParamsAddNeighborTargetPos(builder, t.NeighborTargetPos)
	GetCollisionGuardStateParamsAddNeighborBrakingDist(builder, t.NeighborBrakingDist)
	GetCollisionGuardStateParamsAddHysteresisDist(builder, t.HysteresisDist)
	GetCollisionGuardStateParamsAddPrevState(builder, t.PrevState)
	return GetCollisionGuardStateParamsEnd(builder)
}

func (rcv *GetCollisionGuardStateParams) UnPackTo(t *GetCollisionGuardStateParamsT) {
	t.EmergencyDist = rcv.EmergencyDist()
	t.MinDist = rcv.MinDist()
	t.SlowDownDist = rcv.SlowDownDist()
	t.SyncDist = rcv.SyncDist()
	t.ActPos = rcv.ActPos()
	t.TargetPos = rcv.TargetPos()
	t.ActBrakingDist = rcv.ActBrakingDist()
	t.NeighborPos = rcv.NeighborPos()
	t.NeighborTargetPos = rcv.NeighborTargetPos()
	t.NeighborBrakingDist = rcv.NeighborBrakingDist()
	t.HysteresisDist = rcv.HysteresisDist()
	t.PrevState = rcv.PrevState()
}

func (rcv *GetCollisionGuardStateParams) UnPack() *GetCollisionGuardStateParamsT {
	if rcv == nil { return nil }
	t := &GetCollisionGuardStateParamsT{}
	rcv.UnPackTo(t)
	return t
}

type GetCollisionGuardStateParams struct {
	_tab flatbuffers.Table
}

func GetRootAsGetCollisionGuardStateParams(buf []byte, offset flatbuffers.UOffsetT) *GetCollisionGuardStateParams {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GetCollisionGuardStateParams{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsGetCollisionGuardStateParams(buf []byte, offset flatbuffers.UOffsetT) *GetCollisionGuardStateParams {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &GetCollisionGuardStateParams{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *GetCollisionGuardStateParams) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GetCollisionGuardStateParams) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GetCollisionGuardStateParams) EmergencyDist() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GetCollisionGuardStateParams) MutateEmergencyDist(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

func (rcv *GetCollisionGuardStateParams) MinDist() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GetCollisionGuardStateParams) MutateMinDist(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func (rcv *GetCollisionGuardStateParams) SlowDownDist() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GetCollisionGuardStateParams) MutateSlowDownDist(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func (rcv *GetCollisionGuardStateParams) SyncDist() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GetCollisionGuardStateParams) MutateSyncDist(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

func (rcv *GetCollisionGuardStateParams) ActPos() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GetCollisionGuardStateParams) MutateActPos(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

func (rcv *GetCollisionGuardStateParams) TargetPos() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GetCollisionGuardStateParams) MutateTargetPos(n float64) bool {
	return rcv._tab.MutateFloat64Slot(14, n)
}

func (rcv *GetCollisionGuardStateParams) ActBrakingDist() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GetCollisionGuardStateParams) MutateActBrakingDist(n float64) bool {
	return rcv._tab.MutateFloat64Slot(16, n)
}

func (rcv *GetCollisionGuardStateParams) NeighborPos() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GetCollisionGuardStateParams) MutateNeighborPos(n float64) bool {
	return rcv._tab.MutateFloat64Slot(18, n)
}

func (rcv *GetCollisionGuardStateParams) NeighborTargetPos() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GetCollisionGuardStateParams) MutateNeighborTargetPos(n float64) bool {
	return rcv._tab.MutateFloat64Slot(20, n)
}

func (rcv *GetCollisionGuardStateParams) NeighborBrakingDist() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GetCollisionGuardStateParams) MutateNeighborBrakingDist(n float64) bool {
	return rcv._tab.MutateFloat64Slot(22, n)
}

func (rcv *GetCollisionGuardStateParams) HysteresisDist() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GetCollisionGuardStateParams) MutateHysteresisDist(n float64) bool {
	return rcv._tab.MutateFloat64Slot(24, n)
}

func (rcv *GetCollisionGuardStateParams) PrevState() CollisionGuardState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return CollisionGuardState(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *GetCollisionGuardStateParams) MutatePrevState(n CollisionGuardState) bool {
	return rcv._tab.MutateInt8Slot(26, int8(n))
}

func GetCollisionGuardStateParamsStart(builder *flatbuffers.Builder) {
	builder.StartObject(12)
}
func GetCollisionGuardStateParamsAddEmergencyDist(builder *flatbuffers.Builder, emergencyDist float64) {
	builder.PrependFloat64Slot(0, emergencyDist, 0.0)
}
func GetCollisionGuardStateParamsAddMinDist(builder *flatbuffers.Builder, minDist float64) {
	builder.PrependFloat64Slot(1, minDist, 0.0)
}
func GetCollisionGuardStateParamsAddSlowDownDist(builder *flatbuffers.Builder, slowDownDist float64) {
	builder.PrependFloat64Slot(2, slowDownDist, 0.0)
}
func GetCollisionGuardStateParamsAddSyncDist(builder *flatbuffers.Builder, syncDist float64) {
	builder.PrependFloat64Slot(3, syncDist, 0.0)
}
func GetCollisionGuardStateParamsAddActPos(builder *flatbuffers.Builder, actPos float64) {
	builder.PrependFloat64Slot(4, actPos, 0.0)
}
func GetCollisionGuardStateParamsAddTargetPos(builder *flatbuffers.Builder, targetPos float64) {
	builder.PrependFloat64Slot(5, targetPos, 0.0)
}
func GetCollisionGuardStateParamsAddActBrakingDist(builder *flatbuffers.Builder, actBrakingDist float64) {
	builder.PrependFloat64Slot(6, actBrakingDist, 0.0)
}
func GetCollisionGuardStateParamsAddNeighborPos(builder *flatbuffers.Builder, neighborPos float64) {
	builder.PrependFloat64Slot(7, neighborPos, 0.0)
}
func GetCollisionGuardStateParamsAddNeighborTargetPos(builder *flatbuffers.Builder, neighborTargetPos float64) {
	builder.PrependFloat64Slot(8, neighborTargetPos, 0.0)
}
func GetCollisionGuardStateParamsAddNeighborBrakingDist(builder *flatbuffers.Builder, neighborBrakingDist float64) {
	builder.PrependFloat64Slot(9, neighborBrakingDist, 0.0)
}
func GetCollisionGuardStateParamsAddHysteresisDist(builder *flatbuffers.Builder, hysteresisDist float64) {
	builder.PrependFloat64Slot(10, hysteresisDist, 0.0)
}
func GetCollisionGuardStateParamsAddPrevState(builder *flatbuffers.Builder, prevState CollisionGuardState) {
	builder.PrependInt8Slot(11, int8(prevState), 0)
}
func GetCollisionGuardStateParamsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
