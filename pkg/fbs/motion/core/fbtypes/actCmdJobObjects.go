// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// get informations of a single active command
type actCmdJobObjectsT struct {
	CmdName string `json:"cmdName"`
	JobObjects []string `json:"jobObjects"`
}

func (t *actCmdJobObjectsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	cmdNameOffset := flatbuffers.UOffsetT(0)
	if t.CmdName != "" {
		cmdNameOffset = builder.CreateString(t.CmdName)
	}
	jobObjectsOffset := flatbuffers.UOffsetT(0)
	if t.JobObjects != nil {
		jobObjectsLength := len(t.JobObjects)
		jobObjectsOffsets := make([]flatbuffers.UOffsetT, jobObjectsLength)
		for j := 0; j < jobObjectsLength; j++ {
			jobObjectsOffsets[j] = builder.CreateString(t.JobObjects[j])
		}
		actCmdJobObjectsStartJobObjectsVector(builder, jobObjectsLength)
		for j := jobObjectsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(jobObjectsOffsets[j])
		}
		jobObjectsOffset = builder.EndVector(jobObjectsLength)
	}
	actCmdJobObjectsStart(builder)
	actCmdJobObjectsAddCmdName(builder, cmdNameOffset)
	actCmdJobObjectsAddJobObjects(builder, jobObjectsOffset)
	return actCmdJobObjectsEnd(builder)
}

func (rcv *actCmdJobObjects) UnPackTo(t *actCmdJobObjectsT) {
	t.CmdName = string(rcv.CmdName())
	jobObjectsLength := rcv.JobObjectsLength()
	t.JobObjects = make([]string, jobObjectsLength)
	for j := 0; j < jobObjectsLength; j++ {
		t.JobObjects[j] = string(rcv.JobObjects(j))
	}
}

func (rcv *actCmdJobObjects) UnPack() *actCmdJobObjectsT {
	if rcv == nil { return nil }
	t := &actCmdJobObjectsT{}
	rcv.UnPackTo(t)
	return t
}

type actCmdJobObjects struct {
	_tab flatbuffers.Table
}

func GetRootAsactCmdJobObjects(buf []byte, offset flatbuffers.UOffsetT) *actCmdJobObjects {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &actCmdJobObjects{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsactCmdJobObjects(buf []byte, offset flatbuffers.UOffsetT) *actCmdJobObjects {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &actCmdJobObjects{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *actCmdJobObjects) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *actCmdJobObjects) Table() flatbuffers.Table {
	return rcv._tab
}

/// command type name (e.g. PosAbs)
func (rcv *actCmdJobObjects) CmdName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// command type name (e.g. PosAbs)
/// array of the strings of the jobObjects (including parameters)
func (rcv *actCmdJobObjects) JobObjects(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *actCmdJobObjects) JobObjectsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// array of the strings of the jobObjects (including parameters)
func actCmdJobObjectsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func actCmdJobObjectsAddCmdName(builder *flatbuffers.Builder, cmdName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(cmdName), 0)
}
func actCmdJobObjectsAddJobObjects(builder *flatbuffers.Builder, jobObjects flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(jobObjects), 0)
}
func actCmdJobObjectsStartJobObjectsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func actCmdJobObjectsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
