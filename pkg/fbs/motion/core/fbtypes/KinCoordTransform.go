// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// coordinate transformation based on the currently active transformations of the kinematics
type KinCoordTransformT struct {
	InPos []float64
	InCoordSys string
	OutPos []float64
	OutCoordSys string
}

func (t *KinCoordTransformT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	inPosOffset := flatbuffers.UOffsetT(0)
	if t.InPos != nil {
		inPosLength := len(t.InPos)
		KinCoordTransformStartInPosVector(builder, inPosLength)
		for j := inPosLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.InPos[j])
		}
		inPosOffset = builder.EndVector(inPosLength)
	}
	inCoordSysOffset := builder.CreateString(t.InCoordSys)
	outPosOffset := flatbuffers.UOffsetT(0)
	if t.OutPos != nil {
		outPosLength := len(t.OutPos)
		KinCoordTransformStartOutPosVector(builder, outPosLength)
		for j := outPosLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.OutPos[j])
		}
		outPosOffset = builder.EndVector(outPosLength)
	}
	outCoordSysOffset := builder.CreateString(t.OutCoordSys)
	KinCoordTransformStart(builder)
	KinCoordTransformAddInPos(builder, inPosOffset)
	KinCoordTransformAddInCoordSys(builder, inCoordSysOffset)
	KinCoordTransformAddOutPos(builder, outPosOffset)
	KinCoordTransformAddOutCoordSys(builder, outCoordSysOffset)
	return KinCoordTransformEnd(builder)
}

func (rcv *KinCoordTransform) UnPackTo(t *KinCoordTransformT) {
	inPosLength := rcv.InPosLength()
	t.InPos = make([]float64, inPosLength)
	for j := 0; j < inPosLength; j++ {
		t.InPos[j] = rcv.InPos(j)
	}
	t.InCoordSys = string(rcv.InCoordSys())
	outPosLength := rcv.OutPosLength()
	t.OutPos = make([]float64, outPosLength)
	for j := 0; j < outPosLength; j++ {
		t.OutPos[j] = rcv.OutPos(j)
	}
	t.OutCoordSys = string(rcv.OutCoordSys())
}

func (rcv *KinCoordTransform) UnPack() *KinCoordTransformT {
	if rcv == nil { return nil }
	t := &KinCoordTransformT{}
	rcv.UnPackTo(t)
	return t
}

type KinCoordTransform struct {
	_tab flatbuffers.Table
}

func GetRootAsKinCoordTransform(buf []byte, offset flatbuffers.UOffsetT) *KinCoordTransform {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinCoordTransform{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinCoordTransform(buf []byte, offset flatbuffers.UOffsetT) *KinCoordTransform {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinCoordTransform{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinCoordTransform) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinCoordTransform) Table() flatbuffers.Table {
	return rcv._tab
}

/// input coordinates
func (rcv *KinCoordTransform) InPos(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *KinCoordTransform) InPosLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// input coordinates
func (rcv *KinCoordTransform) MutateInPos(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

/// coordinate system of the input coordinates (default is "PCS")
func (rcv *KinCoordTransform) InCoordSys() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// coordinate system of the input coordinates (default is "PCS")
/// output coordinates, should be left out in the request
func (rcv *KinCoordTransform) OutPos(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *KinCoordTransform) OutPosLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// output coordinates, should be left out in the request
func (rcv *KinCoordTransform) MutateOutPos(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

/// coordinate system of the output coordinate (default is "ACS")
func (rcv *KinCoordTransform) OutCoordSys() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// coordinate system of the output coordinate (default is "ACS")
func KinCoordTransformStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func KinCoordTransformAddInPos(builder *flatbuffers.Builder, inPos flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(inPos), 0)
}
func KinCoordTransformStartInPosVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func KinCoordTransformAddInCoordSys(builder *flatbuffers.Builder, inCoordSys flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(inCoordSys), 0)
}
func KinCoordTransformAddOutPos(builder *flatbuffers.Builder, outPos flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(outPos), 0)
}
func KinCoordTransformStartOutPosVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func KinCoordTransformAddOutCoordSys(builder *flatbuffers.Builder, outCoordSys flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(outCoordSys), 0)
}
func KinCoordTransformEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
