// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type OscilloscopeStateEnum int8

const (
	OscilloscopeStateEnumNOT_CONFIGURED      OscilloscopeStateEnum = 0
	OscilloscopeStateEnumCONFIGURED          OscilloscopeStateEnum = 1
	OscilloscopeStateEnumSTARTING            OscilloscopeStateEnum = 2
	OscilloscopeStateEnumWAIT_FOR_TRIGGER    OscilloscopeStateEnum = 3
	OscilloscopeStateEnumTRIGGERED           OscilloscopeStateEnum = 4
	OscilloscopeStateEnumRECORDING_COMPLETED OscilloscopeStateEnum = 5
	OscilloscopeStateEnumRECORDING           OscilloscopeStateEnum = 6
	OscilloscopeStateEnumERROR               OscilloscopeStateEnum = 7
)

var EnumNamesOscilloscopeStateEnum = map[OscilloscopeStateEnum]string{
	OscilloscopeStateEnumNOT_CONFIGURED:      "NOT_CONFIGURED",
	OscilloscopeStateEnumCONFIGURED:          "CONFIGURED",
	OscilloscopeStateEnumSTARTING:            "STARTING",
	OscilloscopeStateEnumWAIT_FOR_TRIGGER:    "WAIT_FOR_TRIGGER",
	OscilloscopeStateEnumTRIGGERED:           "TRIGGERED",
	OscilloscopeStateEnumRECORDING_COMPLETED: "RECORDING_COMPLETED",
	OscilloscopeStateEnumRECORDING:           "RECORDING",
	OscilloscopeStateEnumERROR:               "ERROR",
}

var EnumValuesOscilloscopeStateEnum = map[string]OscilloscopeStateEnum{
	"NOT_CONFIGURED":      OscilloscopeStateEnumNOT_CONFIGURED,
	"CONFIGURED":          OscilloscopeStateEnumCONFIGURED,
	"STARTING":            OscilloscopeStateEnumSTARTING,
	"WAIT_FOR_TRIGGER":    OscilloscopeStateEnumWAIT_FOR_TRIGGER,
	"TRIGGERED":           OscilloscopeStateEnumTRIGGERED,
	"RECORDING_COMPLETED": OscilloscopeStateEnumRECORDING_COMPLETED,
	"RECORDING":           OscilloscopeStateEnumRECORDING,
	"ERROR":               OscilloscopeStateEnumERROR,
}

func (v OscilloscopeStateEnum) String() string {
	if s, ok := EnumNamesOscilloscopeStateEnum[v]; ok {
		return s
	}
	return "OscilloscopeStateEnum(" + strconv.FormatInt(int64(v), 10) + ")"
}
