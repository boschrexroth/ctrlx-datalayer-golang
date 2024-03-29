// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MailboxStatisticCountersT struct {
	Response *MailboxStatisticCountersResponseT `json:"response"`
}

func (t *MailboxStatisticCountersT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	responseOffset := t.Response.Pack(builder)
	MailboxStatisticCountersStart(builder)
	MailboxStatisticCountersAddResponse(builder, responseOffset)
	return MailboxStatisticCountersEnd(builder)
}

func (rcv *MailboxStatisticCounters) UnPackTo(t *MailboxStatisticCountersT) {
	t.Response = rcv.Response(nil).UnPack()
}

func (rcv *MailboxStatisticCounters) UnPack() *MailboxStatisticCountersT {
	if rcv == nil { return nil }
	t := &MailboxStatisticCountersT{}
	rcv.UnPackTo(t)
	return t
}

type MailboxStatisticCounters struct {
	_tab flatbuffers.Table
}

func GetRootAsMailboxStatisticCounters(buf []byte, offset flatbuffers.UOffsetT) *MailboxStatisticCounters {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MailboxStatisticCounters{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMailboxStatisticCounters(buf []byte, offset flatbuffers.UOffsetT) *MailboxStatisticCounters {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MailboxStatisticCounters{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MailboxStatisticCounters) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MailboxStatisticCounters) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MailboxStatisticCounters) Response(obj *MailboxStatisticCountersResponse) *MailboxStatisticCountersResponse {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MailboxStatisticCountersResponse)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func MailboxStatisticCountersStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MailboxStatisticCountersAddResponse(builder *flatbuffers.Builder, response flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(response), 0)
}
func MailboxStatisticCountersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
