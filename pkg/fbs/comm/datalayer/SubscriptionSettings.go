// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SubscriptionSettingsT struct {
	MinimumPublishInterval uint32 `json:"minimumPublishInterval"`
	MinimumSampleInterval uint64 `json:"minimumSampleInterval"`
	MaximumBufferSize uint32 `json:"maximumBufferSize"`
	MinimumErrorInterval uint32 `json:"minimumErrorInterval"`
	MaximumRtsubscribedNodes uint32 `json:"maximumRTSubscribedNodes"`
}

func (t *SubscriptionSettingsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	SubscriptionSettingsStart(builder)
	SubscriptionSettingsAddMinimumPublishInterval(builder, t.MinimumPublishInterval)
	SubscriptionSettingsAddMinimumSampleInterval(builder, t.MinimumSampleInterval)
	SubscriptionSettingsAddMaximumBufferSize(builder, t.MaximumBufferSize)
	SubscriptionSettingsAddMinimumErrorInterval(builder, t.MinimumErrorInterval)
	SubscriptionSettingsAddMaximumRtsubscribedNodes(builder, t.MaximumRtsubscribedNodes)
	return SubscriptionSettingsEnd(builder)
}

func (rcv *SubscriptionSettings) UnPackTo(t *SubscriptionSettingsT) {
	t.MinimumPublishInterval = rcv.MinimumPublishInterval()
	t.MinimumSampleInterval = rcv.MinimumSampleInterval()
	t.MaximumBufferSize = rcv.MaximumBufferSize()
	t.MinimumErrorInterval = rcv.MinimumErrorInterval()
	t.MaximumRtsubscribedNodes = rcv.MaximumRtsubscribedNodes()
}

func (rcv *SubscriptionSettings) UnPack() *SubscriptionSettingsT {
	if rcv == nil { return nil }
	t := &SubscriptionSettingsT{}
	rcv.UnPackTo(t)
	return t
}

type SubscriptionSettings struct {
	_tab flatbuffers.Table
}

func GetRootAsSubscriptionSettings(buf []byte, offset flatbuffers.UOffsetT) *SubscriptionSettings {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SubscriptionSettings{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSubscriptionSettings(buf []byte, offset flatbuffers.UOffsetT) *SubscriptionSettings {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SubscriptionSettings{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SubscriptionSettings) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SubscriptionSettings) Table() flatbuffers.Table {
	return rcv._tab
}

/// minimum publish interval in milliseconds
func (rcv *SubscriptionSettings) MinimumPublishInterval() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 50
}

/// minimum publish interval in milliseconds
func (rcv *SubscriptionSettings) MutateMinimumPublishInterval(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

/// minimum sampling interval in microseconds
func (rcv *SubscriptionSettings) MinimumSampleInterval() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 100000
}

/// minimum sampling interval in microseconds
func (rcv *SubscriptionSettings) MutateMinimumSampleInterval(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

/// maximum size of buffer
func (rcv *SubscriptionSettings) MaximumBufferSize() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 50
}

/// maximum size of buffer
func (rcv *SubscriptionSettings) MutateMaximumBufferSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

/// minimum error interval
func (rcv *SubscriptionSettings) MinimumErrorInterval() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 10000
}

/// minimum error interval
func (rcv *SubscriptionSettings) MutateMinimumErrorInterval(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

/// Maximum allowed count of RT Subscriptions to addresses
func (rcv *SubscriptionSettings) MaximumRtsubscribedNodes() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 50
}

/// Maximum allowed count of RT Subscriptions to addresses
func (rcv *SubscriptionSettings) MutateMaximumRtsubscribedNodes(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func SubscriptionSettingsStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func SubscriptionSettingsAddMinimumPublishInterval(builder *flatbuffers.Builder, minimumPublishInterval uint32) {
	builder.PrependUint32Slot(0, minimumPublishInterval, 50)
}
func SubscriptionSettingsAddMinimumSampleInterval(builder *flatbuffers.Builder, minimumSampleInterval uint64) {
	builder.PrependUint64Slot(1, minimumSampleInterval, 100000)
}
func SubscriptionSettingsAddMaximumBufferSize(builder *flatbuffers.Builder, maximumBufferSize uint32) {
	builder.PrependUint32Slot(2, maximumBufferSize, 50)
}
func SubscriptionSettingsAddMinimumErrorInterval(builder *flatbuffers.Builder, minimumErrorInterval uint32) {
	builder.PrependUint32Slot(3, minimumErrorInterval, 10000)
}
func SubscriptionSettingsAddMaximumRtsubscribedNodes(builder *flatbuffers.Builder, maximumRtsubscribedNodes uint32) {
	builder.PrependUint32Slot(4, maximumRtsubscribedNodes, 50)
}
func SubscriptionSettingsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
