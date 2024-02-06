// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// type of probe index
type ProbeIndexType int8

const (
	///< probe1 channel
	ProbeIndexTypePROBE1 ProbeIndexType = 0
	///< probe2 channel
	ProbeIndexTypePROBE2 ProbeIndexType = 1
)

var EnumNamesProbeIndexType = map[ProbeIndexType]string{
	ProbeIndexTypePROBE1: "PROBE1",
	ProbeIndexTypePROBE2: "PROBE2",
}

var EnumValuesProbeIndexType = map[string]ProbeIndexType{
	"PROBE1": ProbeIndexTypePROBE1,
	"PROBE2": ProbeIndexTypePROBE2,
}

func (v ProbeIndexType) String() string {
	if s, ok := EnumNamesProbeIndexType[v]; ok {
		return s
	}
	return "ProbeIndexType(" + strconv.FormatInt(int64(v), 10) + ")"
}
