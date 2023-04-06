// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// This table defines a list of main or detailed diagnostics that should be registered.
type DiagnosisElementsT struct {
	DiagnosisElements []*DiagnosisElementT
}

func (t *DiagnosisElementsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	diagnosisElementsOffset := flatbuffers.UOffsetT(0)
	if t.DiagnosisElements != nil {
		diagnosisElementsLength := len(t.DiagnosisElements)
		diagnosisElementsOffsets := make([]flatbuffers.UOffsetT, diagnosisElementsLength)
		for j := 0; j < diagnosisElementsLength; j++ {
			diagnosisElementsOffsets[j] = t.DiagnosisElements[j].Pack(builder)
		}
		DiagnosisElementsStartDiagnosisElementsVector(builder, diagnosisElementsLength)
		for j := diagnosisElementsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(diagnosisElementsOffsets[j])
		}
		diagnosisElementsOffset = builder.EndVector(diagnosisElementsLength)
	}
	DiagnosisElementsStart(builder)
	DiagnosisElementsAddDiagnosisElements(builder, diagnosisElementsOffset)
	return DiagnosisElementsEnd(builder)
}

func (rcv *DiagnosisElements) UnPackTo(t *DiagnosisElementsT) {
	diagnosisElementsLength := rcv.DiagnosisElementsLength()
	t.DiagnosisElements = make([]*DiagnosisElementT, diagnosisElementsLength)
	for j := 0; j < diagnosisElementsLength; j++ {
		x := DiagnosisElement{}
		rcv.DiagnosisElements(&x, j)
		t.DiagnosisElements[j] = x.UnPack()
	}
}

func (rcv *DiagnosisElements) UnPack() *DiagnosisElementsT {
	if rcv == nil { return nil }
	t := &DiagnosisElementsT{}
	rcv.UnPackTo(t)
	return t
}

type DiagnosisElements struct {
	_tab flatbuffers.Table
}

func GetRootAsDiagnosisElements(buf []byte, offset flatbuffers.UOffsetT) *DiagnosisElements {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DiagnosisElements{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDiagnosisElements(buf []byte, offset flatbuffers.UOffsetT) *DiagnosisElements {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DiagnosisElements{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DiagnosisElements) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DiagnosisElements) Table() flatbuffers.Table {
	return rcv._tab
}

/// List of main or detailed diagnostics.
func (rcv *DiagnosisElements) DiagnosisElements(obj *DiagnosisElement, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *DiagnosisElements) DiagnosisElementsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// List of main or detailed diagnostics.
func DiagnosisElementsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func DiagnosisElementsAddDiagnosisElements(builder *flatbuffers.Builder, diagnosisElements flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(diagnosisElements), 0)
}
func DiagnosisElementsStartDiagnosisElementsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DiagnosisElementsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
