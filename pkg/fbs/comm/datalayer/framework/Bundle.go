// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package framework

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BundleT struct {
	Name string
	Version string
	Location string
	Id int64
	Components []*ComponentT
	State string
	Active bool
	Installed bool
}

func (t *BundleT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	versionOffset := builder.CreateString(t.Version)
	locationOffset := builder.CreateString(t.Location)
	componentsOffset := flatbuffers.UOffsetT(0)
	if t.Components != nil {
		componentsLength := len(t.Components)
		componentsOffsets := make([]flatbuffers.UOffsetT, componentsLength)
		for j := 0; j < componentsLength; j++ {
			componentsOffsets[j] = t.Components[j].Pack(builder)
		}
		BundleStartComponentsVector(builder, componentsLength)
		for j := componentsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(componentsOffsets[j])
		}
		componentsOffset = builder.EndVector(componentsLength)
	}
	stateOffset := builder.CreateString(t.State)
	BundleStart(builder)
	BundleAddName(builder, nameOffset)
	BundleAddVersion(builder, versionOffset)
	BundleAddLocation(builder, locationOffset)
	BundleAddId(builder, t.Id)
	BundleAddComponents(builder, componentsOffset)
	BundleAddState(builder, stateOffset)
	BundleAddActive(builder, t.Active)
	BundleAddInstalled(builder, t.Installed)
	return BundleEnd(builder)
}

func (rcv *Bundle) UnPackTo(t *BundleT) {
	t.Name = string(rcv.Name())
	t.Version = string(rcv.Version())
	t.Location = string(rcv.Location())
	t.Id = rcv.Id()
	componentsLength := rcv.ComponentsLength()
	t.Components = make([]*ComponentT, componentsLength)
	for j := 0; j < componentsLength; j++ {
		x := Component{}
		rcv.Components(&x, j)
		t.Components[j] = x.UnPack()
	}
	t.State = string(rcv.State())
	t.Active = rcv.Active()
	t.Installed = rcv.Installed()
}

func (rcv *Bundle) UnPack() *BundleT {
	if rcv == nil { return nil }
	t := &BundleT{}
	rcv.UnPackTo(t)
	return t
}

type Bundle struct {
	_tab flatbuffers.Table
}

func GetRootAsBundle(buf []byte, offset flatbuffers.UOffsetT) *Bundle {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Bundle{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBundle(buf []byte, offset flatbuffers.UOffsetT) *Bundle {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Bundle{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Bundle) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Bundle) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Bundle) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Bundle) Version() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Bundle) Location() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Bundle) Id() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Bundle) MutateId(n int64) bool {
	return rcv._tab.MutateInt64Slot(10, n)
}

func (rcv *Bundle) Components(obj *Component, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Bundle) ComponentsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Bundle) State() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Bundle) Active() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Bundle) MutateActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(16, n)
}

func (rcv *Bundle) Installed() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Bundle) MutateInstalled(n bool) bool {
	return rcv._tab.MutateBoolSlot(18, n)
}

func BundleStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func BundleAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func BundleAddVersion(builder *flatbuffers.Builder, version flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(version), 0)
}
func BundleAddLocation(builder *flatbuffers.Builder, location flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(location), 0)
}
func BundleAddId(builder *flatbuffers.Builder, id int64) {
	builder.PrependInt64Slot(3, id, 0)
}
func BundleAddComponents(builder *flatbuffers.Builder, components flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(components), 0)
}
func BundleStartComponentsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BundleAddState(builder *flatbuffers.Builder, state flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(state), 0)
}
func BundleAddActive(builder *flatbuffers.Builder, active bool) {
	builder.PrependBoolSlot(6, active, false)
}
func BundleAddInstalled(builder *flatbuffers.Builder, installed bool) {
	builder.PrependBoolSlot(7, installed, false)
}
func BundleEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
