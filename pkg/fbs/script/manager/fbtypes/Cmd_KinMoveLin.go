// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd_KinMoveLinT struct {
	Base *Cmd_BaseT
	Pos []float64
	CoordSys string
	Lim *Cmd_DynLimitsT
	Buffered bool
}

func (t *Cmd_KinMoveLinT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	posOffset := flatbuffers.UOffsetT(0)
	if t.Pos != nil {
		posLength := len(t.Pos)
		Cmd_KinMoveLinStartPosVector(builder, posLength)
		for j := posLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.Pos[j])
		}
		posOffset = builder.EndVector(posLength)
	}
	coordSysOffset := builder.CreateString(t.CoordSys)
	limOffset := t.Lim.Pack(builder)
	Cmd_KinMoveLinStart(builder)
	Cmd_KinMoveLinAddBase(builder, baseOffset)
	Cmd_KinMoveLinAddPos(builder, posOffset)
	Cmd_KinMoveLinAddCoordSys(builder, coordSysOffset)
	Cmd_KinMoveLinAddLim(builder, limOffset)
	Cmd_KinMoveLinAddBuffered(builder, t.Buffered)
	return Cmd_KinMoveLinEnd(builder)
}

func (rcv *Cmd_KinMoveLin) UnPackTo(t *Cmd_KinMoveLinT) {
	t.Base = rcv.Base(nil).UnPack()
	posLength := rcv.PosLength()
	t.Pos = make([]float64, posLength)
	for j := 0; j < posLength; j++ {
		t.Pos[j] = rcv.Pos(j)
	}
	t.CoordSys = string(rcv.CoordSys())
	t.Lim = rcv.Lim(nil).UnPack()
	t.Buffered = rcv.Buffered()
}

func (rcv *Cmd_KinMoveLin) UnPack() *Cmd_KinMoveLinT {
	if rcv == nil { return nil }
	t := &Cmd_KinMoveLinT{}
	rcv.UnPackTo(t)
	return t
}

type Cmd_KinMoveLin struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd_KinMoveLin(buf []byte, offset flatbuffers.UOffsetT) *Cmd_KinMoveLin {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd_KinMoveLin{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd_KinMoveLin(buf []byte, offset flatbuffers.UOffsetT) *Cmd_KinMoveLin {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd_KinMoveLin{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd_KinMoveLin) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd_KinMoveLin) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd_KinMoveLin) Base(obj *Cmd_Base) *Cmd_Base {
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

func (rcv *Cmd_KinMoveLin) Pos(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *Cmd_KinMoveLin) PosLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Cmd_KinMoveLin) MutatePos(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func (rcv *Cmd_KinMoveLin) CoordSys() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Cmd_KinMoveLin) Lim(obj *Cmd_DynLimits) *Cmd_DynLimits {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Cmd_DynLimits)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Cmd_KinMoveLin) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *Cmd_KinMoveLin) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

func Cmd_KinMoveLinStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func Cmd_KinMoveLinAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func Cmd_KinMoveLinAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(pos), 0)
}
func Cmd_KinMoveLinStartPosVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func Cmd_KinMoveLinAddCoordSys(builder *flatbuffers.Builder, coordSys flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(coordSys), 0)
}
func Cmd_KinMoveLinAddLim(builder *flatbuffers.Builder, lim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(lim), 0)
}
func Cmd_KinMoveLinAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(4, buffered, true)
}
func Cmd_KinMoveLinEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
