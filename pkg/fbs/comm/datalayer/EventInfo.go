// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type EventInfoT struct {
	EventType string `json:"eventType"`
	SourceNode string `json:"sourceNode"`
	SourceName string `json:"sourceName"`
	Timestamp uint64 `json:"timestamp"`
	SequenceNumber uint64 `json:"sequenceNumber"`
}

func (t *EventInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	eventTypeOffset := flatbuffers.UOffsetT(0)
	if t.EventType != "" {
		eventTypeOffset = builder.CreateString(t.EventType)
	}
	sourceNodeOffset := flatbuffers.UOffsetT(0)
	if t.SourceNode != "" {
		sourceNodeOffset = builder.CreateString(t.SourceNode)
	}
	sourceNameOffset := flatbuffers.UOffsetT(0)
	if t.SourceName != "" {
		sourceNameOffset = builder.CreateString(t.SourceName)
	}
	EventInfoStart(builder)
	EventInfoAddEventType(builder, eventTypeOffset)
	EventInfoAddSourceNode(builder, sourceNodeOffset)
	EventInfoAddSourceName(builder, sourceNameOffset)
	EventInfoAddTimestamp(builder, t.Timestamp)
	EventInfoAddSequenceNumber(builder, t.SequenceNumber)
	return EventInfoEnd(builder)
}

func (rcv *EventInfo) UnPackTo(t *EventInfoT) {
	t.EventType = string(rcv.EventType())
	t.SourceNode = string(rcv.SourceNode())
	t.SourceName = string(rcv.SourceName())
	t.Timestamp = rcv.Timestamp()
	t.SequenceNumber = rcv.SequenceNumber()
}

func (rcv *EventInfo) UnPack() *EventInfoT {
	if rcv == nil { return nil }
	t := &EventInfoT{}
	rcv.UnPackTo(t)
	return t
}

type EventInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsEventInfo(buf []byte, offset flatbuffers.UOffsetT) *EventInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EventInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEventInfo(buf []byte, offset flatbuffers.UOffsetT) *EventInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &EventInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *EventInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EventInfo) Table() flatbuffers.Table {
	return rcv._tab
}

/// The type of event, that is transfered.
/// E.g.: "types/events/ExampleEvent"
func (rcv *EventInfo) EventType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The type of event, that is transfered.
/// E.g.: "types/events/ExampleEvent"
/// The node, that fired the event.
/// E.g.: diagnosis/warning
func (rcv *EventInfo) SourceNode() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The node, that fired the event.
/// E.g.: diagnosis/warning
/// Description of the source of the event.
func (rcv *EventInfo) SourceName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Description of the source of the event.
/// optional, otherwise filled by provider.
/// Format is FILETIME
func (rcv *EventInfo) Timestamp() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// optional, otherwise filled by provider.
/// Format is FILETIME
func (rcv *EventInfo) MutateTimestamp(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

/// optional, otherwise filled by provider.
func (rcv *EventInfo) SequenceNumber() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// optional, otherwise filled by provider.
func (rcv *EventInfo) MutateSequenceNumber(n uint64) bool {
	return rcv._tab.MutateUint64Slot(12, n)
}

func EventInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func EventInfoAddEventType(builder *flatbuffers.Builder, eventType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(eventType), 0)
}
func EventInfoAddSourceNode(builder *flatbuffers.Builder, sourceNode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(sourceNode), 0)
}
func EventInfoAddSourceName(builder *flatbuffers.Builder, sourceName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(sourceName), 0)
}
func EventInfoAddTimestamp(builder *flatbuffers.Builder, timestamp uint64) {
	builder.PrependUint64Slot(3, timestamp, 0)
}
func EventInfoAddSequenceNumber(builder *flatbuffers.Builder, sequenceNumber uint64) {
	builder.PrependUint64Slot(4, sequenceNumber, 0)
}
func EventInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}