// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Data of all axes that are currently added to the kinematics
type KinAxsAllT struct {
	Info []*KinAxsSingleT `json:"info"`
}

func (t *KinAxsAllT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	infoOffset := flatbuffers.UOffsetT(0)
	if t.Info != nil {
		infoLength := len(t.Info)
		infoOffsets := make([]flatbuffers.UOffsetT, infoLength)
		for j := 0; j < infoLength; j++ {
			infoOffsets[j] = t.Info[j].Pack(builder)
		}
		KinAxsAllStartInfoVector(builder, infoLength)
		for j := infoLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(infoOffsets[j])
		}
		infoOffset = builder.EndVector(infoLength)
	}
	KinAxsAllStart(builder)
	KinAxsAllAddInfo(builder, infoOffset)
	return KinAxsAllEnd(builder)
}

func (rcv *KinAxsAll) UnPackTo(t *KinAxsAllT) {
	infoLength := rcv.InfoLength()
	t.Info = make([]*KinAxsSingleT, infoLength)
	for j := 0; j < infoLength; j++ {
		x := KinAxsSingle{}
		rcv.Info(&x, j)
		t.Info[j] = x.UnPack()
	}
}

func (rcv *KinAxsAll) UnPack() *KinAxsAllT {
	if rcv == nil { return nil }
	t := &KinAxsAllT{}
	rcv.UnPackTo(t)
	return t
}

type KinAxsAll struct {
	_tab flatbuffers.Table
}

func GetRootAsKinAxsAll(buf []byte, offset flatbuffers.UOffsetT) *KinAxsAll {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinAxsAll{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinAxsAll(buf []byte, offset flatbuffers.UOffsetT) *KinAxsAll {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinAxsAll{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinAxsAll) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinAxsAll) Table() flatbuffers.Table {
	return rcv._tab
}

/// vector of all axes that are currently added to the kinematics
func (rcv *KinAxsAll) Info(obj *KinAxsSingle, j int) bool {
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

func (rcv *KinAxsAll) InfoByKey(obj *KinAxsSingle, key string) bool{
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *KinAxsAll) InfoLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of all axes that are currently added to the kinematics
func KinAxsAllStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func KinAxsAllAddInfo(builder *flatbuffers.Builder, info flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(info), 0)
}
func KinAxsAllStartInfoVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KinAxsAllEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
