// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TriggerOptionsT struct {
	QueueSize uint16 `json:"queueSize"`
}

func (t *TriggerOptionsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	TriggerOptionsStart(builder)
	TriggerOptionsAddQueueSize(builder, t.QueueSize)
	return TriggerOptionsEnd(builder)
}

func (rcv *TriggerOptions) UnPackTo(t *TriggerOptionsT) {
	t.QueueSize = rcv.QueueSize()
}

func (rcv *TriggerOptions) UnPack() *TriggerOptionsT {
	if rcv == nil { return nil }
	t := &TriggerOptionsT{}
	rcv.UnPackTo(t)
	return t
}

type TriggerOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsTriggerOptions(buf []byte, offset flatbuffers.UOffsetT) *TriggerOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TriggerOptions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTriggerOptions(buf []byte, offset flatbuffers.UOffsetT) *TriggerOptions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TriggerOptions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TriggerOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TriggerOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TriggerOptions) QueueSize() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 1
}

func (rcv *TriggerOptions) MutateQueueSize(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func TriggerOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TriggerOptionsAddQueueSize(builder *flatbuffers.Builder, queueSize uint16) {
	builder.PrependUint16Slot(0, queueSize, 1)
}
func TriggerOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
