// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type InstancesCreateRequestT struct {
	InstanceName string `json:"instanceName"`
	Port string `json:"port"`
	LinkLayer string `json:"linkLayer"`
	Arguments string `json:"arguments"`
}

func (t *InstancesCreateRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	instanceNameOffset := flatbuffers.UOffsetT(0)
	if t.InstanceName != "" {
		instanceNameOffset = builder.CreateString(t.InstanceName)
	}
	portOffset := flatbuffers.UOffsetT(0)
	if t.Port != "" {
		portOffset = builder.CreateString(t.Port)
	}
	linkLayerOffset := flatbuffers.UOffsetT(0)
	if t.LinkLayer != "" {
		linkLayerOffset = builder.CreateString(t.LinkLayer)
	}
	argumentsOffset := flatbuffers.UOffsetT(0)
	if t.Arguments != "" {
		argumentsOffset = builder.CreateString(t.Arguments)
	}
	InstancesCreateRequestStart(builder)
	InstancesCreateRequestAddInstanceName(builder, instanceNameOffset)
	InstancesCreateRequestAddPort(builder, portOffset)
	InstancesCreateRequestAddLinkLayer(builder, linkLayerOffset)
	InstancesCreateRequestAddArguments(builder, argumentsOffset)
	return InstancesCreateRequestEnd(builder)
}

func (rcv *InstancesCreateRequest) UnPackTo(t *InstancesCreateRequestT) {
	t.InstanceName = string(rcv.InstanceName())
	t.Port = string(rcv.Port())
	t.LinkLayer = string(rcv.LinkLayer())
	t.Arguments = string(rcv.Arguments())
}

func (rcv *InstancesCreateRequest) UnPack() *InstancesCreateRequestT {
	if rcv == nil { return nil }
	t := &InstancesCreateRequestT{}
	rcv.UnPackTo(t)
	return t
}

type InstancesCreateRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsInstancesCreateRequest(buf []byte, offset flatbuffers.UOffsetT) *InstancesCreateRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &InstancesCreateRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInstancesCreateRequest(buf []byte, offset flatbuffers.UOffsetT) *InstancesCreateRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &InstancesCreateRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *InstancesCreateRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *InstancesCreateRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *InstancesCreateRequest) InstanceName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *InstancesCreateRequest) Port() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *InstancesCreateRequest) LinkLayer() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *InstancesCreateRequest) Arguments() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func InstancesCreateRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func InstancesCreateRequestAddInstanceName(builder *flatbuffers.Builder, instanceName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(instanceName), 0)
}
func InstancesCreateRequestAddPort(builder *flatbuffers.Builder, port flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(port), 0)
}
func InstancesCreateRequestAddLinkLayer(builder *flatbuffers.Builder, linkLayer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(linkLayer), 0)
}
func InstancesCreateRequestAddArguments(builder *flatbuffers.Builder, arguments flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(arguments), 0)
}
func InstancesCreateRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
