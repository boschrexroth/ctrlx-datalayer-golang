// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters for the command option ToolLengths (tool lenght correction) for kinematics
type KinCmdOptToolLengthDataT struct {
	PermType string `json:"permType"`
	SetName string `json:"setName"`
}

func (t *KinCmdOptToolLengthDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	permTypeOffset := flatbuffers.UOffsetT(0)
	if t.PermType != "" {
		permTypeOffset = builder.CreateString(t.PermType)
	}
	setNameOffset := flatbuffers.UOffsetT(0)
	if t.SetName != "" {
		setNameOffset = builder.CreateString(t.SetName)
	}
	KinCmdOptToolLengthDataStart(builder)
	KinCmdOptToolLengthDataAddPermType(builder, permTypeOffset)
	KinCmdOptToolLengthDataAddSetName(builder, setNameOffset)
	return KinCmdOptToolLengthDataEnd(builder)
}

func (rcv *KinCmdOptToolLengthData) UnPackTo(t *KinCmdOptToolLengthDataT) {
	t.PermType = string(rcv.PermType())
	t.SetName = string(rcv.SetName())
}

func (rcv *KinCmdOptToolLengthData) UnPack() *KinCmdOptToolLengthDataT {
	if rcv == nil { return nil }
	t := &KinCmdOptToolLengthDataT{}
	rcv.UnPackTo(t)
	return t
}

type KinCmdOptToolLengthData struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCmdOptToolLengthData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdOptToolLengthData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCmdOptToolLengthData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCmdOptToolLengthData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdOptToolLengthData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCmdOptToolLengthData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCmdOptToolLengthData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCmdOptToolLengthData) Table() flatbuffers.Table {
	return rcv._tab
}

/// permanent type (either "PermOn" or any other string to switch off)
func (rcv *KinCmdOptToolLengthData) PermType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// permanent type (either "PermOn" or any other string to switch off)
/// name of the tool that should become active
func (rcv *KinCmdOptToolLengthData) SetName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the tool that should become active
func KinCmdOptToolLengthDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func KinCmdOptToolLengthDataAddPermType(builder *flatbuffers.Builder, permType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(permType), 0)
}
func KinCmdOptToolLengthDataAddSetName(builder *flatbuffers.Builder, setName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(setName), 0)
}
func KinCmdOptToolLengthDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
