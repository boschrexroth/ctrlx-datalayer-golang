// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package axis

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AxisDetailsT struct {
	SubDeviceName string `json:"subDeviceName"`
	SubDeviceType string `json:"subDeviceType"`
	CommAddress string `json:"commAddress"`
	PhysicalAddress string `json:"physicalAddress"`
}

func (t *AxisDetailsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	subDeviceNameOffset := flatbuffers.UOffsetT(0)
	if t.SubDeviceName != "" {
		subDeviceNameOffset = builder.CreateString(t.SubDeviceName)
	}
	subDeviceTypeOffset := flatbuffers.UOffsetT(0)
	if t.SubDeviceType != "" {
		subDeviceTypeOffset = builder.CreateString(t.SubDeviceType)
	}
	commAddressOffset := flatbuffers.UOffsetT(0)
	if t.CommAddress != "" {
		commAddressOffset = builder.CreateString(t.CommAddress)
	}
	physicalAddressOffset := flatbuffers.UOffsetT(0)
	if t.PhysicalAddress != "" {
		physicalAddressOffset = builder.CreateString(t.PhysicalAddress)
	}
	AxisDetailsStart(builder)
	AxisDetailsAddSubDeviceName(builder, subDeviceNameOffset)
	AxisDetailsAddSubDeviceType(builder, subDeviceTypeOffset)
	AxisDetailsAddCommAddress(builder, commAddressOffset)
	AxisDetailsAddPhysicalAddress(builder, physicalAddressOffset)
	return AxisDetailsEnd(builder)
}

func (rcv *AxisDetails) UnPackTo(t *AxisDetailsT) {
	t.SubDeviceName = string(rcv.SubDeviceName())
	t.SubDeviceType = string(rcv.SubDeviceType())
	t.CommAddress = string(rcv.CommAddress())
	t.PhysicalAddress = string(rcv.PhysicalAddress())
}

func (rcv *AxisDetails) UnPack() *AxisDetailsT {
	if rcv == nil { return nil }
	t := &AxisDetailsT{}
	rcv.UnPackTo(t)
	return t
}

type AxisDetails struct {
	_tab flatbuffers.Table
}

func GetRootAsAxisDetails(buf []byte, offset flatbuffers.UOffsetT) *AxisDetails {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxisDetails{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxisDetails(buf []byte, offset flatbuffers.UOffsetT) *AxisDetails {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxisDetails{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxisDetails) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxisDetails) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AxisDetails) SubDeviceName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AxisDetails) SubDeviceType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AxisDetails) CommAddress() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *AxisDetails) PhysicalAddress() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func AxisDetailsStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func AxisDetailsAddSubDeviceName(builder *flatbuffers.Builder, subDeviceName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(subDeviceName), 0)
}
func AxisDetailsAddSubDeviceType(builder *flatbuffers.Builder, subDeviceType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(subDeviceType), 0)
}
func AxisDetailsAddCommAddress(builder *flatbuffers.Builder, commAddress flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(commAddress), 0)
}
func AxisDetailsAddPhysicalAddress(builder *flatbuffers.Builder, physicalAddress flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(physicalAddress), 0)
}
func AxisDetailsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
