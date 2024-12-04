// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package daq

import (
	"bytes"
	flatbuffers "github.com/google/flatbuffers/go"
)

/// equidistant DAQ data
type EquidistantT struct {
	Name string `json:"name"`
	Tags []*TagT `json:"tags"`
	TimestampNs uint64 `json:"timestampNs"`
	IntervalNs uint64 `json:"intervalNs"`
	Array *ArrayT `json:"array"`
}

func (t *EquidistantT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	tagsOffset := flatbuffers.UOffsetT(0)
	if t.Tags != nil {
		tagsLength := len(t.Tags)
		tagsOffsets := make([]flatbuffers.UOffsetT, tagsLength)
		for j := 0; j < tagsLength; j++ {
			tagsOffsets[j] = t.Tags[j].Pack(builder)
		}
		EquidistantStartTagsVector(builder, tagsLength)
		for j := tagsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(tagsOffsets[j])
		}
		tagsOffset = builder.EndVector(tagsLength)
	}
	arrayOffset := t.Array.Pack(builder)
	
	EquidistantStart(builder)
	EquidistantAddName(builder, nameOffset)
	EquidistantAddTags(builder, tagsOffset)
	EquidistantAddTimestampNs(builder, t.TimestampNs)
	EquidistantAddIntervalNs(builder, t.IntervalNs)
	if t.Array != nil {
		EquidistantAddArrayType(builder, t.Array.Type)
	}
	EquidistantAddArray(builder, arrayOffset)
	return EquidistantEnd(builder)
}

func (rcv *Equidistant) UnPackTo(t *EquidistantT) {
	t.Name = string(rcv.Name())
	tagsLength := rcv.TagsLength()
	t.Tags = make([]*TagT, tagsLength)
	for j := 0; j < tagsLength; j++ {
		x := Tag{}
		rcv.Tags(&x, j)
		t.Tags[j] = x.UnPack()
	}
	t.TimestampNs = rcv.TimestampNs()
	t.IntervalNs = rcv.IntervalNs()
	arrayTable := flatbuffers.Table{}
	if rcv.Array(&arrayTable) {
		t.Array = rcv.ArrayType().UnPack(arrayTable)
	}
}

func (rcv *Equidistant) UnPack() *EquidistantT {
	if rcv == nil { return nil }
	t := &EquidistantT{}
	rcv.UnPackTo(t)
	return t
}

type Equidistant struct {
	_tab flatbuffers.Table
}

func GetRootAsEquidistant(buf []byte, offset flatbuffers.UOffsetT) *Equidistant {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Equidistant{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEquidistant(buf []byte, offset flatbuffers.UOffsetT) *Equidistant {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Equidistant{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Equidistant) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Equidistant) Table() flatbuffers.Table {
	return rcv._tab
}

/// unique name of the equidistant data
func (rcv *Equidistant) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unique name of the equidistant data
func EquidistantKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &Equidistant{}
	obj2 := &Equidistant{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return string(obj1.Name()) < string(obj2.Name())
}

func (rcv *Equidistant) LookupByKey(key string, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	bKey := []byte(key)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &Equidistant{}
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

/// array of tags for meta data
func (rcv *Equidistant) Tags(obj *Tag, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Equidistant) TagsByKey(obj *Tag, key string) bool{
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *Equidistant) TagsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// array of tags for meta data
/// unix epoch time (since 1.1.1970) for the 1st value in array in [ns]
func (rcv *Equidistant) TimestampNs() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// unix epoch time (since 1.1.1970) for the 1st value in array in [ns]
func (rcv *Equidistant) MutateTimestampNs(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

/// sampling interval in [ns]
func (rcv *Equidistant) IntervalNs() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// sampling interval in [ns]
func (rcv *Equidistant) MutateIntervalNs(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func (rcv *Equidistant) ArrayType() Array {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return Array(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Equidistant) MutateArrayType(n Array) bool {
	return rcv._tab.MutateByteSlot(12, byte(n))
}

/// array of equidistant sampled values starting at timestamp [ns] with interval [ns]
func (rcv *Equidistant) Array(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

/// array of equidistant sampled values starting at timestamp [ns] with interval [ns]
func EquidistantStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func EquidistantAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func EquidistantAddTags(builder *flatbuffers.Builder, tags flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(tags), 0)
}
func EquidistantStartTagsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EquidistantAddTimestampNs(builder *flatbuffers.Builder, timestampNs uint64) {
	builder.PrependUint64Slot(2, timestampNs, 0)
}
func EquidistantAddIntervalNs(builder *flatbuffers.Builder, intervalNs uint64) {
	builder.PrependUint64Slot(3, intervalNs, 0)
}
func EquidistantAddArrayType(builder *flatbuffers.Builder, arrayType Array) {
	builder.PrependByteSlot(4, byte(arrayType), 0)
}
func EquidistantAddArray(builder *flatbuffers.Builder, array flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(array), 0)
}
func EquidistantEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}