// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type WatchlistT struct {
	Name string `json:"name"`
	Items []string `json:"items"`
}

func (t *WatchlistT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	itemsOffset := flatbuffers.UOffsetT(0)
	if t.Items != nil {
		itemsLength := len(t.Items)
		itemsOffsets := make([]flatbuffers.UOffsetT, itemsLength)
		for j := 0; j < itemsLength; j++ {
			itemsOffsets[j] = builder.CreateString(t.Items[j])
		}
		WatchlistStartItemsVector(builder, itemsLength)
		for j := itemsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(itemsOffsets[j])
		}
		itemsOffset = builder.EndVector(itemsLength)
	}
	WatchlistStart(builder)
	WatchlistAddName(builder, nameOffset)
	WatchlistAddItems(builder, itemsOffset)
	return WatchlistEnd(builder)
}

func (rcv *Watchlist) UnPackTo(t *WatchlistT) {
	t.Name = string(rcv.Name())
	itemsLength := rcv.ItemsLength()
	t.Items = make([]string, itemsLength)
	for j := 0; j < itemsLength; j++ {
		t.Items[j] = string(rcv.Items(j))
	}
}

func (rcv *Watchlist) UnPack() *WatchlistT {
	if rcv == nil { return nil }
	t := &WatchlistT{}
	rcv.UnPackTo(t)
	return t
}

type Watchlist struct {
	_tab flatbuffers.Table
}

func GetRootAsWatchlist(buf []byte, offset flatbuffers.UOffsetT) *Watchlist {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Watchlist{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsWatchlist(buf []byte, offset flatbuffers.UOffsetT) *Watchlist {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Watchlist{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Watchlist) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Watchlist) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the watchlist
func (rcv *Watchlist) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the watchlist
/// items of the watchlist
func (rcv *Watchlist) Items(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Watchlist) ItemsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// items of the watchlist
func WatchlistStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func WatchlistAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func WatchlistAddItems(builder *flatbuffers.Builder, items flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(items), 0)
}
func WatchlistStartItemsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func WatchlistEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
