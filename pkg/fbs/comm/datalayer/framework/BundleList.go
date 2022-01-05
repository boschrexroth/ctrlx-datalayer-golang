// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package framework

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BundleList struct {
	_tab flatbuffers.Table
}

func GetRootAsBundleList(buf []byte, offset flatbuffers.UOffsetT) *BundleList {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BundleList{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBundleList(buf []byte, offset flatbuffers.UOffsetT) *BundleList {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BundleList{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *BundleList) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BundleList) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BundleList) Bundles(obj *Bundle, j int) bool {
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

func (rcv *BundleList) BundlesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func BundleListStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func BundleListAddBundles(builder *flatbuffers.Builder, bundles flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(bundles), 0)
}
func BundleListStartBundlesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BundleListEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
