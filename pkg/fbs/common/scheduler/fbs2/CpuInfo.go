// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs2

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Information about CPU cores of ctrlX CORE
type CpuInfoT struct {
	CpuCoresTotal []uint32 `json:"cpuCoresTotal"`
	CpuCoresActive []uint32 `json:"cpuCoresActive"`
	CpuCoresRealtime []uint32 `json:"cpuCoresRealtime"`
	CpuCoresNonRealtime []uint32 `json:"cpuCoresNonRealtime"`
	CpuCoreRealtimeMax int32 `json:"cpuCoreRealtimeMax"`
	CpuCoreRealtimeMin int32 `json:"cpuCoreRealtimeMin"`
	CpuCoreRealtimeDefault int32 `json:"cpuCoreRealtimeDefault"`
	CpuCoreNonRealtimeMax int32 `json:"cpuCoreNonRealtimeMax"`
	CpuCoreNonRealtimeMin int32 `json:"cpuCoreNonRealtimeMin"`
	CpuCoreNonRealtimeDefault int32 `json:"cpuCoreNonRealtimeDefault"`
	VariationId string `json:"variationId"`
	CpuCoreHwWdg uint32 `json:"cpuCoreHwWdg"`
	CpuCorePtpTimer uint32 `json:"cpuCorePtpTimer"`
	CpuCoreScheduler uint32 `json:"cpuCoreScheduler"`
	CpuCoreAutomation uint32 `json:"cpuCoreAutomation"`
}

func (t *CpuInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	cpuCoresTotalOffset := flatbuffers.UOffsetT(0)
	if t.CpuCoresTotal != nil {
		cpuCoresTotalLength := len(t.CpuCoresTotal)
		CpuInfoStartCpuCoresTotalVector(builder, cpuCoresTotalLength)
		for j := cpuCoresTotalLength - 1; j >= 0; j-- {
			builder.PrependUint32(t.CpuCoresTotal[j])
		}
		cpuCoresTotalOffset = builder.EndVector(cpuCoresTotalLength)
	}
	cpuCoresActiveOffset := flatbuffers.UOffsetT(0)
	if t.CpuCoresActive != nil {
		cpuCoresActiveLength := len(t.CpuCoresActive)
		CpuInfoStartCpuCoresActiveVector(builder, cpuCoresActiveLength)
		for j := cpuCoresActiveLength - 1; j >= 0; j-- {
			builder.PrependUint32(t.CpuCoresActive[j])
		}
		cpuCoresActiveOffset = builder.EndVector(cpuCoresActiveLength)
	}
	cpuCoresRealtimeOffset := flatbuffers.UOffsetT(0)
	if t.CpuCoresRealtime != nil {
		cpuCoresRealtimeLength := len(t.CpuCoresRealtime)
		CpuInfoStartCpuCoresRealtimeVector(builder, cpuCoresRealtimeLength)
		for j := cpuCoresRealtimeLength - 1; j >= 0; j-- {
			builder.PrependUint32(t.CpuCoresRealtime[j])
		}
		cpuCoresRealtimeOffset = builder.EndVector(cpuCoresRealtimeLength)
	}
	cpuCoresNonRealtimeOffset := flatbuffers.UOffsetT(0)
	if t.CpuCoresNonRealtime != nil {
		cpuCoresNonRealtimeLength := len(t.CpuCoresNonRealtime)
		CpuInfoStartCpuCoresNonRealtimeVector(builder, cpuCoresNonRealtimeLength)
		for j := cpuCoresNonRealtimeLength - 1; j >= 0; j-- {
			builder.PrependUint32(t.CpuCoresNonRealtime[j])
		}
		cpuCoresNonRealtimeOffset = builder.EndVector(cpuCoresNonRealtimeLength)
	}
	variationIdOffset := flatbuffers.UOffsetT(0)
	if t.VariationId != "" {
		variationIdOffset = builder.CreateString(t.VariationId)
	}
	CpuInfoStart(builder)
	CpuInfoAddCpuCoresTotal(builder, cpuCoresTotalOffset)
	CpuInfoAddCpuCoresActive(builder, cpuCoresActiveOffset)
	CpuInfoAddCpuCoresRealtime(builder, cpuCoresRealtimeOffset)
	CpuInfoAddCpuCoresNonRealtime(builder, cpuCoresNonRealtimeOffset)
	CpuInfoAddCpuCoreRealtimeMax(builder, t.CpuCoreRealtimeMax)
	CpuInfoAddCpuCoreRealtimeMin(builder, t.CpuCoreRealtimeMin)
	CpuInfoAddCpuCoreRealtimeDefault(builder, t.CpuCoreRealtimeDefault)
	CpuInfoAddCpuCoreNonRealtimeMax(builder, t.CpuCoreNonRealtimeMax)
	CpuInfoAddCpuCoreNonRealtimeMin(builder, t.CpuCoreNonRealtimeMin)
	CpuInfoAddCpuCoreNonRealtimeDefault(builder, t.CpuCoreNonRealtimeDefault)
	CpuInfoAddVariationId(builder, variationIdOffset)
	CpuInfoAddCpuCoreHwWdg(builder, t.CpuCoreHwWdg)
	CpuInfoAddCpuCorePtpTimer(builder, t.CpuCorePtpTimer)
	CpuInfoAddCpuCoreScheduler(builder, t.CpuCoreScheduler)
	CpuInfoAddCpuCoreAutomation(builder, t.CpuCoreAutomation)
	return CpuInfoEnd(builder)
}

