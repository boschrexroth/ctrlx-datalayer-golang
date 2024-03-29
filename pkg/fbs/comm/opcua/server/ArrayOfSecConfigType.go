// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package server

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfSecConfigTypeT struct {
	SecConfigArray []*SecConfigTypeT `json:"secConfigArray"`
}

func (t *ArrayOfSecConfigTypeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	secConfigArrayOffset := flatbuffers.UOffsetT(0)
	if t.SecConfigArray != nil {
		secConfigArrayLength := len(t.SecConfigArray)
		secConfigArrayOffsets := make([]flatbuffers.UOffsetT, secConfigArrayLength)
		for j := 0; j < secConfigArrayLength; j++ {
			secConfigArrayOffsets[j] = t.SecConfigArray[j].Pack(builder)
		}
		ArrayOfSecConfigTypeStartSecConfigArrayVector(builder, secConfigArrayLength)
		for j := secConfigArrayLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(secConfigArrayOffsets[j])
		}
		secConfigArrayOffset = builder.EndVector(secConfigArrayLength)
	}
	ArrayOfSecConfigTypeStart(builder)
	ArrayOfSecConfigTypeAddSecConfigArray(builder, secConfigArrayOffset)
	return ArrayOfSecConfigTypeEnd(builder)
}

func (rcv *ArrayOfSecConfigType) UnPackTo(t *ArrayOfSecConfigTypeT) {
	secConfigArrayLength := rcv.SecConfigArrayLength()
	t.SecConfigArray = make([]*SecConfigTypeT, secConfigArrayLength)
	for j := 0; j < secConfigArrayLength; j++ {
		x := SecConfigType{}
		rcv.SecConfigArray(&x, j)
		t.SecConfigArray[j] = x.UnPack()
	}
}

func (rcv *ArrayOfSecConfigType) UnPack() *ArrayOfSecConfigTypeT {
	if rcv == nil { return nil }
	t := &ArrayOfSecConfigTypeT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayOfSecConfigType struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfSecConfigType(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfSecConfigType {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfSecConfigType{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfSecConfigType(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfSecConfigType {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfSecConfigType{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfSecConfigType) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfSecConfigType) Table() flatbuffers.Table {
	return rcv._tab
}

/// The endpoint security configurations of the OPC UA Server in an array
/// Each array element describes one security configuration of OPC UA Server endpoint
func (rcv *ArrayOfSecConfigType) SecConfigArray(obj *SecConfigType, j int) bool {
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

func (rcv *ArrayOfSecConfigType) SecConfigArrayLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// The endpoint security configurations of the OPC UA Server in an array
/// Each array element describes one security configuration of OPC UA Server endpoint
func ArrayOfSecConfigTypeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfSecConfigTypeAddSecConfigArray(builder *flatbuffers.Builder, secConfigArray flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(secConfigArray), 0)
}
func ArrayOfSecConfigTypeStartSecConfigArrayVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ArrayOfSecConfigTypeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
