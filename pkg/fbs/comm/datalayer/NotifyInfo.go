// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NotifyInfoT struct {
	Node string `json:"node"`
	Timestamp uint64 `json:"timestamp"`
	NotifyType NotifyType `json:"notifyType"`
}

func (t *NotifyInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nodeOffset := flatbuffers.UOffsetT(0)
	if t.Node != "" {
		nodeOffset = builder.CreateString(t.Node)
	}
	NotifyInfoStart(builder)
	NotifyInfoAddNode(builder, nodeOffset)
	NotifyInfoAddTimestamp(builder, t.Timestamp)
	NotifyInfoAddNotifyType(builder, t.NotifyType)
	return NotifyInfoEnd(builder)
}

func (rcv *NotifyInfo) UnPackTo(t *NotifyInfoT) {
	t.Node = string(rcv.Node())
	t.Timestamp = rcv.Timestamp()
	t.NotifyType = rcv.NotifyType()
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

func NotifyInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
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
func NotifyInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
