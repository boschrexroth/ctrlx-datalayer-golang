// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import "strconv"

type AccessType int8

const (
	AccessTypeUnknown      AccessType = 0
	AccessTypeRaw          AccessType = 1
	AccessTypeTripleBuffer AccessType = 2
	AccessTypeNTelBuf      AccessType = 3
)

var EnumNamesAccessType = map[AccessType]string{
	AccessTypeUnknown:      "Unknown",
	AccessTypeRaw:          "Raw",
	AccessTypeTripleBuffer: "TripleBuffer",
	AccessTypeNTelBuf:      "NTelBuf",
}

var EnumValuesAccessType = map[string]AccessType{
	"Unknown":      AccessTypeUnknown,
	"Raw":          AccessTypeRaw,
	"TripleBuffer": AccessTypeTripleBuffer,
	"NTelBuf":      AccessTypeNTelBuf,
}

func (v AccessType) String() string {
	if s, ok := EnumNamesAccessType[v]; ok {
		return s
	}
	return "AccessType(" + strconv.FormatInt(int64(v), 10) + ")"
}
