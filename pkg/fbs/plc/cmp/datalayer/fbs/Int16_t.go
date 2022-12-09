// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Int16_tT struct {
	Data int16
}

func (t *Int16_tT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Int16_tStart(builder)
	Int16_tAddData(builder, t.Data)
	return Int16_tEnd(builder)
}

func (rcv *Int16_t) UnPackTo(t *Int16_tT) {
	t.Data = rcv.Data()
}

func (rcv *Int16_t) UnPack() *Int16_tT {
	if rcv == nil { return nil }
	t := &Int16_tT{}
	rcv.UnPackTo(t)
	return t
}

type Int16_t struct {
	_tab flatbuffers.Table
}

func GetRootAsInt16_t(buf []byte, offset flatbuffers.UOffsetT) *Int16_t {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Int16_t{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInt16_t(buf []byte, offset flatbuffers.UOffsetT) *Int16_t {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Int16_t{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Int16_t) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Int16_t) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Int16_t) Data() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Int16_t) MutateData(n int16) bool {
	return rcv._tab.MutateInt16Slot(4, n)
}

func Int16_tStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Int16_tAddData(builder *flatbuffers.Builder, data int16) {
	builder.PrependInt16Slot(0, data, 0)
}
func Int16_tEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
