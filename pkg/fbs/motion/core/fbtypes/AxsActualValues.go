// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// actual values from the drives
type AxsActualValues struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsActualValues(buf []byte, offset flatbuffers.UOffsetT) *AxsActualValues {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsActualValues{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsActualValues(buf []byte, offset flatbuffers.UOffsetT) *AxsActualValues {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsActualValues{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsActualValues) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsActualValues) Table() flatbuffers.Table {
	return rcv._tab
}

/// actual drive position
func (rcv *AxsActualValues) ActualPos() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// actual drive position
func (rcv *AxsActualValues) MutateActualPos(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// actual drive velocity (currently not supported for real drives)
func (rcv *AxsActualValues) ActualVel() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// actual drive velocity (currently not supported for real drives)
func (rcv *AxsActualValues) MutateActualVel(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// actual drive acceleration (currently not supported for real drives)
func (rcv *AxsActualValues) ActualAcc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// actual drive acceleration (currently not supported for real drives)
func (rcv *AxsActualValues) MutateActualAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

/// actual drive torque (currently not supported for real drives)
func (rcv *AxsActualValues) ActualTorque() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// actual drive torque (currently not supported for real drives)
func (rcv *AxsActualValues) MutateActualTorque(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

/// actual distance left to the commanded target (currently not supported for real drives)
func (rcv *AxsActualValues) DistLeft() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// actual distance left to the commanded target (currently not supported for real drives)
func (rcv *AxsActualValues) MutateDistLeft(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

/// unit of the actual drive position
func (rcv *AxsActualValues) ActualPosUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the actual drive position
/// unit of the actual drive velocity (currently not supported for real drives)
func (rcv *AxsActualValues) ActualVelUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the actual drive velocity (currently not supported for real drives)
/// unit of the actual drive acceleration (currently not supported for real drives)
func (rcv *AxsActualValues) ActualAccUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the actual drive acceleration (currently not supported for real drives)
/// unit of the actual drive torque (currently not supported for real drives)
func (rcv *AxsActualValues) ActualTorqueUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the actual drive torque (currently not supported for real drives)
/// unit of the actual distance left to the commanded target (currently not supported for real drives)
func (rcv *AxsActualValues) DistLeftUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the actual distance left to the commanded target (currently not supported for real drives)
func AxsActualValuesStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func AxsActualValuesAddActualPos(builder *flatbuffers.Builder, actualPos float64) {
	builder.PrependFloat64Slot(0, actualPos, 0.0)
}
func AxsActualValuesAddActualVel(builder *flatbuffers.Builder, actualVel float64) {
	builder.PrependFloat64Slot(1, actualVel, 0.0)
}
func AxsActualValuesAddActualAcc(builder *flatbuffers.Builder, actualAcc float64) {
	builder.PrependFloat64Slot(2, actualAcc, 0.0)
}
func AxsActualValuesAddActualTorque(builder *flatbuffers.Builder, actualTorque float64) {
	builder.PrependFloat64Slot(3, actualTorque, 0.0)
}
func AxsActualValuesAddDistLeft(builder *flatbuffers.Builder, distLeft float64) {
	builder.PrependFloat64Slot(4, distLeft, 0.0)
}
func AxsActualValuesAddActualPosUnit(builder *flatbuffers.Builder, actualPosUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(actualPosUnit), 0)
}
func AxsActualValuesAddActualVelUnit(builder *flatbuffers.Builder, actualVelUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(actualVelUnit), 0)
}
func AxsActualValuesAddActualAccUnit(builder *flatbuffers.Builder, actualAccUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(actualAccUnit), 0)
}
func AxsActualValuesAddActualTorqueUnit(builder *flatbuffers.Builder, actualTorqueUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(actualTorqueUnit), 0)
}
func AxsActualValuesAddDistLeftUnit(builder *flatbuffers.Builder, distLeftUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(distLeftUnit), 0)
}
func AxsActualValuesEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
