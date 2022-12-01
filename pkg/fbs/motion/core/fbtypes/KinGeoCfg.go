// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// common configuration of orientation
type KinGeoCfgT struct {
	Orientation *KinOriCfgT
}

func (t *KinGeoCfgT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	orientationOffset := t.Orientation.Pack(builder)
	KinGeoCfgStart(builder)
	KinGeoCfgAddOrientation(builder, orientationOffset)
	return KinGeoCfgEnd(builder)
}

func (rcv *KinGeoCfg) UnPackTo(t *KinGeoCfgT) {
	t.Orientation = rcv.Orientation(nil).UnPack()
}

func (rcv *KinGeoCfg) UnPack() *KinGeoCfgT {
	if rcv == nil { return nil }
	t := &KinGeoCfgT{}
	rcv.UnPackTo(t)
	return t
}

type KinGeoCfg struct {
	_tab flatbuffers.Table
}

func GetRootAsKinGeoCfg(buf []byte, offset flatbuffers.UOffsetT) *KinGeoCfg {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &KinGeoCfg{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsKinGeoCfg(buf []byte, offset flatbuffers.UOffsetT) *KinGeoCfg {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &KinGeoCfg{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *KinGeoCfg) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KinGeoCfg) Table() flatbuffers.Table {
	return rcv._tab
}

/// orientation config
func (rcv *KinGeoCfg) Orientation(obj *KinOriCfg) *KinOriCfg {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(KinOriCfg)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// orientation config
func KinGeoCfgStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func KinGeoCfgAddOrientation(builder *flatbuffers.Builder, orientation flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(orientation), 0)
}
func KinGeoCfgEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}