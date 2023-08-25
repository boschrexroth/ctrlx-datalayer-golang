// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	"strconv"
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Special extensions to modify Scheduler and system behavior e.g. activation of DEBUG mode
type Controls byte

const (
	ControlsNONE  Controls = 0
	/// Deactivation of hardware watchdog and task watchdogs to enabling debugging of machine
	ControlsDebug Controls = 1
)

var EnumNamesControls = map[Controls]string{
	ControlsNONE:  "NONE",
	ControlsDebug: "Debug",
}

var EnumValuesControls = map[string]Controls{
	"NONE":  ControlsNONE,
	"Debug": ControlsDebug,
}

func (v Controls) String() string {
	if s, ok := EnumNamesControls[v]; ok {
		return s
	}
	return "Controls(" + strconv.FormatInt(int64(v), 10) + ")"
}

type ControlsT struct {
	Type Controls
	Value interface{}
}

func (t *ControlsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case ControlsDebug:
		return t.Value.(*DebugT).Pack(builder)
	}
	return 0
}

func (rcv Controls) UnPack(table flatbuffers.Table) *ControlsT {
	switch rcv {
	case ControlsDebug:
		var x Debug
		x.Init(table.Bytes, table.Pos)
		return &ControlsT{ Type: ControlsDebug, Value: x.UnPack() }
	}
	return nil
}
