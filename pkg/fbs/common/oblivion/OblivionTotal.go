// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package oblivion

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OblivionTotalT struct {
	Overhead uint32 `json:"overhead"`
	Unobserved uint32 `json:"unobserved"`
}

func (t *OblivionTotalT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	OblivionTotalStart(builder)
	OblivionTotalAddOverhead(builder, t.Overhead)
	OblivionTotalAddUnobserved(builder, t.Unobserved)
	return OblivionTotalEnd(builder)
}

func (rcv *OblivionTotal) UnPackTo(t *OblivionTotalT) {
	t.Overhead = rcv.Overhead()
	t.Unobserved = rcv.Unobserved()
}

func (rcv *OblivionTotal) UnPack() *OblivionTotalT {
	if rcv == nil { return nil }
	t := &OblivionTotalT{}
	rcv.UnPackTo(t)
	return t
}

type OblivionTotal struct {
	_tab flatbuffers.Table
}

func GetRootAsOblivionTotal(buf []byte, offset flatbuffers.UOffsetT) *OblivionTotal {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OblivionTotal{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsOblivionTotal(buf []byte, offset flatbuffers.UOffsetT) *OblivionTotal {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &OblivionTotal{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *OblivionTotal) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OblivionTotal) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OblivionTotal) Overhead() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OblivionTotal) MutateOverhead(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *OblivionTotal) Unobserved() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *OblivionTotal) MutateUnobserved(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func OblivionTotalStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func OblivionTotalAddOverhead(builder *flatbuffers.Builder, overhead uint32) {
	builder.PrependUint32Slot(0, overhead, 0)
}
func OblivionTotalAddUnobserved(builder *flatbuffers.Builder, unobserved uint32) {
	builder.PrependUint32Slot(1, unobserved, 0)
}
func OblivionTotalEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
