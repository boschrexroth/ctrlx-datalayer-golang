// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

///Slave online information data response
type SlaveOnlineInfoResponseT struct {
	AutoIncAddr uint16 `json:"autoIncAddr"`
	EthercatAddr uint16 `json:"ethercatAddr"`
	StationAlias uint16 `json:"stationAlias"`
	IdentifyValue uint16 `json:"identifyValue"`
	SlaveHandle uint32 `json:"slaveHandle"`
	PortSlaveHandles []uint32 `json:"portSlaveHandles"`
	SlaveIdentity *EthercatIdentityInfoT `json:"slaveIdentity"`
	EscType byte `json:"escType"`
	EscRevision byte `json:"escRevision"`
	EscBuild uint16 `json:"escBuild"`
	EscFeatures uint16 `json:"escFeatures"`
	PortDescriptor byte `json:"portDescriptor"`
	Reserved01 byte `json:"reserved01"`
	AlStatus uint16 `json:"alStatus"`
	AlStatusCode uint16 `json:"alStatusCode"`
	MbxProtocols uint16 `json:"mbxProtocols"`
	DlStatus uint16 `json:"dlStatus"`
	PortState uint16 `json:"portState"`
	PreviousPort uint16 `json:"previousPort"`
	SystemTimeDifference uint32 `json:"systemTimeDifference"`
	SlaveDelay uint32 `json:"slaveDelay"`
	PropagationDelay uint32 `json:"propagationDelay"`
	Reserved02 []uint32 `json:"reserved02"`
	DcSupport bool `json:"dcSupport"`
	Dc64Support bool `json:"dc64Support"`
	IsRefClock bool `json:"isRefClock"`
	LineCrossed bool `json:"lineCrossed"`
	LineCrossedFlags uint32 `json:"lineCrossedFlags"`
	CyclicWkcErrorCnt uint32 `json:"cyclicWkcErrorCnt"`
	SlaveAbsentCnt uint32 `json:"slaveAbsentCnt"`
	UnexpectedAbsentCnt uint32 `json:"unexpectedAbsentCnt"`
	IsDeviceEmulation bool `json:"isDeviceEmulation"`
}

