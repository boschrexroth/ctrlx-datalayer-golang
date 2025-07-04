// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// general system configuration
type SysCfgT struct {
	Pcs *SysCfgPcsAllT `json:"pcs"`
	Function *SysCfgFunctionT `json:"function"`
	Internal *SysCfgInternalT `json:"internal"`
	SafeAreas *SysCfgSafeAreaAllT `json:"safeAreas"`
	RtInputs *RTInputsCfgT `json:"rtInputs"`
	ToolData *SysCfgToolDataAllT `json:"toolData"`
	BeltCs *SysCfgBeltCsAllT `json:"beltCS"`
	Ccs *SysCfgCcsAllT `json:"ccs"`
	UiData string `json:"uiData"`
}

func (t *SysCfgT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	pcsOffset := t.Pcs.Pack(builder)
	functionOffset := t.Function.Pack(builder)
	internalOffset := t.Internal.Pack(builder)
	safeAreasOffset := t.SafeAreas.Pack(builder)
	rtInputsOffset := t.RtInputs.Pack(builder)
	toolDataOffset := t.ToolData.Pack(builder)
	beltCsOffset := t.BeltCs.Pack(builder)
	ccsOffset := t.Ccs.Pack(builder)
	uiDataOffset := flatbuffers.UOffsetT(0)
	if t.UiData != "" {
		uiDataOffset = builder.CreateString(t.UiData)
	}
	SysCfgStart(builder)
	SysCfgAddPcs(builder, pcsOffset)
	SysCfgAddFunction(builder, functionOffset)
	SysCfgAddInternal(builder, internalOffset)
	SysCfgAddSafeAreas(builder, safeAreasOffset)
	SysCfgAddRtInputs(builder, rtInputsOffset)
	SysCfgAddToolData(builder, toolDataOffset)
	SysCfgAddBeltCs(builder, beltCsOffset)
	SysCfgAddCcs(builder, ccsOffset)
	SysCfgAddUiData(builder, uiDataOffset)
	return SysCfgEnd(builder)
}

func (rcv *SysCfg) UnPackTo(t *SysCfgT) {
	t.Pcs = rcv.Pcs(nil).UnPack()
	t.Function = rcv.Function(nil).UnPack()
	t.Internal = rcv.Internal(nil).UnPack()
	t.SafeAreas = rcv.SafeAreas(nil).UnPack()
	t.RtInputs = rcv.RtInputs(nil).UnPack()
	t.ToolData = rcv.ToolData(nil).UnPack()
	t.BeltCs = rcv.BeltCs(nil).UnPack()
	t.Ccs = rcv.Ccs(nil).UnPack()
	t.UiData = string(rcv.UiData())
}

func (rcv *SysCfg) UnPack() *SysCfgT {
	if rcv == nil { return nil }
	t := &SysCfgT{}
	rcv.UnPackTo(t)
	return t
}

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
/// configuration of the safe and work areas
func (rcv *SysCfg) SafeAreas(obj *SysCfgSafeAreaAll) *SysCfgSafeAreaAll {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SysCfgSafeAreaAll)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// configuration of the safe and work areas
/// configuration of the real-time inputs of the kinematics
func (rcv *SysCfg) RtInputs(obj *RTInputsCfg) *RTInputsCfg {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(RTInputsCfg)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// configuration of the real-time inputs of the kinematics
/// configuration of the tool data
func (rcv *SysCfg) ToolData(obj *SysCfgToolDataAll) *SysCfgToolDataAll {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SysCfgToolDataAll)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// configuration of the tool data
/// configuration of the belt cs data
func (rcv *SysCfg) BeltCs(obj *SysCfgBeltCsAll) *SysCfgBeltCsAll {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SysCfgBeltCsAll)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// configuration of the belt cs data
/// configuration of the camera coordinate system
func (rcv *SysCfg) Ccs(obj *SysCfgCcsAll) *SysCfgCcsAll {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SysCfgCcsAll)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// configuration of the camera coordinate system
/// System UI data
func (rcv *SysCfg) UiData() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// System UI data
func SysCfgStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
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
func SysCfgAddSafeAreas(builder *flatbuffers.Builder, safeAreas flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(safeAreas), 0)
}
func SysCfgAddRtInputs(builder *flatbuffers.Builder, rtInputs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(rtInputs), 0)
}
func SysCfgAddToolData(builder *flatbuffers.Builder, toolData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(toolData), 0)
}
func SysCfgAddBeltCs(builder *flatbuffers.Builder, beltCs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(beltCs), 0)
}
func SysCfgAddCcs(builder *flatbuffers.Builder, ccs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(ccs), 0)
}
func SysCfgAddUiData(builder *flatbuffers.Builder, uiData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(uiData), 0)
}
func SysCfgEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
