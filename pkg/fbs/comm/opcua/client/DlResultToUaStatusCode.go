// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package client

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DlResultToUaStatusCode struct {
	_tab flatbuffers.Table
}

func GetRootAsDlResultToUaStatusCode(buf []byte, offset flatbuffers.UOffsetT) *DlResultToUaStatusCode {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DlResultToUaStatusCode{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDlResultToUaStatusCode(buf []byte, offset flatbuffers.UOffsetT) *DlResultToUaStatusCode {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DlResultToUaStatusCode{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DlResultToUaStatusCode) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DlResultToUaStatusCode) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DlResultToUaStatusCode) DlResult() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DlResultToUaStatusCode) MutateDlResult(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *DlResultToUaStatusCode) UaStatusCode() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DlResultToUaStatusCode) MutateUaStatusCode(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *DlResultToUaStatusCode) UaStatusCodeName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DlResultToUaStatusCode) UaStatusCodeDescr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DlResultToUaStatusCodeStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DlResultToUaStatusCodeAddDlResult(builder *flatbuffers.Builder, dlResult uint32) {
	builder.PrependUint32Slot(0, dlResult, 0)
}
func DlResultToUaStatusCodeAddUaStatusCode(builder *flatbuffers.Builder, uaStatusCode uint32) {
	builder.PrependUint32Slot(1, uaStatusCode, 0)
}
func DlResultToUaStatusCodeAddUaStatusCodeName(builder *flatbuffers.Builder, uaStatusCodeName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(uaStatusCodeName), 0)
}
func DlResultToUaStatusCodeAddUaStatusCodeDescr(builder *flatbuffers.Builder, uaStatusCodeDescr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(uaStatusCodeDescr), 0)
}
func DlResultToUaStatusCodeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
