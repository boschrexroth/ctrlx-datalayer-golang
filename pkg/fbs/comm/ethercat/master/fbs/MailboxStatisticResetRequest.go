// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MailboxStatisticResetRequestT struct {
	Aoe *ProtocolStatisticResetT `json:"aoe"`
	Coe *ProtocolStatisticResetT `json:"coe"`
	Foe *ProtocolStatisticResetT `json:"foe"`
	Eoe *ProtocolStatisticResetT `json:"eoe"`
	Soe *ProtocolStatisticResetT `json:"soe"`
	Voe *ProtocolStatisticResetT `json:"voe"`
	Raw *ProtocolStatisticResetT `json:"raw"`
}

func (t *MailboxStatisticResetRequestT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	MailboxStatisticResetRequestStart(builder)
	aoeOffset := t.Aoe.Pack(builder)
	MailboxStatisticResetRequestAddAoe(builder, aoeOffset)
	coeOffset := t.Coe.Pack(builder)
	MailboxStatisticResetRequestAddCoe(builder, coeOffset)
	foeOffset := t.Foe.Pack(builder)
	MailboxStatisticResetRequestAddFoe(builder, foeOffset)
	eoeOffset := t.Eoe.Pack(builder)
	MailboxStatisticResetRequestAddEoe(builder, eoeOffset)
	soeOffset := t.Soe.Pack(builder)
	MailboxStatisticResetRequestAddSoe(builder, soeOffset)
	voeOffset := t.Voe.Pack(builder)
	MailboxStatisticResetRequestAddVoe(builder, voeOffset)
	rawOffset := t.Raw.Pack(builder)
	MailboxStatisticResetRequestAddRaw(builder, rawOffset)
	return MailboxStatisticResetRequestEnd(builder)
}

func (rcv *MailboxStatisticResetRequest) UnPackTo(t *MailboxStatisticResetRequestT) {
	t.Aoe = rcv.Aoe(nil).UnPack()
	t.Coe = rcv.Coe(nil).UnPack()
	t.Foe = rcv.Foe(nil).UnPack()
	t.Eoe = rcv.Eoe(nil).UnPack()
	t.Soe = rcv.Soe(nil).UnPack()
	t.Voe = rcv.Voe(nil).UnPack()
	t.Raw = rcv.Raw(nil).UnPack()
}

func (rcv *MailboxStatisticResetRequest) UnPack() *MailboxStatisticResetRequestT {
	if rcv == nil { return nil }
	t := &MailboxStatisticResetRequestT{}
	rcv.UnPackTo(t)
	return t
}

type MailboxStatisticResetRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsMailboxStatisticResetRequest(buf []byte, offset flatbuffers.UOffsetT) *MailboxStatisticResetRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MailboxStatisticResetRequest{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMailboxStatisticResetRequest(buf []byte, offset flatbuffers.UOffsetT) *MailboxStatisticResetRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MailboxStatisticResetRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MailboxStatisticResetRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MailboxStatisticResetRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MailboxStatisticResetRequest) Aoe(obj *ProtocolStatisticReset) *ProtocolStatisticReset {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatisticReset)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *MailboxStatisticResetRequest) Coe(obj *ProtocolStatisticReset) *ProtocolStatisticReset {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatisticReset)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *MailboxStatisticResetRequest) Foe(obj *ProtocolStatisticReset) *ProtocolStatisticReset {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatisticReset)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *MailboxStatisticResetRequest) Eoe(obj *ProtocolStatisticReset) *ProtocolStatisticReset {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatisticReset)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *MailboxStatisticResetRequest) Soe(obj *ProtocolStatisticReset) *ProtocolStatisticReset {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatisticReset)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *MailboxStatisticResetRequest) Voe(obj *ProtocolStatisticReset) *ProtocolStatisticReset {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatisticReset)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *MailboxStatisticResetRequest) Raw(obj *ProtocolStatisticReset) *ProtocolStatisticReset {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatisticReset)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func MailboxStatisticResetRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func MailboxStatisticResetRequestAddAoe(builder *flatbuffers.Builder, aoe flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(aoe), 0)
}
func MailboxStatisticResetRequestAddCoe(builder *flatbuffers.Builder, coe flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(coe), 0)
}
func MailboxStatisticResetRequestAddFoe(builder *flatbuffers.Builder, foe flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(foe), 0)
}
func MailboxStatisticResetRequestAddEoe(builder *flatbuffers.Builder, eoe flatbuffers.UOffsetT) {
	builder.PrependStructSlot(3, flatbuffers.UOffsetT(eoe), 0)
}
func MailboxStatisticResetRequestAddSoe(builder *flatbuffers.Builder, soe flatbuffers.UOffsetT) {
	builder.PrependStructSlot(4, flatbuffers.UOffsetT(soe), 0)
}
func MailboxStatisticResetRequestAddVoe(builder *flatbuffers.Builder, voe flatbuffers.UOffsetT) {
	builder.PrependStructSlot(5, flatbuffers.UOffsetT(voe), 0)
}
func MailboxStatisticResetRequestAddRaw(builder *flatbuffers.Builder, raw flatbuffers.UOffsetT) {
	builder.PrependStructSlot(6, flatbuffers.UOffsetT(raw), 0)
}
func MailboxStatisticResetRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
