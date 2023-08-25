// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// state information of the execution
type AxsStateFlexProfileExecutionT struct {
	ActiveProfile string `json:"activeProfile"`
	ActiveSegment uint32 `json:"activeSegment"`
	WaitForSwitchPos bool `json:"waitForSwitchPos"`
	SingleExecDone bool `json:"singleExecDone"`
}

func (t *AxsStateFlexProfileExecutionT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	activeProfileOffset := flatbuffers.UOffsetT(0)
	if t.ActiveProfile != "" {
		activeProfileOffset = builder.CreateString(t.ActiveProfile)
	}
	AxsStateFlexProfileExecutionStart(builder)
	AxsStateFlexProfileExecutionAddActiveProfile(builder, activeProfileOffset)
	AxsStateFlexProfileExecutionAddActiveSegment(builder, t.ActiveSegment)
	AxsStateFlexProfileExecutionAddWaitForSwitchPos(builder, t.WaitForSwitchPos)
	AxsStateFlexProfileExecutionAddSingleExecDone(builder, t.SingleExecDone)
	return AxsStateFlexProfileExecutionEnd(builder)
}

func (rcv *AxsStateFlexProfileExecution) UnPackTo(t *AxsStateFlexProfileExecutionT) {
	t.ActiveProfile = string(rcv.ActiveProfile())
	t.ActiveSegment = rcv.ActiveSegment()
	t.WaitForSwitchPos = rcv.WaitForSwitchPos()
	t.SingleExecDone = rcv.SingleExecDone()
}

func (rcv *AxsStateFlexProfileExecution) UnPack() *AxsStateFlexProfileExecutionT {
	if rcv == nil { return nil }
	t := &AxsStateFlexProfileExecutionT{}
	rcv.UnPackTo(t)
	return t
}

type AxsStateFlexProfileExecution struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsStateFlexProfileExecution(buf []byte, offset flatbuffers.UOffsetT) *AxsStateFlexProfileExecution {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsStateFlexProfileExecution{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsStateFlexProfileExecution(buf []byte, offset flatbuffers.UOffsetT) *AxsStateFlexProfileExecution {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsStateFlexProfileExecution{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsStateFlexProfileExecution) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsStateFlexProfileExecution) Table() flatbuffers.Table {
	return rcv._tab
}

/// currently active profile from the axis
func (rcv *AxsStateFlexProfileExecution) ActiveProfile() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// currently active profile from the axis
/// currently active segment of the profile from the axis
func (rcv *AxsStateFlexProfileExecution) ActiveSegment() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// currently active segment of the profile from the axis
func (rcv *AxsStateFlexProfileExecution) MutateActiveSegment(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

/// the profile is waiting for the specified switching position
func (rcv *AxsStateFlexProfileExecution) WaitForSwitchPos() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// the profile is waiting for the specified switching position
func (rcv *AxsStateFlexProfileExecution) MutateWaitForSwitchPos(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

/// the one-time execution of the profile was completed (only for execution mode "SINGLE")
func (rcv *AxsStateFlexProfileExecution) SingleExecDone() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// the one-time execution of the profile was completed (only for execution mode "SINGLE")
func (rcv *AxsStateFlexProfileExecution) MutateSingleExecDone(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func AxsStateFlexProfileExecutionStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func AxsStateFlexProfileExecutionAddActiveProfile(builder *flatbuffers.Builder, activeProfile flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(activeProfile), 0)
}
func AxsStateFlexProfileExecutionAddActiveSegment(builder *flatbuffers.Builder, activeSegment uint32) {
	builder.PrependUint32Slot(1, activeSegment, 0)
}
func AxsStateFlexProfileExecutionAddWaitForSwitchPos(builder *flatbuffers.Builder, waitForSwitchPos bool) {
	builder.PrependBoolSlot(2, waitForSwitchPos, false)
}
func AxsStateFlexProfileExecutionAddSingleExecDone(builder *flatbuffers.Builder, singleExecDone bool) {
	builder.PrependBoolSlot(3, singleExecDone, false)
}
func AxsStateFlexProfileExecutionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
