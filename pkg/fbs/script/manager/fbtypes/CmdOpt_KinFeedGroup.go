// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CmdOpt_KinFeedGroupT struct {
	Base *CmdOpt_BaseT `json:"base"`
	FeedGroup FeedGroup `json:"feedGroup"`
}

func (t *CmdOpt_KinFeedGroupT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	baseOffset := t.Base.Pack(builder)
	CmdOpt_KinFeedGroupStart(builder)
	CmdOpt_KinFeedGroupAddBase(builder, baseOffset)
	CmdOpt_KinFeedGroupAddFeedGroup(builder, t.FeedGroup)
	return CmdOpt_KinFeedGroupEnd(builder)
}

func (rcv *CmdOpt_KinFeedGroup) UnPackTo(t *CmdOpt_KinFeedGroupT) {
	t.Base = rcv.Base(nil).UnPack()
	t.FeedGroup = rcv.FeedGroup()
}

func (rcv *CmdOpt_KinFeedGroup) UnPack() *CmdOpt_KinFeedGroupT {
	if rcv == nil { return nil }
	t := &CmdOpt_KinFeedGroupT{}
	rcv.UnPackTo(t)
	return t
}

type CmdOpt_KinFeedGroup struct {
	_tab flatbuffers.Table
}

func GetRootAsCmdOpt_KinFeedGroup(buf []byte, offset flatbuffers.UOffsetT) *CmdOpt_KinFeedGroup {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CmdOpt_KinFeedGroup{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmdOpt_KinFeedGroup(buf []byte, offset flatbuffers.UOffsetT) *CmdOpt_KinFeedGroup {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CmdOpt_KinFeedGroup{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CmdOpt_KinFeedGroup) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CmdOpt_KinFeedGroup) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CmdOpt_KinFeedGroup) Base(obj *CmdOpt_Base) *CmdOpt_Base {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(CmdOpt_Base)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *CmdOpt_KinFeedGroup) FeedGroup() FeedGroup {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return FeedGroup(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *CmdOpt_KinFeedGroup) MutateFeedGroup(n FeedGroup) bool {
	return rcv._tab.MutateInt8Slot(6, int8(n))
}

func CmdOpt_KinFeedGroupStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func CmdOpt_KinFeedGroupAddBase(builder *flatbuffers.Builder, base flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(base), 0)
}
func CmdOpt_KinFeedGroupAddFeedGroup(builder *flatbuffers.Builder, feedGroup FeedGroup) {
	builder.PrependInt8Slot(1, int8(feedGroup), 0)
}
func CmdOpt_KinFeedGroupEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
