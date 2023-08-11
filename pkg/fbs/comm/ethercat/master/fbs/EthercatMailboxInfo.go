// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type EthercatMailboxInfoT struct {
	SizeIn uint32 `json:"sizeIn"`
	SizeOut uint32 `json:"sizeOut"`
}

func (t *EthercatMailboxInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	return CreateEthercatMailboxInfo(builder, t.SizeIn, t.SizeOut)
}
func (rcv *EthercatMailboxInfo) UnPackTo(t *EthercatMailboxInfoT) {
	t.SizeIn = rcv.SizeIn()
	t.SizeOut = rcv.SizeOut()
}

func (rcv *EthercatMailboxInfo) UnPack() *EthercatMailboxInfoT {
	if rcv == nil { return nil }
	t := &EthercatMailboxInfoT{}
	rcv.UnPackTo(t)
	return t
}

type EthercatMailboxInfo struct {
	_tab flatbuffers.Struct
}

func (rcv *EthercatMailboxInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EthercatMailboxInfo) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *EthercatMailboxInfo) SizeIn() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *EthercatMailboxInfo) MutateSizeIn(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *EthercatMailboxInfo) SizeOut() uint32 {
	return rcv._tab.GetUint32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *EthercatMailboxInfo) MutateSizeOut(n uint32) bool {
	return rcv._tab.MutateUint32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func CreateEthercatMailboxInfo(builder *flatbuffers.Builder, sizeIn uint32, sizeOut uint32) flatbuffers.UOffsetT {
	builder.Prep(4, 8)
	builder.PrependUint32(sizeOut)
	builder.PrependUint32(sizeIn)
	return builder.Offset()
}
