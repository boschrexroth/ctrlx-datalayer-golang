// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Entry description response
type EntryDescriptionResponseT struct {
	Data []byte `json:"data"`
}

func (t *EntryDescriptionResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	EntryDescriptionResponseStart(builder)
	EntryDescriptionResponseAddData(builder, dataOffset)
	return EntryDescriptionResponseEnd(builder)
}

func (rcv *EntryDescriptionResponse) UnPackTo(t *EntryDescriptionResponseT) {
	t.Data = rcv.DataBytes()
}

func (rcv *EntryDescriptionResponse) UnPack() *EntryDescriptionResponseT {
	if rcv == nil { return nil }
	t := &EntryDescriptionResponseT{}
	rcv.UnPackTo(t)
	return t
}

type EntryDescriptionResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsEntryDescriptionResponse(buf []byte, offset flatbuffers.UOffsetT) *EntryDescriptionResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EntryDescriptionResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEntryDescriptionResponse(buf []byte, offset flatbuffers.UOffsetT) *EntryDescriptionResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &EntryDescriptionResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *EntryDescriptionResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EntryDescriptionResponse) Table() flatbuffers.Table {
	return rcv._tab
}

///Response data in binary format
func (rcv *EntryDescriptionResponse) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *EntryDescriptionResponse) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *EntryDescriptionResponse) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

///Response data in binary format
func (rcv *EntryDescriptionResponse) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func EntryDescriptionResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func EntryDescriptionResponseAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(data), 0)
}
func EntryDescriptionResponseStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func EntryDescriptionResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
