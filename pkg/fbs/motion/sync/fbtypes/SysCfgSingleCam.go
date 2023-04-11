// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configuration of a single cam table
type SysCfgSingleCamT struct {
	Name string
	Points []float64
	Interpolation CamTableInterpolation
	CamBuilderData string
}

func (t *SysCfgSingleCamT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	pointsOffset := flatbuffers.UOffsetT(0)
	if t.Points != nil {
		pointsLength := len(t.Points)
		SysCfgSingleCamStartPointsVector(builder, pointsLength)
		for j := pointsLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.Points[j])
		}
		pointsOffset = builder.EndVector(pointsLength)
	}
	camBuilderDataOffset := builder.CreateString(t.CamBuilderData)
	SysCfgSingleCamStart(builder)
	SysCfgSingleCamAddName(builder, nameOffset)
	SysCfgSingleCamAddPoints(builder, pointsOffset)
	SysCfgSingleCamAddInterpolation(builder, t.Interpolation)
	SysCfgSingleCamAddCamBuilderData(builder, camBuilderDataOffset)
	return SysCfgSingleCamEnd(builder)
}

func (rcv *SysCfgSingleCam) UnPackTo(t *SysCfgSingleCamT) {
	t.Name = string(rcv.Name())
	pointsLength := rcv.PointsLength()
	t.Points = make([]float64, pointsLength)
	for j := 0; j < pointsLength; j++ {
		t.Points[j] = rcv.Points(j)
	}
	t.Interpolation = rcv.Interpolation()
	t.CamBuilderData = string(rcv.CamBuilderData())
}

func (rcv *SysCfgSingleCam) UnPack() *SysCfgSingleCamT {
	if rcv == nil { return nil }
	t := &SysCfgSingleCamT{}
	rcv.UnPackTo(t)
	return t
}

type SysCfgSingleCam struct {
	_tab flatbuffers.Table
}

func GetRootAsSysCfgSingleCam(buf []byte, offset flatbuffers.UOffsetT) *SysCfgSingleCam {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SysCfgSingleCam{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSysCfgSingleCam(buf []byte, offset flatbuffers.UOffsetT) *SysCfgSingleCam {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SysCfgSingleCam{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SysCfgSingleCam) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SysCfgSingleCam) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the cam table (filled when reading; optional for writing [content is always ignored])
func (rcv *SysCfgSingleCam) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the cam table (filled when reading; optional for writing [content is always ignored])
/// interpolation points of the cam table
func (rcv *SysCfgSingleCam) Points(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *SysCfgSingleCam) PointsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// interpolation points of the cam table
func (rcv *SysCfgSingleCam) MutatePoints(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

/// interpolation type for this cam table
func (rcv *SysCfgSingleCam) Interpolation() CamTableInterpolation {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return CamTableInterpolation(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// interpolation type for this cam table
func (rcv *SysCfgSingleCam) MutateInterpolation(n CamTableInterpolation) bool {
	return rcv._tab.MutateInt8Slot(8, int8(n))
}

/// data of the CamBuilder
func (rcv *SysCfgSingleCam) CamBuilderData() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// data of the CamBuilder
func SysCfgSingleCamStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func SysCfgSingleCamAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func SysCfgSingleCamAddPoints(builder *flatbuffers.Builder, points flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(points), 0)
}
func SysCfgSingleCamStartPointsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func SysCfgSingleCamAddInterpolation(builder *flatbuffers.Builder, interpolation CamTableInterpolation) {
	builder.PrependInt8Slot(2, int8(interpolation), 0)
}
func SysCfgSingleCamAddCamBuilderData(builder *flatbuffers.Builder, camBuilderData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(camBuilderData), 0)
}
func SysCfgSingleCamEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
