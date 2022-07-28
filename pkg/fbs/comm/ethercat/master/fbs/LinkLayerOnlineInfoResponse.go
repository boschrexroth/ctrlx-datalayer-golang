// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LinkLayerOnlineInfoResponseT struct {
	Port string
	LinkLayer string
	Arguments string
	Message string
	MacAddress []byte
	FrameRepeatCntSupport bool
	LinkSpeed uint32
	LinkMode LinkMode
	LinkStatus LinkStatus
	RefClockWidth uint32
	SystemTime uint64
	PhysicalErrorCnt uint32
	TelegramErrorCnt uint32
}

func (t *LinkLayerOnlineInfoResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	portOffset := builder.CreateString(t.Port)
	linkLayerOffset := builder.CreateString(t.LinkLayer)
	argumentsOffset := builder.CreateString(t.Arguments)
	messageOffset := builder.CreateString(t.Message)
	macAddressOffset := flatbuffers.UOffsetT(0)
	if t.MacAddress != nil {
		macAddressOffset = builder.CreateByteString(t.MacAddress)
	}
	LinkLayerOnlineInfoResponseStart(builder)
	LinkLayerOnlineInfoResponseAddPort(builder, portOffset)
	LinkLayerOnlineInfoResponseAddLinkLayer(builder, linkLayerOffset)
	LinkLayerOnlineInfoResponseAddArguments(builder, argumentsOffset)
	LinkLayerOnlineInfoResponseAddMessage(builder, messageOffset)
	LinkLayerOnlineInfoResponseAddMacAddress(builder, macAddressOffset)
	LinkLayerOnlineInfoResponseAddFrameRepeatCntSupport(builder, t.FrameRepeatCntSupport)
	LinkLayerOnlineInfoResponseAddLinkSpeed(builder, t.LinkSpeed)
	LinkLayerOnlineInfoResponseAddLinkMode(builder, t.LinkMode)
	LinkLayerOnlineInfoResponseAddLinkStatus(builder, t.LinkStatus)
	LinkLayerOnlineInfoResponseAddRefClockWidth(builder, t.RefClockWidth)
	LinkLayerOnlineInfoResponseAddSystemTime(builder, t.SystemTime)
	LinkLayerOnlineInfoResponseAddPhysicalErrorCnt(builder, t.PhysicalErrorCnt)
	LinkLayerOnlineInfoResponseAddTelegramErrorCnt(builder, t.TelegramErrorCnt)
	return LinkLayerOnlineInfoResponseEnd(builder)
}

func (rcv *LinkLayerOnlineInfoResponse) UnPackTo(t *LinkLayerOnlineInfoResponseT) {
	t.Port = string(rcv.Port())
	t.LinkLayer = string(rcv.LinkLayer())
	t.Arguments = string(rcv.Arguments())
	t.Message = string(rcv.Message())
	t.MacAddress = rcv.MacAddressBytes()
	t.FrameRepeatCntSupport = rcv.FrameRepeatCntSupport()
	t.LinkSpeed = rcv.LinkSpeed()
	t.LinkMode = rcv.LinkMode()
	t.LinkStatus = rcv.LinkStatus()
	t.RefClockWidth = rcv.RefClockWidth()
	t.SystemTime = rcv.SystemTime()
	t.PhysicalErrorCnt = rcv.PhysicalErrorCnt()
	t.TelegramErrorCnt = rcv.TelegramErrorCnt()
}

func (rcv *LinkLayerOnlineInfoResponse) UnPack() *LinkLayerOnlineInfoResponseT {
	if rcv == nil { return nil }
	t := &LinkLayerOnlineInfoResponseT{}
	rcv.UnPackTo(t)
	return t
}

type LinkLayerOnlineInfoResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsLinkLayerOnlineInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *LinkLayerOnlineInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LinkLayerOnlineInfoResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLinkLayerOnlineInfoResponse(buf []byte, offset flatbuffers.UOffsetT) *LinkLayerOnlineInfoResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LinkLayerOnlineInfoResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LinkLayerOnlineInfoResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LinkLayerOnlineInfoResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LinkLayerOnlineInfoResponse) Port() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LinkLayerOnlineInfoResponse) LinkLayer() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LinkLayerOnlineInfoResponse) Arguments() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LinkLayerOnlineInfoResponse) Message() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LinkLayerOnlineInfoResponse) MacAddress(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *LinkLayerOnlineInfoResponse) MacAddressLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *LinkLayerOnlineInfoResponse) MacAddressBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LinkLayerOnlineInfoResponse) MutateMacAddress(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *LinkLayerOnlineInfoResponse) FrameRepeatCntSupport() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *LinkLayerOnlineInfoResponse) MutateFrameRepeatCntSupport(n bool) bool {
	return rcv._tab.MutateBoolSlot(14, n)
}

