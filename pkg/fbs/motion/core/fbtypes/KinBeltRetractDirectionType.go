// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

/// definition of the direction of the retract motion in case of an error in the kinematics function 'belt synchronization'
type KinBeltRetractDirectionType int8

const (
	/// retract to X-plane in positive direction
	KinBeltRetractDirectionTypeX_POSITIVE KinBeltRetractDirectionType = 0
	/// retract to X-plane in negative direction
	KinBeltRetractDirectionTypeX_NEGATIVE KinBeltRetractDirectionType = 1
	/// retract to Y-plane in positive direction
	KinBeltRetractDirectionTypeY_POSITIVE KinBeltRetractDirectionType = 2
	/// retract to Y-plane in negative direction
	KinBeltRetractDirectionTypeY_NEGATIVE KinBeltRetractDirectionType = 3
	/// retract to Z-plane in positive direction
	KinBeltRetractDirectionTypeZ_POSITIVE KinBeltRetractDirectionType = 4
	/// retract to Z-plane in negative direction
	KinBeltRetractDirectionTypeZ_NEGATIVE KinBeltRetractDirectionType = 5
)

var EnumNamesKinBeltRetractDirectionType = map[KinBeltRetractDirectionType]string{
	KinBeltRetractDirectionTypeX_POSITIVE: "X_POSITIVE",
	KinBeltRetractDirectionTypeX_NEGATIVE: "X_NEGATIVE",
	KinBeltRetractDirectionTypeY_POSITIVE: "Y_POSITIVE",
	KinBeltRetractDirectionTypeY_NEGATIVE: "Y_NEGATIVE",
	KinBeltRetractDirectionTypeZ_POSITIVE: "Z_POSITIVE",
	KinBeltRetractDirectionTypeZ_NEGATIVE: "Z_NEGATIVE",
}

var EnumValuesKinBeltRetractDirectionType = map[string]KinBeltRetractDirectionType{
	"X_POSITIVE": KinBeltRetractDirectionTypeX_POSITIVE,
	"X_NEGATIVE": KinBeltRetractDirectionTypeX_NEGATIVE,
	"Y_POSITIVE": KinBeltRetractDirectionTypeY_POSITIVE,
	"Y_NEGATIVE": KinBeltRetractDirectionTypeY_NEGATIVE,
	"Z_POSITIVE": KinBeltRetractDirectionTypeZ_POSITIVE,
	"Z_NEGATIVE": KinBeltRetractDirectionTypeZ_NEGATIVE,
}

func (v KinBeltRetractDirectionType) String() string {
	if s, ok := EnumNamesKinBeltRetractDirectionType[v]; ok {
		return s
	}
	return "KinBeltRetractDirectionType(" + strconv.FormatInt(int64(v), 10) + ")"
}
