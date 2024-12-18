// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// type of the ACS input
type AcsInputType int8

const (
	/// the interpolated axis positions are used
	AcsInputTypeINTERPOLATED_AXIS_POS AcsInputType = 0
	/// the actual axis positions are used
	AcsInputTypeACTUAL_AXIS_POS       AcsInputType = 1
	/// the given axis positions are used
	AcsInputTypeGIVEN_AXIS_POS        AcsInputType = 2
)

var EnumNamesAcsInputType = map[AcsInputType]string{
	AcsInputTypeINTERPOLATED_AXIS_POS: "INTERPOLATED_AXIS_POS",
	AcsInputTypeACTUAL_AXIS_POS:       "ACTUAL_AXIS_POS",
	AcsInputTypeGIVEN_AXIS_POS:        "GIVEN_AXIS_POS",
}

var EnumValuesAcsInputType = map[string]AcsInputType{
	"INTERPOLATED_AXIS_POS": AcsInputTypeINTERPOLATED_AXIS_POS,
	"ACTUAL_AXIS_POS":       AcsInputTypeACTUAL_AXIS_POS,
	"GIVEN_AXIS_POS":        AcsInputTypeGIVEN_AXIS_POS,
}

func (v AcsInputType) String() string {
	if s, ok := EnumNamesAcsInputType[v]; ok {
		return s
	}
	return "AcsInputType(" + strconv.FormatInt(int64(v), 10) + ")"
}
