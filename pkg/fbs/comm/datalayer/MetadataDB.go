// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import (
	"bytes"
	flatbuffers "github.com/google/flatbuffers/go"
)

type MetadataDBT struct {
	Address string `json:"address"`
	Childs []*MetadataDBT `json:"childs"`
	Asterisk *MetadataDBT `json:"asterisk"`
	Metadata *MetadataT `json:"metadata"`
	Scopes []string `json:"scopes"`
}

func (t *MetadataDBT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	addressOffset := flatbuffers.UOffsetT(0)
	if t.Address != "" {
		addressOffset = builder.CreateString(t.Address)
	}
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
	scopesOffset := flatbuffers.UOffsetT(0)
	if t.Scopes != nil {
		scopesLength := len(t.Scopes)
		scopesOffsets := make([]flatbuffers.UOffsetT, scopesLength)
		for j := 0; j < scopesLength; j++ {
			scopesOffsets[j] = builder.CreateString(t.Scopes[j])
		}
		MetadataDBStartScopesVector(builder, scopesLength)
		for j := scopesLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(scopesOffsets[j])
		}
		scopesOffset = builder.EndVector(scopesLength)
	}
	MetadataDBStart(builder)
	MetadataDBAddAddress(builder, addressOffset)
	MetadataDBAddChilds(builder, childsOffset)
	MetadataDBAddAsterisk(builder, asteriskOffset)
	MetadataDBAddMetadata(builder, metadataOffset)
	MetadataDBAddScopes(builder, scopesOffset)
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
	scopesLength := rcv.ScopesLength()
	t.Scopes = make([]string, scopesLength)
	for j := 0; j < scopesLength; j++ {
		t.Scopes[j] = string(rcv.Scopes(j))
	}
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

func MetadataDBKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &MetadataDB{}
	obj2 := &MetadataDB{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return string(obj1.Address()) < string(obj2.Address())
}

func (rcv *MetadataDB) LookupByKey(key string, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	bKey := []byte(key)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &MetadataDB{}
		obj.Init(buf, tableOffset)
		comp := bytes.Compare(obj.Address(), bKey)
		if comp > 0 {
			span = middle
		} else if comp < 0 {
			middle += 1
			start += middle
			span -= middle
		} else {
			rcv.Init(buf, tableOffset)
			return true
		}
	}
	return false
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

func (rcv *MetadataDB) ChildsByKey(obj *MetadataDB, key string) bool{
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
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

func (rcv *MetadataDB) Scopes(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *MetadataDB) ScopesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MetadataDBStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
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
func MetadataDBAddScopes(builder *flatbuffers.Builder, scopes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(scopes), 0)
}
func MetadataDBStartScopesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MetadataDBEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
