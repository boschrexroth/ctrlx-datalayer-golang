// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ProtocolStatisticPeriodResetT struct {
	Total *ProtocolStatisticResetFlagsT
	LastSecond *ProtocolStatisticResetFlagsT
}

func (t *ProtocolStatisticPeriodResetT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	return CreateProtocolStatisticPeriodReset(builder, t.Total.NumRequests, t.Total.NumBytes, t.LastSecond.NumRequests, t.LastSecond.NumBytes)
}
func (rcv *ProtocolStatisticPeriodReset) UnPackTo(t *ProtocolStatisticPeriodResetT) {
	t.Total = rcv.Total(nil).UnPack()
	t.LastSecond = rcv.LastSecond(nil).UnPack()
}

func (rcv *ProtocolStatisticPeriodReset) UnPack() *ProtocolStatisticPeriodResetT {
	if rcv == nil { return nil }
	t := &ProtocolStatisticPeriodResetT{}
	rcv.UnPackTo(t)
	return t
}

type ProtocolStatisticPeriodReset struct {
	_tab flatbuffers.Struct
}

func (rcv *ProtocolStatisticPeriodReset) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ProtocolStatisticPeriodReset) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *ProtocolStatisticPeriodReset) Total(obj *ProtocolStatisticResetFlags) *ProtocolStatisticResetFlags {
	if obj == nil {
		obj = new(ProtocolStatisticResetFlags)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+0)
	return obj
}
func (rcv *ProtocolStatisticPeriodReset) LastSecond(obj *ProtocolStatisticResetFlags) *ProtocolStatisticResetFlags {
	if obj == nil {
		obj = new(ProtocolStatisticResetFlags)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+2)
	return obj
}

func CreateProtocolStatisticPeriodReset(builder *flatbuffers.Builder, total_numRequests bool, total_numBytes bool, lastSecond_numRequests bool, lastSecond_numBytes bool) flatbuffers.UOffsetT {
	builder.Prep(1, 4)
	builder.Prep(1, 2)
	builder.PrependBool(lastSecond_numBytes)
	builder.PrependBool(lastSecond_numRequests)
	builder.Prep(1, 2)
	builder.PrependBool(total_numBytes)
	builder.PrependBool(total_numRequests)
	return builder.Offset()
}
