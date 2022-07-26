// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MessageDetailT struct {
	Timestamp uint64
}

func (t *MessageDetailT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	MessageDetailStart(builder)
	MessageDetailAddTimestamp(builder, t.Timestamp)
	return MessageDetailEnd(builder)
}

func (rcv *MessageDetail) UnPackTo(t *MessageDetailT) {
	t.Timestamp = rcv.Timestamp()
}

func (rcv *MessageDetail) UnPack() *MessageDetailT {
	if rcv == nil { return nil }
	t := &MessageDetailT{}
	rcv.UnPackTo(t)
	return t
}

type MessageDetail struct {
	_tab flatbuffers.Table
}

func GetRootAsMessageDetail(buf []byte, offset flatbuffers.UOffsetT) *MessageDetail {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MessageDetail{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMessageDetail(buf []byte, offset flatbuffers.UOffsetT) *MessageDetail {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MessageDetail{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MessageDetail) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MessageDetail) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MessageDetail) Timestamp() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MessageDetail) MutateTimestamp(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func MessageDetailStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MessageDetailAddTimestamp(builder *flatbuffers.Builder, timestamp uint64) {
	builder.PrependUint64Slot(0, timestamp, 0)
}
func MessageDetailEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
