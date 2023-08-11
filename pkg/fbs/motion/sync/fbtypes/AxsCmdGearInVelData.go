// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters of the axis GearInVel commands
type AxsCmdGearInVelDataT struct {
	Master string `json:"master"`
	SyncSource SyncSource `json:"syncSource"`
	Parameters *AxsCmdGearInVelParamsT `json:"parameters"`
	Buffered bool `json:"buffered"`
}

func (t *AxsCmdGearInVelDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	masterOffset := flatbuffers.UOffsetT(0)
	if t.Master != "" {
		masterOffset = builder.CreateString(t.Master)
	}
	parametersOffset := t.Parameters.Pack(builder)
	AxsCmdGearInVelDataStart(builder)
	AxsCmdGearInVelDataAddMaster(builder, masterOffset)
	AxsCmdGearInVelDataAddSyncSource(builder, t.SyncSource)
	AxsCmdGearInVelDataAddParameters(builder, parametersOffset)
	AxsCmdGearInVelDataAddBuffered(builder, t.Buffered)
	return AxsCmdGearInVelDataEnd(builder)
}

func (rcv *AxsCmdGearInVelData) UnPackTo(t *AxsCmdGearInVelDataT) {
	t.Master = string(rcv.Master())
	t.SyncSource = rcv.SyncSource()
	t.Parameters = rcv.Parameters(nil).UnPack()
	t.Buffered = rcv.Buffered()
}

func (rcv *AxsCmdGearInVelData) UnPack() *AxsCmdGearInVelDataT {
	if rcv == nil { return nil }
	t := &AxsCmdGearInVelDataT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCmdGearInVelData struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCmdGearInVelData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdGearInVelData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCmdGearInVelData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCmdGearInVelData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdGearInVelData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCmdGearInVelData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCmdGearInVelData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCmdGearInVelData) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the master axis
func (rcv *AxsCmdGearInVelData) Master() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the master axis
/// sync source (Actual/Setpoint)
func (rcv *AxsCmdGearInVelData) SyncSource() SyncSource {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return SyncSource(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// sync source (Actual/Setpoint)
func (rcv *AxsCmdGearInVelData) MutateSyncSource(n SyncSource) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

/// gear in vel parameters (master offset, slave offset, master ratio, slave ratio, fine adjust)
func (rcv *AxsCmdGearInVelData) Parameters(obj *AxsCmdGearInVelParams) *AxsCmdGearInVelParams {
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

/// gear in vel parameters (master offset, slave offset, master ratio, slave ratio, fine adjust)
/// should this be a buffered command?
func (rcv *AxsCmdGearInVelData) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// should this be a buffered command?
func (rcv *AxsCmdGearInVelData) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func AxsCmdGearInVelDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func AxsCmdGearInVelDataAddMaster(builder *flatbuffers.Builder, master flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(master), 0)
}
func AxsCmdGearInVelDataAddSyncSource(builder *flatbuffers.Builder, syncSource SyncSource) {
	builder.PrependInt8Slot(1, int8(syncSource), 0)
}
func AxsCmdGearInVelDataAddParameters(builder *flatbuffers.Builder, parameters flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(parameters), 0)
}
func AxsCmdGearInVelDataAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(3, buffered, false)
}
func AxsCmdGearInVelDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
