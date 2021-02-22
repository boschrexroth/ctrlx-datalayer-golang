// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DataChangeFilter struct {
	_tab flatbuffers.Table
}

func GetRootAsDataChangeFilter(buf []byte, offset flatbuffers.UOffsetT) *DataChangeFilter {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DataChangeFilter{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDataChangeFilter(buf []byte, offset flatbuffers.UOffsetT) *DataChangeFilter {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DataChangeFilter{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DataChangeFilter) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DataChangeFilter) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DataChangeFilter) DeadBandValue() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *DataChangeFilter) MutateDeadBandValue(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func DataChangeFilterStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DataChangeFilterAddDeadBandValue(builder *flatbuffers.Builder, deadBandValue float32) {
	builder.PrependFloat32Slot(0, deadBandValue, 0.0)
}
func DataChangeFilterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
