// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// state of a FlexProfile object
type AxsStateFlexProfileObjT struct {
	Profiles []*AxsStateSingleFlexProfileT `json:"profiles"`
	Events *AxsStateFlexProfileAllEventsT `json:"events"`
	Execution *AxsStateFlexProfileExecutionT `json:"execution"`
}

func (t *AxsStateFlexProfileObjT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	profilesOffset := flatbuffers.UOffsetT(0)
	if t.Profiles != nil {
		profilesLength := len(t.Profiles)
		profilesOffsets := make([]flatbuffers.UOffsetT, profilesLength)
		for j := 0; j < profilesLength; j++ {
			profilesOffsets[j] = t.Profiles[j].Pack(builder)
		}
		AxsStateFlexProfileObjStartProfilesVector(builder, profilesLength)
		for j := profilesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(profilesOffsets[j])
		}
		profilesOffset = builder.EndVector(profilesLength)
	}
	eventsOffset := t.Events.Pack(builder)
	executionOffset := t.Execution.Pack(builder)
	AxsStateFlexProfileObjStart(builder)
	AxsStateFlexProfileObjAddProfiles(builder, profilesOffset)
	AxsStateFlexProfileObjAddEvents(builder, eventsOffset)
	AxsStateFlexProfileObjAddExecution(builder, executionOffset)
	return AxsStateFlexProfileObjEnd(builder)
}

func (rcv *AxsStateFlexProfileObj) UnPackTo(t *AxsStateFlexProfileObjT) {
	profilesLength := rcv.ProfilesLength()
	t.Profiles = make([]*AxsStateSingleFlexProfileT, profilesLength)
	for j := 0; j < profilesLength; j++ {
		x := AxsStateSingleFlexProfile{}
		rcv.Profiles(&x, j)
		t.Profiles[j] = x.UnPack()
	}
	t.Events = rcv.Events(nil).UnPack()
	t.Execution = rcv.Execution(nil).UnPack()
}

func (rcv *AxsStateFlexProfileObj) UnPack() *AxsStateFlexProfileObjT {
	if rcv == nil { return nil }
	t := &AxsStateFlexProfileObjT{}
	rcv.UnPackTo(t)
	return t
}

type AxsStateFlexProfileObj struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsStateFlexProfileObj(buf []byte, offset flatbuffers.UOffsetT) *AxsStateFlexProfileObj {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsStateFlexProfileObj{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsStateFlexProfileObj(buf []byte, offset flatbuffers.UOffsetT) *AxsStateFlexProfileObj {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsStateFlexProfileObj{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsStateFlexProfileObj) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsStateFlexProfileObj) Table() flatbuffers.Table {
	return rcv._tab
}

/// state information of single FlexProfiles
func (rcv *AxsStateFlexProfileObj) Profiles(obj *AxsStateSingleFlexProfile, j int) bool {
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

func (rcv *AxsStateFlexProfileObj) ProfilesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// state information of single FlexProfiles
/// state information of all events
func (rcv *AxsStateFlexProfileObj) Events(obj *AxsStateFlexProfileAllEvents) *AxsStateFlexProfileAllEvents {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AxsStateFlexProfileAllEvents)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// state information of all events
/// state information of execution
func (rcv *AxsStateFlexProfileObj) Execution(obj *AxsStateFlexProfileExecution) *AxsStateFlexProfileExecution {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AxsStateFlexProfileExecution)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// state information of execution
func AxsStateFlexProfileObjStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func AxsStateFlexProfileObjAddProfiles(builder *flatbuffers.Builder, profiles flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(profiles), 0)
}
func AxsStateFlexProfileObjStartProfilesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func AxsStateFlexProfileObjAddEvents(builder *flatbuffers.Builder, events flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(events), 0)
}
func AxsStateFlexProfileObjAddExecution(builder *flatbuffers.Builder, execution flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(execution), 0)
}
func AxsStateFlexProfileObjEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
