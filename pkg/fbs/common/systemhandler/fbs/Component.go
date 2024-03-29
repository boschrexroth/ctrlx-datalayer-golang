// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"

	comm__datalayer "github.com/boschrexroth/ctrlx-datalayer-golang/v2/pkg/fbs/comm/datalayer"
)

/// Information about the component
type ComponentT struct {
	Name string `json:"name"`
	State *StateT `json:"state"`
	Metadata *comm__datalayer.MetadataT `json:"metadata"`
}

func (t *ComponentT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	stateOffset := t.State.Pack(builder)
	metadataOffset := t.Metadata.Pack(builder)
	ComponentStart(builder)
	ComponentAddName(builder, nameOffset)
	ComponentAddState(builder, stateOffset)
	ComponentAddMetadata(builder, metadataOffset)
	return ComponentEnd(builder)
}

func (rcv *Component) UnPackTo(t *ComponentT) {
	t.Name = string(rcv.Name())
	t.State = rcv.State(nil).UnPack()
	t.Metadata = rcv.Metadata(nil).UnPack()
}

func (rcv *Component) UnPack() *ComponentT {
	if rcv == nil { return nil }
	t := &ComponentT{}
	rcv.UnPackTo(t)
	return t
}

type Component struct {
	_tab flatbuffers.Table
}

func GetRootAsComponent(buf []byte, offset flatbuffers.UOffsetT) *Component {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Component{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsComponent(buf []byte, offset flatbuffers.UOffsetT) *Component {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Component{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Component) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Component) Table() flatbuffers.Table {
	return rcv._tab
}

/// Name of component, must be unique
func (rcv *Component) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Name of component, must be unique
/// Current state of component
func (rcv *Component) State(obj *State) *State {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(State)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Current state of component
/// Metadata information about the Data Layer child node which will be created for the component
func (rcv *Component) Metadata(obj *comm__datalayer.Metadata) *comm__datalayer.Metadata {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(comm__datalayer.Metadata)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// Metadata information about the Data Layer child node which will be created for the component
func ComponentStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ComponentAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func ComponentAddState(builder *flatbuffers.Builder, state flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(state), 0)
}
func ComponentAddMetadata(builder *flatbuffers.Builder, metadata flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(metadata), 0)
}
func ComponentEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
