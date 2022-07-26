// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SlaveConfigInfoResponseT struct {
	AutoIncAddr uint16
	EthercatAddr uint16
	IdentifyAdo uint16
	IdentifyValue uint16
	SlaveHandle uint32
	HcGroupIdx uint32
	PreviousEthercatAddr uint16
	PreviousPort uint16
	SlaveIdentity *EthercatIdentityInfoT
	SlaveName string
	MbxProtocols uint32
	MbxStandard *EthercatMailboxInfoT
	MbxBootstrap *EthercatMailboxInfoT
	ProcessDataIn []*EthercatMemoryInfoT
	ProcessDataOut []*EthercatMemoryInfoT
	NumProcessVarsIn uint16
	NumProcessVarsOut uint16
	PortDescriptor byte
	Reserved01 []byte
	WkcStateDiagOffsIn []uint16
	WkcStateDiagOffsOut []uint16
	Reserved02 []uint32
	IsPresent bool
	IsHcGroupPresent bool
	DcSupport bool
}

func (t *SlaveConfigInfoResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	slaveNameOffset := builder.CreateString(t.SlaveName)
	processDataInOffset := flatbuffers.UOffsetT(0)
	if t.ProcessDataIn != nil {
		processDataInLength := len(t.ProcessDataIn)
		SlaveConfigInfoResponseStartProcessDataInVector(builder, processDataInLength)
		for j := processDataInLength - 1; j >= 0; j-- {
			t.ProcessDataIn[j].Pack(builder)
		}
		processDataInOffset = builder.EndVector(processDataInLength)
	}
	processDataOutOffset := flatbuffers.UOffsetT(0)
	if t.ProcessDataOut != nil {
		processDataOutLength := len(t.ProcessDataOut)
		SlaveConfigInfoResponseStartProcessDataOutVector(builder, processDataOutLength)
		for j := processDataOutLength - 1; j >= 0; j-- {
			t.ProcessDataOut[j].Pack(builder)
		}
		processDataOutOffset = builder.EndVector(processDataOutLength)
	}
	reserved01Offset := flatbuffers.UOffsetT(0)
	if t.Reserved01 != nil {
		reserved01Offset = builder.CreateByteString(t.Reserved01)
	}
	wkcStateDiagOffsInOffset := flatbuffers.UOffsetT(0)
	if t.WkcStateDiagOffsIn != nil {
		wkcStateDiagOffsInLength := len(t.WkcStateDiagOffsIn)
		SlaveConfigInfoResponseStartWkcStateDiagOffsInVector(builder, wkcStateDiagOffsInLength)
		for j := wkcStateDiagOffsInLength - 1; j >= 0; j-- {
			builder.PrependUint16(t.WkcStateDiagOffsIn[j])
		}
		wkcStateDiagOffsInOffset = builder.EndVector(wkcStateDiagOffsInLength)
	}
	wkcStateDiagOffsOutOffset := flatbuffers.UOffsetT(0)
	if t.WkcStateDiagOffsOut != nil {
		wkcStateDiagOffsOutLength := len(t.WkcStateDiagOffsOut)
		SlaveConfigInfoResponseStartWkcStateDiagOffsOutVector(builder, wkcStateDiagOffsOutLength)
		for j := wkcStateDiagOffsOutLength - 1; j >= 0; j-- {
			builder.PrependUint16(t.WkcStateDiagOffsOut[j])
		}
		wkcStateDiagOffsOutOffset = builder.EndVector(wkcStateDiagOffsOutLength)
	}
	reserved02Offset := flatbuffers.UOffsetT(0)
	if t.Reserved02 != nil {
		reserved02Length := len(t.Reserved02)
		SlaveConfigInfoResponseStartReserved02Vector(builder, reserved02Length)
		for j := reserved02Length - 1; j >= 0; j-- {
			builder.PrependUint32(t.Reserved02[j])
		}
		reserved02Offset = builder.EndVector(reserved02Length)
	}
	SlaveConfigInfoResponseStart(builder)
	SlaveConfigInfoResponseAddAutoIncAddr(builder, t.AutoIncAddr)
	SlaveConfigInfoResponseAddEthercatAddr(builder, t.EthercatAddr)
	SlaveConfigInfoResponseAddIdentifyAdo(builder, t.IdentifyAdo)
	SlaveConfigInfoResponseAddIdentifyValue(builder, t.IdentifyValue)
	SlaveConfigInfoResponseAddSlaveHandle(builder, t.SlaveHandle)
	SlaveConfigInfoResponseAddHcGroupIdx(builder, t.HcGroupIdx)
	SlaveConfigInfoResponseAddPreviousEthercatAddr(builder, t.PreviousEthercatAddr)
	SlaveConfigInfoResponseAddPreviousPort(builder, t.PreviousPort)
	slaveIdentityOffset := t.SlaveIdentity.Pack(builder)
	SlaveConfigInfoResponseAddSlaveIdentity(builder, slaveIdentityOffset)
	SlaveConfigInfoResponseAddSlaveName(builder, slaveNameOffset)
	SlaveConfigInfoResponseAddMbxProtocols(builder, t.MbxProtocols)
	mbxStandardOffset := t.MbxStandard.Pack(builder)
	SlaveConfigInfoResponseAddMbxStandard(builder, mbxStandardOffset)
	mbxBootstrapOffset := t.MbxBootstrap.Pack(builder)
	SlaveConfigInfoResponseAddMbxBootstrap(builder, mbxBootstrapOffset)
	SlaveConfigInfoResponseAddProcessDataIn(builder, processDataInOffset)
	SlaveConfigInfoResponseAddProcessDataOut(builder, processDataOutOffset)
	SlaveConfigInfoResponseAddNumProcessVarsIn(builder, t.NumProcessVarsIn)
	SlaveConfigInfoResponseAddNumProcessVarsOut(builder, t.NumProcessVarsOut)
	SlaveConfigInfoResponseAddPortDescriptor(builder, t.PortDescriptor)
	SlaveConfigInfoResponseAddReserved01(builder, reserved01Offset)
	SlaveConfigInfoResponseAddWkcStateDiagOffsIn(builder, wkcStateDiagOffsInOffset)
	SlaveConfigInfoResponseAddWkcStateDiagOffsOut(builder, wkcStateDiagOffsOutOffset)
	SlaveConfigInfoResponseAddReserved02(builder, reserved02Offset)
	SlaveConfigInfoResponseAddIsPresent(builder, t.IsPresent)
	SlaveConfigInfoResponseAddIsHcGroupPresent(builder, t.IsHcGroupPresent)
	SlaveConfigInfoResponseAddDcSupport(builder, t.DcSupport)
	return SlaveConfigInfoResponseEnd(builder)
}

