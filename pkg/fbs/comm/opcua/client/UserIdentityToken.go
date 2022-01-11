// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package client

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UserIdentityToken struct {
	_tab flatbuffers.Table
}

func GetRootAsUserIdentityToken(buf []byte, offset flatbuffers.UOffsetT) *UserIdentityToken {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &UserIdentityToken{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUserIdentityToken(buf []byte, offset flatbuffers.UOffsetT) *UserIdentityToken {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &UserIdentityToken{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *UserIdentityToken) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UserIdentityToken) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *UserIdentityToken) UserIdentityTokenType() UserIdentityTokenUnion {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return UserIdentityTokenUnion(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *UserIdentityToken) MutateUserIdentityTokenType(n UserIdentityTokenUnion) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *UserIdentityToken) UserIdentityToken(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func UserIdentityTokenStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func UserIdentityTokenAddUserIdentityTokenType(builder *flatbuffers.Builder, userIdentityTokenType UserIdentityTokenUnion) {
	builder.PrependByteSlot(0, byte(userIdentityTokenType), 0)
}
func UserIdentityTokenAddUserIdentityToken(builder *flatbuffers.Builder, userIdentityToken flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(userIdentityToken), 0)
}
func UserIdentityTokenEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}