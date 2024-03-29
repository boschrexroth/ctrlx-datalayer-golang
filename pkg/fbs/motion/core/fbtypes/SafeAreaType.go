// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// type of the safe area
type SafeAreaType int8

const (
	/// its a safe area
	SafeAreaTypeSAFE_AREA      SafeAreaType = 0
	/// its an exclusive work area
	SafeAreaTypeWORK_AREA_EXCL SafeAreaType = 1
	/// its an inclusive work area
	SafeAreaTypeWORK_AREA_INCL SafeAreaType = 2
)

var EnumNamesSafeAreaType = map[SafeAreaType]string{
	SafeAreaTypeSAFE_AREA:      "SAFE_AREA",
	SafeAreaTypeWORK_AREA_EXCL: "WORK_AREA_EXCL",
	SafeAreaTypeWORK_AREA_INCL: "WORK_AREA_INCL",
}

var EnumValuesSafeAreaType = map[string]SafeAreaType{
	"SAFE_AREA":      SafeAreaTypeSAFE_AREA,
	"WORK_AREA_EXCL": SafeAreaTypeWORK_AREA_EXCL,
	"WORK_AREA_INCL": SafeAreaTypeWORK_AREA_INCL,
}

func (v SafeAreaType) String() string {
	if s, ok := EnumNamesSafeAreaType[v]; ok {
		return s
	}
	return "SafeAreaType(" + strconv.FormatInt(int64(v), 10) + ")"
}
