// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// axis braking distance
type AxsBrakingDistanceExT struct {
	SelectedTypes []BrakingDistanceType
	DistanceUnit string
	Distance float64
	DistanceType BrakingDistanceType
}

func (t *AxsBrakingDistanceExT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	selectedTypesOffset := flatbuffers.UOffsetT(0)
	if t.SelectedTypes != nil {
		selectedTypesLength := len(t.SelectedTypes)
		AxsBrakingDistanceExStartSelectedTypesVector(builder, selectedTypesLength)
		for j := selectedTypesLength - 1; j >= 0; j-- {
			builder.PrependInt8(int8(t.SelectedTypes[j]))
		}
		selectedTypesOffset = builder.EndVector(selectedTypesLength)
	}
	distanceUnitOffset := builder.CreateString(t.DistanceUnit)
	AxsBrakingDistanceExStart(builder)
	AxsBrakingDistanceExAddSelectedTypes(builder, selectedTypesOffset)
	AxsBrakingDistanceExAddDistanceUnit(builder, distanceUnitOffset)
	AxsBrakingDistanceExAddDistance(builder, t.Distance)
	AxsBrakingDistanceExAddDistanceType(builder, t.DistanceType)
	return AxsBrakingDistanceExEnd(builder)
}

func (rcv *AxsBrakingDistanceEx) UnPackTo(t *AxsBrakingDistanceExT) {
	selectedTypesLength := rcv.SelectedTypesLength()
	t.SelectedTypes = make([]BrakingDistanceType, selectedTypesLength)
	for j := 0; j < selectedTypesLength; j++ {
		t.SelectedTypes[j] = rcv.SelectedTypes(j)
	}
	t.DistanceUnit = string(rcv.DistanceUnit())
	t.Distance = rcv.Distance()
	t.DistanceType = rcv.DistanceType()
}

func (rcv *AxsBrakingDistanceEx) UnPack() *AxsBrakingDistanceExT {
	if rcv == nil { return nil }
	t := &AxsBrakingDistanceExT{}
	rcv.UnPackTo(t)
	return t
}

type AxsBrakingDistanceEx struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsBrakingDistanceEx(buf []byte, offset flatbuffers.UOffsetT) *AxsBrakingDistanceEx {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsBrakingDistanceEx{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsBrakingDistanceEx(buf []byte, offset flatbuffers.UOffsetT) *AxsBrakingDistanceEx {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsBrakingDistanceEx{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsBrakingDistanceEx) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsBrakingDistanceEx) Table() flatbuffers.Table {
	return rcv._tab
}

/// selected braking distance types in the calculation
/// default value is [SOFT_STOP, ESTOP]
func (rcv *AxsBrakingDistanceEx) SelectedTypes(j int) BrakingDistanceType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return BrakingDistanceType(rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1)))
	}
	return 0
}

func (rcv *AxsBrakingDistanceEx) SelectedTypesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// selected braking distance types in the calculation
/// default value is [SOFT_STOP, ESTOP]
func (rcv *AxsBrakingDistanceEx) MutateSelectedTypes(j int, n BrakingDistanceType) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), int8(n))
	}
	return false
}

/// unit in which braking distance value should be calculated
/// default value is the configured unit for the axis
func (rcv *AxsBrakingDistanceEx) DistanceUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit in which braking distance value should be calculated
/// default value is the configured unit for the axis
/// calculated braking distance
func (rcv *AxsBrakingDistanceEx) Distance() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// calculated braking distance
func (rcv *AxsBrakingDistanceEx) MutateDistance(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

/// the distance type of the calculated braking distance
func (rcv *AxsBrakingDistanceEx) DistanceType() BrakingDistanceType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return BrakingDistanceType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// the distance type of the calculated braking distance
func (rcv *AxsBrakingDistanceEx) MutateDistanceType(n BrakingDistanceType) bool {
	return rcv._tab.MutateInt8Slot(10, int8(n))
}

func AxsBrakingDistanceExStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func AxsBrakingDistanceExAddSelectedTypes(builder *flatbuffers.Builder, selectedTypes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(selectedTypes), 0)
}
func AxsBrakingDistanceExStartSelectedTypesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func AxsBrakingDistanceExAddDistanceUnit(builder *flatbuffers.Builder, distanceUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(distanceUnit), 0)
}
func AxsBrakingDistanceExAddDistance(builder *flatbuffers.Builder, distance float64) {
	builder.PrependFloat64Slot(2, distance, 0.0)
}
func AxsBrakingDistanceExAddDistanceType(builder *flatbuffers.Builder, distanceType BrakingDistanceType) {
	builder.PrependInt8Slot(3, int8(distanceType), 0)
}
func AxsBrakingDistanceExEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
