// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type BrakingDistanceTypeReq int8

const (
	BrakingDistanceTypeReqUNDEFINED BrakingDistanceTypeReq = 0
	BrakingDistanceTypeReqESTOP     BrakingDistanceTypeReq = 1
	BrakingDistanceTypeReqSOFT_STOP BrakingDistanceTypeReq = 2
)

var EnumNamesBrakingDistanceTypeReq = map[BrakingDistanceTypeReq]string{
	BrakingDistanceTypeReqUNDEFINED: "UNDEFINED",
	BrakingDistanceTypeReqESTOP:     "ESTOP",
	BrakingDistanceTypeReqSOFT_STOP: "SOFT_STOP",
}

var EnumValuesBrakingDistanceTypeReq = map[string]BrakingDistanceTypeReq{
	"UNDEFINED": BrakingDistanceTypeReqUNDEFINED,
	"ESTOP":     BrakingDistanceTypeReqESTOP,
	"SOFT_STOP": BrakingDistanceTypeReqSOFT_STOP,
}

func (v BrakingDistanceTypeReq) String() string {
	if s, ok := EnumNamesBrakingDistanceTypeReq[v]; ok {
		return s
	}
	return "BrakingDistanceTypeReq(" + strconv.FormatInt(int64(v), 10) + ")"
}
