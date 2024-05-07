// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Configuration of the dynamic synchronisation window modulo of a single axis
type AxsCfgDynSynchronisationSyncWinModT struct {
	SyncWindowModulo float64 `json:"syncWindowModulo"`
	SyncWindowModuloUnit string `json:"syncWindowModuloUnit"`
}

func (t *AxsCfgDynSynchronisationSyncWinModT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	syncWindowModuloUnitOffset := flatbuffers.UOffsetT(0)
	if t.SyncWindowModuloUnit != "" {
		syncWindowModuloUnitOffset = builder.CreateString(t.SyncWindowModuloUnit)
	}
	AxsCfgDynSynchronisationSyncWinModStart(builder)
	AxsCfgDynSynchronisationSyncWinModAddSyncWindowModulo(builder, t.SyncWindowModulo)
	AxsCfgDynSynchronisationSyncWinModAddSyncWindowModuloUnit(builder, syncWindowModuloUnitOffset)
	return AxsCfgDynSynchronisationSyncWinModEnd(builder)
}

func (rcv *AxsCfgDynSynchronisationSyncWinMod) UnPackTo(t *AxsCfgDynSynchronisationSyncWinModT) {
	t.SyncWindowModulo = rcv.SyncWindowModulo()
	t.SyncWindowModuloUnit = string(rcv.SyncWindowModuloUnit())
}

func (rcv *AxsCfgDynSynchronisationSyncWinMod) UnPack() *AxsCfgDynSynchronisationSyncWinModT {
	if rcv == nil { return nil }
	t := &AxsCfgDynSynchronisationSyncWinModT{}
	rcv.UnPackTo(t)
	return t
}

type AxsCfgDynSynchronisationSyncWinMod struct {
	_tab flatbuffers.Table
}

func GetRootAsAxsCfgDynSynchronisationSyncWinMod(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgDynSynchronisationSyncWinMod {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AxsCfgDynSynchronisationSyncWinMod{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAxsCfgDynSynchronisationSyncWinMod(buf []byte, offset flatbuffers.UOffsetT) *AxsCfgDynSynchronisationSyncWinMod {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &AxsCfgDynSynchronisationSyncWinMod{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *AxsCfgDynSynchronisationSyncWinMod) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AxsCfgDynSynchronisationSyncWinMod) Table() flatbuffers.Table {
	return rcv._tab
}

/// Dynamic synchronisation window modulo value
func (rcv *AxsCfgDynSynchronisationSyncWinMod) SyncWindowModulo() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Dynamic synchronisation window modulo value
func (rcv *AxsCfgDynSynchronisationSyncWinMod) MutateSyncWindowModulo(n float64) bool {
	return rcv._tab.MutateFloat64Slot(4, n)
}

/// Unit of syncWindowModulo
func (rcv *AxsCfgDynSynchronisationSyncWinMod) SyncWindowModuloUnit() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Unit of syncWindowModulo
func AxsCfgDynSynchronisationSyncWinModStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func AxsCfgDynSynchronisationSyncWinModAddSyncWindowModulo(builder *flatbuffers.Builder, syncWindowModulo float64) {
	builder.PrependFloat64Slot(0, syncWindowModulo, 0.0)
}
func AxsCfgDynSynchronisationSyncWinModAddSyncWindowModuloUnit(builder *flatbuffers.Builder, syncWindowModuloUnit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(syncWindowModuloUnit), 0)
}
func AxsCfgDynSynchronisationSyncWinModEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
