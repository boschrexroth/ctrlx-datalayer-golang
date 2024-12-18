// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters for the command option belt sync
type CmdOptionBeltSyncT struct {
	PermType string `json:"permType"`
	BeltName string `json:"beltName"`
	BeltValue float64 `json:"beltValue"`
}

func (t *CmdOptionBeltSyncT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	permTypeOffset := flatbuffers.UOffsetT(0)
	if t.PermType != "" {
		permTypeOffset = builder.CreateString(t.PermType)
	}
	beltNameOffset := flatbuffers.UOffsetT(0)
	if t.BeltName != "" {
		beltNameOffset = builder.CreateString(t.BeltName)
	}
	CmdOptionBeltSyncStart(builder)
	CmdOptionBeltSyncAddPermType(builder, permTypeOffset)
	CmdOptionBeltSyncAddBeltName(builder, beltNameOffset)
	CmdOptionBeltSyncAddBeltValue(builder, t.BeltValue)
	return CmdOptionBeltSyncEnd(builder)
}

func (rcv *CmdOptionBeltSync) UnPackTo(t *CmdOptionBeltSyncT) {
	t.PermType = string(rcv.PermType())
	t.BeltName = string(rcv.BeltName())
	t.BeltValue = rcv.BeltValue()
}

func (rcv *CmdOptionBeltSync) UnPack() *CmdOptionBeltSyncT {
	if rcv == nil { return nil }
	t := &CmdOptionBeltSyncT{}
	rcv.UnPackTo(t)
	return t
}

type CmdOptionBeltSync struct {
	_tab flatbuffers.Table
}

func GetRootAsCmdOptionBeltSync(buf []byte, offset flatbuffers.UOffsetT) *CmdOptionBeltSync {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CmdOptionBeltSync{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCmdOptionBeltSync(buf []byte, offset flatbuffers.UOffsetT) *CmdOptionBeltSync {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CmdOptionBeltSync{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CmdOptionBeltSync) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CmdOptionBeltSync) Table() flatbuffers.Table {
	return rcv._tab
}

/// permanent type (e.g. "PermOn", "PermOff")
func (rcv *CmdOptionBeltSync) PermType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// permanent type (e.g. "PermOn", "PermOff")
/// belt name
func (rcv *CmdOptionBeltSync) BeltName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// belt name
/// belt value
func (rcv *CmdOptionBeltSync) BeltValue() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// belt value
func (rcv *CmdOptionBeltSync) MutateBeltValue(n float64) bool {
	return rcv._tab.MutateFloat64Slot(8, n)
}

func CmdOptionBeltSyncStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func CmdOptionBeltSyncAddPermType(builder *flatbuffers.Builder, permType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(permType), 0)
}
func CmdOptionBeltSyncAddBeltName(builder *flatbuffers.Builder, beltName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(beltName), 0)
}
func CmdOptionBeltSyncAddBeltValue(builder *flatbuffers.Builder, beltValue float64) {
	builder.PrependFloat64Slot(2, beltValue, 0.0)
}
func CmdOptionBeltSyncEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
