// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters of the axis GearInVelEx command
type AxsCmdGearInVelExDataT struct {
	Master string `json:"master"`
	SyncSource SyncSource `json:"syncSource"`
	Parameters *AxsCmdGearInVelParamsT `json:"parameters"`
	DlParameters *AxsCmdGearInVelDlParamsT `json:"dlParameters"`
	Buffered bool `json:"buffered"`
}

func (t *AxsCmdGearInVelExDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	masterOffset := flatbuffers.UOffsetT(0)
	if t.Master != "" {
		masterOffset = builder.CreateString(t.Master)
	}
	parametersOffset := t.Parameters.Pack(builder)
	dlParametersOffset := t.DlParameters.Pack(builder)
	AxsCmdGearInVelExDataStart(builder)
	AxsCmdGearInVelExDataAddMaster(builder, masterOffset)
	AxsCmdGearInVelExDataAddSyncSource(builder, t.SyncSource)
	AxsCmdGearInVelExDataAddParameters(builder, parametersOffset)
	AxsCmdGearInVelExDataAddDlParameters(builder, dlParametersOffset)
	AxsCmdGearInVelExDataAddBuffered(builder, t.Buffered)
	return AxsCmdGearInVelExDataEnd(builder)
}

func (rcv *AxsCmdGearInVelExData) UnPackTo(t *AxsCmdGearInVelExDataT) {
	t.Master = string(rcv.Master())
	t.SyncSource = rcv.SyncSource()
	t.Parameters = rcv.Parameters(nil).UnPack()
	t.DlParameters = rcv.DlParameters(nil).UnPack()
	t.Buffered = rcv.Buffered()
}

func (rcv *AxsCmdGearInVelExData) UnPack() *AxsCmdGearInVelExDataT {
	if rcv == nil { return nil }
	t := &AxsCmdGearInVelExDataT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCmdGearInVelExData struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCmdGearInVelExData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdGearInVelExData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCmdGearInVelExData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCmdGearInVelExData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdGearInVelExData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCmdGearInVelExData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCmdGearInVelExData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCmdGearInVelExData) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the master axis
func (rcv *AxsCmdGearInVelExData) Master() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the master axis
/// sync source (Actual/Setpoint/Actual-extrapolated)
func (rcv *AxsCmdGearInVelExData) SyncSource() SyncSource {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return SyncSource(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// sync source (Actual/Setpoint/Actual-extrapolated)
func (rcv *AxsCmdGearInVelExData) MutateSyncSource(n SyncSource) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

/// gear in vel parameters
func (rcv *AxsCmdGearInVelExData) Parameters(obj *AxsCmdGearInVelParams) *AxsCmdGearInVelParams {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AxsCmdGearInVelParams)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// gear in vel parameters
/// gear in vel data layer parameters
func (rcv *AxsCmdGearInVelExData) DlParameters(obj *AxsCmdGearInVelDlParams) *AxsCmdGearInVelDlParams {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AxsCmdGearInVelDlParams)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// gear in vel data layer parameters
/// should this be a buffered command?
func (rcv *AxsCmdGearInVelExData) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// should this be a buffered command?
func (rcv *AxsCmdGearInVelExData) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

func AxsCmdGearInVelExDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func AxsCmdGearInVelExDataAddMaster(builder *flatbuffers.Builder, master flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(master), 0)
}
func AxsCmdGearInVelExDataAddSyncSource(builder *flatbuffers.Builder, syncSource SyncSource) {
	builder.PrependInt8Slot(1, int8(syncSource), 0)
}
func AxsCmdGearInVelExDataAddParameters(builder *flatbuffers.Builder, parameters flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(parameters), 0)
}
func AxsCmdGearInVelExDataAddDlParameters(builder *flatbuffers.Builder, dlParameters flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(dlParameters), 0)
}
func AxsCmdGearInVelExDataAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(4, buffered, false)
}
func AxsCmdGearInVelExDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
