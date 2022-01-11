// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// configuration for specific functions of this axis
type AxsCfgFunctions struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgFunctions(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgFunctions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgFunctions{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgFunctions(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgFunctions {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgFunctions{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgFunctions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgFunctions) Table() flatbuffers.Table {
	return rcv._tab
}

/// configuration for coupling functions for a single axis
func (rcv *AxsCfgFunctions) Coupling(obj *AxsCfgCoupling) *AxsCfgCoupling {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(AxsCfgCoupling)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// configuration for coupling functions for a single axis
func AxsCfgFunctionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func AxsCfgFunctionsAddCoupling(builder *flatbuffers.Builder, coupling flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(coupling), 0)
}
func AxsCfgFunctionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}