// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Cmd_KinMoveDirectAsync struct {
	_tab flatbuffers.Table
}

func GetRootAsCmd_KinMoveDirectAsync(buf []byte, offset flatbuffers.UOffsetT) *Cmd_KinMoveDirectAsync {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Cmd_KinMoveDirectAsync{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmd_KinMoveDirectAsync(buf []byte, offset flatbuffers.UOffsetT) *Cmd_KinMoveDirectAsync {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Cmd_KinMoveDirectAsync{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Cmd_KinMoveDirectAsync) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Cmd_KinMoveDirectAsync) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Cmd_KinMoveDirectAsync) Base(obj *Cmd_Base) *Cmd_Base {
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

func (rcv *Cmd_KinMoveDirectAsync) Pos(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *Cmd_KinMoveDirectAsync) PosLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Cmd_KinMoveDirectAsync) MutatePos(j int, n float64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func (rcv *Cmd_KinMoveDirectAsync) CoordSys() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Cmd_KinMoveDirectAsync) DynLimFactors(obj *Cmd_DynLimits) *Cmd_DynLimits {
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

func (rcv *Cmd_KinMoveDirectAsync) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *Cmd_KinMoveDirectAsync) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

func Cmd_KinMoveDirectAsyncStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func Cmd_KinMoveDirectAsyncAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func Cmd_KinMoveDirectAsyncAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(pos), 0)
}
func Cmd_KinMoveDirectAsyncStartPosVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func Cmd_KinMoveDirectAsyncAddCoordSys(builder *flatbuffers.Builder, coordSys flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(coordSys), 0)
}
func Cmd_KinMoveDirectAsyncAddDynLimFactors(builder *flatbuffers.Builder, dynLimFactors flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(dynLimFactors), 0)
}
func Cmd_KinMoveDirectAsyncAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(4, buffered, true)
}
func Cmd_KinMoveDirectAsyncEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
