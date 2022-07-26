// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters for the command option PCS (product coordinate system) for kinematics
type KinCmdOptPCSDataT struct {
	PermType string
	SetName string
}

func (t *KinCmdOptPCSDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	permTypeOffset := builder.CreateString(t.PermType)
	setNameOffset := builder.CreateString(t.SetName)
	KinCmdOptPCSDataStart(builder)
	KinCmdOptPCSDataAddPermType(builder, permTypeOffset)
	KinCmdOptPCSDataAddSetName(builder, setNameOffset)
	return KinCmdOptPCSDataEnd(builder)
}

func (rcv *KinCmdOptPCSData) UnPackTo(t *KinCmdOptPCSDataT) {
	t.PermType = string(rcv.PermType())
	t.SetName = string(rcv.SetName())
}

func (rcv *KinCmdOptPCSData) UnPack() *KinCmdOptPCSDataT {
	if rcv == nil { return nil }
	t := &KinCmdOptPCSDataT{}
	rcv.UnPackTo(t)
	return t
}

type KinCmdOptPCSData struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCmdOptPCSData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdOptPCSData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCmdOptPCSData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCmdOptPCSData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdOptPCSData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCmdOptPCSData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCmdOptPCSData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCmdOptPCSData) Table() flatbuffers.Table {
	return rcv._tab
}

/// permanent type (e.g. "PermOn")
func (rcv *KinCmdOptPCSData) PermType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// permanent type (e.g. "PermOn")
/// name of the set/group that should become active
func (rcv *KinCmdOptPCSData) SetName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the set/group that should become active
func KinCmdOptPCSDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func KinCmdOptPCSDataAddPermType(builder *flatbuffers.Builder, permType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(permType), 0)
}
func KinCmdOptPCSDataAddSetName(builder *flatbuffers.Builder, setName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(setName), 0)
}
func KinCmdOptPCSDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
