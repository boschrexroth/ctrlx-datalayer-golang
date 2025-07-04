// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type CategoryOfAxis int8

const (
	/// drive axis
	CategoryOfAxisDRIVEAXS   CategoryOfAxis = 0
	/// encoder axis
	CategoryOfAxisENCODERAXS CategoryOfAxis = 1
	/// control to control
	CategoryOfAxisC2C        CategoryOfAxis = 2
	/// mover axs (related to the ctrlX FLOW products)
	CategoryOfAxisMOVERAXS   CategoryOfAxis = 3
)

var EnumNamesCategoryOfAxis = map[CategoryOfAxis]string{
	CategoryOfAxisDRIVEAXS:   "DRIVEAXS",
	CategoryOfAxisENCODERAXS: "ENCODERAXS",
	CategoryOfAxisC2C:        "C2C",
	CategoryOfAxisMOVERAXS:   "MOVERAXS",
}

var EnumValuesCategoryOfAxis = map[string]CategoryOfAxis{
	"DRIVEAXS":   CategoryOfAxisDRIVEAXS,
	"ENCODERAXS": CategoryOfAxisENCODERAXS,
	"C2C":        CategoryOfAxisC2C,
	"MOVERAXS":   CategoryOfAxisMOVERAXS,
}

func (v CategoryOfAxis) String() string {
	if s, ok := EnumNamesCategoryOfAxis[v]; ok {
		return s
	}
	return "CategoryOfAxis(" + strconv.FormatInt(int64(v), 10) + ")"
}
