// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configuration for reaction to device error
type AxsCfgDeviceErrorReaction struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgDeviceErrorReaction(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgDeviceErrorReaction {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgDeviceErrorReaction{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgDeviceErrorReaction(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgDeviceErrorReaction {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgDeviceErrorReaction{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgDeviceErrorReaction) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgDeviceErrorReaction) Table() flatbuffers.Table {
	return rcv._tab
}

/// user E-Stop dynamic limitations
func (rcv *AxsCfgDeviceErrorReaction) UserEstopLim(obj *AxsCfgEstopDynamicLimits) *AxsCfgEstopDynamicLimits {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AxsCfgEstopDynamicLimits)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// user E-Stop dynamic limitations
/// E-Stop setting to different device errors
func (rcv *AxsCfgDeviceErrorReaction) EstopSettings(obj *AxsCfgEstopSettings) *AxsCfgEstopSettings {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AxsCfgEstopSettings)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// E-Stop setting to different device errors
func AxsCfgDeviceErrorReactionStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AxsCfgDeviceErrorReactionAddUserEstopLim(builder *flatbuffers.Builder, userEstopLim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(userEstopLim), 0)
}
func AxsCfgDeviceErrorReactionAddEstopSettings(builder *flatbuffers.Builder, estopSettings flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(estopSettings), 0)
}
func AxsCfgDeviceErrorReactionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}