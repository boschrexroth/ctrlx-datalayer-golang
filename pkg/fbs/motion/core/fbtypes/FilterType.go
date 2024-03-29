// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// definition of filter type on encoder value
type FilterType int8

const (
	/// without any filter, using raw value
	FilterTypeNONE             FilterType = 0
	/// 1st order low pass filter
	FilterTypeLOW_PASS_ORDER_1 FilterType = 1
)

var EnumNamesFilterType = map[FilterType]string{
	FilterTypeNONE:             "NONE",
	FilterTypeLOW_PASS_ORDER_1: "LOW_PASS_ORDER_1",
}

var EnumValuesFilterType = map[string]FilterType{
	"NONE":             FilterTypeNONE,
	"LOW_PASS_ORDER_1": FilterTypeLOW_PASS_ORDER_1,
}

func (v FilterType) String() string {
	if s, ok := EnumNamesFilterType[v]; ok {
		return s
	}
	return "FilterType(" + strconv.FormatInt(int64(v), 10) + ")"
}