func (rcv *SlaveConfigInfoResponse) UnPackTo(t *SlaveConfigInfoResponseT) {
	t.AutoIncAddr = rcv.AutoIncAddr()
	t.EthercatAddr = rcv.EthercatAddr()
	t.IdentifyAdo = rcv.IdentifyAdo()
	t.IdentifyValue = rcv.IdentifyValue()
	t.SlaveHandle = rcv.SlaveHandle()
	t.HcGroupIdx = rcv.HcGroupIdx()
	t.PreviousEthercatAddr = rcv.PreviousEthercatAddr()
	t.PreviousPort = rcv.PreviousPort()
	t.SlaveIdentity = rcv.SlaveIdentity(nil).UnPack()
	t.SlaveName = string(rcv.SlaveName())
	t.MbxProtocols = rcv.MbxProtocols()
	t.MbxStandard = rcv.MbxStandard(nil).UnPack()
	t.MbxBootstrap = rcv.MbxBootstrap(nil).UnPack()
	processDataInLength := rcv.ProcessDataInLength()
	t.ProcessDataIn = make([]*EthercatMemoryInfoT, processDataInLength)
	for j := 0; j < processDataInLength; j++ {
		x := EthercatMemoryInfo{}
		rcv.ProcessDataIn(&x, j)
		t.ProcessDataIn[j] = x.UnPack()
	}
	processDataOutLength := rcv.ProcessDataOutLength()
	t.ProcessDataOut = make([]*EthercatMemoryInfoT, processDataOutLength)
	for j := 0; j < processDataOutLength; j++ {
		x := EthercatMemoryInfo{}
		rcv.ProcessDataOut(&x, j)
		t.ProcessDataOut[j] = x.UnPack()
	}
	t.NumProcessVarsIn = rcv.NumProcessVarsIn()
	t.NumProcessVarsOut = rcv.NumProcessVarsOut()
	t.PortDescriptor = rcv.PortDescriptor()
	t.Reserved01 = rcv.Reserved01Bytes()
	wkcStateDiagOffsInLength := rcv.WkcStateDiagOffsInLength()
	t.WkcStateDiagOffsIn = make([]uint16, wkcStateDiagOffsInLength)
	for j := 0; j < wkcStateDiagOffsInLength; j++ {
		t.WkcStateDiagOffsIn[j] = rcv.WkcStateDiagOffsIn(j)
	}
	wkcStateDiagOffsOutLength := rcv.WkcStateDiagOffsOutLength()
	t.WkcStateDiagOffsOut = make([]uint16, wkcStateDiagOffsOutLength)
	for j := 0; j < wkcStateDiagOffsOutLength; j++ {
		t.WkcStateDiagOffsOut[j] = rcv.WkcStateDiagOffsOut(j)
	}
	reserved02Length := rcv.Reserved02Length()
	t.Reserved02 = make([]uint32, reserved02Length)
	for j := 0; j < reserved02Length; j++ {
		t.Reserved02[j] = rcv.Reserved02(j)
	}
	t.IsPresent = rcv.IsPresent()
	t.IsHcGroupPresent = rcv.IsHcGroupPresent()
	t.DcSupport = rcv.DcSupport()
}

