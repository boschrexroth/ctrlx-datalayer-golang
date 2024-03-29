// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type State_KinCoordTransformT struct {
	ObjName string `json:"objName"`
	PosIn []float64 `json:"posIn"`
	CoordSysIn string `json:"coordSysIn"`
	CoordSysOut string `json:"coordSysOut"`
}

func (t *State_KinCoordTransformT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	objNameOffset := flatbuffers.UOffsetT(0)
	if t.ObjName != "" {
		objNameOffset = builder.CreateString(t.ObjName)
	}
	posInOffset := flatbuffers.UOffsetT(0)
	if t.PosIn != nil {
		posInLength := len(t.PosIn)
		State_KinCoordTransformStartPosInVector(builder, posInLength)
		for j := posInLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.PosIn[j])
		}
		posInOffset = builder.EndVector(posInLength)
	}
	coordSysInOffset := flatbuffers.UOffsetT(0)
	if t.CoordSysIn != "" {
		coordSysInOffset = builder.CreateString(t.CoordSysIn)
	}
	coordSysOutOffset := flatbuffers.UOffsetT(0)
	if t.CoordSysOut != "" {
		coordSysOutOffset = builder.CreateString(t.CoordSysOut)
	}
	State_KinCoordTransformStart(builder)
	State_KinCoordTransformAddObjName(builder, objNameOffset)
	State_KinCoordTransformAddPosIn(builder, posInOffset)
	State_KinCoordTransformAddCoordSysIn(builder, coordSysInOffset)
	State_KinCoordTransformAddCoordSysOut(builder, coordSysOutOffset)
	return State_KinCoordTransformEnd(builder)
}

func (rcv *State_KinCoordTransform) UnPackTo(t *State_KinCoordTransformT) {
	t.ObjName = string(rcv.ObjName())
	posInLength := rcv.PosInLength()
	t.PosIn = make([]float64, posInLength)
	for j := 0; j < posInLength; j++ {
		t.PosIn[j] = rcv.PosIn(j)
	}
	t.CoordSysIn = string(rcv.CoordSysIn())
	t.CoordSysOut = string(rcv.CoordSysOut())
}

func (rcv *State_KinCoordTransform) UnPack() *State_KinCoordTransformT {
	if rcv == nil { return nil }
	t := &State_KinCoordTransformT{}
	rcv.UnPackTo(t)
	return t
}

type State_KinCoordTransform struct {
	_tab flatbuffers.Table
}

func GetRootAsState_KinCoordTransform(buf []byte, offset flatbuffers.UOffsetT) *State_KinCoordTransform {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &State_KinCoordTransform{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsState_KinCoordTransform(buf []byte, offset flatbuffers.UOffsetT) *State_KinCoordTransform {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &State_KinCoordTransform{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *State_KinCoordTransform) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *State_KinCoordTransform) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *State_KinCoordTransform) ObjName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *State_KinCoordTransform) PosIn(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *State_KinCoordTransform) PosInLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *State_KinCoordTransform) MutatePosIn(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func (rcv *State_KinCoordTransform) CoordSysIn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *State_KinCoordTransform) CoordSysOut() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func State_KinCoordTransformStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func State_KinCoordTransformAddObjName(builder *flatbuffers.Builder, objName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(objName), 0)
}
func State_KinCoordTransformAddPosIn(builder *flatbuffers.Builder, posIn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(posIn), 0)
}
func State_KinCoordTransformStartPosInVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func State_KinCoordTransformAddCoordSysIn(builder *flatbuffers.Builder, coordSysIn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(coordSysIn), 0)
}
func State_KinCoordTransformAddCoordSysOut(builder *flatbuffers.Builder, coordSysOut flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(coordSysOut), 0)
}
func State_KinCoordTransformEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
