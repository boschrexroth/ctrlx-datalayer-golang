// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// general system configuration
type SysCfg struct {
	_tab flatbuffers.Table
}

func GetRootAsSysCfg(buf []byte, offset flatbuffers.UOffsetT) *SysCfg {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SysCfg{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSysCfg(buf []byte, offset flatbuffers.UOffsetT) *SysCfg {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SysCfg{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SysCfg) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SysCfg) Table() flatbuffers.Table {
	return rcv._tab
}

/// configuration of the product coordinate systems 
func (rcv *SysCfg) Pcs(obj *SysCfgPcsAll) *SysCfgPcsAll {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SysCfgPcsAll)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// configuration of the product coordinate systems 
/// system function configuration
func (rcv *SysCfg) Function(obj *SysCfgFunction) *SysCfgFunction {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SysCfgFunction)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// system function configuration
/// internal system configuration
func (rcv *SysCfg) Internal(obj *SysCfgInternal) *SysCfgInternal {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SysCfgInternal)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// internal system configuration
func SysCfgStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func SysCfgAddPcs(builder *flatbuffers.Builder, pcs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(pcs), 0)
}
func SysCfgAddFunction(builder *flatbuffers.Builder, function flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(function), 0)
}
func SysCfgAddInternal(builder *flatbuffers.Builder, internal flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(internal), 0)
}
func SysCfgEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
