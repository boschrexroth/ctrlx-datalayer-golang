// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd_KinMoveDirectT struct {
	Base *Cmd_BaseT
	Pos []float64
	CoordSys string
	Buffered bool
}

func (t *Cmd_KinMoveDirectT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	posOffset := flatbuffers.UOffsetT(0)
	if t.Pos != nil {
		posLength := len(t.Pos)
		Cmd_KinMoveDirectStartPosVector(builder, posLength)
		for j := posLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.Pos[j])
		}
		posOffset = builder.EndVector(posLength)
	}
	coordSysOffset := builder.CreateString(t.CoordSys)
	Cmd_KinMoveDirectStart(builder)
	Cmd_KinMoveDirectAddBase(builder, baseOffset)
	Cmd_KinMoveDirectAddPos(builder, posOffset)
	Cmd_KinMoveDirectAddCoordSys(builder, coordSysOffset)
	Cmd_KinMoveDirectAddBuffered(builder, t.Buffered)
	return Cmd_KinMoveDirectEnd(builder)
}

func (rcv *Cmd_KinMoveDirect) UnPackTo(t *Cmd_KinMoveDirectT) {
	t.Base = rcv.Base(nil).UnPack()
	posLength := rcv.PosLength()
	t.Pos = make([]float64, posLength)
	for j := 0; j < posLength; j++ {
		t.Pos[j] = rcv.Pos(j)
	}
	t.CoordSys = string(rcv.CoordSys())
	t.Buffered = rcv.Buffered()
}

func (rcv *Cmd_KinMoveDirect) UnPack() *Cmd_KinMoveDirectT {
	if rcv == nil { return nil }
	t := &Cmd_KinMoveDirectT{}
	rcv.UnPackTo(t)
	return t
}

type Cmd_KinMoveDirect struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd_KinMoveDirect(buf []byte, offset flatbuffers.UOffsetT) *Cmd_KinMoveDirect {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd_KinMoveDirect{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd_KinMoveDirect(buf []byte, offset flatbuffers.UOffsetT) *Cmd_KinMoveDirect {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd_KinMoveDirect{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd_KinMoveDirect) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd_KinMoveDirect) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd_KinMoveDirect) Base(obj *Cmd_Base) *Cmd_Base {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Cmd_Base)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Cmd_KinMoveDirect) Pos(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *Cmd_KinMoveDirect) PosLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Cmd_KinMoveDirect) MutatePos(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func (rcv *Cmd_KinMoveDirect) CoordSys() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Cmd_KinMoveDirect) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *Cmd_KinMoveDirect) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func Cmd_KinMoveDirectStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func Cmd_KinMoveDirectAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func Cmd_KinMoveDirectAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(pos), 0)
}
func Cmd_KinMoveDirectStartPosVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func Cmd_KinMoveDirectAddCoordSys(builder *flatbuffers.Builder, coordSys flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(coordSys), 0)
}
func Cmd_KinMoveDirectAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(3, buffered, true)
}
func Cmd_KinMoveDirectEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
