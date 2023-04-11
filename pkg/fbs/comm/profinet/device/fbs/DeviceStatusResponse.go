// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DeviceStatusResponseT struct {
	Status uint32
}

func (t *DeviceStatusResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	DeviceStatusResponseStart(builder)
	DeviceStatusResponseAddStatus(builder, t.Status)
	return DeviceStatusResponseEnd(builder)
}

func (rcv *DeviceStatusResponse) UnPackTo(t *DeviceStatusResponseT) {
	t.Status = rcv.Status()
}

func (rcv *DeviceStatusResponse) UnPack() *DeviceStatusResponseT {
	if rcv == nil { return nil }
	t := &DeviceStatusResponseT{}
	rcv.UnPackTo(t)
	return t
}

type DeviceStatusResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsDeviceStatusResponse(buf []byte, offset flatbuffers.UOffsetT) *DeviceStatusResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DeviceStatusResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDeviceStatusResponse(buf []byte, offset flatbuffers.UOffsetT) *DeviceStatusResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DeviceStatusResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DeviceStatusResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DeviceStatusResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DeviceStatusResponse) Status() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DeviceStatusResponse) MutateStatus(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func DeviceStatusResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DeviceStatusResponseAddStatus(builder *flatbuffers.Builder, status uint32) {
	builder.PrependUint32Slot(0, status, 0)
}
func DeviceStatusResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
