// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BusLoadStatisticCountersT struct {
	Response *BusLoadStatisticCountersResponseT
}

func (t *BusLoadStatisticCountersT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	responseOffset := t.Response.Pack(builder)
	BusLoadStatisticCountersStart(builder)
	BusLoadStatisticCountersAddResponse(builder, responseOffset)
	return BusLoadStatisticCountersEnd(builder)
}

func (rcv *BusLoadStatisticCounters) UnPackTo(t *BusLoadStatisticCountersT) {
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *BusLoadStatisticCounters) UnPack() *BusLoadStatisticCountersT {
	if rcv == nil { return nil }
	t := &BusLoadStatisticCountersT{}
	rcv.UnPackTo(t)
	return t
}

type BusLoadStatisticCounters struct {
	_tab flatbuffers.Table
}

func GetRootAsBusLoadStatisticCounters(buf []byte, offset flatbuffers.UOffsetT) *BusLoadStatisticCounters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BusLoadStatisticCounters{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBusLoadStatisticCounters(buf []byte, offset flatbuffers.UOffsetT) *BusLoadStatisticCounters {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BusLoadStatisticCounters{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *BusLoadStatisticCounters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BusLoadStatisticCounters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BusLoadStatisticCounters) Response(obj *BusLoadStatisticCountersResponse) *BusLoadStatisticCountersResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BusLoadStatisticCountersResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func BusLoadStatisticCountersStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func BusLoadStatisticCountersAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(response), 0)
}
func BusLoadStatisticCountersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
