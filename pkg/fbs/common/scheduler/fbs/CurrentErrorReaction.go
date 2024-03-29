// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import "strconv"

/// Error reaction in case of not all configured components are available on startup
type CurrentErrorReaction int8

const (
	/// The Scheduler aborts further startup
	CurrentErrorReactionSTOP     CurrentErrorReaction = 0
	/// The Scheduler continues the startup without the missing components
	CurrentErrorReactionCONTINUE CurrentErrorReaction = 1
)

var EnumNamesCurrentErrorReaction = map[CurrentErrorReaction]string{
	CurrentErrorReactionSTOP:     "STOP",
	CurrentErrorReactionCONTINUE: "CONTINUE",
}

var EnumValuesCurrentErrorReaction = map[string]CurrentErrorReaction{
	"STOP":     CurrentErrorReactionSTOP,
	"CONTINUE": CurrentErrorReactionCONTINUE,
}

func (v CurrentErrorReaction) String() string {
	if s, ok := EnumNamesCurrentErrorReaction[v]; ok {
		return s
	}
	return "CurrentErrorReaction(" + strconv.FormatInt(int64(v), 10) + ")"
}
