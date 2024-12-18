// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package daq

import (
	"bytes"
	flatbuffers "github.com/google/flatbuffers/go"
)

/// tag as key-value-pair for meta data
type TagT struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

func (t *TagT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	keyOffset := flatbuffers.UOffsetT(0)
	if t.Key != "" {
		keyOffset = builder.CreateString(t.Key)
	}
	valueOffset := flatbuffers.UOffsetT(0)
	if t.Value != "" {
		valueOffset = builder.CreateString(t.Value)
	}
	TagStart(builder)
	TagAddKey(builder, keyOffset)
	TagAddValue(builder, valueOffset)
	return TagEnd(builder)
}

func (rcv *Tag) UnPackTo(t *TagT) {
	t.Key = string(rcv.Key())
	t.Value = string(rcv.Value())
}

func (rcv *Tag) UnPack() *TagT {
	if rcv == nil { return nil }
	t := &TagT{}
	rcv.UnPackTo(t)
	return t
}

type Tag struct {
	_tab flatbuffers.Table
}

func GetRootAsTag(buf []byte, offset flatbuffers.UOffsetT) *Tag {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Tag{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTag(buf []byte, offset flatbuffers.UOffsetT) *Tag {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Tag{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Tag) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Tag) Table() flatbuffers.Table {
	return rcv._tab
}

/// unique key of the tag
func (rcv *Tag) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unique key of the tag
func TagKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &Tag{}
	obj2 := &Tag{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return string(obj1.Key()) < string(obj2.Key())
}

func (rcv *Tag) LookupByKey(key string, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	bKey := []byte(key)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &Tag{}
		obj.Init(buf, tableOffset)
		comp := bytes.Compare(obj.Key(), bKey)
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

/// value of the tag
func (rcv *Tag) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// value of the tag
func TagStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func TagAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func TagAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func TagEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
