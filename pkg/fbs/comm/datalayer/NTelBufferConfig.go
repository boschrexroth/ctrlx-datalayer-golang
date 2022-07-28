// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type NTelBufferConfigT struct {
	DefaultN uint16
}

func (t *NTelBufferConfigT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	NTelBufferConfigStart(builder)
	NTelBufferConfigAddDefaultN(builder, t.DefaultN)
	return NTelBufferConfigEnd(builder)
}

func (rcv *NTelBufferConfig) UnPackTo(t *NTelBufferConfigT) {
	t.DefaultN = rcv.DefaultN()
}

func (rcv *NTelBufferConfig) UnPack() *NTelBufferConfigT {
	if rcv == nil { return nil }
	t := &NTelBufferConfigT{}
	rcv.UnPackTo(t)
	return t
}

type NTelBufferConfig struct {
	_tab flatbuffers.Table
}

func GetRootAsNTelBufferConfig(buf []byte, offset flatbuffers.UOffsetT) *NTelBufferConfig {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NTelBufferConfig{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsNTelBufferConfig(buf []byte, offset flatbuffers.UOffsetT) *NTelBufferConfig {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &NTelBufferConfig{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *NTelBufferConfig) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NTelBufferConfig) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NTelBufferConfig) DefaultN() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 6
}

func (rcv *NTelBufferConfig) MutateDefaultN(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func NTelBufferConfigStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func NTelBufferConfigAddDefaultN(builder *flatbuffers.Builder, defaultN uint16) {
	builder.PrependUint16Slot(0, defaultN, 6)
}
func NTelBufferConfigEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}