// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type profileDeviceScalingTypeT struct {
	ScalingType deviceScalingType `json:"scalingType"`
}

func (t *profileDeviceScalingTypeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	profileDeviceScalingTypeStart(builder)
	profileDeviceScalingTypeAddScalingType(builder, t.ScalingType)
	return profileDeviceScalingTypeEnd(builder)
}

func (rcv *profileDeviceScalingType) UnPackTo(t *profileDeviceScalingTypeT) {
	t.ScalingType = rcv.ScalingType()
}

func (rcv *profileDeviceScalingType) UnPack() *profileDeviceScalingTypeT {
	if rcv == nil { return nil }
	t := &profileDeviceScalingTypeT{}
	rcv.UnPackTo(t)
	return t
}

type profileDeviceScalingType struct {
	_tab flatbuffers.Table
}

func GetRootAsprofileDeviceScalingType(buf []byte, offset flatbuffers.UOffsetT) *profileDeviceScalingType {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &profileDeviceScalingType{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsprofileDeviceScalingType(buf []byte, offset flatbuffers.UOffsetT) *profileDeviceScalingType {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &profileDeviceScalingType{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *profileDeviceScalingType) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *profileDeviceScalingType) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *profileDeviceScalingType) ScalingType() deviceScalingType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return deviceScalingType(rcv._tab.GetInt32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *profileDeviceScalingType) MutateScalingType(n deviceScalingType) bool {
	return rcv._tab.MutateInt32Slot(4, int32(n))
}

func profileDeviceScalingTypeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func profileDeviceScalingTypeAddScalingType(builder *flatbuffers.Builder, scalingType deviceScalingType) {
	builder.PrependInt32Slot(0, int32(scalingType), 0)
}
func profileDeviceScalingTypeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
