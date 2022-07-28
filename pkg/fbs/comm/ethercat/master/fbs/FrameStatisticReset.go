// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FrameStatisticResetT struct {
	Request *FrameStatisticResetRequestT
}

func (t *FrameStatisticResetT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	requestOffset := t.Request.Pack(builder)
	FrameStatisticResetStart(builder)
	FrameStatisticResetAddRequest(builder, requestOffset)
	return FrameStatisticResetEnd(builder)
}

func (rcv *FrameStatisticReset) UnPackTo(t *FrameStatisticResetT) {
	t.Request = rcv.Request(nil).UnPack()
}

func (rcv *FrameStatisticReset) UnPack() *FrameStatisticResetT {
	if rcv == nil { return nil }
	t := &FrameStatisticResetT{}
	rcv.UnPackTo(t)
	return t
}

type FrameStatisticReset struct {
	_tab flatbuffers.Table
}

func GetRootAsFrameStatisticReset(buf []byte, offset flatbuffers.UOffsetT) *FrameStatisticReset {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FrameStatisticReset{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFrameStatisticReset(buf []byte, offset flatbuffers.UOffsetT) *FrameStatisticReset {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FrameStatisticReset{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FrameStatisticReset) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FrameStatisticReset) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FrameStatisticReset) Request(obj *FrameStatisticResetRequest) *FrameStatisticResetRequest {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(FrameStatisticResetRequest)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func FrameStatisticResetStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func FrameStatisticResetAddRequest(builder *flatbuffers.Builder, request flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(request), 0)
}
func FrameStatisticResetEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}