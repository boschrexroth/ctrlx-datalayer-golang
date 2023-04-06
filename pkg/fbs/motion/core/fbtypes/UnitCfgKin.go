// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// General unit configuration for a kinematics object
type UnitCfgKinT struct {
	Default []*UnitCfgObjSingleT
	Position []string
}

func (t *UnitCfgKinT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	defaultOffset := flatbuffers.UOffsetT(0)
	if t.Default != nil {
		defaultLength := len(t.Default)
		defaultOffsets := make([]flatbuffers.UOffsetT, defaultLength)
		for j := 0; j < defaultLength; j++ {
			defaultOffsets[j] = t.Default[j].Pack(builder)
		}
		UnitCfgKinStartDefaultVector(builder, defaultLength)
		for j := defaultLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(defaultOffsets[j])
		}
		defaultOffset = builder.EndVector(defaultLength)
	}
	positionOffset := flatbuffers.UOffsetT(0)
	if t.Position != nil {
		positionLength := len(t.Position)
		positionOffsets := make([]flatbuffers.UOffsetT, positionLength)
		for j := 0; j < positionLength; j++ {
			positionOffsets[j] = builder.CreateString(t.Position[j])
		}
		UnitCfgKinStartPositionVector(builder, positionLength)
		for j := positionLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(positionOffsets[j])
		}
		positionOffset = builder.EndVector(positionLength)
	}
	UnitCfgKinStart(builder)
	UnitCfgKinAddDefault(builder, defaultOffset)
	UnitCfgKinAddPosition(builder, positionOffset)
	return UnitCfgKinEnd(builder)
}

func (rcv *UnitCfgKin) UnPackTo(t *UnitCfgKinT) {
	defaultLength := rcv.DefaultLength()
	t.Default = make([]*UnitCfgObjSingleT, defaultLength)
	for j := 0; j < defaultLength; j++ {
		x := UnitCfgObjSingle{}
		rcv.Default(&x, j)
		t.Default[j] = x.UnPack()
	}
	positionLength := rcv.PositionLength()
	t.Position = make([]string, positionLength)
	for j := 0; j < positionLength; j++ {
		t.Position[j] = string(rcv.Position(j))
	}
}

func (rcv *UnitCfgKin) UnPack() *UnitCfgKinT {
	if rcv == nil { return nil }
	t := &UnitCfgKinT{}
	rcv.UnPackTo(t)
	return t
}

type UnitCfgKin struct {
	_tab flatbuffers.Table
}

func GetRootAsUnitCfgKin(buf []byte, offset flatbuffers.UOffsetT) *UnitCfgKin {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UnitCfgKin{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUnitCfgKin(buf []byte, offset flatbuffers.UOffsetT) *UnitCfgKin {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &UnitCfgKin{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *UnitCfgKin) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UnitCfgKin) Table() flatbuffers.Table {
	return rcv._tab
}

/// default units
func (rcv *UnitCfgKin) Default(obj *UnitCfgObjSingle, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *UnitCfgKin) DefaultLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// default units
/// position unit abbreviations (must be 16 entries)
func (rcv *UnitCfgKin) Position(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *UnitCfgKin) PositionLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// position unit abbreviations (must be 16 entries)
func UnitCfgKinStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func UnitCfgKinAddDefault(builder *flatbuffers.Builder, default_ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(default_), 0)
}
func UnitCfgKinStartDefaultVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func UnitCfgKinAddPosition(builder *flatbuffers.Builder, position flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(position), 0)
}
func UnitCfgKinStartPositionVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func UnitCfgKinEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
