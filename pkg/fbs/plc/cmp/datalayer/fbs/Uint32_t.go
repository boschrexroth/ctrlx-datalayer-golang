// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Uint32_tT struct {
	Data uint32
}

func (t *Uint32_tT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Uint32_tStart(builder)
	Uint32_tAddData(builder, t.Data)
	return Uint32_tEnd(builder)
}

func (rcv *Uint32_t) UnPackTo(t *Uint32_tT) {
	t.Data = rcv.Data()
}

func (rcv *Uint32_t) UnPack() *Uint32_tT {
	if rcv == nil { return nil }
	t := &Uint32_tT{}
	rcv.UnPackTo(t)
	return t
}

type Uint32_t struct {
	_tab flatbuffers.Table
}

func GetRootAsUint32_t(buf []byte, offset flatbuffers.UOffsetT) *Uint32_t {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Uint32_t{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUint32_t(buf []byte, offset flatbuffers.UOffsetT) *Uint32_t {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Uint32_t{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Uint32_t) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Uint32_t) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Uint32_t) Data() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Uint32_t) MutateData(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func Uint32_tStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Uint32_tAddData(builder *flatbuffers.Builder, data uint32) {
	builder.PrependUint32Slot(0, data, 0)
}
func Uint32_tEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
