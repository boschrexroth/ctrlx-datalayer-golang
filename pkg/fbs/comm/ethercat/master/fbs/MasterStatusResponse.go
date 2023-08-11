// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MasterStatusResponseT struct {
	Status uint32 `json:"status"`
}

func (t *MasterStatusResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	MasterStatusResponseStart(builder)
	MasterStatusResponseAddStatus(builder, t.Status)
	return MasterStatusResponseEnd(builder)
}

func (rcv *MasterStatusResponse) UnPackTo(t *MasterStatusResponseT) {
	t.Status = rcv.Status()
}

func (rcv *MasterStatusResponse) UnPack() *MasterStatusResponseT {
	if rcv == nil { return nil }
	t := &MasterStatusResponseT{}
	rcv.UnPackTo(t)
	return t
}

type MasterStatusResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsMasterStatusResponse(buf []byte, offset flatbuffers.UOffsetT) *MasterStatusResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MasterStatusResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMasterStatusResponse(buf []byte, offset flatbuffers.UOffsetT) *MasterStatusResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MasterStatusResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MasterStatusResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MasterStatusResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MasterStatusResponse) Status() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MasterStatusResponse) MutateStatus(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func MasterStatusResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MasterStatusResponseAddStatus(builder *flatbuffers.Builder, status uint32) {
	builder.PrependUint32Slot(0, status, 0)
}
func MasterStatusResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
