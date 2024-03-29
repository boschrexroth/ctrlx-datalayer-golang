// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import "strconv"

/// Types of system operation state
type CurrentState int8

const (
	/// 
	CurrentStateINIT              CurrentState = 0
	///
	CurrentStateSTOP              CurrentState = 1
	///
	CurrentStateRUN               CurrentState = 2
	/// At least one error is pending 
	CurrentStateERROR             CurrentState = 3
	///
	CurrentStateSYSERROR          CurrentState = 4
	///
	CurrentStateSHUTDOWN          CurrentState = 5
	/// At least one warning is pending 
	CurrentStateWARNING           CurrentState = 6
	/// User tasks and callables are not executed, action with influence on the real-time behavior can be performed
	CurrentStateSERVICE           CurrentState = 7
	/// User tasks and callables will executed, changes to the configuration are permitted, action with influence on the real-time behavior (like app installation) are not allowed
	CurrentStateSETUP             CurrentState = 8
	/// At least one but not all user tasks and callables are in OPERATING state
	CurrentStatePARTIAL_OPERATING CurrentState = 9
	/// User tasks and callables will executed, changes to the configuration might not be permitted, action with influence on the real-time behavior (like app installation) are not allowed
	CurrentStateOPERATING         CurrentState = 10
)

var EnumNamesCurrentState = map[CurrentState]string{
	CurrentStateINIT:              "INIT",
	CurrentStateSTOP:              "STOP",
	CurrentStateRUN:               "RUN",
	CurrentStateERROR:             "ERROR",
	CurrentStateSYSERROR:          "SYSERROR",
	CurrentStateSHUTDOWN:          "SHUTDOWN",
	CurrentStateWARNING:           "WARNING",
	CurrentStateSERVICE:           "SERVICE",
	CurrentStateSETUP:             "SETUP",
	CurrentStatePARTIAL_OPERATING: "PARTIAL_OPERATING",
	CurrentStateOPERATING:         "OPERATING",
}

var EnumValuesCurrentState = map[string]CurrentState{
	"INIT":              CurrentStateINIT,
	"STOP":              CurrentStateSTOP,
	"RUN":               CurrentStateRUN,
	"ERROR":             CurrentStateERROR,
	"SYSERROR":          CurrentStateSYSERROR,
	"SHUTDOWN":          CurrentStateSHUTDOWN,
	"WARNING":           CurrentStateWARNING,
	"SERVICE":           CurrentStateSERVICE,
	"SETUP":             CurrentStateSETUP,
	"PARTIAL_OPERATING": CurrentStatePARTIAL_OPERATING,
	"OPERATING":         CurrentStateOPERATING,
}

func (v CurrentState) String() string {
	if s, ok := EnumNamesCurrentState[v]; ok {
		return s
	}
	return "CurrentState(" + strconv.FormatInt(int64(v), 10) + ")"
}
