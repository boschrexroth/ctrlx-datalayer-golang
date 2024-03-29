// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// type of the law
type SegmentLawType uint32

const (
	/// Rest in rest, inclined sine curve
	SegmentLawTypeREST_IN_REST_INCLINED_SINE         SegmentLawType = 0
	/// Rest in rest, polynomial 5th order
	SegmentLawTypeREST_IN_REST_POLYNOMIAL_5          SegmentLawType = 1
	/// Rest in velocity, polynomial 5th order
	SegmentLawTypeREST_IN_VELOCITY_POLYNOMIAL_5      SegmentLawType = 2
	/// Rest in velocity, polynomial 7th order
	SegmentLawTypeREST_IN_VELOCITY_POLYNOMIAL_7      SegmentLawType = 3
	/// Velocity in rest, polynomial 5th order
	SegmentLawTypeVELOCITY_IN_REST_POLYNOMIAL_5      SegmentLawType = 4
	/// Velocity in rest, polynomial 7th order
	SegmentLawTypeVELOCITY_IN_REST_POLYNOMIAL_7      SegmentLawType = 5
	/// Velocity in velocity, constant velocity
	SegmentLawTypeVELOCITY_IN_VELOCITY_LINEAR        SegmentLawType = 6
	/// Velocity in velocity, polynomial 5th order
	SegmentLawTypeVELOCITY_IN_VELOCITY_POLYNOMIAL_5  SegmentLawType = 7
	/// Rest in rest, standstill
	SegmentLawTypeREST_IN_REST_LINEAR                SegmentLawType = 8
	/// Rest in rest, polynomial 7th order
	SegmentLawTypeREST_IN_REST_POLYNOMIAL_7          SegmentLawType = 9
	/// Rest in rest, sine curve
	SegmentLawTypeREST_IN_REST_SINE                  SegmentLawType = 10
	/// Rest in rest, sinusoid of Gutman
	SegmentLawTypeREST_IN_REST_SINE_GUTMAN           SegmentLawType = 11
	/// Rest in rest, acceleration-optimal inclined sine curve
	SegmentLawTypeREST_IN_REST_SINE_ACC              SegmentLawType = 12
	/// Rest in rest, torque-inclined sine curve
	SegmentLawTypeREST_IN_REST_SINE_TORQUE           SegmentLawType = 13
	/// Rest in rest, modified acceleration trapezoid
	SegmentLawTypeREST_IN_REST_MOD_TRAPEZ            SegmentLawType = 14
	/// Rest in rest, modified sine curve
	SegmentLawTypeREST_IN_REST_MOD_SINE              SegmentLawType = 15
	/// Velocity in velocity, polynomial 7th order
	SegmentLawTypeVELOCITY_IN_VELOCITY_POLYNOMIAL_7  SegmentLawType = 16
	/// Velocity in velocity, modified sine curve
	SegmentLawTypeVELOCITY_IN_VELOCITY_MOD_SINE      SegmentLawType = 17
	/// Rest in rest, velocity-limited polynomial 5th order
	SegmentLawTypeREST_IN_REST_POLYNOMIAL_5_VLIM     SegmentLawType = 18
	/// Rest in rest, quadratic parabola
	SegmentLawTypeREST_IN_REST_PARABOLA              SegmentLawType = 19
	/// Rest in rest, polynomial 8th order
	SegmentLawTypeREST_IN_REST_POLYNOMIAL_8          SegmentLawType = 20
	/// Motion in motion, polynomial 5th order
	SegmentLawTypeMOTION_IN_MOTION_POLYNOMIAL_5      SegmentLawType = 21
	/// Motion in motion, polynomial 7th order
	SegmentLawTypeMOTION_IN_MOTION_POLYNOMIAL_7      SegmentLawType = 22
	/// Polynomial 5th order
	SegmentLawTypeCOMMON_POLYNOMIAL_5                SegmentLawType = 23
	/// Polynomial 7th order
	SegmentLawTypeCOMMON_POLYNOMIAL_7                SegmentLawType = 24
	/// Polynomial 2nd order
	SegmentLawTypeCOMMON_POLYNOMIAL_2                SegmentLawType = 25
	/// Polynomial 3rd order
	SegmentLawTypeCOMMON_POLYNOMIAL_3                SegmentLawType = 26
	/// Polynomial 4th order
	SegmentLawTypeCOMMON_POLYNOMIAL_4                SegmentLawType = 27
	/// Polynomial 8th order
	SegmentLawTypeCOMMON_POLYNOMIAL_8                SegmentLawType = 28
	/// Velocity in velocity, Acceleration-limited (trapezoid profile)
	SegmentLawTypeVELOCITY_IN_VELOCITY_TRAPEZ_ALIM   SegmentLawType = 29
	/// Motion in motion, Velocity-limited polynomial 5th order
	SegmentLawTypeMOTION_IN_MOTION_POLYNOMIAL_5_VLIM SegmentLawType = 30
	/// Motion in motion, Free of harmonics polynomial 5th order
	SegmentLawTypeMOTION_IN_MOTION_POLYNOMIAL_5_SLIM SegmentLawType = 31
	/// Acceleration-limited motion (trapezoid profile)
	SegmentLawTypeFIT_VEL_TRAPEZE_ALIM               SegmentLawType = 32
	/// Acceleration-limited motion (sinusoid profile)
	SegmentLawTypeFIT_SINE_TRAPEZE_ALIM              SegmentLawType = 33
	/// Jerk-limited motion
	SegmentLawTypeFIT_ACC_TRAPEZE_JLIM               SegmentLawType = 34
	/// Linear Acceleration
	SegmentLawTypeMOTION_IN_MOTION_ACAM              SegmentLawType = 35
	/// Linear Velocity
	SegmentLawTypeMOTION_IN_MOTION_VCAM              SegmentLawType = 36
	/// Velocity 2nd order (Start acceleration zero)
	SegmentLawTypeMOTION_IN_MOTION_VCAM2_A           SegmentLawType = 37
	/// Velocity 2nd order (End acceleration zero)
	SegmentLawTypeMOTION_IN_MOTION_VCAM2_B           SegmentLawType = 38
	/// Points-table with interpolation
	SegmentLawTypePOINTS_TABLE_AUTO                  SegmentLawType = 39
)

