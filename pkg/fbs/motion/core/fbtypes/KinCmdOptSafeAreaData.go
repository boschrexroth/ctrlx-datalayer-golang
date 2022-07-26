// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters for the command option SafeArea (monitoring of safe zones and work areas) for kinematics
type KinCmdOptSafeAreaDataT struct {
	PermType string
	SafeArea string
}

func (t *KinCmdOptSafeAreaDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	permTypeOffset := builder.CreateString(t.PermType)
	safeAreaOffset := builder.CreateString(t.SafeArea)
	KinCmdOptSafeAreaDataStart(builder)
	KinCmdOptSafeAreaDataAddPermType(builder, permTypeOffset)
	KinCmdOptSafeAreaDataAddSafeArea(builder, safeAreaOffset)
	return KinCmdOptSafeAreaDataEnd(builder)
}

func (rcv *KinCmdOptSafeAreaData) UnPackTo(t *KinCmdOptSafeAreaDataT) {
	t.PermType = string(rcv.PermType())
	t.SafeArea = string(rcv.SafeArea())
}

func (rcv *KinCmdOptSafeAreaData) UnPack() *KinCmdOptSafeAreaDataT {
	if rcv == nil { return nil }
	t := &KinCmdOptSafeAreaDataT{}
	rcv.UnPackTo(t)
	return t
}

type KinCmdOptSafeAreaData struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCmdOptSafeAreaData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdOptSafeAreaData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCmdOptSafeAreaData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCmdOptSafeAreaData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdOptSafeAreaData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCmdOptSafeAreaData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCmdOptSafeAreaData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCmdOptSafeAreaData) Table() flatbuffers.Table {
	return rcv._tab
}

/// permanent type (e.g. "PermOn")
func (rcv *KinCmdOptSafeAreaData) PermType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// permanent type (e.g. "PermOn")
/// name of the safe zone or work area that should become active/disabled as set in the configuration
func (rcv *KinCmdOptSafeAreaData) SafeArea() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the safe zone or work area that should become active/disabled as set in the configuration
func KinCmdOptSafeAreaDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func KinCmdOptSafeAreaDataAddPermType(builder *flatbuffers.Builder, permType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(permType), 0)
}
func KinCmdOptSafeAreaDataAddSafeArea(builder *flatbuffers.Builder, safeArea flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(safeArea), 0)
}
func KinCmdOptSafeAreaDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