func (rcv *LinkLayerOnlineInfoResponse) LinkSpeed() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LinkLayerOnlineInfoResponse) MutateLinkSpeed(n uint32) bool {
	return rcv._tab.MutateUint32Slot(16, n)
}

func (rcv *LinkLayerOnlineInfoResponse) LinkMode() LinkMode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return LinkMode(rcv._tab.GetUint32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *LinkLayerOnlineInfoResponse) MutateLinkMode(n LinkMode) bool {
	return rcv._tab.MutateUint32Slot(18, uint32(n))
}

func (rcv *LinkLayerOnlineInfoResponse) LinkStatus() LinkStatus {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return LinkStatus(rcv._tab.GetUint32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *LinkLayerOnlineInfoResponse) MutateLinkStatus(n LinkStatus) bool {
	return rcv._tab.MutateUint32Slot(20, uint32(n))
}

func (rcv *LinkLayerOnlineInfoResponse) RefClockWidth() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LinkLayerOnlineInfoResponse) MutateRefClockWidth(n uint32) bool {
	return rcv._tab.MutateUint32Slot(22, n)
}

func (rcv *LinkLayerOnlineInfoResponse) SystemTime() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LinkLayerOnlineInfoResponse) MutateSystemTime(n uint64) bool {
	return rcv._tab.MutateUint64Slot(24, n)
}

func (rcv *LinkLayerOnlineInfoResponse) PhysicalErrorCnt() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LinkLayerOnlineInfoResponse) MutatePhysicalErrorCnt(n uint32) bool {
	return rcv._tab.MutateUint32Slot(26, n)
}

func (rcv *LinkLayerOnlineInfoResponse) TelegramErrorCnt() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LinkLayerOnlineInfoResponse) MutateTelegramErrorCnt(n uint32) bool {
	return rcv._tab.MutateUint32Slot(28, n)
}

func LinkLayerOnlineInfoResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(13)
}
func LinkLayerOnlineInfoResponseAddPort(builder *flatbuffers.Builder, port flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(port), 0)
}
func LinkLayerOnlineInfoResponseAddLinkLayer(builder *flatbuffers.Builder, linkLayer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(linkLayer), 0)
}
func LinkLayerOnlineInfoResponseAddArguments(builder *flatbuffers.Builder, arguments flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(arguments), 0)
}
func LinkLayerOnlineInfoResponseAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(message), 0)
}
func LinkLayerOnlineInfoResponseAddMacAddress(builder *flatbuffers.Builder, macAddress flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(macAddress), 0)
}
func LinkLayerOnlineInfoResponseStartMacAddressVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func LinkLayerOnlineInfoResponseAddFrameRepeatCntSupport(builder *flatbuffers.Builder, frameRepeatCntSupport bool) {
	builder.PrependBoolSlot(5, frameRepeatCntSupport, false)
}
func LinkLayerOnlineInfoResponseAddLinkSpeed(builder *flatbuffers.Builder, linkSpeed uint32) {
	builder.PrependUint32Slot(6, linkSpeed, 0)
}
func LinkLayerOnlineInfoResponseAddLinkMode(builder *flatbuffers.Builder, linkMode LinkMode) {
	builder.PrependUint32Slot(7, uint32(linkMode), 0)
}
func LinkLayerOnlineInfoResponseAddLinkStatus(builder *flatbuffers.Builder, linkStatus LinkStatus) {
	builder.PrependUint32Slot(8, uint32(linkStatus), 0)
}
func LinkLayerOnlineInfoResponseAddRefClockWidth(builder *flatbuffers.Builder, refClockWidth uint32) {
	builder.PrependUint32Slot(9, refClockWidth, 0)
}
func LinkLayerOnlineInfoResponseAddSystemTime(builder *flatbuffers.Builder, systemTime uint64) {
	builder.PrependUint64Slot(10, systemTime, 0)
}
func LinkLayerOnlineInfoResponseAddPhysicalErrorCnt(builder *flatbuffers.Builder, physicalErrorCnt uint32) {
	builder.PrependUint32Slot(11, physicalErrorCnt, 0)
}
func LinkLayerOnlineInfoResponseAddTelegramErrorCnt(builder *flatbuffers.Builder, telegramErrorCnt uint32) {
	builder.PrependUint32Slot(12, telegramErrorCnt, 0)
}
func LinkLayerOnlineInfoResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
