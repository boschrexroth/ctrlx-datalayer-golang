// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters for the contour command for a kinematics
type KinCmdContourDataT struct {
	IsStart bool `json:"isStart"`
	PrepCmds uint32 `json:"prepCmds"`
}

func (t *KinCmdContourDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	KinCmdContourDataStart(builder)
	KinCmdContourDataAddIsStart(builder, t.IsStart)
	KinCmdContourDataAddPrepCmds(builder, t.PrepCmds)
	return KinCmdContourDataEnd(builder)
}

func (rcv *KinCmdContourData) UnPackTo(t *KinCmdContourDataT) {
	t.IsStart = rcv.IsStart()
	t.PrepCmds = rcv.PrepCmds()
}

func (rcv *KinCmdContourData) UnPack() *KinCmdContourDataT {
	if rcv == nil { return nil }
	t := &KinCmdContourDataT{}
	rcv.UnPackTo(t)
	return t
}

type KinCmdContourData struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCmdContourData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdContourData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCmdContourData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCmdContourData(buf []byte, offset flatbuffers.UOffsetT) *KinCmdContourData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCmdContourData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCmdContourData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCmdContourData) Table() flatbuffers.Table {
	return rcv._tab
}

/// is this the start of the contour? 
func (rcv *KinCmdContourData) IsStart() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

/// is this the start of the contour? 
func (rcv *KinCmdContourData) MutateIsStart(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

/// set the number of commands that should be prepared completely (only relevant when isStart=true)
func (rcv *KinCmdContourData) PrepCmds() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// set the number of commands that should be prepared completely (only relevant when isStart=true)
func (rcv *KinCmdContourData) MutatePrepCmds(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func KinCmdContourDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func KinCmdContourDataAddIsStart(builder *flatbuffers.Builder, isStart bool) {
	builder.PrependBoolSlot(0, isStart, true)
}
func KinCmdContourDataAddPrepCmds(builder *flatbuffers.Builder, prepCmds uint32) {
	builder.PrependUint32Slot(1, prepCmds, 0)
}
func KinCmdContourDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
