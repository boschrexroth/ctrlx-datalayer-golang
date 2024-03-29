// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package framework

import (
	"bytes"
	flatbuffers "github.com/google/flatbuffers/go"
)

type DependencyT struct {
	Name string `json:"name"`
	Available bool `json:"available"`
	Required bool `json:"required"`
	Filter string `json:"filter"`
}

func (t *DependencyT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	filterOffset := flatbuffers.UOffsetT(0)
	if t.Filter != "" {
		filterOffset = builder.CreateString(t.Filter)
	}
	DependencyStart(builder)
	DependencyAddName(builder, nameOffset)
	DependencyAddAvailable(builder, t.Available)
	DependencyAddRequired(builder, t.Required)
	DependencyAddFilter(builder, filterOffset)
	return DependencyEnd(builder)
}

func (rcv *Dependency) UnPackTo(t *DependencyT) {
	t.Name = string(rcv.Name())
	t.Available = rcv.Available()
	t.Required = rcv.Required()
	t.Filter = string(rcv.Filter())
}

func (rcv *Dependency) UnPack() *DependencyT {
	if rcv == nil { return nil }
	t := &DependencyT{}
	rcv.UnPackTo(t)
	return t
}

type Dependency struct {
	_tab flatbuffers.Table
}

func GetRootAsDependency(buf []byte, offset flatbuffers.UOffsetT) *Dependency {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Dependency{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDependency(buf []byte, offset flatbuffers.UOffsetT) *Dependency {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Dependency{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Dependency) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Dependency) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Dependency) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DependencyKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &Dependency{}
	obj2 := &Dependency{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return string(obj1.Name()) < string(obj2.Name())
}

func (rcv *Dependency) LookupByKey(key string, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	bKey := []byte(key)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &Dependency{}
		obj.Init(buf, tableOffset)
		comp := bytes.Compare(obj.Name(), bKey)
		if comp > 0 {
			span = middle
		} else if comp < 0 {
			middle += 1
			start += middle
			span -= middle
		} else {
			rcv.Init(buf, tableOffset)
			return true
		}
	}
	return false
}

func (rcv *Dependency) Available() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Dependency) MutateAvailable(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *Dependency) Required() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Dependency) MutateRequired(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func (rcv *Dependency) Filter() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DependencyStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DependencyAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func DependencyAddAvailable(builder *flatbuffers.Builder, available bool) {
	builder.PrependBoolSlot(1, available, false)
}
func DependencyAddRequired(builder *flatbuffers.Builder, required bool) {
	builder.PrependBoolSlot(2, required, false)
}
func DependencyAddFilter(builder *flatbuffers.Builder, filter flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(filter), 0)
}
func DependencyEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
