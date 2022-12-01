// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ServerSettingsT struct {
	ServerIdlePingTimeout uint32
	ServerWaitResponseTimeout uint32
	ServerMaxMessageSize uint32
	DebugAddress string
	ServerMaxRtSize uint32
}

func (t *ServerSettingsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	debugAddressOffset := builder.CreateString(t.DebugAddress)
	ServerSettingsStart(builder)
	ServerSettingsAddServerIdlePingTimeout(builder, t.ServerIdlePingTimeout)
	ServerSettingsAddServerWaitResponseTimeout(builder, t.ServerWaitResponseTimeout)
	ServerSettingsAddServerMaxMessageSize(builder, t.ServerMaxMessageSize)
	ServerSettingsAddDebugAddress(builder, debugAddressOffset)
	ServerSettingsAddServerMaxRtSize(builder, t.ServerMaxRtSize)
	return ServerSettingsEnd(builder)
}

func (rcv *ServerSettings) UnPackTo(t *ServerSettingsT) {
	t.ServerIdlePingTimeout = rcv.ServerIdlePingTimeout()
	t.ServerWaitResponseTimeout = rcv.ServerWaitResponseTimeout()
	t.ServerMaxMessageSize = rcv.ServerMaxMessageSize()
	t.DebugAddress = string(rcv.DebugAddress())
	t.ServerMaxRtSize = rcv.ServerMaxRtSize()
}

func (rcv *ServerSettings) UnPack() *ServerSettingsT {
	if rcv == nil { return nil }
	t := &ServerSettingsT{}
	rcv.UnPackTo(t)
	return t
}

type ServerSettings struct {
	_tab flatbuffers.Table
}

func GetRootAsServerSettings(buf []byte, offset flatbuffers.UOffsetT) *ServerSettings {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ServerSettings{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsServerSettings(buf []byte, offset flatbuffers.UOffsetT) *ServerSettings {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ServerSettings{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ServerSettings) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ServerSettings) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ServerSettings) ServerIdlePingTimeout() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 30000
}

func (rcv *ServerSettings) MutateServerIdlePingTimeout(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *ServerSettings) ServerWaitResponseTimeout() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 3000
}

func (rcv *ServerSettings) MutateServerWaitResponseTimeout(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *ServerSettings) ServerMaxMessageSize() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 52428800
}

func (rcv *ServerSettings) MutateServerMaxMessageSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *ServerSettings) DebugAddress() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ServerSettings) ServerMaxRtSize() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 1048576
}

func (rcv *ServerSettings) MutateServerMaxRtSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func ServerSettingsStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ServerSettingsAddServerIdlePingTimeout(builder *flatbuffers.Builder, serverIdlePingTimeout uint32) {
	builder.PrependUint32Slot(0, serverIdlePingTimeout, 30000)
}
func ServerSettingsAddServerWaitResponseTimeout(builder *flatbuffers.Builder, serverWaitResponseTimeout uint32) {
	builder.PrependUint32Slot(1, serverWaitResponseTimeout, 3000)
}
func ServerSettingsAddServerMaxMessageSize(builder *flatbuffers.Builder, serverMaxMessageSize uint32) {
	builder.PrependUint32Slot(2, serverMaxMessageSize, 52428800)
}
func ServerSettingsAddDebugAddress(builder *flatbuffers.Builder, debugAddress flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(debugAddress), 0)
}
func ServerSettingsAddServerMaxRtSize(builder *flatbuffers.Builder, serverMaxRtSize uint32) {
	builder.PrependUint32Slot(4, serverMaxRtSize, 1048576)
}
func ServerSettingsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
