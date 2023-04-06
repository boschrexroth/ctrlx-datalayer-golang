// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Synchronization points to specify the call sequence, alternatively to the use of run index
/// Notes:
///   - The names of the synchronization points of a callables have to be unique
///   - The names of the synchronization points of a callables have match to the Data Layer compliance guide lines meaning any alphanumeric character are allowed [a-zA-Z_][a-zA-Z0-9-._]+
///   - To ensure that other callable can run after or before this callable it's recommended to set at least one synchronization point in each list
type SyncPointsT struct {
	After []string
	Before []string
}

func (t *SyncPointsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	afterOffset := flatbuffers.UOffsetT(0)
	if t.After != nil {
		afterLength := len(t.After)
		afterOffsets := make([]flatbuffers.UOffsetT, afterLength)
		for j := 0; j < afterLength; j++ {
			afterOffsets[j] = builder.CreateString(t.After[j])
		}
		SyncPointsStartAfterVector(builder, afterLength)
		for j := afterLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(afterOffsets[j])
		}
		afterOffset = builder.EndVector(afterLength)
	}
	beforeOffset := flatbuffers.UOffsetT(0)
	if t.Before != nil {
		beforeLength := len(t.Before)
		beforeOffsets := make([]flatbuffers.UOffsetT, beforeLength)
		for j := 0; j < beforeLength; j++ {
			beforeOffsets[j] = builder.CreateString(t.Before[j])
		}
		SyncPointsStartBeforeVector(builder, beforeLength)
		for j := beforeLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(beforeOffsets[j])
		}
		beforeOffset = builder.EndVector(beforeLength)
	}
	SyncPointsStart(builder)
	SyncPointsAddAfter(builder, afterOffset)
	SyncPointsAddBefore(builder, beforeOffset)
	return SyncPointsEnd(builder)
}

func (rcv *SyncPoints) UnPackTo(t *SyncPointsT) {
	afterLength := rcv.AfterLength()
	t.After = make([]string, afterLength)
	for j := 0; j < afterLength; j++ {
		t.After[j] = string(rcv.After(j))
	}
	beforeLength := rcv.BeforeLength()
	t.Before = make([]string, beforeLength)
	for j := 0; j < beforeLength; j++ {
		t.Before[j] = string(rcv.Before(j))
	}
}

func (rcv *SyncPoints) UnPack() *SyncPointsT {
	if rcv == nil { return nil }
	t := &SyncPointsT{}
	rcv.UnPackTo(t)
	return t
}

type SyncPoints struct {
	_tab flatbuffers.Table
}

func GetRootAsSyncPoints(buf []byte, offset flatbuffers.UOffsetT) *SyncPoints {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SyncPoints{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSyncPoints(buf []byte, offset flatbuffers.UOffsetT) *SyncPoints {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SyncPoints{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SyncPoints) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SyncPoints) Table() flatbuffers.Table {
	return rcv._tab
}

/// User defined synchronization points, execute callable in order after these points, 
func (rcv *SyncPoints) After(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *SyncPoints) AfterLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// User defined synchronization points, execute callable in order after these points, 
/// User defined synchronization points, execute callable in order before these points, any alphanumeric character
func (rcv *SyncPoints) Before(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *SyncPoints) BeforeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// User defined synchronization points, execute callable in order before these points, any alphanumeric character
func SyncPointsStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SyncPointsAddAfter(builder *flatbuffers.Builder, after flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(after), 0)
}
func SyncPointsStartAfterVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SyncPointsAddBefore(builder *flatbuffers.Builder, before flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(before), 0)
}
func SyncPointsStartBeforeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SyncPointsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
