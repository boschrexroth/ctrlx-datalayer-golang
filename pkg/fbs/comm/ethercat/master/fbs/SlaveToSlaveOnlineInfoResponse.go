// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Slave to slave online information response
type SlaveToSlaveOnlineInfoResponseT struct {
	ProcessDataCopy bool `json:"processDataCopy"`
	ProcessDataCopyMode SlaveToSlaveCopyMode `json:"processDataCopyMode"`
}

func (t *SlaveToSlaveOnlineInfoResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	SlaveToSlaveOnlineInfoResponseStart(builder)
	SlaveToSlaveOnlineInfoResponseAddProcessDataCopy(builder, t.ProcessDataCopy)
	SlaveToSlaveOnlineInfoResponseAddProcessDataCopyMode(builder, t.ProcessDataCopyMode)
	return SlaveToSlaveOnlineInfoResponseEnd(builder)
}

func (rcv *SlaveToSlaveOnlineInfoResponse) UnPackTo(t *SlaveToSlaveOnlineInfoResponseT) {
	t.ProcessDataCopy = rcv.ProcessDataCopy()
	t.ProcessDataCopyMode = rcv.ProcessDataCopyMode()
}

func (rcv *SlaveToSlaveOnlineInfoResponse) UnPack() *SlaveToSlaveOnlineInfoResponseT {
	if rcv == nil { return nil }
	t := &SlaveToSlaveOnlineInfoResponseT{}
	rcv.UnPackTo(t)
	return t
}

type SlaveToSlaveOnlineInfoResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsSlaveToSlaveOnlineInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *SlaveToSlaveOnlineInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SlaveToSlaveOnlineInfoResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSlaveToSlaveOnlineInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *SlaveToSlaveOnlineInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SlaveToSlaveOnlineInfoResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SlaveToSlaveOnlineInfoResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SlaveToSlaveOnlineInfoResponse) Table() flatbuffers.Table {
	return rcv._tab
}

///Slave to slave copy active
func (rcv *SlaveToSlaveOnlineInfoResponse) ProcessDataCopy() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

///Slave to slave copy active
func (rcv *SlaveToSlaveOnlineInfoResponse) MutateProcessDataCopy(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

///Slave to slave copy mode for process data communication
func (rcv *SlaveToSlaveOnlineInfoResponse) ProcessDataCopyMode() SlaveToSlaveCopyMode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return SlaveToSlaveCopyMode(rcv._tab.GetUint32(o + rcv._tab.Pos))
	}
	return 0
}

///Slave to slave copy mode for process data communication
func (rcv *SlaveToSlaveOnlineInfoResponse) MutateProcessDataCopyMode(n SlaveToSlaveCopyMode) bool {
	return rcv._tab.MutateUint32Slot(6, uint32(n))
}

func SlaveToSlaveOnlineInfoResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SlaveToSlaveOnlineInfoResponseAddProcessDataCopy(builder *flatbuffers.Builder, processDataCopy bool) {
	builder.PrependBoolSlot(0, processDataCopy, false)
}
func SlaveToSlaveOnlineInfoResponseAddProcessDataCopyMode(builder *flatbuffers.Builder, processDataCopyMode SlaveToSlaveCopyMode) {
	builder.PrependUint32Slot(1, uint32(processDataCopyMode), 0)
}
func SlaveToSlaveOnlineInfoResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
