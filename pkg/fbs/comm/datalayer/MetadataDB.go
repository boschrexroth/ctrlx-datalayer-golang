// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MetadataDBT struct {
	Address string
	Childs []*MetadataDBT
	Asterisk *MetadataDBT
	Metadata *MetadataT
}

func (t *MetadataDBT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	addressOffset := builder.CreateString(t.Address)
	childsOffset := flatbuffers.UOffsetT(0)
	if t.Childs != nil {
		childsLength := len(t.Childs)
		childsOffsets := make([]flatbuffers.UOffsetT, childsLength)
		for j := 0; j < childsLength; j++ {
			childsOffsets[j] = t.Childs[j].Pack(builder)
		}
		MetadataDBStartChildsVector(builder, childsLength)
		for j := childsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(childsOffsets[j])
		}
		childsOffset = builder.EndVector(childsLength)
	}
	asteriskOffset := t.Asterisk.Pack(builder)
	metadataOffset := t.Metadata.Pack(builder)
	MetadataDBStart(builder)
	MetadataDBAddAddress(builder, addressOffset)
	MetadataDBAddChilds(builder, childsOffset)
	MetadataDBAddAsterisk(builder, asteriskOffset)
	MetadataDBAddMetadata(builder, metadataOffset)
	return MetadataDBEnd(builder)
}

func (rcv *MetadataDB) UnPackTo(t *MetadataDBT) {
	t.Address = string(rcv.Address())
	childsLength := rcv.ChildsLength()
	t.Childs = make([]*MetadataDBT, childsLength)
	for j := 0; j < childsLength; j++ {
		x := MetadataDB{}
		rcv.Childs(&x, j)
		t.Childs[j] = x.UnPack()
	}
	t.Asterisk = rcv.Asterisk(nil).UnPack()
	t.Metadata = rcv.Metadata(nil).UnPack()
}

func (rcv *MetadataDB) UnPack() *MetadataDBT {
	if rcv == nil { return nil }
	t := &MetadataDBT{}
	rcv.UnPackTo(t)
	return t
}

type MetadataDB struct {
	_tab flatbuffers.Table
}

func GetRootAsMetadataDB(buf []byte, offset flatbuffers.UOffsetT) *MetadataDB {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MetadataDB{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMetadataDB(buf []byte, offset flatbuffers.UOffsetT) *MetadataDB {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MetadataDB{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MetadataDB) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MetadataDB) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MetadataDB) Address() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MetadataDB) Childs(obj *MetadataDB, j int) bool {
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

func (rcv *MetadataDB) ChildsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MetadataDB) Asterisk(obj *MetadataDB) *MetadataDB {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MetadataDB)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *MetadataDB) Metadata(obj *Metadata) *Metadata {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Metadata)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func MetadataDBStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func MetadataDBAddAddress(builder *flatbuffers.Builder, address flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(address), 0)
}
func MetadataDBAddChilds(builder *flatbuffers.Builder, childs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(childs), 0)
}
func MetadataDBStartChildsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MetadataDBAddAsterisk(builder *flatbuffers.Builder, asterisk flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(asterisk), 0)
}
func MetadataDBAddMetadata(builder *flatbuffers.Builder, metadata flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(metadata), 0)
}
func MetadataDBEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
