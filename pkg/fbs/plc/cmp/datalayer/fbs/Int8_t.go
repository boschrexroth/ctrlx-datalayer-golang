// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Int8_tT struct {
	Data int8
}

func (t *Int8_tT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Int8_tStart(builder)
	Int8_tAddData(builder, t.Data)
	return Int8_tEnd(builder)
}

func (rcv *Int8_t) UnPackTo(t *Int8_tT) {
	t.Data = rcv.Data()
}

func (rcv *Int8_t) UnPack() *Int8_tT {
	if rcv == nil { return nil }
	t := &Int8_tT{}
	rcv.UnPackTo(t)
	return t
}

type Int8_t struct {
	_tab flatbuffers.Table
}

func GetRootAsInt8_t(buf []byte, offset flatbuffers.UOffsetT) *Int8_t {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Int8_t{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInt8_t(buf []byte, offset flatbuffers.UOffsetT) *Int8_t {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Int8_t{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Int8_t) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Int8_t) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Int8_t) Data() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Int8_t) MutateData(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func Int8_tStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Int8_tAddData(builder *flatbuffers.Builder, data int8) {
	builder.PrependInt8Slot(0, data, 0)
}
func Int8_tEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