func (rcv *CpuInfo) UnPackTo(t *CpuInfoT) {
	cpuCoresTotalLength := rcv.CpuCoresTotalLength()
	t.CpuCoresTotal = make([]uint32, cpuCoresTotalLength)
	for j := 0; j < cpuCoresTotalLength; j++ {
		t.CpuCoresTotal[j] = rcv.CpuCoresTotal(j)
	}
	cpuCoresActiveLength := rcv.CpuCoresActiveLength()
	t.CpuCoresActive = make([]uint32, cpuCoresActiveLength)
	for j := 0; j < cpuCoresActiveLength; j++ {
		t.CpuCoresActive[j] = rcv.CpuCoresActive(j)
	}
	cpuCoresRealtimeLength := rcv.CpuCoresRealtimeLength()
	t.CpuCoresRealtime = make([]uint32, cpuCoresRealtimeLength)
	for j := 0; j < cpuCoresRealtimeLength; j++ {
		t.CpuCoresRealtime[j] = rcv.CpuCoresRealtime(j)
	}
	cpuCoresNonRealtimeLength := rcv.CpuCoresNonRealtimeLength()
	t.CpuCoresNonRealtime = make([]uint32, cpuCoresNonRealtimeLength)
	for j := 0; j < cpuCoresNonRealtimeLength; j++ {
		t.CpuCoresNonRealtime[j] = rcv.CpuCoresNonRealtime(j)
	}
	t.CpuCoreRealtimeMax = rcv.CpuCoreRealtimeMax()
	t.CpuCoreRealtimeMin = rcv.CpuCoreRealtimeMin()
	t.CpuCoreRealtimeDefault = rcv.CpuCoreRealtimeDefault()
	t.CpuCoreNonRealtimeMax = rcv.CpuCoreNonRealtimeMax()
	t.CpuCoreNonRealtimeMin = rcv.CpuCoreNonRealtimeMin()
	t.CpuCoreNonRealtimeDefault = rcv.CpuCoreNonRealtimeDefault()
	t.VariationId = string(rcv.VariationId())
	t.CpuCoreHwWdg = rcv.CpuCoreHwWdg()
	t.CpuCorePtpTimer = rcv.CpuCorePtpTimer()
	t.CpuCoreScheduler = rcv.CpuCoreScheduler()
	t.CpuCoreAutomation = rcv.CpuCoreAutomation()
}

func (rcv *CpuInfo) UnPack() *CpuInfoT {
	if rcv == nil { return nil }
	t := &CpuInfoT{}
	rcv.UnPackTo(t)
	return t
}

type CpuInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsCpuInfo(buf []byte, offset flatbuffers.UOffsetT) *CpuInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CpuInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCpuInfo(buf []byte, offset flatbuffers.UOffsetT) *CpuInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CpuInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CpuInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CpuInfo) Table() flatbuffers.Table {
	return rcv._tab
}

