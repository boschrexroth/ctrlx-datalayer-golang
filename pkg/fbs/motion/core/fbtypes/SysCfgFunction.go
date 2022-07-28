// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// system function configuration items
type SysCfgFunctionT struct {
	IgnoreAxisProfiles bool
}

func (t *SysCfgFunctionT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	SysCfgFunctionStart(builder)
	SysCfgFunctionAddIgnoreAxisProfiles(builder, t.IgnoreAxisProfiles)
	return SysCfgFunctionEnd(builder)
}

func (rcv *SysCfgFunction) UnPackTo(t *SysCfgFunctionT) {
	t.IgnoreAxisProfiles = rcv.IgnoreAxisProfiles()
}

func (rcv *SysCfgFunction) UnPack() *SysCfgFunctionT {
	if rcv == nil { return nil }
	t := &SysCfgFunctionT{}
	rcv.UnPackTo(t)
	return t
}

type SysCfgFunction struct {
	_tab flatbuffers.Table
}

func GetRootAsSysCfgFunction(buf []byte, offset flatbuffers.UOffsetT) *SysCfgFunction {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SysCfgFunction{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSysCfgFunction(buf []byte, offset flatbuffers.UOffsetT) *SysCfgFunction {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SysCfgFunction{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SysCfgFunction) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SysCfgFunction) Table() flatbuffers.Table {
	return rcv._tab
}

/// ignore the configured axis profiles in BOOTING?
func (rcv *SysCfgFunction) IgnoreAxisProfiles() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// ignore the configured axis profiles in BOOTING?
func (rcv *SysCfgFunction) MutateIgnoreAxisProfiles(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func SysCfgFunctionStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SysCfgFunctionAddIgnoreAxisProfiles(builder *flatbuffers.Builder, ignoreAxisProfiles bool) {
	builder.PrependBoolSlot(0, ignoreAxisProfiles, false)
}
func SysCfgFunctionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
