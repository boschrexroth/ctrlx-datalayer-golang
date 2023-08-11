// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package automation

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StatsT struct {
	Sse *StatsSSET `json:"sse"`
	DlClient *StatsDlClientT `json:"dlClient"`
	Restbed *StatsRestbedT `json:"restbed"`
	RequestsV1 uint32 `json:"requestsV1"`
	RequestsV2 uint32 `json:"requestsV2"`
}

func (t *StatsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	sseOffset := t.Sse.Pack(builder)
	dlClientOffset := t.DlClient.Pack(builder)
	restbedOffset := t.Restbed.Pack(builder)
	StatsStart(builder)
	StatsAddSse(builder, sseOffset)
	StatsAddDlClient(builder, dlClientOffset)
	StatsAddRestbed(builder, restbedOffset)
	StatsAddRequestsV1(builder, t.RequestsV1)
	StatsAddRequestsV2(builder, t.RequestsV2)
	return StatsEnd(builder)
}

func (rcv *Stats) UnPackTo(t *StatsT) {
	t.Sse = rcv.Sse(nil).UnPack()
	t.DlClient = rcv.DlClient(nil).UnPack()
	t.Restbed = rcv.Restbed(nil).UnPack()
	t.RequestsV1 = rcv.RequestsV1()
	t.RequestsV2 = rcv.RequestsV2()
}

func (rcv *Stats) UnPack() *StatsT {
	if rcv == nil { return nil }
	t := &StatsT{}
	rcv.UnPackTo(t)
	return t
}

type Stats struct {
	_tab flatbuffers.Table
}

func GetRootAsStats(buf []byte, offset flatbuffers.UOffsetT) *Stats {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Stats{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStats(buf []byte, offset flatbuffers.UOffsetT) *Stats {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Stats{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Stats) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Stats) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Stats) Sse(obj *StatsSSE) *StatsSSE {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(StatsSSE)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Stats) DlClient(obj *StatsDlClient) *StatsDlClient {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(StatsDlClient)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Stats) Restbed(obj *StatsRestbed) *StatsRestbed {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(StatsRestbed)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Stats) RequestsV1() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateRequestsV1(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *Stats) RequestsV2() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stats) MutateRequestsV2(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func StatsStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func StatsAddSse(builder *flatbuffers.Builder, sse flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(sse), 0)
}
func StatsAddDlClient(builder *flatbuffers.Builder, dlClient flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(dlClient), 0)
}
func StatsAddRestbed(builder *flatbuffers.Builder, restbed flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(restbed), 0)
}
func StatsAddRequestsV1(builder *flatbuffers.Builder, requestsV1 uint32) {
	builder.PrependUint32Slot(3, requestsV1, 0)
}
func StatsAddRequestsV2(builder *flatbuffers.Builder, requestsV2 uint32) {
	builder.PrependUint32Slot(4, requestsV2, 0)
}
func StatsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
