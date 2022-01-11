// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Data of all axes that are currently added to the kinematics
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