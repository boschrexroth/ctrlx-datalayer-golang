// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TimeT struct {
	Value uint64
}

func (t *TimeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	TimeStart(builder)
	TimeAddValue(builder, t.Value)
	return TimeEnd(builder)
}

func (rcv *Time) UnPackTo(t *TimeT) {
	t.Value = rcv.Value()
}

func (rcv *Time) UnPack() *TimeT {
	if rcv == nil { return nil }
	t := &TimeT{}
	rcv.UnPackTo(t)
	return t
}

type Time struct {
	_tab flatbuffers.Table
}

func GetRootAsTime(buf []byte, offset flatbuffers.UOffsetT) *Time {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Time{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTime(buf []byte, offset flatbuffers.UOffsetT) *Time {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Time{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Time) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Time) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Time) Value() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Time) MutateValue(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func TimeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TimeAddValue(builder *flatbuffers.Builder, value uint64) {
	builder.PrependUint64Slot(0, value, 0)
}
func TimeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
