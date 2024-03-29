// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	"bytes"
	flatbuffers "github.com/google/flatbuffers/go"
)

type mappingEntryT struct {
	ValueId string `json:"valueID"`
	DatalayerUri string `json:"datalayerURI"`
	ProfileVar variableType `json:"profileVar"`
	Required bool `json:"required"`
}

func (t *mappingEntryT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	valueIdOffset := flatbuffers.UOffsetT(0)
	if t.ValueId != "" {
		valueIdOffset = builder.CreateString(t.ValueId)
	}
	datalayerUriOffset := flatbuffers.UOffsetT(0)
	if t.DatalayerUri != "" {
		datalayerUriOffset = builder.CreateString(t.DatalayerUri)
	}
	mappingEntryStart(builder)
	mappingEntryAddValueId(builder, valueIdOffset)
	mappingEntryAddDatalayerUri(builder, datalayerUriOffset)
	mappingEntryAddProfileVar(builder, t.ProfileVar)
	mappingEntryAddRequired(builder, t.Required)
	return mappingEntryEnd(builder)
}

func (rcv *mappingEntry) UnPackTo(t *mappingEntryT) {
	t.ValueId = string(rcv.ValueId())
	t.DatalayerUri = string(rcv.DatalayerUri())
	t.ProfileVar = rcv.ProfileVar()
	t.Required = rcv.Required()
}

func (rcv *mappingEntry) UnPack() *mappingEntryT {
	if rcv == nil { return nil }
	t := &mappingEntryT{}
	rcv.UnPackTo(t)
	return t
}

type mappingEntry struct {
	_tab flatbuffers.Table
}

func GetRootAsmappingEntry(buf []byte, offset flatbuffers.UOffsetT) *mappingEntry {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &mappingEntry{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsmappingEntry(buf []byte, offset flatbuffers.UOffsetT) *mappingEntry {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &mappingEntry{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *mappingEntry) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *mappingEntry) Table() flatbuffers.Table {
	return rcv._tab
}

/// Id-string of the mapping entry
func (rcv *mappingEntry) ValueId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Id-string of the mapping entry
func mappingEntryKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &mappingEntry{}
	obj2 := &mappingEntry{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return string(obj1.ValueId()) < string(obj2.ValueId())
}

func (rcv *mappingEntry) LookupByKey(key string, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	bKey := []byte(key)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &mappingEntry{}
		obj.Init(buf, tableOffset)
		comp := bytes.Compare(obj.ValueId(), bKey)
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

/// datalayer uri of the mapping entry
func (rcv *mappingEntry) DatalayerUri() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// datalayer uri of the mapping entry
/// variable type of mapped variable
func (rcv *mappingEntry) ProfileVar() variableType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return variableType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

/// variable type of mapped variable
func (rcv *mappingEntry) MutateProfileVar(n variableType) bool {
	return rcv._tab.MutateInt8Slot(8, int8(n))
}

/// is current item is required
func (rcv *mappingEntry) Required() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// is current item is required
func (rcv *mappingEntry) MutateRequired(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func mappingEntryStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func mappingEntryAddValueId(builder *flatbuffers.Builder, valueId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(valueId), 0)
}
func mappingEntryAddDatalayerUri(builder *flatbuffers.Builder, datalayerUri flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(datalayerUri), 0)
}
func mappingEntryAddProfileVar(builder *flatbuffers.Builder, profileVar variableType) {
	builder.PrependInt8Slot(2, int8(profileVar), 0)
}
func mappingEntryAddRequired(builder *flatbuffers.Builder, required bool) {
	builder.PrependBoolSlot(3, required, false)
}
func mappingEntryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