/// Identification indices of all available CPU cores
func (rcv *CpuInfo) CpuCoresTotal(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *CpuInfo) CpuCoresTotalLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Identification indices of all available CPU cores
func (rcv *CpuInfo) MutateCpuCoresTotal(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

/// Identification indices of all available active CPU cores
func (rcv *CpuInfo) CpuCoresActive(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *CpuInfo) CpuCoresActiveLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Identification indices of all available active CPU cores
func (rcv *CpuInfo) MutateCpuCoresActive(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

/// Identification indices of the available CPU cores assigned to real-time processing
func (rcv *CpuInfo) CpuCoresRealtime(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *CpuInfo) CpuCoresRealtimeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Identification indices of the available CPU cores assigned to real-time processing
func (rcv *CpuInfo) MutateCpuCoresRealtime(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

/// Identification indices of the available CPU cores assigned to non real-time processing
func (rcv *CpuInfo) CpuCoresNonRealtime(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *CpuInfo) CpuCoresNonRealtimeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Identification indices of the available CPU cores assigned to non real-time processing
func (rcv *CpuInfo) MutateCpuCoresNonRealtime(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

/// Highest available identification index of the CPU core for real-time processing
func (rcv *CpuInfo) CpuCoreRealtimeMax() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return -1
}

/// Highest available identification index of the CPU core for real-time processing
func (rcv *CpuInfo) MutateCpuCoreRealtimeMax(n int32) bool {
	return rcv._tab.MutateInt32Slot(12, n)
}

/// Lowest available identification index of the CPU core for real-time processing
func (rcv *CpuInfo) CpuCoreRealtimeMin() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return -1
}

/// Lowest available identification index of the CPU core for real-time processing
func (rcv *CpuInfo) MutateCpuCoreRealtimeMin(n int32) bool {
	return rcv._tab.MutateInt32Slot(14, n)
}

/// Identification index of the CPU core assigned to real-time processing by default
func (rcv *CpuInfo) CpuCoreRealtimeDefault() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return -1
}

/// Identification index of the CPU core assigned to real-time processing by default
func (rcv *CpuInfo) MutateCpuCoreRealtimeDefault(n int32) bool {
	return rcv._tab.MutateInt32Slot(16, n)
}

/// Highest available identification index of the CPU core for non real-time processing
func (rcv *CpuInfo) CpuCoreNonRealtimeMax() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return -1
}

/// Highest available identification index of the CPU core for non real-time processing
func (rcv *CpuInfo) MutateCpuCoreNonRealtimeMax(n int32) bool {
	return rcv._tab.MutateInt32Slot(18, n)
}

/// Lowest available identification index of the CPU core for non real-time processing
func (rcv *CpuInfo) CpuCoreNonRealtimeMin() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return -1
}

/// Lowest available identification index of the CPU core for non real-time processing
func (rcv *CpuInfo) MutateCpuCoreNonRealtimeMin(n int32) bool {
	return rcv._tab.MutateInt32Slot(20, n)
}

/// Identification index of the CPU core assigned to non real-time processing by default
func (rcv *CpuInfo) CpuCoreNonRealtimeDefault() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return -1
}

/// Identification index of the CPU core assigned to non real-time processing by default
func (rcv *CpuInfo) MutateCpuCoreNonRealtimeDefault(n int32) bool {
	return rcv._tab.MutateInt32Slot(22, n)
}

/// Variation ID to manage special handling of CPU type
func (rcv *CpuInfo) VariationId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Variation ID to manage special handling of CPU type
/// Identification index of the CPU core on which the hardware watchdog is handled
func (rcv *CpuInfo) CpuCoreHwWdg() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// Identification index of the CPU core on which the hardware watchdog is handled
func (rcv *CpuInfo) MutateCpuCoreHwWdg(n uint32) bool {
	return rcv._tab.MutateUint32Slot(26, n)
}

/// Identification index of the CPU core on which the FPGA interrupt is handled
func (rcv *CpuInfo) CpuCorePtpTimer() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// Identification index of the CPU core on which the FPGA interrupt is handled
func (rcv *CpuInfo) MutateCpuCorePtpTimer(n uint32) bool {
	return rcv._tab.MutateUint32Slot(28, n)
}

