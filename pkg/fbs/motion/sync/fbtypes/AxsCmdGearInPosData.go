// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters of the axis GearInPos command
type AxsCmdGearInPosDataT struct {
	Master string `json:"master"`
	SyncSource SyncSource `json:"syncSource"`
	SyncMode SyncMode `json:"syncMode"`
	DynSyncDirection DynSyncDirection `json:"dynSyncDirection"`
	Parameters *AxsCmdGearInPosParamsT `json:"parameters"`
	Buffered bool `json:"buffered"`
}

func (t *AxsCmdGearInPosDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	masterOffset := flatbuffers.UOffsetT(0)
	if t.Master != "" {
		masterOffset = builder.CreateString(t.Master)
	}
	parametersOffset := t.Parameters.Pack(builder)
	AxsCmdGearInPosDataStart(builder)
	AxsCmdGearInPosDataAddMaster(builder, masterOffset)
	AxsCmdGearInPosDataAddSyncSource(builder, t.SyncSource)
	AxsCmdGearInPosDataAddSyncMode(builder, t.SyncMode)
	AxsCmdGearInPosDataAddDynSyncDirection(builder, t.DynSyncDirection)
	AxsCmdGearInPosDataAddParameters(builder, parametersOffset)
	AxsCmdGearInPosDataAddBuffered(builder, t.Buffered)
	return AxsCmdGearInPosDataEnd(builder)
}

func (rcv *AxsCmdGearInPosData) UnPackTo(t *AxsCmdGearInPosDataT) {
	t.Master = string(rcv.Master())
	t.SyncSource = rcv.SyncSource()
	t.SyncMode = rcv.SyncMode()
	t.DynSyncDirection = rcv.DynSyncDirection()
	t.Parameters = rcv.Parameters(nil).UnPack()
	t.Buffered = rcv.Buffered()
}

func (rcv *AxsCmdGearInPosData) UnPack() *AxsCmdGearInPosDataT {
	if rcv == nil { return nil }
	t := &AxsCmdGearInPosDataT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCmdGearInPosData struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCmdGearInPosData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdGearInPosData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCmdGearInPosData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCmdGearInPosData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdGearInPosData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCmdGearInPosData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCmdGearInPosData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCmdGearInPosData) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the master axis
func (rcv *AxsCmdGearInPosData) Master() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the master axis
/// sync source (Actual/Setpoint/Actual-extrapolated)
func (rcv *AxsCmdGearInPosData) SyncSource() SyncSource {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return SyncSource(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// sync source (Actual/Setpoint/Actual-extrapolated)
func (rcv *AxsCmdGearInPosData) MutateSyncSource(n SyncSource) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

/// sync mode (Abs/Rel; Init/KeepState)
func (rcv *AxsCmdGearInPosData) SyncMode() SyncMode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return SyncMode(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// sync mode (Abs/Rel; Init/KeepState)
func (rcv *AxsCmdGearInPosData) MutateSyncMode(n SyncMode) bool {
	return rcv._tab.MutateInt8Slot(8, int8(n))
}

/// dynamic sync direction (Positive/Negative/ShortestWay)
func (rcv *AxsCmdGearInPosData) DynSyncDirection() DynSyncDirection {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return DynSyncDirection(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// dynamic sync direction (Positive/Negative/ShortestWay)
func (rcv *AxsCmdGearInPosData) MutateDynSyncDirection(n DynSyncDirection) bool {
	return rcv._tab.MutateInt8Slot(10, int8(n))
}

/// gear in pos parameters (master offset, slave offset, ratio numerator, ratio denominator, fine adjust)
func (rcv *AxsCmdGearInPosData) Parameters(obj *AxsCmdGearInPosParams) *AxsCmdGearInPosParams {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AxsCmdGearInPosParams)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// gear in pos parameters (master offset, slave offset, ratio numerator, ratio denominator, fine adjust)
/// should this be a buffered command?
func (rcv *AxsCmdGearInPosData) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// should this be a buffered command?
func (rcv *AxsCmdGearInPosData) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(14, n)
}

func AxsCmdGearInPosDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func AxsCmdGearInPosDataAddMaster(builder *flatbuffers.Builder, master flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(master), 0)
}
func AxsCmdGearInPosDataAddSyncSource(builder *flatbuffers.Builder, syncSource SyncSource) {
	builder.PrependInt8Slot(1, int8(syncSource), 0)
}
func AxsCmdGearInPosDataAddSyncMode(builder *flatbuffers.Builder, syncMode SyncMode) {
	builder.PrependInt8Slot(2, int8(syncMode), 0)
}
func AxsCmdGearInPosDataAddDynSyncDirection(builder *flatbuffers.Builder, dynSyncDirection DynSyncDirection) {
	builder.PrependInt8Slot(3, int8(dynSyncDirection), 0)
}
func AxsCmdGearInPosDataAddParameters(builder *flatbuffers.Builder, parameters flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(parameters), 0)
}
func AxsCmdGearInPosDataAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(5, buffered, false)
}
func AxsCmdGearInPosDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
