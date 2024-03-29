// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// internal system configuration items
type SysCfgInternalT struct {
	Save bool `json:"save"`
	UseIpoIvaj bool `json:"useIpoIVAJ"`
	TimeMeasurement bool `json:"timeMeasurement"`
	DriveResetTimeout float64 `json:"driveResetTimeout"`
}

func (t *SysCfgInternalT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	SysCfgInternalStart(builder)
	SysCfgInternalAddSave(builder, t.Save)
	SysCfgInternalAddUseIpoIvaj(builder, t.UseIpoIvaj)
	SysCfgInternalAddTimeMeasurement(builder, t.TimeMeasurement)
	SysCfgInternalAddDriveResetTimeout(builder, t.DriveResetTimeout)
	return SysCfgInternalEnd(builder)
}

func (rcv *SysCfgInternal) UnPackTo(t *SysCfgInternalT) {
	t.Save = rcv.Save()
	t.UseIpoIvaj = rcv.UseIpoIvaj()
	t.TimeMeasurement = rcv.TimeMeasurement()
	t.DriveResetTimeout = rcv.DriveResetTimeout()
}

func (rcv *SysCfgInternal) UnPack() *SysCfgInternalT {
	if rcv == nil { return nil }
	t := &SysCfgInternalT{}
	rcv.UnPackTo(t)
	return t
}

type SysCfgInternal struct {
	_tab flatbuffers.Table
}

func GetRootAsSysCfgInternal(buf []byte, offset flatbuffers.UOffsetT) *SysCfgInternal {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SysCfgInternal{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSysCfgInternal(buf []byte, offset flatbuffers.UOffsetT) *SysCfgInternal {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SysCfgInternal{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SysCfgInternal) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SysCfgInternal) Table() flatbuffers.Table {
	return rcv._tab
}

/// save the internal system configuration items also, when saving the system config to file?
func (rcv *SysCfgInternal) Save() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// save the internal system configuration items also, when saving the system config to file?
func (rcv *SysCfgInternal) MutateSave(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

/// use the new ipo IVAJ?
func (rcv *SysCfgInternal) UseIpoIvaj() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// use the new ipo IVAJ?
func (rcv *SysCfgInternal) MutateUseIpoIvaj(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

/// enable time measurement?
func (rcv *SysCfgInternal) TimeMeasurement() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// enable time measurement?
func (rcv *SysCfgInternal) MutateTimeMeasurement(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

/// Change the timeout for the drive reset
func (rcv *SysCfgInternal) DriveResetTimeout() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 4.0
}

/// Change the timeout for the drive reset
func (rcv *SysCfgInternal) MutateDriveResetTimeout(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

func SysCfgInternalStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func SysCfgInternalAddSave(builder *flatbuffers.Builder, save bool) {
	builder.PrependBoolSlot(0, save, false)
}
func SysCfgInternalAddUseIpoIvaj(builder *flatbuffers.Builder, useIpoIvaj bool) {
	builder.PrependBoolSlot(1, useIpoIvaj, false)
}
func SysCfgInternalAddTimeMeasurement(builder *flatbuffers.Builder, timeMeasurement bool) {
	builder.PrependBoolSlot(2, timeMeasurement, false)
}
func SysCfgInternalAddDriveResetTimeout(builder *flatbuffers.Builder, driveResetTimeout float64) {
	builder.PrependFloat64Slot(3, driveResetTimeout, 4.0)
}
func SysCfgInternalEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
