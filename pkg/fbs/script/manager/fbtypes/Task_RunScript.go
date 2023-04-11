// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Task_RunScriptT struct {
	Name string
	Param []string
}

func (t *Task_RunScriptT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	paramOffset := flatbuffers.UOffsetT(0)
	if t.Param != nil {
		paramLength := len(t.Param)
		paramOffsets := make([]flatbuffers.UOffsetT, paramLength)
		for j := 0; j < paramLength; j++ {
			paramOffsets[j] = builder.CreateString(t.Param[j])
		}
		Task_RunScriptStartParamVector(builder, paramLength)
		for j := paramLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(paramOffsets[j])
		}
		paramOffset = builder.EndVector(paramLength)
	}
	Task_RunScriptStart(builder)
	Task_RunScriptAddName(builder, nameOffset)
	Task_RunScriptAddParam(builder, paramOffset)
	return Task_RunScriptEnd(builder)
}

func (rcv *Task_RunScript) UnPackTo(t *Task_RunScriptT) {
	t.Name = string(rcv.Name())
	paramLength := rcv.ParamLength()
	t.Param = make([]string, paramLength)
	for j := 0; j < paramLength; j++ {
		t.Param[j] = string(rcv.Param(j))
	}
}

func (rcv *Task_RunScript) UnPack() *Task_RunScriptT {
	if rcv == nil { return nil }
	t := &Task_RunScriptT{}
	rcv.UnPackTo(t)
	return t
}

type Task_RunScript struct {
	_tab flatbuffers.Table
}

func GetRootAsTask_RunScript(buf []byte, offset flatbuffers.UOffsetT) *Task_RunScript {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Task_RunScript{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTask_RunScript(buf []byte, offset flatbuffers.UOffsetT) *Task_RunScript {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Task_RunScript{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Task_RunScript) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Task_RunScript) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Task_RunScript) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Task_RunScript) Param(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Task_RunScript) ParamLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func Task_RunScriptStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Task_RunScriptAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func Task_RunScriptAddParam(builder *flatbuffers.Builder, param flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(param), 0)
}
func Task_RunScriptStartParamVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func Task_RunScriptEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
