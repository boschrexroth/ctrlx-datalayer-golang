// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package testbuddy

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TestResultT struct {
	NumWarnings int32
	NumErrors int32
	Log []*LogT
	Data []*DataT
}

func (t *TestResultT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	logOffset := flatbuffers.UOffsetT(0)
	if t.Log != nil {
		logLength := len(t.Log)
		logOffsets := make([]flatbuffers.UOffsetT, logLength)
		for j := 0; j < logLength; j++ {
			logOffsets[j] = t.Log[j].Pack(builder)
		}
		TestResultStartLogVector(builder, logLength)
		for j := logLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(logOffsets[j])
		}
		logOffset = builder.EndVector(logLength)
	}
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataLength := len(t.Data)
		dataOffsets := make([]flatbuffers.UOffsetT, dataLength)
		for j := 0; j < dataLength; j++ {
			dataOffsets[j] = t.Data[j].Pack(builder)
		}
		TestResultStartDataVector(builder, dataLength)
		for j := dataLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(dataOffsets[j])
		}
		dataOffset = builder.EndVector(dataLength)
	}
	TestResultStart(builder)
	TestResultAddNumWarnings(builder, t.NumWarnings)
	TestResultAddNumErrors(builder, t.NumErrors)
	TestResultAddLog(builder, logOffset)
	TestResultAddData(builder, dataOffset)
	return TestResultEnd(builder)
}

func (rcv *TestResult) UnPackTo(t *TestResultT) {
	t.NumWarnings = rcv.NumWarnings()
	t.NumErrors = rcv.NumErrors()
	logLength := rcv.LogLength()
	t.Log = make([]*LogT, logLength)
	for j := 0; j < logLength; j++ {
		x := Log{}
		rcv.Log(&x, j)
		t.Log[j] = x.UnPack()
	}
	dataLength := rcv.DataLength()
	t.Data = make([]*DataT, dataLength)
	for j := 0; j < dataLength; j++ {
		x := Data{}
		rcv.Data(&x, j)
		t.Data[j] = x.UnPack()
	}
}

func (rcv *TestResult) UnPack() *TestResultT {
	if rcv == nil { return nil }
	t := &TestResultT{}
	rcv.UnPackTo(t)
	return t
}

type TestResult struct {
	_tab flatbuffers.Table
}

func GetRootAsTestResult(buf []byte, offset flatbuffers.UOffsetT) *TestResult {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TestResult{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTestResult(buf []byte, offset flatbuffers.UOffsetT) *TestResult {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TestResult{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TestResult) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TestResult) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TestResult) NumWarnings() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TestResult) MutateNumWarnings(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *TestResult) NumErrors() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TestResult) MutateNumErrors(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *TestResult) Log(obj *Log, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *TestResult) LogLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TestResult) Data(obj *Data, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *TestResult) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func TestResultStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func TestResultAddNumWarnings(builder *flatbuffers.Builder, numWarnings int32) {
	builder.PrependInt32Slot(0, numWarnings, 0)
}
func TestResultAddNumErrors(builder *flatbuffers.Builder, numErrors int32) {
	builder.PrependInt32Slot(1, numErrors, 0)
}
func TestResultAddLog(builder *flatbuffers.Builder, log flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(log), 0)
}
func TestResultStartLogVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TestResultAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(data), 0)
}
func TestResultStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TestResultEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}