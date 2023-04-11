// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DeviceConfigT struct {
	Name string
	Description string
}

func (t *DeviceConfigT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	descriptionOffset := builder.CreateString(t.Description)
	DeviceConfigStart(builder)
	DeviceConfigAddName(builder, nameOffset)
	DeviceConfigAddDescription(builder, descriptionOffset)
	return DeviceConfigEnd(builder)
}

func (rcv *DeviceConfig) UnPackTo(t *DeviceConfigT) {
	t.Name = string(rcv.Name())
	t.Description = string(rcv.Description())
}

func (rcv *DeviceConfig) UnPack() *DeviceConfigT {
	if rcv == nil { return nil }
	t := &DeviceConfigT{}
	rcv.UnPackTo(t)
	return t
}

type DeviceConfig struct {
	_tab flatbuffers.Table
}

func GetRootAsDeviceConfig(buf []byte, offset flatbuffers.UOffsetT) *DeviceConfig {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DeviceConfig{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDeviceConfig(buf []byte, offset flatbuffers.UOffsetT) *DeviceConfig {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DeviceConfig{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DeviceConfig) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DeviceConfig) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DeviceConfig) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DeviceConfig) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DeviceConfigStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func DeviceConfigAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func DeviceConfigAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(description), 0)
}
func DeviceConfigEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
