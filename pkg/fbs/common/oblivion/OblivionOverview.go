// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package oblivion

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OblivionOverviewT struct {
	Active bool `json:"active"`
	Current *OblivionCurrentT `json:"current"`
	Total *OblivionTotalT `json:"total"`
}

func (t *OblivionOverviewT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	currentOffset := t.Current.Pack(builder)
	totalOffset := t.Total.Pack(builder)
	OblivionOverviewStart(builder)
	OblivionOverviewAddActive(builder, t.Active)
	OblivionOverviewAddCurrent(builder, currentOffset)
	OblivionOverviewAddTotal(builder, totalOffset)
	return OblivionOverviewEnd(builder)
}

func (rcv *OblivionOverview) UnPackTo(t *OblivionOverviewT) {
	t.Active = rcv.Active()
	t.Current = rcv.Current(nil).UnPack()
	t.Total = rcv.Total(nil).UnPack()
}

func (rcv *OblivionOverview) UnPack() *OblivionOverviewT {
	if rcv == nil { return nil }
	t := &OblivionOverviewT{}
	rcv.UnPackTo(t)
	return t
}

type OblivionOverview struct {
	_tab flatbuffers.Table
}

func GetRootAsOblivionOverview(buf []byte, offset flatbuffers.UOffsetT) *OblivionOverview {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OblivionOverview{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsOblivionOverview(buf []byte, offset flatbuffers.UOffsetT) *OblivionOverview {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &OblivionOverview{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *OblivionOverview) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OblivionOverview) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OblivionOverview) Active() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *OblivionOverview) MutateActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *OblivionOverview) Current(obj *OblivionCurrent) *OblivionCurrent {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(OblivionCurrent)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *OblivionOverview) Total(obj *OblivionTotal) *OblivionTotal {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(OblivionTotal)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func OblivionOverviewStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func OblivionOverviewAddActive(builder *flatbuffers.Builder, active bool) {
	builder.PrependBoolSlot(0, active, false)
}
func OblivionOverviewAddCurrent(builder *flatbuffers.Builder, current flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(current), 0)
}
func OblivionOverviewAddTotal(builder *flatbuffers.Builder, total flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(total), 0)
}
func OblivionOverviewEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
