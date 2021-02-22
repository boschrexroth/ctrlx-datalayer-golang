// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbtypes

import "strconv"

type ProfileType uint32

const (
	ProfileTypeUnknown                  ProfileType = 0
	ProfileTypeServodrive_over_Ethercat ProfileType = 1
)

var EnumNamesProfileType = map[ProfileType]string{
	ProfileTypeUnknown:                  "Unknown",
	ProfileTypeServodrive_over_Ethercat: "Servodrive_over_Ethercat",
}

var EnumValuesProfileType = map[string]ProfileType{
	"Unknown":                  ProfileTypeUnknown,
	"Servodrive_over_Ethercat": ProfileTypeServodrive_over_Ethercat,
}

func (v ProfileType) String() string {
	if s, ok := EnumNamesProfileType[v]; ok {
		return s
	}
	return "ProfileType(" + strconv.FormatInt(int64(v), 10) + ")"
}
