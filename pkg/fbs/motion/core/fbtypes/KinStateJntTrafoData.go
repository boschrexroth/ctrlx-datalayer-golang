// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// data of all registered joint transformations when reading all data of implemented joint transformations
type KinStateJntTrafoDataT struct {
	JntTrafoData []*KinStateJntTrafoDataSingleT `json:"jntTrafoData"`
}

func (t *KinStateJntTrafoDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	jntTrafoDataOffset := flatbuffers.UOffsetT(0)
	if t.JntTrafoData != nil {
		jntTrafoDataLength := len(t.JntTrafoData)
		jntTrafoDataOffsets := make([]flatbuffers.UOffsetT, jntTrafoDataLength)
		for j := 0; j < jntTrafoDataLength; j++ {
			jntTrafoDataOffsets[j] = t.JntTrafoData[j].Pack(builder)
		}
		KinStateJntTrafoDataStartJntTrafoDataVector(builder, jntTrafoDataLength)
		for j := jntTrafoDataLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(jntTrafoDataOffsets[j])
		}
		jntTrafoDataOffset = builder.EndVector(jntTrafoDataLength)
	}
	KinStateJntTrafoDataStart(builder)
	KinStateJntTrafoDataAddJntTrafoData(builder, jntTrafoDataOffset)
	return KinStateJntTrafoDataEnd(builder)
}

func (rcv *KinStateJntTrafoData) UnPackTo(t *KinStateJntTrafoDataT) {
	jntTrafoDataLength := rcv.JntTrafoDataLength()
	t.JntTrafoData = make([]*KinStateJntTrafoDataSingleT, jntTrafoDataLength)
	for j := 0; j < jntTrafoDataLength; j++ {
		x := KinStateJntTrafoDataSingle{}
		rcv.JntTrafoData(&x, j)
		t.JntTrafoData[j] = x.UnPack()
	}
}

func (rcv *KinStateJntTrafoData) UnPack() *KinStateJntTrafoDataT {
	if rcv == nil { return nil }
	t := &KinStateJntTrafoDataT{}
	rcv.UnPackTo(t)
	return t
}

type KinStateJntTrafoData struct {
	_tab flatbuffers.Table
}

func GetRootAsKinStateJntTrafoData(buf []byte, offset flatbuffers.UOffsetT) *KinStateJntTrafoData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinStateJntTrafoData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinStateJntTrafoData(buf []byte, offset flatbuffers.UOffsetT) *KinStateJntTrafoData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinStateJntTrafoData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinStateJntTrafoData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinStateJntTrafoData) Table() flatbuffers.Table {
	return rcv._tab
}

/// data of all registered joint transformations
func (rcv *KinStateJntTrafoData) JntTrafoData(obj *KinStateJntTrafoDataSingle, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *KinStateJntTrafoData) JntTrafoDataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// data of all registered joint transformations
func KinStateJntTrafoDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func KinStateJntTrafoDataAddJntTrafoData(builder *flatbuffers.Builder, jntTrafoData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(jntTrafoData), 0)
}
func KinStateJntTrafoDataStartJntTrafoDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KinStateJntTrafoDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
