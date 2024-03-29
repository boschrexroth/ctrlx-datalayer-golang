// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DL_WriteT struct {
	Path string `json:"path"`
	Type int32 `json:"type"`
	S string `json:"s"`
	I int64 `json:"i"`
	B bool `json:"b"`
	D float64 `json:"d"`
}

func (t *DL_WriteT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	pathOffset := flatbuffers.UOffsetT(0)
	if t.Path != "" {
		pathOffset = builder.CreateString(t.Path)
	}
	sOffset := flatbuffers.UOffsetT(0)
	if t.S != "" {
		sOffset = builder.CreateString(t.S)
	}
	DL_WriteStart(builder)
	DL_WriteAddPath(builder, pathOffset)
	DL_WriteAddType(builder, t.Type)
	DL_WriteAddS(builder, sOffset)
	DL_WriteAddI(builder, t.I)
	DL_WriteAddB(builder, t.B)
	DL_WriteAddD(builder, t.D)
	return DL_WriteEnd(builder)
}

func (rcv *DL_Write) UnPackTo(t *DL_WriteT) {
	t.Path = string(rcv.Path())
	t.Type = rcv.Type()
	t.S = string(rcv.S())
	t.I = rcv.I()
	t.B = rcv.B()
	t.D = rcv.D()
}

func (rcv *DL_Write) UnPack() *DL_WriteT {
	if rcv == nil { return nil }
	t := &DL_WriteT{}
	rcv.UnPackTo(t)
	return t
}

type DL_Write struct {
	_tab flatbuffers.Table
}

func GetRootAsDL_Write(buf []byte, offset flatbuffers.UOffsetT) *DL_Write {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DL_Write{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDL_Write(buf []byte, offset flatbuffers.UOffsetT) *DL_Write {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DL_Write{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DL_Write) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DL_Write) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DL_Write) Path() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DL_Write) Type() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DL_Write) MutateType(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *DL_Write) S() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DL_Write) I() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DL_Write) MutateI(n int64) bool {
	return rcv._tab.MutateInt64Slot(10, n)
}

func (rcv *DL_Write) B() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *DL_Write) MutateB(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

func (rcv *DL_Write) D() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *DL_Write) MutateD(n float64) bool {
	return rcv._tab.MutateFloat64Slot(14, n)
}

func DL_WriteStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func DL_WriteAddPath(builder *flatbuffers.Builder, path flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(path), 0)
}
func DL_WriteAddType(builder *flatbuffers.Builder, type_ int32) {
	builder.PrependInt32Slot(1, type_, 0)
}
func DL_WriteAddS(builder *flatbuffers.Builder, s flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(s), 0)
}
func DL_WriteAddI(builder *flatbuffers.Builder, i int64) {
	builder.PrependInt64Slot(3, i, 0)
}
func DL_WriteAddB(builder *flatbuffers.Builder, b bool) {
	builder.PrependBoolSlot(4, b, false)
}
func DL_WriteAddD(builder *flatbuffers.Builder, d float64) {
	builder.PrependFloat64Slot(5, d, 0.0)
}
func DL_WriteEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
