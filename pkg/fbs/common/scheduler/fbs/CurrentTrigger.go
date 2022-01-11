// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import "strconv"

type CurrentTrigger int8

const (
	CurrentTriggerTIMER     CurrentTrigger = 0
	CurrentTriggerINTERRUPT CurrentTrigger = 1
)

var EnumNamesCurrentTrigger = map[CurrentTrigger]string{
	CurrentTriggerTIMER:     "TIMER",
	CurrentTriggerINTERRUPT: "INTERRUPT",
}

var EnumValuesCurrentTrigger = map[string]CurrentTrigger{
	"TIMER":     CurrentTriggerTIMER,
	"INTERRUPT": CurrentTriggerINTERRUPT,
}

func (v CurrentTrigger) String() string {
	if s, ok := EnumNamesCurrentTrigger[v]; ok {
		return s
	}
	return "CurrentTrigger(" + strconv.FormatInt(int64(v), 10) + ")"
}