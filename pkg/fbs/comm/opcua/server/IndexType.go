// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package server

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type IndexTypeT struct {
	Index uint32
}

func (t *IndexTypeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	IndexTypeStart(builder)
	IndexTypeAddIndex(builder, t.Index)
	return IndexTypeEnd(builder)
}

func (rcv *IndexType) UnPackTo(t *IndexTypeT) {
	t.Index = rcv.Index()
}

func (rcv *IndexType) UnPack() *IndexTypeT {
	if rcv == nil { return nil }
	t := &IndexTypeT{}
	rcv.UnPackTo(t)
	return t
}

type IndexType struct {
	_tab flatbuffers.Table
}

func GetRootAsIndexType(buf []byte, offset flatbuffers.UOffsetT) *IndexType {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IndexType{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsIndexType(buf []byte, offset flatbuffers.UOffsetT) *IndexType {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &IndexType{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *IndexType) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IndexType) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *IndexType) Index() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *IndexType) MutateIndex(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func IndexTypeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func IndexTypeAddIndex(builder *flatbuffers.Builder, index uint32) {
	builder.PrependUint32Slot(0, index, 0)
}
func IndexTypeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
