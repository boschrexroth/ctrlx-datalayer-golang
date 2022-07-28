// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs2

import "strconv"

type CurrentState int8

const (
	CurrentStateOPERATING CurrentState = 0
	CurrentStateSETUP     CurrentState = 1
	CurrentStateSERVICE   CurrentState = 2
	CurrentStateEXIT      CurrentState = 3
)

var EnumNamesCurrentState = map[CurrentState]string{
	CurrentStateOPERATING: "OPERATING",
	CurrentStateSETUP:     "SETUP",
	CurrentStateSERVICE:   "SERVICE",
	CurrentStateEXIT:      "EXIT",
}

var EnumValuesCurrentState = map[string]CurrentState{
	"OPERATING": CurrentStateOPERATING,
	"SETUP":     CurrentStateSETUP,
	"SERVICE":   CurrentStateSERVICE,
	"EXIT":      CurrentStateEXIT,
}

func (v CurrentState) String() string {
	if s, ok := EnumNamesCurrentState[v]; ok {
		return s
	}
	return "CurrentState(" + strconv.FormatInt(int64(v), 10) + ")"
}
