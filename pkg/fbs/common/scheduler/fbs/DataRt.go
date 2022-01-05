// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DataRt struct {
	_tab flatbuffers.Table
}

func GetRootAsDataRt(buf []byte, offset flatbuffers.UOffsetT) *DataRt {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DataRt{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDataRt(buf []byte, offset flatbuffers.UOffsetT) *DataRt {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DataRt{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DataRt) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DataRt) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DataRt) StartTime() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DataRt) MutateStartTime(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *DataRt) Counter() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DataRt) MutateCounter(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func DataRtStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func DataRtAddStartTime(builder *flatbuffers.Builder, startTime uint64) {
	builder.PrependUint64Slot(0, startTime, 0)
}
func DataRtAddCounter(builder *flatbuffers.Builder, counter uint64) {
	builder.PrependUint64Slot(1, counter, 0)
}
func DataRtEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
