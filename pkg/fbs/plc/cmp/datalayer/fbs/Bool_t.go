// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Bool_tT struct {
	Data bool
}

func (t *Bool_tT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Bool_tStart(builder)
	Bool_tAddData(builder, t.Data)
	return Bool_tEnd(builder)
}

func (rcv *Bool_t) UnPackTo(t *Bool_tT) {
	t.Data = rcv.Data()
}

func (rcv *Bool_t) UnPack() *Bool_tT {
	if rcv == nil { return nil }
	t := &Bool_tT{}
	rcv.UnPackTo(t)
	return t
}

type Bool_t struct {
	_tab flatbuffers.Table
}

func GetRootAsBool_t(buf []byte, offset flatbuffers.UOffsetT) *Bool_t {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Bool_t{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBool_t(buf []byte, offset flatbuffers.UOffsetT) *Bool_t {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Bool_t{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Bool_t) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Bool_t) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Bool_t) Data() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Bool_t) MutateData(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func Bool_tStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Bool_tAddData(builder *flatbuffers.Builder, data bool) {
	builder.PrependBoolSlot(0, data, false)
}
func Bool_tEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
