// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Float32_tT struct {
	Data float32
}

func (t *Float32_tT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	Float32_tStart(builder)
	Float32_tAddData(builder, t.Data)
	return Float32_tEnd(builder)
}

func (rcv *Float32_t) UnPackTo(t *Float32_tT) {
	t.Data = rcv.Data()
}

func (rcv *Float32_t) UnPack() *Float32_tT {
	if rcv == nil { return nil }
	t := &Float32_tT{}
	rcv.UnPackTo(t)
	return t
}

type Float32_t struct {
	_tab flatbuffers.Table
}

func GetRootAsFloat32_t(buf []byte, offset flatbuffers.UOffsetT) *Float32_t {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Float32_t{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFloat32_t(buf []byte, offset flatbuffers.UOffsetT) *Float32_t {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Float32_t{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Float32_t) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Float32_t) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Float32_t) Data() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Float32_t) MutateData(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func Float32_tStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func Float32_tAddData(builder *flatbuffers.Builder, data float32) {
	builder.PrependFloat32Slot(0, data, 0.0)
}
func Float32_tEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}