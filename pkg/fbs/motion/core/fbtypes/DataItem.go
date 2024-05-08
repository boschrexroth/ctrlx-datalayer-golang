// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// information on a single data item
type DataItemT struct {
	Uri string `json:"uri"`
	Value float64 `json:"value"`
	Unit string `json:"unit"`
	Status string `json:"status"`
}

func (t *DataItemT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	uriOffset := flatbuffers.UOffsetT(0)
	if t.Uri != "" {
		uriOffset = builder.CreateString(t.Uri)
	}
	unitOffset := flatbuffers.UOffsetT(0)
	if t.Unit != "" {
		unitOffset = builder.CreateString(t.Unit)
	}
	statusOffset := flatbuffers.UOffsetT(0)
	if t.Status != "" {
		statusOffset = builder.CreateString(t.Status)
	}
	DataItemStart(builder)
	DataItemAddUri(builder, uriOffset)
	DataItemAddValue(builder, t.Value)
	DataItemAddUnit(builder, unitOffset)
	DataItemAddStatus(builder, statusOffset)
	return DataItemEnd(builder)
}

func (rcv *DataItem) UnPackTo(t *DataItemT) {
	t.Uri = string(rcv.Uri())
	t.Value = rcv.Value()
	t.Unit = string(rcv.Unit())
	t.Status = string(rcv.Status())
}

func (rcv *DataItem) UnPack() *DataItemT {
	if rcv == nil { return nil }
	t := &DataItemT{}
	rcv.UnPackTo(t)
	return t
}

type DataItem struct {
	_tab flatbuffers.Table
}

func GetRootAsDataItem(buf []byte, offset flatbuffers.UOffsetT) *DataItem {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DataItem{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDataItem(buf []byte, offset flatbuffers.UOffsetT) *DataItem {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DataItem{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DataItem) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DataItem) Table() flatbuffers.Table {
	return rcv._tab
}

/// URI of the data item
func (rcv *DataItem) Uri() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// URI of the data item
/// value of the data item at the time of the request
func (rcv *DataItem) Value() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// value of the data item at the time of the request
func (rcv *DataItem) MutateValue(n float64) bool {
	return rcv._tab.MutateFloat64Slot(6, n)
}

/// unit of the value (or empty for unit-less values)
func (rcv *DataItem) Unit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// unit of the value (or empty for unit-less values)
/// result of the read-value-call (value is only set on STS_OK)
func (rcv *DataItem) Status() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// result of the read-value-call (value is only set on STS_OK)
func DataItemStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DataItemAddUri(builder *flatbuffers.Builder, uri flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(uri), 0)
}
func DataItemAddValue(builder *flatbuffers.Builder, value float64) {
	builder.PrependFloat64Slot(1, value, 0.0)
}
func DataItemAddUnit(builder *flatbuffers.Builder, unit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(unit), 0)
}
func DataItemAddStatus(builder *flatbuffers.Builder, status flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(status), 0)
}
func DataItemEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}