// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SamplingT struct {
	SamplingInterval uint64 `json:"samplingInterval"`
}

func (t *SamplingT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	SamplingStart(builder)
	SamplingAddSamplingInterval(builder, t.SamplingInterval)
	return SamplingEnd(builder)
}

func (rcv *Sampling) UnPackTo(t *SamplingT) {
	t.SamplingInterval = rcv.SamplingInterval()
}

func (rcv *Sampling) UnPack() *SamplingT {
	if rcv == nil { return nil }
	t := &SamplingT{}
	rcv.UnPackTo(t)
	return t
}

type Sampling struct {
	_tab flatbuffers.Table
}

func GetRootAsSampling(buf []byte, offset flatbuffers.UOffsetT) *Sampling {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Sampling{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSampling(buf []byte, offset flatbuffers.UOffsetT) *Sampling {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Sampling{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Sampling) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Sampling) Table() flatbuffers.Table {
	return rcv._tab
}

/// sample time in µ seconds for minimum sampling of data - currently only multiples of 1000 are supported
func (rcv *Sampling) SamplingInterval() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 1000000
}

/// sample time in µ seconds for minimum sampling of data - currently only multiples of 1000 are supported
func (rcv *Sampling) MutateSamplingInterval(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func SamplingStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SamplingAddSamplingInterval(builder *flatbuffers.Builder, samplingInterval uint64) {
	builder.PrependUint64Slot(0, samplingInterval, 1000000)
}
func SamplingEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
