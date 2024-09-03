// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StationNameResponseT struct {
	Name string `json:"name"`
}

func (t *StationNameResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	StationNameResponseStart(builder)
	StationNameResponseAddName(builder, nameOffset)
	return StationNameResponseEnd(builder)
}

func (rcv *StationNameResponse) UnPackTo(t *StationNameResponseT) {
	t.Name = string(rcv.Name())
}

func (rcv *StationNameResponse) UnPack() *StationNameResponseT {
	if rcv == nil { return nil }
	t := &StationNameResponseT{}
	rcv.UnPackTo(t)
	return t
}

type StationNameResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsStationNameResponse(buf []byte, offset flatbuffers.UOffsetT) *StationNameResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StationNameResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStationNameResponse(buf []byte, offset flatbuffers.UOffsetT) *StationNameResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StationNameResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StationNameResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StationNameResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StationNameResponse) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func StationNameResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func StationNameResponseAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func StationNameResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
