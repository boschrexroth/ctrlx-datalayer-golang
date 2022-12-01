// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import "strconv"

type Area uint32

const (
	AreaMARKER Area = 0
	AreaINPUT  Area = 1
	AreaOUTPUT Area = 2
)

var EnumNamesArea = map[Area]string{
	AreaMARKER: "MARKER",
	AreaINPUT:  "INPUT",
	AreaOUTPUT: "OUTPUT",
}

var EnumValuesArea = map[string]Area{
	"MARKER": AreaMARKER,
	"INPUT":  AreaINPUT,
	"OUTPUT": AreaOUTPUT,
}

func (v Area) String() string {
	if s, ok := EnumNamesArea[v]; ok {
		return s
	}
	return "Area(" + strconv.FormatInt(int64(v), 10) + ")"
}