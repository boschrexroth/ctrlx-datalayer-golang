// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd_KinJogT struct {
	Base *Cmd_BaseT
	Dir []float64
	CoordSys string
	Incr float64
	Lim *Cmd_DynLimitsT
}

func (t *Cmd_KinJogT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	dirOffset := flatbuffers.UOffsetT(0)
	if t.Dir != nil {
		dirLength := len(t.Dir)
		Cmd_KinJogStartDirVector(builder, dirLength)
		for j := dirLength - 1; j >= 0; j-- {
			builder.PrependFloat64(t.Dir[j])
		}
		dirOffset = builder.EndVector(dirLength)
	}
	coordSysOffset := builder.CreateString(t.CoordSys)
	limOffset := t.Lim.Pack(builder)
	Cmd_KinJogStart(builder)
	Cmd_KinJogAddBase(builder, baseOffset)
	Cmd_KinJogAddDir(builder, dirOffset)
	Cmd_KinJogAddCoordSys(builder, coordSysOffset)
	Cmd_KinJogAddIncr(builder, t.Incr)
	Cmd_KinJogAddLim(builder, limOffset)
	return Cmd_KinJogEnd(builder)
}

func (rcv *Cmd_KinJog) UnPackTo(t *Cmd_KinJogT) {
	t.Base = rcv.Base(nil).UnPack()
	dirLength := rcv.DirLength()
	t.Dir = make([]float64, dirLength)
	for j := 0; j < dirLength; j++ {
		t.Dir[j] = rcv.Dir(j)
	}
	t.CoordSys = string(rcv.CoordSys())
	t.Incr = rcv.Incr()
	t.Lim = rcv.Lim(nil).UnPack()
}

func (rcv *Cmd_KinJog) UnPack() *Cmd_KinJogT {
	if rcv == nil { return nil }
	t := &Cmd_KinJogT{}
	rcv.UnPackTo(t)
	return t
}

type Cmd_KinJog struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd_KinJog(buf []byte, offset flatbuffers.UOffsetT) *Cmd_KinJog {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd_KinJog{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd_KinJog(buf []byte, offset flatbuffers.UOffsetT) *Cmd_KinJog {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd_KinJog{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd_KinJog) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd_KinJog) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd_KinJog) Base(obj *Cmd_Base) *Cmd_Base {
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

func (rcv *Cmd_KinJog) Dir(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *Cmd_KinJog) DirLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Cmd_KinJog) MutateDir(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func (rcv *Cmd_KinJog) CoordSys() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Cmd_KinJog) Incr() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Cmd_KinJog) MutateIncr(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

func (rcv *Cmd_KinJog) Lim(obj *Cmd_DynLimits) *Cmd_DynLimits {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
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

func Cmd_KinJogStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func Cmd_KinJogAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func Cmd_KinJogAddDir(builder *flatbuffers.Builder, dir flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(dir), 0)
}
func Cmd_KinJogStartDirVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func Cmd_KinJogAddCoordSys(builder *flatbuffers.Builder, coordSys flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(coordSys), 0)
}
func Cmd_KinJogAddIncr(builder *flatbuffers.Builder, incr float64) {
	builder.PrependFloat64Slot(3, incr, 0.0)
}
func Cmd_KinJogAddLim(builder *flatbuffers.Builder, lim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(lim), 0)
}
func Cmd_KinJogEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}