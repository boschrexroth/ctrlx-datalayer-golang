// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ProgramsT struct {
	Programs []*ProgramT
}

func (t *ProgramsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	programsOffset := flatbuffers.UOffsetT(0)
	if t.Programs != nil {
		programsLength := len(t.Programs)
		programsOffsets := make([]flatbuffers.UOffsetT, programsLength)
		for j := 0; j < programsLength; j++ {
			programsOffsets[j] = t.Programs[j].Pack(builder)
		}
		ProgramsStartProgramsVector(builder, programsLength)
		for j := programsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(programsOffsets[j])
		}
		programsOffset = builder.EndVector(programsLength)
	}
	ProgramsStart(builder)
	ProgramsAddPrograms(builder, programsOffset)
	return ProgramsEnd(builder)
}

func (rcv *Programs) UnPackTo(t *ProgramsT) {
	programsLength := rcv.ProgramsLength()
	t.Programs = make([]*ProgramT, programsLength)
	for j := 0; j < programsLength; j++ {
		x := Program{}
		rcv.Programs(&x, j)
		t.Programs[j] = x.UnPack()
	}
}

func (rcv *Programs) UnPack() *ProgramsT {
	if rcv == nil { return nil }
	t := &ProgramsT{}
	rcv.UnPackTo(t)
	return t
}

type Programs struct {
	_tab flatbuffers.Table
}

func GetRootAsPrograms(buf []byte, offset flatbuffers.UOffsetT) *Programs {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Programs{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPrograms(buf []byte, offset flatbuffers.UOffsetT) *Programs {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Programs{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Programs) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Programs) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Programs) Programs(obj *Program, j int) bool {
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

func (rcv *Programs) ProgramsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ProgramsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ProgramsAddPrograms(builder *flatbuffers.Builder, programs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(programs), 0)
}
func ProgramsStartProgramsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ProgramsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
