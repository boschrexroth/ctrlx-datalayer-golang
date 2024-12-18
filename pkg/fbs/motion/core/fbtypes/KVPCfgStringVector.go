// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Table for a vector of string values
type KVPCfgStringVectorT struct {
	StringVector []string `json:"stringVector"`
}

func (t *KVPCfgStringVectorT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	stringVectorOffset := flatbuffers.UOffsetT(0)
	if t.StringVector != nil {
		stringVectorLength := len(t.StringVector)
		stringVectorOffsets := make([]flatbuffers.UOffsetT, stringVectorLength)
		for j := 0; j < stringVectorLength; j++ {
			stringVectorOffsets[j] = builder.CreateString(t.StringVector[j])
		}
		KVPCfgStringVectorStartStringVectorVector(builder, stringVectorLength)
		for j := stringVectorLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(stringVectorOffsets[j])
		}
		stringVectorOffset = builder.EndVector(stringVectorLength)
	}
	KVPCfgStringVectorStart(builder)
	KVPCfgStringVectorAddStringVector(builder, stringVectorOffset)
	return KVPCfgStringVectorEnd(builder)
}

func (rcv *KVPCfgStringVector) UnPackTo(t *KVPCfgStringVectorT) {
	stringVectorLength := rcv.StringVectorLength()
	t.StringVector = make([]string, stringVectorLength)
	for j := 0; j < stringVectorLength; j++ {
		t.StringVector[j] = string(rcv.StringVector(j))
	}
}

func (rcv *KVPCfgStringVector) UnPack() *KVPCfgStringVectorT {
	if rcv == nil { return nil }
	t := &KVPCfgStringVectorT{}
	rcv.UnPackTo(t)
	return t
}

type KVPCfgStringVector struct {
	_tab flatbuffers.Table
}

func GetRootAsKVPCfgStringVector(buf []byte, offset flatbuffers.UOffsetT) *KVPCfgStringVector {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KVPCfgStringVector{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKVPCfgStringVector(buf []byte, offset flatbuffers.UOffsetT) *KVPCfgStringVector {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KVPCfgStringVector{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KVPCfgStringVector) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KVPCfgStringVector) Table() flatbuffers.Table {
	return rcv._tab
}

/// vector of strings
func (rcv *KVPCfgStringVector) StringVector(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *KVPCfgStringVector) StringVectorLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// vector of strings
func KVPCfgStringVectorStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func KVPCfgStringVectorAddStringVector(builder *flatbuffers.Builder, stringVector flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(stringVector), 0)
}
func KVPCfgStringVectorStartStringVectorVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KVPCfgStringVectorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