func (rcv *SlaveConfigInfoResponse) UnPack() *SlaveConfigInfoResponseT {
	if rcv == nil { return nil }
	t := &SlaveConfigInfoResponseT{}
	rcv.UnPackTo(t)
	return t
}

type SlaveConfigInfoResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsSlaveConfigInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *SlaveConfigInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SlaveConfigInfoResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSlaveConfigInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *SlaveConfigInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SlaveConfigInfoResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SlaveConfigInfoResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SlaveConfigInfoResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SlaveConfigInfoResponse) AutoIncAddr() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutateAutoIncAddr(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func (rcv *SlaveConfigInfoResponse) EthercatAddr() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutateEthercatAddr(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

func (rcv *SlaveConfigInfoResponse) IdentifyAdo() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutateIdentifyAdo(n uint16) bool {
	return rcv._tab.MutateUint16Slot(8, n)
}

func (rcv *SlaveConfigInfoResponse) IdentifyValue() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutateIdentifyValue(n uint16) bool {
	return rcv._tab.MutateUint16Slot(10, n)
}

func (rcv *SlaveConfigInfoResponse) SlaveHandle() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutateSlaveHandle(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

func (rcv *SlaveConfigInfoResponse) HcGroupIdx() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutateHcGroupIdx(n uint32) bool {
	return rcv._tab.MutateUint32Slot(14, n)
}

func (rcv *SlaveConfigInfoResponse) PreviousEthercatAddr() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutatePreviousEthercatAddr(n uint16) bool {
	return rcv._tab.MutateUint16Slot(16, n)
}

func (rcv *SlaveConfigInfoResponse) PreviousPort() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutatePreviousPort(n uint16) bool {
	return rcv._tab.MutateUint16Slot(18, n)
}

func (rcv *SlaveConfigInfoResponse) SlaveIdentity(obj *EthercatIdentityInfo) *EthercatIdentityInfo {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(EthercatIdentityInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SlaveConfigInfoResponse) SlaveName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *SlaveConfigInfoResponse) MbxProtocols() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutateMbxProtocols(n uint32) bool {
	return rcv._tab.MutateUint32Slot(24, n)
}

func (rcv *SlaveConfigInfoResponse) MbxStandard(obj *EthercatMailboxInfo) *EthercatMailboxInfo {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(EthercatMailboxInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SlaveConfigInfoResponse) MbxBootstrap(obj *EthercatMailboxInfo) *EthercatMailboxInfo {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(EthercatMailboxInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SlaveConfigInfoResponse) ProcessDataIn(obj *EthercatMemoryInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 8
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *SlaveConfigInfoResponse) ProcessDataInLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) ProcessDataOut(obj *EthercatMemoryInfo, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 8
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *SlaveConfigInfoResponse) ProcessDataOutLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) NumProcessVarsIn() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutateNumProcessVarsIn(n uint16) bool {
	return rcv._tab.MutateUint16Slot(34, n)
}

func (rcv *SlaveConfigInfoResponse) NumProcessVarsOut() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutateNumProcessVarsOut(n uint16) bool {
	return rcv._tab.MutateUint16Slot(36, n)
}

func (rcv *SlaveConfigInfoResponse) PortDescriptor() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutatePortDescriptor(n byte) bool {
	return rcv._tab.MutateByteSlot(38, n)
}

func (rcv *SlaveConfigInfoResponse) Reserved01(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) Reserved01Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) Reserved01Bytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *SlaveConfigInfoResponse) MutateReserved01(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *SlaveConfigInfoResponse) WkcStateDiagOffsIn(j int) uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint16(a + flatbuffers.UOffsetT(j*2))
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) WkcStateDiagOffsInLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutateWkcStateDiagOffsIn(j int, n uint16) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint16(a+flatbuffers.UOffsetT(j*2), n)
	}
	return false
}

