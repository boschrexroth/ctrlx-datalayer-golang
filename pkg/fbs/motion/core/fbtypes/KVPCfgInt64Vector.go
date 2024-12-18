// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Table for a vector of int64 values
type KVPCfgInt64VectorT struct {
	IntVector []int64 `json:"intVector"`
}

func (t *KVPCfgInt64VectorT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	intVectorOffset := flatbuffers.UOffsetT(0)
	if t.IntVector != nil {
		intVectorLength := len(t.IntVector)
		KVPCfgInt64VectorStartIntVectorVector(builder, intVectorLength)
		for j := intVectorLength - 1; j >= 0; j-- {
			builder.PrependInt64(t.IntVector[j])
		}
		intVectorOffset = builder.EndVector(intVectorLength)
	}
	KVPCfgInt64VectorStart(builder)
	KVPCfgInt64VectorAddIntVector(builder, intVectorOffset)
	return KVPCfgInt64VectorEnd(builder)
}

func (rcv *KVPCfgInt64Vector) UnPackTo(t *KVPCfgInt64VectorT) {
	intVectorLength := rcv.IntVectorLength()
	t.IntVector = make([]int64, intVectorLength)
	for j := 0; j < intVectorLength; j++ {
		t.IntVector[j] = rcv.IntVector(j)
	}
}

func (rcv *KVPCfgInt64Vector) UnPack() *KVPCfgInt64VectorT {
	if rcv == nil { return nil }
	t := &KVPCfgInt64VectorT{}
	rcv.UnPackTo(t)
	return t
}

type KVPCfgInt64Vector struct {
	_tab flatbuffers.Table
}

func GetRootAsKVPCfgInt64Vector(buf []byte, offset flatbuffers.UOffsetT) *KVPCfgInt64Vector {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KVPCfgInt64Vector{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKVPCfgInt64Vector(buf []byte, offset flatbuffers.UOffsetT) *KVPCfgInt64Vector {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KVPCfgInt64Vector{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KVPCfgInt64Vector) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KVPCfgInt64Vector) Table() flatbuffers.Table {
	return rcv._tab
}

/// vector of int64 numbers
func (rcv *KVPCfgInt64Vector) IntVector(j int) int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *KVPCfgInt64Vector) IntVectorLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of int64 numbers
func (rcv *KVPCfgInt64Vector) MutateIntVector(j int, n int64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func KVPCfgInt64VectorStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func KVPCfgInt64VectorAddIntVector(builder *flatbuffers.Builder, intVector flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(intVector), 0)
}
func KVPCfgInt64VectorStartIntVectorVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func KVPCfgInt64VectorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
