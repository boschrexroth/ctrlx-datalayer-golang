// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AppOperationStateT struct {
	ProgramLoaded bool `json:"programLoaded"`
	DownloadActive bool `json:"downloadActive"`
	OnlineChangeActive bool `json:"onlineChangeActive"`
	StoreBootprojectActive bool `json:"storeBootprojectActive"`
	ForceVariablesActive bool `json:"forceVariablesActive"`
	Exception bool `json:"exception"`
	InitializeActive bool `json:"initializeActive"`
	StoreBootprojectOnlyActive bool `json:"storeBootprojectOnlyActive"`
	ExitActive bool `json:"exitActive"`
	Deleted bool `json:"deleted"`
	ResetActive bool `json:"resetActive"`
	RetainMismatch bool `json:"retainMismatch"`
	BootprojectValid bool `json:"bootprojectValid"`
	LoadBootprojectActive bool `json:"loadBootprojectActive"`
	FlowControlActive bool `json:"flowControlActive"`
	RunInFlash bool `json:"runInFlash"`
	ResetOutputs bool `json:"resetOutputs"`
	CoreDumpLoaded bool `json:"coreDumpLoaded"`
	ExecutenPointsActive bool `json:"executenPointsActive"`
	CoreDumpCreating bool `json:"coreDumpCreating"`
	SingelCycleActive bool `json:"singelCycleActive"`
	ResetDisabled bool `json:"resetDisabled"`
}

func (t *AppOperationStateT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	AppOperationStateStart(builder)
	AppOperationStateAddProgramLoaded(builder, t.ProgramLoaded)
	AppOperationStateAddDownloadActive(builder, t.DownloadActive)
	AppOperationStateAddOnlineChangeActive(builder, t.OnlineChangeActive)
	AppOperationStateAddStoreBootprojectActive(builder, t.StoreBootprojectActive)
	AppOperationStateAddForceVariablesActive(builder, t.ForceVariablesActive)
	AppOperationStateAddException(builder, t.Exception)
	AppOperationStateAddInitializeActive(builder, t.InitializeActive)
	AppOperationStateAddStoreBootprojectOnlyActive(builder, t.StoreBootprojectOnlyActive)
	AppOperationStateAddExitActive(builder, t.ExitActive)
	AppOperationStateAddDeleted(builder, t.Deleted)
	AppOperationStateAddResetActive(builder, t.ResetActive)
	AppOperationStateAddRetainMismatch(builder, t.RetainMismatch)
	AppOperationStateAddBootprojectValid(builder, t.BootprojectValid)
	AppOperationStateAddLoadBootprojectActive(builder, t.LoadBootprojectActive)
	AppOperationStateAddFlowControlActive(builder, t.FlowControlActive)
	AppOperationStateAddRunInFlash(builder, t.RunInFlash)
	AppOperationStateAddResetOutputs(builder, t.ResetOutputs)
	AppOperationStateAddCoreDumpLoaded(builder, t.CoreDumpLoaded)
	AppOperationStateAddExecutenPointsActive(builder, t.ExecutenPointsActive)
	AppOperationStateAddCoreDumpCreating(builder, t.CoreDumpCreating)
	AppOperationStateAddSingelCycleActive(builder, t.SingelCycleActive)
	AppOperationStateAddResetDisabled(builder, t.ResetDisabled)
	return AppOperationStateEnd(builder)
}

