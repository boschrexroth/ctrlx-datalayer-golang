// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package framework

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DependencyListT struct {
	Dependencies []*DependencyT `json:"dependencies"`
}

func (t *DependencyListT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dependenciesOffset := flatbuffers.UOffsetT(0)
	if t.Dependencies != nil {
		dependenciesLength := len(t.Dependencies)
		dependenciesOffsets := make([]flatbuffers.UOffsetT, dependenciesLength)
		for j := 0; j < dependenciesLength; j++ {
			dependenciesOffsets[j] = t.Dependencies[j].Pack(builder)
		}
		DependencyListStartDependenciesVector(builder, dependenciesLength)
		for j := dependenciesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(dependenciesOffsets[j])
		}
		dependenciesOffset = builder.EndVector(dependenciesLength)
	}
	DependencyListStart(builder)
	DependencyListAddDependencies(builder, dependenciesOffset)
	return DependencyListEnd(builder)
}

func (rcv *DependencyList) UnPackTo(t *DependencyListT) {
	dependenciesLength := rcv.DependenciesLength()
	t.Dependencies = make([]*DependencyT, dependenciesLength)
	for j := 0; j < dependenciesLength; j++ {
		x := Dependency{}
		rcv.Dependencies(&x, j)
		t.Dependencies[j] = x.UnPack()
	}
}

func (rcv *DependencyList) UnPack() *DependencyListT {
	if rcv == nil { return nil }
	t := &DependencyListT{}
	rcv.UnPackTo(t)
	return t
}

type DependencyList struct {
	_tab flatbuffers.Table
}

func GetRootAsDependencyList(buf []byte, offset flatbuffers.UOffsetT) *DependencyList {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DependencyList{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDependencyList(buf []byte, offset flatbuffers.UOffsetT) *DependencyList {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DependencyList{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DependencyList) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DependencyList) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DependencyList) Dependencies(obj *Dependency, j int) bool {
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

func (rcv *DependencyList) DependenciesByKey(obj *Dependency, key string) bool{
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *DependencyList) DependenciesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DependencyListStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DependencyListAddDependencies(builder *flatbuffers.Builder, dependencies flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(dependencies), 0)
}
func DependencyListStartDependenciesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DependencyListEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
