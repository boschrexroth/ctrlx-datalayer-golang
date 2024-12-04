// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package enumeration

import "strconv"

type OverrideValueHandling int32

const (
	OverrideValueHandlingDisabled        OverrideValueHandling = 0
	OverrideValueHandlingLastUsableValue OverrideValueHandling = 1
	OverrideValueHandlingOverrideValue   OverrideValueHandling = 2
)

var EnumNamesOverrideValueHandling = map[OverrideValueHandling]string{
	OverrideValueHandlingDisabled:        "Disabled",
	OverrideValueHandlingLastUsableValue: "LastUsableValue",
	OverrideValueHandlingOverrideValue:   "OverrideValue",
}

var EnumValuesOverrideValueHandling = map[string]OverrideValueHandling{
	"Disabled":        OverrideValueHandlingDisabled,
	"LastUsableValue": OverrideValueHandlingLastUsableValue,
	"OverrideValue":   OverrideValueHandlingOverrideValue,
}

func (v OverrideValueHandling) String() string {
	if s, ok := EnumNamesOverrideValueHandling[v]; ok {
		return s
	}
	return "OverrideValueHandling(" + strconv.FormatInt(int64(v), 10) + ")"
}
