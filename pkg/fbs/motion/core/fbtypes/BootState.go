// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// return type of requests of the current boot state
type BootState struct {
	_tab flatbuffers.Table
}

func GetRootAsBootState(buf []byte, offset flatbuffers.UOffsetT) *BootState {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BootState{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBootState(buf []byte, offset flatbuffers.UOffsetT) *BootState {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BootState{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *BootState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BootState) Table() flatbuffers.Table {
	return rcv._tab
}

/// text of the boot step (can be shown in an HMI)
func (rcv *BootState) Text() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// text of the boot step (can be shown in an HMI)
/// current boot step (starts with 0, ends with maxSteps)
func (rcv *BootState) ActStep() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// current boot step (starts with 0, ends with maxSteps)
func (rcv *BootState) MutateActStep(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

/// number of boot steps (when actStep == maxSteps then the booting is finished)
func (rcv *BootState) MaxSteps() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// number of boot steps (when actStep == maxSteps then the booting is finished)
func (rcv *BootState) MutateMaxSteps(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func BootStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func BootStateAddText(builder *flatbuffers.Builder, text flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(text), 0)
}
func BootStateAddActStep(builder *flatbuffers.Builder, actStep uint32) {
	builder.PrependUint32Slot(1, actStep, 0)
}
func BootStateAddMaxSteps(builder *flatbuffers.Builder, maxSteps uint32) {
	builder.PrependUint32Slot(2, maxSteps, 0)
}
func BootStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}