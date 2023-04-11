// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AllowedOperationsT struct {
	Read bool
	Write bool
	Create bool
	Delete bool
	Browse bool
}

func (t *AllowedOperationsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	AllowedOperationsStart(builder)
	AllowedOperationsAddRead(builder, t.Read)
	AllowedOperationsAddWrite(builder, t.Write)
	AllowedOperationsAddCreate(builder, t.Create)
	AllowedOperationsAddDelete(builder, t.Delete)
	AllowedOperationsAddBrowse(builder, t.Browse)
	return AllowedOperationsEnd(builder)
}

func (rcv *AllowedOperations) UnPackTo(t *AllowedOperationsT) {
	t.Read = rcv.Read()
	t.Write = rcv.Write()
	t.Create = rcv.Create()
	t.Delete = rcv.Delete()
	t.Browse = rcv.Browse()
}

func (rcv *AllowedOperations) UnPack() *AllowedOperationsT {
	if rcv == nil { return nil }
	t := &AllowedOperationsT{}
	rcv.UnPackTo(t)
	return t
}

type AllowedOperations struct {
	_tab flatbuffers.Table
}

func GetRootAsAllowedOperations(buf []byte, offset flatbuffers.UOffsetT) *AllowedOperations {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AllowedOperations{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAllowedOperations(buf []byte, offset flatbuffers.UOffsetT) *AllowedOperations {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AllowedOperations{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AllowedOperations) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AllowedOperations) Table() flatbuffers.Table {
	return rcv._tab
}

/// get
func (rcv *AllowedOperations) Read() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// get
func (rcv *AllowedOperations) MutateRead(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

/// put
func (rcv *AllowedOperations) Write() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// put
func (rcv *AllowedOperations) MutateWrite(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

/// post
func (rcv *AllowedOperations) Create() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// post
func (rcv *AllowedOperations) MutateCreate(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

/// delete
func (rcv *AllowedOperations) Delete() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// delete
func (rcv *AllowedOperations) MutateDelete(n bool) bool {
	return rcv._tab.MutateBoolSlot(10, n)
}

/// browse
func (rcv *AllowedOperations) Browse() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

/// browse
func (rcv *AllowedOperations) MutateBrowse(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

func AllowedOperationsStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func AllowedOperationsAddRead(builder *flatbuffers.Builder, read bool) {
	builder.PrependBoolSlot(0, read, false)
}
func AllowedOperationsAddWrite(builder *flatbuffers.Builder, write bool) {
	builder.PrependBoolSlot(1, write, false)
}
func AllowedOperationsAddCreate(builder *flatbuffers.Builder, create bool) {
	builder.PrependBoolSlot(2, create, false)
}
func AllowedOperationsAddDelete(builder *flatbuffers.Builder, delete bool) {
	builder.PrependBoolSlot(3, delete, false)
}
func AllowedOperationsAddBrowse(builder *flatbuffers.Builder, browse bool) {
	builder.PrependBoolSlot(4, browse, true)
}
func AllowedOperationsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
