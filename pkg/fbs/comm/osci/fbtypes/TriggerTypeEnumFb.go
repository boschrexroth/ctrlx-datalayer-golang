// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type TriggerTypeEnumFb int8

const (
	TriggerTypeEnumFbManual        TriggerTypeEnumFb = 0
	TriggerTypeEnumFbRisingEdge    TriggerTypeEnumFb = 1
	TriggerTypeEnumFbFallingEdge   TriggerTypeEnumFb = 2
	TriggerTypeEnumFbLevel         TriggerTypeEnumFb = 3
	TriggerTypeEnumFbRisingLevel   TriggerTypeEnumFb = 4
	TriggerTypeEnumFbFallingLevel  TriggerTypeEnumFb = 5
	TriggerTypeEnumFbRisingFalling TriggerTypeEnumFb = 6
)

var EnumNamesTriggerTypeEnumFb = map[TriggerTypeEnumFb]string{
	TriggerTypeEnumFbManual:        "Manual",
	TriggerTypeEnumFbRisingEdge:    "RisingEdge",
	TriggerTypeEnumFbFallingEdge:   "FallingEdge",
	TriggerTypeEnumFbLevel:         "Level",
	TriggerTypeEnumFbRisingLevel:   "RisingLevel",
	TriggerTypeEnumFbFallingLevel:  "FallingLevel",
	TriggerTypeEnumFbRisingFalling: "RisingFalling",
}

var EnumValuesTriggerTypeEnumFb = map[string]TriggerTypeEnumFb{
	"Manual":        TriggerTypeEnumFbManual,
	"RisingEdge":    TriggerTypeEnumFbRisingEdge,
	"FallingEdge":   TriggerTypeEnumFbFallingEdge,
	"Level":         TriggerTypeEnumFbLevel,
	"RisingLevel":   TriggerTypeEnumFbRisingLevel,
	"FallingLevel":  TriggerTypeEnumFbFallingLevel,
	"RisingFalling": TriggerTypeEnumFbRisingFalling,
}

func (v TriggerTypeEnumFb) String() string {
	if s, ok := EnumNamesTriggerTypeEnumFb[v]; ok {
		return s
	}
	return "TriggerTypeEnumFb(" + strconv.FormatInt(int64(v), 10) + ")"
}