// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// information on all available data items
type StateDataItemsT struct {
	Items []*DataItemT `json:"items"`
}

func (t *StateDataItemsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	itemsOffset := flatbuffers.UOffsetT(0)
	if t.Items != nil {
		itemsLength := len(t.Items)
		itemsOffsets := make([]flatbuffers.UOffsetT, itemsLength)
		for j := 0; j < itemsLength; j++ {
			itemsOffsets[j] = t.Items[j].Pack(builder)
		}
		StateDataItemsStartItemsVector(builder, itemsLength)
		for j := itemsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(itemsOffsets[j])
		}
		itemsOffset = builder.EndVector(itemsLength)
	}
	StateDataItemsStart(builder)
	StateDataItemsAddItems(builder, itemsOffset)
	return StateDataItemsEnd(builder)
}

func (rcv *StateDataItems) UnPackTo(t *StateDataItemsT) {
	itemsLength := rcv.ItemsLength()
	t.Items = make([]*DataItemT, itemsLength)
	for j := 0; j < itemsLength; j++ {
		x := DataItem{}
		rcv.Items(&x, j)
		t.Items[j] = x.UnPack()
	}
}

func (rcv *StateDataItems) UnPack() *StateDataItemsT {
	if rcv == nil { return nil }
	t := &StateDataItemsT{}
	rcv.UnPackTo(t)
	return t
}

type StateDataItems struct {
	_tab flatbuffers.Table
}

func GetRootAsStateDataItems(buf []byte, offset flatbuffers.UOffsetT) *StateDataItems {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StateDataItems{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStateDataItems(buf []byte, offset flatbuffers.UOffsetT) *StateDataItems {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StateDataItems{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StateDataItems) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StateDataItems) Table() flatbuffers.Table {
	return rcv._tab
}

/// all available data items (URI and current value)
func (rcv *StateDataItems) Items(obj *DataItem, j int) bool {
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

func (rcv *StateDataItems) ItemsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// all available data items (URI and current value)
func StateDataItemsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func StateDataItemsAddItems(builder *flatbuffers.Builder, items flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(items), 0)
}
func StateDataItemsStartItemsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func StateDataItemsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
