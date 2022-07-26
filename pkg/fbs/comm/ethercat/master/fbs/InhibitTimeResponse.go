// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type InhibitTimeResponseT struct {
	Time uint32
}

func (t *InhibitTimeResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	InhibitTimeResponseStart(builder)
	InhibitTimeResponseAddTime(builder, t.Time)
	return InhibitTimeResponseEnd(builder)
}

func (rcv *InhibitTimeResponse) UnPackTo(t *InhibitTimeResponseT) {
	t.Time = rcv.Time()
}

func (rcv *InhibitTimeResponse) UnPack() *InhibitTimeResponseT {
	if rcv == nil { return nil }
	t := &InhibitTimeResponseT{}
	rcv.UnPackTo(t)
	return t
}

type InhibitTimeResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsInhibitTimeResponse(buf []byte, offset flatbuffers.UOffsetT) *InhibitTimeResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &InhibitTimeResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInhibitTimeResponse(buf []byte, offset flatbuffers.UOffsetT) *InhibitTimeResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &InhibitTimeResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *InhibitTimeResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *InhibitTimeResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *InhibitTimeResponse) Time() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *InhibitTimeResponse) MutateTime(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func InhibitTimeResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func InhibitTimeResponseAddTime(builder *flatbuffers.Builder, time uint32) {
	builder.PrependUint32Slot(0, time, 0)
}
func InhibitTimeResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
