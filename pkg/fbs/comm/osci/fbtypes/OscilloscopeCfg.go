// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type OscilloscopeCfgT struct {
	Name string
	Channels []*ChannelCfgT
	Buffer *BufferCfgT
	Trigger *TriggerCfgT
	Diagramview []*DiagramCfgT
}

func (t *OscilloscopeCfgT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := builder.CreateString(t.Name)
	channelsOffset := flatbuffers.UOffsetT(0)
	if t.Channels != nil {
		channelsLength := len(t.Channels)
		channelsOffsets := make([]flatbuffers.UOffsetT, channelsLength)
		for j := 0; j < channelsLength; j++ {
			channelsOffsets[j] = t.Channels[j].Pack(builder)
		}
		OscilloscopeCfgStartChannelsVector(builder, channelsLength)
		for j := channelsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(channelsOffsets[j])
		}
		channelsOffset = builder.EndVector(channelsLength)
	}
	bufferOffset := t.Buffer.Pack(builder)
	triggerOffset := t.Trigger.Pack(builder)
	diagramviewOffset := flatbuffers.UOffsetT(0)
	if t.Diagramview != nil {
		diagramviewLength := len(t.Diagramview)
		diagramviewOffsets := make([]flatbuffers.UOffsetT, diagramviewLength)
		for j := 0; j < diagramviewLength; j++ {
			diagramviewOffsets[j] = t.Diagramview[j].Pack(builder)
		}
		OscilloscopeCfgStartDiagramviewVector(builder, diagramviewLength)
		for j := diagramviewLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(diagramviewOffsets[j])
		}
		diagramviewOffset = builder.EndVector(diagramviewLength)
	}
	OscilloscopeCfgStart(builder)
	OscilloscopeCfgAddName(builder, nameOffset)
	OscilloscopeCfgAddChannels(builder, channelsOffset)
	OscilloscopeCfgAddBuffer(builder, bufferOffset)
	OscilloscopeCfgAddTrigger(builder, triggerOffset)
	OscilloscopeCfgAddDiagramview(builder, diagramviewOffset)
	return OscilloscopeCfgEnd(builder)
}

func (rcv *OscilloscopeCfg) UnPackTo(t *OscilloscopeCfgT) {
	t.Name = string(rcv.Name())
	channelsLength := rcv.ChannelsLength()
	t.Channels = make([]*ChannelCfgT, channelsLength)
	for j := 0; j < channelsLength; j++ {
		x := ChannelCfg{}
		rcv.Channels(&x, j)
		t.Channels[j] = x.UnPack()
	}
	t.Buffer = rcv.Buffer(nil).UnPack()
	t.Trigger = rcv.Trigger(nil).UnPack()
	diagramviewLength := rcv.DiagramviewLength()
	t.Diagramview = make([]*DiagramCfgT, diagramviewLength)
	for j := 0; j < diagramviewLength; j++ {
		x := DiagramCfg{}
		rcv.Diagramview(&x, j)
		t.Diagramview[j] = x.UnPack()
	}
}

func (rcv *OscilloscopeCfg) UnPack() *OscilloscopeCfgT {
	if rcv == nil { return nil }
	t := &OscilloscopeCfgT{}
	rcv.UnPackTo(t)
	return t
}

type OscilloscopeCfg struct {
	_tab flatbuffers.Table
}

func GetRootAsOscilloscopeCfg(buf []byte, offset flatbuffers.UOffsetT) *OscilloscopeCfg {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &OscilloscopeCfg{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsOscilloscopeCfg(buf []byte, offset flatbuffers.UOffsetT) *OscilloscopeCfg {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &OscilloscopeCfg{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *OscilloscopeCfg) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *OscilloscopeCfg) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *OscilloscopeCfg) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *OscilloscopeCfg) Channels(obj *ChannelCfg, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *OscilloscopeCfg) ChannelsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *OscilloscopeCfg) Buffer(obj *BufferCfg) *BufferCfg {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BufferCfg)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *OscilloscopeCfg) Trigger(obj *TriggerCfg) *TriggerCfg {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(TriggerCfg)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *OscilloscopeCfg) Diagramview(obj *DiagramCfg, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *OscilloscopeCfg) DiagramviewLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func OscilloscopeCfgStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func OscilloscopeCfgAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func OscilloscopeCfgAddChannels(builder *flatbuffers.Builder, channels flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(channels), 0)
}
func OscilloscopeCfgStartChannelsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func OscilloscopeCfgAddBuffer(builder *flatbuffers.Builder, buffer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(buffer), 0)
}
func OscilloscopeCfgAddTrigger(builder *flatbuffers.Builder, trigger flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(trigger), 0)
}
func OscilloscopeCfgAddDiagramview(builder *flatbuffers.Builder, diagramview flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(diagramview), 0)
}
func OscilloscopeCfgStartDiagramviewVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func OscilloscopeCfgEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
