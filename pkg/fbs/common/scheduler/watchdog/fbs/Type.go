// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Type of watchdog of the task
type TypeT struct {
	Type CurrentType `json:"type"`
}

func (t *TypeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	TypeStart(builder)
	TypeAddType(builder, t.Type)
	return TypeEnd(builder)
}

func (rcv *Type) UnPackTo(t *TypeT) {
	t.Type = rcv.Type()
}

func (rcv *Type) UnPack() *TypeT {
	if rcv == nil { return nil }
	t := &TypeT{}
	rcv.UnPackTo(t)
	return t
}

type Type struct {
	_tab flatbuffers.Table
}

func GetRootAsType(buf []byte, offset flatbuffers.UOffsetT) *Type {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Type{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsType(buf []byte, offset flatbuffers.UOffsetT) *Type {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Type{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Type) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Type) Table() flatbuffers.Table {
	return rcv._tab
}

/// Type of watchdog of the task
func (rcv *Type) Type() CurrentType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return CurrentType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// Type of watchdog of the task
func (rcv *Type) MutateType(n CurrentType) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func TypeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TypeAddType(builder *flatbuffers.Builder, type_ CurrentType) {
	builder.PrependInt8Slot(0, int8(type_), 0)
}
func TypeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
