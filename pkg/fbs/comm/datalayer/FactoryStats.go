// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FactoryStatsT struct {
	NumClients uint32
	NumProviders uint32
	OpenClientRequests uint32
	OpenProviderRequests uint32
}

func (t *FactoryStatsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	FactoryStatsStart(builder)
	FactoryStatsAddNumClients(builder, t.NumClients)
	FactoryStatsAddNumProviders(builder, t.NumProviders)
	FactoryStatsAddOpenClientRequests(builder, t.OpenClientRequests)
	FactoryStatsAddOpenProviderRequests(builder, t.OpenProviderRequests)
	return FactoryStatsEnd(builder)
}

func (rcv *FactoryStats) UnPackTo(t *FactoryStatsT) {
	t.NumClients = rcv.NumClients()
	t.NumProviders = rcv.NumProviders()
	t.OpenClientRequests = rcv.OpenClientRequests()
	t.OpenProviderRequests = rcv.OpenProviderRequests()
}

func (rcv *FactoryStats) UnPack() *FactoryStatsT {
	if rcv == nil { return nil }
	t := &FactoryStatsT{}
	rcv.UnPackTo(t)
	return t
}

type FactoryStats struct {
	_tab flatbuffers.Table
}

func GetRootAsFactoryStats(buf []byte, offset flatbuffers.UOffsetT) *FactoryStats {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FactoryStats{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFactoryStats(buf []byte, offset flatbuffers.UOffsetT) *FactoryStats {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FactoryStats{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FactoryStats) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FactoryStats) Table() flatbuffers.Table {
	return rcv._tab
}

/// number of clients
func (rcv *FactoryStats) NumClients() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// number of clients
func (rcv *FactoryStats) MutateNumClients(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

/// number of provider
func (rcv *FactoryStats) NumProviders() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// number of provider
func (rcv *FactoryStats) MutateNumProviders(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

/// number of open requests from clients
func (rcv *FactoryStats) OpenClientRequests() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// number of open requests from clients
func (rcv *FactoryStats) MutateOpenClientRequests(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

/// number of open requests from provider
func (rcv *FactoryStats) OpenProviderRequests() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// number of open requests from provider
func (rcv *FactoryStats) MutateOpenProviderRequests(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func FactoryStatsStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func FactoryStatsAddNumClients(builder *flatbuffers.Builder, numClients uint32) {
	builder.PrependUint32Slot(0, numClients, 0)
}
func FactoryStatsAddNumProviders(builder *flatbuffers.Builder, numProviders uint32) {
	builder.PrependUint32Slot(1, numProviders, 0)
}
func FactoryStatsAddOpenClientRequests(builder *flatbuffers.Builder, openClientRequests uint32) {
	builder.PrependUint32Slot(2, openClientRequests, 0)
}
func FactoryStatsAddOpenProviderRequests(builder *flatbuffers.Builder, openProviderRequests uint32) {
	builder.PrependUint32Slot(3, openProviderRequests, 0)
}
func FactoryStatsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
