// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import "strconv"

type SyncMode uint32

const (
	SyncModefreerun              SyncMode = 0
	SyncModedcmAuto              SyncMode = 1
	SyncModedcEnable             SyncMode = 2
	SyncModedcmBusShift          SyncMode = 3
	SyncModedcmMasterShift       SyncMode = 4
	SyncModedcmLinkLayerRefClock SyncMode = 5
	SyncModeunknown              SyncMode = 4294967295
)

var EnumNamesSyncMode = map[SyncMode]string{
	SyncModefreerun:              "freerun",
	SyncModedcmAuto:              "dcmAuto",
	SyncModedcEnable:             "dcEnable",
	SyncModedcmBusShift:          "dcmBusShift",
	SyncModedcmMasterShift:       "dcmMasterShift",
	SyncModedcmLinkLayerRefClock: "dcmLinkLayerRefClock",
	SyncModeunknown:              "unknown",
}

var EnumValuesSyncMode = map[string]SyncMode{
	"freerun":              SyncModefreerun,
	"dcmAuto":              SyncModedcmAuto,
	"dcEnable":             SyncModedcEnable,
	"dcmBusShift":          SyncModedcmBusShift,
	"dcmMasterShift":       SyncModedcmMasterShift,
	"dcmLinkLayerRefClock": SyncModedcmLinkLayerRefClock,
	"unknown":              SyncModeunknown,
}

func (v SyncMode) String() string {
	if s, ok := EnumNamesSyncMode[v]; ok {
		return s
	}
	return "SyncMode(" + strconv.FormatInt(int64(v), 10) + ")"
}
