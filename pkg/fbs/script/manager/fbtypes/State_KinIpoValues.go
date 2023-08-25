// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type State_KinIpoValuesT struct {
	Pos []float64 `json:"pos"`
	Vel float64 `json:"vel"`
	Acc float64 `json:"acc"`
	Jrk float64 `json:"jrk"`
}

func (t *State_KinIpoValuesT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	posOffset := flatbuffers.UOffsetT(0)
	if t.Pos != nil {
		posLength := len(t.Pos)
		State_KinIpoValuesStartPosVector(builder, posLength)
		for j := posLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.Pos[j])
		}
		posOffset = builder.EndVector(posLength)
	}
	State_KinIpoValuesStart(builder)
	State_KinIpoValuesAddPos(builder, posOffset)
	State_KinIpoValuesAddVel(builder, t.Vel)
	State_KinIpoValuesAddAcc(builder, t.Acc)
	State_KinIpoValuesAddJrk(builder, t.Jrk)
	return State_KinIpoValuesEnd(builder)
}

func (rcv *State_KinIpoValues) UnPackTo(t *State_KinIpoValuesT) {
	posLength := rcv.PosLength()
	t.Pos = make([]float64, posLength)
	for j := 0; j < posLength; j++ {
		t.Pos[j] = rcv.Pos(j)
	}
	t.Vel = rcv.Vel()
	t.Acc = rcv.Acc()
	t.Jrk = rcv.Jrk()
}

func (rcv *State_KinIpoValues) UnPack() *State_KinIpoValuesT {
	if rcv == nil { return nil }
	t := &State_KinIpoValuesT{}
	rcv.UnPackTo(t)
	return t
}

type State_KinIpoValues struct {
	_tab flatbuffers.Table
}

func GetRootAsState_KinIpoValues(buf []byte, offset flatbuffers.UOffsetT) *State_KinIpoValues {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &State_KinIpoValues{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsState_KinIpoValues(buf []byte, offset flatbuffers.UOffsetT) *State_KinIpoValues {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &State_KinIpoValues{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *State_KinIpoValues) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *State_KinIpoValues) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *State_KinIpoValues) Pos(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *State_KinIpoValues) PosLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *State_KinIpoValues) MutatePos(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func (rcv *State_KinIpoValues) Vel() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *State_KinIpoValues) MutateVel(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func (rcv *State_KinIpoValues) Acc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *State_KinIpoValues) MutateAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func (rcv *State_KinIpoValues) Jrk() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *State_KinIpoValues) MutateJrk(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

func State_KinIpoValuesStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func State_KinIpoValuesAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(pos), 0)
}
func State_KinIpoValuesStartPosVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func State_KinIpoValuesAddVel(builder *flatbuffers.Builder, vel float64) {
	builder.PrependFloat64Slot(1, vel, 0.0)
}
func State_KinIpoValuesAddAcc(builder *flatbuffers.Builder, acc float64) {
	builder.PrependFloat64Slot(2, acc, 0.0)
}
func State_KinIpoValuesAddJrk(builder *flatbuffers.Builder, jrk float64) {
	builder.PrependFloat64Slot(3, jrk, 0.0)
}
func State_KinIpoValuesEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