var EnumNamesSegmentLawType = map[SegmentLawType]string{
	SegmentLawTypeREST_IN_REST_INCLINED_SINE:         "REST_IN_REST_INCLINED_SINE",
	SegmentLawTypeREST_IN_REST_POLYNOMIAL_5:          "REST_IN_REST_POLYNOMIAL_5",
	SegmentLawTypeREST_IN_VELOCITY_POLYNOMIAL_5:      "REST_IN_VELOCITY_POLYNOMIAL_5",
	SegmentLawTypeREST_IN_VELOCITY_POLYNOMIAL_7:      "REST_IN_VELOCITY_POLYNOMIAL_7",
	SegmentLawTypeVELOCITY_IN_REST_POLYNOMIAL_5:      "VELOCITY_IN_REST_POLYNOMIAL_5",
	SegmentLawTypeVELOCITY_IN_REST_POLYNOMIAL_7:      "VELOCITY_IN_REST_POLYNOMIAL_7",
	SegmentLawTypeVELOCITY_IN_VELOCITY_LINEAR:        "VELOCITY_IN_VELOCITY_LINEAR",
	SegmentLawTypeVELOCITY_IN_VELOCITY_POLYNOMIAL_5:  "VELOCITY_IN_VELOCITY_POLYNOMIAL_5",
	SegmentLawTypeREST_IN_REST_LINEAR:                "REST_IN_REST_LINEAR",
	SegmentLawTypeREST_IN_REST_POLYNOMIAL_7:          "REST_IN_REST_POLYNOMIAL_7",
	SegmentLawTypeREST_IN_REST_SINE:                  "REST_IN_REST_SINE",
	SegmentLawTypeREST_IN_REST_SINE_GUTMAN:           "REST_IN_REST_SINE_GUTMAN",
	SegmentLawTypeREST_IN_REST_SINE_ACC:              "REST_IN_REST_SINE_ACC",
	SegmentLawTypeREST_IN_REST_SINE_TORQUE:           "REST_IN_REST_SINE_TORQUE",
	SegmentLawTypeREST_IN_REST_MOD_TRAPEZ:            "REST_IN_REST_MOD_TRAPEZ",
	SegmentLawTypeREST_IN_REST_MOD_SINE:              "REST_IN_REST_MOD_SINE",
	SegmentLawTypeVELOCITY_IN_VELOCITY_POLYNOMIAL_7:  "VELOCITY_IN_VELOCITY_POLYNOMIAL_7",
	SegmentLawTypeVELOCITY_IN_VELOCITY_MOD_SINE:      "VELOCITY_IN_VELOCITY_MOD_SINE",
	SegmentLawTypeREST_IN_REST_POLYNOMIAL_5_VLIM:     "REST_IN_REST_POLYNOMIAL_5_VLIM",
	SegmentLawTypeREST_IN_REST_PARABOLA:              "REST_IN_REST_PARABOLA",
	SegmentLawTypeREST_IN_REST_POLYNOMIAL_8:          "REST_IN_REST_POLYNOMIAL_8",
	SegmentLawTypeMOTION_IN_MOTION_POLYNOMIAL_5:      "MOTION_IN_MOTION_POLYNOMIAL_5",
	SegmentLawTypeMOTION_IN_MOTION_POLYNOMIAL_7:      "MOTION_IN_MOTION_POLYNOMIAL_7",
	SegmentLawTypeCOMMON_POLYNOMIAL_5:                "COMMON_POLYNOMIAL_5",
	SegmentLawTypeCOMMON_POLYNOMIAL_7:                "COMMON_POLYNOMIAL_7",
	SegmentLawTypeCOMMON_POLYNOMIAL_2:                "COMMON_POLYNOMIAL_2",
	SegmentLawTypeCOMMON_POLYNOMIAL_3:                "COMMON_POLYNOMIAL_3",
	SegmentLawTypeCOMMON_POLYNOMIAL_4:                "COMMON_POLYNOMIAL_4",
	SegmentLawTypeCOMMON_POLYNOMIAL_8:                "COMMON_POLYNOMIAL_8",
	SegmentLawTypeVELOCITY_IN_VELOCITY_TRAPEZ_ALIM:   "VELOCITY_IN_VELOCITY_TRAPEZ_ALIM",
	SegmentLawTypeMOTION_IN_MOTION_POLYNOMIAL_5_VLIM: "MOTION_IN_MOTION_POLYNOMIAL_5_VLIM",
	SegmentLawTypeMOTION_IN_MOTION_POLYNOMIAL_5_SLIM: "MOTION_IN_MOTION_POLYNOMIAL_5_SLIM",
	SegmentLawTypeFIT_VEL_TRAPEZE_ALIM:               "FIT_VEL_TRAPEZE_ALIM",
	SegmentLawTypeFIT_SINE_TRAPEZE_ALIM:              "FIT_SINE_TRAPEZE_ALIM",
	SegmentLawTypeFIT_ACC_TRAPEZE_JLIM:               "FIT_ACC_TRAPEZE_JLIM",
	SegmentLawTypeMOTION_IN_MOTION_ACAM:              "MOTION_IN_MOTION_ACAM",
	SegmentLawTypeMOTION_IN_MOTION_VCAM:              "MOTION_IN_MOTION_VCAM",
	SegmentLawTypeMOTION_IN_MOTION_VCAM2_A:           "MOTION_IN_MOTION_VCAM2_A",
	SegmentLawTypeMOTION_IN_MOTION_VCAM2_B:           "MOTION_IN_MOTION_VCAM2_B",
	SegmentLawTypePOINTS_TABLE_AUTO:                  "POINTS_TABLE_AUTO",
}

