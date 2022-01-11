// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// parameters of the axis command add-to-kin
type AxsCmdAddToKinData struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCmdAddToKinData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdAddToKinData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCmdAddToKinData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCmdAddToKinData(buf []byte, offset flatbuffers.UOffsetT) *AxsCmdAddToKinData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCmdAddToKinData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCmdAddToKinData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCmdAddToKinData) Table() flatbuffers.Table {
	return rcv._tab
}

/// name of the kinematics to which this axis should be added
func (rcv *AxsCmdAddToKinData) KinName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name of the kinematics to which this axis should be added
/// should this be a buffered command?
func (rcv *AxsCmdAddToKinData) Buffered() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// should this be a buffered command?
func (rcv *AxsCmdAddToKinData) MutateBuffered(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func AxsCmdAddToKinDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AxsCmdAddToKinDataAddKinName(builder *flatbuffers.Builder, kinName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(kinName), 0)
}
func AxsCmdAddToKinDataAddBuffered(builder *flatbuffers.Builder, buffered bool) {
	builder.PrependBoolSlot(1, buffered, false)
}
func AxsCmdAddToKinDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}