// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// currently interpolated values
type AxsIpoValues struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsIpoValues(buf []byte, offset flatbuffers.UOffsetT) *AxsIpoValues {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsIpoValues{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsIpoValues(buf []byte, offset flatbuffers.UOffsetT) *AxsIpoValues {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsIpoValues{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsIpoValues) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsIpoValues) Table() flatbuffers.Table {
	return rcv._tab
}

/// currently interpolated position
func (rcv *AxsIpoValues) IpoPos() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// currently interpolated position
func (rcv *AxsIpoValues) MutateIpoPos(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// currently interpolated velocity
func (rcv *AxsIpoValues) IpoVel() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// currently interpolated velocity
func (rcv *AxsIpoValues) MutateIpoVel(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// currently interpolated acceleration
func (rcv *AxsIpoValues) IpoAcc() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// currently interpolated acceleration
func (rcv *AxsIpoValues) MutateIpoAcc(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

/// currently interpolated jerk
func (rcv *AxsIpoValues) IpoJrk() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// currently interpolated jerk
func (rcv *AxsIpoValues) MutateIpoJrk(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

/// unit of currently interpolated position
func (rcv *AxsIpoValues) IpoPosUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of currently interpolated position
/// unit of currently interpolated velocity
func (rcv *AxsIpoValues) IpoVelUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of currently interpolated velocity
/// unit of currently interpolated acceleration
func (rcv *AxsIpoValues) IpoAccUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of currently interpolated acceleration
/// unit of currently interpolated jerk
func (rcv *AxsIpoValues) IpoJrkUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of currently interpolated jerk
func AxsIpoValuesStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func AxsIpoValuesAddIpoPos(builder *flatbuffers.Builder, ipoPos float64) {
	builder.PrependFloat64Slot(0, ipoPos, 0.0)
}
func AxsIpoValuesAddIpoVel(builder *flatbuffers.Builder, ipoVel float64) {
	builder.PrependFloat64Slot(1, ipoVel, 0.0)
}
func AxsIpoValuesAddIpoAcc(builder *flatbuffers.Builder, ipoAcc float64) {
	builder.PrependFloat64Slot(2, ipoAcc, 0.0)
}
func AxsIpoValuesAddIpoJrk(builder *flatbuffers.Builder, ipoJrk float64) {
	builder.PrependFloat64Slot(3, ipoJrk, 0.0)
}
func AxsIpoValuesAddIpoPosUnit(builder *flatbuffers.Builder, ipoPosUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(ipoPosUnit), 0)
}
func AxsIpoValuesAddIpoVelUnit(builder *flatbuffers.Builder, ipoVelUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(ipoVelUnit), 0)
}
func AxsIpoValuesAddIpoAccUnit(builder *flatbuffers.Builder, ipoAccUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(ipoAccUnit), 0)
}
func AxsIpoValuesAddIpoJrkUnit(builder *flatbuffers.Builder, ipoJrkUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(ipoJrkUnit), 0)
}
func AxsIpoValuesEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
