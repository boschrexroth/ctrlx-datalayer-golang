// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Configuration struct {
	_tab flatbuffers.Table
}

func GetRootAsConfiguration(buf []byte, offset flatbuffers.UOffsetT) *Configuration {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Configuration{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsConfiguration(buf []byte, offset flatbuffers.UOffsetT) *Configuration {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Configuration{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Configuration) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Configuration) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Configuration) Type() CurrentConfiguration {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return CurrentConfiguration(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Configuration) MutateType(n CurrentConfiguration) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func ConfigurationStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ConfigurationAddType(builder *flatbuffers.Builder, type_ CurrentConfiguration) {
	builder.PrependInt8Slot(0, int8(type_), 0)
}
func ConfigurationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
