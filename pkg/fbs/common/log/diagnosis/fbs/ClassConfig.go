// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// This table defines which log class is selected.
type ClassConfigT struct {
	ClassConfig Class
}

func (t *ClassConfigT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ClassConfigStart(builder)
	ClassConfigAddClassConfig(builder, t.ClassConfig)
	return ClassConfigEnd(builder)
}

func (rcv *ClassConfig) UnPackTo(t *ClassConfigT) {
	t.ClassConfig = rcv.ClassConfig()
}

func (rcv *ClassConfig) UnPack() *ClassConfigT {
	if rcv == nil { return nil }
	t := &ClassConfigT{}
	rcv.UnPackTo(t)
	return t
}

type ClassConfig struct {
	_tab flatbuffers.Table
}

func GetRootAsClassConfig(buf []byte, offset flatbuffers.UOffsetT) *ClassConfig {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ClassConfig{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsClassConfig(buf []byte, offset flatbuffers.UOffsetT) *ClassConfig {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ClassConfig{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ClassConfig) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ClassConfig) Table() flatbuffers.Table {
	return rcv._tab
}

/// Selection of log class.
func (rcv *ClassConfig) ClassConfig() Class {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return Class(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 2
}

/// Selection of log class.
func (rcv *ClassConfig) MutateClassConfig(n Class) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func ClassConfigStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ClassConfigAddClassConfig(builder *flatbuffers.Builder, classConfig Class) {
	builder.PrependInt8Slot(0, int8(classConfig), 2)
}
func ClassConfigEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
