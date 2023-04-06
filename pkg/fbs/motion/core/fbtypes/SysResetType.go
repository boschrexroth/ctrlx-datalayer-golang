// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// Type of the system reset request
type SysResetType int8

const (
	/// reset all motion objects, that are currently in an error state and have their error reaction finished
	SysResetTypeResetAllMotionObjectsWithError SysResetType = 0
	/// reset all motion object (not yet supported)
	SysResetTypeResetAllMotionObjects          SysResetType = 1
)

var EnumNamesSysResetType = map[SysResetType]string{
	SysResetTypeResetAllMotionObjectsWithError: "ResetAllMotionObjectsWithError",
	SysResetTypeResetAllMotionObjects:          "ResetAllMotionObjects",
}

var EnumValuesSysResetType = map[string]SysResetType{
	"ResetAllMotionObjectsWithError": SysResetTypeResetAllMotionObjectsWithError,
	"ResetAllMotionObjects":          SysResetTypeResetAllMotionObjects,
}

func (v SysResetType) String() string {
	if s, ok := EnumNamesSysResetType[v]; ok {
		return s
	}
	return "SysResetType(" + strconv.FormatInt(int64(v), 10) + ")"
}
