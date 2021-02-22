// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Problem struct {
	_tab flatbuffers.Table
}

func GetRootAsProblem(buf []byte, offset flatbuffers.UOffsetT) *Problem {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Problem{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsProblem(buf []byte, offset flatbuffers.UOffsetT) *Problem {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Problem{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Problem) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Problem) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Problem) Type() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Problem) Title() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Problem) Status() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Problem) MutateStatus(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *Problem) Detail() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Problem) Instance() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Problem) MainDiagnosisCode() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Problem) DetailedDiagnosisCode() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Problem) DynamicDescription() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Problem) Severity() Severity {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return Severity(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Problem) MutateSeverity(n Severity) bool {
	return rcv._tab.MutateInt8Slot(20, int8(n))
}

func (rcv *Problem) Links(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Problem) LinksLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ProblemStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func ProblemAddType(builder *flatbuffers.Builder, type_ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(type_), 0)
}
func ProblemAddTitle(builder *flatbuffers.Builder, title flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(title), 0)
}
func ProblemAddStatus(builder *flatbuffers.Builder, status int32) {
	builder.PrependInt32Slot(2, status, 0)
}
func ProblemAddDetail(builder *flatbuffers.Builder, detail flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(detail), 0)
}
func ProblemAddInstance(builder *flatbuffers.Builder, instance flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(instance), 0)
}
func ProblemAddMainDiagnosisCode(builder *flatbuffers.Builder, mainDiagnosisCode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(mainDiagnosisCode), 0)
}
func ProblemAddDetailedDiagnosisCode(builder *flatbuffers.Builder, detailedDiagnosisCode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(detailedDiagnosisCode), 0)
}
func ProblemAddDynamicDescription(builder *flatbuffers.Builder, dynamicDescription flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(dynamicDescription), 0)
}
func ProblemAddSeverity(builder *flatbuffers.Builder, severity Severity) {
	builder.PrependInt8Slot(8, int8(severity), 0)
}
func ProblemAddLinks(builder *flatbuffers.Builder, links flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(links), 0)
}
func ProblemStartLinksVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ProblemEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
