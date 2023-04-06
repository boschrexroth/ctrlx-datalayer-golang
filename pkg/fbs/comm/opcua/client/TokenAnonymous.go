// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package client

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TokenAnonymousT struct {
}

func (t *TokenAnonymousT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	TokenAnonymousStart(builder)
	return TokenAnonymousEnd(builder)
}

func (rcv *TokenAnonymous) UnPackTo(t *TokenAnonymousT) {
}

func (rcv *TokenAnonymous) UnPack() *TokenAnonymousT {
	if rcv == nil { return nil }
	t := &TokenAnonymousT{}
	rcv.UnPackTo(t)
	return t
}

type TokenAnonymous struct {
	_tab flatbuffers.Table
}

func GetRootAsTokenAnonymous(buf []byte, offset flatbuffers.UOffsetT) *TokenAnonymous {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TokenAnonymous{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTokenAnonymous(buf []byte, offset flatbuffers.UOffsetT) *TokenAnonymous {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TokenAnonymous{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TokenAnonymous) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TokenAnonymous) Table() flatbuffers.Table {
	return rcv._tab
}

func TokenAnonymousStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func TokenAnonymousEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