func (t *SlaveOnlineInfoResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	portSlaveHandlesOffset := flatbuffers.UOffsetT(0)
	if t.PortSlaveHandles != nil {
		portSlaveHandlesLength := len(t.PortSlaveHandles)
		SlaveOnlineInfoResponseStartPortSlaveHandlesVector(builder, portSlaveHandlesLength)
		for j := portSlaveHandlesLength - 1; j >= 0; j-- {
			builder.PrependUint32(t.PortSlaveHandles[j])
		}
		portSlaveHandlesOffset = builder.EndVector(portSlaveHandlesLength)
	}
	reserved02Offset := flatbuffers.UOffsetT(0)
	if t.Reserved02 != nil {
		reserved02Length := len(t.Reserved02)
		SlaveOnlineInfoResponseStartReserved02Vector(builder, reserved02Length)
		for j := reserved02Length - 1; j >= 0; j-- {
			builder.PrependUint32(t.Reserved02[j])
		}
		reserved02Offset = builder.EndVector(reserved02Length)
	}
	SlaveOnlineInfoResponseStart(builder)
	SlaveOnlineInfoResponseAddAutoIncAddr(builder, t.AutoIncAddr)
	SlaveOnlineInfoResponseAddEthercatAddr(builder, t.EthercatAddr)
	SlaveOnlineInfoResponseAddStationAlias(builder, t.StationAlias)
	SlaveOnlineInfoResponseAddIdentifyValue(builder, t.IdentifyValue)
	SlaveOnlineInfoResponseAddSlaveHandle(builder, t.SlaveHandle)
	SlaveOnlineInfoResponseAddPortSlaveHandles(builder, portSlaveHandlesOffset)
	slaveIdentityOffset := t.SlaveIdentity.Pack(builder)
	SlaveOnlineInfoResponseAddSlaveIdentity(builder, slaveIdentityOffset)
	SlaveOnlineInfoResponseAddEscType(builder, t.EscType)
	SlaveOnlineInfoResponseAddEscRevision(builder, t.EscRevision)
	SlaveOnlineInfoResponseAddEscBuild(builder, t.EscBuild)
	SlaveOnlineInfoResponseAddEscFeatures(builder, t.EscFeatures)
	SlaveOnlineInfoResponseAddPortDescriptor(builder, t.PortDescriptor)
	SlaveOnlineInfoResponseAddReserved01(builder, t.Reserved01)
	SlaveOnlineInfoResponseAddAlStatus(builder, t.AlStatus)
	SlaveOnlineInfoResponseAddAlStatusCode(builder, t.AlStatusCode)
	SlaveOnlineInfoResponseAddMbxProtocols(builder, t.MbxProtocols)
	SlaveOnlineInfoResponseAddDlStatus(builder, t.DlStatus)
	SlaveOnlineInfoResponseAddPortState(builder, t.PortState)
	SlaveOnlineInfoResponseAddPreviousPort(builder, t.PreviousPort)
	SlaveOnlineInfoResponseAddSystemTimeDifference(builder, t.SystemTimeDifference)
	SlaveOnlineInfoResponseAddSlaveDelay(builder, t.SlaveDelay)
	SlaveOnlineInfoResponseAddPropagationDelay(builder, t.PropagationDelay)
	SlaveOnlineInfoResponseAddReserved02(builder, reserved02Offset)
	SlaveOnlineInfoResponseAddDcSupport(builder, t.DcSupport)
	SlaveOnlineInfoResponseAddDc64Support(builder, t.Dc64Support)
	SlaveOnlineInfoResponseAddIsRefClock(builder, t.IsRefClock)
	SlaveOnlineInfoResponseAddLineCrossed(builder, t.LineCrossed)
	SlaveOnlineInfoResponseAddLineCrossedFlags(builder, t.LineCrossedFlags)
	SlaveOnlineInfoResponseAddCyclicWkcErrorCnt(builder, t.CyclicWkcErrorCnt)
	SlaveOnlineInfoResponseAddSlaveAbsentCnt(builder, t.SlaveAbsentCnt)
	SlaveOnlineInfoResponseAddUnexpectedAbsentCnt(builder, t.UnexpectedAbsentCnt)
	SlaveOnlineInfoResponseAddIsDeviceEmulation(builder, t.IsDeviceEmulation)
	return SlaveOnlineInfoResponseEnd(builder)
}

func (rcv *SlaveOnlineInfoResponse) UnPackTo(t *SlaveOnlineInfoResponseT) {
	t.AutoIncAddr = rcv.AutoIncAddr()
	t.EthercatAddr = rcv.EthercatAddr()
	t.StationAlias = rcv.StationAlias()
	t.IdentifyValue = rcv.IdentifyValue()
	t.SlaveHandle = rcv.SlaveHandle()
	portSlaveHandlesLength := rcv.PortSlaveHandlesLength()
	t.PortSlaveHandles = make([]uint32, portSlaveHandlesLength)
	for j := 0; j < portSlaveHandlesLength; j++ {
		t.PortSlaveHandles[j] = rcv.PortSlaveHandles(j)
	}
	t.SlaveIdentity = rcv.SlaveIdentity(nil).UnPack()
	t.EscType = rcv.EscType()
	t.EscRevision = rcv.EscRevision()
	t.EscBuild = rcv.EscBuild()
	t.EscFeatures = rcv.EscFeatures()
	t.PortDescriptor = rcv.PortDescriptor()
	t.Reserved01 = rcv.Reserved01()
	t.AlStatus = rcv.AlStatus()
	t.AlStatusCode = rcv.AlStatusCode()
	t.MbxProtocols = rcv.MbxProtocols()
	t.DlStatus = rcv.DlStatus()
	t.PortState = rcv.PortState()
	t.PreviousPort = rcv.PreviousPort()
	t.SystemTimeDifference = rcv.SystemTimeDifference()
	t.SlaveDelay = rcv.SlaveDelay()
	t.PropagationDelay = rcv.PropagationDelay()
	reserved02Length := rcv.Reserved02Length()
	t.Reserved02 = make([]uint32, reserved02Length)
	for j := 0; j < reserved02Length; j++ {
		t.Reserved02[j] = rcv.Reserved02(j)
	}
	t.DcSupport = rcv.DcSupport()
	t.Dc64Support = rcv.Dc64Support()
	t.IsRefClock = rcv.IsRefClock()
	t.LineCrossed = rcv.LineCrossed()
	t.LineCrossedFlags = rcv.LineCrossedFlags()
	t.CyclicWkcErrorCnt = rcv.CyclicWkcErrorCnt()
	t.SlaveAbsentCnt = rcv.SlaveAbsentCnt()
	t.UnexpectedAbsentCnt = rcv.UnexpectedAbsentCnt()
	t.IsDeviceEmulation = rcv.IsDeviceEmulation()
}

