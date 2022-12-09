// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package automation

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StatsDlClientT struct {
	Converter uint32
}

func (t *StatsDlClientT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	StatsDlClientStart(builder)
	StatsDlClientAddConverter(builder, t.Converter)
	return StatsDlClientEnd(builder)
}

func (rcv *StatsDlClient) UnPackTo(t *StatsDlClientT) {
	t.Converter = rcv.Converter()
}

func (rcv *StatsDlClient) UnPack() *StatsDlClientT {
	if rcv == nil { return nil }
	t := &StatsDlClientT{}
	rcv.UnPackTo(t)
	return t
}

type StatsDlClient struct {
	_tab flatbuffers.Table
}

func GetRootAsStatsDlClient(buf []byte, offset flatbuffers.UOffsetT) *StatsDlClient {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StatsDlClient{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStatsDlClient(buf []byte, offset flatbuffers.UOffsetT) *StatsDlClient {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StatsDlClient{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StatsDlClient) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StatsDlClient) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StatsDlClient) Converter() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *StatsDlClient) MutateConverter(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func StatsDlClientStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func StatsDlClientAddConverter(builder *flatbuffers.Builder, converter uint32) {
	builder.PrependUint32Slot(0, converter, 0)
}
func StatsDlClientEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
