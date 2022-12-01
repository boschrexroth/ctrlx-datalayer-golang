// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BufferCfgT struct {
	RecordingTime *TimeT
	BufferType BufferTypeEnumFb
	RecordingInterval *TimeT
}

func (t *BufferCfgT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	recordingTimeOffset := t.RecordingTime.Pack(builder)
	recordingIntervalOffset := t.RecordingInterval.Pack(builder)
	BufferCfgStart(builder)
	BufferCfgAddRecordingTime(builder, recordingTimeOffset)
	BufferCfgAddBufferType(builder, t.BufferType)
	BufferCfgAddRecordingInterval(builder, recordingIntervalOffset)
	return BufferCfgEnd(builder)
}

func (rcv *BufferCfg) UnPackTo(t *BufferCfgT) {
	t.RecordingTime = rcv.RecordingTime(nil).UnPack()
	t.BufferType = rcv.BufferType()
	t.RecordingInterval = rcv.RecordingInterval(nil).UnPack()
}

func (rcv *BufferCfg) UnPack() *BufferCfgT {
	if rcv == nil { return nil }
	t := &BufferCfgT{}
	rcv.UnPackTo(t)
	return t
}

type BufferCfg struct {
	_tab flatbuffers.Table
}

func GetRootAsBufferCfg(buf []byte, offset flatbuffers.UOffsetT) *BufferCfg {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BufferCfg{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBufferCfg(buf []byte, offset flatbuffers.UOffsetT) *BufferCfg {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BufferCfg{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *BufferCfg) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BufferCfg) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BufferCfg) RecordingTime(obj *Time) *Time {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Time)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *BufferCfg) BufferType() BufferTypeEnumFb {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return BufferTypeEnumFb(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *BufferCfg) MutateBufferType(n BufferTypeEnumFb) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

func (rcv *BufferCfg) RecordingInterval(obj *Time) *Time {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Time)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func BufferCfgStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func BufferCfgAddRecordingTime(builder *flatbuffers.Builder, recordingTime flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(recordingTime), 0)
}
func BufferCfgAddBufferType(builder *flatbuffers.Builder, bufferType BufferTypeEnumFb) {
	builder.PrependInt8Slot(1, int8(bufferType), 0)
}
func BufferCfgAddRecordingInterval(builder *flatbuffers.Builder, recordingInterval flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(recordingInterval), 0)
}
func BufferCfgEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
