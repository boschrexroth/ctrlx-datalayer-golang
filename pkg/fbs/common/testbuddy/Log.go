// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package testbuddy

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Log struct {
	_tab flatbuffers.Table
}

func GetRootAsLog(buf []byte, offset flatbuffers.UOffsetT) *Log {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Log{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLog(buf []byte, offset flatbuffers.UOffsetT) *Log {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Log{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Log) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Log) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Log) Index() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Log) MutateIndex(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Log) Type() LogType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return LogType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Log) MutateType(n LogType) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

func (rcv *Log) Log() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func LogStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func LogAddIndex(builder *flatbuffers.Builder, index uint32) {
	builder.PrependUint32Slot(0, index, 0)
}
func LogAddType(builder *flatbuffers.Builder, type_ LogType) {
	builder.PrependInt8Slot(1, int8(type_), 0)
}
func LogAddLog(builder *flatbuffers.Builder, log flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(log), 0)
}
func LogEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}