// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import "strconv"

type CurrentState int8

const (
	CurrentStateRUN    CurrentState = 0
	CurrentStateCONFIG CurrentState = 1
	CurrentStateINIT   CurrentState = 2
	CurrentStateEXIT   CurrentState = 3
)

var EnumNamesCurrentState = map[CurrentState]string{
	CurrentStateRUN:    "RUN",
	CurrentStateCONFIG: "CONFIG",
	CurrentStateINIT:   "INIT",
	CurrentStateEXIT:   "EXIT",
}

var EnumValuesCurrentState = map[string]CurrentState{
	"RUN":    CurrentStateRUN,
	"CONFIG": CurrentStateCONFIG,
	"INIT":   CurrentStateINIT,
	"EXIT":   CurrentStateEXIT,
}

func (v CurrentState) String() string {
	if s, ok := EnumNamesCurrentState[v]; ok {
		return s
	}
	return "CurrentState(" + strconv.FormatInt(int64(v), 10) + ")"
}
