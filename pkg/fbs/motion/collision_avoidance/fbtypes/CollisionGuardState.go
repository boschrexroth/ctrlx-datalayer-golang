// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// Possible collision guard states
type CollisionGuardState int8

const (
	CollisionGuardStateMONITORING       CollisionGuardState = 0
	CollisionGuardStateTARGETS_CROSSING CollisionGuardState = 1
	CollisionGuardStateSYNC_ADVICE      CollisionGuardState = 2
	CollisionGuardStateSLOWDOWN_ADVICE  CollisionGuardState = 3
	CollisionGuardStateSTOP             CollisionGuardState = 4
	CollisionGuardStateEMERGENCY_STOP   CollisionGuardState = 5
	CollisionGuardStateUNKNOWN          CollisionGuardState = 6
)

var EnumNamesCollisionGuardState = map[CollisionGuardState]string{
	CollisionGuardStateMONITORING:       "MONITORING",
	CollisionGuardStateTARGETS_CROSSING: "TARGETS_CROSSING",
	CollisionGuardStateSYNC_ADVICE:      "SYNC_ADVICE",
	CollisionGuardStateSLOWDOWN_ADVICE:  "SLOWDOWN_ADVICE",
	CollisionGuardStateSTOP:             "STOP",
	CollisionGuardStateEMERGENCY_STOP:   "EMERGENCY_STOP",
	CollisionGuardStateUNKNOWN:          "UNKNOWN",
}

var EnumValuesCollisionGuardState = map[string]CollisionGuardState{
	"MONITORING":       CollisionGuardStateMONITORING,
	"TARGETS_CROSSING": CollisionGuardStateTARGETS_CROSSING,
	"SYNC_ADVICE":      CollisionGuardStateSYNC_ADVICE,
	"SLOWDOWN_ADVICE":  CollisionGuardStateSLOWDOWN_ADVICE,
	"STOP":             CollisionGuardStateSTOP,
	"EMERGENCY_STOP":   CollisionGuardStateEMERGENCY_STOP,
	"UNKNOWN":          CollisionGuardStateUNKNOWN,
}

func (v CollisionGuardState) String() string {
	if s, ok := EnumNamesCollisionGuardState[v]; ok {
		return s
	}
	return "CollisionGuardState(" + strconv.FormatInt(int64(v), 10) + ")"
}