func (rcv *SlaveConfigInfoResponse) WkcStateDiagOffsOut(j int) uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint16(a + flatbuffers.UOffsetT(j*2))
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) WkcStateDiagOffsOutLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutateWkcStateDiagOffsOut(j int, n uint16) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint16(a+flatbuffers.UOffsetT(j*2), n)
	}
	return false
}

func (rcv *SlaveConfigInfoResponse) Reserved02(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) Reserved02Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SlaveConfigInfoResponse) MutateReserved02(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *SlaveConfigInfoResponse) IsPresent() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *SlaveConfigInfoResponse) MutateIsPresent(n bool) bool {
	return rcv._tab.MutateBoolSlot(48, n)
}

func (rcv *SlaveConfigInfoResponse) IsHcGroupPresent() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *SlaveConfigInfoResponse) MutateIsHcGroupPresent(n bool) bool {
	return rcv._tab.MutateBoolSlot(50, n)
}

func (rcv *SlaveConfigInfoResponse) DcSupport() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *SlaveConfigInfoResponse) MutateDcSupport(n bool) bool {
	return rcv._tab.MutateBoolSlot(52, n)
}

func SlaveConfigInfoResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(25)
}
func SlaveConfigInfoResponseAddAutoIncAddr(builder *flatbuffers.Builder, autoIncAddr uint16) {
	builder.PrependUint16Slot(0, autoIncAddr, 0)
}
func SlaveConfigInfoResponseAddEthercatAddr(builder *flatbuffers.Builder, ethercatAddr uint16) {
	builder.PrependUint16Slot(1, ethercatAddr, 0)
}
func SlaveConfigInfoResponseAddIdentifyAdo(builder *flatbuffers.Builder, identifyAdo uint16) {
	builder.PrependUint16Slot(2, identifyAdo, 0)
}
func SlaveConfigInfoResponseAddIdentifyValue(builder *flatbuffers.Builder, identifyValue uint16) {
	builder.PrependUint16Slot(3, identifyValue, 0)
}
func SlaveConfigInfoResponseAddSlaveHandle(builder *flatbuffers.Builder, slaveHandle uint32) {
	builder.PrependUint32Slot(4, slaveHandle, 0)
}
func SlaveConfigInfoResponseAddHcGroupIdx(builder *flatbuffers.Builder, hcGroupIdx uint32) {
	builder.PrependUint32Slot(5, hcGroupIdx, 0)
}
func SlaveConfigInfoResponseAddPreviousEthercatAddr(builder *flatbuffers.Builder, previousEthercatAddr uint16) {
	builder.PrependUint16Slot(6, previousEthercatAddr, 0)
}
func SlaveConfigInfoResponseAddPreviousPort(builder *flatbuffers.Builder, previousPort uint16) {
	builder.PrependUint16Slot(7, previousPort, 0)
}
func SlaveConfigInfoResponseAddSlaveIdentity(builder *flatbuffers.Builder, slaveIdentity flatbuffers.UOffsetT) {
	builder.PrependStructSlot(8, flatbuffers.UOffsetT(slaveIdentity), 0)
}
func SlaveConfigInfoResponseAddSlaveName(builder *flatbuffers.Builder, slaveName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(slaveName), 0)
}
func SlaveConfigInfoResponseAddMbxProtocols(builder *flatbuffers.Builder, mbxProtocols uint32) {
	builder.PrependUint32Slot(10, mbxProtocols, 0)
}
func SlaveConfigInfoResponseAddMbxStandard(builder *flatbuffers.Builder, mbxStandard flatbuffers.UOffsetT) {
	builder.PrependStructSlot(11, flatbuffers.UOffsetT(mbxStandard), 0)
}
func SlaveConfigInfoResponseAddMbxBootstrap(builder *flatbuffers.Builder, mbxBootstrap flatbuffers.UOffsetT) {
	builder.PrependStructSlot(12, flatbuffers.UOffsetT(mbxBootstrap), 0)
}
func SlaveConfigInfoResponseAddProcessDataIn(builder *flatbuffers.Builder, processDataIn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(processDataIn), 0)
}
func SlaveConfigInfoResponseStartProcessDataInVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 4)
}
func SlaveConfigInfoResponseAddProcessDataOut(builder *flatbuffers.Builder, processDataOut flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(14, flatbuffers.UOffsetT(processDataOut), 0)
}
func SlaveConfigInfoResponseStartProcessDataOutVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 4)
}
func SlaveConfigInfoResponseAddNumProcessVarsIn(builder *flatbuffers.Builder, numProcessVarsIn uint16) {
	builder.PrependUint16Slot(15, numProcessVarsIn, 0)
}
func SlaveConfigInfoResponseAddNumProcessVarsOut(builder *flatbuffers.Builder, numProcessVarsOut uint16) {
	builder.PrependUint16Slot(16, numProcessVarsOut, 0)
}
func SlaveConfigInfoResponseAddPortDescriptor(builder *flatbuffers.Builder, portDescriptor byte) {
	builder.PrependByteSlot(17, portDescriptor, 0)
}
func SlaveConfigInfoResponseAddReserved01(builder *flatbuffers.Builder, reserved01 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(18, flatbuffers.UOffsetT(reserved01), 0)
}
func SlaveConfigInfoResponseStartReserved01Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func SlaveConfigInfoResponseAddWkcStateDiagOffsIn(builder *flatbuffers.Builder, wkcStateDiagOffsIn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(19, flatbuffers.UOffsetT(wkcStateDiagOffsIn), 0)
}
func SlaveConfigInfoResponseStartWkcStateDiagOffsInVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(2, numElems, 2)
}
func SlaveConfigInfoResponseAddWkcStateDiagOffsOut(builder *flatbuffers.Builder, wkcStateDiagOffsOut flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(20, flatbuffers.UOffsetT(wkcStateDiagOffsOut), 0)
}
func SlaveConfigInfoResponseStartWkcStateDiagOffsOutVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(2, numElems, 2)
}
func SlaveConfigInfoResponseAddReserved02(builder *flatbuffers.Builder, reserved02 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(21, flatbuffers.UOffsetT(reserved02), 0)
}
func SlaveConfigInfoResponseStartReserved02Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SlaveConfigInfoResponseAddIsPresent(builder *flatbuffers.Builder, isPresent bool) {
	builder.PrependBoolSlot(22, isPresent, false)
}
func SlaveConfigInfoResponseAddIsHcGroupPresent(builder *flatbuffers.Builder, isHcGroupPresent bool) {
	builder.PrependBoolSlot(23, isHcGroupPresent, false)
}
func SlaveConfigInfoResponseAddDcSupport(builder *flatbuffers.Builder, dcSupport bool) {
	builder.PrependBoolSlot(24, dcSupport, false)
}
func SlaveConfigInfoResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
