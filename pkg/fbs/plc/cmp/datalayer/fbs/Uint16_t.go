// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Uint16_tT struct {
	Data uint16 `json:"data"`
}

func (t *Uint16_tT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Uint16_tStart(builder)
	Uint16_tAddData(builder, t.Data)
	return Uint16_tEnd(builder)
}

func (rcv *Uint16_t) UnPackTo(t *Uint16_tT) {
	t.Data = rcv.Data()
}

func (rcv *Uint16_t) UnPack() *Uint16_tT {
	if rcv == nil { return nil }
	t := &Uint16_tT{}
	rcv.UnPackTo(t)
	return t
}

type Uint16_t struct {
	_tab flatbuffers.Table
}

func GetRootAsUint16_t(buf []byte, offset flatbuffers.UOffsetT) *Uint16_t {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Uint16_t{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUint16_t(buf []byte, offset flatbuffers.UOffsetT) *Uint16_t {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Uint16_t{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Uint16_t) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Uint16_t) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Uint16_t) Data() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Uint16_t) MutateData(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func Uint16_tStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Uint16_tAddData(builder *flatbuffers.Builder, data uint16) {
	builder.PrependUint16Slot(0, data, 0)
}
func Uint16_tEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
