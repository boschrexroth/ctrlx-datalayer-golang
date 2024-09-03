// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// data of a single entry of a required axis of a joint transformation when reading all data of an implemented joint transformation
type KinStateJntTrafoDataReqAxisT struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	DocRef string `json:"docRef"`
	Type KinStateJntTrafoDataReqAxisType `json:"type"`
	Mandatory bool `json:"mandatory"`
	AddInfo []*KVPCfgSingleItemT `json:"addInfo"`
}

func (t *KinStateJntTrafoDataReqAxisT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	descriptionOffset := flatbuffers.UOffsetT(0)
	if t.Description != "" {
		descriptionOffset = builder.CreateString(t.Description)
	}
	imageOffset := flatbuffers.UOffsetT(0)
	if t.Image != "" {
		imageOffset = builder.CreateString(t.Image)
	}
	docRefOffset := flatbuffers.UOffsetT(0)
	if t.DocRef != "" {
		docRefOffset = builder.CreateString(t.DocRef)
	}
	addInfoOffset := flatbuffers.UOffsetT(0)
	if t.AddInfo != nil {
		addInfoLength := len(t.AddInfo)
		addInfoOffsets := make([]flatbuffers.UOffsetT, addInfoLength)
		for j := 0; j < addInfoLength; j++ {
			addInfoOffsets[j] = t.AddInfo[j].Pack(builder)
		}
		KinStateJntTrafoDataReqAxisStartAddInfoVector(builder, addInfoLength)
		for j := addInfoLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(addInfoOffsets[j])
		}
		addInfoOffset = builder.EndVector(addInfoLength)
	}
	KinStateJntTrafoDataReqAxisStart(builder)
	KinStateJntTrafoDataReqAxisAddName(builder, nameOffset)
	KinStateJntTrafoDataReqAxisAddDescription(builder, descriptionOffset)
	KinStateJntTrafoDataReqAxisAddImage(builder, imageOffset)
	KinStateJntTrafoDataReqAxisAddDocRef(builder, docRefOffset)
	KinStateJntTrafoDataReqAxisAddType(builder, t.Type)
	KinStateJntTrafoDataReqAxisAddMandatory(builder, t.Mandatory)
	KinStateJntTrafoDataReqAxisAddAddInfo(builder, addInfoOffset)
	return KinStateJntTrafoDataReqAxisEnd(builder)
}

func (rcv *KinStateJntTrafoDataReqAxis) UnPackTo(t *KinStateJntTrafoDataReqAxisT) {
	t.Name = string(rcv.Name())
	t.Description = string(rcv.Description())
	t.Image = string(rcv.Image())
	t.DocRef = string(rcv.DocRef())
	t.Type = rcv.Type()
	t.Mandatory = rcv.Mandatory()
	addInfoLength := rcv.AddInfoLength()
	t.AddInfo = make([]*KVPCfgSingleItemT, addInfoLength)
	for j := 0; j < addInfoLength; j++ {
		x := KVPCfgSingleItem{}
		rcv.AddInfo(&x, j)
		t.AddInfo[j] = x.UnPack()
	}
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

/// Additional information of the axis (as key-value-pair)
func (rcv *KinStateJntTrafoDataReqAxis) AddInfo(obj *KVPCfgSingleItem, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *KinStateJntTrafoDataReqAxis) AddInfoLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Additional information of the axis (as key-value-pair)
func KinStateJntTrafoDataReqAxisStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
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
func KinStateJntTrafoDataReqAxisAddAddInfo(builder *flatbuffers.Builder, addInfo flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(addInfo), 0)
}
func KinStateJntTrafoDataReqAxisStartAddInfoVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KinStateJntTrafoDataReqAxisEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
