// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package framework

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RegisterT struct {
	Ip string `json:"IP"`
	Sp string `json:"SP"`
	Bp string `json:"BP"`
}

func (t *RegisterT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ipOffset := flatbuffers.UOffsetT(0)
	if t.Ip != "" {
		ipOffset = builder.CreateString(t.Ip)
	}
	spOffset := flatbuffers.UOffsetT(0)
	if t.Sp != "" {
		spOffset = builder.CreateString(t.Sp)
	}
	bpOffset := flatbuffers.UOffsetT(0)
	if t.Bp != "" {
		bpOffset = builder.CreateString(t.Bp)
	}
	RegisterStart(builder)
	RegisterAddIp(builder, ipOffset)
	RegisterAddSp(builder, spOffset)
	RegisterAddBp(builder, bpOffset)
	return RegisterEnd(builder)
}

func (rcv *Register) UnPackTo(t *RegisterT) {
	t.Ip = string(rcv.Ip())
	t.Sp = string(rcv.Sp())
	t.Bp = string(rcv.Bp())
}

func (rcv *Register) UnPack() *RegisterT {
	if rcv == nil { return nil }
	t := &RegisterT{}
	rcv.UnPackTo(t)
	return t
}

type Register struct {
	_tab flatbuffers.Table
}

func GetRootAsRegister(buf []byte, offset flatbuffers.UOffsetT) *Register {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Register{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRegister(buf []byte, offset flatbuffers.UOffsetT) *Register {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Register{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Register) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Register) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Register) Ip() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Register) Sp() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Register) Bp() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func RegisterStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func RegisterAddIp(builder *flatbuffers.Builder, ip flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ip), 0)
}
func RegisterAddSp(builder *flatbuffers.Builder, sp flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(sp), 0)
}
func RegisterAddBp(builder *flatbuffers.Builder, bp flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(bp), 0)
}
func RegisterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
