// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// information of a single acquired license
type SingleLicense struct {
	_tab flatbuffers.Table
}

func GetRootAsSingleLicense(buf []byte, offset flatbuffers.UOffsetT) *SingleLicense {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SingleLicense{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSingleLicense(buf []byte, offset flatbuffers.UOffsetT) *SingleLicense {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SingleLicense{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SingleLicense) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SingleLicense) Table() flatbuffers.Table {
	return rcv._tab
}

/// name (material number) of the license
func (rcv *SingleLicense) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// name (material number) of the license
/// version number of the license
func (rcv *SingleLicense) Version() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// version number of the license
func SingleLicenseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SingleLicenseAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func SingleLicenseAddVersion(builder *flatbuffers.Builder, version flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(version), 0)
}
func SingleLicenseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}