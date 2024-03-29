// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configuration for axis estop to device error
type AxsCfgEstopSettingsT struct {
	EstopToNoDevErr EstopType `json:"estopToNoDevErr"`
	EstopToNoDeviceInfo EstopType `json:"estopToNoDeviceInfo"`
	EstopToDevLightErr EstopType `json:"estopToDevLightErr"`
	EstopToFieldbusErr EstopType `json:"estopToFieldbusErr"`
	EstopToDevCriticalErr EstopType `json:"estopToDevCriticalErr"`
	EstopToDevFatalErr EstopType `json:"estopToDevFatalErr"`
}

func (t *AxsCfgEstopSettingsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	AxsCfgEstopSettingsStart(builder)
	AxsCfgEstopSettingsAddEstopToNoDevErr(builder, t.EstopToNoDevErr)
	AxsCfgEstopSettingsAddEstopToNoDeviceInfo(builder, t.EstopToNoDeviceInfo)
	AxsCfgEstopSettingsAddEstopToDevLightErr(builder, t.EstopToDevLightErr)
	AxsCfgEstopSettingsAddEstopToFieldbusErr(builder, t.EstopToFieldbusErr)
	AxsCfgEstopSettingsAddEstopToDevCriticalErr(builder, t.EstopToDevCriticalErr)
	AxsCfgEstopSettingsAddEstopToDevFatalErr(builder, t.EstopToDevFatalErr)
	return AxsCfgEstopSettingsEnd(builder)
}

func (rcv *AxsCfgEstopSettings) UnPackTo(t *AxsCfgEstopSettingsT) {
	t.EstopToNoDevErr = rcv.EstopToNoDevErr()
	t.EstopToNoDeviceInfo = rcv.EstopToNoDeviceInfo()
	t.EstopToDevLightErr = rcv.EstopToDevLightErr()
	t.EstopToFieldbusErr = rcv.EstopToFieldbusErr()
	t.EstopToDevCriticalErr = rcv.EstopToDevCriticalErr()
	t.EstopToDevFatalErr = rcv.EstopToDevFatalErr()
}

func (rcv *AxsCfgEstopSettings) UnPack() *AxsCfgEstopSettingsT {
	if rcv == nil { return nil }
	t := &AxsCfgEstopSettingsT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCfgEstopSettings struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgEstopSettings(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgEstopSettings {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgEstopSettings{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgEstopSettings(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgEstopSettings {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgEstopSettings{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgEstopSettings) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgEstopSettings) Table() flatbuffers.Table {
	return rcv._tab
}

/// estop when device has no error
func (rcv *AxsCfgEstopSettings) EstopToNoDevErr() EstopType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return EstopType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// estop when device has no error
func (rcv *AxsCfgEstopSettings) MutateEstopToNoDevErr(n EstopType) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

/// estop when no device error information, e.g. virtual axis
func (rcv *AxsCfgEstopSettings) EstopToNoDeviceInfo() EstopType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return EstopType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// estop when no device error information, e.g. virtual axis
func (rcv *AxsCfgEstopSettings) MutateEstopToNoDeviceInfo(n EstopType) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

/// estop when get device light error
func (rcv *AxsCfgEstopSettings) EstopToDevLightErr() EstopType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return EstopType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// estop when get device light error
func (rcv *AxsCfgEstopSettings) MutateEstopToDevLightErr(n EstopType) bool {
	return rcv._tab.MutateInt8Slot(8, int8(n))
}

/// estop when get device light error
func (rcv *AxsCfgEstopSettings) EstopToFieldbusErr() EstopType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return EstopType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// estop when get device light error
func (rcv *AxsCfgEstopSettings) MutateEstopToFieldbusErr(n EstopType) bool {
	return rcv._tab.MutateInt8Slot(10, int8(n))
}

/// estop when get device critical error
func (rcv *AxsCfgEstopSettings) EstopToDevCriticalErr() EstopType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return EstopType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// estop when get device critical error
func (rcv *AxsCfgEstopSettings) MutateEstopToDevCriticalErr(n EstopType) bool {
	return rcv._tab.MutateInt8Slot(12, int8(n))
}

/// estop when get device fatal error
func (rcv *AxsCfgEstopSettings) EstopToDevFatalErr() EstopType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return EstopType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// estop when get device fatal error
func (rcv *AxsCfgEstopSettings) MutateEstopToDevFatalErr(n EstopType) bool {
	return rcv._tab.MutateInt8Slot(14, int8(n))
}

func AxsCfgEstopSettingsStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func AxsCfgEstopSettingsAddEstopToNoDevErr(builder *flatbuffers.Builder, estopToNoDevErr EstopType) {
	builder.PrependInt8Slot(0, int8(estopToNoDevErr), 0)
}
func AxsCfgEstopSettingsAddEstopToNoDeviceInfo(builder *flatbuffers.Builder, estopToNoDeviceInfo EstopType) {
	builder.PrependInt8Slot(1, int8(estopToNoDeviceInfo), 0)
}
func AxsCfgEstopSettingsAddEstopToDevLightErr(builder *flatbuffers.Builder, estopToDevLightErr EstopType) {
	builder.PrependInt8Slot(2, int8(estopToDevLightErr), 0)
}
func AxsCfgEstopSettingsAddEstopToFieldbusErr(builder *flatbuffers.Builder, estopToFieldbusErr EstopType) {
	builder.PrependInt8Slot(3, int8(estopToFieldbusErr), 0)
}
func AxsCfgEstopSettingsAddEstopToDevCriticalErr(builder *flatbuffers.Builder, estopToDevCriticalErr EstopType) {
	builder.PrependInt8Slot(4, int8(estopToDevCriticalErr), 0)
}
func AxsCfgEstopSettingsAddEstopToDevFatalErr(builder *flatbuffers.Builder, estopToDevFatalErr EstopType) {
	builder.PrependInt8Slot(5, int8(estopToDevFatalErr), 0)
}
func AxsCfgEstopSettingsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
