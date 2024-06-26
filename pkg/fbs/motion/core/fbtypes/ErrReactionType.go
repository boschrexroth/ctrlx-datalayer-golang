// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// Error reaction type
type ErrReactionType uint32

const (
	/// no reaction but single trace message
	ErrReactionTypeNONE         ErrReactionType = 1
	/// single axis message
	ErrReactionTypeMESSAGE      ErrReactionType = 2
	/// single axis warning
	ErrReactionTypeWARNING      ErrReactionType = 3
	/// axis goes to error state with severe error
	ErrReactionTypeSEVERE_ERROR ErrReactionType = 4
)

var EnumNamesErrReactionType = map[ErrReactionType]string{
	ErrReactionTypeNONE:         "NONE",
	ErrReactionTypeMESSAGE:      "MESSAGE",
	ErrReactionTypeWARNING:      "WARNING",
	ErrReactionTypeSEVERE_ERROR: "SEVERE_ERROR",
}

var EnumValuesErrReactionType = map[string]ErrReactionType{
	"NONE":         ErrReactionTypeNONE,
	"MESSAGE":      ErrReactionTypeMESSAGE,
	"WARNING":      ErrReactionTypeWARNING,
	"SEVERE_ERROR": ErrReactionTypeSEVERE_ERROR,
}

func (v ErrReactionType) String() string {
	if s, ok := EnumNamesErrReactionType[v]; ok {
		return s
	}
	return "ErrReactionType(" + strconv.FormatInt(int64(v), 10) + ")"
}