func (rcv *AppOperationState) UnPackTo(t *AppOperationStateT) {
	t.ProgramLoaded = rcv.ProgramLoaded()
	t.DownloadActive = rcv.DownloadActive()
	t.OnlineChangeActive = rcv.OnlineChangeActive()
	t.StoreBootprojectActive = rcv.StoreBootprojectActive()
	t.ForceVariablesActive = rcv.ForceVariablesActive()
	t.Exception = rcv.Exception()
	t.InitializeActive = rcv.InitializeActive()
	t.StoreBootprojectOnlyActive = rcv.StoreBootprojectOnlyActive()
	t.ExitActive = rcv.ExitActive()
	t.Deleted = rcv.Deleted()
	t.ResetActive = rcv.ResetActive()
	t.RetainMismatch = rcv.RetainMismatch()
	t.BootprojectValid = rcv.BootprojectValid()
	t.LoadBootprojectActive = rcv.LoadBootprojectActive()
	t.FlowControlActive = rcv.FlowControlActive()
	t.RunInFlash = rcv.RunInFlash()
	t.ResetOutputs = rcv.ResetOutputs()
	t.CoreDumpLoaded = rcv.CoreDumpLoaded()
	t.ExecutenPointsActive = rcv.ExecutenPointsActive()
	t.CoreDumpCreating = rcv.CoreDumpCreating()
	t.SingelCycleActive = rcv.SingelCycleActive()
	t.ResetDisabled = rcv.ResetDisabled()
}

func (rcv *AppOperationState) UnPack() *AppOperationStateT {
	if rcv == nil { return nil }
	t := &AppOperationStateT{}
	rcv.UnPackTo(t)
	return t
}

type AppOperationState struct {
	_tab flatbuffers.Table
}

func GetRootAsAppOperationState(buf []byte, offset flatbuffers.UOffsetT) *AppOperationState {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AppOperationState{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAppOperationState(buf []byte, offset flatbuffers.UOffsetT) *AppOperationState {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AppOperationState{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AppOperationState) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AppOperationState) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AppOperationState) ProgramLoaded() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateProgramLoaded(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *AppOperationState) DownloadActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateDownloadActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *AppOperationState) OnlineChangeActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateOnlineChangeActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func (rcv *AppOperationState) StoreBootprojectActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateStoreBootprojectActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

func (rcv *AppOperationState) ForceVariablesActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateForceVariablesActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

func (rcv *AppOperationState) Exception() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateException(n bool) bool {
	return rcv._tab.MutateBoolSlot(14, n)
}

func (rcv *AppOperationState) InitializeActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateInitializeActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(16, n)
}

func (rcv *AppOperationState) StoreBootprojectOnlyActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateStoreBootprojectOnlyActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(18, n)
}

func (rcv *AppOperationState) ExitActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateExitActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(20, n)
}

func (rcv *AppOperationState) Deleted() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateDeleted(n bool) bool {
	return rcv._tab.MutateBoolSlot(22, n)
}

func (rcv *AppOperationState) ResetActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateResetActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(24, n)
}

func (rcv *AppOperationState) RetainMismatch() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateRetainMismatch(n bool) bool {
	return rcv._tab.MutateBoolSlot(26, n)
}

func (rcv *AppOperationState) BootprojectValid() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateBootprojectValid(n bool) bool {
	return rcv._tab.MutateBoolSlot(28, n)
}

func (rcv *AppOperationState) LoadBootprojectActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateLoadBootprojectActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(30, n)
}

func (rcv *AppOperationState) FlowControlActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateFlowControlActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(32, n)
}

func (rcv *AppOperationState) RunInFlash() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateRunInFlash(n bool) bool {
	return rcv._tab.MutateBoolSlot(34, n)
}

func (rcv *AppOperationState) ResetOutputs() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateResetOutputs(n bool) bool {
	return rcv._tab.MutateBoolSlot(36, n)
}

func (rcv *AppOperationState) CoreDumpLoaded() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateCoreDumpLoaded(n bool) bool {
	return rcv._tab.MutateBoolSlot(38, n)
}

func (rcv *AppOperationState) ExecutenPointsActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateExecutenPointsActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(40, n)
}

func (rcv *AppOperationState) CoreDumpCreating() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateCoreDumpCreating(n bool) bool {
	return rcv._tab.MutateBoolSlot(42, n)
}

func (rcv *AppOperationState) SingelCycleActive() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateSingelCycleActive(n bool) bool {
	return rcv._tab.MutateBoolSlot(44, n)
}

func (rcv *AppOperationState) ResetDisabled() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *AppOperationState) MutateResetDisabled(n bool) bool {
	return rcv._tab.MutateBoolSlot(46, n)
}

