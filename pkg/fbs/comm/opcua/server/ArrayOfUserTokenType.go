// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package server

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ArrayOfUserTokenTypeT struct {
	UserTokenArray []*UserTokenTypeT
}

func (t *ArrayOfUserTokenTypeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	userTokenArrayOffset := flatbuffers.UOffsetT(0)
	if t.UserTokenArray != nil {
		userTokenArrayLength := len(t.UserTokenArray)
		userTokenArrayOffsets := make([]flatbuffers.UOffsetT, userTokenArrayLength)
		for j := 0; j < userTokenArrayLength; j++ {
			userTokenArrayOffsets[j] = t.UserTokenArray[j].Pack(builder)
		}
		ArrayOfUserTokenTypeStartUserTokenArrayVector(builder, userTokenArrayLength)
		for j := userTokenArrayLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(userTokenArrayOffsets[j])
		}
		userTokenArrayOffset = builder.EndVector(userTokenArrayLength)
	}
	ArrayOfUserTokenTypeStart(builder)
	ArrayOfUserTokenTypeAddUserTokenArray(builder, userTokenArrayOffset)
	return ArrayOfUserTokenTypeEnd(builder)
}

func (rcv *ArrayOfUserTokenType) UnPackTo(t *ArrayOfUserTokenTypeT) {
	userTokenArrayLength := rcv.UserTokenArrayLength()
	t.UserTokenArray = make([]*UserTokenTypeT, userTokenArrayLength)
	for j := 0; j < userTokenArrayLength; j++ {
		x := UserTokenType{}
		rcv.UserTokenArray(&x, j)
		t.UserTokenArray[j] = x.UnPack()
	}
}

func (rcv *ArrayOfUserTokenType) UnPack() *ArrayOfUserTokenTypeT {
	if rcv == nil { return nil }
	t := &ArrayOfUserTokenTypeT{}
	rcv.UnPackTo(t)
	return t
}

type ArrayOfUserTokenType struct {
	_tab flatbuffers.Table
}

func GetRootAsArrayOfUserTokenType(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfUserTokenType {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ArrayOfUserTokenType{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsArrayOfUserTokenType(buf []byte, offset flatbuffers.UOffsetT) *ArrayOfUserTokenType {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ArrayOfUserTokenType{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ArrayOfUserTokenType) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ArrayOfUserTokenType) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ArrayOfUserTokenType) UserTokenArray(obj *UserTokenType, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ArrayOfUserTokenType) UserTokenArrayLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ArrayOfUserTokenTypeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ArrayOfUserTokenTypeAddUserTokenArray(builder *flatbuffers.Builder, userTokenArray flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(userTokenArray), 0)
}
func ArrayOfUserTokenTypeStartUserTokenArrayVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ArrayOfUserTokenTypeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
