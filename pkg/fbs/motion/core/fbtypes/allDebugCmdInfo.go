// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// get informations of all active commands and commands that were recently executed of a single motion object
type allDebugCmdInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsallDebugCmdInfo(buf []byte, offset flatbuffers.UOffsetT) *allDebugCmdInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &allDebugCmdInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsallDebugCmdInfo(buf []byte, offset flatbuffers.UOffsetT) *allDebugCmdInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &allDebugCmdInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *allDebugCmdInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *allDebugCmdInfo) Table() flatbuffers.Table {
	return rcv._tab
}

/// vector of all main motion commands (starting with the most recent active command)
func (rcv *allDebugCmdInfo) MainCmds(obj *debugCmdInfo, j int) bool {
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

func (rcv *allDebugCmdInfo) MainCmdsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of all main motion commands (starting with the most recent active command)
/// vector of all additional active motion commands
func (rcv *allDebugCmdInfo) AddCmds(obj *debugCmdInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *allDebugCmdInfo) AddCmdsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of all additional active motion commands
func allDebugCmdInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func allDebugCmdInfoAddMainCmds(builder *flatbuffers.Builder, mainCmds flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(mainCmds), 0)
}
func allDebugCmdInfoStartMainCmdsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func allDebugCmdInfoAddAddCmds(builder *flatbuffers.Builder, addCmds flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(addCmds), 0)
}
func allDebugCmdInfoStartAddCmdsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func allDebugCmdInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
