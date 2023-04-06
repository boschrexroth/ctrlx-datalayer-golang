// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ChangeEventsT struct {
	ValueChange DataChangeTrigger
	BrowselistChange bool
	MetadataChange bool
}

func (t *ChangeEventsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ChangeEventsStart(builder)
	ChangeEventsAddValueChange(builder, t.ValueChange)
	ChangeEventsAddBrowselistChange(builder, t.BrowselistChange)
	ChangeEventsAddMetadataChange(builder, t.MetadataChange)
	return ChangeEventsEnd(builder)
}

func (rcv *ChangeEvents) UnPackTo(t *ChangeEventsT) {
	t.ValueChange = rcv.ValueChange()
	t.BrowselistChange = rcv.BrowselistChange()
	t.MetadataChange = rcv.MetadataChange()
}

func (rcv *ChangeEvents) UnPack() *ChangeEventsT {
	if rcv == nil { return nil }
	t := &ChangeEventsT{}
	rcv.UnPackTo(t)
	return t
}

type ChangeEvents struct {
	_tab flatbuffers.Table
}

func GetRootAsChangeEvents(buf []byte, offset flatbuffers.UOffsetT) *ChangeEvents {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ChangeEvents{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsChangeEvents(buf []byte, offset flatbuffers.UOffsetT) *ChangeEvents {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ChangeEvents{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ChangeEvents) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ChangeEvents) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ChangeEvents) ValueChange() DataChangeTrigger {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return DataChangeTrigger(rcv._tab.GetInt32(o + rcv._tab.Pos))
	}
	return 1
}

func (rcv *ChangeEvents) MutateValueChange(n DataChangeTrigger) bool {
	return rcv._tab.MutateInt32Slot(4, int32(n))
}

func (rcv *ChangeEvents) BrowselistChange() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ChangeEvents) MutateBrowselistChange(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *ChangeEvents) MetadataChange() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ChangeEvents) MutateMetadataChange(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func ChangeEventsStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ChangeEventsAddValueChange(builder *flatbuffers.Builder, valueChange DataChangeTrigger) {
	builder.PrependInt32Slot(0, int32(valueChange), 1)
}
func ChangeEventsAddBrowselistChange(builder *flatbuffers.Builder, browselistChange bool) {
	builder.PrependBoolSlot(1, browselistChange, false)
}
func ChangeEventsAddMetadataChange(builder *flatbuffers.Builder, metadataChange bool) {
	builder.PrependBoolSlot(2, metadataChange, false)
}
func ChangeEventsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
