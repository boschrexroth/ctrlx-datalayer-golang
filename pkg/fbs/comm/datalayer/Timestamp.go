// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TimestampT struct {
	Value uint64
}

func (t *TimestampT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	TimestampStart(builder)
	TimestampAddValue(builder, t.Value)
	return TimestampEnd(builder)
}

func (rcv *Timestamp) UnPackTo(t *TimestampT) {
	t.Value = rcv.Value()
}

func (rcv *Timestamp) UnPack() *TimestampT {
	if rcv == nil { return nil }
	t := &TimestampT{}
	rcv.UnPackTo(t)
	return t
}

type Timestamp struct {
	_tab flatbuffers.Table
}

func GetRootAsTimestamp(buf []byte, offset flatbuffers.UOffsetT) *Timestamp {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Timestamp{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTimestamp(buf []byte, offset flatbuffers.UOffsetT) *Timestamp {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Timestamp{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Timestamp) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Timestamp) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Timestamp) Value() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Timestamp) MutateValue(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func TimestampStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TimestampAddValue(builder *flatbuffers.Builder, value uint64) {
	builder.PrependUint64Slot(0, value, 0)
}
func TimestampEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
