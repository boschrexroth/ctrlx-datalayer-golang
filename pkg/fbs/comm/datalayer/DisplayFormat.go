// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package datalayer

import "strconv"

type DisplayFormat int8

const (
	DisplayFormatAuto DisplayFormat = 0
	DisplayFormatBin  DisplayFormat = 1
	DisplayFormatOct  DisplayFormat = 2
	DisplayFormatDec  DisplayFormat = 3
	DisplayFormatHex  DisplayFormat = 4
)

var EnumNamesDisplayFormat = map[DisplayFormat]string{
	DisplayFormatAuto: "Auto",
	DisplayFormatBin:  "Bin",
	DisplayFormatOct:  "Oct",
	DisplayFormatDec:  "Dec",
	DisplayFormatHex:  "Hex",
}

var EnumValuesDisplayFormat = map[string]DisplayFormat{
	"Auto": DisplayFormatAuto,
	"Bin":  DisplayFormatBin,
	"Oct":  DisplayFormatOct,
	"Dec":  DisplayFormatDec,
	"Hex":  DisplayFormatHex,
}

func (v DisplayFormat) String() string {
	if s, ok := EnumNamesDisplayFormat[v]; ok {
		return s
	}
	return "DisplayFormat(" + strconv.FormatInt(int64(v), 10) + ")"
}