func (rcv *SlaveOnlineInfoResponse) UnPack() *SlaveOnlineInfoResponseT {
	if rcv == nil { return nil }
	t := &SlaveOnlineInfoResponseT{}
	rcv.UnPackTo(t)
	return t
}

type SlaveOnlineInfoResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsSlaveOnlineInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *SlaveOnlineInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SlaveOnlineInfoResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSlaveOnlineInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *SlaveOnlineInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SlaveOnlineInfoResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SlaveOnlineInfoResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SlaveOnlineInfoResponse) Table() flatbuffers.Table {
	return rcv._tab
}

///Auto increment address
func (rcv *SlaveOnlineInfoResponse) AutoIncAddr() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Auto increment address
func (rcv *SlaveOnlineInfoResponse) MutateAutoIncAddr(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

///EtherCAT address, fixed physical (ESC register 0x0010)
func (rcv *SlaveOnlineInfoResponse) EthercatAddr() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///EtherCAT address, fixed physical (ESC register 0x0010)
func (rcv *SlaveOnlineInfoResponse) MutateEthercatAddr(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

///Station alias (second slave address) (ESC register 0x0012) 
func (rcv *SlaveOnlineInfoResponse) StationAlias() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Station alias (second slave address) (ESC register 0x0012) 
func (rcv *SlaveOnlineInfoResponse) MutateStationAlias(n uint16) bool {
	return rcv._tab.MutateUint16Slot(8, n)
}

///Last read identification value (only if slave is configured with an ‘IdentifyAdo’)
func (rcv *SlaveOnlineInfoResponse) IdentifyValue() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Last read identification value (only if slave is configured with an ‘IdentifyAdo’)
func (rcv *SlaveOnlineInfoResponse) MutateIdentifyValue(n uint16) bool {
	return rcv._tab.MutateUint16Slot(10, n)
}

///Internal slave id (internal use)
func (rcv *SlaveOnlineInfoResponse) SlaveHandle() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Internal slave id (internal use)
func (rcv *SlaveOnlineInfoResponse) MutateSlaveHandle(n uint32) bool {
	return rcv._tab.MutateUint32Slot(12, n)
}

///Link to connected slaves (via SlaveHandle)
func (rcv *SlaveOnlineInfoResponse) PortSlaveHandles(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *SlaveOnlineInfoResponse) PortSlaveHandlesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

///Link to connected slaves (via SlaveHandle)
func (rcv *SlaveOnlineInfoResponse) MutatePortSlaveHandles(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

///Slave identity info (EEPROM Offset 0x0008)
func (rcv *SlaveOnlineInfoResponse) SlaveIdentity(obj *EthercatIdentityInfo) *EthercatIdentityInfo {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
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

///Slave identity info (EEPROM Offset 0x0008)
///EtherCAT slave controller (ESC) type (ESC register 0x0000)
func (rcv *SlaveOnlineInfoResponse) EscType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

///EtherCAT slave controller (ESC) type (ESC register 0x0000)
func (rcv *SlaveOnlineInfoResponse) MutateEscType(n byte) bool {
	return rcv._tab.MutateByteSlot(18, n)
}

///EtherCAT slave controller (ESC) revision (ESC register 0x0001)
func (rcv *SlaveOnlineInfoResponse) EscRevision() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

///EtherCAT slave controller (ESC) revision (ESC register 0x0001)
func (rcv *SlaveOnlineInfoResponse) MutateEscRevision(n byte) bool {
	return rcv._tab.MutateByteSlot(20, n)
}

///EtherCAT slave controller (ESC) build (ESC register 0x0002)
func (rcv *SlaveOnlineInfoResponse) EscBuild() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///EtherCAT slave controller (ESC) build (ESC register 0x0002)
func (rcv *SlaveOnlineInfoResponse) MutateEscBuild(n uint16) bool {
	return rcv._tab.MutateUint16Slot(22, n)
}

///Supported EtherCAT slave controller (ESC) Features (ESC register 0x0008)
func (rcv *SlaveOnlineInfoResponse) EscFeatures() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Supported EtherCAT slave controller (ESC) Features (ESC register 0x0008)
func (rcv *SlaveOnlineInfoResponse) MutateEscFeatures(n uint16) bool {
	return rcv._tab.MutateUint16Slot(24, n)
}

///Port descriptor (ESC register 0x0007) (0b10 = EBUS, 0b11 = MII/..) (Port0^Bit0:1, Port1^Bit2:3, Port2^Bit4:5, Port3^Bit6:7)
func (rcv *SlaveOnlineInfoResponse) PortDescriptor() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

///Port descriptor (ESC register 0x0007) (0b10 = EBUS, 0b11 = MII/..) (Port0^Bit0:1, Port1^Bit2:3, Port2^Bit4:5, Port3^Bit6:7)
func (rcv *SlaveOnlineInfoResponse) MutatePortDescriptor(n byte) bool {
	return rcv._tab.MutateByteSlot(26, n)
}

///Reserved
func (rcv *SlaveOnlineInfoResponse) Reserved01() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

///Reserved
func (rcv *SlaveOnlineInfoResponse) MutateReserved01(n byte) bool {
	return rcv._tab.MutateByteSlot(28, n)
}

///Application layer status (ESC register 0x0130)
func (rcv *SlaveOnlineInfoResponse) AlStatus() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Application layer status (ESC register 0x0130)
func (rcv *SlaveOnlineInfoResponse) MutateAlStatus(n uint16) bool {
	return rcv._tab.MutateUint16Slot(30, n)
}

///Application layer status Code (ESC register 0x0134)
func (rcv *SlaveOnlineInfoResponse) AlStatusCode() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Application layer status Code (ESC register 0x0134)
func (rcv *SlaveOnlineInfoResponse) MutateAlStatusCode(n uint16) bool {
	return rcv._tab.MutateUint16Slot(32, n)
}

///Mailbox supported Protocols (Online data)
///Bit 0: AoE (ADS over EtherCAT) (ADS: Automation Device Specification)
///Bit 1: EoE (Ethernet over EtherCAT)
///Bit 2: CoE (CAN application protocol over EtherCAT)
///Bit 3: FoE (File access over EtherCAT)
///Bit 4: SoE (Servo drive over EtherCAT)
///Bit 5: VoE (Vendor specific protocol over EtherCAT)
///Bit 6 to 15: Reserved
func (rcv *SlaveOnlineInfoResponse) MbxProtocols() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Mailbox supported Protocols (Online data)
///Bit 0: AoE (ADS over EtherCAT) (ADS: Automation Device Specification)
///Bit 1: EoE (Ethernet over EtherCAT)
///Bit 2: CoE (CAN application protocol over EtherCAT)
///Bit 3: FoE (File access over EtherCAT)
///Bit 4: SoE (Servo drive over EtherCAT)
///Bit 5: VoE (Vendor specific protocol over EtherCAT)
///Bit 6 to 15: Reserved
func (rcv *SlaveOnlineInfoResponse) MutateMbxProtocols(n uint16) bool {
	return rcv._tab.MutateUint16Slot(34, n)
}

///Data link status (ESC register 0x0110:0x0111)
func (rcv *SlaveOnlineInfoResponse) DlStatus() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Data link status (ESC register 0x0110:0x0111)
func (rcv *SlaveOnlineInfoResponse) MutateDlStatus(n uint16) bool {
	return rcv._tab.MutateUint16Slot(36, n)
}

///In comparison to the DL Status the ‘Port-State’ represents only the physical ports. 
///The "ConnectionPortX" diagnosis bit is a logical result of the other status bits.
func (rcv *SlaveOnlineInfoResponse) PortState() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///In comparison to the DL Status the ‘Port-State’ represents only the physical ports. 
///The "ConnectionPortX" diagnosis bit is a logical result of the other status bits.
func (rcv *SlaveOnlineInfoResponse) MutatePortState(n uint16) bool {
	return rcv._tab.MutateUint16Slot(38, n)
}

///Connected Port of the previous slave
func (rcv *SlaveOnlineInfoResponse) PreviousPort() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

///Connected Port of the previous slave
func (rcv *SlaveOnlineInfoResponse) MutatePreviousPort(n uint16) bool {
	return rcv._tab.MutateUint16Slot(40, n)
}

///System time difference in nanoseconds (only if DC is configured) (ESC register 0x092C)
func (rcv *SlaveOnlineInfoResponse) SystemTimeDifference() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///System time difference in nanoseconds (only if DC is configured) (ESC register 0x092C)
func (rcv *SlaveOnlineInfoResponse) MutateSystemTimeDifference(n uint32) bool {
	return rcv._tab.MutateUint32Slot(42, n)
}

///Time delay "behind" slave in nanoseconds (only if DC is configured)
func (rcv *SlaveOnlineInfoResponse) SlaveDelay() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Time delay "behind" slave in nanoseconds (only if DC is configured)
func (rcv *SlaveOnlineInfoResponse) MutateSlaveDelay(n uint32) bool {
	return rcv._tab.MutateUint32Slot(44, n)
}

///Propagation delay in nanoseconds (only if DC is configured) (ESC register 0x0928)
func (rcv *SlaveOnlineInfoResponse) PropagationDelay() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Propagation delay in nanoseconds (only if DC is configured) (ESC register 0x0928)
func (rcv *SlaveOnlineInfoResponse) MutatePropagationDelay(n uint32) bool {
	return rcv._tab.MutateUint32Slot(46, n)
}

///Reserved
func (rcv *SlaveOnlineInfoResponse) Reserved02(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *SlaveOnlineInfoResponse) Reserved02Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

///Reserved
func (rcv *SlaveOnlineInfoResponse) MutateReserved02(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

///Slave supports DC (Distributed Clock)
func (rcv *SlaveOnlineInfoResponse) DcSupport() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

///Slave supports DC (Distributed Clock)
func (rcv *SlaveOnlineInfoResponse) MutateDcSupport(n bool) bool {
	return rcv._tab.MutateBoolSlot(50, n)
}

func (rcv *SlaveOnlineInfoResponse) Dc64Support() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *SlaveOnlineInfoResponse) MutateDc64Support(n bool) bool {
	return rcv._tab.MutateBoolSlot(52, n)
}

///Slave is reference clock
func (rcv *SlaveOnlineInfoResponse) IsRefClock() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

///Slave is reference clock
func (rcv *SlaveOnlineInfoResponse) MutateIsRefClock(n bool) bool {
	return rcv._tab.MutateBoolSlot(54, n)
}

///Line crossed detected at this slave (e.g. IN-OUT Port interchanged)
func (rcv *SlaveOnlineInfoResponse) LineCrossed() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

///Line crossed detected at this slave (e.g. IN-OUT Port interchanged)
func (rcv *SlaveOnlineInfoResponse) MutateLineCrossed(n bool) bool {
	return rcv._tab.MutateBoolSlot(56, n)
}

///Line crossed flags indicate in which way(s) lines have been crossed
///Bit 0:Not connected port A
///Bit 1:Unexpected input port
///Bit 2:Unexpected junction port
///Bit 3:Unresolved port connection
///Bit 4:Hidden slave connected
///Bit 5:Physic mismatch
///Bit 6:Invalid port connection
///Bit 7 to 31:Reserved   
func (rcv *SlaveOnlineInfoResponse) LineCrossedFlags() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Line crossed flags indicate in which way(s) lines have been crossed
///Bit 0:Not connected port A
///Bit 1:Unexpected input port
///Bit 2:Unexpected junction port
///Bit 3:Unresolved port connection
///Bit 4:Hidden slave connected
///Bit 5:Physic mismatch
///Bit 6:Invalid port connection
///Bit 7 to 31:Reserved   
func (rcv *SlaveOnlineInfoResponse) MutateLineCrossedFlags(n uint32) bool {
	return rcv._tab.MutateUint32Slot(58, n)
}

///Counter for Cyclic WKC Errors (over all pd sections)
func (rcv *SlaveOnlineInfoResponse) CyclicWkcErrorCnt() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Counter for Cyclic WKC Errors (over all pd sections)
func (rcv *SlaveOnlineInfoResponse) MutateCyclicWkcErrorCnt(n uint32) bool {
	return rcv._tab.MutateUint32Slot(60, n)
}

///Counter for Absent/Not Present Slaves
func (rcv *SlaveOnlineInfoResponse) SlaveAbsentCnt() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(62))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Counter for Absent/Not Present Slaves
func (rcv *SlaveOnlineInfoResponse) MutateSlaveAbsentCnt(n uint32) bool {
	return rcv._tab.MutateUint32Slot(62, n)
}

///Counter for Abnormal State Change
func (rcv *SlaveOnlineInfoResponse) UnexpectedAbsentCnt() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(64))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

///Counter for Abnormal State Change
func (rcv *SlaveOnlineInfoResponse) MutateUnexpectedAbsentCnt(n uint32) bool {
	return rcv._tab.MutateUint32Slot(64, n)
}

///Slave without Firmware. ESC register 0x0141, enabled by EEPROM offset 0x0000.8
func (rcv *SlaveOnlineInfoResponse) IsDeviceEmulation() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(66))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

///Slave without Firmware. ESC register 0x0141, enabled by EEPROM offset 0x0000.8
func (rcv *SlaveOnlineInfoResponse) MutateIsDeviceEmulation(n bool) bool {
	return rcv._tab.MutateBoolSlot(66, n)
}

func SlaveOnlineInfoResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(32)
}
func SlaveOnlineInfoResponseAddAutoIncAddr(builder *flatbuffers.Builder, autoIncAddr uint16) {
	builder.PrependUint16Slot(0, autoIncAddr, 0)
}
func SlaveOnlineInfoResponseAddEthercatAddr(builder *flatbuffers.Builder, ethercatAddr uint16) {
	builder.PrependUint16Slot(1, ethercatAddr, 0)
}
func SlaveOnlineInfoResponseAddStationAlias(builder *flatbuffers.Builder, stationAlias uint16) {
	builder.PrependUint16Slot(2, stationAlias, 0)
}
func SlaveOnlineInfoResponseAddIdentifyValue(builder *flatbuffers.Builder, identifyValue uint16) {
	builder.PrependUint16Slot(3, identifyValue, 0)
}
func SlaveOnlineInfoResponseAddSlaveHandle(builder *flatbuffers.Builder, slaveHandle uint32) {
	builder.PrependUint32Slot(4, slaveHandle, 0)
}
func SlaveOnlineInfoResponseAddPortSlaveHandles(builder *flatbuffers.Builder, portSlaveHandles flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(portSlaveHandles), 0)
}
func SlaveOnlineInfoResponseStartPortSlaveHandlesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SlaveOnlineInfoResponseAddSlaveIdentity(builder *flatbuffers.Builder, slaveIdentity flatbuffers.UOffsetT) {
	builder.PrependStructSlot(6, flatbuffers.UOffsetT(slaveIdentity), 0)
}
func SlaveOnlineInfoResponseAddEscType(builder *flatbuffers.Builder, escType byte) {
	builder.PrependByteSlot(7, escType, 0)
}
func SlaveOnlineInfoResponseAddEscRevision(builder *flatbuffers.Builder, escRevision byte) {
	builder.PrependByteSlot(8, escRevision, 0)
}
func SlaveOnlineInfoResponseAddEscBuild(builder *flatbuffers.Builder, escBuild uint16) {
	builder.PrependUint16Slot(9, escBuild, 0)
}
func SlaveOnlineInfoResponseAddEscFeatures(builder *flatbuffers.Builder, escFeatures uint16) {
	builder.PrependUint16Slot(10, escFeatures, 0)
}
func SlaveOnlineInfoResponseAddPortDescriptor(builder *flatbuffers.Builder, portDescriptor byte) {
	builder.PrependByteSlot(11, portDescriptor, 0)
}
func SlaveOnlineInfoResponseAddReserved01(builder *flatbuffers.Builder, reserved01 byte) {
	builder.PrependByteSlot(12, reserved01, 0)
}
func SlaveOnlineInfoResponseAddAlStatus(builder *flatbuffers.Builder, alStatus uint16) {
	builder.PrependUint16Slot(13, alStatus, 0)
}
func SlaveOnlineInfoResponseAddAlStatusCode(builder *flatbuffers.Builder, alStatusCode uint16) {
	builder.PrependUint16Slot(14, alStatusCode, 0)
}
func SlaveOnlineInfoResponseAddMbxProtocols(builder *flatbuffers.Builder, mbxProtocols uint16) {
	builder.PrependUint16Slot(15, mbxProtocols, 0)
}
func SlaveOnlineInfoResponseAddDlStatus(builder *flatbuffers.Builder, dlStatus uint16) {
	builder.PrependUint16Slot(16, dlStatus, 0)
}
func SlaveOnlineInfoResponseAddPortState(builder *flatbuffers.Builder, portState uint16) {
	builder.PrependUint16Slot(17, portState, 0)
}
func SlaveOnlineInfoResponseAddPreviousPort(builder *flatbuffers.Builder, previousPort uint16) {
	builder.PrependUint16Slot(18, previousPort, 0)
}
func SlaveOnlineInfoResponseAddSystemTimeDifference(builder *flatbuffers.Builder, systemTimeDifference uint32) {
	builder.PrependUint32Slot(19, systemTimeDifference, 0)
}
func SlaveOnlineInfoResponseAddSlaveDelay(builder *flatbuffers.Builder, slaveDelay uint32) {
	builder.PrependUint32Slot(20, slaveDelay, 0)
}
func SlaveOnlineInfoResponseAddPropagationDelay(builder *flatbuffers.Builder, propagationDelay uint32) {
	builder.PrependUint32Slot(21, propagationDelay, 0)
}
func SlaveOnlineInfoResponseAddReserved02(builder *flatbuffers.Builder, reserved02 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(22, flatbuffers.UOffsetT(reserved02), 0)
}
func SlaveOnlineInfoResponseStartReserved02Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SlaveOnlineInfoResponseAddDcSupport(builder *flatbuffers.Builder, dcSupport bool) {
	builder.PrependBoolSlot(23, dcSupport, false)
}
func SlaveOnlineInfoResponseAddDc64Support(builder *flatbuffers.Builder, dc64Support bool) {
	builder.PrependBoolSlot(24, dc64Support, false)
}
func SlaveOnlineInfoResponseAddIsRefClock(builder *flatbuffers.Builder, isRefClock bool) {
	builder.PrependBoolSlot(25, isRefClock, false)
}
func SlaveOnlineInfoResponseAddLineCrossed(builder *flatbuffers.Builder, lineCrossed bool) {
	builder.PrependBoolSlot(26, lineCrossed, false)
}
func SlaveOnlineInfoResponseAddLineCrossedFlags(builder *flatbuffers.Builder, lineCrossedFlags uint32) {
	builder.PrependUint32Slot(27, lineCrossedFlags, 0)
}
func SlaveOnlineInfoResponseAddCyclicWkcErrorCnt(builder *flatbuffers.Builder, cyclicWkcErrorCnt uint32) {
	builder.PrependUint32Slot(28, cyclicWkcErrorCnt, 0)
}
func SlaveOnlineInfoResponseAddSlaveAbsentCnt(builder *flatbuffers.Builder, slaveAbsentCnt uint32) {
	builder.PrependUint32Slot(29, slaveAbsentCnt, 0)
}
func SlaveOnlineInfoResponseAddUnexpectedAbsentCnt(builder *flatbuffers.Builder, unexpectedAbsentCnt uint32) {
	builder.PrependUint32Slot(30, unexpectedAbsentCnt, 0)
}
func SlaveOnlineInfoResponseAddIsDeviceEmulation(builder *flatbuffers.Builder, isDeviceEmulation bool) {
	builder.PrependBoolSlot(31, isDeviceEmulation, false)
}
func SlaveOnlineInfoResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
