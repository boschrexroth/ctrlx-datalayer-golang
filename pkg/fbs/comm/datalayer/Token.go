// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TokenT struct {
	Claims []*ClaimT `json:"claims"`
	Id string `json:"id"`
	Iat uint64 `json:"iat"`
	Exp uint64 `json:"exp"`
	Name string `json:"name"`
	Plchandle uint64 `json:"plchandle"`
	Scope []string `json:"scope"`
}

func (t *TokenT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	claimsOffset := flatbuffers.UOffsetT(0)
	if t.Claims != nil {
		claimsLength := len(t.Claims)
		claimsOffsets := make([]flatbuffers.UOffsetT, claimsLength)
		for j := 0; j < claimsLength; j++ {
			claimsOffsets[j] = t.Claims[j].Pack(builder)
		}
		TokenStartClaimsVector(builder, claimsLength)
		for j := claimsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(claimsOffsets[j])
		}
		claimsOffset = builder.EndVector(claimsLength)
	}
	idOffset := flatbuffers.UOffsetT(0)
	if t.Id != "" {
		idOffset = builder.CreateString(t.Id)
	}
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	scopeOffset := flatbuffers.UOffsetT(0)
	if t.Scope != nil {
		scopeLength := len(t.Scope)
		scopeOffsets := make([]flatbuffers.UOffsetT, scopeLength)
		for j := 0; j < scopeLength; j++ {
			scopeOffsets[j] = builder.CreateString(t.Scope[j])
		}
		TokenStartScopeVector(builder, scopeLength)
		for j := scopeLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(scopeOffsets[j])
		}
		scopeOffset = builder.EndVector(scopeLength)
	}
	TokenStart(builder)
	TokenAddClaims(builder, claimsOffset)
	TokenAddId(builder, idOffset)
	TokenAddIat(builder, t.Iat)
	TokenAddExp(builder, t.Exp)
	TokenAddName(builder, nameOffset)
	TokenAddPlchandle(builder, t.Plchandle)
	TokenAddScope(builder, scopeOffset)
	return TokenEnd(builder)
}

func (rcv *Token) UnPackTo(t *TokenT) {
	claimsLength := rcv.ClaimsLength()
	t.Claims = make([]*ClaimT, claimsLength)
	for j := 0; j < claimsLength; j++ {
		x := Claim{}
		rcv.Claims(&x, j)
		t.Claims[j] = x.UnPack()
	}
	t.Id = string(rcv.Id())
	t.Iat = rcv.Iat()
	t.Exp = rcv.Exp()
	t.Name = string(rcv.Name())
	t.Plchandle = rcv.Plchandle()
	scopeLength := rcv.ScopeLength()
	t.Scope = make([]string, scopeLength)
	for j := 0; j < scopeLength; j++ {
		t.Scope[j] = string(rcv.Scope(j))
	}
}

func (rcv *Token) UnPack() *TokenT {
	if rcv == nil { return nil }
	t := &TokenT{}
	rcv.UnPackTo(t)
	return t
}

type Token struct {
	_tab flatbuffers.Table
}

func GetRootAsToken(buf []byte, offset flatbuffers.UOffsetT) *Token {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Token{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsToken(buf []byte, offset flatbuffers.UOffsetT) *Token {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Token{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Token) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Token) Table() flatbuffers.Table {
	return rcv._tab
}

/// for all unknown claims - not in fields below
func (rcv *Token) Claims(obj *Claim, j int) bool {
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

func (rcv *Token) ClaimsByKey(obj *Claim, key string) bool{
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *Token) ClaimsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// for all unknown claims - not in fields below
func (rcv *Token) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Token) Iat() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Token) MutateIat(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

func (rcv *Token) Exp() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Token) MutateExp(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func (rcv *Token) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Token) Plchandle() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Token) MutatePlchandle(n uint64) bool {
	return rcv._tab.MutateUint64Slot(14, n)
}

func (rcv *Token) Scope(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Token) ScopeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func TokenStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func TokenAddClaims(builder *flatbuffers.Builder, claims flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(claims), 0)
}
func TokenStartClaimsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TokenAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(id), 0)
}
func TokenAddIat(builder *flatbuffers.Builder, iat uint64) {
	builder.PrependUint64Slot(2, iat, 0)
}
func TokenAddExp(builder *flatbuffers.Builder, exp uint64) {
	builder.PrependUint64Slot(3, exp, 0)
}
func TokenAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(name), 0)
}
func TokenAddPlchandle(builder *flatbuffers.Builder, plchandle uint64) {
	builder.PrependUint64Slot(5, plchandle, 0)
}
func TokenAddScope(builder *flatbuffers.Builder, scope flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(scope), 0)
}
func TokenStartScopeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TokenEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
