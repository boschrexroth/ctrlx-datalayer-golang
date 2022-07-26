// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters of the axis position commands
type AxsCmdHomingDataT struct {
	Buffered bool
	NewRefPos float64
}

func (t *AxsCmdHomingDataT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	AxsCmdHomingDataStart(builder)
	AxsCmdHomingDataAddBuffered(builder, t.Buffered)
	AxsCmdHomingDataAddNewRefPos(builder, t.NewRefPos)
	return AxsCmdHomingDataEnd(builder)
}

func (rcv *AxsCmdHomingData) UnPackTo(t *AxsCmdHomingDataT) {
	t.Buffered = rcv.Buffered()
	t.NewRefPos = rcv.NewRefPos()
}

func (rcv *AxsCmdHomingData) UnPack() *AxsCmdHomingDataT {
	if rcv == nil { return nil }
	t := &AxsCmdHomingDataT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCmdHomingData struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCmdHomingData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdHomingData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCmdHomingData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCmdHomingData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdHomingData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCmdHomingData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCmdHomingData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCmdHomingData) Table() flatbuffers.Table {
	return rcv._tab
}

/// buffered type for generic command
func (rcv *AxsCmdHomingData) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// buffered type for generic command
func (rcv *AxsCmdHomingData) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

/// commanded target position (or increment for relative position command)
func (rcv *AxsCmdHomingData) NewRefPos() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// commanded target position (or increment for relative position command)
func (rcv *AxsCmdHomingData) MutateNewRefPos(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

func AxsCmdHomingDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AxsCmdHomingDataAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(0, buffered, false)
}
func AxsCmdHomingDataAddNewRefPos(builder *flatbuffers.Builder, newRefPos float64) {
	builder.PrependFloat64Slot(1, newRefPos, 0.0)
}
func AxsCmdHomingDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
