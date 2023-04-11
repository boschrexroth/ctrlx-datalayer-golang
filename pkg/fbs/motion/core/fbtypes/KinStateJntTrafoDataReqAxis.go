// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// data of a single entry of a required axis of a joint transformation when reading all data of an implemented joint transformation
type KinStateJntTrafoDataReqAxisT struct {
	Name string
	Description string
	Image string
	DocRef string
	Type KinStateJntTrafoDataReqAxisType
	Mandatory bool
}

func (t *KinStateJntTrafoDataReqAxisT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	descriptionOffset := builder.CreateString(t.Description)
	imageOffset := builder.CreateString(t.Image)
	docRefOffset := builder.CreateString(t.DocRef)
	KinStateJntTrafoDataReqAxisStart(builder)
	KinStateJntTrafoDataReqAxisAddName(builder, nameOffset)
	KinStateJntTrafoDataReqAxisAddDescription(builder, descriptionOffset)
	KinStateJntTrafoDataReqAxisAddImage(builder, imageOffset)
	KinStateJntTrafoDataReqAxisAddDocRef(builder, docRefOffset)
	KinStateJntTrafoDataReqAxisAddType(builder, t.Type)
	KinStateJntTrafoDataReqAxisAddMandatory(builder, t.Mandatory)
	return KinStateJntTrafoDataReqAxisEnd(builder)
}

func (rcv *KinStateJntTrafoDataReqAxis) UnPackTo(t *KinStateJntTrafoDataReqAxisT) {
	t.Name = string(rcv.Name())
	t.Description = string(rcv.Description())
	t.Image = string(rcv.Image())
	t.DocRef = string(rcv.DocRef())
	t.Type = rcv.Type()
	t.Mandatory = rcv.Mandatory()
}

func (rcv *KinStateJntTrafoDataReqAxis) UnPack() *KinStateJntTrafoDataReqAxisT {
	if rcv == nil { return nil }
	t := &KinStateJntTrafoDataReqAxisT{}
	rcv.UnPackTo(t)
	return t
}

type KinStateJntTrafoDataReqAxis struct {
	_tab flatbuffers.Table
}

func GetRootAsKinStateJntTrafoDataReqAxis(buf []byte, offset flatbuffers.UOffsetT) *KinStateJntTrafoDataReqAxis {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinStateJntTrafoDataReqAxis{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinStateJntTrafoDataReqAxis(buf []byte, offset flatbuffers.UOffsetT) *KinStateJntTrafoDataReqAxis {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinStateJntTrafoDataReqAxis{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinStateJntTrafoDataReqAxis) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinStateJntTrafoDataReqAxis) Table() flatbuffers.Table {
	return rcv._tab
}

/// Name of the axis (as used in the joint transformation)
func (rcv *KinStateJntTrafoDataReqAxis) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Name of the axis (as used in the joint transformation)
/// Description of the axis
func (rcv *KinStateJntTrafoDataReqAxis) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Description of the axis
/// Image (link) of the axis
func (rcv *KinStateJntTrafoDataReqAxis) Image() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Image (link) of the axis
/// Reference to the documentation of the axis
func (rcv *KinStateJntTrafoDataReqAxis) DocRef() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Reference to the documentation of the axis
/// Type of the axis (linear, rotational, any)
func (rcv *KinStateJntTrafoDataReqAxis) Type() KinStateJntTrafoDataReqAxisType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return KinStateJntTrafoDataReqAxisType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Type of the axis (linear, rotational, any)
func (rcv *KinStateJntTrafoDataReqAxis) MutateType(n KinStateJntTrafoDataReqAxisType) bool {
	return rcv._tab.MutateInt8Slot(12, int8(n))
}

/// Is this axis mandatory?
func (rcv *KinStateJntTrafoDataReqAxis) Mandatory() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Is this axis mandatory?
func (rcv *KinStateJntTrafoDataReqAxis) MutateMandatory(n bool) bool {
	return rcv._tab.MutateBoolSlot(14, n)
}

func KinStateJntTrafoDataReqAxisStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func KinStateJntTrafoDataReqAxisAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func KinStateJntTrafoDataReqAxisAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(description), 0)
}
func KinStateJntTrafoDataReqAxisAddImage(builder *flatbuffers.Builder, image flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(image), 0)
}
func KinStateJntTrafoDataReqAxisAddDocRef(builder *flatbuffers.Builder, docRef flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(docRef), 0)
}
func KinStateJntTrafoDataReqAxisAddType(builder *flatbuffers.Builder, type_ KinStateJntTrafoDataReqAxisType) {
	builder.PrependInt8Slot(4, int8(type_), 0)
}
func KinStateJntTrafoDataReqAxisAddMandatory(builder *flatbuffers.Builder, mandatory bool) {
	builder.PrependBoolSlot(5, mandatory, false)
}
func KinStateJntTrafoDataReqAxisEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