/// Identification index of the CPU core on which the Scheduler tick task is executed
func (rcv *CpuInfo) CpuCoreScheduler() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// Identification index of the CPU core on which the Scheduler tick task is executed
func (rcv *CpuInfo) MutateCpuCoreScheduler(n uint32) bool {
	return rcv._tab.MutateUint32Slot(30, n)
}

/// Identification index of the CPU core on which the user task 'ctrlXAutomation' is executed by default
func (rcv *CpuInfo) CpuCoreAutomation() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

/// Identification index of the CPU core on which the user task 'ctrlXAutomation' is executed by default
func (rcv *CpuInfo) MutateCpuCoreAutomation(n uint32) bool {
	return rcv._tab.MutateUint32Slot(32, n)
}

func CpuInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(15)
}
func CpuInfoAddCpuCoresTotal(builder *flatbuffers.Builder, cpuCoresTotal flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(cpuCoresTotal), 0)
}
func CpuInfoStartCpuCoresTotalVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CpuInfoAddCpuCoresActive(builder *flatbuffers.Builder, cpuCoresActive flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(cpuCoresActive), 0)
}
func CpuInfoStartCpuCoresActiveVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CpuInfoAddCpuCoresRealtime(builder *flatbuffers.Builder, cpuCoresRealtime flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(cpuCoresRealtime), 0)
}
func CpuInfoStartCpuCoresRealtimeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CpuInfoAddCpuCoresNonRealtime(builder *flatbuffers.Builder, cpuCoresNonRealtime flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(cpuCoresNonRealtime), 0)
}
func CpuInfoStartCpuCoresNonRealtimeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CpuInfoAddCpuCoreRealtimeMax(builder *flatbuffers.Builder, cpuCoreRealtimeMax int32) {
	builder.PrependInt32Slot(4, cpuCoreRealtimeMax, -1)
}
func CpuInfoAddCpuCoreRealtimeMin(builder *flatbuffers.Builder, cpuCoreRealtimeMin int32) {
	builder.PrependInt32Slot(5, cpuCoreRealtimeMin, -1)
}
func CpuInfoAddCpuCoreRealtimeDefault(builder *flatbuffers.Builder, cpuCoreRealtimeDefault int32) {
	builder.PrependInt32Slot(6, cpuCoreRealtimeDefault, -1)
}
func CpuInfoAddCpuCoreNonRealtimeMax(builder *flatbuffers.Builder, cpuCoreNonRealtimeMax int32) {
	builder.PrependInt32Slot(7, cpuCoreNonRealtimeMax, -1)
}
func CpuInfoAddCpuCoreNonRealtimeMin(builder *flatbuffers.Builder, cpuCoreNonRealtimeMin int32) {
	builder.PrependInt32Slot(8, cpuCoreNonRealtimeMin, -1)
}
func CpuInfoAddCpuCoreNonRealtimeDefault(builder *flatbuffers.Builder, cpuCoreNonRealtimeDefault int32) {
	builder.PrependInt32Slot(9, cpuCoreNonRealtimeDefault, -1)
}
func CpuInfoAddVariationId(builder *flatbuffers.Builder, variationId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(variationId), 0)
}
func CpuInfoAddCpuCoreHwWdg(builder *flatbuffers.Builder, cpuCoreHwWdg uint32) {
	builder.PrependUint32Slot(11, cpuCoreHwWdg, 0)
}
func CpuInfoAddCpuCorePtpTimer(builder *flatbuffers.Builder, cpuCorePtpTimer uint32) {
	builder.PrependUint32Slot(12, cpuCorePtpTimer, 0)
}
func CpuInfoAddCpuCoreScheduler(builder *flatbuffers.Builder, cpuCoreScheduler uint32) {
	builder.PrependUint32Slot(13, cpuCoreScheduler, 0)
}
func CpuInfoAddCpuCoreAutomation(builder *flatbuffers.Builder, cpuCoreAutomation uint32) {
	builder.PrependUint32Slot(14, cpuCoreAutomation, 0)
}
func CpuInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
