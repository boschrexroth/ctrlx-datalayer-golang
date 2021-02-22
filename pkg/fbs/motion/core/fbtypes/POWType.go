// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type POWType int8

const (
	POWTypeBASE        POWType = 1
	POWTypeORIENTATION POWType = 2
	POWTypeADDITIONAL  POWType = 3
)

var EnumNamesPOWType = map[POWType]string{
	POWTypeBASE:        "BASE",
	POWTypeORIENTATION: "ORIENTATION",
	POWTypeADDITIONAL:  "ADDITIONAL",
}

var EnumValuesPOWType = map[string]POWType{
	"BASE":        POWTypeBASE,
	"ORIENTATION": POWTypeORIENTATION,
	"ADDITIONAL":  POWTypeADDITIONAL,
}

func (v POWType) String() string {
	if s, ok := EnumNamesPOWType[v]; ok {
		return s
	}
	return "POWType(" + strconv.FormatInt(int64(v), 10) + ")"
}
