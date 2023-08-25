// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package framework

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type InterfaceList_T struct {
	Interfaces []*Interface_T `json:"interfaces"`
}

func (t *InterfaceList_T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	interfacesOffset := flatbuffers.UOffsetT(0)
	if t.Interfaces != nil {
		interfacesLength := len(t.Interfaces)
		interfacesOffsets := make([]flatbuffers.UOffsetT, interfacesLength)
		for j := 0; j < interfacesLength; j++ {
			interfacesOffsets[j] = t.Interfaces[j].Pack(builder)
		}
		InterfaceList_StartInterfacesVector(builder, interfacesLength)
		for j := interfacesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(interfacesOffsets[j])
		}
		interfacesOffset = builder.EndVector(interfacesLength)
	}
	InterfaceList_Start(builder)
	InterfaceList_AddInterfaces(builder, interfacesOffset)
	return InterfaceList_End(builder)
}

func (rcv *InterfaceList_) UnPackTo(t *InterfaceList_T) {
	interfacesLength := rcv.InterfacesLength()
	t.Interfaces = make([]*Interface_T, interfacesLength)
	for j := 0; j < interfacesLength; j++ {
		x := Interface_{}
		rcv.Interfaces(&x, j)
		t.Interfaces[j] = x.UnPack()
	}
}

func (rcv *InterfaceList_) UnPack() *InterfaceList_T {
	if rcv == nil { return nil }
	t := &InterfaceList_T{}
	rcv.UnPackTo(t)
	return t
}

type InterfaceList_ struct {
	_tab flatbuffers.Table
}

func GetRootAsInterfaceList_(buf []byte, offset flatbuffers.UOffsetT) *InterfaceList_ {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &InterfaceList_{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInterfaceList_(buf []byte, offset flatbuffers.UOffsetT) *InterfaceList_ {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &InterfaceList_{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *InterfaceList_) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *InterfaceList_) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *InterfaceList_) Interfaces(obj *Interface_, j int) bool {
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

func (rcv *InterfaceList_) InterfacesByKey(obj *Interface_, key string) bool{
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *InterfaceList_) InterfacesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func InterfaceList_Start(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func InterfaceList_AddInterfaces(builder *flatbuffers.Builder, interfaces flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(interfaces), 0)
}
func InterfaceList_StartInterfacesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func InterfaceList_End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
