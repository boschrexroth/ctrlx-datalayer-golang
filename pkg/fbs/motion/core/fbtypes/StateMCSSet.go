// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// informations on a single MCS set
type StateMCSSetT struct {
	Valid bool `json:"valid"`
	LastCheckDiag []*JointTrafoCheckDiagT `json:"lastCheckDiag"`
}

func (t *StateMCSSetT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	lastCheckDiagOffset := flatbuffers.UOffsetT(0)
	if t.LastCheckDiag != nil {
		lastCheckDiagLength := len(t.LastCheckDiag)
		lastCheckDiagOffsets := make([]flatbuffers.UOffsetT, lastCheckDiagLength)
		for j := 0; j < lastCheckDiagLength; j++ {
			lastCheckDiagOffsets[j] = t.LastCheckDiag[j].Pack(builder)
		}
		StateMCSSetStartLastCheckDiagVector(builder, lastCheckDiagLength)
		for j := lastCheckDiagLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(lastCheckDiagOffsets[j])
		}
		lastCheckDiagOffset = builder.EndVector(lastCheckDiagLength)
	}
	StateMCSSetStart(builder)
	StateMCSSetAddValid(builder, t.Valid)
	StateMCSSetAddLastCheckDiag(builder, lastCheckDiagOffset)
	return StateMCSSetEnd(builder)
}

func (rcv *StateMCSSet) UnPackTo(t *StateMCSSetT) {
	t.Valid = rcv.Valid()
	lastCheckDiagLength := rcv.LastCheckDiagLength()
	t.LastCheckDiag = make([]*JointTrafoCheckDiagT, lastCheckDiagLength)
	for j := 0; j < lastCheckDiagLength; j++ {
		x := JointTrafoCheckDiag{}
		rcv.LastCheckDiag(&x, j)
		t.LastCheckDiag[j] = x.UnPack()
	}
}

func (rcv *StateMCSSet) UnPack() *StateMCSSetT {
	if rcv == nil { return nil }
	t := &StateMCSSetT{}
	rcv.UnPackTo(t)
	return t
}

type StateMCSSet struct {
	_tab flatbuffers.Table
}

func GetRootAsStateMCSSet(buf []byte, offset flatbuffers.UOffsetT) *StateMCSSet {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StateMCSSet{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStateMCSSet(buf []byte, offset flatbuffers.UOffsetT) *StateMCSSet {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StateMCSSet{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StateMCSSet) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StateMCSSet) Table() flatbuffers.Table {
	return rcv._tab
}

/// is the MCS set valid and can be used?
func (rcv *StateMCSSet) Valid() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// is the MCS set valid and can be used?
func (rcv *StateMCSSet) MutateValid(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

/// all diagnoses, that were created in the last validation of the MCS set
func (rcv *StateMCSSet) LastCheckDiag(obj *JointTrafoCheckDiag, j int) bool {
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

func (rcv *StateMCSSet) LastCheckDiagLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// all diagnoses, that were created in the last validation of the MCS set
func StateMCSSetStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func StateMCSSetAddValid(builder *flatbuffers.Builder, valid bool) {
	builder.PrependBoolSlot(0, valid, false)
}
func StateMCSSetAddLastCheckDiag(builder *flatbuffers.Builder, lastCheckDiag flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(lastCheckDiag), 0)
}
func StateMCSSetStartLastCheckDiagVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func StateMCSSetEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}