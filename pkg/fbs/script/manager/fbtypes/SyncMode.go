// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// sync mode (must fit the enum class motion::sync::cmd::SyncMode)
type SyncMode int8

const (
	/// The slave is moved relative and the calculation pipeline is initialized new.
	SyncModeSYNC_RELATIVE            SyncMode = 0
	/// The slave is moved relative and the calculation pipeline keeps its state.
	SyncModeSYNC_RELATIVE_KEEP_STATE SyncMode = 1
	/// The slave is moved absolute and the calculation pipeline is initialized new.
	SyncModeSYNC_ABSOLUTE            SyncMode = 2
	/// The slave is moved absolute and the calculation pipeline keeps its state.
	SyncModeSYNC_ABSOLUTE_KEEP_STATE SyncMode = 3
)

var EnumNamesSyncMode = map[SyncMode]string{
	SyncModeSYNC_RELATIVE:            "SYNC_RELATIVE",
	SyncModeSYNC_RELATIVE_KEEP_STATE: "SYNC_RELATIVE_KEEP_STATE",
	SyncModeSYNC_ABSOLUTE:            "SYNC_ABSOLUTE",
	SyncModeSYNC_ABSOLUTE_KEEP_STATE: "SYNC_ABSOLUTE_KEEP_STATE",
}

var EnumValuesSyncMode = map[string]SyncMode{
	"SYNC_RELATIVE":            SyncModeSYNC_RELATIVE,
	"SYNC_RELATIVE_KEEP_STATE": SyncModeSYNC_RELATIVE_KEEP_STATE,
	"SYNC_ABSOLUTE":            SyncModeSYNC_ABSOLUTE,
	"SYNC_ABSOLUTE_KEEP_STATE": SyncModeSYNC_ABSOLUTE_KEEP_STATE,
}

func (v SyncMode) String() string {
	if s, ok := EnumNamesSyncMode[v]; ok {
		return s
	}
	return "SyncMode(" + strconv.FormatInt(int64(v), 10) + ")"
}