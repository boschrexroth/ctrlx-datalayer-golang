// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package server

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CertificateConfigTypeT struct {
	Certificate string
	Key string
	Store uint32
	DaysValid uint32
	Algorithm string
	KeyUsage string
	KeyLength uint32
	IssuerIndex uint32
	Create bool
}

func (t *CertificateConfigTypeT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	certificateOffset := builder.CreateString(t.Certificate)
	keyOffset := builder.CreateString(t.Key)
	algorithmOffset := builder.CreateString(t.Algorithm)
	keyUsageOffset := builder.CreateString(t.KeyUsage)
	CertificateConfigTypeStart(builder)
	CertificateConfigTypeAddCertificate(builder, certificateOffset)
	CertificateConfigTypeAddKey(builder, keyOffset)
	CertificateConfigTypeAddStore(builder, t.Store)
	CertificateConfigTypeAddDaysValid(builder, t.DaysValid)
	CertificateConfigTypeAddAlgorithm(builder, algorithmOffset)
	CertificateConfigTypeAddKeyUsage(builder, keyUsageOffset)
	CertificateConfigTypeAddKeyLength(builder, t.KeyLength)
	CertificateConfigTypeAddIssuerIndex(builder, t.IssuerIndex)
	CertificateConfigTypeAddCreate(builder, t.Create)
	return CertificateConfigTypeEnd(builder)
}

func (rcv *CertificateConfigType) UnPackTo(t *CertificateConfigTypeT) {
	t.Certificate = string(rcv.Certificate())
	t.Key = string(rcv.Key())
	t.Store = rcv.Store()
	t.DaysValid = rcv.DaysValid()
	t.Algorithm = string(rcv.Algorithm())
	t.KeyUsage = string(rcv.KeyUsage())
	t.KeyLength = rcv.KeyLength()
	t.IssuerIndex = rcv.IssuerIndex()
	t.Create = rcv.Create()
}

func (rcv *CertificateConfigType) UnPack() *CertificateConfigTypeT {
	if rcv == nil { return nil }
	t := &CertificateConfigTypeT{}
	rcv.UnPackTo(t)
	return t
}

type CertificateConfigType struct {
	_tab flatbuffers.Table
}

func GetRootAsCertificateConfigType(buf []byte, offset flatbuffers.UOffsetT) *CertificateConfigType {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CertificateConfigType{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCertificateConfigType(buf []byte, offset flatbuffers.UOffsetT) *CertificateConfigType {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CertificateConfigType{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CertificateConfigType) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CertificateConfigType) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CertificateConfigType) Certificate() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CertificateConfigType) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CertificateConfigType) Store() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CertificateConfigType) MutateStore(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *CertificateConfigType) DaysValid() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CertificateConfigType) MutateDaysValid(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *CertificateConfigType) Algorithm() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CertificateConfigType) KeyUsage() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CertificateConfigType) KeyLength() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CertificateConfigType) MutateKeyLength(n uint32) bool {
	return rcv._tab.MutateUint32Slot(16, n)
}

func (rcv *CertificateConfigType) IssuerIndex() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CertificateConfigType) MutateIssuerIndex(n uint32) bool {
	return rcv._tab.MutateUint32Slot(18, n)
}

func (rcv *CertificateConfigType) Create() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *CertificateConfigType) MutateCreate(n bool) bool {
	return rcv._tab.MutateBoolSlot(20, n)
}

func CertificateConfigTypeStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
}
func CertificateConfigTypeAddCertificate(builder *flatbuffers.Builder, certificate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(certificate), 0)
}
func CertificateConfigTypeAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(key), 0)
}
func CertificateConfigTypeAddStore(builder *flatbuffers.Builder, store uint32) {
	builder.PrependUint32Slot(2, store, 0)
}
func CertificateConfigTypeAddDaysValid(builder *flatbuffers.Builder, daysValid uint32) {
	builder.PrependUint32Slot(3, daysValid, 0)
}
func CertificateConfigTypeAddAlgorithm(builder *flatbuffers.Builder, algorithm flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(algorithm), 0)
}
func CertificateConfigTypeAddKeyUsage(builder *flatbuffers.Builder, keyUsage flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(keyUsage), 0)
}
func CertificateConfigTypeAddKeyLength(builder *flatbuffers.Builder, keyLength uint32) {
	builder.PrependUint32Slot(6, keyLength, 0)
}
func CertificateConfigTypeAddIssuerIndex(builder *flatbuffers.Builder, issuerIndex uint32) {
	builder.PrependUint32Slot(7, issuerIndex, 0)
}
func CertificateConfigTypeAddCreate(builder *flatbuffers.Builder, create bool) {
	builder.PrependBoolSlot(8, create, false)
}
func CertificateConfigTypeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
