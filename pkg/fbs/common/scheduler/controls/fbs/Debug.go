// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Machine ID to use the DEBUG configuration only on the intended machine
type DebugT struct {
	Machine string `json:"machine"`
}

func (t *DebugT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	machineOffset := flatbuffers.UOffsetT(0)
	if t.Machine != "" {
		machineOffset = builder.CreateString(t.Machine)
	}
	DebugStart(builder)
	DebugAddMachine(builder, machineOffset)
	return DebugEnd(builder)
}

func (rcv *Debug) UnPackTo(t *DebugT) {
	t.Machine = string(rcv.Machine())
}

func (rcv *Debug) UnPack() *DebugT {
	if rcv == nil { return nil }
	t := &DebugT{}
	rcv.UnPackTo(t)
	return t
}

type Debug struct {
	_tab flatbuffers.Table
}

func GetRootAsDebug(buf []byte, offset flatbuffers.UOffsetT) *Debug {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Debug{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDebug(buf []byte, offset flatbuffers.UOffsetT) *Debug {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Debug{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Debug) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Debug) Table() flatbuffers.Table {
	return rcv._tab
}

/// Machine ID to use the DEBUG configuration only on the intended machine
func (rcv *Debug) Machine() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Machine ID to use the DEBUG configuration only on the intended machine
func DebugStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DebugAddMachine(builder *flatbuffers.Builder, machine flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(machine), 0)
}
func DebugEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
