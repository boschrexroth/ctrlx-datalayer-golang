// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import "strconv"

/// Type of watchdog of the task
type CurrentType int8

const (
	/// Task will be monitored regarding the consumed cycle time
	CurrentTypeCYCLE CurrentType = 0
	/// No watchdog used
	CurrentTypeNONE  CurrentType = 1
)

var EnumNamesCurrentType = map[CurrentType]string{
	CurrentTypeCYCLE: "CYCLE",
	CurrentTypeNONE:  "NONE",
}

var EnumValuesCurrentType = map[string]CurrentType{
	"CYCLE": CurrentTypeCYCLE,
	"NONE":  CurrentTypeNONE,
}

func (v CurrentType) String() string {
	if s, ok := EnumNamesCurrentType[v]; ok {
		return s
	}
	return "CurrentType(" + strconv.FormatInt(int64(v), 10) + ")"
}
