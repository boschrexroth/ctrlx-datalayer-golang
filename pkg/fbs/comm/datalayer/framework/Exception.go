// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package framework

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ExceptionT struct {
	Date string
	Name string
	Signal string
	Code string
	Register *RegisterT
	Stack []*StackentryT
}

func (t *ExceptionT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	dateOffset := builder.CreateString(t.Date)
	nameOffset := builder.CreateString(t.Name)
	signalOffset := builder.CreateString(t.Signal)
	codeOffset := builder.CreateString(t.Code)
	registerOffset := t.Register.Pack(builder)
	stackOffset := flatbuffers.UOffsetT(0)
	if t.Stack != nil {
		stackLength := len(t.Stack)
		stackOffsets := make([]flatbuffers.UOffsetT, stackLength)
		for j := 0; j < stackLength; j++ {
			stackOffsets[j] = t.Stack[j].Pack(builder)
		}
		ExceptionStartStackVector(builder, stackLength)
		for j := stackLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(stackOffsets[j])
		}
		stackOffset = builder.EndVector(stackLength)
	}
	ExceptionStart(builder)
	ExceptionAddDate(builder, dateOffset)
	ExceptionAddName(builder, nameOffset)
	ExceptionAddSignal(builder, signalOffset)
	ExceptionAddCode(builder, codeOffset)
	ExceptionAddRegister(builder, registerOffset)
	ExceptionAddStack(builder, stackOffset)
	return ExceptionEnd(builder)
}

func (rcv *Exception) UnPackTo(t *ExceptionT) {
	t.Date = string(rcv.Date())
	t.Name = string(rcv.Name())
	t.Signal = string(rcv.Signal())
	t.Code = string(rcv.Code())
	t.Register = rcv.Register(nil).UnPack()
	stackLength := rcv.StackLength()
	t.Stack = make([]*StackentryT, stackLength)
	for j := 0; j < stackLength; j++ {
		x := Stackentry{}
		rcv.Stack(&x, j)
		t.Stack[j] = x.UnPack()
	}
}

func (rcv *Exception) UnPack() *ExceptionT {
	if rcv == nil { return nil }
	t := &ExceptionT{}
	rcv.UnPackTo(t)
	return t
}

type Exception struct {
	_tab flatbuffers.Table
}

func GetRootAsException(buf []byte, offset flatbuffers.UOffsetT) *Exception {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Exception{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsException(buf []byte, offset flatbuffers.UOffsetT) *Exception {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Exception{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Exception) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Exception) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Exception) Date() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Exception) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Exception) Signal() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Exception) Code() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Exception) Register(obj *Register) *Register {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Register)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Exception) Stack(obj *Stackentry, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Exception) StackLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ExceptionStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func ExceptionAddDate(builder *flatbuffers.Builder, date flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(date), 0)
}
func ExceptionAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func ExceptionAddSignal(builder *flatbuffers.Builder, signal flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(signal), 0)
}
func ExceptionAddCode(builder *flatbuffers.Builder, code flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(code), 0)
}
func ExceptionAddRegister(builder *flatbuffers.Builder, register flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(register), 0)
}
func ExceptionAddStack(builder *flatbuffers.Builder, stack flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(stack), 0)
}
func ExceptionStartStackVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ExceptionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}