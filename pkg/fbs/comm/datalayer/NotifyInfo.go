// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NotifyInfoT struct {
	Node string `json:"node"`
	Timestamp uint64 `json:"timestamp"`
	NotifyType NotifyType `json:"notifyType"`
	EventType string `json:"eventType"`
	SequenceNumber uint64 `json:"sequenceNumber"`
	SourceName string `json:"sourceName"`
}

func (t *NotifyInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nodeOffset := flatbuffers.UOffsetT(0)
	if t.Node != "" {
		nodeOffset = builder.CreateString(t.Node)
	}
	eventTypeOffset := flatbuffers.UOffsetT(0)
	if t.EventType != "" {
		eventTypeOffset = builder.CreateString(t.EventType)
	}
	sourceNameOffset := flatbuffers.UOffsetT(0)
	if t.SourceName != "" {
		sourceNameOffset = builder.CreateString(t.SourceName)
	}
	NotifyInfoStart(builder)
	NotifyInfoAddNode(builder, nodeOffset)
	NotifyInfoAddTimestamp(builder, t.Timestamp)
	NotifyInfoAddNotifyType(builder, t.NotifyType)
	NotifyInfoAddEventType(builder, eventTypeOffset)
	NotifyInfoAddSequenceNumber(builder, t.SequenceNumber)
	NotifyInfoAddSourceName(builder, sourceNameOffset)
	return NotifyInfoEnd(builder)
}

func (rcv *NotifyInfo) UnPackTo(t *NotifyInfoT) {
	t.Node = string(rcv.Node())
	t.Timestamp = rcv.Timestamp()
	t.NotifyType = rcv.NotifyType()
	t.EventType = string(rcv.EventType())
	t.SequenceNumber = rcv.SequenceNumber()
	t.SourceName = string(rcv.SourceName())
}

func (rcv *NotifyInfo) UnPack() *NotifyInfoT {
	if rcv == nil { return nil }
	t := &NotifyInfoT{}
	rcv.UnPackTo(t)
	return t
}

type NotifyInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsNotifyInfo(buf []byte, offset flatbuffers.UOffsetT) *NotifyInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NotifyInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNotifyInfo(buf []byte, offset flatbuffers.UOffsetT) *NotifyInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &NotifyInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *NotifyInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NotifyInfo) Table() flatbuffers.Table {
	return rcv._tab
}

/// Node address
func (rcv *NotifyInfo) Node() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Node address
/// Filetime: Contains a 64-bit value representing the number of 100-nanosecond intervals since January 1, 1601 (UTC).
func (rcv *NotifyInfo) Timestamp() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Filetime: Contains a 64-bit value representing the number of 100-nanosecond intervals since January 1, 1601 (UTC).
func (rcv *NotifyInfo) MutateTimestamp(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *NotifyInfo) NotifyType() NotifyType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return NotifyType(rcv._tab.GetInt32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *NotifyInfo) MutateNotifyType(n NotifyType) bool {
	return rcv._tab.MutateInt32Slot(8, int32(n))
}

/// In case of an event, this string contains the information what EventType has been fired. E.g.: "types/events/ExampleEvent"
func (rcv *NotifyInfo) EventType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// In case of an event, this string contains the information what EventType has been fired. E.g.: "types/events/ExampleEvent"
/// sequence number of an event
func (rcv *NotifyInfo) SequenceNumber() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// sequence number of an event
func (rcv *NotifyInfo) MutateSequenceNumber(n uint64) bool {
	return rcv._tab.MutateUint64Slot(12, n)
}

/// Description of the source of an event
func (rcv *NotifyInfo) SourceName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Description of the source of an event
func NotifyInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func NotifyInfoAddNode(builder *flatbuffers.Builder, node flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(node), 0)
}
func NotifyInfoAddTimestamp(builder *flatbuffers.Builder, timestamp uint64) {
	builder.PrependUint64Slot(1, timestamp, 0)
}
func NotifyInfoAddNotifyType(builder *flatbuffers.Builder, notifyType NotifyType) {
	builder.PrependInt32Slot(2, int32(notifyType), 0)
}
func NotifyInfoAddEventType(builder *flatbuffers.Builder, eventType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(eventType), 0)
}
func NotifyInfoAddSequenceNumber(builder *flatbuffers.Builder, sequenceNumber uint64) {
	builder.PrependUint64Slot(4, sequenceNumber, 0)
}
func NotifyInfoAddSourceName(builder *flatbuffers.Builder, sourceName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(sourceName), 0)
}
func NotifyInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