var EnumValuesSegmentLawType = map[string]SegmentLawType{
	"REST_IN_REST_INCLINED_SINE":         SegmentLawTypeREST_IN_REST_INCLINED_SINE,
	"REST_IN_REST_POLYNOMIAL_5":          SegmentLawTypeREST_IN_REST_POLYNOMIAL_5,
	"REST_IN_VELOCITY_POLYNOMIAL_5":      SegmentLawTypeREST_IN_VELOCITY_POLYNOMIAL_5,
	"REST_IN_VELOCITY_POLYNOMIAL_7":      SegmentLawTypeREST_IN_VELOCITY_POLYNOMIAL_7,
	"VELOCITY_IN_REST_POLYNOMIAL_5":      SegmentLawTypeVELOCITY_IN_REST_POLYNOMIAL_5,
	"VELOCITY_IN_REST_POLYNOMIAL_7":      SegmentLawTypeVELOCITY_IN_REST_POLYNOMIAL_7,
	"VELOCITY_IN_VELOCITY_LINEAR":        SegmentLawTypeVELOCITY_IN_VELOCITY_LINEAR,
	"VELOCITY_IN_VELOCITY_POLYNOMIAL_5":  SegmentLawTypeVELOCITY_IN_VELOCITY_POLYNOMIAL_5,
	"REST_IN_REST_LINEAR":                SegmentLawTypeREST_IN_REST_LINEAR,
	"REST_IN_REST_POLYNOMIAL_7":          SegmentLawTypeREST_IN_REST_POLYNOMIAL_7,
	"REST_IN_REST_SINE":                  SegmentLawTypeREST_IN_REST_SINE,
	"REST_IN_REST_SINE_GUTMAN":           SegmentLawTypeREST_IN_REST_SINE_GUTMAN,
	"REST_IN_REST_SINE_ACC":              SegmentLawTypeREST_IN_REST_SINE_ACC,
	"REST_IN_REST_SINE_TORQUE":           SegmentLawTypeREST_IN_REST_SINE_TORQUE,
	"REST_IN_REST_MOD_TRAPEZ":            SegmentLawTypeREST_IN_REST_MOD_TRAPEZ,
	"REST_IN_REST_MOD_SINE":              SegmentLawTypeREST_IN_REST_MOD_SINE,
	"VELOCITY_IN_VELOCITY_POLYNOMIAL_7":  SegmentLawTypeVELOCITY_IN_VELOCITY_POLYNOMIAL_7,
	"VELOCITY_IN_VELOCITY_MOD_SINE":      SegmentLawTypeVELOCITY_IN_VELOCITY_MOD_SINE,
	"REST_IN_REST_POLYNOMIAL_5_VLIM":     SegmentLawTypeREST_IN_REST_POLYNOMIAL_5_VLIM,
	"REST_IN_REST_PARABOLA":              SegmentLawTypeREST_IN_REST_PARABOLA,
	"REST_IN_REST_POLYNOMIAL_8":          SegmentLawTypeREST_IN_REST_POLYNOMIAL_8,
	"MOTION_IN_MOTION_POLYNOMIAL_5":      SegmentLawTypeMOTION_IN_MOTION_POLYNOMIAL_5,
	"MOTION_IN_MOTION_POLYNOMIAL_7":      SegmentLawTypeMOTION_IN_MOTION_POLYNOMIAL_7,
	"COMMON_POLYNOMIAL_5":                SegmentLawTypeCOMMON_POLYNOMIAL_5,
	"COMMON_POLYNOMIAL_7":                SegmentLawTypeCOMMON_POLYNOMIAL_7,
	"COMMON_POLYNOMIAL_2":                SegmentLawTypeCOMMON_POLYNOMIAL_2,
	"COMMON_POLYNOMIAL_3":                SegmentLawTypeCOMMON_POLYNOMIAL_3,
	"COMMON_POLYNOMIAL_4":                SegmentLawTypeCOMMON_POLYNOMIAL_4,
	"COMMON_POLYNOMIAL_8":                SegmentLawTypeCOMMON_POLYNOMIAL_8,
	"VELOCITY_IN_VELOCITY_TRAPEZ_ALIM":   SegmentLawTypeVELOCITY_IN_VELOCITY_TRAPEZ_ALIM,
	"MOTION_IN_MOTION_POLYNOMIAL_5_VLIM": SegmentLawTypeMOTION_IN_MOTION_POLYNOMIAL_5_VLIM,
	"MOTION_IN_MOTION_POLYNOMIAL_5_SLIM": SegmentLawTypeMOTION_IN_MOTION_POLYNOMIAL_5_SLIM,
	"FIT_VEL_TRAPEZE_ALIM":               SegmentLawTypeFIT_VEL_TRAPEZE_ALIM,
	"FIT_SINE_TRAPEZE_ALIM":              SegmentLawTypeFIT_SINE_TRAPEZE_ALIM,
	"FIT_ACC_TRAPEZE_JLIM":               SegmentLawTypeFIT_ACC_TRAPEZE_JLIM,
	"MOTION_IN_MOTION_ACAM":              SegmentLawTypeMOTION_IN_MOTION_ACAM,
	"MOTION_IN_MOTION_VCAM":              SegmentLawTypeMOTION_IN_MOTION_VCAM,
	"MOTION_IN_MOTION_VCAM2_A":           SegmentLawTypeMOTION_IN_MOTION_VCAM2_A,
	"MOTION_IN_MOTION_VCAM2_B":           SegmentLawTypeMOTION_IN_MOTION_VCAM2_B,
	"POINTS_TABLE_AUTO":                  SegmentLawTypePOINTS_TABLE_AUTO,
}

func (v SegmentLawType) String() string {
	if s, ok := EnumNamesSegmentLawType[v]; ok {
		return s
	}
	return "SegmentLawType(" + strconv.FormatInt(int64(v), 10) + ")"
}
