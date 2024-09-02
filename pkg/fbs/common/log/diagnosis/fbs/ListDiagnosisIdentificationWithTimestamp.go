// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// This table defines a list of uniquely identifiable diagnostic logs with the related additional log information (e.g. it is used for the alarms).
type ListDiagnosisIdentificationWithTimestampT struct {
	ListDiagnosisIdentificationWithTimestamp []*DiagnosisIdentificationWithTimestampT `json:"listDiagnosisIdentificationWithTimestamp"`
}

func (t *ListDiagnosisIdentificationWithTimestampT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	listDiagnosisIdentificationWithTimestampOffset := flatbuffers.UOffsetT(0)
	if t.ListDiagnosisIdentificationWithTimestamp != nil {
		listDiagnosisIdentificationWithTimestampLength := len(t.ListDiagnosisIdentificationWithTimestamp)
		listDiagnosisIdentificationWithTimestampOffsets := make([]flatbuffers.UOffsetT, listDiagnosisIdentificationWithTimestampLength)
		for j := 0; j < listDiagnosisIdentificationWithTimestampLength; j++ {
			listDiagnosisIdentificationWithTimestampOffsets[j] = t.ListDiagnosisIdentificationWithTimestamp[j].Pack(builder)
		}
		ListDiagnosisIdentificationWithTimestampStartListDiagnosisIdentificationWithTimestampVector(builder, listDiagnosisIdentificationWithTimestampLength)
		for j := listDiagnosisIdentificationWithTimestampLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(listDiagnosisIdentificationWithTimestampOffsets[j])
		}
		listDiagnosisIdentificationWithTimestampOffset = builder.EndVector(listDiagnosisIdentificationWithTimestampLength)
	}
	ListDiagnosisIdentificationWithTimestampStart(builder)
	ListDiagnosisIdentificationWithTimestampAddListDiagnosisIdentificationWithTimestamp(builder, listDiagnosisIdentificationWithTimestampOffset)
	return ListDiagnosisIdentificationWithTimestampEnd(builder)
}

func (rcv *ListDiagnosisIdentificationWithTimestamp) UnPackTo(t *ListDiagnosisIdentificationWithTimestampT) {
	listDiagnosisIdentificationWithTimestampLength := rcv.ListDiagnosisIdentificationWithTimestampLength()
	t.ListDiagnosisIdentificationWithTimestamp = make([]*DiagnosisIdentificationWithTimestampT, listDiagnosisIdentificationWithTimestampLength)
	for j := 0; j < listDiagnosisIdentificationWithTimestampLength; j++ {
		x := DiagnosisIdentificationWithTimestamp{}
		rcv.ListDiagnosisIdentificationWithTimestamp(&x, j)
		t.ListDiagnosisIdentificationWithTimestamp[j] = x.UnPack()
	}
}

func (rcv *ListDiagnosisIdentificationWithTimestamp) UnPack() *ListDiagnosisIdentificationWithTimestampT {
	if rcv == nil { return nil }
	t := &ListDiagnosisIdentificationWithTimestampT{}
	rcv.UnPackTo(t)
	return t
}

type ListDiagnosisIdentificationWithTimestamp struct {
	_tab flatbuffers.Table
}

func GetRootAsListDiagnosisIdentificationWithTimestamp(buf []byte, offset flatbuffers.UOffsetT) *ListDiagnosisIdentificationWithTimestamp {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ListDiagnosisIdentificationWithTimestamp{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsListDiagnosisIdentificationWithTimestamp(buf []byte, offset flatbuffers.UOffsetT) *ListDiagnosisIdentificationWithTimestamp {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ListDiagnosisIdentificationWithTimestamp{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ListDiagnosisIdentificationWithTimestamp) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ListDiagnosisIdentificationWithTimestamp) Table() flatbuffers.Table {
	return rcv._tab
}

/// List of uniquely identifiable diagnostic logs with the related timestamp.
func (rcv *ListDiagnosisIdentificationWithTimestamp) ListDiagnosisIdentificationWithTimestamp(obj *DiagnosisIdentificationWithTimestamp, j int) bool {
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

func (rcv *ListDiagnosisIdentificationWithTimestamp) ListDiagnosisIdentificationWithTimestampLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// List of uniquely identifiable diagnostic logs with the related timestamp.
func ListDiagnosisIdentificationWithTimestampStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ListDiagnosisIdentificationWithTimestampAddListDiagnosisIdentificationWithTimestamp(builder *flatbuffers.Builder, listDiagnosisIdentificationWithTimestamp flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(listDiagnosisIdentificationWithTimestamp), 0)
}
func ListDiagnosisIdentificationWithTimestampStartListDiagnosisIdentificationWithTimestampVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ListDiagnosisIdentificationWithTimestampEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
