// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// data of a all axis zero point offsets for a joint transformation
type KinCfgJntTrafoAxisOffsetsT struct {
	Offsets []*KinCfgJntTrafoSingleAxisOffsetT `json:"offsets"`
}

func (t *KinCfgJntTrafoAxisOffsetsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	offsetsOffset := flatbuffers.UOffsetT(0)
	if t.Offsets != nil {
		offsetsLength := len(t.Offsets)
		offsetsOffsets := make([]flatbuffers.UOffsetT, offsetsLength)
		for j := 0; j < offsetsLength; j++ {
			offsetsOffsets[j] = t.Offsets[j].Pack(builder)
		}
		KinCfgJntTrafoAxisOffsetsStartOffsetsVector(builder, offsetsLength)
		for j := offsetsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(offsetsOffsets[j])
		}
		offsetsOffset = builder.EndVector(offsetsLength)
	}
	KinCfgJntTrafoAxisOffsetsStart(builder)
	KinCfgJntTrafoAxisOffsetsAddOffsets(builder, offsetsOffset)
	return KinCfgJntTrafoAxisOffsetsEnd(builder)
}

func (rcv *KinCfgJntTrafoAxisOffsets) UnPackTo(t *KinCfgJntTrafoAxisOffsetsT) {
	offsetsLength := rcv.OffsetsLength()
	t.Offsets = make([]*KinCfgJntTrafoSingleAxisOffsetT, offsetsLength)
	for j := 0; j < offsetsLength; j++ {
		x := KinCfgJntTrafoSingleAxisOffset{}
		rcv.Offsets(&x, j)
		t.Offsets[j] = x.UnPack()
	}
}

func (rcv *KinCfgJntTrafoAxisOffsets) UnPack() *KinCfgJntTrafoAxisOffsetsT {
	if rcv == nil { return nil }
	t := &KinCfgJntTrafoAxisOffsetsT{}
	rcv.UnPackTo(t)
	return t
}

type KinCfgJntTrafoAxisOffsets struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCfgJntTrafoAxisOffsets(buf []byte, offset flatbuffers.UOffsetT) *KinCfgJntTrafoAxisOffsets {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCfgJntTrafoAxisOffsets{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCfgJntTrafoAxisOffsets(buf []byte, offset flatbuffers.UOffsetT) *KinCfgJntTrafoAxisOffsets {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCfgJntTrafoAxisOffsets{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCfgJntTrafoAxisOffsets) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCfgJntTrafoAxisOffsets) Table() flatbuffers.Table {
	return rcv._tab
}

/// assignment as pairs of <axis object name; zero point offset>
func (rcv *KinCfgJntTrafoAxisOffsets) Offsets(obj *KinCfgJntTrafoSingleAxisOffset, j int) bool {
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

func (rcv *KinCfgJntTrafoAxisOffsets) OffsetsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// assignment as pairs of <axis object name; zero point offset>
func KinCfgJntTrafoAxisOffsetsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func KinCfgJntTrafoAxisOffsetsAddOffsets(builder *flatbuffers.Builder, offsets flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(offsets), 0)
}
func KinCfgJntTrafoAxisOffsetsStartOffsetsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func KinCfgJntTrafoAxisOffsetsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
