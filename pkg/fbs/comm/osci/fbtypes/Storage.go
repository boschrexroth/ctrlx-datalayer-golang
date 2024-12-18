// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StorageT struct {
	InfluxDbcfg *InfluxDBT `json:"influxDBCfg"`
	Local bool `json:"local"`
}

func (t *StorageT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	influxDbcfgOffset := t.InfluxDbcfg.Pack(builder)
	StorageStart(builder)
	StorageAddInfluxDbcfg(builder, influxDbcfgOffset)
	StorageAddLocal(builder, t.Local)
	return StorageEnd(builder)
}

func (rcv *Storage) UnPackTo(t *StorageT) {
	t.InfluxDbcfg = rcv.InfluxDbcfg(nil).UnPack()
	t.Local = rcv.Local()
}

func (rcv *Storage) UnPack() *StorageT {
	if rcv == nil { return nil }
	t := &StorageT{}
	rcv.UnPackTo(t)
	return t
}

type Storage struct {
	_tab flatbuffers.Table
}

func GetRootAsStorage(buf []byte, offset flatbuffers.UOffsetT) *Storage {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Storage{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStorage(buf []byte, offset flatbuffers.UOffsetT) *Storage {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Storage{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Storage) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Storage) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Storage) InfluxDbcfg(obj *InfluxDB) *InfluxDB {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(InfluxDB)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Storage) Local() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Storage) MutateLocal(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func StorageStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func StorageAddInfluxDbcfg(builder *flatbuffers.Builder, influxDbcfg flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(influxDbcfg), 0)
}
func StorageAddLocal(builder *flatbuffers.Builder, local bool) {
	builder.PrependBoolSlot(1, local, false)
}
func StorageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
