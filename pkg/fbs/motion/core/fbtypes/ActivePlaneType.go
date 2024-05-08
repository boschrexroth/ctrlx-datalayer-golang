// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type ActivePlaneType int8

const (
	/// XY plane with Z perpendicular
	ActivePlaneTypeXY_PLANE ActivePlaneType = 0
	/// YZ plane with X perpendicular
	ActivePlaneTypeYZ_PLANE ActivePlaneType = 1
	/// ZX plane with Y perpendicular
	ActivePlaneTypeZX_PLANE ActivePlaneType = 2
)

var EnumNamesActivePlaneType = map[ActivePlaneType]string{
	ActivePlaneTypeXY_PLANE: "XY_PLANE",
	ActivePlaneTypeYZ_PLANE: "YZ_PLANE",
	ActivePlaneTypeZX_PLANE: "ZX_PLANE",
}

var EnumValuesActivePlaneType = map[string]ActivePlaneType{
	"XY_PLANE": ActivePlaneTypeXY_PLANE,
	"YZ_PLANE": ActivePlaneTypeYZ_PLANE,
	"ZX_PLANE": ActivePlaneTypeZX_PLANE,
}

func (v ActivePlaneType) String() string {
	if s, ok := EnumNamesActivePlaneType[v]; ok {
		return s
	}
	return "ActivePlaneType(" + strconv.FormatInt(int64(v), 10) + ")"
}