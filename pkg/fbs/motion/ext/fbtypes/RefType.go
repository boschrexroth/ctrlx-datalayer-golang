// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type RefType int8

const (
	RefTypeSTOPPED_POS RefType = 0
	RefTypeMARKER_POS  RefType = 1
)

var EnumNamesRefType = map[RefType]string{
	RefTypeSTOPPED_POS: "STOPPED_POS",
	RefTypeMARKER_POS:  "MARKER_POS",
}

var EnumValuesRefType = map[string]RefType{
	"STOPPED_POS": RefTypeSTOPPED_POS,
	"MARKER_POS":  RefTypeMARKER_POS,
}

func (v RefType) String() string {
	if s, ok := EnumNamesRefType[v]; ok {
		return s
	}
	return "RefType(" + strconv.FormatInt(int64(v), 10) + ")"
}