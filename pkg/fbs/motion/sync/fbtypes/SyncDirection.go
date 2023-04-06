// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// sync direction mode (must fit the enum class motion::sync::cmd::SyncDirection)
type SyncDirection int8

const (
	/// shortest path synchronization way
	SyncDirectionSYNC_SHORTEST_PATH SyncDirection = 0
	/// catch up synchronization way
	SyncDirectionSYNC_CATCH_UP      SyncDirection = 1
	/// slow down synchronization way
	SyncDirectionSYNC_SLOW_DOWN     SyncDirection = 2
)

var EnumNamesSyncDirection = map[SyncDirection]string{
	SyncDirectionSYNC_SHORTEST_PATH: "SYNC_SHORTEST_PATH",
	SyncDirectionSYNC_CATCH_UP:      "SYNC_CATCH_UP",
	SyncDirectionSYNC_SLOW_DOWN:     "SYNC_SLOW_DOWN",
}

var EnumValuesSyncDirection = map[string]SyncDirection{
	"SYNC_SHORTEST_PATH": SyncDirectionSYNC_SHORTEST_PATH,
	"SYNC_CATCH_UP":      SyncDirectionSYNC_CATCH_UP,
	"SYNC_SLOW_DOWN":     SyncDirectionSYNC_SLOW_DOWN,
}

func (v SyncDirection) String() string {
	if s, ok := EnumNamesSyncDirection[v]; ok {
		return s
	}
	return "SyncDirection(" + strconv.FormatInt(int64(v), 10) + ")"
}
