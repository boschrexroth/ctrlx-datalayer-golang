// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// braking distance type
type BrakingDistanceType int8

const (
	/// invalid braking distance type
	BrakingDistanceTypeUNDEFINED BrakingDistanceType = 0
	/// emergency stop distance
	BrakingDistanceTypeESTOP     BrakingDistanceType = 1
	/// soft stop distance (corresponding to setting override to 0)
	BrakingDistanceTypeSOFT_STOP BrakingDistanceType = 2
)

var EnumNamesBrakingDistanceType = map[BrakingDistanceType]string{
	BrakingDistanceTypeUNDEFINED: "UNDEFINED",
	BrakingDistanceTypeESTOP:     "ESTOP",
	BrakingDistanceTypeSOFT_STOP: "SOFT_STOP",
}

var EnumValuesBrakingDistanceType = map[string]BrakingDistanceType{
	"UNDEFINED": BrakingDistanceTypeUNDEFINED,
	"ESTOP":     BrakingDistanceTypeESTOP,
	"SOFT_STOP": BrakingDistanceTypeSOFT_STOP,
}

func (v BrakingDistanceType) String() string {
	if s, ok := EnumNamesBrakingDistanceType[v]; ok {
		return s
	}
	return "BrakingDistanceType(" + strconv.FormatInt(int64(v), 10) + ")"
}
