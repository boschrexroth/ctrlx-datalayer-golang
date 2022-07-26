// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ClaimT struct {
	Claim string
	Value string
}

func (t *ClaimT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	claimOffset := builder.CreateString(t.Claim)
	valueOffset := builder.CreateString(t.Value)
	ClaimStart(builder)
	ClaimAddClaim(builder, claimOffset)
	ClaimAddValue(builder, valueOffset)
	return ClaimEnd(builder)
}

func (rcv *Claim) UnPackTo(t *ClaimT) {
	t.Claim = string(rcv.Claim())
	t.Value = string(rcv.Value())
}

func (rcv *Claim) UnPack() *ClaimT {
	if rcv == nil { return nil }
	t := &ClaimT{}
	rcv.UnPackTo(t)
	return t
}

type Claim struct {
	_tab flatbuffers.Table
}

func GetRootAsClaim(buf []byte, offset flatbuffers.UOffsetT) *Claim {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Claim{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsClaim(buf []byte, offset flatbuffers.UOffsetT) *Claim {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Claim{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Claim) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Claim) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Claim) Claim() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Claim) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ClaimStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ClaimAddClaim(builder *flatbuffers.Builder, claim flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(claim), 0)
}
func ClaimAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func ClaimEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
