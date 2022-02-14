// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Diagnosis struct {
	_tab flatbuffers.Table
}

func GetRootAsDiagnosis(buf []byte, offset flatbuffers.UOffsetT) *Diagnosis {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Diagnosis{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDiagnosis(buf []byte, offset flatbuffers.UOffsetT) *Diagnosis {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Diagnosis{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Diagnosis) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Diagnosis) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Diagnosis) MainDiagnosisCode() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Diagnosis) MutateMainDiagnosisCode(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Diagnosis) DetailedDiagnosisCode() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Diagnosis) MutateDetailedDiagnosisCode(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *Diagnosis) DynamicDescription() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Diagnosis) Entity() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DiagnosisStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func DiagnosisAddMainDiagnosisCode(builder *flatbuffers.Builder, mainDiagnosisCode uint32) {
	builder.PrependUint32Slot(0, mainDiagnosisCode, 0)
}
func DiagnosisAddDetailedDiagnosisCode(builder *flatbuffers.Builder, detailedDiagnosisCode uint32) {
	builder.PrependUint32Slot(1, detailedDiagnosisCode, 0)
}
func DiagnosisAddDynamicDescription(builder *flatbuffers.Builder, dynamicDescription flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(dynamicDescription), 0)
}
func DiagnosisAddEntity(builder *flatbuffers.Builder, entity flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(entity), 0)
}
func DiagnosisEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
