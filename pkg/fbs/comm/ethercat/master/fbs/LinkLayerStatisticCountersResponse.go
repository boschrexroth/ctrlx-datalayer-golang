// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Link layer statistic counters
type LinkLayerStatisticCountersResponseT struct {
	PhysicalErrorCnt uint32 `json:"physicalErrorCnt"`
	TelegramErrorCnt uint32 `json:"telegramErrorCnt"`
	SendOffsetMeasurementActive bool `json:"sendOffsetMeasurementActive"`
	SendOffset *MinActMaxValuesT `json:"sendOffset"`
	NetworkDelayMeasurementActive bool `json:"networkDelayMeasurementActive"`
	NetworkDelay *MinActMaxValuesT `json:"networkDelay"`
}

func (t *LinkLayerStatisticCountersResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	LinkLayerStatisticCountersResponseStart(builder)
	LinkLayerStatisticCountersResponseAddPhysicalErrorCnt(builder, t.PhysicalErrorCnt)
	LinkLayerStatisticCountersResponseAddTelegramErrorCnt(builder, t.TelegramErrorCnt)
	LinkLayerStatisticCountersResponseAddSendOffsetMeasurementActive(builder, t.SendOffsetMeasurementActive)
	sendOffsetOffset := t.SendOffset.Pack(builder)
	LinkLayerStatisticCountersResponseAddSendOffset(builder, sendOffsetOffset)
	LinkLayerStatisticCountersResponseAddNetworkDelayMeasurementActive(builder, t.NetworkDelayMeasurementActive)
	networkDelayOffset := t.NetworkDelay.Pack(builder)
	LinkLayerStatisticCountersResponseAddNetworkDelay(builder, networkDelayOffset)
	return LinkLayerStatisticCountersResponseEnd(builder)
}

func (rcv *LinkLayerStatisticCountersResponse) UnPackTo(t *LinkLayerStatisticCountersResponseT) {
	t.PhysicalErrorCnt = rcv.PhysicalErrorCnt()
	t.TelegramErrorCnt = rcv.TelegramErrorCnt()
	t.SendOffsetMeasurementActive = rcv.SendOffsetMeasurementActive()
	t.SendOffset = rcv.SendOffset(nil).UnPack()
	t.NetworkDelayMeasurementActive = rcv.NetworkDelayMeasurementActive()
	t.NetworkDelay = rcv.NetworkDelay(nil).UnPack()
}

func (rcv *LinkLayerStatisticCountersResponse) UnPack() *LinkLayerStatisticCountersResponseT {
	if rcv == nil { return nil }
	t := &LinkLayerStatisticCountersResponseT{}
	rcv.UnPackTo(t)
	return t
}

type LinkLayerStatisticCountersResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsLinkLayerStatisticCountersResponse(buf []byte, offset flatbuffers.UOffsetT) *LinkLayerStatisticCountersResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LinkLayerStatisticCountersResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLinkLayerStatisticCountersResponse(buf []byte, offset flatbuffers.UOffsetT) *LinkLayerStatisticCountersResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LinkLayerStatisticCountersResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LinkLayerStatisticCountersResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LinkLayerStatisticCountersResponse) Table() flatbuffers.Table {
	return rcv._tab
}

///Physical error counter
func (rcv *LinkLayerStatisticCountersResponse) PhysicalErrorCnt() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Physical error counter
func (rcv *LinkLayerStatisticCountersResponse) MutatePhysicalErrorCnt(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

///Telegram error counter
func (rcv *LinkLayerStatisticCountersResponse) TelegramErrorCnt() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Telegram error counter
func (rcv *LinkLayerStatisticCountersResponse) MutateTelegramErrorCnt(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

///Measurement of send offset is active
func (rcv *LinkLayerStatisticCountersResponse) SendOffsetMeasurementActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

///Measurement of send offset is active
func (rcv *LinkLayerStatisticCountersResponse) MutateSendOffsetMeasurementActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

/// Send offset values in nanoseconds
func (rcv *LinkLayerStatisticCountersResponse) SendOffset(obj *MinActMaxValues) *MinActMaxValues {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(MinActMaxValues)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Send offset values in nanoseconds
///Measurement of network delay is active
func (rcv *LinkLayerStatisticCountersResponse) NetworkDelayMeasurementActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

///Measurement of network delay is active
func (rcv *LinkLayerStatisticCountersResponse) MutateNetworkDelayMeasurementActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

/// Network delay values in nanoseconds
func (rcv *LinkLayerStatisticCountersResponse) NetworkDelay(obj *MinActMaxValues) *MinActMaxValues {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(MinActMaxValues)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Network delay values in nanoseconds
func LinkLayerStatisticCountersResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func LinkLayerStatisticCountersResponseAddPhysicalErrorCnt(builder *flatbuffers.Builder, physicalErrorCnt uint32) {
	builder.PrependUint32Slot(0, physicalErrorCnt, 0)
}
func LinkLayerStatisticCountersResponseAddTelegramErrorCnt(builder *flatbuffers.Builder, telegramErrorCnt uint32) {
	builder.PrependUint32Slot(1, telegramErrorCnt, 0)
}
func LinkLayerStatisticCountersResponseAddSendOffsetMeasurementActive(builder *flatbuffers.Builder, sendOffsetMeasurementActive bool) {
	builder.PrependBoolSlot(2, sendOffsetMeasurementActive, false)
}
func LinkLayerStatisticCountersResponseAddSendOffset(builder *flatbuffers.Builder, sendOffset flatbuffers.UOffsetT) {
	builder.PrependStructSlot(3, flatbuffers.UOffsetT(sendOffset), 0)
}
func LinkLayerStatisticCountersResponseAddNetworkDelayMeasurementActive(builder *flatbuffers.Builder, networkDelayMeasurementActive bool) {
	builder.PrependBoolSlot(4, networkDelayMeasurementActive, false)
}
func LinkLayerStatisticCountersResponseAddNetworkDelay(builder *flatbuffers.Builder, networkDelay flatbuffers.UOffsetT) {
	builder.PrependStructSlot(5, flatbuffers.UOffsetT(networkDelay), 0)
}
func LinkLayerStatisticCountersResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