func AppOperationStateStart(builder *flatbuffers.Builder) {
	builder.StartObject(22)
}
func AppOperationStateAddProgramLoaded(builder *flatbuffers.Builder, programLoaded bool) {
	builder.PrependBoolSlot(0, programLoaded, false)
}
func AppOperationStateAddDownloadActive(builder *flatbuffers.Builder, downloadActive bool) {
	builder.PrependBoolSlot(1, downloadActive, false)
}
func AppOperationStateAddOnlineChangeActive(builder *flatbuffers.Builder, onlineChangeActive bool) {
	builder.PrependBoolSlot(2, onlineChangeActive, false)
}
func AppOperationStateAddStoreBootprojectActive(builder *flatbuffers.Builder, storeBootprojectActive bool) {
	builder.PrependBoolSlot(3, storeBootprojectActive, false)
}
func AppOperationStateAddForceVariablesActive(builder *flatbuffers.Builder, forceVariablesActive bool) {
	builder.PrependBoolSlot(4, forceVariablesActive, false)
}
func AppOperationStateAddException(builder *flatbuffers.Builder, exception bool) {
	builder.PrependBoolSlot(5, exception, false)
}
func AppOperationStateAddInitializeActive(builder *flatbuffers.Builder, initializeActive bool) {
	builder.PrependBoolSlot(6, initializeActive, false)
}
func AppOperationStateAddStoreBootprojectOnlyActive(builder *flatbuffers.Builder, storeBootprojectOnlyActive bool) {
	builder.PrependBoolSlot(7, storeBootprojectOnlyActive, false)
}
func AppOperationStateAddExitActive(builder *flatbuffers.Builder, exitActive bool) {
	builder.PrependBoolSlot(8, exitActive, false)
}
func AppOperationStateAddDeleted(builder *flatbuffers.Builder, deleted bool) {
	builder.PrependBoolSlot(9, deleted, false)
}
func AppOperationStateAddResetActive(builder *flatbuffers.Builder, resetActive bool) {
	builder.PrependBoolSlot(10, resetActive, false)
}
func AppOperationStateAddRetainMismatch(builder *flatbuffers.Builder, retainMismatch bool) {
	builder.PrependBoolSlot(11, retainMismatch, false)
}
func AppOperationStateAddBootprojectValid(builder *flatbuffers.Builder, bootprojectValid bool) {
	builder.PrependBoolSlot(12, bootprojectValid, false)
}
func AppOperationStateAddLoadBootprojectActive(builder *flatbuffers.Builder, loadBootprojectActive bool) {
	builder.PrependBoolSlot(13, loadBootprojectActive, false)
}
func AppOperationStateAddFlowControlActive(builder *flatbuffers.Builder, flowControlActive bool) {
	builder.PrependBoolSlot(14, flowControlActive, false)
}
func AppOperationStateAddRunInFlash(builder *flatbuffers.Builder, runInFlash bool) {
	builder.PrependBoolSlot(15, runInFlash, false)
}
func AppOperationStateAddResetOutputs(builder *flatbuffers.Builder, resetOutputs bool) {
	builder.PrependBoolSlot(16, resetOutputs, false)
}
func AppOperationStateAddCoreDumpLoaded(builder *flatbuffers.Builder, coreDumpLoaded bool) {
	builder.PrependBoolSlot(17, coreDumpLoaded, false)
}
func AppOperationStateAddExecutenPointsActive(builder *flatbuffers.Builder, executenPointsActive bool) {
	builder.PrependBoolSlot(18, executenPointsActive, false)
}
func AppOperationStateAddCoreDumpCreating(builder *flatbuffers.Builder, coreDumpCreating bool) {
	builder.PrependBoolSlot(19, coreDumpCreating, false)
}
func AppOperationStateAddSingelCycleActive(builder *flatbuffers.Builder, singelCycleActive bool) {
	builder.PrependBoolSlot(20, singelCycleActive, false)
}
func AppOperationStateAddResetDisabled(builder *flatbuffers.Builder, resetDisabled bool) {
	builder.PrependBoolSlot(21, resetDisabled, false)
}
func AppOperationStateEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
