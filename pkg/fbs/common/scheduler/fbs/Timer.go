// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import "strconv"

/// Types of duration measurements
type Timer int8

const (
	/// Start, stop or reset all available types of duration measurements
	TimerALL          Timer = 0
	/// Start, stop or reset duration measurement of the task including interruptions by other tasks
	TimerTOTAL        Timer = 1
	/// Start, stop or reset duration measurement of the task without interruptions by other tasks
	TimerTASK         Timer = 2
	/// Start, stop or reset duration measurement of the interruptions of the task by other tasks
	TimerOTHER        Timer = 3
	/// Start, stop or reset duration measurement of equidistance of the task from one cycle to an other
	TimerEQUIDISTANCE Timer = 4
	/// Start, stop or reset duration measurement of deviation of the task to the expected start time
	TimerDEVIATION    Timer = 5
	/// Start, stop or reset duration measurement of remaining time from the end of the task to the begin of the next cycle of it
	TimerREMAINING    Timer = 6
)

var EnumNamesTimer = map[Timer]string{
	TimerALL:          "ALL",
	TimerTOTAL:        "TOTAL",
	TimerTASK:         "TASK",
	TimerOTHER:        "OTHER",
	TimerEQUIDISTANCE: "EQUIDISTANCE",
	TimerDEVIATION:    "DEVIATION",
	TimerREMAINING:    "REMAINING",
}

var EnumValuesTimer = map[string]Timer{
	"ALL":          TimerALL,
	"TOTAL":        TimerTOTAL,
	"TASK":         TimerTASK,
	"OTHER":        TimerOTHER,
	"EQUIDISTANCE": TimerEQUIDISTANCE,
	"DEVIATION":    TimerDEVIATION,
	"REMAINING":    TimerREMAINING,
}

func (v Timer) String() string {
	if s, ok := EnumNamesTimer[v]; ok {
		return s
	}
	return "Timer(" + strconv.FormatInt(int64(v), 10) + ")"
}
