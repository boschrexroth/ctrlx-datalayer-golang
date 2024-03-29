// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters and data of the active command
type AxsCmdValuesT struct {
	TargetPos float64 `json:"targetPos"`
	TargetVel float64 `json:"targetVel"`
	TargetTrq float64 `json:"targetTrq"`
	Lim *DynamicLimitsStateT `json:"lim"`
	CmdId uint64 `json:"cmdId"`
	Src *CmdSourceT `json:"src"`
	TargetPosUnit string `json:"targetPosUnit"`
	TargetVelUnit string `json:"targetVelUnit"`
	TargetTrqUnit string `json:"targetTrqUnit"`
}

func (t *AxsCmdValuesT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	limOffset := t.Lim.Pack(builder)
	srcOffset := t.Src.Pack(builder)
	targetPosUnitOffset := flatbuffers.UOffsetT(0)
	if t.TargetPosUnit != "" {
		targetPosUnitOffset = builder.CreateString(t.TargetPosUnit)
	}
	targetVelUnitOffset := flatbuffers.UOffsetT(0)
	if t.TargetVelUnit != "" {
		targetVelUnitOffset = builder.CreateString(t.TargetVelUnit)
	}
	targetTrqUnitOffset := flatbuffers.UOffsetT(0)
	if t.TargetTrqUnit != "" {
		targetTrqUnitOffset = builder.CreateString(t.TargetTrqUnit)
	}
	AxsCmdValuesStart(builder)
	AxsCmdValuesAddTargetPos(builder, t.TargetPos)
	AxsCmdValuesAddTargetVel(builder, t.TargetVel)
	AxsCmdValuesAddTargetTrq(builder, t.TargetTrq)
	AxsCmdValuesAddLim(builder, limOffset)
	AxsCmdValuesAddCmdId(builder, t.CmdId)
	AxsCmdValuesAddSrc(builder, srcOffset)
	AxsCmdValuesAddTargetPosUnit(builder, targetPosUnitOffset)
	AxsCmdValuesAddTargetVelUnit(builder, targetVelUnitOffset)
	AxsCmdValuesAddTargetTrqUnit(builder, targetTrqUnitOffset)
	return AxsCmdValuesEnd(builder)
}

func (rcv *AxsCmdValues) UnPackTo(t *AxsCmdValuesT) {
	t.TargetPos = rcv.TargetPos()
	t.TargetVel = rcv.TargetVel()
	t.TargetTrq = rcv.TargetTrq()
	t.Lim = rcv.Lim(nil).UnPack()
	t.CmdId = rcv.CmdId()
	t.Src = rcv.Src(nil).UnPack()
	t.TargetPosUnit = string(rcv.TargetPosUnit())
	t.TargetVelUnit = string(rcv.TargetVelUnit())
	t.TargetTrqUnit = string(rcv.TargetTrqUnit())
}

func (rcv *AxsCmdValues) UnPack() *AxsCmdValuesT {
	if rcv == nil { return nil }
	t := &AxsCmdValuesT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCmdValues struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCmdValues(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdValues {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCmdValues{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCmdValues(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdValues {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCmdValues{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCmdValues) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCmdValues) Table() flatbuffers.Table {
	return rcv._tab
}

/// commanded target position
func (rcv *AxsCmdValues) TargetPos() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// commanded target position
func (rcv *AxsCmdValues) MutateTargetPos(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// commanded target velocity (currently not supported)
func (rcv *AxsCmdValues) TargetVel() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// commanded target velocity (currently not supported)
func (rcv *AxsCmdValues) MutateTargetVel(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// commanded target torque (currently not supported)
func (rcv *AxsCmdValues) TargetTrq() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// commanded target torque (currently not supported)
func (rcv *AxsCmdValues) MutateTargetTrq(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

/// dynamic limits for the motion of this command
func (rcv *AxsCmdValues) Lim(obj *DynamicLimitsState) *DynamicLimitsState {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(DynamicLimitsState)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// dynamic limits for the motion of this command
/// command ID of the active command
func (rcv *AxsCmdValues) CmdId() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// command ID of the active command
func (rcv *AxsCmdValues) MutateCmdId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(12, n)
}

/// command source (by which interface was this command inserted into the system (e.g. "PLC"))
func (rcv *AxsCmdValues) Src(obj *CmdSource) *CmdSource {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(CmdSource)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// command source (by which interface was this command inserted into the system (e.g. "PLC"))
/// unit of the commanded target position
func (rcv *AxsCmdValues) TargetPosUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the commanded target position
/// unit of the commanded target velocity (currently not supported)
func (rcv *AxsCmdValues) TargetVelUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the commanded target velocity (currently not supported)
/// unit of the commanded target torque (currently not supported)
func (rcv *AxsCmdValues) TargetTrqUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the commanded target torque (currently not supported)
func AxsCmdValuesStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
}
func AxsCmdValuesAddTargetPos(builder *flatbuffers.Builder, targetPos float64) {
	builder.PrependFloat64Slot(0, targetPos, 0.0)
}
func AxsCmdValuesAddTargetVel(builder *flatbuffers.Builder, targetVel float64) {
	builder.PrependFloat64Slot(1, targetVel, 0.0)
}
func AxsCmdValuesAddTargetTrq(builder *flatbuffers.Builder, targetTrq float64) {
	builder.PrependFloat64Slot(2, targetTrq, 0.0)
}
func AxsCmdValuesAddLim(builder *flatbuffers.Builder, lim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(lim), 0)
}
func AxsCmdValuesAddCmdId(builder *flatbuffers.Builder, cmdId uint64) {
	builder.PrependUint64Slot(4, cmdId, 0)
}
func AxsCmdValuesAddSrc(builder *flatbuffers.Builder, src flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(src), 0)
}
func AxsCmdValuesAddTargetPosUnit(builder *flatbuffers.Builder, targetPosUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(targetPosUnit), 0)
}
func AxsCmdValuesAddTargetVelUnit(builder *flatbuffers.Builder, targetVelUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(targetVelUnit), 0)
}
func AxsCmdValuesAddTargetTrqUnit(builder *flatbuffers.Builder, targetTrqUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(targetTrqUnit), 0)
}
func AxsCmdValuesEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
