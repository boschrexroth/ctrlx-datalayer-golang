// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SlaveStatisticCountersResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsSlaveStatisticCountersResponse(buf []byte, offset flatbuffers.UOffsetT) *SlaveStatisticCountersResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SlaveStatisticCountersResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSlaveStatisticCountersResponse(buf []byte, offset flatbuffers.UOffsetT) *SlaveStatisticCountersResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SlaveStatisticCountersResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SlaveStatisticCountersResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SlaveStatisticCountersResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SlaveStatisticCountersResponse) AlStatusCode() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveStatisticCountersResponse) MutateAlStatusCode(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func (rcv *SlaveStatisticCountersResponse) ProcUnitErrorCounter() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveStatisticCountersResponse) MutateProcUnitErrorCounter(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *SlaveStatisticCountersResponse) PdiErrorCounter() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveStatisticCountersResponse) MutatePdiErrorCounter(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *SlaveStatisticCountersResponse) PortErrorCounters(obj *PortErrorCounters, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *SlaveStatisticCountersResponse) PortErrorCountersLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func SlaveStatisticCountersResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func SlaveStatisticCountersResponseAddAlStatusCode(builder *flatbuffers.Builder, alStatusCode uint16) {
	builder.PrependUint16Slot(0, alStatusCode, 0)
}
func SlaveStatisticCountersResponseAddProcUnitErrorCounter(builder *flatbuffers.Builder, procUnitErrorCounter byte) {
	builder.PrependByteSlot(1, procUnitErrorCounter, 0)
}
func SlaveStatisticCountersResponseAddPdiErrorCounter(builder *flatbuffers.Builder, pdiErrorCounter byte) {
	builder.PrependByteSlot(2, pdiErrorCounter, 0)
}
func SlaveStatisticCountersResponseAddPortErrorCounters(builder *flatbuffers.Builder, portErrorCounters flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(portErrorCounters), 0)
}
func SlaveStatisticCountersResponseStartPortErrorCountersVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 1)
}
func SlaveStatisticCountersResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}