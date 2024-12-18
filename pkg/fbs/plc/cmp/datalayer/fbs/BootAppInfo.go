// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// This object describes the boot application properties
type BootAppInfoT struct {
	Name string `json:"Name"`
	LastModified uint64 `json:"LastModified"`
}

func (t *BootAppInfoT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	nameOffset := flatbuffers.UOffsetT(0)
	if t.Name != "" {
		nameOffset = builder.CreateString(t.Name)
	}
	BootAppInfoStart(builder)
	BootAppInfoAddName(builder, nameOffset)
	BootAppInfoAddLastModified(builder, t.LastModified)
	return BootAppInfoEnd(builder)
}

func (rcv *BootAppInfo) UnPackTo(t *BootAppInfoT) {
	t.Name = string(rcv.Name())
	t.LastModified = rcv.LastModified()
}

func (rcv *BootAppInfo) UnPack() *BootAppInfoT {
	if rcv == nil { return nil }
	t := &BootAppInfoT{}
	rcv.UnPackTo(t)
	return t
}

type BootAppInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsBootAppInfo(buf []byte, offset flatbuffers.UOffsetT) *BootAppInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BootAppInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBootAppInfo(buf []byte, offset flatbuffers.UOffsetT) *BootAppInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BootAppInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *BootAppInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BootAppInfo) Table() flatbuffers.Table {
	return rcv._tab
}

/// Name of the boot application
func (rcv *BootAppInfo) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Name of the boot application
/// Timestamp (Filetime) from the last time this application was modified
func (rcv *BootAppInfo) LastModified() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// Timestamp (Filetime) from the last time this application was modified
func (rcv *BootAppInfo) MutateLastModified(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func BootAppInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func BootAppInfoAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func BootAppInfoAddLastModified(builder *flatbuffers.Builder, lastModified uint64) {
	builder.PrependUint64Slot(1, lastModified, 0)
}
func BootAppInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
