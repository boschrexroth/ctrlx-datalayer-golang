// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Configuration of the units in the flexprofile
type CfgFlexProfileUnitsT struct {
	PosUnitMaster string `json:"posUnitMaster"`
	VelUnitMaster string `json:"velUnitMaster"`
	PosUnitSlave string `json:"posUnitSlave"`
	VelUnitSlave string `json:"velUnitSlave"`
	AccUnitSlave string `json:"accUnitSlave"`
	JrkUnitSlave string `json:"jrkUnitSlave"`
	ActivateUnits bool `json:"activateUnits"`
}

func (t *CfgFlexProfileUnitsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	posUnitMasterOffset := flatbuffers.UOffsetT(0)
	if t.PosUnitMaster != "" {
		posUnitMasterOffset = builder.CreateString(t.PosUnitMaster)
	}
	velUnitMasterOffset := flatbuffers.UOffsetT(0)
	if t.VelUnitMaster != "" {
		velUnitMasterOffset = builder.CreateString(t.VelUnitMaster)
	}
	posUnitSlaveOffset := flatbuffers.UOffsetT(0)
	if t.PosUnitSlave != "" {
		posUnitSlaveOffset = builder.CreateString(t.PosUnitSlave)
	}
	velUnitSlaveOffset := flatbuffers.UOffsetT(0)
	if t.VelUnitSlave != "" {
		velUnitSlaveOffset = builder.CreateString(t.VelUnitSlave)
	}
	accUnitSlaveOffset := flatbuffers.UOffsetT(0)
	if t.AccUnitSlave != "" {
		accUnitSlaveOffset = builder.CreateString(t.AccUnitSlave)
	}
	jrkUnitSlaveOffset := flatbuffers.UOffsetT(0)
	if t.JrkUnitSlave != "" {
		jrkUnitSlaveOffset = builder.CreateString(t.JrkUnitSlave)
	}
	CfgFlexProfileUnitsStart(builder)
	CfgFlexProfileUnitsAddPosUnitMaster(builder, posUnitMasterOffset)
	CfgFlexProfileUnitsAddVelUnitMaster(builder, velUnitMasterOffset)
	CfgFlexProfileUnitsAddPosUnitSlave(builder, posUnitSlaveOffset)
	CfgFlexProfileUnitsAddVelUnitSlave(builder, velUnitSlaveOffset)
	CfgFlexProfileUnitsAddAccUnitSlave(builder, accUnitSlaveOffset)
	CfgFlexProfileUnitsAddJrkUnitSlave(builder, jrkUnitSlaveOffset)
	CfgFlexProfileUnitsAddActivateUnits(builder, t.ActivateUnits)
	return CfgFlexProfileUnitsEnd(builder)
}

func (rcv *CfgFlexProfileUnits) UnPackTo(t *CfgFlexProfileUnitsT) {
	t.PosUnitMaster = string(rcv.PosUnitMaster())
	t.VelUnitMaster = string(rcv.VelUnitMaster())
	t.PosUnitSlave = string(rcv.PosUnitSlave())
	t.VelUnitSlave = string(rcv.VelUnitSlave())
	t.AccUnitSlave = string(rcv.AccUnitSlave())
	t.JrkUnitSlave = string(rcv.JrkUnitSlave())
	t.ActivateUnits = rcv.ActivateUnits()
}

func (rcv *CfgFlexProfileUnits) UnPack() *CfgFlexProfileUnitsT {
	if rcv == nil { return nil }
	t := &CfgFlexProfileUnitsT{}
	rcv.UnPackTo(t)
	return t
}

type CfgFlexProfileUnits struct {
	_tab flatbuffers.Table
}

func GetRootAsCfgFlexProfileUnits(buf []byte, offset flatbuffers.UOffsetT) *CfgFlexProfileUnits {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CfgFlexProfileUnits{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCfgFlexProfileUnits(buf []byte, offset flatbuffers.UOffsetT) *CfgFlexProfileUnits {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CfgFlexProfileUnits{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CfgFlexProfileUnits) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CfgFlexProfileUnits) Table() flatbuffers.Table {
	return rcv._tab
}

/// position unit of the master axis
func (rcv *CfgFlexProfileUnits) PosUnitMaster() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// position unit of the master axis
/// velocity unit of the master axis
func (rcv *CfgFlexProfileUnits) VelUnitMaster() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// velocity unit of the master axis
/// position unit of the slave axis
func (rcv *CfgFlexProfileUnits) PosUnitSlave() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// position unit of the slave axis
/// velocity unit of the slave axis
func (rcv *CfgFlexProfileUnits) VelUnitSlave() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// velocity unit of the slave axis
/// acceleration of the slave axis
func (rcv *CfgFlexProfileUnits) AccUnitSlave() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// acceleration of the slave axis
/// jerk unit of the slave axis
func (rcv *CfgFlexProfileUnits) JrkUnitSlave() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// jerk unit of the slave axis
/// units should be considered
func (rcv *CfgFlexProfileUnits) ActivateUnits() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// units should be considered
func (rcv *CfgFlexProfileUnits) MutateActivateUnits(n bool) bool {
	return rcv._tab.MutateBoolSlot(16, n)
}

func CfgFlexProfileUnitsStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func CfgFlexProfileUnitsAddPosUnitMaster(builder *flatbuffers.Builder, posUnitMaster flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(posUnitMaster), 0)
}
func CfgFlexProfileUnitsAddVelUnitMaster(builder *flatbuffers.Builder, velUnitMaster flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(velUnitMaster), 0)
}
func CfgFlexProfileUnitsAddPosUnitSlave(builder *flatbuffers.Builder, posUnitSlave flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(posUnitSlave), 0)
}
func CfgFlexProfileUnitsAddVelUnitSlave(builder *flatbuffers.Builder, velUnitSlave flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(velUnitSlave), 0)
}
func CfgFlexProfileUnitsAddAccUnitSlave(builder *flatbuffers.Builder, accUnitSlave flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(accUnitSlave), 0)
}
func CfgFlexProfileUnitsAddJrkUnitSlave(builder *flatbuffers.Builder, jrkUnitSlave flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(jrkUnitSlave), 0)
}
func CfgFlexProfileUnitsAddActivateUnits(builder *flatbuffers.Builder, activateUnits bool) {
	builder.PrependBoolSlot(6, activateUnits, false)
}
func CfgFlexProfileUnitsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}