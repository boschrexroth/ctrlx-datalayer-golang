// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// get informations of a single active command
type debugCmdInfoT struct {
	CmdName string
	JobObjects []string
	State string
	CmdID uint64
	PrepLevel string
}

func (t *debugCmdInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	cmdNameOffset := builder.CreateString(t.CmdName)
	jobObjectsOffset := flatbuffers.UOffsetT(0)
	if t.JobObjects != nil {
		jobObjectsLength := len(t.JobObjects)
		jobObjectsOffsets := make([]flatbuffers.UOffsetT, jobObjectsLength)
		for j := 0; j < jobObjectsLength; j++ {
			jobObjectsOffsets[j] = builder.CreateString(t.JobObjects[j])
		}
		debugCmdInfoStartJobObjectsVector(builder, jobObjectsLength)
		for j := jobObjectsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(jobObjectsOffsets[j])
		}
		jobObjectsOffset = builder.EndVector(jobObjectsLength)
	}
	stateOffset := builder.CreateString(t.State)
	prepLevelOffset := builder.CreateString(t.PrepLevel)
	debugCmdInfoStart(builder)
	debugCmdInfoAddCmdName(builder, cmdNameOffset)
	debugCmdInfoAddJobObjects(builder, jobObjectsOffset)
	debugCmdInfoAddState(builder, stateOffset)
	debugCmdInfoAddCmdID(builder, t.CmdID)
	debugCmdInfoAddPrepLevel(builder, prepLevelOffset)
	return debugCmdInfoEnd(builder)
}

func (rcv *debugCmdInfo) UnPackTo(t *debugCmdInfoT) {
	t.CmdName = string(rcv.CmdName())
	jobObjectsLength := rcv.JobObjectsLength()
	t.JobObjects = make([]string, jobObjectsLength)
	for j := 0; j < jobObjectsLength; j++ {
		t.JobObjects[j] = string(rcv.JobObjects(j))
	}
	t.State = string(rcv.State())
	t.CmdID = rcv.CmdID()
	t.PrepLevel = string(rcv.PrepLevel())
}

func (rcv *debugCmdInfo) UnPack() *debugCmdInfoT {
	if rcv == nil { return nil }
	t := &debugCmdInfoT{}
	rcv.UnPackTo(t)
	return t
}

type debugCmdInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsdebugCmdInfo(buf []byte, offset flatbuffers.UOffsetT) *debugCmdInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &debugCmdInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsdebugCmdInfo(buf []byte, offset flatbuffers.UOffsetT) *debugCmdInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &debugCmdInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *debugCmdInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *debugCmdInfo) Table() flatbuffers.Table {
	return rcv._tab
}

/// command type name (e.g. PosAbs)
func (rcv *debugCmdInfo) CmdName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// command type name (e.g. PosAbs)
/// array of the strings of the jobObjects (including parameters)
func (rcv *debugCmdInfo) JobObjects(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *debugCmdInfo) JobObjectsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// array of the strings of the jobObjects (including parameters)
/// command state as string
func (rcv *debugCmdInfo) State() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// command state as string
/// command ID
func (rcv *debugCmdInfo) CmdID() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// command ID
func (rcv *debugCmdInfo) MutateCmdID(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

/// preparation level (in fact, a section name)
func (rcv *debugCmdInfo) PrepLevel() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// preparation level (in fact, a section name)
func debugCmdInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func debugCmdInfoAddCmdName(builder *flatbuffers.Builder, cmdName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(cmdName), 0)
}
func debugCmdInfoAddJobObjects(builder *flatbuffers.Builder, jobObjects flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(jobObjects), 0)
}
func debugCmdInfoStartJobObjectsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func debugCmdInfoAddState(builder *flatbuffers.Builder, state flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(state), 0)
}
func debugCmdInfoAddCmdID(builder *flatbuffers.Builder, cmdID uint64) {
	builder.PrependUint64Slot(3, cmdID, 0)
}
func debugCmdInfoAddPrepLevel(builder *flatbuffers.Builder, prepLevel flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(prepLevel), 0)
}
func debugCmdInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
