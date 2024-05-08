// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type MotionComponentType int8

const (
	/// motion.core component
	MotionComponentTypeMOTION_CORE             MotionComponentType = 0
	/// a sub library of motion.core
	MotionComponentTypeCORE_LIB                MotionComponentType = 1
	/// an motion extension bundle, that is part of the app.motion
	MotionComponentTypeEXTENSION_BUNDLE_MOTION MotionComponentType = 2
	/// an motion extension bundle, that was built with the MExA SDK, but its not part of app.motion
	MotionComponentTypeEXTENSION_BUNDLE_SDK    MotionComponentType = 3
	/// other bundle, that is part of app.motion
	MotionComponentTypeOTHER_BUNDLE            MotionComponentType = 4
)

var EnumNamesMotionComponentType = map[MotionComponentType]string{
	MotionComponentTypeMOTION_CORE:             "MOTION_CORE",
	MotionComponentTypeCORE_LIB:                "CORE_LIB",
	MotionComponentTypeEXTENSION_BUNDLE_MOTION: "EXTENSION_BUNDLE_MOTION",
	MotionComponentTypeEXTENSION_BUNDLE_SDK:    "EXTENSION_BUNDLE_SDK",
	MotionComponentTypeOTHER_BUNDLE:            "OTHER_BUNDLE",
}

var EnumValuesMotionComponentType = map[string]MotionComponentType{
	"MOTION_CORE":             MotionComponentTypeMOTION_CORE,
	"CORE_LIB":                MotionComponentTypeCORE_LIB,
	"EXTENSION_BUNDLE_MOTION": MotionComponentTypeEXTENSION_BUNDLE_MOTION,
	"EXTENSION_BUNDLE_SDK":    MotionComponentTypeEXTENSION_BUNDLE_SDK,
	"OTHER_BUNDLE":            MotionComponentTypeOTHER_BUNDLE,
}

func (v MotionComponentType) String() string {
	if s, ok := EnumNamesMotionComponentType[v]; ok {
		return s
	}
	return "MotionComponentType(" + strconv.FormatInt(int64(v), 10) + ")"
}