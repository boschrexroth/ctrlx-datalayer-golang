// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// probe measuring type
type ShotType int8

const (
	///< measuring type undefined
	ShotTypeUNDEFINE    ShotType = 0
	///< continuous measurement
	ShotTypeCONTINUOUS  ShotType = 1
	///< single-shot measurement
	ShotTypeSINGLE_SHOT ShotType = 2
)

var EnumNamesShotType = map[ShotType]string{
	ShotTypeUNDEFINE:    "UNDEFINE",
	ShotTypeCONTINUOUS:  "CONTINUOUS",
	ShotTypeSINGLE_SHOT: "SINGLE_SHOT",
}

var EnumValuesShotType = map[string]ShotType{
	"UNDEFINE":    ShotTypeUNDEFINE,
	"CONTINUOUS":  ShotTypeCONTINUOUS,
	"SINGLE_SHOT": ShotTypeSINGLE_SHOT,
}

func (v ShotType) String() string {
	if s, ok := EnumNamesShotType[v]; ok {
		return s
	}
	return "ShotType(" + strconv.FormatInt(int64(v), 10) + ")"
}
