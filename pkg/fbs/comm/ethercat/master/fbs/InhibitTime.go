// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///This read only node returns the inhibit time of the slave statistic single shot update mode
type InhibitTimeT struct {
	Response *InhibitTimeResponseT `json:"response"`
}

func (t *InhibitTimeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	responseOffset := t.Response.Pack(builder)
	InhibitTimeStart(builder)
	InhibitTimeAddResponse(builder, responseOffset)
	return InhibitTimeEnd(builder)
}

func (rcv *InhibitTime) UnPackTo(t *InhibitTimeT) {
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *InhibitTime) UnPack() *InhibitTimeT {
	if rcv == nil { return nil }
	t := &InhibitTimeT{}
	rcv.UnPackTo(t)
	return t
}

type InhibitTime struct {
	_tab flatbuffers.Table
}

func GetRootAsInhibitTime(buf []byte, offset flatbuffers.UOffsetT) *InhibitTime {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &InhibitTime{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInhibitTime(buf []byte, offset flatbuffers.UOffsetT) *InhibitTime {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &InhibitTime{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *InhibitTime) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *InhibitTime) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *InhibitTime) Response(obj *InhibitTimeResponse) *InhibitTimeResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(InhibitTimeResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func InhibitTimeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func InhibitTimeAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(response), 0)
}
func InhibitTimeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
