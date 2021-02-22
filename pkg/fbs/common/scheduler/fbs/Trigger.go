// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Trigger struct {
	_tab flatbuffers.Table
}

func GetRootAsTrigger(buf []byte, offset flatbuffers.UOffsetT) *Trigger {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Trigger{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTrigger(buf []byte, offset flatbuffers.UOffsetT) *Trigger {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Trigger{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Trigger) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Trigger) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Trigger) Trigger() CurrentTrigger {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return CurrentTrigger(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Trigger) MutateTrigger(n CurrentTrigger) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func TriggerStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TriggerAddTrigger(builder *flatbuffers.Builder, trigger CurrentTrigger) {
	builder.PrependInt8Slot(0, int8(trigger), 0)
}
func TriggerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
