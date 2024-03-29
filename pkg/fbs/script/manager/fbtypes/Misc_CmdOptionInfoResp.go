// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Misc_CmdOptionInfoRespT struct {
	Bfbs string `json:"bfbs"`
	Description string `json:"description"`
}

func (t *Misc_CmdOptionInfoRespT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	bfbsOffset := flatbuffers.UOffsetT(0)
	if t.Bfbs != "" {
		bfbsOffset = builder.CreateString(t.Bfbs)
	}
	descriptionOffset := flatbuffers.UOffsetT(0)
	if t.Description != "" {
		descriptionOffset = builder.CreateString(t.Description)
	}
	Misc_CmdOptionInfoRespStart(builder)
	Misc_CmdOptionInfoRespAddBfbs(builder, bfbsOffset)
	Misc_CmdOptionInfoRespAddDescription(builder, descriptionOffset)
	return Misc_CmdOptionInfoRespEnd(builder)
}

func (rcv *Misc_CmdOptionInfoResp) UnPackTo(t *Misc_CmdOptionInfoRespT) {
	t.Bfbs = string(rcv.Bfbs())
	t.Description = string(rcv.Description())
}

func (rcv *Misc_CmdOptionInfoResp) UnPack() *Misc_CmdOptionInfoRespT {
	if rcv == nil { return nil }
	t := &Misc_CmdOptionInfoRespT{}
	rcv.UnPackTo(t)
	return t
}

type Misc_CmdOptionInfoResp struct {
	_tab flatbuffers.Table
}

func GetRootAsMisc_CmdOptionInfoResp(buf []byte, offset flatbuffers.UOffsetT) *Misc_CmdOptionInfoResp {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Misc_CmdOptionInfoResp{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMisc_CmdOptionInfoResp(buf []byte, offset flatbuffers.UOffsetT) *Misc_CmdOptionInfoResp {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Misc_CmdOptionInfoResp{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Misc_CmdOptionInfoResp) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Misc_CmdOptionInfoResp) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Misc_CmdOptionInfoResp) Bfbs() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Misc_CmdOptionInfoResp) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func Misc_CmdOptionInfoRespStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Misc_CmdOptionInfoRespAddBfbs(builder *flatbuffers.Builder, bfbs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(bfbs), 0)
}
func Misc_CmdOptionInfoRespAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(description), 0)
}
func Misc_CmdOptionInfoRespEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
