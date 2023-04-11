// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SlaveStatusResponseT struct {
	Status uint32
}

func (t *SlaveStatusResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	SlaveStatusResponseStart(builder)
	SlaveStatusResponseAddStatus(builder, t.Status)
	return SlaveStatusResponseEnd(builder)
}

func (rcv *SlaveStatusResponse) UnPackTo(t *SlaveStatusResponseT) {
	t.Status = rcv.Status()
}

func (rcv *SlaveStatusResponse) UnPack() *SlaveStatusResponseT {
	if rcv == nil { return nil }
	t := &SlaveStatusResponseT{}
	rcv.UnPackTo(t)
	return t
}

type SlaveStatusResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsSlaveStatusResponse(buf []byte, offset flatbuffers.UOffsetT) *SlaveStatusResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SlaveStatusResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSlaveStatusResponse(buf []byte, offset flatbuffers.UOffsetT) *SlaveStatusResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SlaveStatusResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SlaveStatusResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SlaveStatusResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SlaveStatusResponse) Status() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveStatusResponse) MutateStatus(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func SlaveStatusResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SlaveStatusResponseAddStatus(builder *flatbuffers.Builder, status uint32) {
	builder.PrependUint32Slot(0, status, 0)
}
func SlaveStatusResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
