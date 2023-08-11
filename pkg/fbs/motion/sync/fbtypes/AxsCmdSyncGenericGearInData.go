// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters of the axis generic GearIn command
type AxsCmdSyncGenericGearInDataT struct {
	Master string `json:"master"`
	Pipeline string `json:"pipeline"`
	SyncSource SyncSource `json:"syncSource"`
	Buffered bool `json:"buffered"`
}

func (t *AxsCmdSyncGenericGearInDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	masterOffset := flatbuffers.UOffsetT(0)
	if t.Master != "" {
		masterOffset = builder.CreateString(t.Master)
	}
	pipelineOffset := flatbuffers.UOffsetT(0)
	if t.Pipeline != "" {
		pipelineOffset = builder.CreateString(t.Pipeline)
	}
	AxsCmdSyncGenericGearInDataStart(builder)
	AxsCmdSyncGenericGearInDataAddMaster(builder, masterOffset)
	AxsCmdSyncGenericGearInDataAddPipeline(builder, pipelineOffset)
	AxsCmdSyncGenericGearInDataAddSyncSource(builder, t.SyncSource)
	AxsCmdSyncGenericGearInDataAddBuffered(builder, t.Buffered)
	return AxsCmdSyncGenericGearInDataEnd(builder)
}

func (rcv *AxsCmdSyncGenericGearInData) UnPackTo(t *AxsCmdSyncGenericGearInDataT) {
	t.Master = string(rcv.Master())
	t.Pipeline = string(rcv.Pipeline())
	t.SyncSource = rcv.SyncSource()
	t.Buffered = rcv.Buffered()
}

func (rcv *AxsCmdSyncGenericGearInData) UnPack() *AxsCmdSyncGenericGearInDataT {
	if rcv == nil { return nil }
	t := &AxsCmdSyncGenericGearInDataT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCmdSyncGenericGearInData struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCmdSyncGenericGearInData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdSyncGenericGearInData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCmdSyncGenericGearInData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCmdSyncGenericGearInData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdSyncGenericGearInData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCmdSyncGenericGearInData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCmdSyncGenericGearInData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCmdSyncGenericGearInData) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the master axis
func (rcv *AxsCmdSyncGenericGearInData) Master() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the master axis
/// name of the pipeline
func (rcv *AxsCmdSyncGenericGearInData) Pipeline() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the pipeline
/// Sync source
func (rcv *AxsCmdSyncGenericGearInData) SyncSource() SyncSource {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return SyncSource(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Sync source
func (rcv *AxsCmdSyncGenericGearInData) MutateSyncSource(n SyncSource) bool {
	return rcv._tab.MutateInt8Slot(8, int8(n))
}

/// should this be a buffered command?
func (rcv *AxsCmdSyncGenericGearInData) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// should this be a buffered command?
func (rcv *AxsCmdSyncGenericGearInData) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func AxsCmdSyncGenericGearInDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func AxsCmdSyncGenericGearInDataAddMaster(builder *flatbuffers.Builder, master flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(master), 0)
}
func AxsCmdSyncGenericGearInDataAddPipeline(builder *flatbuffers.Builder, pipeline flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(pipeline), 0)
}
func AxsCmdSyncGenericGearInDataAddSyncSource(builder *flatbuffers.Builder, syncSource SyncSource) {
	builder.PrependInt8Slot(2, int8(syncSource), 0)
}
func AxsCmdSyncGenericGearInDataAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(3, buffered, false)
}
func AxsCmdSyncGenericGearInDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
