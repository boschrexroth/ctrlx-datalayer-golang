// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type variableType int8

const (
	/// unknown type
	variableTypeUNKNOWN      variableType = -1
	/// base type
	variableTypeBASE         variableType = 0
	/// velocity type
	variableTypeVELOCITY     variableType = 1
	/// homing type
	variableTypeHOMING       variableType = 2
	/// probe type
	variableTypePROBE        variableType = 3
	/// torque/force type
	variableTypeTORQUE_FORCE variableType = 4
)

var EnumNamesvariableType = map[variableType]string{
	variableTypeUNKNOWN:      "UNKNOWN",
	variableTypeBASE:         "BASE",
	variableTypeVELOCITY:     "VELOCITY",
	variableTypeHOMING:       "HOMING",
	variableTypePROBE:        "PROBE",
	variableTypeTORQUE_FORCE: "TORQUE_FORCE",
}

var EnumValuesvariableType = map[string]variableType{
	"UNKNOWN":      variableTypeUNKNOWN,
	"BASE":         variableTypeBASE,
	"VELOCITY":     variableTypeVELOCITY,
	"HOMING":       variableTypeHOMING,
	"PROBE":        variableTypePROBE,
	"TORQUE_FORCE": variableTypeTORQUE_FORCE,
}

func (v variableType) String() string {
	if s, ok := EnumNamesvariableType[v]; ok {
		return s
	}
	return "variableType(" + strconv.FormatInt(int64(v), 10) + ")"
}
