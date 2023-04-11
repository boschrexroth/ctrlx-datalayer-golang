// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package server

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BuildInfoTypeT struct {
	BuildDate uint64
	BuildNumber string
	ManufacturerName string
	ProductName string
	ProductUri string
	SoftwareVersion string
}

func (t *BuildInfoTypeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	buildNumberOffset := builder.CreateString(t.BuildNumber)
	manufacturerNameOffset := builder.CreateString(t.ManufacturerName)
	productNameOffset := builder.CreateString(t.ProductName)
	productUriOffset := builder.CreateString(t.ProductUri)
	softwareVersionOffset := builder.CreateString(t.SoftwareVersion)
	BuildInfoTypeStart(builder)
	BuildInfoTypeAddBuildDate(builder, t.BuildDate)
	BuildInfoTypeAddBuildNumber(builder, buildNumberOffset)
	BuildInfoTypeAddManufacturerName(builder, manufacturerNameOffset)
	BuildInfoTypeAddProductName(builder, productNameOffset)
	BuildInfoTypeAddProductUri(builder, productUriOffset)
	BuildInfoTypeAddSoftwareVersion(builder, softwareVersionOffset)
	return BuildInfoTypeEnd(builder)
}

func (rcv *BuildInfoType) UnPackTo(t *BuildInfoTypeT) {
	t.BuildDate = rcv.BuildDate()
	t.BuildNumber = string(rcv.BuildNumber())
	t.ManufacturerName = string(rcv.ManufacturerName())
	t.ProductName = string(rcv.ProductName())
	t.ProductUri = string(rcv.ProductUri())
	t.SoftwareVersion = string(rcv.SoftwareVersion())
}

func (rcv *BuildInfoType) UnPack() *BuildInfoTypeT {
	if rcv == nil { return nil }
	t := &BuildInfoTypeT{}
	rcv.UnPackTo(t)
	return t
}

type BuildInfoType struct {
	_tab flatbuffers.Table
}

func GetRootAsBuildInfoType(buf []byte, offset flatbuffers.UOffsetT) *BuildInfoType {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BuildInfoType{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsBuildInfoType(buf []byte, offset flatbuffers.UOffsetT) *BuildInfoType {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &BuildInfoType{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *BuildInfoType) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BuildInfoType) Table() flatbuffers.Table {
	return rcv._tab
}

/// The build date of the OPC UA Server
func (rcv *BuildInfoType) BuildDate() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

/// The build date of the OPC UA Server
func (rcv *BuildInfoType) MutateBuildDate(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

/// The build number of the OPC UA Server
func (rcv *BuildInfoType) BuildNumber() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The build number of the OPC UA Server
/// The manufacturer name of the OPC UA Server
func (rcv *BuildInfoType) ManufacturerName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The manufacturer name of the OPC UA Server
/// The product name of the OPC UA Server
func (rcv *BuildInfoType) ProductName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The product name of the OPC UA Server
/// The product uri of the OPC UA Server
func (rcv *BuildInfoType) ProductUri() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The product uri of the OPC UA Server
/// The software version of the OPC UA Server
func (rcv *BuildInfoType) SoftwareVersion() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// The software version of the OPC UA Server
func BuildInfoTypeStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func BuildInfoTypeAddBuildDate(builder *flatbuffers.Builder, buildDate uint64) {
	builder.PrependUint64Slot(0, buildDate, 0)
}
func BuildInfoTypeAddBuildNumber(builder *flatbuffers.Builder, buildNumber flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(buildNumber), 0)
}
func BuildInfoTypeAddManufacturerName(builder *flatbuffers.Builder, manufacturerName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(manufacturerName), 0)
}
func BuildInfoTypeAddProductName(builder *flatbuffers.Builder, productName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(productName), 0)
}
func BuildInfoTypeAddProductUri(builder *flatbuffers.Builder, productUri flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(productUri), 0)
}
func BuildInfoTypeAddSoftwareVersion(builder *flatbuffers.Builder, softwareVersion flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(softwareVersion), 0)
}
func BuildInfoTypeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
