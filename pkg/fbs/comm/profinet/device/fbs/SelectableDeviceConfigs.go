// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SelectableDeviceConfigsT struct {
	DeviceConfigsList []*DeviceConfigT `json:"deviceConfigsList"`
}

func (t *SelectableDeviceConfigsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	deviceConfigsListOffset := flatbuffers.UOffsetT(0)
	if t.DeviceConfigsList != nil {
		deviceConfigsListLength := len(t.DeviceConfigsList)
		deviceConfigsListOffsets := make([]flatbuffers.UOffsetT, deviceConfigsListLength)
		for j := 0; j < deviceConfigsListLength; j++ {
			deviceConfigsListOffsets[j] = t.DeviceConfigsList[j].Pack(builder)
		}
		SelectableDeviceConfigsStartDeviceConfigsListVector(builder, deviceConfigsListLength)
		for j := deviceConfigsListLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(deviceConfigsListOffsets[j])
		}
		deviceConfigsListOffset = builder.EndVector(deviceConfigsListLength)
	}
	SelectableDeviceConfigsStart(builder)
	SelectableDeviceConfigsAddDeviceConfigsList(builder, deviceConfigsListOffset)
	return SelectableDeviceConfigsEnd(builder)
}

func (rcv *SelectableDeviceConfigs) UnPackTo(t *SelectableDeviceConfigsT) {
	deviceConfigsListLength := rcv.DeviceConfigsListLength()
	t.DeviceConfigsList = make([]*DeviceConfigT, deviceConfigsListLength)
	for j := 0; j < deviceConfigsListLength; j++ {
		x := DeviceConfig{}
		rcv.DeviceConfigsList(&x, j)
		t.DeviceConfigsList[j] = x.UnPack()
	}
}

func (rcv *SelectableDeviceConfigs) UnPack() *SelectableDeviceConfigsT {
	if rcv == nil { return nil }
	t := &SelectableDeviceConfigsT{}
	rcv.UnPackTo(t)
	return t
}

type SelectableDeviceConfigs struct {
	_tab flatbuffers.Table
}

func GetRootAsSelectableDeviceConfigs(buf []byte, offset flatbuffers.UOffsetT) *SelectableDeviceConfigs {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SelectableDeviceConfigs{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSelectableDeviceConfigs(buf []byte, offset flatbuffers.UOffsetT) *SelectableDeviceConfigs {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SelectableDeviceConfigs{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SelectableDeviceConfigs) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SelectableDeviceConfigs) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SelectableDeviceConfigs) DeviceConfigsList(obj *DeviceConfig, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *SelectableDeviceConfigs) DeviceConfigsListLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func SelectableDeviceConfigsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SelectableDeviceConfigsAddDeviceConfigsList(builder *flatbuffers.Builder, deviceConfigsList flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(deviceConfigsList), 0)
}
func SelectableDeviceConfigsStartDeviceConfigsListVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SelectableDeviceConfigsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
