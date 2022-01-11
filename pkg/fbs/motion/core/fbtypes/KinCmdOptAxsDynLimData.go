// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters for the command option SafeArea (monitoring of safe zones and work areas) for kinematics
type KinCmdOptAxsDynLimData struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCmdOptAxsDynLimData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdOptAxsDynLimData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCmdOptAxsDynLimData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCmdOptAxsDynLimData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdOptAxsDynLimData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCmdOptAxsDynLimData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCmdOptAxsDynLimData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCmdOptAxsDynLimData) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the kinematics axis, that dynamic limits should be reduced
func (rcv *KinCmdOptAxsDynLimData) AxsName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the kinematics axis, that dynamic limits should be reduced
/// dynamic limits for the following motion commands for this axis (optional)
/// leave it out to disable the command option
func (rcv *KinCmdOptAxsDynLimData) Lim(obj *DynamicLimits) *DynamicLimits {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(DynamicLimits)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// dynamic limits for the following motion commands for this axis (optional)
/// leave it out to disable the command option
func KinCmdOptAxsDynLimDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func KinCmdOptAxsDynLimDataAddAxsName(builder *flatbuffers.Builder, axsName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(axsName), 0)
}
func KinCmdOptAxsDynLimDataAddLim(builder *flatbuffers.Builder, lim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(lim), 0)
}
func KinCmdOptAxsDynLimDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}