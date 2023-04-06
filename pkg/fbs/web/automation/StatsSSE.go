// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package automation

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StatsSSET struct {
	OpenSubscriptions uint32
	OpenSSE uint32
	RequestsSSE uint32
}

func (t *StatsSSET) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	StatsSSEStart(builder)
	StatsSSEAddOpenSubscriptions(builder, t.OpenSubscriptions)
	StatsSSEAddOpenSSE(builder, t.OpenSSE)
	StatsSSEAddRequestsSSE(builder, t.RequestsSSE)
	return StatsSSEEnd(builder)
}

func (rcv *StatsSSE) UnPackTo(t *StatsSSET) {
	t.OpenSubscriptions = rcv.OpenSubscriptions()
	t.OpenSSE = rcv.OpenSSE()
	t.RequestsSSE = rcv.RequestsSSE()
}

func (rcv *StatsSSE) UnPack() *StatsSSET {
	if rcv == nil { return nil }
	t := &StatsSSET{}
	rcv.UnPackTo(t)
	return t
}

type StatsSSE struct {
	_tab flatbuffers.Table
}

func GetRootAsStatsSSE(buf []byte, offset flatbuffers.UOffsetT) *StatsSSE {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StatsSSE{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStatsSSE(buf []byte, offset flatbuffers.UOffsetT) *StatsSSE {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StatsSSE{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StatsSSE) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StatsSSE) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StatsSSE) OpenSubscriptions() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *StatsSSE) MutateOpenSubscriptions(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *StatsSSE) OpenSSE() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *StatsSSE) MutateOpenSSE(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *StatsSSE) RequestsSSE() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *StatsSSE) MutateRequestsSSE(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func StatsSSEStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func StatsSSEAddOpenSubscriptions(builder *flatbuffers.Builder, openSubscriptions uint32) {
	builder.PrependUint32Slot(0, openSubscriptions, 0)
}
func StatsSSEAddOpenSSE(builder *flatbuffers.Builder, openSSE uint32) {
	builder.PrependUint32Slot(1, openSSE, 0)
}
func StatsSSEAddRequestsSSE(builder *flatbuffers.Builder, requestsSSE uint32) {
	builder.PrependUint32Slot(2, requestsSSE, 0)
}
func StatsSSEEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
