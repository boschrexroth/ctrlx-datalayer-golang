// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// Meaning for commanded 2d circle
type CircleMeaning2d int8

const (
	/// radius of the circle commanded
	CircleMeaning2dRADIUS CircleMeaning2d = 0
	/// center point of the circle X-coordinate
	CircleMeaning2dCP_X   CircleMeaning2d = 1
	/// center point of the circle Y-coordinate
	CircleMeaning2dCP_Y   CircleMeaning2d = 2
	/// center point of the circle Z-coordinate
	CircleMeaning2dCP_Z   CircleMeaning2d = 3
	/// circle revolutions
	CircleMeaning2dREV    CircleMeaning2d = 4
)

var EnumNamesCircleMeaning2d = map[CircleMeaning2d]string{
	CircleMeaning2dRADIUS: "RADIUS",
	CircleMeaning2dCP_X:   "CP_X",
	CircleMeaning2dCP_Y:   "CP_Y",
	CircleMeaning2dCP_Z:   "CP_Z",
	CircleMeaning2dREV:    "REV",
}

var EnumValuesCircleMeaning2d = map[string]CircleMeaning2d{
	"RADIUS": CircleMeaning2dRADIUS,
	"CP_X":   CircleMeaning2dCP_X,
	"CP_Y":   CircleMeaning2dCP_Y,
	"CP_Z":   CircleMeaning2dCP_Z,
	"REV":    CircleMeaning2dREV,
}

func (v CircleMeaning2d) String() string {
	if s, ok := EnumNamesCircleMeaning2d[v]; ok {
		return s
	}
	return "CircleMeaning2d(" + strconv.FormatInt(int64(v), 10) + ")"
}
