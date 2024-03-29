// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package parameter

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type DataDetailsT struct {
	CommonErrorCode uint16 `json:"commonErrorCode"`
	SpecificErrorCode uint32 `json:"specificErrorCode"`
	ValidElements uint16 `json:"validElements"`
	Data []byte `json:"data"`
}

func (t *DataDetailsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataOffset = builder.CreateByteString(t.Data)
	}
	DataDetailsStart(builder)
	DataDetailsAddCommonErrorCode(builder, t.CommonErrorCode)
	DataDetailsAddSpecificErrorCode(builder, t.SpecificErrorCode)
	DataDetailsAddValidElements(builder, t.ValidElements)
	DataDetailsAddData(builder, dataOffset)
	return DataDetailsEnd(builder)
}

func (rcv *DataDetails) UnPackTo(t *DataDetailsT) {
	t.CommonErrorCode = rcv.CommonErrorCode()
	t.SpecificErrorCode = rcv.SpecificErrorCode()
	t.ValidElements = rcv.ValidElements()
	t.Data = rcv.DataBytes()
}

func (rcv *DataDetails) UnPack() *DataDetailsT {
	if rcv == nil { return nil }
	t := &DataDetailsT{}
	rcv.UnPackTo(t)
	return t
}

type DataDetails struct {
	_tab flatbuffers.Table
}

func GetRootAsDataDetails(buf []byte, offset flatbuffers.UOffsetT) *DataDetails {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DataDetails{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDataDetails(buf []byte, offset flatbuffers.UOffsetT) *DataDetails {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DataDetails{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DataDetails) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DataDetails) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DataDetails) CommonErrorCode() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DataDetails) MutateCommonErrorCode(n uint16) bool {
	return rcv._tab.MutateUint16Slot(4, n)
}

func (rcv *DataDetails) SpecificErrorCode() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DataDetails) MutateSpecificErrorCode(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *DataDetails) ValidElements() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DataDetails) MutateValidElements(n uint16) bool {
	return rcv._tab.MutateUint16Slot(8, n)
}

func (rcv *DataDetails) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *DataDetails) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DataDetails) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DataDetails) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func DataDetailsStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DataDetailsAddCommonErrorCode(builder *flatbuffers.Builder, commonErrorCode uint16) {
	builder.PrependUint16Slot(0, commonErrorCode, 0)
}
func DataDetailsAddSpecificErrorCode(builder *flatbuffers.Builder, specificErrorCode uint32) {
	builder.PrependUint32Slot(1, specificErrorCode, 0)
}
func DataDetailsAddValidElements(builder *flatbuffers.Builder, validElements uint16) {
	builder.PrependUint16Slot(2, validElements, 0)
}
func DataDetailsAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(data), 0)
}
func DataDetailsStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func DataDetailsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
