// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package enumeration

import "strconv"

type UadpDataSetMessageContentMask uint32

const (
	UadpDataSetMessageContentMaskTimestamp      UadpDataSetMessageContentMask = 1
	UadpDataSetMessageContentMaskPicoSeconds    UadpDataSetMessageContentMask = 2
	UadpDataSetMessageContentMaskStatus         UadpDataSetMessageContentMask = 4
	UadpDataSetMessageContentMaskMajorVersion   UadpDataSetMessageContentMask = 8
	UadpDataSetMessageContentMaskMinorVersion   UadpDataSetMessageContentMask = 16
	UadpDataSetMessageContentMaskSequenceNumber UadpDataSetMessageContentMask = 32
)

var EnumNamesUadpDataSetMessageContentMask = map[UadpDataSetMessageContentMask]string{
	UadpDataSetMessageContentMaskTimestamp:      "Timestamp",
	UadpDataSetMessageContentMaskPicoSeconds:    "PicoSeconds",
	UadpDataSetMessageContentMaskStatus:         "Status",
	UadpDataSetMessageContentMaskMajorVersion:   "MajorVersion",
	UadpDataSetMessageContentMaskMinorVersion:   "MinorVersion",
	UadpDataSetMessageContentMaskSequenceNumber: "SequenceNumber",
}

var EnumValuesUadpDataSetMessageContentMask = map[string]UadpDataSetMessageContentMask{
	"Timestamp":      UadpDataSetMessageContentMaskTimestamp,
	"PicoSeconds":    UadpDataSetMessageContentMaskPicoSeconds,
	"Status":         UadpDataSetMessageContentMaskStatus,
	"MajorVersion":   UadpDataSetMessageContentMaskMajorVersion,
	"MinorVersion":   UadpDataSetMessageContentMaskMinorVersion,
	"SequenceNumber": UadpDataSetMessageContentMaskSequenceNumber,
}

func (v UadpDataSetMessageContentMask) String() string {
	if s, ok := EnumNamesUadpDataSetMessageContentMask[v]; ok {
		return s
	}
	return "UadpDataSetMessageContentMask(" + strconv.FormatInt(int64(v), 10) + ")"
}