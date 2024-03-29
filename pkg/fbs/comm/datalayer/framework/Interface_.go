// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package framework

import (
	"bytes"
	flatbuffers "github.com/google/flatbuffers/go"
)

type Interface_T struct {
	Name string `json:"name"`
	Version string `json:"version"`
	Language string `json:"language"`
	Properties []*PropertyT `json:"properties"`
}

func (t *Interface_T) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	versionOffset := flatbuffers.UOffsetT(0)
	if t.Version != "" {
		versionOffset = builder.CreateString(t.Version)
	}
	languageOffset := flatbuffers.UOffsetT(0)
	if t.Language != "" {
		languageOffset = builder.CreateString(t.Language)
	}
	propertiesOffset := flatbuffers.UOffsetT(0)
	if t.Properties != nil {
		propertiesLength := len(t.Properties)
		propertiesOffsets := make([]flatbuffers.UOffsetT, propertiesLength)
		for j := 0; j < propertiesLength; j++ {
			propertiesOffsets[j] = t.Properties[j].Pack(builder)
		}
		Interface_StartPropertiesVector(builder, propertiesLength)
		for j := propertiesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(propertiesOffsets[j])
		}
		propertiesOffset = builder.EndVector(propertiesLength)
	}
	Interface_Start(builder)
	Interface_AddName(builder, nameOffset)
	Interface_AddVersion(builder, versionOffset)
	Interface_AddLanguage(builder, languageOffset)
	Interface_AddProperties(builder, propertiesOffset)
	return Interface_End(builder)
}

func (rcv *Interface_) UnPackTo(t *Interface_T) {
	t.Name = string(rcv.Name())
	t.Version = string(rcv.Version())
	t.Language = string(rcv.Language())
	propertiesLength := rcv.PropertiesLength()
	t.Properties = make([]*PropertyT, propertiesLength)
	for j := 0; j < propertiesLength; j++ {
		x := Property{}
		rcv.Properties(&x, j)
		t.Properties[j] = x.UnPack()
	}
}

func (rcv *Interface_) UnPack() *Interface_T {
	if rcv == nil { return nil }
	t := &Interface_T{}
	rcv.UnPackTo(t)
	return t
}

type Interface_ struct {
	_tab flatbuffers.Table
}

func GetRootAsInterface_(buf []byte, offset flatbuffers.UOffsetT) *Interface_ {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Interface_{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInterface_(buf []byte, offset flatbuffers.UOffsetT) *Interface_ {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Interface_{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Interface_) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Interface_) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Interface_) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func Interface_KeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &Interface_{}
	obj2 := &Interface_{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return string(obj1.Name()) < string(obj2.Name())
}

func (rcv *Interface_) LookupByKey(key string, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	bKey := []byte(key)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &Interface_{}
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

func (rcv *Interface_) Version() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Interface_) Language() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Interface_) Properties(obj *Property, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Interface_) PropertiesByKey(obj *Property, key string) bool{
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *Interface_) PropertiesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func Interface_Start(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func Interface_AddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func Interface_AddVersion(builder *flatbuffers.Builder, version flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(version), 0)
}
func Interface_AddLanguage(builder *flatbuffers.Builder, language flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(language), 0)
}
func Interface_AddProperties(builder *flatbuffers.Builder, properties flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(properties), 0)
}
func Interface_StartPropertiesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func Interface_End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
