// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// Definition of diagnosis type, when SlaveCantFollowErrorReaction is TRY_TO_RESYNC or KEEP_SYNCHRONIZED
type SlaveCantFollowMessageType int8

const (
	/// don't create any diagnosis
	SlaveCantFollowMessageTypeNONE    SlaveCantFollowMessageType = 0
	/// create diagnosis message
	SlaveCantFollowMessageTypeMESSAGE SlaveCantFollowMessageType = 1
	/// create diagnosis warning
	SlaveCantFollowMessageTypeWARNING SlaveCantFollowMessageType = 2
)

var EnumNamesSlaveCantFollowMessageType = map[SlaveCantFollowMessageType]string{
	SlaveCantFollowMessageTypeNONE:    "NONE",
	SlaveCantFollowMessageTypeMESSAGE: "MESSAGE",
	SlaveCantFollowMessageTypeWARNING: "WARNING",
}

var EnumValuesSlaveCantFollowMessageType = map[string]SlaveCantFollowMessageType{
	"NONE":    SlaveCantFollowMessageTypeNONE,
	"MESSAGE": SlaveCantFollowMessageTypeMESSAGE,
	"WARNING": SlaveCantFollowMessageTypeWARNING,
}

func (v SlaveCantFollowMessageType) String() string {
	if s, ok := EnumNamesSlaveCantFollowMessageType[v]; ok {
		return s
	}
	return "SlaveCantFollowMessageType(" + strconv.FormatInt(int64(v), 10) + ")"
}
