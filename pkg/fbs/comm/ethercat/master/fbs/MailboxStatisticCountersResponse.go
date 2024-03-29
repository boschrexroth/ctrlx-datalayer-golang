// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Mailbox statistic counters response
type MailboxStatisticCountersResponseT struct {
	Aoe *ProtocolStatisticT `json:"aoe"`
	Coe *ProtocolStatisticT `json:"coe"`
	Eoe *ProtocolStatisticT `json:"eoe"`
	Foe *ProtocolStatisticT `json:"foe"`
	Soe *ProtocolStatisticT `json:"soe"`
	Voe *ProtocolStatisticT `json:"voe"`
	Raw *ProtocolStatisticT `json:"raw"`
}

func (t *MailboxStatisticCountersResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	MailboxStatisticCountersResponseStart(builder)
	aoeOffset := t.Aoe.Pack(builder)
	MailboxStatisticCountersResponseAddAoe(builder, aoeOffset)
	coeOffset := t.Coe.Pack(builder)
	MailboxStatisticCountersResponseAddCoe(builder, coeOffset)
	eoeOffset := t.Eoe.Pack(builder)
	MailboxStatisticCountersResponseAddEoe(builder, eoeOffset)
	foeOffset := t.Foe.Pack(builder)
	MailboxStatisticCountersResponseAddFoe(builder, foeOffset)
	soeOffset := t.Soe.Pack(builder)
	MailboxStatisticCountersResponseAddSoe(builder, soeOffset)
	voeOffset := t.Voe.Pack(builder)
	MailboxStatisticCountersResponseAddVoe(builder, voeOffset)
	rawOffset := t.Raw.Pack(builder)
	MailboxStatisticCountersResponseAddRaw(builder, rawOffset)
	return MailboxStatisticCountersResponseEnd(builder)
}

func (rcv *MailboxStatisticCountersResponse) UnPackTo(t *MailboxStatisticCountersResponseT) {
	t.Aoe = rcv.Aoe(nil).UnPack()
	t.Coe = rcv.Coe(nil).UnPack()
	t.Eoe = rcv.Eoe(nil).UnPack()
	t.Foe = rcv.Foe(nil).UnPack()
	t.Soe = rcv.Soe(nil).UnPack()
	t.Voe = rcv.Voe(nil).UnPack()
	t.Raw = rcv.Raw(nil).UnPack()
}

func (rcv *MailboxStatisticCountersResponse) UnPack() *MailboxStatisticCountersResponseT {
	if rcv == nil { return nil }
	t := &MailboxStatisticCountersResponseT{}
	rcv.UnPackTo(t)
	return t
}

type MailboxStatisticCountersResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsMailboxStatisticCountersResponse(buf []byte, offset flatbuffers.UOffsetT) *MailboxStatisticCountersResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MailboxStatisticCountersResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMailboxStatisticCountersResponse(buf []byte, offset flatbuffers.UOffsetT) *MailboxStatisticCountersResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MailboxStatisticCountersResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MailboxStatisticCountersResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MailboxStatisticCountersResponse) Table() flatbuffers.Table {
	return rcv._tab
}

///AoE protocol statistic
func (rcv *MailboxStatisticCountersResponse) Aoe(obj *ProtocolStatistic) *ProtocolStatistic {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatistic)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

///AoE protocol statistic
///CoE protocol statistic
func (rcv *MailboxStatisticCountersResponse) Coe(obj *ProtocolStatistic) *ProtocolStatistic {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatistic)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

///CoE protocol statistic
///EoE protocol statistic
func (rcv *MailboxStatisticCountersResponse) Eoe(obj *ProtocolStatistic) *ProtocolStatistic {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatistic)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

///EoE protocol statistic
///FoE protocol statistic
func (rcv *MailboxStatisticCountersResponse) Foe(obj *ProtocolStatistic) *ProtocolStatistic {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatistic)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

///FoE protocol statistic
///SoE protocol statistic
func (rcv *MailboxStatisticCountersResponse) Soe(obj *ProtocolStatistic) *ProtocolStatistic {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatistic)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

///SoE protocol statistic
///VoE protocol statistic
func (rcv *MailboxStatisticCountersResponse) Voe(obj *ProtocolStatistic) *ProtocolStatistic {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatistic)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

///VoE protocol statistic
///Raw protocol statistic
func (rcv *MailboxStatisticCountersResponse) Raw(obj *ProtocolStatistic) *ProtocolStatistic {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ProtocolStatistic)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

///Raw protocol statistic
func MailboxStatisticCountersResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func MailboxStatisticCountersResponseAddAoe(builder *flatbuffers.Builder, aoe flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(aoe), 0)
}
func MailboxStatisticCountersResponseAddCoe(builder *flatbuffers.Builder, coe flatbuffers.UOffsetT) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(coe), 0)
}
func MailboxStatisticCountersResponseAddEoe(builder *flatbuffers.Builder, eoe flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(eoe), 0)
}
func MailboxStatisticCountersResponseAddFoe(builder *flatbuffers.Builder, foe flatbuffers.UOffsetT) {
	builder.PrependStructSlot(3, flatbuffers.UOffsetT(foe), 0)
}
func MailboxStatisticCountersResponseAddSoe(builder *flatbuffers.Builder, soe flatbuffers.UOffsetT) {
	builder.PrependStructSlot(4, flatbuffers.UOffsetT(soe), 0)
}
func MailboxStatisticCountersResponseAddVoe(builder *flatbuffers.Builder, voe flatbuffers.UOffsetT) {
	builder.PrependStructSlot(5, flatbuffers.UOffsetT(voe), 0)
}
func MailboxStatisticCountersResponseAddRaw(builder *flatbuffers.Builder, raw flatbuffers.UOffsetT) {
	builder.PrependStructSlot(6, flatbuffers.UOffsetT(raw), 0)
}
func MailboxStatisticCountersResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
