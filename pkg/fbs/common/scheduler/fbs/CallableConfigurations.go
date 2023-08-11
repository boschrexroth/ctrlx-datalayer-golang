// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Callable configurations of a callable factory
type CallableConfigurationsT struct {
	Configurations []*CallableConfigurationT `json:"configurations"`
}

func (t *CallableConfigurationsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	configurationsOffset := flatbuffers.UOffsetT(0)
	if t.Configurations != nil {
		configurationsLength := len(t.Configurations)
		configurationsOffsets := make([]flatbuffers.UOffsetT, configurationsLength)
		for j := 0; j < configurationsLength; j++ {
			configurationsOffsets[j] = t.Configurations[j].Pack(builder)
		}
		CallableConfigurationsStartConfigurationsVector(builder, configurationsLength)
		for j := configurationsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(configurationsOffsets[j])
		}
		configurationsOffset = builder.EndVector(configurationsLength)
	}
	CallableConfigurationsStart(builder)
	CallableConfigurationsAddConfigurations(builder, configurationsOffset)
	return CallableConfigurationsEnd(builder)
}

func (rcv *CallableConfigurations) UnPackTo(t *CallableConfigurationsT) {
	configurationsLength := rcv.ConfigurationsLength()
	t.Configurations = make([]*CallableConfigurationT, configurationsLength)
	for j := 0; j < configurationsLength; j++ {
		x := CallableConfiguration{}
		rcv.Configurations(&x, j)
		t.Configurations[j] = x.UnPack()
	}
}

func (rcv *CallableConfigurations) UnPack() *CallableConfigurationsT {
	if rcv == nil { return nil }
	t := &CallableConfigurationsT{}
	rcv.UnPackTo(t)
	return t
}

type CallableConfigurations struct {
	_tab flatbuffers.Table
}

func GetRootAsCallableConfigurations(buf []byte, offset flatbuffers.UOffsetT) *CallableConfigurations {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CallableConfigurations{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCallableConfigurations(buf []byte, offset flatbuffers.UOffsetT) *CallableConfigurations {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CallableConfigurations{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CallableConfigurations) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CallableConfigurations) Table() flatbuffers.Table {
	return rcv._tab
}

/// Callable configurations of a callable factory
func (rcv *CallableConfigurations) Configurations(obj *CallableConfiguration, j int) bool {
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

func (rcv *CallableConfigurations) ConfigurationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Callable configurations of a callable factory
func CallableConfigurationsStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func CallableConfigurationsAddConfigurations(builder *flatbuffers.Builder, configurations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(configurations), 0)
}
func CallableConfigurationsStartConfigurationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func CallableConfigurationsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
